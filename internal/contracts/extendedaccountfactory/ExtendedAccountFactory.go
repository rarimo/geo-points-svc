// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package extendedaccountfactory

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

// ExtendedAccountFactoryMetaData contains all meta data concerning the ExtendedAccountFactory contract.
var ExtendedAccountFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AbstractAccountDeployed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"abstractAccountImplementation_\",\"type\":\"address\"}],\"name\":\"__AbstractAccountFactory_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"abstractAccountImplementation_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"negRiskAdapter_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ctfExchange_\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"points_\",\"type\":\"address\"},{\"internalType\":\"contractIERC1155\",\"name\":\"ctf_\",\"type\":\"address\"}],\"name\":\"__ExtendedAccountFactory_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"abstractAccounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ctf\",\"outputs\":[{\"internalType\":\"contractIERC1155\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ctfExchange\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nullifier_\",\"type\":\"bytes32\"}],\"name\":\"deployAbstractAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nullifier_\",\"type\":\"bytes32\"}],\"name\":\"getAbstractAccount\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAbstractAccountImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"negRiskAdapter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"points\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"nullifier_\",\"type\":\"bytes32\"}],\"name\":\"predictAbstractAccountAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"setAbstractAccountImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// ExtendedAccountFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use ExtendedAccountFactoryMetaData.ABI instead.
var ExtendedAccountFactoryABI = ExtendedAccountFactoryMetaData.ABI

// ExtendedAccountFactory is an auto generated Go binding around an Ethereum contract.
type ExtendedAccountFactory struct {
	ExtendedAccountFactoryCaller     // Read-only binding to the contract
	ExtendedAccountFactoryTransactor // Write-only binding to the contract
	ExtendedAccountFactoryFilterer   // Log filterer for contract events
}

// ExtendedAccountFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExtendedAccountFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtendedAccountFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExtendedAccountFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtendedAccountFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExtendedAccountFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExtendedAccountFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExtendedAccountFactorySession struct {
	Contract     *ExtendedAccountFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ExtendedAccountFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExtendedAccountFactoryCallerSession struct {
	Contract *ExtendedAccountFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ExtendedAccountFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExtendedAccountFactoryTransactorSession struct {
	Contract     *ExtendedAccountFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ExtendedAccountFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExtendedAccountFactoryRaw struct {
	Contract *ExtendedAccountFactory // Generic contract binding to access the raw methods on
}

// ExtendedAccountFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExtendedAccountFactoryCallerRaw struct {
	Contract *ExtendedAccountFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ExtendedAccountFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExtendedAccountFactoryTransactorRaw struct {
	Contract *ExtendedAccountFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExtendedAccountFactory creates a new instance of ExtendedAccountFactory, bound to a specific deployed contract.
func NewExtendedAccountFactory(address common.Address, backend bind.ContractBackend) (*ExtendedAccountFactory, error) {
	contract, err := bindExtendedAccountFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExtendedAccountFactory{ExtendedAccountFactoryCaller: ExtendedAccountFactoryCaller{contract: contract}, ExtendedAccountFactoryTransactor: ExtendedAccountFactoryTransactor{contract: contract}, ExtendedAccountFactoryFilterer: ExtendedAccountFactoryFilterer{contract: contract}}, nil
}

// NewExtendedAccountFactoryCaller creates a new read-only instance of ExtendedAccountFactory, bound to a specific deployed contract.
func NewExtendedAccountFactoryCaller(address common.Address, caller bind.ContractCaller) (*ExtendedAccountFactoryCaller, error) {
	contract, err := bindExtendedAccountFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExtendedAccountFactoryCaller{contract: contract}, nil
}

// NewExtendedAccountFactoryTransactor creates a new write-only instance of ExtendedAccountFactory, bound to a specific deployed contract.
func NewExtendedAccountFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ExtendedAccountFactoryTransactor, error) {
	contract, err := bindExtendedAccountFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExtendedAccountFactoryTransactor{contract: contract}, nil
}

// NewExtendedAccountFactoryFilterer creates a new log filterer instance of ExtendedAccountFactory, bound to a specific deployed contract.
func NewExtendedAccountFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*ExtendedAccountFactoryFilterer, error) {
	contract, err := bindExtendedAccountFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExtendedAccountFactoryFilterer{contract: contract}, nil
}

// bindExtendedAccountFactory binds a generic wrapper to an already deployed contract.
func bindExtendedAccountFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ExtendedAccountFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtendedAccountFactory *ExtendedAccountFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtendedAccountFactory.Contract.ExtendedAccountFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtendedAccountFactory *ExtendedAccountFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.ExtendedAccountFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtendedAccountFactory *ExtendedAccountFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.ExtendedAccountFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExtendedAccountFactory *ExtendedAccountFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ExtendedAccountFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExtendedAccountFactory *ExtendedAccountFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExtendedAccountFactory *ExtendedAccountFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.contract.Transact(opts, method, params...)
}

