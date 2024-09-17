// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package points

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

// DiamondFacet is an auto generated low-level Go binding around an user-defined struct.
type DiamondFacet struct {
	FacetAddress      common.Address
	Action            uint8
	FunctionSelectors [][4]byte
}

// DiamondStorageFacetInfo is an auto generated low-level Go binding around an user-defined struct.
type DiamondStorageFacetInfo struct {
	FacetAddress      common.Address
	FunctionSelectors [][4]byte
}

// PointsMetaData contains all meta data concerning the Points contract.
var PointsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumDiamond.FacetAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"indexed\":false,\"internalType\":\"structDiamond.Facet[]\",\"name\":\"facets\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"initFacet\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"initData\",\"type\":\"bytes\"}],\"name\":\"DiamondCut\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"storageSlot\",\"type\":\"bytes32\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"AGENT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURN_SELECTOR\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DIAMOND_ACCESS_CONTROL_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DIAMOND_ERC20_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DIAMOND_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EIP712_DIAMOND_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FORCED_TRANSFER_SELECTOR\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINT_SELECTOR\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POINTS_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECOVERY_SELECTOR\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_FROM_SELECTOR\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_SELECTOR\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TYPE_HASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"regulatoryCompliance_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"kycCompliance_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initRegulatory_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"initKYC_\",\"type\":\"bytes\"}],\"name\":\"__Points_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender_\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"checkRole\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumDiamond.FacetAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structDiamond.Facet[]\",\"name\":\"modules_\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"initModule_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"initData_\",\"type\":\"bytes\"}],\"name\":\"diamondCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"enumDiamond.FacetAction\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structDiamond.Facet[]\",\"name\":\"modules_\",\"type\":\"tuple[]\"}],\"name\":\"diamondCut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector_\",\"type\":\"bytes4\"}],\"name\":\"facetAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facet_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facetAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facets_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"facet_\",\"type\":\"address\"}],\"name\":\"facetFunctionSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"selectors_\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"facets\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"facetAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"functionSelectors\",\"type\":\"bytes4[]\"}],\"internalType\":\"structDiamondStorage.FacetInfo[]\",\"name\":\"facets_\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"forcedTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role_\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oldAccount_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAccount_\",\"type\":\"address\"}],\"name\":\"recovery\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role_\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account_\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PointsABI is the input ABI used to generate the binding from.
// Deprecated: Use PointsMetaData.ABI instead.
var PointsABI = PointsMetaData.ABI

// Points is an auto generated Go binding around an Ethereum contract.
type Points struct {
	PointsCaller     // Read-only binding to the contract
	PointsTransactor // Write-only binding to the contract
	PointsFilterer   // Log filterer for contract events
}

