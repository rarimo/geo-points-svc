// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rarimarketaccountfactory

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// RarimarketAccountFactoryMetaData contains all meta data concerning the RarimarketAccountFactory contract.
var RarimarketAccountFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RarimarketAccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"rarimarketAccountImplementation_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"negRiskAdapter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ctfExchange_\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"points_\",\"type\":\"address\"},{\"internalType\":\"contractIERC1155\",\"name\":\"ctf_\",\"type\":\"address\"}],\"name\":\"__RarimarketAccountFactory_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ctf\",\"outputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ctfExchange\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nullifier_\",\"type\":\"bytes32\"}],\"name\":\"deployRarimarketAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nullifier_\",\"type\":\"bytes32\"}],\"name\":\"getRarimarketAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRarimarketAccountImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"negRiskAdapter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"points\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nullifier_\",\"type\":\"bytes32\"}],\"name\":\"predictRarimarketAccountAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"rarimarketAccounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"setRarimarketAccountImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// RarimarketAccountFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use RarimarketAccountFactoryMetaData.ABI instead.
var RarimarketAccountFactoryABI = RarimarketAccountFactoryMetaData.ABI

// RarimarketAccountFactory is an auto generated Go binding around an Ethereum contract.
type RarimarketAccountFactory struct {
	RarimarketAccountFactoryCaller     // Read-only binding to the contract
	RarimarketAccountFactoryTransactor // Write-only binding to the contract
	RarimarketAccountFactoryFilterer   // Log filterer for contract events
}

// RarimarketAccountFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type RarimarketAccountFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RarimarketAccountFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RarimarketAccountFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RarimarketAccountFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RarimarketAccountFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RarimarketAccountFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RarimarketAccountFactorySession struct {
	Contract     *RarimarketAccountFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// RarimarketAccountFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RarimarketAccountFactoryCallerSession struct {
	Contract *RarimarketAccountFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// RarimarketAccountFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RarimarketAccountFactoryTransactorSession struct {
	Contract     *RarimarketAccountFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// RarimarketAccountFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type RarimarketAccountFactoryRaw struct {
	Contract *RarimarketAccountFactory // Generic contract binding to access the raw methods on
}

// RarimarketAccountFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RarimarketAccountFactoryCallerRaw struct {
	Contract *RarimarketAccountFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// RarimarketAccountFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RarimarketAccountFactoryTransactorRaw struct {
	Contract *RarimarketAccountFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRarimarketAccountFactory creates a new instance of RarimarketAccountFactory, bound to a specific deployed contract.
func NewRarimarketAccountFactory(address common.Address, backend bind.ContractBackend) (*RarimarketAccountFactory, error) {
	contract, err := bindRarimarketAccountFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RarimarketAccountFactory{RarimarketAccountFactoryCaller: RarimarketAccountFactoryCaller{contract: contract}, RarimarketAccountFactoryTransactor: RarimarketAccountFactoryTransactor{contract: contract}, RarimarketAccountFactoryFilterer: RarimarketAccountFactoryFilterer{contract: contract}}, nil
}

// NewRarimarketAccountFactoryCaller creates a new read-only instance of RarimarketAccountFactory, bound to a specific deployed contract.
func NewRarimarketAccountFactoryCaller(address common.Address, caller bind.ContractCaller) (*RarimarketAccountFactoryCaller, error) {
	contract, err := bindRarimarketAccountFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RarimarketAccountFactoryCaller{contract: contract}, nil
}

// NewRarimarketAccountFactoryTransactor creates a new write-only instance of RarimarketAccountFactory, bound to a specific deployed contract.
func NewRarimarketAccountFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*RarimarketAccountFactoryTransactor, error) {
	contract, err := bindRarimarketAccountFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RarimarketAccountFactoryTransactor{contract: contract}, nil
}

// NewRarimarketAccountFactoryFilterer creates a new log filterer instance of RarimarketAccountFactory, bound to a specific deployed contract.
func NewRarimarketAccountFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*RarimarketAccountFactoryFilterer, error) {
	contract, err := bindRarimarketAccountFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RarimarketAccountFactoryFilterer{contract: contract}, nil
}