// AbstractAccounts is a free data retrieval call binding the contract method 0xf5c2d47b.
//
// Solidity: function abstractAccounts(bytes32 ) view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCaller) AbstractAccounts(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ExtendedAccountFactory.contract.Call(opts, &out, "abstractAccounts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AbstractAccounts is a free data retrieval call binding the contract method 0xf5c2d47b.
//
// Solidity: function abstractAccounts(bytes32 ) view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) AbstractAccounts(arg0 [32]byte) (common.Address, error) {
	return _ExtendedAccountFactory.Contract.AbstractAccounts(&_ExtendedAccountFactory.CallOpts, arg0)
}

// AbstractAccounts is a free data retrieval call binding the contract method 0xf5c2d47b.
//
// Solidity: function abstractAccounts(bytes32 ) view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCallerSession) AbstractAccounts(arg0 [32]byte) (common.Address, error) {
	return _ExtendedAccountFactory.Contract.AbstractAccounts(&_ExtendedAccountFactory.CallOpts, arg0)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCaller) Ctf(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExtendedAccountFactory.contract.Call(opts, &out, "ctf")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) Ctf() (common.Address, error) {
	return _ExtendedAccountFactory.Contract.Ctf(&_ExtendedAccountFactory.CallOpts)
}

// Ctf is a free data retrieval call binding the contract method 0x22a9339f.
//
// Solidity: function ctf() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCallerSession) Ctf() (common.Address, error) {
	return _ExtendedAccountFactory.Contract.Ctf(&_ExtendedAccountFactory.CallOpts)
}

// CtfExchange is a free data retrieval call binding the contract method 0x4b8295e7.
//
// Solidity: function ctfExchange() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCaller) CtfExchange(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExtendedAccountFactory.contract.Call(opts, &out, "ctfExchange")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CtfExchange is a free data retrieval call binding the contract method 0x4b8295e7.
//
// Solidity: function ctfExchange() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) CtfExchange() (common.Address, error) {
	return _ExtendedAccountFactory.Contract.CtfExchange(&_ExtendedAccountFactory.CallOpts)
}

// CtfExchange is a free data retrieval call binding the contract method 0x4b8295e7.
//
// Solidity: function ctfExchange() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCallerSession) CtfExchange() (common.Address, error) {
	return _ExtendedAccountFactory.Contract.CtfExchange(&_ExtendedAccountFactory.CallOpts)
}

// GetAbstractAccount is a free data retrieval call binding the contract method 0x433e0e30.
//
// Solidity: function getAbstractAccount(bytes32 nullifier_) view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCaller) GetAbstractAccount(opts *bind.CallOpts, nullifier_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ExtendedAccountFactory.contract.Call(opts, &out, "getAbstractAccount", nullifier_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAbstractAccount is a free data retrieval call binding the contract method 0x433e0e30.
//
// Solidity: function getAbstractAccount(bytes32 nullifier_) view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) GetAbstractAccount(nullifier_ [32]byte) (common.Address, error) {
	return _ExtendedAccountFactory.Contract.GetAbstractAccount(&_ExtendedAccountFactory.CallOpts, nullifier_)
}

