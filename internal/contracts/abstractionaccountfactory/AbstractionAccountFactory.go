// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abstractionaccountfactory

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

// AbstractionAccountFactoryMetaData contains all meta data concerning the AbstractionAccountFactory contract.
var AbstractionAccountFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AbstractionAccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"abstractionAccountImplementation_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"negRiskAdapter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ctfExchange_\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"points_\",\"type\":\"address\"},{\"internalType\":\"contractIERC1155\",\"name\":\"ctf_\",\"type\":\"address\"}],\"name\":\"__AbstractionAccountFactory_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ctf\",\"outputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ctfExchange\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nullifier_\",\"type\":\"bytes32\"}],\"name\":\"deployAbstractionAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nullifier_\",\"type\":\"bytes32\"}],\"name\":\"getAbstractionAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAbstractionAccountImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"negRiskAdapter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"points\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nullifier_\",\"type\":\"bytes32\"}],\"name\":\"predictAbstractionAccountAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"abstractionAccounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"setAbstractionAccountImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// AbstractionAccountFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use AbstractionAccountFactoryMetaData.ABI instead.
var AbstractionAccountFactoryABI = AbstractionAccountFactoryMetaData.ABI

// AbstractionAccountFactory is an auto generated Go binding around an Ethereum contract.
type AbstractionAccountFactory struct {
	AbstractionAccountFactoryCaller     // Read-only binding to the contract
	AbstractionAccountFactoryTransactor // Write-only binding to the contract
	AbstractionAccountFactoryFilterer   // Log filterer for contract events
}

// AbstractionAccountFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbstractionAccountFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractionAccountFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbstractionAccountFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractionAccountFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbstractionAccountFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractionAccountFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbstractionAccountFactorySession struct {
	Contract     *AbstractionAccountFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AbstractionAccountFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbstractionAccountFactoryCallerSession struct {
	Contract *AbstractionAccountFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// AbstractionAccountFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbstractionAccountFactoryTransactorSession struct {
	Contract     *AbstractionAccountFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// AbstractionAccountFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbstractionAccountFactoryRaw struct {
	Contract *AbstractionAccountFactory // Generic contract binding to access the raw methods on
}

// AbstractionAccountFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbstractionAccountFactoryCallerRaw struct {
	Contract *AbstractionAccountFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// AbstractionAccountFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbstractionAccountFactoryTransactorRaw struct {
	Contract *AbstractionAccountFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbstractionAccountFactory creates a new instance of AbstractionAccountFactory, bound to a specific deployed contract.
func NewAbstractionAccountFactory(address common.Address, backend bind.ContractBackend) (*AbstractionAccountFactory, error) {
	contract, err := bindAbstractionAccountFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AbstractionAccountFactory{AbstractionAccountFactoryCaller: AbstractionAccountFactoryCaller{contract: contract}, AbstractionAccountFactoryTransactor: AbstractionAccountFactoryTransactor{contract: contract}, AbstractionAccountFactoryFilterer: AbstractionAccountFactoryFilterer{contract: contract}}, nil
}

// NewAbstractionAccountFactoryCaller creates a new read-only instance of AbstractionAccountFactory, bound to a specific deployed contract.
func NewAbstractionAccountFactoryCaller(address common.Address, caller bind.ContractCaller) (*AbstractionAccountFactoryCaller, error) {
	contract, err := bindAbstractionAccountFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractionAccountFactoryCaller{contract: contract}, nil
}

// NewAbstractionAccountFactoryTransactor creates a new write-only instance of AbstractionAccountFactory, bound to a specific deployed contract.
func NewAbstractionAccountFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*AbstractionAccountFactoryTransactor, error) {
	contract, err := bindAbstractionAccountFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractionAccountFactoryTransactor{contract: contract}, nil
}

// NewAbstractionAccountFactoryFilterer creates a new log filterer instance of AbstractionAccountFactory, bound to a specific deployed contract.
func NewAbstractionAccountFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*AbstractionAccountFactoryFilterer, error) {
	contract, err := bindAbstractionAccountFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbstractionAccountFactoryFilterer{contract: contract}, nil
}