// PointsCaller is an auto generated read-only Go binding around an Ethereum contract.
type PointsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PointsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PointsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PointsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PointsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PointsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PointsSession struct {
	Contract     *Points           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PointsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PointsCallerSession struct {
	Contract *PointsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PointsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PointsTransactorSession struct {
	Contract     *PointsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PointsRaw is an auto generated low-level Go binding around an Ethereum contract.
type PointsRaw struct {
	Contract *Points // Generic contract binding to access the raw methods on
}

// PointsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PointsCallerRaw struct {
	Contract *PointsCaller // Generic read-only contract binding to access the raw methods on
}

// PointsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PointsTransactorRaw struct {
	Contract *PointsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoints creates a new instance of Points, bound to a specific deployed contract.
func NewPoints(address common.Address, backend bind.ContractBackend) (*Points, error) {
	contract, err := bindPoints(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Points{PointsCaller: PointsCaller{contract: contract}, PointsTransactor: PointsTransactor{contract: contract}, PointsFilterer: PointsFilterer{contract: contract}}, nil
}

// NewPointsCaller creates a new read-only instance of Points, bound to a specific deployed contract.
func NewPointsCaller(address common.Address, caller bind.ContractCaller) (*PointsCaller, error) {
	contract, err := bindPoints(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PointsCaller{contract: contract}, nil
}

// NewPointsTransactor creates a new write-only instance of Points, bound to a specific deployed contract.
func NewPointsTransactor(address common.Address, transactor bind.ContractTransactor) (*PointsTransactor, error) {
	contract, err := bindPoints(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PointsTransactor{contract: contract}, nil
}

// NewPointsFilterer creates a new log filterer instance of Points, bound to a specific deployed contract.
func NewPointsFilterer(address common.Address, filterer bind.ContractFilterer) (*PointsFilterer, error) {
	contract, err := bindPoints(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PointsFilterer{contract: contract}, nil
}

// bindPoints binds a generic wrapper to an already deployed contract.
func bindPoints(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PointsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Points *PointsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Points.Contract.PointsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Points *PointsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Points.Contract.PointsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Points *PointsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Points.Contract.PointsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Points *PointsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Points.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Points *PointsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Points.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Points *PointsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Points.Contract.contract.Transact(opts, method, params...)
}

// AGENTROLE is a free data retrieval call binding the contract method 0x22459e18.
//
// Solidity: function AGENT_ROLE() view returns(bytes32)
func (_Points *PointsCaller) AGENTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "AGENT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AGENTROLE is a free data retrieval call binding the contract method 0x22459e18.
//
// Solidity: function AGENT_ROLE() view returns(bytes32)
func (_Points *PointsSession) AGENTROLE() ([32]byte, error) {
	return _Points.Contract.AGENTROLE(&_Points.CallOpts)
}

// AGENTROLE is a free data retrieval call binding the contract method 0x22459e18.
//
// Solidity: function AGENT_ROLE() view returns(bytes32)
func (_Points *PointsCallerSession) AGENTROLE() ([32]byte, error) {
	return _Points.Contract.AGENTROLE(&_Points.CallOpts)
}

// BURNSELECTOR is a free data retrieval call binding the contract method 0x87a131bb.
//
// Solidity: function BURN_SELECTOR() view returns(bytes4)
func (_Points *PointsCaller) BURNSELECTOR(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "BURN_SELECTOR")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// BURNSELECTOR is a free data retrieval call binding the contract method 0x87a131bb.
//
// Solidity: function BURN_SELECTOR() view returns(bytes4)
func (_Points *PointsSession) BURNSELECTOR() ([4]byte, error) {
	return _Points.Contract.BURNSELECTOR(&_Points.CallOpts)
}

// BURNSELECTOR is a free data retrieval call binding the contract method 0x87a131bb.
//
// Solidity: function BURN_SELECTOR() view returns(bytes4)
func (_Points *PointsCallerSession) BURNSELECTOR() ([4]byte, error) {
	return _Points.Contract.BURNSELECTOR(&_Points.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Points *PointsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Points *PointsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Points.Contract.DEFAULTADMINROLE(&_Points.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Points *PointsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Points.Contract.DEFAULTADMINROLE(&_Points.CallOpts)
}

// DIAMONDACCESSCONTROLSTORAGESLOT is a free data retrieval call binding the contract method 0xb29b3bc0.
//
// Solidity: function DIAMOND_ACCESS_CONTROL_STORAGE_SLOT() view returns(bytes32)
func (_Points *PointsCaller) DIAMONDACCESSCONTROLSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "DIAMOND_ACCESS_CONTROL_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DIAMONDACCESSCONTROLSTORAGESLOT is a free data retrieval call binding the contract method 0xb29b3bc0.
//
// Solidity: function DIAMOND_ACCESS_CONTROL_STORAGE_SLOT() view returns(bytes32)
func (_Points *PointsSession) DIAMONDACCESSCONTROLSTORAGESLOT() ([32]byte, error) {
	return _Points.Contract.DIAMONDACCESSCONTROLSTORAGESLOT(&_Points.CallOpts)
}

// DIAMONDACCESSCONTROLSTORAGESLOT is a free data retrieval call binding the contract method 0xb29b3bc0.
//
// Solidity: function DIAMOND_ACCESS_CONTROL_STORAGE_SLOT() view returns(bytes32)
func (_Points *PointsCallerSession) DIAMONDACCESSCONTROLSTORAGESLOT() ([32]byte, error) {
	return _Points.Contract.DIAMONDACCESSCONTROLSTORAGESLOT(&_Points.CallOpts)
}

// DIAMONDERC20STORAGESLOT is a free data retrieval call binding the contract method 0x6812cb2a.
//
// Solidity: function DIAMOND_ERC20_STORAGE_SLOT() view returns(bytes32)
func (_Points *PointsCaller) DIAMONDERC20STORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "DIAMOND_ERC20_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DIAMONDERC20STORAGESLOT is a free data retrieval call binding the contract method 0x6812cb2a.
//
// Solidity: function DIAMOND_ERC20_STORAGE_SLOT() view returns(bytes32)
func (_Points *PointsSession) DIAMONDERC20STORAGESLOT() ([32]byte, error) {
	return _Points.Contract.DIAMONDERC20STORAGESLOT(&_Points.CallOpts)
}

// DIAMONDERC20STORAGESLOT is a free data retrieval call binding the contract method 0x6812cb2a.
//
// Solidity: function DIAMOND_ERC20_STORAGE_SLOT() view returns(bytes32)
func (_Points *PointsCallerSession) DIAMONDERC20STORAGESLOT() ([32]byte, error) {
	return _Points.Contract.DIAMONDERC20STORAGESLOT(&_Points.CallOpts)
}

// DIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe828c51f.
//
// Solidity: function DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_Points *PointsCaller) DIAMONDSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "DIAMOND_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe828c51f.
//
// Solidity: function DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_Points *PointsSession) DIAMONDSTORAGESLOT() ([32]byte, error) {
	return _Points.Contract.DIAMONDSTORAGESLOT(&_Points.CallOpts)
}

// DIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe828c51f.
//
// Solidity: function DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_Points *PointsCallerSession) DIAMONDSTORAGESLOT() ([32]byte, error) {
	return _Points.Contract.DIAMONDSTORAGESLOT(&_Points.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Points *PointsCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Points *PointsSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Points.Contract.DOMAINSEPARATOR(&_Points.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Points *PointsCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Points.Contract.DOMAINSEPARATOR(&_Points.CallOpts)
}

// EIP712DIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0x20cb691b.
//
// Solidity: function EIP712_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_Points *PointsCaller) EIP712DIAMONDSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "EIP712_DIAMOND_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0x20cb691b.
//
// Solidity: function EIP712_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_Points *PointsSession) EIP712DIAMONDSTORAGESLOT() ([32]byte, error) {
	return _Points.Contract.EIP712DIAMONDSTORAGESLOT(&_Points.CallOpts)
}

// EIP712DIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0x20cb691b.
//
// Solidity: function EIP712_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_Points *PointsCallerSession) EIP712DIAMONDSTORAGESLOT() ([32]byte, error) {
	return _Points.Contract.EIP712DIAMONDSTORAGESLOT(&_Points.CallOpts)
}

// FORCEDTRANSFERSELECTOR is a free data retrieval call binding the contract method 0xc7ff5a47.
//
// Solidity: function FORCED_TRANSFER_SELECTOR() view returns(bytes4)
func (_Points *PointsCaller) FORCEDTRANSFERSELECTOR(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "FORCED_TRANSFER_SELECTOR")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// FORCEDTRANSFERSELECTOR is a free data retrieval call binding the contract method 0xc7ff5a47.
//
// Solidity: function FORCED_TRANSFER_SELECTOR() view returns(bytes4)
func (_Points *PointsSession) FORCEDTRANSFERSELECTOR() ([4]byte, error) {
	return _Points.Contract.FORCEDTRANSFERSELECTOR(&_Points.CallOpts)
}

// FORCEDTRANSFERSELECTOR is a free data retrieval call binding the contract method 0xc7ff5a47.
//
// Solidity: function FORCED_TRANSFER_SELECTOR() view returns(bytes4)
func (_Points *PointsCallerSession) FORCEDTRANSFERSELECTOR() ([4]byte, error) {
	return _Points.Contract.FORCEDTRANSFERSELECTOR(&_Points.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Points *PointsCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Points *PointsSession) MINTERROLE() ([32]byte, error) {
	return _Points.Contract.MINTERROLE(&_Points.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Points *PointsCallerSession) MINTERROLE() ([32]byte, error) {
	return _Points.Contract.MINTERROLE(&_Points.CallOpts)
}

// MINTSELECTOR is a free data retrieval call binding the contract method 0x544fe464.
//
// Solidity: function MINT_SELECTOR() view returns(bytes4)
func (_Points *PointsCaller) MINTSELECTOR(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "MINT_SELECTOR")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// MINTSELECTOR is a free data retrieval call binding the contract method 0x544fe464.
//
// Solidity: function MINT_SELECTOR() view returns(bytes4)
func (_Points *PointsSession) MINTSELECTOR() ([4]byte, error) {
	return _Points.Contract.MINTSELECTOR(&_Points.CallOpts)
}

// MINTSELECTOR is a free data retrieval call binding the contract method 0x544fe464.
//
// Solidity: function MINT_SELECTOR() view returns(bytes4)
func (_Points *PointsCallerSession) MINTSELECTOR() ([4]byte, error) {
	return _Points.Contract.MINTSELECTOR(&_Points.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Points *PointsCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Points *PointsSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Points.Contract.PERMITTYPEHASH(&_Points.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Points *PointsCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Points.Contract.PERMITTYPEHASH(&_Points.CallOpts)
}

// POINTSSTORAGESLOT is a free data retrieval call binding the contract method 0x616e3487.
//
// Solidity: function POINTS_STORAGE_SLOT() view returns(bytes32)
func (_Points *PointsCaller) POINTSSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "POINTS_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// POINTSSTORAGESLOT is a free data retrieval call binding the contract method 0x616e3487.
//
// Solidity: function POINTS_STORAGE_SLOT() view returns(bytes32)
func (_Points *PointsSession) POINTSSTORAGESLOT() ([32]byte, error) {
	return _Points.Contract.POINTSSTORAGESLOT(&_Points.CallOpts)
}

// POINTSSTORAGESLOT is a free data retrieval call binding the contract method 0x616e3487.
//
// Solidity: function POINTS_STORAGE_SLOT() view returns(bytes32)
func (_Points *PointsCallerSession) POINTSSTORAGESLOT() ([32]byte, error) {
	return _Points.Contract.POINTSSTORAGESLOT(&_Points.CallOpts)
}

// RECOVERYSELECTOR is a free data retrieval call binding the contract method 0xe4c38584.
//
// Solidity: function RECOVERY_SELECTOR() view returns(bytes4)
func (_Points *PointsCaller) RECOVERYSELECTOR(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "RECOVERY_SELECTOR")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// RECOVERYSELECTOR is a free data retrieval call binding the contract method 0xe4c38584.
//
// Solidity: function RECOVERY_SELECTOR() view returns(bytes4)
func (_Points *PointsSession) RECOVERYSELECTOR() ([4]byte, error) {
	return _Points.Contract.RECOVERYSELECTOR(&_Points.CallOpts)
}

// RECOVERYSELECTOR is a free data retrieval call binding the contract method 0xe4c38584.
//
// Solidity: function RECOVERY_SELECTOR() view returns(bytes4)
func (_Points *PointsCallerSession) RECOVERYSELECTOR() ([4]byte, error) {
	return _Points.Contract.RECOVERYSELECTOR(&_Points.CallOpts)
}

// TRANSFERFROMSELECTOR is a free data retrieval call binding the contract method 0xa53016a3.
//
// Solidity: function TRANSFER_FROM_SELECTOR() view returns(bytes4)
func (_Points *PointsCaller) TRANSFERFROMSELECTOR(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "TRANSFER_FROM_SELECTOR")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// TRANSFERFROMSELECTOR is a free data retrieval call binding the contract method 0xa53016a3.
//
// Solidity: function TRANSFER_FROM_SELECTOR() view returns(bytes4)
func (_Points *PointsSession) TRANSFERFROMSELECTOR() ([4]byte, error) {
	return _Points.Contract.TRANSFERFROMSELECTOR(&_Points.CallOpts)
}

// TRANSFERFROMSELECTOR is a free data retrieval call binding the contract method 0xa53016a3.
//
// Solidity: function TRANSFER_FROM_SELECTOR() view returns(bytes4)
func (_Points *PointsCallerSession) TRANSFERFROMSELECTOR() ([4]byte, error) {
	return _Points.Contract.TRANSFERFROMSELECTOR(&_Points.CallOpts)
}

// TRANSFERSELECTOR is a free data retrieval call binding the contract method 0x72e93225.
//
// Solidity: function TRANSFER_SELECTOR() view returns(bytes4)
func (_Points *PointsCaller) TRANSFERSELECTOR(opts *bind.CallOpts) ([4]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "TRANSFER_SELECTOR")

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// TRANSFERSELECTOR is a free data retrieval call binding the contract method 0x72e93225.
//
// Solidity: function TRANSFER_SELECTOR() view returns(bytes4)
func (_Points *PointsSession) TRANSFERSELECTOR() ([4]byte, error) {
	return _Points.Contract.TRANSFERSELECTOR(&_Points.CallOpts)
}

// TRANSFERSELECTOR is a free data retrieval call binding the contract method 0x72e93225.
//
// Solidity: function TRANSFER_SELECTOR() view returns(bytes4)
func (_Points *PointsCallerSession) TRANSFERSELECTOR() ([4]byte, error) {
	return _Points.Contract.TRANSFERSELECTOR(&_Points.CallOpts)
}

// TYPEHASH is a free data retrieval call binding the contract method 0x64d4c819.
//
// Solidity: function TYPE_HASH() view returns(bytes32)
func (_Points *PointsCaller) TYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "TYPE_HASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TYPEHASH is a free data retrieval call binding the contract method 0x64d4c819.
//
// Solidity: function TYPE_HASH() view returns(bytes32)
func (_Points *PointsSession) TYPEHASH() ([32]byte, error) {
	return _Points.Contract.TYPEHASH(&_Points.CallOpts)
}

// TYPEHASH is a free data retrieval call binding the contract method 0x64d4c819.
//
// Solidity: function TYPE_HASH() view returns(bytes32)
func (_Points *PointsCallerSession) TYPEHASH() ([32]byte, error) {
	return _Points.Contract.TYPEHASH(&_Points.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Points *PointsCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Points *PointsSession) UPGRADERROLE() ([32]byte, error) {
	return _Points.Contract.UPGRADERROLE(&_Points.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Points *PointsCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _Points.Contract.UPGRADERROLE(&_Points.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner_, address spender_) view returns(uint256)
func (_Points *PointsCaller) Allowance(opts *bind.CallOpts, owner_ common.Address, spender_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "allowance", owner_, spender_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner_, address spender_) view returns(uint256)
func (_Points *PointsSession) Allowance(owner_ common.Address, spender_ common.Address) (*big.Int, error) {
	return _Points.Contract.Allowance(&_Points.CallOpts, owner_, spender_)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner_, address spender_) view returns(uint256)
func (_Points *PointsCallerSession) Allowance(owner_ common.Address, spender_ common.Address) (*big.Int, error) {
	return _Points.Contract.Allowance(&_Points.CallOpts, owner_, spender_)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account_) view returns(uint256)
func (_Points *PointsCaller) BalanceOf(opts *bind.CallOpts, account_ common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "balanceOf", account_)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account_) view returns(uint256)
func (_Points *PointsSession) BalanceOf(account_ common.Address) (*big.Int, error) {
	return _Points.Contract.BalanceOf(&_Points.CallOpts, account_)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account_) view returns(uint256)
func (_Points *PointsCallerSession) BalanceOf(account_ common.Address) (*big.Int, error) {
	return _Points.Contract.BalanceOf(&_Points.CallOpts, account_)
}

// CheckRole is a free data retrieval call binding the contract method 0x12d9a6ad.
//
// Solidity: function checkRole(bytes32 role_, address account_) view returns()
func (_Points *PointsCaller) CheckRole(opts *bind.CallOpts, role_ [32]byte, account_ common.Address) error {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "checkRole", role_, account_)

	if err != nil {
		return err
	}

	return err

}

// CheckRole is a free data retrieval call binding the contract method 0x12d9a6ad.
//
// Solidity: function checkRole(bytes32 role_, address account_) view returns()
func (_Points *PointsSession) CheckRole(role_ [32]byte, account_ common.Address) error {
	return _Points.Contract.CheckRole(&_Points.CallOpts, role_, account_)
}

// CheckRole is a free data retrieval call binding the contract method 0x12d9a6ad.
//
// Solidity: function checkRole(bytes32 role_, address account_) view returns()
func (_Points *PointsCallerSession) CheckRole(role_ [32]byte, account_ common.Address) error {
	return _Points.Contract.CheckRole(&_Points.CallOpts, role_, account_)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Points *PointsCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Points *PointsSession) Decimals() (uint8, error) {
	return _Points.Contract.Decimals(&_Points.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Points *PointsCallerSession) Decimals() (uint8, error) {
	return _Points.Contract.Decimals(&_Points.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Points *PointsCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Points *PointsSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Points.Contract.Eip712Domain(&_Points.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Points *PointsCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Points.Contract.Eip712Domain(&_Points.CallOpts)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 selector_) view returns(address facet_)
func (_Points *PointsCaller) FacetAddress(opts *bind.CallOpts, selector_ [4]byte) (common.Address, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "facetAddress", selector_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 selector_) view returns(address facet_)
func (_Points *PointsSession) FacetAddress(selector_ [4]byte) (common.Address, error) {
	return _Points.Contract.FacetAddress(&_Points.CallOpts, selector_)
}

// FacetAddress is a free data retrieval call binding the contract method 0xcdffacc6.
//
// Solidity: function facetAddress(bytes4 selector_) view returns(address facet_)
func (_Points *PointsCallerSession) FacetAddress(selector_ [4]byte) (common.Address, error) {
	return _Points.Contract.FacetAddress(&_Points.CallOpts, selector_)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facets_)
func (_Points *PointsCaller) FacetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "facetAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facets_)
func (_Points *PointsSession) FacetAddresses() ([]common.Address, error) {
	return _Points.Contract.FacetAddresses(&_Points.CallOpts)
}

// FacetAddresses is a free data retrieval call binding the contract method 0x52ef6b2c.
//
// Solidity: function facetAddresses() view returns(address[] facets_)
func (_Points *PointsCallerSession) FacetAddresses() ([]common.Address, error) {
	return _Points.Contract.FacetAddresses(&_Points.CallOpts)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address facet_) view returns(bytes4[] selectors_)
func (_Points *PointsCaller) FacetFunctionSelectors(opts *bind.CallOpts, facet_ common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "facetFunctionSelectors", facet_)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address facet_) view returns(bytes4[] selectors_)
func (_Points *PointsSession) FacetFunctionSelectors(facet_ common.Address) ([][4]byte, error) {
	return _Points.Contract.FacetFunctionSelectors(&_Points.CallOpts, facet_)
}

// FacetFunctionSelectors is a free data retrieval call binding the contract method 0xadfca15e.
//
// Solidity: function facetFunctionSelectors(address facet_) view returns(bytes4[] selectors_)
func (_Points *PointsCallerSession) FacetFunctionSelectors(facet_ common.Address) ([][4]byte, error) {
	return _Points.Contract.FacetFunctionSelectors(&_Points.CallOpts, facet_)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_Points *PointsCaller) Facets(opts *bind.CallOpts) ([]DiamondStorageFacetInfo, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "facets")

	if err != nil {
		return *new([]DiamondStorageFacetInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]DiamondStorageFacetInfo)).(*[]DiamondStorageFacetInfo)

	return out0, err

}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_Points *PointsSession) Facets() ([]DiamondStorageFacetInfo, error) {
	return _Points.Contract.Facets(&_Points.CallOpts)
}

// Facets is a free data retrieval call binding the contract method 0x7a0ed627.
//
// Solidity: function facets() view returns((address,bytes4[])[] facets_)
func (_Points *PointsCallerSession) Facets() ([]DiamondStorageFacetInfo, error) {
	return _Points.Contract.Facets(&_Points.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role_) view returns(bytes32)
func (_Points *PointsCaller) GetRoleAdmin(opts *bind.CallOpts, role_ [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "getRoleAdmin", role_)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role_) view returns(bytes32)
func (_Points *PointsSession) GetRoleAdmin(role_ [32]byte) ([32]byte, error) {
	return _Points.Contract.GetRoleAdmin(&_Points.CallOpts, role_)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role_) view returns(bytes32)
func (_Points *PointsCallerSession) GetRoleAdmin(role_ [32]byte) ([32]byte, error) {
	return _Points.Contract.GetRoleAdmin(&_Points.CallOpts, role_)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role_, address account_) view returns(bool)
func (_Points *PointsCaller) HasRole(opts *bind.CallOpts, role_ [32]byte, account_ common.Address) (bool, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "hasRole", role_, account_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role_, address account_) view returns(bool)
func (_Points *PointsSession) HasRole(role_ [32]byte, account_ common.Address) (bool, error) {
	return _Points.Contract.HasRole(&_Points.CallOpts, role_, account_)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role_, address account_) view returns(bool)
func (_Points *PointsCallerSession) HasRole(role_ [32]byte, account_ common.Address) (bool, error) {
	return _Points.Contract.HasRole(&_Points.CallOpts, role_, account_)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Points *PointsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Points *PointsSession) Name() (string, error) {
	return _Points.Contract.Name(&_Points.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Points *PointsCallerSession) Name() (string, error) {
	return _Points.Contract.Name(&_Points.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Points *PointsCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Points *PointsSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Points.Contract.Nonces(&_Points.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Points *PointsCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Points.Contract.Nonces(&_Points.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Points *PointsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Points *PointsSession) Symbol() (string, error) {
	return _Points.Contract.Symbol(&_Points.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Points *PointsCallerSession) Symbol() (string, error) {
	return _Points.Contract.Symbol(&_Points.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Points *PointsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Points.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Points *PointsSession) TotalSupply() (*big.Int, error) {
	return _Points.Contract.TotalSupply(&_Points.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Points *PointsCallerSession) TotalSupply() (*big.Int, error) {
	return _Points.Contract.TotalSupply(&_Points.CallOpts)
}

// PointsInit is a paid mutator transaction binding the contract method 0xa7fe5e30.
//
// Solidity: function __Points_init(string name_, string symbol_, uint8 decimals_, address regulatoryCompliance_, address kycCompliance_, bytes initRegulatory_, bytes initKYC_) returns()
func (_Points *PointsTransactor) PointsInit(opts *bind.TransactOpts, name_ string, symbol_ string, decimals_ uint8, regulatoryCompliance_ common.Address, kycCompliance_ common.Address, initRegulatory_ []byte, initKYC_ []byte) (*types.Transaction, error) {
	return _Points.contract.Transact(opts, "__Points_init", name_, symbol_, decimals_, regulatoryCompliance_, kycCompliance_, initRegulatory_, initKYC_)
}

// PointsInit is a paid mutator transaction binding the contract method 0xa7fe5e30.
//
// Solidity: function __Points_init(string name_, string symbol_, uint8 decimals_, address regulatoryCompliance_, address kycCompliance_, bytes initRegulatory_, bytes initKYC_) returns()
func (_Points *PointsSession) PointsInit(name_ string, symbol_ string, decimals_ uint8, regulatoryCompliance_ common.Address, kycCompliance_ common.Address, initRegulatory_ []byte, initKYC_ []byte) (*types.Transaction, error) {
	return _Points.Contract.PointsInit(&_Points.TransactOpts, name_, symbol_, decimals_, regulatoryCompliance_, kycCompliance_, initRegulatory_, initKYC_)
}

// PointsInit is a paid mutator transaction binding the contract method 0xa7fe5e30.
//
// Solidity: function __Points_init(string name_, string symbol_, uint8 decimals_, address regulatoryCompliance_, address kycCompliance_, bytes initRegulatory_, bytes initKYC_) returns()
func (_Points *PointsTransactorSession) PointsInit(name_ string, symbol_ string, decimals_ uint8, regulatoryCompliance_ common.Address, kycCompliance_ common.Address, initRegulatory_ []byte, initKYC_ []byte) (*types.Transaction, error) {
	return _Points.Contract.PointsInit(&_Points.TransactOpts, name_, symbol_, decimals_, regulatoryCompliance_, kycCompliance_, initRegulatory_, initKYC_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender_, uint256 amount_) returns(bool)
func (_Points *PointsTransactor) Approve(opts *bind.TransactOpts, spender_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Points.contract.Transact(opts, "approve", spender_, amount_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender_, uint256 amount_) returns(bool)
func (_Points *PointsSession) Approve(spender_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Points.Contract.Approve(&_Points.TransactOpts, spender_, amount_)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender_, uint256 amount_) returns(bool)
func (_Points *PointsTransactorSession) Approve(spender_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Points.Contract.Approve(&_Points.TransactOpts, spender_, amount_)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account_, uint256 amount_) returns(bool)
func (_Points *PointsTransactor) Burn(opts *bind.TransactOpts, account_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Points.contract.Transact(opts, "burn", account_, amount_)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account_, uint256 amount_) returns(bool)
func (_Points *PointsSession) Burn(account_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Points.Contract.Burn(&_Points.TransactOpts, account_, amount_)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account_, uint256 amount_) returns(bool)
func (_Points *PointsTransactorSession) Burn(account_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Points.Contract.Burn(&_Points.TransactOpts, account_, amount_)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] modules_, address initModule_, bytes initData_) returns()
func (_Points *PointsTransactor) DiamondCut(opts *bind.TransactOpts, modules_ []DiamondFacet, initModule_ common.Address, initData_ []byte) (*types.Transaction, error) {
	return _Points.contract.Transact(opts, "diamondCut", modules_, initModule_, initData_)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] modules_, address initModule_, bytes initData_) returns()
func (_Points *PointsSession) DiamondCut(modules_ []DiamondFacet, initModule_ common.Address, initData_ []byte) (*types.Transaction, error) {
	return _Points.Contract.DiamondCut(&_Points.TransactOpts, modules_, initModule_, initData_)
}

// DiamondCut is a paid mutator transaction binding the contract method 0x1f931c1c.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] modules_, address initModule_, bytes initData_) returns()
func (_Points *PointsTransactorSession) DiamondCut(modules_ []DiamondFacet, initModule_ common.Address, initData_ []byte) (*types.Transaction, error) {
	return _Points.Contract.DiamondCut(&_Points.TransactOpts, modules_, initModule_, initData_)
}

// DiamondCut0 is a paid mutator transaction binding the contract method 0xe57e69c6.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] modules_) returns()
func (_Points *PointsTransactor) DiamondCut0(opts *bind.TransactOpts, modules_ []DiamondFacet) (*types.Transaction, error) {
	return _Points.contract.Transact(opts, "diamondCut0", modules_)
}

// DiamondCut0 is a paid mutator transaction binding the contract method 0xe57e69c6.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] modules_) returns()
func (_Points *PointsSession) DiamondCut0(modules_ []DiamondFacet) (*types.Transaction, error) {
	return _Points.Contract.DiamondCut0(&_Points.TransactOpts, modules_)
}

// DiamondCut0 is a paid mutator transaction binding the contract method 0xe57e69c6.
//
// Solidity: function diamondCut((address,uint8,bytes4[])[] modules_) returns()
func (_Points *PointsTransactorSession) DiamondCut0(modules_ []DiamondFacet) (*types.Transaction, error) {
	return _Points.Contract.DiamondCut0(&_Points.TransactOpts, modules_)
}

// ForcedTransfer is a paid mutator transaction binding the contract method 0x9fc1d0e7.
//
// Solidity: function forcedTransfer(address from_, address to_, uint256 amount_) returns(bool)
func (_Points *PointsTransactor) ForcedTransfer(opts *bind.TransactOpts, from_ common.Address, to_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Points.contract.Transact(opts, "forcedTransfer", from_, to_, amount_)
}

// ForcedTransfer is a paid mutator transaction binding the contract method 0x9fc1d0e7.
//
// Solidity: function forcedTransfer(address from_, address to_, uint256 amount_) returns(bool)
func (_Points *PointsSession) ForcedTransfer(from_ common.Address, to_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Points.Contract.ForcedTransfer(&_Points.TransactOpts, from_, to_, amount_)
}

// ForcedTransfer is a paid mutator transaction binding the contract method 0x9fc1d0e7.
//
// Solidity: function forcedTransfer(address from_, address to_, uint256 amount_) returns(bool)
func (_Points *PointsTransactorSession) ForcedTransfer(from_ common.Address, to_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Points.Contract.ForcedTransfer(&_Points.TransactOpts, from_, to_, amount_)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role_, address account_) returns()
func (_Points *PointsTransactor) GrantRole(opts *bind.TransactOpts, role_ [32]byte, account_ common.Address) (*types.Transaction, error) {
	return _Points.contract.Transact(opts, "grantRole", role_, account_)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role_, address account_) returns()
func (_Points *PointsSession) GrantRole(role_ [32]byte, account_ common.Address) (*types.Transaction, error) {
	return _Points.Contract.GrantRole(&_Points.TransactOpts, role_, account_)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role_, address account_) returns()
func (_Points *PointsTransactorSession) GrantRole(role_ [32]byte, account_ common.Address) (*types.Transaction, error) {
	return _Points.Contract.GrantRole(&_Points.TransactOpts, role_, account_)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account_, uint256 amount_) returns(bool)
func (_Points *PointsTransactor) Mint(opts *bind.TransactOpts, account_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Points.contract.Transact(opts, "mint", account_, amount_)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account_, uint256 amount_) returns(bool)
func (_Points *PointsSession) Mint(account_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Points.Contract.Mint(&_Points.TransactOpts, account_, amount_)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account_, uint256 amount_) returns(bool)
func (_Points *PointsTransactorSession) Mint(account_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Points.Contract.Mint(&_Points.TransactOpts, account_, amount_)
}

// Permit is a paid mutator transaction binding the contract method 0x9fd5a6cf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, bytes signature) returns()
func (_Points *PointsTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _Points.contract.Transact(opts, "permit", owner, spender, value, deadline, signature)
}

// Permit is a paid mutator transaction binding the contract method 0x9fd5a6cf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, bytes signature) returns()
func (_Points *PointsSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _Points.Contract.Permit(&_Points.TransactOpts, owner, spender, value, deadline, signature)
}

// Permit is a paid mutator transaction binding the contract method 0x9fd5a6cf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, bytes signature) returns()
func (_Points *PointsTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, signature []byte) (*types.Transaction, error) {
	return _Points.Contract.Permit(&_Points.TransactOpts, owner, spender, value, deadline, signature)
}

// Recovery is a paid mutator transaction binding the contract method 0x8d785674.
//
// Solidity: function recovery(address oldAccount_, address newAccount_) returns(bool)
func (_Points *PointsTransactor) Recovery(opts *bind.TransactOpts, oldAccount_ common.Address, newAccount_ common.Address) (*types.Transaction, error) {
	return _Points.contract.Transact(opts, "recovery", oldAccount_, newAccount_)
}

// Recovery is a paid mutator transaction binding the contract method 0x8d785674.
//
// Solidity: function recovery(address oldAccount_, address newAccount_) returns(bool)
func (_Points *PointsSession) Recovery(oldAccount_ common.Address, newAccount_ common.Address) (*types.Transaction, error) {
	return _Points.Contract.Recovery(&_Points.TransactOpts, oldAccount_, newAccount_)
}

// Recovery is a paid mutator transaction binding the contract method 0x8d785674.
//
// Solidity: function recovery(address oldAccount_, address newAccount_) returns(bool)
func (_Points *PointsTransactorSession) Recovery(oldAccount_ common.Address, newAccount_ common.Address) (*types.Transaction, error) {
	return _Points.Contract.Recovery(&_Points.TransactOpts, oldAccount_, newAccount_)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role_, address account_) returns()
func (_Points *PointsTransactor) RenounceRole(opts *bind.TransactOpts, role_ [32]byte, account_ common.Address) (*types.Transaction, error) {
	return _Points.contract.Transact(opts, "renounceRole", role_, account_)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role_, address account_) returns()
func (_Points *PointsSession) RenounceRole(role_ [32]byte, account_ common.Address) (*types.Transaction, error) {
	return _Points.Contract.RenounceRole(&_Points.TransactOpts, role_, account_)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role_, address account_) returns()
func (_Points *PointsTransactorSession) RenounceRole(role_ [32]byte, account_ common.Address) (*types.Transaction, error) {
	return _Points.Contract.RenounceRole(&_Points.TransactOpts, role_, account_)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role_, address account_) returns()
func (_Points *PointsTransactor) RevokeRole(opts *bind.TransactOpts, role_ [32]byte, account_ common.Address) (*types.Transaction, error) {
	return _Points.contract.Transact(opts, "revokeRole", role_, account_)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role_, address account_) returns()
func (_Points *PointsSession) RevokeRole(role_ [32]byte, account_ common.Address) (*types.Transaction, error) {
	return _Points.Contract.RevokeRole(&_Points.TransactOpts, role_, account_)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role_, address account_) returns()
func (_Points *PointsTransactorSession) RevokeRole(role_ [32]byte, account_ common.Address) (*types.Transaction, error) {
	return _Points.Contract.RevokeRole(&_Points.TransactOpts, role_, account_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to_, uint256 amount_) returns(bool)
func (_Points *PointsTransactor) Transfer(opts *bind.TransactOpts, to_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Points.contract.Transact(opts, "transfer", to_, amount_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to_, uint256 amount_) returns(bool)
func (_Points *PointsSession) Transfer(to_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Points.Contract.Transfer(&_Points.TransactOpts, to_, amount_)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to_, uint256 amount_) returns(bool)
func (_Points *PointsTransactorSession) Transfer(to_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Points.Contract.Transfer(&_Points.TransactOpts, to_, amount_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from_, address to_, uint256 amount_) returns(bool)
func (_Points *PointsTransactor) TransferFrom(opts *bind.TransactOpts, from_ common.Address, to_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Points.contract.Transact(opts, "transferFrom", from_, to_, amount_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from_, address to_, uint256 amount_) returns(bool)
func (_Points *PointsSession) TransferFrom(from_ common.Address, to_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Points.Contract.TransferFrom(&_Points.TransactOpts, from_, to_, amount_)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from_, address to_, uint256 amount_) returns(bool)
func (_Points *PointsTransactorSession) TransferFrom(from_ common.Address, to_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _Points.Contract.TransferFrom(&_Points.TransactOpts, from_, to_, amount_)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Points *PointsTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Points.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Points *PointsSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Points.Contract.Fallback(&_Points.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Points *PointsTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Points.Contract.Fallback(&_Points.TransactOpts, calldata)
}

// PointsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Points contract.
type PointsApprovalIterator struct {
	Event *PointsApproval // Event containing the contract specifics and raw log

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
func (it *PointsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PointsApproval)
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
		it.Event = new(PointsApproval)
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
func (it *PointsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PointsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PointsApproval represents a Approval event raised by the Points contract.
type PointsApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Points *PointsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PointsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Points.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PointsApprovalIterator{contract: _Points.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Points *PointsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PointsApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Points.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PointsApproval)
				if err := _Points.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Points *PointsFilterer) ParseApproval(log types.Log) (*PointsApproval, error) {
	event := new(PointsApproval)
	if err := _Points.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PointsDiamondCutIterator is returned from FilterDiamondCut and is used to iterate over the raw logs and unpacked data for DiamondCut events raised by the Points contract.
type PointsDiamondCutIterator struct {
	Event *PointsDiamondCut // Event containing the contract specifics and raw log

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
func (it *PointsDiamondCutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PointsDiamondCut)
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
		it.Event = new(PointsDiamondCut)
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
func (it *PointsDiamondCutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PointsDiamondCutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PointsDiamondCut represents a DiamondCut event raised by the Points contract.
type PointsDiamondCut struct {
	Facets    []DiamondFacet
	InitFacet common.Address
	InitData  []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDiamondCut is a free log retrieval operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] facets, address initFacet, bytes initData)
func (_Points *PointsFilterer) FilterDiamondCut(opts *bind.FilterOpts) (*PointsDiamondCutIterator, error) {

	logs, sub, err := _Points.contract.FilterLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return &PointsDiamondCutIterator{contract: _Points.contract, event: "DiamondCut", logs: logs, sub: sub}, nil
}

// WatchDiamondCut is a free log subscription operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] facets, address initFacet, bytes initData)
func (_Points *PointsFilterer) WatchDiamondCut(opts *bind.WatchOpts, sink chan<- *PointsDiamondCut) (event.Subscription, error) {

	logs, sub, err := _Points.contract.WatchLogs(opts, "DiamondCut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PointsDiamondCut)
				if err := _Points.contract.UnpackLog(event, "DiamondCut", log); err != nil {
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

// ParseDiamondCut is a log parse operation binding the contract event 0x8faa70878671ccd212d20771b795c50af8fd3ff6cf27f4bde57e5d4de0aeb673.
//
// Solidity: event DiamondCut((address,uint8,bytes4[])[] facets, address initFacet, bytes initData)
func (_Points *PointsFilterer) ParseDiamondCut(log types.Log) (*PointsDiamondCut, error) {
	event := new(PointsDiamondCut)
	if err := _Points.contract.UnpackLog(event, "DiamondCut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PointsEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Points contract.
type PointsEIP712DomainChangedIterator struct {
	Event *PointsEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *PointsEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PointsEIP712DomainChanged)
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
		it.Event = new(PointsEIP712DomainChanged)
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
func (it *PointsEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PointsEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PointsEIP712DomainChanged represents a EIP712DomainChanged event raised by the Points contract.
type PointsEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Points *PointsFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*PointsEIP712DomainChangedIterator, error) {

	logs, sub, err := _Points.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &PointsEIP712DomainChangedIterator{contract: _Points.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Points *PointsFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *PointsEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Points.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PointsEIP712DomainChanged)
				if err := _Points.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Points *PointsFilterer) ParseEIP712DomainChanged(log types.Log) (*PointsEIP712DomainChanged, error) {
	event := new(PointsEIP712DomainChanged)
	if err := _Points.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PointsInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Points contract.
type PointsInitializedIterator struct {
	Event *PointsInitialized // Event containing the contract specifics and raw log

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
func (it *PointsInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PointsInitialized)
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
		it.Event = new(PointsInitialized)
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
func (it *PointsInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PointsInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PointsInitialized represents a Initialized event raised by the Points contract.
type PointsInitialized struct {
	StorageSlot [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xdc73717d728bcfa015e8117438a65319aa06e979ca324afa6e1ea645c28ea15d.
//
// Solidity: event Initialized(bytes32 storageSlot)
func (_Points *PointsFilterer) FilterInitialized(opts *bind.FilterOpts) (*PointsInitializedIterator, error) {

	logs, sub, err := _Points.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PointsInitializedIterator{contract: _Points.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xdc73717d728bcfa015e8117438a65319aa06e979ca324afa6e1ea645c28ea15d.
//
// Solidity: event Initialized(bytes32 storageSlot)
func (_Points *PointsFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PointsInitialized) (event.Subscription, error) {

	logs, sub, err := _Points.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PointsInitialized)
				if err := _Points.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xdc73717d728bcfa015e8117438a65319aa06e979ca324afa6e1ea645c28ea15d.
//
// Solidity: event Initialized(bytes32 storageSlot)
func (_Points *PointsFilterer) ParseInitialized(log types.Log) (*PointsInitialized, error) {
	event := new(PointsInitialized)
	if err := _Points.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PointsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Points contract.
type PointsRoleAdminChangedIterator struct {
	Event *PointsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *PointsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PointsRoleAdminChanged)
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
		it.Event = new(PointsRoleAdminChanged)
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
func (it *PointsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PointsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PointsRoleAdminChanged represents a RoleAdminChanged event raised by the Points contract.
type PointsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Points *PointsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*PointsRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Points.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &PointsRoleAdminChangedIterator{contract: _Points.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Points *PointsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *PointsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Points.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PointsRoleAdminChanged)
				if err := _Points.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Points *PointsFilterer) ParseRoleAdminChanged(log types.Log) (*PointsRoleAdminChanged, error) {
	event := new(PointsRoleAdminChanged)
	if err := _Points.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PointsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Points contract.
type PointsRoleGrantedIterator struct {
	Event *PointsRoleGranted // Event containing the contract specifics and raw log

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
func (it *PointsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PointsRoleGranted)
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
		it.Event = new(PointsRoleGranted)
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
func (it *PointsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PointsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PointsRoleGranted represents a RoleGranted event raised by the Points contract.
type PointsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Points *PointsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PointsRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Points.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PointsRoleGrantedIterator{contract: _Points.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Points *PointsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *PointsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Points.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PointsRoleGranted)
				if err := _Points.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Points *PointsFilterer) ParseRoleGranted(log types.Log) (*PointsRoleGranted, error) {
	event := new(PointsRoleGranted)
	if err := _Points.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PointsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Points contract.
type PointsRoleRevokedIterator struct {
	Event *PointsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *PointsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PointsRoleRevoked)
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
		it.Event = new(PointsRoleRevoked)
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
func (it *PointsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PointsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PointsRoleRevoked represents a RoleRevoked event raised by the Points contract.
type PointsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Points *PointsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*PointsRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Points.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PointsRoleRevokedIterator{contract: _Points.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Points *PointsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *PointsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Points.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PointsRoleRevoked)
				if err := _Points.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Points *PointsFilterer) ParseRoleRevoked(log types.Log) (*PointsRoleRevoked, error) {
	event := new(PointsRoleRevoked)
	if err := _Points.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PointsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Points contract.
type PointsTransferIterator struct {
	Event *PointsTransfer // Event containing the contract specifics and raw log

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
func (it *PointsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PointsTransfer)
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
		it.Event = new(PointsTransfer)
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
func (it *PointsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PointsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PointsTransfer represents a Transfer event raised by the Points contract.
type PointsTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Points *PointsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PointsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Points.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PointsTransferIterator{contract: _Points.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Points *PointsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PointsTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Points.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PointsTransfer)
				if err := _Points.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Points *PointsFilterer) ParseTransfer(log types.Log) (*PointsTransfer, error) {
	event := new(PointsTransfer)
	if err := _Points.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
