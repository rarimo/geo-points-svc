package config

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	vaultapi "github.com/hashicorp/vault/api"
	accountFactory "github.com/rarimo/geo-points-svc/internal/contracts/extendedaccountfactory"
	pointTokens "github.com/rarimo/geo-points-svc/internal/contracts/points"

	"gitlab.com/distributed_lab/dig"
	"gitlab.com/distributed_lab/figure/v3"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

var ZeroAddress = common.HexToAddress("0x0000000000000000000000000000000000000000")

type Abstraction interface {
	AbstractionConfig() *AbstractionConfig
}

func NewAbstractionConfig(getter kv.Getter) Abstraction {
	return &abstractionConfig{
		getter: getter,
	}
}

type abstractionConfig struct {
	once   comfig.Once
	getter kv.Getter
}

type AbstractionConfig struct {
	RPC            *ethclient.Client
	AccountFactory common.Address
	PointTokens    common.Address
	ChainID        *big.Int
	PointPrice     *big.Int

	privateKey *ecdsa.PrivateKey
}

func (c *abstractionConfig) AbstractionConfig() *AbstractionConfig {
	return c.once.Do(func() interface{} {
		var cfg struct {
			RPC            *ethclient.Client `fig:"rpc,required"`
			AccountFactory common.Address    `fig:"account_factory,required"`
			PointTokens    common.Address    `fig:"point_tokens,required"`
			PointPrice     int64             `fig:"point_price"`

			VaultAddress   string            `fig:"vault_address"`
			VaultMountPath string            `fig:"vault_mount_path"`
			PrivateKey     *ecdsa.PrivateKey `fig:"private_key"`
		}

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "abstraction")).
			With(figure.EthereumHooks, figure.BaseHooks).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out abstraction config: %w", err))
		}

		privateKey := cfg.PrivateKey
		if privateKey == nil {
			privateKey = extractPrivateKey(cfg.VaultAddress, cfg.VaultMountPath)
		}

		chainID, err := cfg.RPC.ChainID(context.TODO())
		if err != nil {
			panic(fmt.Errorf("failed to get chain id: %w", err))
		}

		if cfg.PointPrice == 0 {
			// Default 1 point price == 1 collateral
			cfg.PointPrice = int64(math.Pow10(9))
		}

		return &AbstractionConfig{
			RPC:            cfg.RPC,
			AccountFactory: cfg.AccountFactory,
			PointTokens:    cfg.PointTokens,
			ChainID:        chainID,
			// Collateral tokens has precission 10^18. Point price will be with precission 10^9
			// This mean that 10^9 point price ~ 10^18 collateral
			PointPrice: new(big.Int).Mul(
				big.NewInt(int64(math.Pow10(9))),
				big.NewInt(cfg.PointPrice)),

			privateKey: privateKey,
		}
	}).(*AbstractionConfig)
}

func (r *AbstractionConfig) CreateAccount(ctx context.Context, nullifier [32]byte) (common.Address, error) {
	signerOpts, err := bind.NewKeyedTransactorWithChainID(r.privateKey, r.ChainID)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get keyed transactor: %w", err)
	}

	accountFactoryInstance, err := accountFactory.NewExtendedAccountFactory(r.AccountFactory, r.RPC)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get account factory: %w", err)
	}

	tx, err := accountFactoryInstance.DeployAbstractAccount(signerOpts, nullifier)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to deploy abstraction account: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	rec, err := bind.WaitMined(ctx, r.RPC, tx)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to wait mined transaction: %w", err)
	}

	abi, err := accountFactory.ExtendedAccountFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get contract abi: %w", err)
	}

	abstractionAccountDeployedTopic := abi.Events["AbstractAccountDeployed"].ID

	var event *accountFactory.ExtendedAccountFactoryAbstractAccountDeployed
	for _, log := range rec.Logs {
		if !bytes.Equal(log.Topics[0][:], abstractionAccountDeployedTopic[:]) {
			continue
		}

		event, err = accountFactoryInstance.ParseAbstractAccountDeployed(*log)
		if err != nil {
			return common.Address{}, fmt.Errorf("failed to unpack log: %w", err)
		}
		break
	}

	return event.Account, nil
}

func (r *AbstractionConfig) GetAccount(nullifier [32]byte) (common.Address, error) {
	accountFactoryInstance, err := accountFactory.NewExtendedAccountFactory(r.AccountFactory, r.RPC)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get account factory: %w", err)
	}

	accountAddress, err := accountFactoryInstance.GetAbstractAccount(nil, nullifier)
	if err != nil {
		return common.Address{}, fmt.Errorf("failed to get abstraction account: %w", err)
	}

	return accountAddress, nil
}

func (r *AbstractionConfig) Mint(ctx context.Context, account common.Address, amount *big.Int) (common.Hash, error) {
	signerOpts, err := bind.NewKeyedTransactorWithChainID(r.privateKey, r.ChainID)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get keyed transactor: %w", err)
	}

	pointTokensInstance, err := pointTokens.NewPoints(r.PointTokens, r.RPC)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to get points instance: %w", err)
	}

	tx, err := pointTokensInstance.Mint(signerOpts, account, amount)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to mint points: %w", err)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	_, err = bind.WaitMined(ctx, r.RPC, tx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to wait mined transaction: %w", err)
	}

	return tx.Hash(), nil
}

func extractPrivateKey(vaultAddress, vaultMountPath string) *ecdsa.PrivateKey {
	conf := vaultapi.DefaultConfig()
	conf.Address = vaultAddress

	vaultClient, err := vaultapi.NewClient(conf)
	if err != nil {
		panic(fmt.Errorf("failed to initialize new client: %w", err))
	}

	token := struct {
		Token string `dig:"VAULT_TOKEN,clear"`
	}{}

	err = dig.Out(&token).Now()
	if err != nil {
		panic(fmt.Errorf("failed to dig out token: %w", err))
	}

	vaultClient.SetToken(token.Token)

	secret, err := vaultClient.KVv2(vaultMountPath).Get(context.Background(), "geo-points")
	if err != nil {
		panic(fmt.Errorf("failed to get secret: %w", err))
	}

	vaultRelayerConf := struct {
		PrivateKey *ecdsa.PrivateKey `fig:"private_key,required"`
	}{}

	if err := figure.
		Out(&vaultRelayerConf).
		With(figure.EthereumHooks).
		From(secret.Data).
		Please(); err != nil {
		panic(fmt.Errorf("failed to figure out private_key: %w", err))
	}

	return vaultRelayerConf.PrivateKey
}