// bindAbstractionAccountFactory binds a generic wrapper to an already deployed contract.
func bindAbstractionAccountFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbstractionAccountFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractionAccountFactory *AbstractionAccountFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbstractionAccountFactory.Contract.AbstractionAccountFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractionAccountFactory *AbstractionAccountFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractionAccountFactory.Contract.AbstractionAccountFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractionAccountFactory *AbstractionAccountFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractionAccountFactory.Contract.AbstractionAccountFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractionAccountFactory *AbstractionAccountFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbstractionAccountFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractionAccountFactory *AbstractionAccountFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractionAccountFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractionAccountFactory *AbstractionAccountFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractionAccountFactory.Contract.contract.Transact(opts, method, params...)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCaller) Ctf(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbstractionAccountFactory.contract.Call(opts, &out, "ctf")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactorySession) Ctf() (common.Address, error) {
	return _AbstractionAccountFactory.Contract.Ctf(&_AbstractionAccountFactory.CallOpts)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCallerSession) Ctf() (common.Address, error) {
	return _AbstractionAccountFactory.Contract.Ctf(&_AbstractionAccountFactory.CallOpts)
}

// CtfExchange is a free data retrieval call binding the contract method 0x4b8295e7.
//
// Solidity: function ctfExchange() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCaller) CtfExchange(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbstractionAccountFactory.contract.Call(opts, &out, "ctfExchange")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CtfExchange is a free data retrieval call binding the contract method 0x4b8295e7.
//
// Solidity: function ctfExchange() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactorySession) CtfExchange() (common.Address, error) {
	return _AbstractionAccountFactory.Contract.CtfExchange(&_AbstractionAccountFactory.CallOpts)
}

// CtfExchange is a free data retrieval call binding the contract method 0x4b8295e7.
//
// Solidity: function ctfExchange() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCallerSession) CtfExchange() (common.Address, error) {
	return _AbstractionAccountFactory.Contract.CtfExchange(&_AbstractionAccountFactory.CallOpts)
}

// GetAbstractionAccount is a free data retrieval call binding the contract method 0x2d1ac65a.
//
// Solidity: function getAbstractionAccount(bytes32 nullifier_) view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCaller) GetAbstractionAccount(opts *bind.CallOpts, nullifier_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AbstractionAccountFactory.contract.Call(opts, &out, "getAbstractionAccount", nullifier_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAbstractionAccount is a free data retrieval call binding the contract method 0x2d1ac65a.
//
// Solidity: function getAbstractionAccount(bytes32 nullifier_) view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactorySession) GetAbstractionAccount(nullifier_ [32]byte) (common.Address, error) {
	return _AbstractionAccountFactory.Contract.GetAbstractionAccount(&_AbstractionAccountFactory.CallOpts, nullifier_)
}

// GetAbstractionAccount is a free data retrieval call binding the contract method 0x2d1ac65a.
//
// Solidity: function getAbstractionAccount(bytes32 nullifier_) view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCallerSession) GetAbstractionAccount(nullifier_ [32]byte) (common.Address, error) {
	return _AbstractionAccountFactory.Contract.GetAbstractionAccount(&_AbstractionAccountFactory.CallOpts, nullifier_)
}

// GetAbstractionAccountImplementation is a free data retrieval call binding the contract method 0xad04299b.
//
// Solidity: function getAbstractionAccountImplementation() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCaller) GetAbstractionAccountImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbstractionAccountFactory.contract.Call(opts, &out, "getAbstractionAccountImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAbstractionAccountImplementation is a free data retrieval call binding the contract method 0xad04299b.
//
// Solidity: function getAbstractionAccountImplementation() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactorySession) GetAbstractionAccountImplementation() (common.Address, error) {
	return _AbstractionAccountFactory.Contract.GetAbstractionAccountImplementation(&_AbstractionAccountFactory.CallOpts)
}

// GetAbstractionAccountImplementation is a free data retrieval call binding the contract method 0xad04299b.
//
// Solidity: function getAbstractionAccountImplementation() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCallerSession) GetAbstractionAccountImplementation() (common.Address, error) {
	return _AbstractionAccountFactory.Contract.GetAbstractionAccountImplementation(&_AbstractionAccountFactory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbstractionAccountFactory.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactorySession) Implementation() (common.Address, error) {
	return _AbstractionAccountFactory.Contract.Implementation(&_AbstractionAccountFactory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCallerSession) Implementation() (common.Address, error) {
	return _AbstractionAccountFactory.Contract.Implementation(&_AbstractionAccountFactory.CallOpts)
}

// NegRiskAdapter is a free data retrieval call binding the contract method 0xf6ef95a1.
//
// Solidity: function negRiskAdapter() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCaller) NegRiskAdapter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbstractionAccountFactory.contract.Call(opts, &out, "negRiskAdapter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NegRiskAdapter is a free data retrieval call binding the contract method 0xf6ef95a1.
//
// Solidity: function negRiskAdapter() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactorySession) NegRiskAdapter() (common.Address, error) {
	return _AbstractionAccountFactory.Contract.NegRiskAdapter(&_AbstractionAccountFactory.CallOpts)
}

