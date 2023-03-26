// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractsv1

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

// MainMetaData contains all meta data concerning the Main contract.
var MainMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"arr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bool1\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"set2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"str1\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"temp\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"val1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MainABI is the input ABI used to generate the binding from.
// Deprecated: Use MainMetaData.ABI instead.
var MainABI = MainMetaData.ABI

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MainMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// Arr is a free data retrieval call binding the contract method 0x71e5ee5f.
//
// Solidity: function arr(uint256 ) view returns(uint256)
func (_Main *MainCaller) Arr(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "arr", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Arr is a free data retrieval call binding the contract method 0x71e5ee5f.
//
// Solidity: function arr(uint256 ) view returns(uint256)
func (_Main *MainSession) Arr(arg0 *big.Int) (*big.Int, error) {
	return _Main.Contract.Arr(&_Main.CallOpts, arg0)
}

// Arr is a free data retrieval call binding the contract method 0x71e5ee5f.
//
// Solidity: function arr(uint256 ) view returns(uint256)
func (_Main *MainCallerSession) Arr(arg0 *big.Int) (*big.Int, error) {
	return _Main.Contract.Arr(&_Main.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Main *MainCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Main *MainSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Main.Contract.Balances(&_Main.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_Main *MainCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _Main.Contract.Balances(&_Main.CallOpts, arg0)
}

// Balances2 is a free data retrieval call binding the contract method 0xd54dea9e.
//
// Solidity: function balances2(address , address ) view returns(uint256)
func (_Main *MainCaller) Balances2(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "balances2", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances2 is a free data retrieval call binding the contract method 0xd54dea9e.
//
// Solidity: function balances2(address , address ) view returns(uint256)
func (_Main *MainSession) Balances2(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Main.Contract.Balances2(&_Main.CallOpts, arg0, arg1)
}

// Balances2 is a free data retrieval call binding the contract method 0xd54dea9e.
//
// Solidity: function balances2(address , address ) view returns(uint256)
func (_Main *MainCallerSession) Balances2(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Main.Contract.Balances2(&_Main.CallOpts, arg0, arg1)
}

// Bool1 is a free data retrieval call binding the contract method 0xec630c70.
//
// Solidity: function bool1() view returns(bool)
func (_Main *MainCaller) Bool1(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "bool1")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Bool1 is a free data retrieval call binding the contract method 0xec630c70.
//
// Solidity: function bool1() view returns(bool)
func (_Main *MainSession) Bool1() (bool, error) {
	return _Main.Contract.Bool1(&_Main.CallOpts)
}

// Bool1 is a free data retrieval call binding the contract method 0xec630c70.
//
// Solidity: function bool1() view returns(bool)
func (_Main *MainCallerSession) Bool1() (bool, error) {
	return _Main.Contract.Bool1(&_Main.CallOpts)
}

// Str1 is a free data retrieval call binding the contract method 0xd118d187.
//
// Solidity: function str1() view returns(string)
func (_Main *MainCaller) Str1(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "str1")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Str1 is a free data retrieval call binding the contract method 0xd118d187.
//
// Solidity: function str1() view returns(string)
func (_Main *MainSession) Str1() (string, error) {
	return _Main.Contract.Str1(&_Main.CallOpts)
}

// Str1 is a free data retrieval call binding the contract method 0xd118d187.
//
// Solidity: function str1() view returns(string)
func (_Main *MainCallerSession) Str1() (string, error) {
	return _Main.Contract.Str1(&_Main.CallOpts)
}

// Temp is a free data retrieval call binding the contract method 0x673402e5.
//
// Solidity: function temp() view returns(string name, uint256 val)
func (_Main *MainCaller) Temp(opts *bind.CallOpts) (struct {
	Name string
	Val  *big.Int
}, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "temp")

	outstruct := new(struct {
		Name string
		Val  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Val = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Temp is a free data retrieval call binding the contract method 0x673402e5.
//
// Solidity: function temp() view returns(string name, uint256 val)
func (_Main *MainSession) Temp() (struct {
	Name string
	Val  *big.Int
}, error) {
	return _Main.Contract.Temp(&_Main.CallOpts)
}

// Temp is a free data retrieval call binding the contract method 0x673402e5.
//
// Solidity: function temp() view returns(string name, uint256 val)
func (_Main *MainCallerSession) Temp() (struct {
	Name string
	Val  *big.Int
}, error) {
	return _Main.Contract.Temp(&_Main.CallOpts)
}

// Val1 is a free data retrieval call binding the contract method 0xc82fdf36.
//
// Solidity: function val1() view returns(uint256)
func (_Main *MainCaller) Val1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Main.contract.Call(opts, &out, "val1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Val1 is a free data retrieval call binding the contract method 0xc82fdf36.
//
// Solidity: function val1() view returns(uint256)
func (_Main *MainSession) Val1() (*big.Int, error) {
	return _Main.Contract.Val1(&_Main.CallOpts)
}

// Val1 is a free data retrieval call binding the contract method 0xc82fdf36.
//
// Solidity: function val1() view returns(uint256)
func (_Main *MainCallerSession) Val1() (*big.Int, error) {
	return _Main.Contract.Val1(&_Main.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0xb8e010de.
//
// Solidity: function set() returns()
func (_Main *MainTransactor) Set(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "set")
}

// Set is a paid mutator transaction binding the contract method 0xb8e010de.
//
// Solidity: function set() returns()
func (_Main *MainSession) Set() (*types.Transaction, error) {
	return _Main.Contract.Set(&_Main.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0xb8e010de.
//
// Solidity: function set() returns()
func (_Main *MainTransactorSession) Set() (*types.Transaction, error) {
	return _Main.Contract.Set(&_Main.TransactOpts)
}

// Set2 is a paid mutator transaction binding the contract method 0x6f315638.
//
// Solidity: function set2() returns()
func (_Main *MainTransactor) Set2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "set2")
}

// Set2 is a paid mutator transaction binding the contract method 0x6f315638.
//
// Solidity: function set2() returns()
func (_Main *MainSession) Set2() (*types.Transaction, error) {
	return _Main.Contract.Set2(&_Main.TransactOpts)
}

// Set2 is a paid mutator transaction binding the contract method 0x6f315638.
//
// Solidity: function set2() returns()
func (_Main *MainTransactorSession) Set2() (*types.Transaction, error) {
	return _Main.Contract.Set2(&_Main.TransactOpts)
}