// GetAbstractAccount is a free data retrieval call binding the contract method 0x433e0e30.
//
// Solidity: function getAbstractAccount(bytes32 nullifier_) view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCallerSession) GetAbstractAccount(nullifier_ [32]byte) (common.Address, error) {
	return _ExtendedAccountFactory.Contract.GetAbstractAccount(&_ExtendedAccountFactory.CallOpts, nullifier_)
}

// GetAbstractAccountImplementation is a free data retrieval call binding the contract method 0xb5ca3ec7.
//
// Solidity: function getAbstractAccountImplementation() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCaller) GetAbstractAccountImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExtendedAccountFactory.contract.Call(opts, &out, "getAbstractAccountImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAbstractAccountImplementation is a free data retrieval call binding the contract method 0xb5ca3ec7.
//
// Solidity: function getAbstractAccountImplementation() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) GetAbstractAccountImplementation() (common.Address, error) {
	return _ExtendedAccountFactory.Contract.GetAbstractAccountImplementation(&_ExtendedAccountFactory.CallOpts)
}

// GetAbstractAccountImplementation is a free data retrieval call binding the contract method 0xb5ca3ec7.
//
// Solidity: function getAbstractAccountImplementation() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCallerSession) GetAbstractAccountImplementation() (common.Address, error) {
	return _ExtendedAccountFactory.Contract.GetAbstractAccountImplementation(&_ExtendedAccountFactory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExtendedAccountFactory.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) Implementation() (common.Address, error) {
	return _ExtendedAccountFactory.Contract.Implementation(&_ExtendedAccountFactory.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCallerSession) Implementation() (common.Address, error) {
	return _ExtendedAccountFactory.Contract.Implementation(&_ExtendedAccountFactory.CallOpts)
}

// NegRiskAdapter is a free data retrieval call binding the contract method 0xf6ef95a1.
//
// Solidity: function negRiskAdapter() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCaller) NegRiskAdapter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExtendedAccountFactory.contract.Call(opts, &out, "negRiskAdapter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NegRiskAdapter is a free data retrieval call binding the contract method 0xf6ef95a1.
//
// Solidity: function negRiskAdapter() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) NegRiskAdapter() (common.Address, error) {
	return _ExtendedAccountFactory.Contract.NegRiskAdapter(&_ExtendedAccountFactory.CallOpts)
}

// NegRiskAdapter is a free data retrieval call binding the contract method 0xf6ef95a1.
//
// Solidity: function negRiskAdapter() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCallerSession) NegRiskAdapter() (common.Address, error) {
	return _ExtendedAccountFactory.Contract.NegRiskAdapter(&_ExtendedAccountFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExtendedAccountFactory.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) Owner() (common.Address, error) {
	return _ExtendedAccountFactory.Contract.Owner(&_ExtendedAccountFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCallerSession) Owner() (common.Address, error) {
	return _ExtendedAccountFactory.Contract.Owner(&_ExtendedAccountFactory.CallOpts)
}

// Points is a free data retrieval call binding the contract method 0x1be6dd64.
//
// Solidity: function points() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCaller) Points(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ExtendedAccountFactory.contract.Call(opts, &out, "points")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Points is a free data retrieval call binding the contract method 0x1be6dd64.
//
// Solidity: function points() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) Points() (common.Address, error) {
	return _ExtendedAccountFactory.Contract.Points(&_ExtendedAccountFactory.CallOpts)
}

// Points is a free data retrieval call binding the contract method 0x1be6dd64.
//
// Solidity: function points() view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCallerSession) Points() (common.Address, error) {
	return _ExtendedAccountFactory.Contract.Points(&_ExtendedAccountFactory.CallOpts)
}