// NegRiskAdapter is a free data retrieval call binding the contract method 0xf6ef95a1.
//
// Solidity: function negRiskAdapter() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCallerSession) NegRiskAdapter() (common.Address, error) {
	return _AbstractionAccountFactory.Contract.NegRiskAdapter(&_AbstractionAccountFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbstractionAccountFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactorySession) Owner() (common.Address, error) {
	return _AbstractionAccountFactory.Contract.Owner(&_AbstractionAccountFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCallerSession) Owner() (common.Address, error) {
	return _AbstractionAccountFactory.Contract.Owner(&_AbstractionAccountFactory.CallOpts)
}

// Points is a free data retrieval call binding the contract method 0x1be6dd64.
//
// Solidity: function points() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCaller) Points(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbstractionAccountFactory.contract.Call(opts, &out, "points")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Points is a free data retrieval call binding the contract method 0x1be6dd64.
//
// Solidity: function points() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactorySession) Points() (common.Address, error) {
	return _AbstractionAccountFactory.Contract.Points(&_AbstractionAccountFactory.CallOpts)
}

// Points is a free data retrieval call binding the contract method 0x1be6dd64.
//
// Solidity: function points() view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCallerSession) Points() (common.Address, error) {
	return _AbstractionAccountFactory.Contract.Points(&_AbstractionAccountFactory.CallOpts)
}

// PredictAbstractionAccountAddress is a free data retrieval call binding the contract method 0xaedbd803.
//
// Solidity: function predictAbstractionAccountAddress(bytes32 nullifier_) view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCaller) PredictAbstractionAccountAddress(opts *bind.CallOpts, nullifier_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AbstractionAccountFactory.contract.Call(opts, &out, "predictAbstractionAccountAddress", nullifier_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PredictAbstractionAccountAddress is a free data retrieval call binding the contract method 0xaedbd803.
//
// Solidity: function predictAbstractionAccountAddress(bytes32 nullifier_) view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactorySession) PredictAbstractionAccountAddress(nullifier_ [32]byte) (common.Address, error) {
	return _AbstractionAccountFactory.Contract.PredictAbstractionAccountAddress(&_AbstractionAccountFactory.CallOpts, nullifier_)
}

// PredictAbstractionAccountAddress is a free data retrieval call binding the contract method 0xaedbd803.
//
// Solidity: function predictAbstractionAccountAddress(bytes32 nullifier_) view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCallerSession) PredictAbstractionAccountAddress(nullifier_ [32]byte) (common.Address, error) {
	return _AbstractionAccountFactory.Contract.PredictAbstractionAccountAddress(&_AbstractionAccountFactory.CallOpts, nullifier_)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AbstractionAccountFactory.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AbstractionAccountFactory *AbstractionAccountFactorySession) ProxiableUUID() ([32]byte, error) {
	return _AbstractionAccountFactory.Contract.ProxiableUUID(&_AbstractionAccountFactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _AbstractionAccountFactory.Contract.ProxiableUUID(&_AbstractionAccountFactory.CallOpts)
}

// AbstractionAccounts is a free data retrieval call binding the contract method 0x60b7ddc0.
//
// Solidity: function abstractionAccounts(bytes32 ) view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCaller) AbstractionAccounts(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AbstractionAccountFactory.contract.Call(opts, &out, "abstractionAccounts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AbstractionAccounts is a free data retrieval call binding the contract method 0x60b7ddc0.
//
// Solidity: function abstractionAccounts(bytes32 ) view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactorySession) AbstractionAccounts(arg0 [32]byte) (common.Address, error) {
	return _AbstractionAccountFactory.Contract.AbstractionAccounts(&_AbstractionAccountFactory.CallOpts, arg0)
}

