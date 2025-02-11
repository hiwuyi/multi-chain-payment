// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package goBind

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IPaymentMinimalTxInfo is an auto generated low-level Go binding around an user-defined struct.
type IPaymentMinimalTxInfo struct {
	Id         string
	Token      common.Address
	MinPayment *big.Int
	LockedFee  *big.Int
	Owner      common.Address
	Recipient  common.Address
	Deadline   *big.Int
	IsExisted  bool
	Size       *big.Int
}

// IPaymentMinimallockPaymentParam is an auto generated low-level Go binding around an user-defined struct.
type IPaymentMinimallockPaymentParam struct {
	Id         string
	MinPayment *big.Int
	Amount     *big.Int
	LockTime   *big.Int
	Recipient  common.Address
	Size       *big.Int
}

// IPaymentMinimalunlockPaymentParam is an auto generated low-level Go binding around an user-defined struct.
type IPaymentMinimalunlockPaymentParam struct {
	Id        string
	OrderId   string
	DealId    string
	Amount    *big.Int
	Recipient common.Address
}

// SwanPaymentABI is the input ABI used to generate the binding from.
const SwanPaymentABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ExpirePayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minPayment\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"LockPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"restToken\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"UnlockPayment\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"NATIVE_TOKEN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cId\",\"type\":\"string\"}],\"name\":\"getLockedPaymentInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockedFee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isExisted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"internalType\":\"structIPaymentMinimal.TxInfo\",\"name\":\"tx\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"ERC20_TOKEN\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"chainlinkOracle\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"internalType\":\"structIPaymentMinimal.lockPaymentParam\",\"name\":\"param\",\"type\":\"tuple\"}],\"name\":\"lockTokenPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string[]\",\"name\":\"cidList\",\"type\":\"string[]\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chainlinkOracle\",\"type\":\"address\"}],\"name\":\"setChainlinkOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceFeed\",\"type\":\"address\"}],\"name\":\"setPriceFeed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"unlockCarPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"dealId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structIPaymentMinimal.unlockPaymentParam\",\"name\":\"param\",\"type\":\"tuple\"}],\"name\":\"unlockTokenPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SwanPayment is an auto generated Go binding around an Ethereum contract.
type SwanPayment struct {
	SwanPaymentCaller     // Read-only binding to the contract
	SwanPaymentTransactor // Write-only binding to the contract
	SwanPaymentFilterer   // Log filterer for contract events
}

// SwanPaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwanPaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwanPaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwanPaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwanPaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwanPaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwanPaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwanPaymentSession struct {
	Contract     *SwanPayment      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwanPaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwanPaymentCallerSession struct {
	Contract *SwanPaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SwanPaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwanPaymentTransactorSession struct {
	Contract     *SwanPaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SwanPaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwanPaymentRaw struct {
	Contract *SwanPayment // Generic contract binding to access the raw methods on
}

// SwanPaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwanPaymentCallerRaw struct {
	Contract *SwanPaymentCaller // Generic read-only contract binding to access the raw methods on
}

// SwanPaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwanPaymentTransactorRaw struct {
	Contract *SwanPaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwanPayment creates a new instance of SwanPayment, bound to a specific deployed contract.
func NewSwanPayment(address common.Address, backend bind.ContractBackend) (*SwanPayment, error) {
	contract, err := bindSwanPayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwanPayment{SwanPaymentCaller: SwanPaymentCaller{contract: contract}, SwanPaymentTransactor: SwanPaymentTransactor{contract: contract}, SwanPaymentFilterer: SwanPaymentFilterer{contract: contract}}, nil
}

// NewSwanPaymentCaller creates a new read-only instance of SwanPayment, bound to a specific deployed contract.
func NewSwanPaymentCaller(address common.Address, caller bind.ContractCaller) (*SwanPaymentCaller, error) {
	contract, err := bindSwanPayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwanPaymentCaller{contract: contract}, nil
}

// NewSwanPaymentTransactor creates a new write-only instance of SwanPayment, bound to a specific deployed contract.
func NewSwanPaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*SwanPaymentTransactor, error) {
	contract, err := bindSwanPayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwanPaymentTransactor{contract: contract}, nil
}

// NewSwanPaymentFilterer creates a new log filterer instance of SwanPayment, bound to a specific deployed contract.
func NewSwanPaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*SwanPaymentFilterer, error) {
	contract, err := bindSwanPayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwanPaymentFilterer{contract: contract}, nil
}