// PredictAbstractAccountAddress is a free data retrieval call binding the contract method 0x54cd3610.
//
// Solidity: function predictAbstractAccountAddress(bytes32 nullifier_) view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCaller) PredictAbstractAccountAddress(opts *bind.CallOpts, nullifier_ [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ExtendedAccountFactory.contract.Call(opts, &out, "predictAbstractAccountAddress", nullifier_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PredictAbstractAccountAddress is a free data retrieval call binding the contract method 0x54cd3610.
//
// Solidity: function predictAbstractAccountAddress(bytes32 nullifier_) view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) PredictAbstractAccountAddress(nullifier_ [32]byte) (common.Address, error) {
	return _ExtendedAccountFactory.Contract.PredictAbstractAccountAddress(&_ExtendedAccountFactory.CallOpts, nullifier_)
}

// PredictAbstractAccountAddress is a free data retrieval call binding the contract method 0x54cd3610.
//
// Solidity: function predictAbstractAccountAddress(bytes32 nullifier_) view returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCallerSession) PredictAbstractAccountAddress(nullifier_ [32]byte) (common.Address, error) {
	return _ExtendedAccountFactory.Contract.PredictAbstractAccountAddress(&_ExtendedAccountFactory.CallOpts, nullifier_)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ExtendedAccountFactory.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) ProxiableUUID() ([32]byte, error) {
	return _ExtendedAccountFactory.Contract.ProxiableUUID(&_ExtendedAccountFactory.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_ExtendedAccountFactory *ExtendedAccountFactoryCallerSession) ProxiableUUID() ([32]byte, error) {
	return _ExtendedAccountFactory.Contract.ProxiableUUID(&_ExtendedAccountFactory.CallOpts)
}

// AbstractAccountFactoryInit is a paid mutator transaction binding the contract method 0x30b5f045.
//
// Solidity: function __AbstractAccountFactory_init(address abstractAccountImplementation_) returns()
func (_ExtendedAccountFactory *ExtendedAccountFactoryTransactor) AbstractAccountFactoryInit(opts *bind.TransactOpts, abstractAccountImplementation_ common.Address) (*types.Transaction, error) {
	return _ExtendedAccountFactory.contract.Transact(opts, "__AbstractAccountFactory_init", abstractAccountImplementation_)
}

// AbstractAccountFactoryInit is a paid mutator transaction binding the contract method 0x30b5f045.
//
// Solidity: function __AbstractAccountFactory_init(address abstractAccountImplementation_) returns()
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) AbstractAccountFactoryInit(abstractAccountImplementation_ common.Address) (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.AbstractAccountFactoryInit(&_ExtendedAccountFactory.TransactOpts, abstractAccountImplementation_)
}

// AbstractAccountFactoryInit is a paid mutator transaction binding the contract method 0x30b5f045.
//
// Solidity: function __AbstractAccountFactory_init(address abstractAccountImplementation_) returns()
func (_ExtendedAccountFactory *ExtendedAccountFactoryTransactorSession) AbstractAccountFactoryInit(abstractAccountImplementation_ common.Address) (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.AbstractAccountFactoryInit(&_ExtendedAccountFactory.TransactOpts, abstractAccountImplementation_)
}

// ExtendedAccountFactoryInit is a paid mutator transaction binding the contract method 0xf225d70f.
//
// Solidity: function __ExtendedAccountFactory_init(address abstractAccountImplementation_, address negRiskAdapter_, address ctfExchange_, address points_, address ctf_) returns()
func (_ExtendedAccountFactory *ExtendedAccountFactoryTransactor) ExtendedAccountFactoryInit(opts *bind.TransactOpts, abstractAccountImplementation_ common.Address, negRiskAdapter_ common.Address, ctfExchange_ common.Address, points_ common.Address, ctf_ common.Address) (*types.Transaction, error) {
	return _ExtendedAccountFactory.contract.Transact(opts, "__ExtendedAccountFactory_init", abstractAccountImplementation_, negRiskAdapter_, ctfExchange_, points_, ctf_)
}