// AbstractionAccounts is a free data retrieval call binding the contract method 0x60b7ddc0.
//
// Solidity: function abstractionAccounts(bytes32 ) view returns(address)
func (_AbstractionAccountFactory *AbstractionAccountFactoryCallerSession) AbstractionAccounts(arg0 [32]byte) (common.Address, error) {
	return _AbstractionAccountFactory.Contract.AbstractionAccounts(&_AbstractionAccountFactory.CallOpts, arg0)
}

// AbstractionAccountFactoryInit is a paid mutator transaction binding the contract method 0xa3a4aa4f.
//
// Solidity: function __AbstractionAccountFactory_init(address abstractionAccountImplementation_, address negRiskAdapter_, address ctfExchange_, address points_, address ctf_) returns()
func (_AbstractionAccountFactory *AbstractionAccountFactoryTransactor) AbstractionAccountFactoryInit(opts *bind.TransactOpts, abstractionAccountImplementation_ common.Address, negRiskAdapter_ common.Address, ctfExchange_ common.Address, points_ common.Address, ctf_ common.Address) (*types.Transaction, error) {
	return _AbstractionAccountFactory.contract.Transact(opts, "__AbstractionAccountFactory_init", abstractionAccountImplementation_, negRiskAdapter_, ctfExchange_, points_, ctf_)
}

// AbstractionAccountFactoryInit is a paid mutator transaction binding the contract method 0xa3a4aa4f.
//
// Solidity: function __AbstractionAccountFactory_init(address abstractionAccountImplementation_, address negRiskAdapter_, address ctfExchange_, address points_, address ctf_) returns()
func (_AbstractionAccountFactory *AbstractionAccountFactorySession) AbstractionAccountFactoryInit(abstractionAccountImplementation_ common.Address, negRiskAdapter_ common.Address, ctfExchange_ common.Address, points_ common.Address, ctf_ common.Address) (*types.Transaction, error) {
	return _AbstractionAccountFactory.Contract.AbstractionAccountFactoryInit(&_AbstractionAccountFactory.TransactOpts, abstractionAccountImplementation_, negRiskAdapter_, ctfExchange_, points_, ctf_)
}

// AbstractionAccountFactoryInit is a paid mutator transaction binding the contract method 0xa3a4aa4f.
//
// Solidity: function __AbstractionAccountFactory_init(address abstractionAccountImplementation_, address negRiskAdapter_, address ctfExchange_, address points_, address ctf_) returns()
func (_AbstractionAccountFactory *AbstractionAccountFactoryTransactorSession) AbstractionAccountFactoryInit(abstractionAccountImplementation_ common.Address, negRiskAdapter_ common.Address, ctfExchange_ common.Address, points_ common.Address, ctf_ common.Address) (*types.Transaction, error) {
	return _AbstractionAccountFactory.Contract.AbstractionAccountFactoryInit(&_AbstractionAccountFactory.TransactOpts, abstractionAccountImplementation_, negRiskAdapter_, ctfExchange_, points_, ctf_)
}

// DeployAbstractionAccount is a paid mutator transaction binding the contract method 0xcd95b7b0.
//
// Solidity: function deployAbstractionAccount(bytes32 nullifier_) returns()
func (_AbstractionAccountFactory *AbstractionAccountFactoryTransactor) DeployAbstractionAccount(opts *bind.TransactOpts, nullifier_ [32]byte) (*types.Transaction, error) {
	return _AbstractionAccountFactory.contract.Transact(opts, "deployAbstractionAccount", nullifier_)
}

// DeployAbstractionAccount is a paid mutator transaction binding the contract method 0xcd95b7b0.
//
// Solidity: function deployAbstractionAccount(bytes32 nullifier_) returns()
func (_AbstractionAccountFactory *AbstractionAccountFactorySession) DeployAbstractionAccount(nullifier_ [32]byte) (*types.Transaction, error) {
	return _AbstractionAccountFactory.Contract.DeployAbstractionAccount(&_AbstractionAccountFactory.TransactOpts, nullifier_)
}

