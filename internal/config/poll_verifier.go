package config

import (
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	zkptypes "github.com/iden3/go-rapidsnark/types"
	zkpverifier "github.com/iden3/go-rapidsnark/verifier"
	"github.com/rarimo/geo-points-svc/internal/contracts/proposalsmt"
	"github.com/rarimo/geo-points-svc/internal/contracts/proposalsstate"

	"gitlab.com/distributed_lab/figure"
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/kv"
)

const pollVerificationKey = "./proof_keys/poll.json"

var (
	ErrInvalidProposalEventID = errors.New("proposal event id mismatching")
	ErrInvalidRoot            = errors.New("invalid root")
)

const (
	PollChallengedNullifier = iota
	PollNullifierTreeRoot
	PollParticipationEventID
	PollChallengedEventID
)

type PollVerifierer interface {
	PollVerifier() *PollVerifier
}

func NewPollVerifier(getter kv.Getter) PollVerifierer {
	return &pollVerifier{
		getter: getter,
	}
}

type pollVerifier struct {
	once   comfig.Once
	getter kv.Getter
}

type PollVerifier struct {
	RPC                  *ethclient.Client `fig:"rpc,required"`
	ProposalStateAddress common.Address    `fig:"proposal_state_address,required"`

	proposalsStateCaller *proposalsstate.ProposalsStateCaller
	verificationKey      []byte
}

func (c *pollVerifier) PollVerifier() *PollVerifier {
	return c.once.Do(func() interface{} {

		var cfg PollVerifier

		err := figure.Out(&cfg).
			From(kv.MustGetStringMap(c.getter, "poll_verifier")).
			Please()
		if err != nil {
			panic(fmt.Errorf("failed to figure out vote verifier config: %w", err))
		}

		cfg.verificationKey, err = os.ReadFile(pollVerificationKey)
		if err != nil {
			panic(fmt.Errorf("failed to read pollVerificationKey: %w", err))
		}

		cfg.proposalsStateCaller, err = proposalsstate.NewProposalsStateCaller(cfg.ProposalStateAddress, cfg.RPC)
		if err != nil {
			panic(fmt.Errorf("failed to create proposals state caller: %w", err))
		}

		return &cfg
	}).(*PollVerifier)
}

func (v *PollVerifier) VerifyProof(proof zkptypes.ZKProof, proposalID, proposalEventID *big.Int) error {
	proposalInfo, err := v.proposalsStateCaller.GetProposalInfo(nil, proposalID)
	if err != nil {
		return fmt.Errorf("failed to get proposal info: %w", err)
	}

	proposalEventIDContract, err := v.proposalsStateCaller.GetProposalEventId(nil, proposalID)
	if err != nil {
		return fmt.Errorf("failed to get proposal event id: %w", err)
	}

	if proposalEventIDContract.Cmp(proposalEventID) != 0 {
		return ErrInvalidProposalEventID
	}

	proposalSMTAddress := proposalInfo.ProposalSMT
	proposalSMTCaller, err := proposalsmt.NewProposalSMTCaller(proposalSMTAddress, v.RPC)
	if err != nil {
		return fmt.Errorf("failed to create proposal smt caller: %w", err)
	}

	root, err := proposalSMTCaller.GetRoot(nil)
	if err != nil {
		return fmt.Errorf("failed to get root: %w", err)
	}

	if new(big.Int).SetBytes(root[:]).String() != proof.PubSignals[PollNullifierTreeRoot] {
		return ErrInvalidRoot
	}

	if err = zkpverifier.VerifyGroth16(proof, v.verificationKey); err != nil {
		return fmt.Errorf("failed to verify proof: %w", err)
	}

	return nil
}