// bindSwanPayment binds a generic wrapper to an already deployed contract.
func bindSwanPayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SwanPaymentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwanPayment *SwanPaymentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwanPayment.Contract.SwanPaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwanPayment *SwanPaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwanPayment.Contract.SwanPaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwanPayment *SwanPaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwanPayment.Contract.SwanPaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwanPayment *SwanPaymentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwanPayment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwanPayment *SwanPaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwanPayment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwanPayment *SwanPaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwanPayment.Contract.contract.Transact(opts, method, params...)
}

// NATIVETOKEN is a free data retrieval call binding the contract method 0x31f7d964.
//
// Solidity: function NATIVE_TOKEN() view returns(address)
func (_SwanPayment *SwanPaymentCaller) NATIVETOKEN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwanPayment.contract.Call(opts, &out, "NATIVE_TOKEN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NATIVETOKEN is a free data retrieval call binding the contract method 0x31f7d964.
//
// Solidity: function NATIVE_TOKEN() view returns(address)
func (_SwanPayment *SwanPaymentSession) NATIVETOKEN() (common.Address, error) {
	return _SwanPayment.Contract.NATIVETOKEN(&_SwanPayment.CallOpts)
}

// NATIVETOKEN is a free data retrieval call binding the contract method 0x31f7d964.
//
// Solidity: function NATIVE_TOKEN() view returns(address)
func (_SwanPayment *SwanPaymentCallerSession) NATIVETOKEN() (common.Address, error) {
	return _SwanPayment.Contract.NATIVETOKEN(&_SwanPayment.CallOpts)
}

// GetLockedPaymentInfo is a free data retrieval call binding the contract method 0xe063922b.
//
// Solidity: function getLockedPaymentInfo(string cId) view returns((string,address,uint256,uint256,address,address,uint256,bool,uint256) tx)
func (_SwanPayment *SwanPaymentCaller) GetLockedPaymentInfo(opts *bind.CallOpts, cId string) (IPaymentMinimalTxInfo, error) {
	var out []interface{}
	err := _SwanPayment.contract.Call(opts, &out, "getLockedPaymentInfo", cId)

	if err != nil {
		return *new(IPaymentMinimalTxInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IPaymentMinimalTxInfo)).(*IPaymentMinimalTxInfo)

	return out0, err

}

// GetLockedPaymentInfo is a free data retrieval call binding the contract method 0xe063922b.
//
// Solidity: function getLockedPaymentInfo(string cId) view returns((string,address,uint256,uint256,address,address,uint256,bool,uint256) tx)
func (_SwanPayment *SwanPaymentSession) GetLockedPaymentInfo(cId string) (IPaymentMinimalTxInfo, error) {
	return _SwanPayment.Contract.GetLockedPaymentInfo(&_SwanPayment.CallOpts, cId)
}

// GetLockedPaymentInfo is a free data retrieval call binding the contract method 0xe063922b.
//
// Solidity: function getLockedPaymentInfo(string cId) view returns((string,address,uint256,uint256,address,address,uint256,bool,uint256) tx)
func (_SwanPayment *SwanPaymentCallerSession) GetLockedPaymentInfo(cId string) (IPaymentMinimalTxInfo, error) {
	return _SwanPayment.Contract.GetLockedPaymentInfo(&_SwanPayment.CallOpts, cId)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address owner, address ERC20_TOKEN, address oracle, address priceFeed, address chainlinkOracle) returns()
func (_SwanPayment *SwanPaymentTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, ERC20_TOKEN common.Address, oracle common.Address, priceFeed common.Address, chainlinkOracle common.Address) (*types.Transaction, error) {
	return _SwanPayment.contract.Transact(opts, "initialize", owner, ERC20_TOKEN, oracle, priceFeed, chainlinkOracle)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address owner, address ERC20_TOKEN, address oracle, address priceFeed, address chainlinkOracle) returns()
func (_SwanPayment *SwanPaymentSession) Initialize(owner common.Address, ERC20_TOKEN common.Address, oracle common.Address, priceFeed common.Address, chainlinkOracle common.Address) (*types.Transaction, error) {
	return _SwanPayment.Contract.Initialize(&_SwanPayment.TransactOpts, owner, ERC20_TOKEN, oracle, priceFeed, chainlinkOracle)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address owner, address ERC20_TOKEN, address oracle, address priceFeed, address chainlinkOracle) returns()
func (_SwanPayment *SwanPaymentTransactorSession) Initialize(owner common.Address, ERC20_TOKEN common.Address, oracle common.Address, priceFeed common.Address, chainlinkOracle common.Address) (*types.Transaction, error) {
	return _SwanPayment.Contract.Initialize(&_SwanPayment.TransactOpts, owner, ERC20_TOKEN, oracle, priceFeed, chainlinkOracle)
}

// LockTokenPayment is a paid mutator transaction binding the contract method 0xc873ff65.
//
// Solidity: function lockTokenPayment((string,uint256,uint256,uint256,address,uint256) param) returns(bool)
func (_SwanPayment *SwanPaymentTransactor) LockTokenPayment(opts *bind.TransactOpts, param IPaymentMinimallockPaymentParam) (*types.Transaction, error) {
	return _SwanPayment.contract.Transact(opts, "lockTokenPayment", param)
}

// LockTokenPayment is a paid mutator transaction binding the contract method 0xc873ff65.
//
// Solidity: function lockTokenPayment((string,uint256,uint256,uint256,address,uint256) param) returns(bool)
func (_SwanPayment *SwanPaymentSession) LockTokenPayment(param IPaymentMinimallockPaymentParam) (*types.Transaction, error) {
	return _SwanPayment.Contract.LockTokenPayment(&_SwanPayment.TransactOpts, param)
}

// LockTokenPayment is a paid mutator transaction binding the contract method 0xc873ff65.
//
// Solidity: function lockTokenPayment((string,uint256,uint256,uint256,address,uint256) param) returns(bool)
func (_SwanPayment *SwanPaymentTransactorSession) LockTokenPayment(param IPaymentMinimallockPaymentParam) (*types.Transaction, error) {
	return _SwanPayment.Contract.LockTokenPayment(&_SwanPayment.TransactOpts, param)
}

// Refund is a paid mutator transaction binding the contract method 0x7d29985b.
//
// Solidity: function refund(string[] cidList) returns()
func (_SwanPayment *SwanPaymentTransactor) Refund(opts *bind.TransactOpts, cidList []string) (*types.Transaction, error) {
	return _SwanPayment.contract.Transact(opts, "refund", cidList)
}

// Refund is a paid mutator transaction binding the contract method 0x7d29985b.
//
// Solidity: function refund(string[] cidList) returns()
func (_SwanPayment *SwanPaymentSession) Refund(cidList []string) (*types.Transaction, error) {
	return _SwanPayment.Contract.Refund(&_SwanPayment.TransactOpts, cidList)
}

// Refund is a paid mutator transaction binding the contract method 0x7d29985b.
//
// Solidity: function refund(string[] cidList) returns()
func (_SwanPayment *SwanPaymentTransactorSession) Refund(cidList []string) (*types.Transaction, error) {
	return _SwanPayment.Contract.Refund(&_SwanPayment.TransactOpts, cidList)
}

// SetChainlinkOracle is a paid mutator transaction binding the contract method 0x7a9b0412.
//
// Solidity: function setChainlinkOracle(address _chainlinkOracle) returns(bool)
func (_SwanPayment *SwanPaymentTransactor) SetChainlinkOracle(opts *bind.TransactOpts, _chainlinkOracle common.Address) (*types.Transaction, error) {
	return _SwanPayment.contract.Transact(opts, "setChainlinkOracle", _chainlinkOracle)
}

// SetChainlinkOracle is a paid mutator transaction binding the contract method 0x7a9b0412.
//
// Solidity: function setChainlinkOracle(address _chainlinkOracle) returns(bool)
func (_SwanPayment *SwanPaymentSession) SetChainlinkOracle(_chainlinkOracle common.Address) (*types.Transaction, error) {
	return _SwanPayment.Contract.SetChainlinkOracle(&_SwanPayment.TransactOpts, _chainlinkOracle)
}

// SetChainlinkOracle is a paid mutator transaction binding the contract method 0x7a9b0412.
//
// Solidity: function setChainlinkOracle(address _chainlinkOracle) returns(bool)
func (_SwanPayment *SwanPaymentTransactorSession) SetChainlinkOracle(_chainlinkOracle common.Address) (*types.Transaction, error) {
	return _SwanPayment.Contract.SetChainlinkOracle(&_SwanPayment.TransactOpts, _chainlinkOracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns(bool)
func (_SwanPayment *SwanPaymentTransactor) SetOracle(opts *bind.TransactOpts, oracle common.Address) (*types.Transaction, error) {
	return _SwanPayment.contract.Transact(opts, "setOracle", oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns(bool)
func (_SwanPayment *SwanPaymentSession) SetOracle(oracle common.Address) (*types.Transaction, error) {
	return _SwanPayment.Contract.SetOracle(&_SwanPayment.TransactOpts, oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle) returns(bool)
func (_SwanPayment *SwanPaymentTransactorSession) SetOracle(oracle common.Address) (*types.Transaction, error) {
	return _SwanPayment.Contract.SetOracle(&_SwanPayment.TransactOpts, oracle)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address priceFeed) returns(bool)
func (_SwanPayment *SwanPaymentTransactor) SetPriceFeed(opts *bind.TransactOpts, priceFeed common.Address) (*types.Transaction, error) {
	return _SwanPayment.contract.Transact(opts, "setPriceFeed", priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address priceFeed) returns(bool)
func (_SwanPayment *SwanPaymentSession) SetPriceFeed(priceFeed common.Address) (*types.Transaction, error) {
	return _SwanPayment.Contract.SetPriceFeed(&_SwanPayment.TransactOpts, priceFeed)
}

// SetPriceFeed is a paid mutator transaction binding the contract method 0x724e78da.
//
// Solidity: function setPriceFeed(address priceFeed) returns(bool)
func (_SwanPayment *SwanPaymentTransactorSession) SetPriceFeed(priceFeed common.Address) (*types.Transaction, error) {
	return _SwanPayment.Contract.SetPriceFeed(&_SwanPayment.TransactOpts, priceFeed)
}

// UnlockCarPayment is a paid mutator transaction binding the contract method 0x38b3c23b.
//
// Solidity: function unlockCarPayment(string dealId, address recipient) returns(bool)
func (_SwanPayment *SwanPaymentTransactor) UnlockCarPayment(opts *bind.TransactOpts, dealId string, recipient common.Address) (*types.Transaction, error) {
	return _SwanPayment.contract.Transact(opts, "unlockCarPayment", dealId, recipient)
}

// UnlockCarPayment is a paid mutator transaction binding the contract method 0x38b3c23b.
//
// Solidity: function unlockCarPayment(string dealId, address recipient) returns(bool)
func (_SwanPayment *SwanPaymentSession) UnlockCarPayment(dealId string, recipient common.Address) (*types.Transaction, error) {
	return _SwanPayment.Contract.UnlockCarPayment(&_SwanPayment.TransactOpts, dealId, recipient)
}

// UnlockCarPayment is a paid mutator transaction binding the contract method 0x38b3c23b.
//
// Solidity: function unlockCarPayment(string dealId, address recipient) returns(bool)
func (_SwanPayment *SwanPaymentTransactorSession) UnlockCarPayment(dealId string, recipient common.Address) (*types.Transaction, error) {
	return _SwanPayment.Contract.UnlockCarPayment(&_SwanPayment.TransactOpts, dealId, recipient)
}

// UnlockTokenPayment is a paid mutator transaction binding the contract method 0x5c95e7e1.
//
// Solidity: function unlockTokenPayment((string,string,string,uint256,address) param) returns(bool)
func (_SwanPayment *SwanPaymentTransactor) UnlockTokenPayment(opts *bind.TransactOpts, param IPaymentMinimalunlockPaymentParam) (*types.Transaction, error) {
	return _SwanPayment.contract.Transact(opts, "unlockTokenPayment", param)
}

// UnlockTokenPayment is a paid mutator transaction binding the contract method 0x5c95e7e1.
//
// Solidity: function unlockTokenPayment((string,string,string,uint256,address) param) returns(bool)
func (_SwanPayment *SwanPaymentSession) UnlockTokenPayment(param IPaymentMinimalunlockPaymentParam) (*types.Transaction, error) {
	return _SwanPayment.Contract.UnlockTokenPayment(&_SwanPayment.TransactOpts, param)
}

// UnlockTokenPayment is a paid mutator transaction binding the contract method 0x5c95e7e1.
//
// Solidity: function unlockTokenPayment((string,string,string,uint256,address) param) returns(bool)
func (_SwanPayment *SwanPaymentTransactorSession) UnlockTokenPayment(param IPaymentMinimalunlockPaymentParam) (*types.Transaction, error) {
	return _SwanPayment.Contract.UnlockTokenPayment(&_SwanPayment.TransactOpts, param)
}

// SwanPaymentExpirePaymentIterator is returned from FilterExpirePayment and is used to iterate over the raw logs and unpacked data for ExpirePayment events raised by the SwanPayment contract.
type SwanPaymentExpirePaymentIterator struct {
	Event *SwanPaymentExpirePayment // Event containing the contract specifics and raw log

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
func (it *SwanPaymentExpirePaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanPaymentExpirePayment)
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
		it.Event = new(SwanPaymentExpirePayment)
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
func (it *SwanPaymentExpirePaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanPaymentExpirePaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanPaymentExpirePayment represents a ExpirePayment event raised by the SwanPayment contract.
type SwanPaymentExpirePayment struct {
	Id     string
	Token  common.Address
	Amount *big.Int
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterExpirePayment is a free log retrieval operation binding the contract event 0xe704d5e6168e602e91f017f25d889b182d9e11a90fd939a489cc2f04734c1f8a.
//
// Solidity: event ExpirePayment(string id, address token, uint256 amount, address owner)
func (_SwanPayment *SwanPaymentFilterer) FilterExpirePayment(opts *bind.FilterOpts) (*SwanPaymentExpirePaymentIterator, error) {

	logs, sub, err := _SwanPayment.contract.FilterLogs(opts, "ExpirePayment")
	if err != nil {
		return nil, err
	}
	return &SwanPaymentExpirePaymentIterator{contract: _SwanPayment.contract, event: "ExpirePayment", logs: logs, sub: sub}, nil
}

// WatchExpirePayment is a free log subscription operation binding the contract event 0xe704d5e6168e602e91f017f25d889b182d9e11a90fd939a489cc2f04734c1f8a.
//
// Solidity: event ExpirePayment(string id, address token, uint256 amount, address owner)
func (_SwanPayment *SwanPaymentFilterer) WatchExpirePayment(opts *bind.WatchOpts, sink chan<- *SwanPaymentExpirePayment) (event.Subscription, error) {

	logs, sub, err := _SwanPayment.contract.WatchLogs(opts, "ExpirePayment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanPaymentExpirePayment)
				if err := _SwanPayment.contract.UnpackLog(event, "ExpirePayment", log); err != nil {
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

// ParseExpirePayment is a log parse operation binding the contract event 0xe704d5e6168e602e91f017f25d889b182d9e11a90fd939a489cc2f04734c1f8a.
//
// Solidity: event ExpirePayment(string id, address token, uint256 amount, address owner)
func (_SwanPayment *SwanPaymentFilterer) ParseExpirePayment(log types.Log) (*SwanPaymentExpirePayment, error) {
	event := new(SwanPaymentExpirePayment)
	if err := _SwanPayment.contract.UnpackLog(event, "ExpirePayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwanPaymentLockPaymentIterator is returned from FilterLockPayment and is used to iterate over the raw logs and unpacked data for LockPayment events raised by the SwanPayment contract.
type SwanPaymentLockPaymentIterator struct {
	Event *SwanPaymentLockPayment // Event containing the contract specifics and raw log

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
func (it *SwanPaymentLockPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanPaymentLockPayment)
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
		it.Event = new(SwanPaymentLockPayment)
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
func (it *SwanPaymentLockPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanPaymentLockPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanPaymentLockPayment represents a LockPayment event raised by the SwanPayment contract.
type SwanPaymentLockPayment struct {
	Id         string
	Token      common.Address
	LockedFee  *big.Int
	MinPayment *big.Int
	Recipient  common.Address
	Deadline   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLockPayment is a free log retrieval operation binding the contract event 0x87b076911cb43e46012dbf762607ad5d0c5d6eb5e1d8fec72be99002451c18ba.
//
// Solidity: event LockPayment(string id, address token, uint256 lockedFee, uint256 minPayment, address recipient, uint256 deadline)
func (_SwanPayment *SwanPaymentFilterer) FilterLockPayment(opts *bind.FilterOpts) (*SwanPaymentLockPaymentIterator, error) {

	logs, sub, err := _SwanPayment.contract.FilterLogs(opts, "LockPayment")
	if err != nil {
		return nil, err
	}
	return &SwanPaymentLockPaymentIterator{contract: _SwanPayment.contract, event: "LockPayment", logs: logs, sub: sub}, nil
}

// WatchLockPayment is a free log subscription operation binding the contract event 0x87b076911cb43e46012dbf762607ad5d0c5d6eb5e1d8fec72be99002451c18ba.
//
// Solidity: event LockPayment(string id, address token, uint256 lockedFee, uint256 minPayment, address recipient, uint256 deadline)
func (_SwanPayment *SwanPaymentFilterer) WatchLockPayment(opts *bind.WatchOpts, sink chan<- *SwanPaymentLockPayment) (event.Subscription, error) {

	logs, sub, err := _SwanPayment.contract.WatchLogs(opts, "LockPayment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanPaymentLockPayment)
				if err := _SwanPayment.contract.UnpackLog(event, "LockPayment", log); err != nil {
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

// ParseLockPayment is a log parse operation binding the contract event 0x87b076911cb43e46012dbf762607ad5d0c5d6eb5e1d8fec72be99002451c18ba.
//
// Solidity: event LockPayment(string id, address token, uint256 lockedFee, uint256 minPayment, address recipient, uint256 deadline)
func (_SwanPayment *SwanPaymentFilterer) ParseLockPayment(log types.Log) (*SwanPaymentLockPayment, error) {
	event := new(SwanPaymentLockPayment)
	if err := _SwanPayment.contract.UnpackLog(event, "LockPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SwanPaymentUnlockPaymentIterator is returned from FilterUnlockPayment and is used to iterate over the raw logs and unpacked data for UnlockPayment events raised by the SwanPayment contract.
type SwanPaymentUnlockPaymentIterator struct {
	Event *SwanPaymentUnlockPayment // Event containing the contract specifics and raw log

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
func (it *SwanPaymentUnlockPaymentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SwanPaymentUnlockPayment)
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
		it.Event = new(SwanPaymentUnlockPayment)
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
func (it *SwanPaymentUnlockPaymentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SwanPaymentUnlockPaymentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SwanPaymentUnlockPayment represents a UnlockPayment event raised by the SwanPayment contract.
type SwanPaymentUnlockPayment struct {
	Id        string
	Token     common.Address
	Cost      *big.Int
	RestToken *big.Int
	Recipient common.Address
	Owner     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnlockPayment is a free log retrieval operation binding the contract event 0xd9aab088cd193ab56c9a682b9f2e42fb39206b42c3a6ba78130dc97a2a8d03ee.
//
// Solidity: event UnlockPayment(string id, address token, uint256 cost, uint256 restToken, address recipient, address owner)
func (_SwanPayment *SwanPaymentFilterer) FilterUnlockPayment(opts *bind.FilterOpts) (*SwanPaymentUnlockPaymentIterator, error) {

	logs, sub, err := _SwanPayment.contract.FilterLogs(opts, "UnlockPayment")
	if err != nil {
		return nil, err
	}
	return &SwanPaymentUnlockPaymentIterator{contract: _SwanPayment.contract, event: "UnlockPayment", logs: logs, sub: sub}, nil
}

// WatchUnlockPayment is a free log subscription operation binding the contract event 0xd9aab088cd193ab56c9a682b9f2e42fb39206b42c3a6ba78130dc97a2a8d03ee.
//
// Solidity: event UnlockPayment(string id, address token, uint256 cost, uint256 restToken, address recipient, address owner)
func (_SwanPayment *SwanPaymentFilterer) WatchUnlockPayment(opts *bind.WatchOpts, sink chan<- *SwanPaymentUnlockPayment) (event.Subscription, error) {

	logs, sub, err := _SwanPayment.contract.WatchLogs(opts, "UnlockPayment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SwanPaymentUnlockPayment)
				if err := _SwanPayment.contract.UnpackLog(event, "UnlockPayment", log); err != nil {
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

// ParseUnlockPayment is a log parse operation binding the contract event 0xd9aab088cd193ab56c9a682b9f2e42fb39206b42c3a6ba78130dc97a2a8d03ee.
//
// Solidity: event UnlockPayment(string id, address token, uint256 cost, uint256 restToken, address recipient, address owner)
func (_SwanPayment *SwanPaymentFilterer) ParseUnlockPayment(log types.Log) (*SwanPaymentUnlockPayment, error) {
	event := new(SwanPaymentUnlockPayment)
	if err := _SwanPayment.contract.UnpackLog(event, "UnlockPayment", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}