// DeployAbstractionAccount is a paid mutator transaction binding the contract method 0xcd95b7b0.
//
// Solidity: function deployAbstractionAccount(bytes32 nullifier_) returns()
func (_AbstractionAccountFactory *AbstractionAccountFactoryTransactorSession) DeployAbstractionAccount(nullifier_ [32]byte) (*types.Transaction, error) {
	return _AbstractionAccountFactory.Contract.DeployAbstractionAccount(&_AbstractionAccountFactory.TransactOpts, nullifier_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AbstractionAccountFactory *AbstractionAccountFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractionAccountFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AbstractionAccountFactory *AbstractionAccountFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _AbstractionAccountFactory.Contract.RenounceOwnership(&_AbstractionAccountFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AbstractionAccountFactory *AbstractionAccountFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AbstractionAccountFactory.Contract.RenounceOwnership(&_AbstractionAccountFactory.TransactOpts)
}

// SetAbstractionAccountImplementation is a paid mutator transaction binding the contract method 0x85a3cbd9.
//
// Solidity: function setAbstractionAccountImplementation(address newImplementation) returns()
func (_AbstractionAccountFactory *AbstractionAccountFactoryTransactor) SetAbstractionAccountImplementation(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _AbstractionAccountFactory.contract.Transact(opts, "setAbstractionAccountImplementation", newImplementation)
}

// SetAbstractionAccountImplementation is a paid mutator transaction binding the contract method 0x85a3cbd9.
//
// Solidity: function setAbstractionAccountImplementation(address newImplementation) returns()
func (_AbstractionAccountFactory *AbstractionAccountFactorySession) SetAbstractionAccountImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _AbstractionAccountFactory.Contract.SetAbstractionAccountImplementation(&_AbstractionAccountFactory.TransactOpts, newImplementation)
}

// SetAbstractionAccountImplementation is a paid mutator transaction binding the contract method 0x85a3cbd9.
//
// Solidity: function setAbstractionAccountImplementation(address newImplementation) returns()
func (_AbstractionAccountFactory *AbstractionAccountFactoryTransactorSession) SetAbstractionAccountImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _AbstractionAccountFactory.Contract.SetAbstractionAccountImplementation(&_AbstractionAccountFactory.TransactOpts, newImplementation)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AbstractionAccountFactory *AbstractionAccountFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AbstractionAccountFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AbstractionAccountFactory *AbstractionAccountFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AbstractionAccountFactory.Contract.TransferOwnership(&_AbstractionAccountFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AbstractionAccountFactory *AbstractionAccountFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AbstractionAccountFactory.Contract.TransferOwnership(&_AbstractionAccountFactory.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AbstractionAccountFactory *AbstractionAccountFactoryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _AbstractionAccountFactory.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AbstractionAccountFactory *AbstractionAccountFactorySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AbstractionAccountFactory.Contract.UpgradeTo(&_AbstractionAccountFactory.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AbstractionAccountFactory *AbstractionAccountFactoryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AbstractionAccountFactory.Contract.UpgradeTo(&_AbstractionAccountFactory.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AbstractionAccountFactory *AbstractionAccountFactoryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AbstractionAccountFactory.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AbstractionAccountFactory *AbstractionAccountFactorySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AbstractionAccountFactory.Contract.UpgradeToAndCall(&_AbstractionAccountFactory.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AbstractionAccountFactory *AbstractionAccountFactoryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AbstractionAccountFactory.Contract.UpgradeToAndCall(&_AbstractionAccountFactory.TransactOpts, newImplementation, data)
}

// AbstractionAccountFactoryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the AbstractionAccountFactory contract.
type AbstractionAccountFactoryAdminChangedIterator struct {
	Event *AbstractionAccountFactoryAdminChanged // Event containing the contract specifics and raw log

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
func (it *AbstractionAccountFactoryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbstractionAccountFactoryAdminChanged)
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
		it.Event = new(AbstractionAccountFactoryAdminChanged)
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
func (it *AbstractionAccountFactoryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbstractionAccountFactoryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbstractionAccountFactoryAdminChanged represents a AdminChanged event raised by the AbstractionAccountFactory contract.
type AbstractionAccountFactoryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AbstractionAccountFactory *AbstractionAccountFactoryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AbstractionAccountFactoryAdminChangedIterator, error) {

	logs, sub, err := _AbstractionAccountFactory.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AbstractionAccountFactoryAdminChangedIterator{contract: _AbstractionAccountFactory.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AbstractionAccountFactory *AbstractionAccountFactoryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AbstractionAccountFactoryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _AbstractionAccountFactory.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbstractionAccountFactoryAdminChanged)
				if err := _AbstractionAccountFactory.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_AbstractionAccountFactory *AbstractionAccountFactoryFilterer) ParseAdminChanged(log types.Log) (*AbstractionAccountFactoryAdminChanged, error) {
	event := new(AbstractionAccountFactoryAdminChanged)
	if err := _AbstractionAccountFactory.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbstractionAccountFactoryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the AbstractionAccountFactory contract.
type AbstractionAccountFactoryBeaconUpgradedIterator struct {
	Event *AbstractionAccountFactoryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *AbstractionAccountFactoryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbstractionAccountFactoryBeaconUpgraded)
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
		it.Event = new(AbstractionAccountFactoryBeaconUpgraded)
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
func (it *AbstractionAccountFactoryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbstractionAccountFactoryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbstractionAccountFactoryBeaconUpgraded represents a BeaconUpgraded event raised by the AbstractionAccountFactory contract.
type AbstractionAccountFactoryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AbstractionAccountFactory *AbstractionAccountFactoryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*AbstractionAccountFactoryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _AbstractionAccountFactory.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &AbstractionAccountFactoryBeaconUpgradedIterator{contract: _AbstractionAccountFactory.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AbstractionAccountFactory *AbstractionAccountFactoryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *AbstractionAccountFactoryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _AbstractionAccountFactory.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbstractionAccountFactoryBeaconUpgraded)
				if err := _AbstractionAccountFactory.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_AbstractionAccountFactory *AbstractionAccountFactoryFilterer) ParseBeaconUpgraded(log types.Log) (*AbstractionAccountFactoryBeaconUpgraded, error) {
	event := new(AbstractionAccountFactoryBeaconUpgraded)
	if err := _AbstractionAccountFactory.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbstractionAccountFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AbstractionAccountFactory contract.
type AbstractionAccountFactoryInitializedIterator struct {
	Event *AbstractionAccountFactoryInitialized // Event containing the contract specifics and raw log

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
func (it *AbstractionAccountFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbstractionAccountFactoryInitialized)
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
		it.Event = new(AbstractionAccountFactoryInitialized)
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
func (it *AbstractionAccountFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbstractionAccountFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbstractionAccountFactoryInitialized represents a Initialized event raised by the AbstractionAccountFactory contract.
type AbstractionAccountFactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AbstractionAccountFactory *AbstractionAccountFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*AbstractionAccountFactoryInitializedIterator, error) {

	logs, sub, err := _AbstractionAccountFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AbstractionAccountFactoryInitializedIterator{contract: _AbstractionAccountFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AbstractionAccountFactory *AbstractionAccountFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AbstractionAccountFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _AbstractionAccountFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbstractionAccountFactoryInitialized)
				if err := _AbstractionAccountFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_AbstractionAccountFactory *AbstractionAccountFactoryFilterer) ParseInitialized(log types.Log) (*AbstractionAccountFactoryInitialized, error) {
	event := new(AbstractionAccountFactoryInitialized)
	if err := _AbstractionAccountFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbstractionAccountFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AbstractionAccountFactory contract.
type AbstractionAccountFactoryOwnershipTransferredIterator struct {
	Event *AbstractionAccountFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AbstractionAccountFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbstractionAccountFactoryOwnershipTransferred)
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
		it.Event = new(AbstractionAccountFactoryOwnershipTransferred)
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
func (it *AbstractionAccountFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbstractionAccountFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbstractionAccountFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the AbstractionAccountFactory contract.
type AbstractionAccountFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AbstractionAccountFactory *AbstractionAccountFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AbstractionAccountFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AbstractionAccountFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AbstractionAccountFactoryOwnershipTransferredIterator{contract: _AbstractionAccountFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AbstractionAccountFactory *AbstractionAccountFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AbstractionAccountFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AbstractionAccountFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbstractionAccountFactoryOwnershipTransferred)
				if err := _AbstractionAccountFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AbstractionAccountFactory *AbstractionAccountFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*AbstractionAccountFactoryOwnershipTransferred, error) {
	event := new(AbstractionAccountFactoryOwnershipTransferred)
	if err := _AbstractionAccountFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbstractionAccountFactoryAbstractionAccountDeployedIterator is returned from FilterAbstractionAccountDeployed and is used to iterate over the raw logs and unpacked data for AbstractionAccountDeployed events raised by the AbstractionAccountFactory contract.
type AbstractionAccountFactoryAbstractionAccountDeployedIterator struct {
	Event *AbstractionAccountFactoryAbstractionAccountDeployed // Event containing the contract specifics and raw log

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
func (it *AbstractionAccountFactoryAbstractionAccountDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbstractionAccountFactoryAbstractionAccountDeployed)
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
		it.Event = new(AbstractionAccountFactoryAbstractionAccountDeployed)
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
func (it *AbstractionAccountFactoryAbstractionAccountDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbstractionAccountFactoryAbstractionAccountDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbstractionAccountFactoryAbstractionAccountDeployed represents a AbstractionAccountDeployed event raised by the AbstractionAccountFactory contract.
type AbstractionAccountFactoryAbstractionAccountDeployed struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAbstractionAccountDeployed is a free log retrieval operation binding the contract event 0xf8078a05821fe3a0f57304f01a507627f803be4212ef8a593a28a395b43763de.
//
// Solidity: event AbstractionAccountDeployed(address indexed account)
func (_AbstractionAccountFactory *AbstractionAccountFactoryFilterer) FilterAbstractionAccountDeployed(opts *bind.FilterOpts, account []common.Address) (*AbstractionAccountFactoryAbstractionAccountDeployedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AbstractionAccountFactory.contract.FilterLogs(opts, "AbstractionAccountDeployed", accountRule)
	if err != nil {
		return nil, err
	}
	return &AbstractionAccountFactoryAbstractionAccountDeployedIterator{contract: _AbstractionAccountFactory.contract, event: "AbstractionAccountDeployed", logs: logs, sub: sub}, nil
}

// WatchAbstractionAccountDeployed is a free log subscription operation binding the contract event 0xf8078a05821fe3a0f57304f01a507627f803be4212ef8a593a28a395b43763de.
//
// Solidity: event AbstractionAccountDeployed(address indexed account)
func (_AbstractionAccountFactory *AbstractionAccountFactoryFilterer) WatchAbstractionAccountDeployed(opts *bind.WatchOpts, sink chan<- *AbstractionAccountFactoryAbstractionAccountDeployed, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _AbstractionAccountFactory.contract.WatchLogs(opts, "AbstractionAccountDeployed", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbstractionAccountFactoryAbstractionAccountDeployed)
				if err := _AbstractionAccountFactory.contract.UnpackLog(event, "AbstractionAccountDeployed", log); err != nil {
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

// ParseAbstractionAccountDeployed is a log parse operation binding the contract event 0xf8078a05821fe3a0f57304f01a507627f803be4212ef8a593a28a395b43763de.
//
// Solidity: event AbstractionAccountDeployed(address indexed account)
func (_AbstractionAccountFactory *AbstractionAccountFactoryFilterer) ParseAbstractionAccountDeployed(log types.Log) (*AbstractionAccountFactoryAbstractionAccountDeployed, error) {
	event := new(AbstractionAccountFactoryAbstractionAccountDeployed)
	if err := _AbstractionAccountFactory.contract.UnpackLog(event, "AbstractionAccountDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AbstractionAccountFactoryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AbstractionAccountFactory contract.
type AbstractionAccountFactoryUpgradedIterator struct {
	Event *AbstractionAccountFactoryUpgraded // Event containing the contract specifics and raw log

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
func (it *AbstractionAccountFactoryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AbstractionAccountFactoryUpgraded)
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
		it.Event = new(AbstractionAccountFactoryUpgraded)
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
func (it *AbstractionAccountFactoryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AbstractionAccountFactoryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AbstractionAccountFactoryUpgraded represents a Upgraded event raised by the AbstractionAccountFactory contract.
type AbstractionAccountFactoryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AbstractionAccountFactory *AbstractionAccountFactoryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AbstractionAccountFactoryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AbstractionAccountFactory.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AbstractionAccountFactoryUpgradedIterator{contract: _AbstractionAccountFactory.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AbstractionAccountFactory *AbstractionAccountFactoryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AbstractionAccountFactoryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AbstractionAccountFactory.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AbstractionAccountFactoryUpgraded)
				if err := _AbstractionAccountFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_AbstractionAccountFactory *AbstractionAccountFactoryFilterer) ParseUpgraded(log types.Log) (*AbstractionAccountFactoryUpgraded, error) {
	event := new(AbstractionAccountFactoryUpgraded)
	if err := _AbstractionAccountFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