// ExtendedAccountFactoryInit is a paid mutator transaction binding the contract method 0xf225d70f.
//
// Solidity: function __ExtendedAccountFactory_init(address abstractAccountImplementation_, address negRiskAdapter_, address ctfExchange_, address points_, address ctf_) returns()
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) ExtendedAccountFactoryInit(abstractAccountImplementation_ common.Address, negRiskAdapter_ common.Address, ctfExchange_ common.Address, points_ common.Address, ctf_ common.Address) (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.ExtendedAccountFactoryInit(&_ExtendedAccountFactory.TransactOpts, abstractAccountImplementation_, negRiskAdapter_, ctfExchange_, points_, ctf_)
}

// ExtendedAccountFactoryInit is a paid mutator transaction binding the contract method 0xf225d70f.
//
// Solidity: function __ExtendedAccountFactory_init(address abstractAccountImplementation_, address negRiskAdapter_, address ctfExchange_, address points_, address ctf_) returns()
func (_ExtendedAccountFactory *ExtendedAccountFactoryTransactorSession) ExtendedAccountFactoryInit(abstractAccountImplementation_ common.Address, negRiskAdapter_ common.Address, ctfExchange_ common.Address, points_ common.Address, ctf_ common.Address) (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.ExtendedAccountFactoryInit(&_ExtendedAccountFactory.TransactOpts, abstractAccountImplementation_, negRiskAdapter_, ctfExchange_, points_, ctf_)
}

// DeployAbstractAccount is a paid mutator transaction binding the contract method 0x37c74c08.
//
// Solidity: function deployAbstractAccount(bytes32 nullifier_) returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryTransactor) DeployAbstractAccount(opts *bind.TransactOpts, nullifier_ [32]byte) (*types.Transaction, error) {
	return _ExtendedAccountFactory.contract.Transact(opts, "deployAbstractAccount", nullifier_)
}

// DeployAbstractAccount is a paid mutator transaction binding the contract method 0x37c74c08.
//
// Solidity: function deployAbstractAccount(bytes32 nullifier_) returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) DeployAbstractAccount(nullifier_ [32]byte) (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.DeployAbstractAccount(&_ExtendedAccountFactory.TransactOpts, nullifier_)
}

// DeployAbstractAccount is a paid mutator transaction binding the contract method 0x37c74c08.
//
// Solidity: function deployAbstractAccount(bytes32 nullifier_) returns(address)
func (_ExtendedAccountFactory *ExtendedAccountFactoryTransactorSession) DeployAbstractAccount(nullifier_ [32]byte) (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.DeployAbstractAccount(&_ExtendedAccountFactory.TransactOpts, nullifier_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExtendedAccountFactory *ExtendedAccountFactoryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExtendedAccountFactory.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) RenounceOwnership() (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.RenounceOwnership(&_ExtendedAccountFactory.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ExtendedAccountFactory *ExtendedAccountFactoryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.RenounceOwnership(&_ExtendedAccountFactory.TransactOpts)
}

// SetAbstractAccountImplementation is a paid mutator transaction binding the contract method 0x87a0bfdb.
//
// Solidity: function setAbstractAccountImplementation(address newImplementation) returns()
func (_ExtendedAccountFactory *ExtendedAccountFactoryTransactor) SetAbstractAccountImplementation(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _ExtendedAccountFactory.contract.Transact(opts, "setAbstractAccountImplementation", newImplementation)
}

// SetAbstractAccountImplementation is a paid mutator transaction binding the contract method 0x87a0bfdb.
//
// Solidity: function setAbstractAccountImplementation(address newImplementation) returns()
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) SetAbstractAccountImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.SetAbstractAccountImplementation(&_ExtendedAccountFactory.TransactOpts, newImplementation)
}