// bindRarimarketAccountFactory binds a generic wrapper to an already deployed contract.
func bindRarimarketAccountFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RarimarketAccountFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RarimarketAccountFactory *RarimarketAccountFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RarimarketAccountFactory.Contract.RarimarketAccountFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RarimarketAccountFactory *RarimarketAccountFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RarimarketAccountFactory.Contract.RarimarketAccountFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RarimarketAccountFactory *RarimarketAccountFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RarimarketAccountFactory.Contract.RarimarketAccountFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RarimarketAccountFactory *RarimarketAccountFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RarimarketAccountFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RarimarketAccountFactory *RarimarketAccountFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RarimarketAccountFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RarimarketAccountFactory *RarimarketAccountFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RarimarketAccountFactory.Contract.contract.Transact(opts, method, params...)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCaller) Ctf(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RarimarketAccountFactory.contract.Call(opts, &out, "ctf")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactorySession) Ctf() (common.Address, error) {
	return _RarimarketAccountFactory.Contract.Ctf(&_RarimarketAccountFactory.CallOpts)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCallerSession) Ctf() (common.Address, error) {
	return _RarimarketAccountFactory.Contract.Ctf(&_RarimarketAccountFactory.CallOpts)
}

// CtfExchange is a free data retrieval call binding the contract method 0x4b8295e7.
//
// Solidity: function ctfExchange() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCaller) CtfExchange(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RarimarketAccountFactory.contract.Call(opts, &out, "ctfExchange")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CtfExchange is a free data retrieval call binding the contract method 0x4b8295e7.
//
// Solidity: function ctfExchange() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactorySession) CtfExchange() (common.Address, error) {
	return _RarimarketAccountFactory.Contract.CtfExchange(&_RarimarketAccountFactory.CallOpts)
}

// CtfExchange is a free data retrieval call binding the contract method 0x4b8295e7.
//
// Solidity: function ctfExchange() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCallerSession) CtfExchange() (common.Address, error) {
	return _RarimarketAccountFactory.Contract.CtfExchange(&_RarimarketAccountFactory.CallOpts)
}

// GetRarimarketAccount is a free data retrieval call binding the contract method 0x2d1ac65a.
//
// Solidity: function getRarimarketAccount(bytes32 nullifier_) view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCaller) GetRarimarketAccount(opts *bind.CallOpts, nullifier_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _RarimarketAccountFactory.contract.Call(opts, &out, "getRarimarketAccount", nullifier_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRarimarketAccount is a free data retrieval call binding the contract method 0x2d1ac65a.
//
// Solidity: function getRarimarketAccount(bytes32 nullifier_) view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactorySession) GetRarimarketAccount(nullifier_ [32]byte) (common.Address, error) {
	return _RarimarketAccountFactory.Contract.GetRarimarketAccount(&_RarimarketAccountFactory.CallOpts, nullifier_)
}

// GetRarimarketAccount is a free data retrieval call binding the contract method 0x2d1ac65a.
//
// Solidity: function getRarimarketAccount(bytes32 nullifier_) view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCallerSession) GetRarimarketAccount(nullifier_ [32]byte) (common.Address, error) {
	return _RarimarketAccountFactory.Contract.GetRarimarketAccount(&_RarimarketAccountFactory.CallOpts, nullifier_)
}

// GetRarimarketAccountImplementation is a free data retrieval call binding the contract method 0xad04299b.
//
// Solidity: function getRarimarketAccountImplementation() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCaller) GetRarimarketAccountImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RarimarketAccountFactory.contract.Call(opts, &out, "getRarimarketAccountImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRarimarketAccountImplementation is a free data retrieval call binding the contract method 0xad04299b.
//
// Solidity: function getRarimarketAccountImplementation() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactorySession) GetRarimarketAccountImplementation() (common.Address, error) {
	return _RarimarketAccountFactory.Contract.GetRarimarketAccountImplementation(&_RarimarketAccountFactory.CallOpts)
}

