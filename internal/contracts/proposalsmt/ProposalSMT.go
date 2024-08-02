// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proposalsmt

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

// SparseMerkleTreeNode is an auto generated low-level Go binding around an user-defined struct.
type SparseMerkleTreeNode struct {
	NodeType   uint8
	ChildLeft  uint64
	ChildRight uint64
	NodeHash   [32]byte
	Key        [32]byte
	Value      [32]byte
}

// SparseMerkleTreeProof is an auto generated low-level Go binding around an user-defined struct.
type SparseMerkleTreeProof struct {
	Root         [32]byte
	Siblings     [][32]byte
	Existence    bool
	Key          [32]byte
	Value        [32]byte
	AuxExistence bool
	AuxKey       [32]byte
	AuxValue     [32]byte
}

// ProposalSMTMetaData contains all meta data concerning the ProposalSMT contract.
var ProposalSMTMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"RootUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"TREE_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposalsState_\",\"type\":\"address\"}],\"name\":\"__ProposalSMT_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyOfElement_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"element_\",\"type\":\"bytes32\"}],\"name\":\"add\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key_\",\"type\":\"bytes32\"}],\"name\":\"getNodeByKey\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSparseMerkleTree.NodeType\",\"name\":\"nodeType\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"childLeft\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"childRight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structSparseMerkleTree.Node\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key_\",\"type\":\"bytes32\"}],\"name\":\"getProof\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"auxKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auxValue\",\"type\":\"bytes32\"}],\"internalType\":\"structSparseMerkleTree.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalsState\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyOfElement_\",\"type\":\"bytes32\"}],\"name\":\"remove\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyOfElement_\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"newElement_\",\"type\":\"bytes32\"}],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ProposalSMTABI is the input ABI used to generate the binding from.
// Deprecated: Use ProposalSMTMetaData.ABI instead.
var ProposalSMTABI = ProposalSMTMetaData.ABI

// ProposalSMT is an auto generated Go binding around an Ethereum contract.
type ProposalSMT struct {
	ProposalSMTCaller     // Read-only binding to the contract
	ProposalSMTTransactor // Write-only binding to the contract
	ProposalSMTFilterer   // Log filterer for contract events
}

// ProposalSMTCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProposalSMTCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalSMTTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProposalSMTTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalSMTFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProposalSMTFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalSMTSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProposalSMTSession struct {
	Contract     *ProposalSMT      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProposalSMTCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProposalSMTCallerSession struct {
	Contract *ProposalSMTCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ProposalSMTTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProposalSMTTransactorSession struct {
	Contract     *ProposalSMTTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ProposalSMTRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProposalSMTRaw struct {
	Contract *ProposalSMT // Generic contract binding to access the raw methods on
}

// ProposalSMTCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProposalSMTCallerRaw struct {
	Contract *ProposalSMTCaller // Generic read-only contract binding to access the raw methods on
}

// ProposalSMTTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProposalSMTTransactorRaw struct {
	Contract *ProposalSMTTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProposalSMT creates a new instance of ProposalSMT, bound to a specific deployed contract.
func NewProposalSMT(address common.Address, backend bind.ContractBackend) (*ProposalSMT, error) {
	contract, err := bindProposalSMT(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProposalSMT{ProposalSMTCaller: ProposalSMTCaller{contract: contract}, ProposalSMTTransactor: ProposalSMTTransactor{contract: contract}, ProposalSMTFilterer: ProposalSMTFilterer{contract: contract}}, nil
}

// NewProposalSMTCaller creates a new read-only instance of ProposalSMT, bound to a specific deployed contract.
func NewProposalSMTCaller(address common.Address, caller bind.ContractCaller) (*ProposalSMTCaller, error) {
	contract, err := bindProposalSMT(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalSMTCaller{contract: contract}, nil
}

// NewProposalSMTTransactor creates a new write-only instance of ProposalSMT, bound to a specific deployed contract.
func NewProposalSMTTransactor(address common.Address, transactor bind.ContractTransactor) (*ProposalSMTTransactor, error) {
	contract, err := bindProposalSMT(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalSMTTransactor{contract: contract}, nil
}

// NewProposalSMTFilterer creates a new log filterer instance of ProposalSMT, bound to a specific deployed contract.
func NewProposalSMTFilterer(address common.Address, filterer bind.ContractFilterer) (*ProposalSMTFilterer, error) {
	contract, err := bindProposalSMT(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProposalSMTFilterer{contract: contract}, nil
}

// bindProposalSMT binds a generic wrapper to an already deployed contract.
func bindProposalSMT(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProposalSMTMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProposalSMT *ProposalSMTRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProposalSMT.Contract.ProposalSMTCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProposalSMT *ProposalSMTRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposalSMT.Contract.ProposalSMTTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProposalSMT *ProposalSMTRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProposalSMT.Contract.ProposalSMTTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProposalSMT *ProposalSMTCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProposalSMT.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProposalSMT *ProposalSMTTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposalSMT.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProposalSMT *ProposalSMTTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProposalSMT.Contract.contract.Transact(opts, method, params...)
}

// TREESIZE is a free data retrieval call binding the contract method 0xf9232cb4.
//
// Solidity: function TREE_SIZE() view returns(uint256)
func (_ProposalSMT *ProposalSMTCaller) TREESIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ProposalSMT.contract.Call(opts, &out, "TREE_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TREESIZE is a free data retrieval call binding the contract method 0xf9232cb4.
//
// Solidity: function TREE_SIZE() view returns(uint256)
func (_ProposalSMT *ProposalSMTSession) TREESIZE() (*big.Int, error) {
	return _ProposalSMT.Contract.TREESIZE(&_ProposalSMT.CallOpts)
}

// TREESIZE is a free data retrieval call binding the contract method 0xf9232cb4.
//
// Solidity: function TREE_SIZE() view returns(uint256)
func (_ProposalSMT *ProposalSMTCallerSession) TREESIZE() (*big.Int, error) {
	return _ProposalSMT.Contract.TREESIZE(&_ProposalSMT.CallOpts)
}

// GetNodeByKey is a free data retrieval call binding the contract method 0x083a8580.
//
// Solidity: function getNodeByKey(bytes32 key_) view returns((uint8,uint64,uint64,bytes32,bytes32,bytes32))
func (_ProposalSMT *ProposalSMTCaller) GetNodeByKey(opts *bind.CallOpts, key_ [32]byte) (SparseMerkleTreeNode, error) {
	var out []interface{}
	err := _ProposalSMT.contract.Call(opts, &out, "getNodeByKey", key_)

	if err != nil {
		return *new(SparseMerkleTreeNode), err
	}

	out0 := *abi.ConvertType(out[0], new(SparseMerkleTreeNode)).(*SparseMerkleTreeNode)

	return out0, err

}

// GetNodeByKey is a free data retrieval call binding the contract method 0x083a8580.
//
// Solidity: function getNodeByKey(bytes32 key_) view returns((uint8,uint64,uint64,bytes32,bytes32,bytes32))
func (_ProposalSMT *ProposalSMTSession) GetNodeByKey(key_ [32]byte) (SparseMerkleTreeNode, error) {
	return _ProposalSMT.Contract.GetNodeByKey(&_ProposalSMT.CallOpts, key_)
}

// GetNodeByKey is a free data retrieval call binding the contract method 0x083a8580.
//
// Solidity: function getNodeByKey(bytes32 key_) view returns((uint8,uint64,uint64,bytes32,bytes32,bytes32))
func (_ProposalSMT *ProposalSMTCallerSession) GetNodeByKey(key_ [32]byte) (SparseMerkleTreeNode, error) {
	return _ProposalSMT.Contract.GetNodeByKey(&_ProposalSMT.CallOpts, key_)
}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 key_) view returns((bytes32,bytes32[],bool,bytes32,bytes32,bool,bytes32,bytes32))
func (_ProposalSMT *ProposalSMTCaller) GetProof(opts *bind.CallOpts, key_ [32]byte) (SparseMerkleTreeProof, error) {
	var out []interface{}
	err := _ProposalSMT.contract.Call(opts, &out, "getProof", key_)

	if err != nil {
		return *new(SparseMerkleTreeProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SparseMerkleTreeProof)).(*SparseMerkleTreeProof)

	return out0, err

}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 key_) view returns((bytes32,bytes32[],bool,bytes32,bytes32,bool,bytes32,bytes32))
func (_ProposalSMT *ProposalSMTSession) GetProof(key_ [32]byte) (SparseMerkleTreeProof, error) {
	return _ProposalSMT.Contract.GetProof(&_ProposalSMT.CallOpts, key_)
}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 key_) view returns((bytes32,bytes32[],bool,bytes32,bytes32,bool,bytes32,bytes32))
func (_ProposalSMT *ProposalSMTCallerSession) GetProof(key_ [32]byte) (SparseMerkleTreeProof, error) {
	return _ProposalSMT.Contract.GetProof(&_ProposalSMT.CallOpts, key_)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_ProposalSMT *ProposalSMTCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ProposalSMT.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_ProposalSMT *ProposalSMTSession) GetRoot() ([32]byte, error) {
	return _ProposalSMT.Contract.GetRoot(&_ProposalSMT.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_ProposalSMT *ProposalSMTCallerSession) GetRoot() ([32]byte, error) {
	return _ProposalSMT.Contract.GetRoot(&_ProposalSMT.CallOpts)
}

// ProposalsState is a free data retrieval call binding the contract method 0x3af4e407.
//
// Solidity: function proposalsState() view returns(address)
func (_ProposalSMT *ProposalSMTCaller) ProposalsState(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ProposalSMT.contract.Call(opts, &out, "proposalsState")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalsState is a free data retrieval call binding the contract method 0x3af4e407.
//
// Solidity: function proposalsState() view returns(address)
func (_ProposalSMT *ProposalSMTSession) ProposalsState() (common.Address, error) {
	return _ProposalSMT.Contract.ProposalsState(&_ProposalSMT.CallOpts)
}

// ProposalsState is a free data retrieval call binding the contract method 0x3af4e407.
//
// Solidity: function proposalsState() view returns(address)
func (_ProposalSMT *ProposalSMTCallerSession) ProposalsState() (common.Address, error) {
	return _ProposalSMT.Contract.ProposalsState(&_ProposalSMT.CallOpts)
}

// ProposalSMTInit is a paid mutator transaction binding the contract method 0x4662d2f7.
//
// Solidity: function __ProposalSMT_init(address proposalsState_) returns()
func (_ProposalSMT *ProposalSMTTransactor) ProposalSMTInit(opts *bind.TransactOpts, proposalsState_ common.Address) (*types.Transaction, error) {
	return _ProposalSMT.contract.Transact(opts, "__ProposalSMT_init", proposalsState_)
}

// ProposalSMTInit is a paid mutator transaction binding the contract method 0x4662d2f7.
//
// Solidity: function __ProposalSMT_init(address proposalsState_) returns()
func (_ProposalSMT *ProposalSMTSession) ProposalSMTInit(proposalsState_ common.Address) (*types.Transaction, error) {
	return _ProposalSMT.Contract.ProposalSMTInit(&_ProposalSMT.TransactOpts, proposalsState_)
}

// ProposalSMTInit is a paid mutator transaction binding the contract method 0x4662d2f7.
//
// Solidity: function __ProposalSMT_init(address proposalsState_) returns()
func (_ProposalSMT *ProposalSMTTransactorSession) ProposalSMTInit(proposalsState_ common.Address) (*types.Transaction, error) {
	return _ProposalSMT.Contract.ProposalSMTInit(&_ProposalSMT.TransactOpts, proposalsState_)
}

// Add is a paid mutator transaction binding the contract method 0xd1de592a.
//
// Solidity: function add(bytes32 keyOfElement_, bytes32 element_) returns()
func (_ProposalSMT *ProposalSMTTransactor) Add(opts *bind.TransactOpts, keyOfElement_ [32]byte, element_ [32]byte) (*types.Transaction, error) {
	return _ProposalSMT.contract.Transact(opts, "add", keyOfElement_, element_)
}

// Add is a paid mutator transaction binding the contract method 0xd1de592a.
//
// Solidity: function add(bytes32 keyOfElement_, bytes32 element_) returns()
func (_ProposalSMT *ProposalSMTSession) Add(keyOfElement_ [32]byte, element_ [32]byte) (*types.Transaction, error) {
	return _ProposalSMT.Contract.Add(&_ProposalSMT.TransactOpts, keyOfElement_, element_)
}

// Add is a paid mutator transaction binding the contract method 0xd1de592a.
//
// Solidity: function add(bytes32 keyOfElement_, bytes32 element_) returns()
func (_ProposalSMT *ProposalSMTTransactorSession) Add(keyOfElement_ [32]byte, element_ [32]byte) (*types.Transaction, error) {
	return _ProposalSMT.Contract.Add(&_ProposalSMT.TransactOpts, keyOfElement_, element_)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 keyOfElement_) returns()
func (_ProposalSMT *ProposalSMTTransactor) Remove(opts *bind.TransactOpts, keyOfElement_ [32]byte) (*types.Transaction, error) {
	return _ProposalSMT.contract.Transact(opts, "remove", keyOfElement_)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 keyOfElement_) returns()
func (_ProposalSMT *ProposalSMTSession) Remove(keyOfElement_ [32]byte) (*types.Transaction, error) {
	return _ProposalSMT.Contract.Remove(&_ProposalSMT.TransactOpts, keyOfElement_)
}

// Remove is a paid mutator transaction binding the contract method 0x95bc2673.
//
// Solidity: function remove(bytes32 keyOfElement_) returns()
func (_ProposalSMT *ProposalSMTTransactorSession) Remove(keyOfElement_ [32]byte) (*types.Transaction, error) {
	return _ProposalSMT.Contract.Remove(&_ProposalSMT.TransactOpts, keyOfElement_)
}

// Update is a paid mutator transaction binding the contract method 0x13f57c3e.
//
// Solidity: function update(bytes32 keyOfElement_, bytes32 newElement_) returns()
func (_ProposalSMT *ProposalSMTTransactor) Update(opts *bind.TransactOpts, keyOfElement_ [32]byte, newElement_ [32]byte) (*types.Transaction, error) {
	return _ProposalSMT.contract.Transact(opts, "update", keyOfElement_, newElement_)
}

// Update is a paid mutator transaction binding the contract method 0x13f57c3e.
//
// Solidity: function update(bytes32 keyOfElement_, bytes32 newElement_) returns()
func (_ProposalSMT *ProposalSMTSession) Update(keyOfElement_ [32]byte, newElement_ [32]byte) (*types.Transaction, error) {
	return _ProposalSMT.Contract.Update(&_ProposalSMT.TransactOpts, keyOfElement_, newElement_)
}

// Update is a paid mutator transaction binding the contract method 0x13f57c3e.
//
// Solidity: function update(bytes32 keyOfElement_, bytes32 newElement_) returns()
func (_ProposalSMT *ProposalSMTTransactorSession) Update(keyOfElement_ [32]byte, newElement_ [32]byte) (*types.Transaction, error) {
	return _ProposalSMT.Contract.Update(&_ProposalSMT.TransactOpts, keyOfElement_, newElement_)
}

// ProposalSMTInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ProposalSMT contract.
type ProposalSMTInitializedIterator struct {
	Event *ProposalSMTInitialized // Event containing the contract specifics and raw log

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
func (it *ProposalSMTInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalSMTInitialized)
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
		it.Event = new(ProposalSMTInitialized)
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
func (it *ProposalSMTInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalSMTInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalSMTInitialized represents a Initialized event raised by the ProposalSMT contract.
type ProposalSMTInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ProposalSMT *ProposalSMTFilterer) FilterInitialized(opts *bind.FilterOpts) (*ProposalSMTInitializedIterator, error) {

	logs, sub, err := _ProposalSMT.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ProposalSMTInitializedIterator{contract: _ProposalSMT.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ProposalSMT *ProposalSMTFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ProposalSMTInitialized) (event.Subscription, error) {

	logs, sub, err := _ProposalSMT.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalSMTInitialized)
				if err := _ProposalSMT.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_ProposalSMT *ProposalSMTFilterer) ParseInitialized(log types.Log) (*ProposalSMTInitialized, error) {
	event := new(ProposalSMTInitialized)
	if err := _ProposalSMT.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalSMTRootUpdatedIterator is returned from FilterRootUpdated and is used to iterate over the raw logs and unpacked data for RootUpdated events raised by the ProposalSMT contract.
type ProposalSMTRootUpdatedIterator struct {
	Event *ProposalSMTRootUpdated // Event containing the contract specifics and raw log

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
func (it *ProposalSMTRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalSMTRootUpdated)
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
		it.Event = new(ProposalSMTRootUpdated)
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
func (it *ProposalSMTRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalSMTRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalSMTRootUpdated represents a RootUpdated event raised by the ProposalSMT contract.
type ProposalSMTRootUpdated struct {
	Root [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRootUpdated is a free log retrieval operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 indexed root)
func (_ProposalSMT *ProposalSMTFilterer) FilterRootUpdated(opts *bind.FilterOpts, root [][32]byte) (*ProposalSMTRootUpdatedIterator, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _ProposalSMT.contract.FilterLogs(opts, "RootUpdated", rootRule)
	if err != nil {
		return nil, err
	}
	return &ProposalSMTRootUpdatedIterator{contract: _ProposalSMT.contract, event: "RootUpdated", logs: logs, sub: sub}, nil
}

// WatchRootUpdated is a free log subscription operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 indexed root)
func (_ProposalSMT *ProposalSMTFilterer) WatchRootUpdated(opts *bind.WatchOpts, sink chan<- *ProposalSMTRootUpdated, root [][32]byte) (event.Subscription, error) {

	var rootRule []interface{}
	for _, rootItem := range root {
		rootRule = append(rootRule, rootItem)
	}

	logs, sub, err := _ProposalSMT.contract.WatchLogs(opts, "RootUpdated", rootRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalSMTRootUpdated)
				if err := _ProposalSMT.contract.UnpackLog(event, "RootUpdated", log); err != nil {
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

// ParseRootUpdated is a log parse operation binding the contract event 0x2cbc14f49c068133583f7cb530018af451c87c1cf1327cf2a4ff4698c4730aa4.
//
// Solidity: event RootUpdated(bytes32 indexed root)
func (_ProposalSMT *ProposalSMTFilterer) ParseRootUpdated(log types.Log) (*ProposalSMTRootUpdated, error) {
	event := new(ProposalSMTRootUpdated)
	if err := _ProposalSMT.contract.UnpackLog(event, "RootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