// SetAbstractAccountImplementation is a paid mutator transaction binding the contract method 0x87a0bfdb.
//
// Solidity: function setAbstractAccountImplementation(address newImplementation) returns()
func (_ExtendedAccountFactory *ExtendedAccountFactoryTransactorSession) SetAbstractAccountImplementation(newImplementation common.Address) (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.SetAbstractAccountImplementation(&_ExtendedAccountFactory.TransactOpts, newImplementation)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExtendedAccountFactory *ExtendedAccountFactoryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ExtendedAccountFactory.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.TransferOwnership(&_ExtendedAccountFactory.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ExtendedAccountFactory *ExtendedAccountFactoryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.TransferOwnership(&_ExtendedAccountFactory.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ExtendedAccountFactory *ExtendedAccountFactoryTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _ExtendedAccountFactory.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.UpgradeTo(&_ExtendedAccountFactory.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_ExtendedAccountFactory *ExtendedAccountFactoryTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.UpgradeTo(&_ExtendedAccountFactory.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ExtendedAccountFactory *ExtendedAccountFactoryTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ExtendedAccountFactory.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ExtendedAccountFactory *ExtendedAccountFactorySession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.UpgradeToAndCall(&_ExtendedAccountFactory.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_ExtendedAccountFactory *ExtendedAccountFactoryTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _ExtendedAccountFactory.Contract.UpgradeToAndCall(&_ExtendedAccountFactory.TransactOpts, newImplementation, data)
}

// ExtendedAccountFactoryAbstractAccountDeployedIterator is returned from FilterAbstractAccountDeployed and is used to iterate over the raw logs and unpacked data for AbstractAccountDeployed events raised by the ExtendedAccountFactory contract.
type ExtendedAccountFactoryAbstractAccountDeployedIterator struct {
	Event *ExtendedAccountFactoryAbstractAccountDeployed // Event containing the contract specifics and raw log

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
func (it *ExtendedAccountFactoryAbstractAccountDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtendedAccountFactoryAbstractAccountDeployed)
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
		it.Event = new(ExtendedAccountFactoryAbstractAccountDeployed)
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
func (it *ExtendedAccountFactoryAbstractAccountDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtendedAccountFactoryAbstractAccountDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtendedAccountFactoryAbstractAccountDeployed represents a AbstractAccountDeployed event raised by the ExtendedAccountFactory contract.
type ExtendedAccountFactoryAbstractAccountDeployed struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAbstractAccountDeployed is a free log retrieval operation binding the contract event 0x04a8395ad3493e53ed150cbb37930a73625d14c0fe82374d1a7a58ceda5f6daa.
//
// Solidity: event AbstractAccountDeployed(address indexed account)
func (_ExtendedAccountFactory *ExtendedAccountFactoryFilterer) FilterAbstractAccountDeployed(opts *bind.FilterOpts, account []common.Address) (*ExtendedAccountFactoryAbstractAccountDeployedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ExtendedAccountFactory.contract.FilterLogs(opts, "AbstractAccountDeployed", accountRule)
	if err != nil {
		return nil, err
	}
	return &ExtendedAccountFactoryAbstractAccountDeployedIterator{contract: _ExtendedAccountFactory.contract, event: "AbstractAccountDeployed", logs: logs, sub: sub}, nil
}

// WatchAbstractAccountDeployed is a free log subscription operation binding the contract event 0x04a8395ad3493e53ed150cbb37930a73625d14c0fe82374d1a7a58ceda5f6daa.
//
// Solidity: event AbstractAccountDeployed(address indexed account)
func (_ExtendedAccountFactory *ExtendedAccountFactoryFilterer) WatchAbstractAccountDeployed(opts *bind.WatchOpts, sink chan<- *ExtendedAccountFactoryAbstractAccountDeployed, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _ExtendedAccountFactory.contract.WatchLogs(opts, "AbstractAccountDeployed", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtendedAccountFactoryAbstractAccountDeployed)
				if err := _ExtendedAccountFactory.contract.UnpackLog(event, "AbstractAccountDeployed", log); err != nil {
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

// ParseAbstractAccountDeployed is a log parse operation binding the contract event 0x04a8395ad3493e53ed150cbb37930a73625d14c0fe82374d1a7a58ceda5f6daa.
//
// Solidity: event AbstractAccountDeployed(address indexed account)
func (_ExtendedAccountFactory *ExtendedAccountFactoryFilterer) ParseAbstractAccountDeployed(log types.Log) (*ExtendedAccountFactoryAbstractAccountDeployed, error) {
	event := new(ExtendedAccountFactoryAbstractAccountDeployed)
	if err := _ExtendedAccountFactory.contract.UnpackLog(event, "AbstractAccountDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtendedAccountFactoryAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ExtendedAccountFactory contract.
type ExtendedAccountFactoryAdminChangedIterator struct {
	Event *ExtendedAccountFactoryAdminChanged // Event containing the contract specifics and raw log

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
func (it *ExtendedAccountFactoryAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtendedAccountFactoryAdminChanged)
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
		it.Event = new(ExtendedAccountFactoryAdminChanged)
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
func (it *ExtendedAccountFactoryAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtendedAccountFactoryAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtendedAccountFactoryAdminChanged represents a AdminChanged event raised by the ExtendedAccountFactory contract.
type ExtendedAccountFactoryAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ExtendedAccountFactory *ExtendedAccountFactoryFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ExtendedAccountFactoryAdminChangedIterator, error) {

	logs, sub, err := _ExtendedAccountFactory.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ExtendedAccountFactoryAdminChangedIterator{contract: _ExtendedAccountFactory.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ExtendedAccountFactory *ExtendedAccountFactoryFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ExtendedAccountFactoryAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ExtendedAccountFactory.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtendedAccountFactoryAdminChanged)
				if err := _ExtendedAccountFactory.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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
func (_ExtendedAccountFactory *ExtendedAccountFactoryFilterer) ParseAdminChanged(log types.Log) (*ExtendedAccountFactoryAdminChanged, error) {
	event := new(ExtendedAccountFactoryAdminChanged)
	if err := _ExtendedAccountFactory.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtendedAccountFactoryBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ExtendedAccountFactory contract.
type ExtendedAccountFactoryBeaconUpgradedIterator struct {
	Event *ExtendedAccountFactoryBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ExtendedAccountFactoryBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtendedAccountFactoryBeaconUpgraded)
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
		it.Event = new(ExtendedAccountFactoryBeaconUpgraded)
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
func (it *ExtendedAccountFactoryBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtendedAccountFactoryBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtendedAccountFactoryBeaconUpgraded represents a BeaconUpgraded event raised by the ExtendedAccountFactory contract.
type ExtendedAccountFactoryBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ExtendedAccountFactory *ExtendedAccountFactoryFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ExtendedAccountFactoryBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ExtendedAccountFactory.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ExtendedAccountFactoryBeaconUpgradedIterator{contract: _ExtendedAccountFactory.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ExtendedAccountFactory *ExtendedAccountFactoryFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ExtendedAccountFactoryBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ExtendedAccountFactory.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtendedAccountFactoryBeaconUpgraded)
				if err := _ExtendedAccountFactory.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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
func (_ExtendedAccountFactory *ExtendedAccountFactoryFilterer) ParseBeaconUpgraded(log types.Log) (*ExtendedAccountFactoryBeaconUpgraded, error) {
	event := new(ExtendedAccountFactoryBeaconUpgraded)
	if err := _ExtendedAccountFactory.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtendedAccountFactoryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ExtendedAccountFactory contract.
type ExtendedAccountFactoryInitializedIterator struct {
	Event *ExtendedAccountFactoryInitialized // Event containing the contract specifics and raw log

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
func (it *ExtendedAccountFactoryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtendedAccountFactoryInitialized)
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
		it.Event = new(ExtendedAccountFactoryInitialized)
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
func (it *ExtendedAccountFactoryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtendedAccountFactoryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtendedAccountFactoryInitialized represents a Initialized event raised by the ExtendedAccountFactory contract.
type ExtendedAccountFactoryInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ExtendedAccountFactory *ExtendedAccountFactoryFilterer) FilterInitialized(opts *bind.FilterOpts) (*ExtendedAccountFactoryInitializedIterator, error) {

	logs, sub, err := _ExtendedAccountFactory.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ExtendedAccountFactoryInitializedIterator{contract: _ExtendedAccountFactory.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ExtendedAccountFactory *ExtendedAccountFactoryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ExtendedAccountFactoryInitialized) (event.Subscription, error) {

	logs, sub, err := _ExtendedAccountFactory.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtendedAccountFactoryInitialized)
				if err := _ExtendedAccountFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ExtendedAccountFactory *ExtendedAccountFactoryFilterer) ParseInitialized(log types.Log) (*ExtendedAccountFactoryInitialized, error) {
	event := new(ExtendedAccountFactoryInitialized)
	if err := _ExtendedAccountFactory.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtendedAccountFactoryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ExtendedAccountFactory contract.
type ExtendedAccountFactoryOwnershipTransferredIterator struct {
	Event *ExtendedAccountFactoryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExtendedAccountFactoryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtendedAccountFactoryOwnershipTransferred)
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
		it.Event = new(ExtendedAccountFactoryOwnershipTransferred)
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
func (it *ExtendedAccountFactoryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtendedAccountFactoryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtendedAccountFactoryOwnershipTransferred represents a OwnershipTransferred event raised by the ExtendedAccountFactory contract.
type ExtendedAccountFactoryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExtendedAccountFactory *ExtendedAccountFactoryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExtendedAccountFactoryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExtendedAccountFactory.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExtendedAccountFactoryOwnershipTransferredIterator{contract: _ExtendedAccountFactory.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ExtendedAccountFactory *ExtendedAccountFactoryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExtendedAccountFactoryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ExtendedAccountFactory.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtendedAccountFactoryOwnershipTransferred)
				if err := _ExtendedAccountFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ExtendedAccountFactory *ExtendedAccountFactoryFilterer) ParseOwnershipTransferred(log types.Log) (*ExtendedAccountFactoryOwnershipTransferred, error) {
	event := new(ExtendedAccountFactoryOwnershipTransferred)
	if err := _ExtendedAccountFactory.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExtendedAccountFactoryUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ExtendedAccountFactory contract.
type ExtendedAccountFactoryUpgradedIterator struct {
	Event *ExtendedAccountFactoryUpgraded // Event containing the contract specifics and raw log

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
func (it *ExtendedAccountFactoryUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExtendedAccountFactoryUpgraded)
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
		it.Event = new(ExtendedAccountFactoryUpgraded)
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
func (it *ExtendedAccountFactoryUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExtendedAccountFactoryUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExtendedAccountFactoryUpgraded represents a Upgraded event raised by the ExtendedAccountFactory contract.
type ExtendedAccountFactoryUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ExtendedAccountFactory *ExtendedAccountFactoryFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ExtendedAccountFactoryUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ExtendedAccountFactory.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ExtendedAccountFactoryUpgradedIterator{contract: _ExtendedAccountFactory.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ExtendedAccountFactory *ExtendedAccountFactoryFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ExtendedAccountFactoryUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ExtendedAccountFactory.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExtendedAccountFactoryUpgraded)
				if err := _ExtendedAccountFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ExtendedAccountFactory *ExtendedAccountFactoryFilterer) ParseUpgraded(log types.Log) (*ExtendedAccountFactoryUpgraded, error) {
	event := new(ExtendedAccountFactoryUpgraded)
	if err := _ExtendedAccountFactory.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