// GetRarimarketAccountImplementation is a free data retrieval call binding the contract method 0xad04299b.
//
// Solidity: function getRarimarketAccountImplementation() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCallerSession) GetRarimarketAccountImplementation() (common.Address, error) {
	return _RarimarketAccountFactory.Contract.GetRarimarketAccountImplementation(&_RarimarketAccountFactory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RarimarketAccountFactory.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactorySession) Implementation() (common.Address, error) {
	return _RarimarketAccountFactory.Contract.Implementation(&_RarimarketAccountFactory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCallerSession) Implementation() (common.Address, error) {
	return _RarimarketAccountFactory.Contract.Implementation(&_RarimarketAccountFactory.CallOpts)
}

// NegRiskAdapter is a free data retrieval call binding the contract method 0xf6ef95a1.
//
// Solidity: function negRiskAdapter() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCaller) NegRiskAdapter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RarimarketAccountFactory.contract.Call(opts, &out, "negRiskAdapter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NegRiskAdapter is a free data retrieval call binding the contract method 0xf6ef95a1.
//
// Solidity: function negRiskAdapter() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactorySession) NegRiskAdapter() (common.Address, error) {
	return _RarimarketAccountFactory.Contract.NegRiskAdapter(&_RarimarketAccountFactory.CallOpts)
}

// NegRiskAdapter is a free data retrieval call binding the contract method 0xf6ef95a1.
//
// Solidity: function negRiskAdapter() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCallerSession) NegRiskAdapter() (common.Address, error) {
	return _RarimarketAccountFactory.Contract.NegRiskAdapter(&_RarimarketAccountFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RarimarketAccountFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactorySession) Owner() (common.Address, error) {
	return _RarimarketAccountFactory.Contract.Owner(&_RarimarketAccountFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCallerSession) Owner() (common.Address, error) {
	return _RarimarketAccountFactory.Contract.Owner(&_RarimarketAccountFactory.CallOpts)
}

// Points is a free data retrieval call binding the contract method 0x1be6dd64.
//
// Solidity: function points() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCaller) Points(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RarimarketAccountFactory.contract.Call(opts, &out, "points")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Points is a free data retrieval call binding the contract method 0x1be6dd64.
//
// Solidity: function points() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactorySession) Points() (common.Address, error) {
	return _RarimarketAccountFactory.Contract.Points(&_RarimarketAccountFactory.CallOpts)
}

// Points is a free data retrieval call binding the contract method 0x1be6dd64.
//
// Solidity: function points() view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCallerSession) Points() (common.Address, error) {
	return _RarimarketAccountFactory.Contract.Points(&_RarimarketAccountFactory.CallOpts)
}

// PredictRarimarketAccountAddress is a free data retrieval call binding the contract method 0xaedbd803.
//
// Solidity: function predictRarimarketAccountAddress(bytes32 nullifier_) view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCaller) PredictRarimarketAccountAddress(opts *bind.CallOpts, nullifier_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _RarimarketAccountFactory.contract.Call(opts, &out, "predictRarimarketAccountAddress", nullifier_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PredictRarimarketAccountAddress is a free data retrieval call binding the contract method 0xaedbd803.
//
// Solidity: function predictRarimarketAccountAddress(bytes32 nullifier_) view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactorySession) PredictRarimarketAccountAddress(nullifier_ [32]byte) (common.Address, error) {
	return _RarimarketAccountFactory.Contract.PredictRarimarketAccountAddress(&_RarimarketAccountFactory.CallOpts, nullifier_)
}

// PredictRarimarketAccountAddress is a free data retrieval call binding the contract method 0xaedbd803.
//
// Solidity: function predictRarimarketAccountAddress(bytes32 nullifier_) view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCallerSession) PredictRarimarketAccountAddress(nullifier_ [32]byte) (common.Address, error) {
	return _RarimarketAccountFactory.Contract.PredictRarimarketAccountAddress(&_RarimarketAccountFactory.CallOpts, nullifier_)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _RarimarketAccountFactory.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RarimarketAccountFactory *RarimarketAccountFactorySession) ProxiableUUID() ([32]byte, error) {
	return _RarimarketAccountFactory.Contract.ProxiableUUID(&_RarimarketAccountFactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _RarimarketAccountFactory.Contract.ProxiableUUID(&_RarimarketAccountFactory.CallOpts)
}

// RarimarketAccounts is a free data retrieval call binding the contract method 0x60b7ddc0.
//
// Solidity: function rarimarketAccounts(bytes32 ) view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCaller) RarimarketAccounts(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _RarimarketAccountFactory.contract.Call(opts, &out, "rarimarketAccounts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RarimarketAccounts is a free data retrieval call binding the contract method 0x60b7ddc0.
//
// Solidity: function rarimarketAccounts(bytes32 ) view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactorySession) RarimarketAccounts(arg0 [32]byte) (common.Address, error) {
	return _RarimarketAccountFactory.Contract.RarimarketAccounts(&_RarimarketAccountFactory.CallOpts, arg0)
}

// RarimarketAccounts is a free data retrieval call binding the contract method 0x60b7ddc0.
//
// Solidity: function rarimarketAccounts(bytes32 ) view returns(address)
func (_RarimarketAccountFactory *RarimarketAccountFactoryCallerSession) RarimarketAccounts(arg0 [32]byte) (common.Address, error) {
	return _RarimarketAccountFactory.Contract.RarimarketAccounts(&_RarimarketAccountFactory.CallOpts, arg0)
}

// RarimarketAccountFactoryInit is a paid mutator transaction binding the contract method 0xa3a4aa4f.
//
// Solidity: function __RarimarketAccountFactory_init(address rarimarketAccountImplementation_, address negRiskAdapter_, address ctfExchange_, address points_, address ctf_) returns()
func (_RarimarketAccountFactory *RarimarketAccountFactoryTransactor) RarimarketAccountFactoryInit(opts *bind.TransactOpts, rarimarketAccountImplementation_ common.Address, negRiskAdapter_ common.Address, ctfExchange_ common.Address, points_ common.Address, ctf_ common.Address) (*types.Transaction, error) {
	return _RarimarketAccountFactory.contract.Transact(opts, "__RarimarketAccountFactory_init", rarimarketAccountImplementation_, negRiskAdapter_, ctfExchange_, points_, ctf_)
}

// RarimarketAccountFactoryInit is a paid mutator transaction binding the contract method 0xa3a4aa4f.
//
// Solidity: function __RarimarketAccountFactory_init(address rarimarketAccountImplementation_, address negRiskAdapter_, address ctfExchange_, address points_, address ctf_) returns()
func (_RarimarketAccountFactory *RarimarketAccountFactorySession) RarimarketAccountFactoryInit(rarimarketAccountImplementation_ common.Address, negRiskAdapter_ common.Address, ctfExchange_ common.Address, points_ common.Address, ctf_ common.Address) (*types.Transaction, error) {
	return _RarimarketAccountFactory.Contract.RarimarketAccountFactoryInit(&_RarimarketAccountFactory.TransactOpts, rarimarketAccountImplementation_, negRiskAdapter_, ctfExchange_, points_, ctf_)
}

// RarimarketAccountFactoryInit is a paid mutator transaction binding the contract method 0xa3a4aa4f.
//
// Solidity: function __RarimarketAccountFactory_init(address rarimarketAccountImplementation_, address negRiskAdapter_, address ctfExchange_, address points_, address ctf_) returns()
func (_RarimarketAccountFactory *RarimarketAccountFactoryTransactorSession) RarimarketAccountFactoryInit(rarimarketAccountImplementation_ common.Address, negRiskAdapter_ common.Address, ctfExchange_ common.Address, points_ common.Address, ctf_ common.Address) (*types.Transaction, error) {
	return _RarimarketAccountFactory.Contract.RarimarketAccountFactoryInit(&_RarimarketAccountFactory.TransactOpts, rarimarketAccountImplementation_, negRiskAdapter_, ctfExchange_, points_, ctf_)
}

// DeployRarimarketAccount is a paid mutator transaction binding the contract method 0xcd95b7b0.
//
// Solidity: function deployRarimarketAccount(bytes32 nullifier_) returns()
func (_RarimarketAccountFactory *RarimarketAccountFactoryTransactor) DeployRarimarketAccount(opts *bind.TransactOpts, nullifier_ [32]byte) (*types.Transaction, error) {
	return _RarimarketAccountFactory.contract.Transact(opts, "deployRarimarketAccount", nullifier_)
}

// DeployRarimarketAccount is a paid mutator transaction binding the contract method 0xcd95b7b0.
//
// Solidity: function deployRarimarketAccount(bytes32 nullifier_) returns()
func (_RarimarketAccountFactory *RarimarketAccountFactorySession) DeployRarimarketAccount(nullifier_ [32]byte) (*types.Transaction, error) {
	return _RarimarketAccountFactory.Contract.DeployRarimarketAccount(&_RarimarketAccountFactory.TransactOpts, nullifier_)
}

// DeployRarimarketAccount is a paid mutator transaction binding the contract method 0xcd95b7b0.
//
// Solidity: function deployRarimarketAccount(bytes32 nullifier_) returns()
func (_RarimarketAccountFactory *RarimarketAccountFactoryTransactorSession) DeployRarimarketAccount(nullifier_ [32]byte) (*types.Transaction, error) {
	return _RarimarketAccountFactory.Contract.DeployRarimarketAccount(&_RarimarketAccountFactory.TransactOpts, nullifier_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RarimarketAccountFactory *RarimarketAccountFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RarimarketAccountFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RarimarketAccountFactory *RarimarketAccountFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _RarimarketAccountFactory.Contract.RenounceOwnership(&_RarimarketAccountFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RarimarketAccountFactory *RarimarketAccountFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RarimarketAccountFactory.Contract.RenounceOwnership(&_RarimarketAccountFactory.TransactOpts)
}

// SetRarimarketAccountImplementation is a paid mutator transaction binding the contract method 0x85a3cbd9.
//
// Solidity: function setRarimarketAccountImplementation(address newImplementation) returns()
func (_RarimarketAccountFactory *RarimarketAccountFactoryTransactor) SetRarimarketAccountImplementation(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _RarimarketAccountFactory.contract.Transact(opts, "setRarimarketAccountImplementation", newImplementation)
}

// SetRarimarketAccountImplementation is a paid mutator transaction binding the contract method 0x85a3cbd9.
//
// Solidity: function setRarimarketAccountImplementation(address newImplementation) returns()
func (_RarimarketAccountFactory *RarimarketAccountFactorySession) SetRarimarketAccountImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _RarimarketAccountFactory.Contract.SetRarimarketAccountImplementation(&_RarimarketAccountFactory.TransactOpts, newImplementation)
}

// SetRarimarketAccountImplementation is a paid mutator transaction binding the contract method 0x85a3cbd9.
//
// Solidity: function setRarimarketAccountImplementation(address newImplementation) returns()
func (_RarimarketAccountFactory *RarimarketAccountFactoryTransactorSession) SetRarimarketAccountImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _RarimarketAccountFactory.Contract.SetRarimarketAccountImplementation(&_RarimarketAccountFactory.TransactOpts, newImplementation)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RarimarketAccountFactory *RarimarketAccountFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RarimarketAccountFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RarimarketAccountFactory *RarimarketAccountFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RarimarketAccountFactory.Contract.TransferOwnership(&_RarimarketAccountFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RarimarketAccountFactory *RarimarketAccountFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RarimarketAccountFactory.Contract.TransferOwnership(&_RarimarketAccountFactory.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RarimarketAccountFactory *RarimarketAccountFactoryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _RarimarketAccountFactory.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RarimarketAccountFactory *RarimarketAccountFactorySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _RarimarketAccountFactory.Contract.UpgradeTo(&_RarimarketAccountFactory.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_RarimarketAccountFactory *RarimarketAccountFactoryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _RarimarketAccountFactory.Contract.UpgradeTo(&_RarimarketAccountFactory.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RarimarketAccountFactory *RarimarketAccountFactoryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RarimarketAccountFactory.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RarimarketAccountFactory *RarimarketAccountFactorySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RarimarketAccountFactory.Contract.UpgradeToAndCall(&_RarimarketAccountFactory.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_RarimarketAccountFactory *RarimarketAccountFactoryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _RarimarketAccountFactory.Contract.UpgradeToAndCall(&_RarimarketAccountFactory.TransactOpts, newImplementation, data)
}

// RarimarketAccountFactoryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the RarimarketAccountFactory contract.
type RarimarketAccountFactoryAdminChangedIterator struct {
	Event *RarimarketAccountFactoryAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RarimarketAccountFactoryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RarimarketAccountFactoryAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RarimarketAccountFactoryAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RarimarketAccountFactoryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RarimarketAccountFactoryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RarimarketAccountFactoryAdminChanged represents a AdminChanged event raised by the RarimarketAccountFactory contract.
type RarimarketAccountFactoryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_RarimarketAccountFactory *RarimarketAccountFactoryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*RarimarketAccountFactoryAdminChangedIterator, error) {

	logs, sub, err := _RarimarketAccountFactory.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &RarimarketAccountFactoryAdminChangedIterator{contract: _RarimarketAccountFactory.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_RarimarketAccountFactory *RarimarketAccountFactoryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *RarimarketAccountFactoryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _RarimarketAccountFactory.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RarimarketAccountFactoryAdminChanged)
				if err := _RarimarketAccountFactory.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_RarimarketAccountFactory *RarimarketAccountFactoryFilterer) ParseAdminChanged(log types.Log) (*RarimarketAccountFactoryAdminChanged, error) {
	event := new(RarimarketAccountFactoryAdminChanged)
	if err := _RarimarketAccountFactory.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RarimarketAccountFactoryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the RarimarketAccountFactory contract.
type RarimarketAccountFactoryBeaconUpgradedIterator struct {
	Event *RarimarketAccountFactoryBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RarimarketAccountFactoryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RarimarketAccountFactoryBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RarimarketAccountFactoryBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RarimarketAccountFactoryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RarimarketAccountFactoryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RarimarketAccountFactoryBeaconUpgraded represents a BeaconUpgraded event raised by the RarimarketAccountFactory contract.
type RarimarketAccountFactoryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_RarimarketAccountFactory *RarimarketAccountFactoryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*RarimarketAccountFactoryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _RarimarketAccountFactory.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &RarimarketAccountFactoryBeaconUpgradedIterator{contract: _RarimarketAccountFactory.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_RarimarketAccountFactory *RarimarketAccountFactoryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *RarimarketAccountFactoryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _RarimarketAccountFactory.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RarimarketAccountFactoryBeaconUpgraded)
				if err := _RarimarketAccountFactory.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_RarimarketAccountFactory *RarimarketAccountFactoryFilterer) ParseBeaconUpgraded(log types.Log) (*RarimarketAccountFactoryBeaconUpgraded, error) {
	event := new(RarimarketAccountFactoryBeaconUpgraded)
	if err := _RarimarketAccountFactory.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RarimarketAccountFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the RarimarketAccountFactory contract.
type RarimarketAccountFactoryInitializedIterator struct {
	Event *RarimarketAccountFactoryInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RarimarketAccountFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RarimarketAccountFactoryInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RarimarketAccountFactoryInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RarimarketAccountFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RarimarketAccountFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RarimarketAccountFactoryInitialized represents a Initialized event raised by the RarimarketAccountFactory contract.
type RarimarketAccountFactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RarimarketAccountFactory *RarimarketAccountFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*RarimarketAccountFactoryInitializedIterator, error) {

	logs, sub, err := _RarimarketAccountFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RarimarketAccountFactoryInitializedIterator{contract: _RarimarketAccountFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RarimarketAccountFactory *RarimarketAccountFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RarimarketAccountFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _RarimarketAccountFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RarimarketAccountFactoryInitialized)
				if err := _RarimarketAccountFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_RarimarketAccountFactory *RarimarketAccountFactoryFilterer) ParseInitialized(log types.Log) (*RarimarketAccountFactoryInitialized, error) {
	event := new(RarimarketAccountFactoryInitialized)
	if err := _RarimarketAccountFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RarimarketAccountFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RarimarketAccountFactory contract.
type RarimarketAccountFactoryOwnershipTransferredIterator struct {
	Event *RarimarketAccountFactoryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RarimarketAccountFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RarimarketAccountFactoryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RarimarketAccountFactoryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RarimarketAccountFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RarimarketAccountFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RarimarketAccountFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the RarimarketAccountFactory contract.
type RarimarketAccountFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RarimarketAccountFactory *RarimarketAccountFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RarimarketAccountFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RarimarketAccountFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RarimarketAccountFactoryOwnershipTransferredIterator{contract: _RarimarketAccountFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RarimarketAccountFactory *RarimarketAccountFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RarimarketAccountFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RarimarketAccountFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RarimarketAccountFactoryOwnershipTransferred)
				if err := _RarimarketAccountFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RarimarketAccountFactory *RarimarketAccountFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*RarimarketAccountFactoryOwnershipTransferred, error) {
	event := new(RarimarketAccountFactoryOwnershipTransferred)
	if err := _RarimarketAccountFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RarimarketAccountFactoryRarimarketAccountDeployedIterator is returned from FilterRarimarketAccountDeployed and is used to iterate over the raw logs and unpacked data for RarimarketAccountDeployed events raised by the RarimarketAccountFactory contract.
type RarimarketAccountFactoryRarimarketAccountDeployedIterator struct {
	Event *RarimarketAccountFactoryRarimarketAccountDeployed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RarimarketAccountFactoryRarimarketAccountDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RarimarketAccountFactoryRarimarketAccountDeployed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RarimarketAccountFactoryRarimarketAccountDeployed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RarimarketAccountFactoryRarimarketAccountDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RarimarketAccountFactoryRarimarketAccountDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RarimarketAccountFactoryRarimarketAccountDeployed represents a RarimarketAccountDeployed event raised by the RarimarketAccountFactory contract.
type RarimarketAccountFactoryRarimarketAccountDeployed struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRarimarketAccountDeployed is a free log retrieval operation binding the contract event 0xf8078a05821fe3a0f57304f01a507627f803be4212ef8a593a28a395b43763de.
//
// Solidity: event RarimarketAccountDeployed(address indexed account)
func (_RarimarketAccountFactory *RarimarketAccountFactoryFilterer) FilterRarimarketAccountDeployed(opts *bind.FilterOpts, account []common.Address) (*RarimarketAccountFactoryRarimarketAccountDeployedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RarimarketAccountFactory.contract.FilterLogs(opts, "RarimarketAccountDeployed", accountRule)
	if err != nil {
		return nil, err
	}
	return &RarimarketAccountFactoryRarimarketAccountDeployedIterator{contract: _RarimarketAccountFactory.contract, event: "RarimarketAccountDeployed", logs: logs, sub: sub}, nil
}

// WatchRarimarketAccountDeployed is a free log subscription operation binding the contract event 0xf8078a05821fe3a0f57304f01a507627f803be4212ef8a593a28a395b43763de.
//
// Solidity: event RarimarketAccountDeployed(address indexed account)
func (_RarimarketAccountFactory *RarimarketAccountFactoryFilterer) WatchRarimarketAccountDeployed(opts *bind.WatchOpts, sink chan<- *RarimarketAccountFactoryRarimarketAccountDeployed, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RarimarketAccountFactory.contract.WatchLogs(opts, "RarimarketAccountDeployed", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RarimarketAccountFactoryRarimarketAccountDeployed)
				if err := _RarimarketAccountFactory.contract.UnpackLog(event, "RarimarketAccountDeployed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRarimarketAccountDeployed is a log parse operation binding the contract event 0xf8078a05821fe3a0f57304f01a507627f803be4212ef8a593a28a395b43763de.
//
// Solidity: event RarimarketAccountDeployed(address indexed account)
func (_RarimarketAccountFactory *RarimarketAccountFactoryFilterer) ParseRarimarketAccountDeployed(log types.Log) (*RarimarketAccountFactoryRarimarketAccountDeployed, error) {
	event := new(RarimarketAccountFactoryRarimarketAccountDeployed)
	if err := _RarimarketAccountFactory.contract.UnpackLog(event, "RarimarketAccountDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RarimarketAccountFactoryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the RarimarketAccountFactory contract.
type RarimarketAccountFactoryUpgradedIterator struct {
	Event *RarimarketAccountFactoryUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RarimarketAccountFactoryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RarimarketAccountFactoryUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RarimarketAccountFactoryUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RarimarketAccountFactoryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RarimarketAccountFactoryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RarimarketAccountFactoryUpgraded represents a Upgraded event raised by the RarimarketAccountFactory contract.
type RarimarketAccountFactoryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RarimarketAccountFactory *RarimarketAccountFactoryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*RarimarketAccountFactoryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RarimarketAccountFactory.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &RarimarketAccountFactoryUpgradedIterator{contract: _RarimarketAccountFactory.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RarimarketAccountFactory *RarimarketAccountFactoryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *RarimarketAccountFactoryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _RarimarketAccountFactory.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RarimarketAccountFactoryUpgraded)
				if err := _RarimarketAccountFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_RarimarketAccountFactory *RarimarketAccountFactoryFilterer) ParseUpgraded(log types.Log) (*RarimarketAccountFactoryUpgraded, error) {
	event := new(RarimarketAccountFactoryUpgraded)
	if err := _RarimarketAccountFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
