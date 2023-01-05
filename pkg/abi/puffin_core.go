// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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
)

// PuffinCoreMetaData contains all meta data concerning the PuffinCore contract.
var PuffinCoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isKYC\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setTier\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"tier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userTier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PuffinCoreABI is the input ABI used to generate the binding from.
// Deprecated: Use PuffinCoreMetaData.ABI instead.
var PuffinCoreABI = PuffinCoreMetaData.ABI

// PuffinCore is an auto generated Go binding around an Ethereum contract.
type PuffinCore struct {
	PuffinCoreCaller     // Read-only binding to the contract
	PuffinCoreTransactor // Write-only binding to the contract
	PuffinCoreFilterer   // Log filterer for contract events
}

// PuffinCoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type PuffinCoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinCoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PuffinCoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinCoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PuffinCoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinCoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PuffinCoreSession struct {
	Contract     *PuffinCore       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PuffinCoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PuffinCoreCallerSession struct {
	Contract *PuffinCoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PuffinCoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PuffinCoreTransactorSession struct {
	Contract     *PuffinCoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PuffinCoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type PuffinCoreRaw struct {
	Contract *PuffinCore // Generic contract binding to access the raw methods on
}

// PuffinCoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PuffinCoreCallerRaw struct {
	Contract *PuffinCoreCaller // Generic read-only contract binding to access the raw methods on
}

// PuffinCoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PuffinCoreTransactorRaw struct {
	Contract *PuffinCoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPuffinCore creates a new instance of PuffinCore, bound to a specific deployed contract.
func NewPuffinCore(address common.Address, backend bind.ContractBackend) (*PuffinCore, error) {
	contract, err := bindPuffinCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PuffinCore{PuffinCoreCaller: PuffinCoreCaller{contract: contract}, PuffinCoreTransactor: PuffinCoreTransactor{contract: contract}, PuffinCoreFilterer: PuffinCoreFilterer{contract: contract}}, nil
}

// NewPuffinCoreCaller creates a new read-only instance of PuffinCore, bound to a specific deployed contract.
func NewPuffinCoreCaller(address common.Address, caller bind.ContractCaller) (*PuffinCoreCaller, error) {
	contract, err := bindPuffinCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PuffinCoreCaller{contract: contract}, nil
}

// NewPuffinCoreTransactor creates a new write-only instance of PuffinCore, bound to a specific deployed contract.
func NewPuffinCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*PuffinCoreTransactor, error) {
	contract, err := bindPuffinCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PuffinCoreTransactor{contract: contract}, nil
}

// NewPuffinCoreFilterer creates a new log filterer instance of PuffinCore, bound to a specific deployed contract.
func NewPuffinCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*PuffinCoreFilterer, error) {
	contract, err := bindPuffinCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PuffinCoreFilterer{contract: contract}, nil
}

// bindPuffinCore binds a generic wrapper to an already deployed contract.
func bindPuffinCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PuffinCoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PuffinCore *PuffinCoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PuffinCore.Contract.PuffinCoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PuffinCore *PuffinCoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinCore.Contract.PuffinCoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PuffinCore *PuffinCoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PuffinCore.Contract.PuffinCoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PuffinCore *PuffinCoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PuffinCore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PuffinCore *PuffinCoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinCore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PuffinCore *PuffinCoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PuffinCore.Contract.contract.Transact(opts, method, params...)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_PuffinCore *PuffinCoreCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PuffinCore.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_PuffinCore *PuffinCoreSession) Count() (*big.Int, error) {
	return _PuffinCore.Contract.Count(&_PuffinCore.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_PuffinCore *PuffinCoreCallerSession) Count() (*big.Int, error) {
	return _PuffinCore.Contract.Count(&_PuffinCore.CallOpts)
}

// IsKYC is a free data retrieval call binding the contract method 0x9944f518.
//
// Solidity: function isKYC(address ) view returns(bool)
func (_PuffinCore *PuffinCoreCaller) IsKYC(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PuffinCore.contract.Call(opts, &out, "isKYC", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsKYC is a free data retrieval call binding the contract method 0x9944f518.
//
// Solidity: function isKYC(address ) view returns(bool)
func (_PuffinCore *PuffinCoreSession) IsKYC(arg0 common.Address) (bool, error) {
	return _PuffinCore.Contract.IsKYC(&_PuffinCore.CallOpts, arg0)
}

// IsKYC is a free data retrieval call binding the contract method 0x9944f518.
//
// Solidity: function isKYC(address ) view returns(bool)
func (_PuffinCore *PuffinCoreCallerSession) IsKYC(arg0 common.Address) (bool, error) {
	return _PuffinCore.Contract.IsKYC(&_PuffinCore.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PuffinCore *PuffinCoreCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PuffinCore.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PuffinCore *PuffinCoreSession) Name() (string, error) {
	return _PuffinCore.Contract.Name(&_PuffinCore.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PuffinCore *PuffinCoreCallerSession) Name() (string, error) {
	return _PuffinCore.Contract.Name(&_PuffinCore.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinCore *PuffinCoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinCore.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinCore *PuffinCoreSession) Owner() (common.Address, error) {
	return _PuffinCore.Contract.Owner(&_PuffinCore.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinCore *PuffinCoreCallerSession) Owner() (common.Address, error) {
	return _PuffinCore.Contract.Owner(&_PuffinCore.CallOpts)
}

// Tier is a free data retrieval call binding the contract method 0x2785f8bb.
//
// Solidity: function tier(address user) view returns(uint256)
func (_PuffinCore *PuffinCoreCaller) Tier(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PuffinCore.contract.Call(opts, &out, "tier", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Tier is a free data retrieval call binding the contract method 0x2785f8bb.
//
// Solidity: function tier(address user) view returns(uint256)
func (_PuffinCore *PuffinCoreSession) Tier(user common.Address) (*big.Int, error) {
	return _PuffinCore.Contract.Tier(&_PuffinCore.CallOpts, user)
}

// Tier is a free data retrieval call binding the contract method 0x2785f8bb.
//
// Solidity: function tier(address user) view returns(uint256)
func (_PuffinCore *PuffinCoreCallerSession) Tier(user common.Address) (*big.Int, error) {
	return _PuffinCore.Contract.Tier(&_PuffinCore.CallOpts, user)
}

// UserTier is a free data retrieval call binding the contract method 0x21c7557c.
//
// Solidity: function userTier(address ) view returns(uint256)
func (_PuffinCore *PuffinCoreCaller) UserTier(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PuffinCore.contract.Call(opts, &out, "userTier", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserTier is a free data retrieval call binding the contract method 0x21c7557c.
//
// Solidity: function userTier(address ) view returns(uint256)
func (_PuffinCore *PuffinCoreSession) UserTier(arg0 common.Address) (*big.Int, error) {
	return _PuffinCore.Contract.UserTier(&_PuffinCore.CallOpts, arg0)
}

// UserTier is a free data retrieval call binding the contract method 0x21c7557c.
//
// Solidity: function userTier(address ) view returns(uint256)
func (_PuffinCore *PuffinCoreCallerSession) UserTier(arg0 common.Address) (*big.Int, error) {
	return _PuffinCore.Contract.UserTier(&_PuffinCore.CallOpts, arg0)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinCore *PuffinCoreTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinCore.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinCore *PuffinCoreSession) RenounceOwnership() (*types.Transaction, error) {
	return _PuffinCore.Contract.RenounceOwnership(&_PuffinCore.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinCore *PuffinCoreTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PuffinCore.Contract.RenounceOwnership(&_PuffinCore.TransactOpts)
}

// SetTier is a paid mutator transaction binding the contract method 0xa914adce.
//
// Solidity: function setTier(address user, uint256 val) returns()
func (_PuffinCore *PuffinCoreTransactor) SetTier(opts *bind.TransactOpts, user common.Address, val *big.Int) (*types.Transaction, error) {
	return _PuffinCore.contract.Transact(opts, "setTier", user, val)
}

// SetTier is a paid mutator transaction binding the contract method 0xa914adce.
//
// Solidity: function setTier(address user, uint256 val) returns()
func (_PuffinCore *PuffinCoreSession) SetTier(user common.Address, val *big.Int) (*types.Transaction, error) {
	return _PuffinCore.Contract.SetTier(&_PuffinCore.TransactOpts, user, val)
}

// SetTier is a paid mutator transaction binding the contract method 0xa914adce.
//
// Solidity: function setTier(address user, uint256 val) returns()
func (_PuffinCore *PuffinCoreTransactorSession) SetTier(user common.Address, val *big.Int) (*types.Transaction, error) {
	return _PuffinCore.Contract.SetTier(&_PuffinCore.TransactOpts, user, val)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinCore *PuffinCoreTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PuffinCore.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinCore *PuffinCoreSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PuffinCore.Contract.TransferOwnership(&_PuffinCore.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinCore *PuffinCoreTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PuffinCore.Contract.TransferOwnership(&_PuffinCore.TransactOpts, newOwner)
}

// PuffinCoreOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PuffinCore contract.
type PuffinCoreOwnershipTransferredIterator struct {
	Event *PuffinCoreOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PuffinCoreOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinCoreOwnershipTransferred)
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
		it.Event = new(PuffinCoreOwnershipTransferred)
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
func (it *PuffinCoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinCoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinCoreOwnershipTransferred represents a OwnershipTransferred event raised by the PuffinCore contract.
type PuffinCoreOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PuffinCore *PuffinCoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PuffinCoreOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PuffinCore.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PuffinCoreOwnershipTransferredIterator{contract: _PuffinCore.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PuffinCore *PuffinCoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PuffinCoreOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PuffinCore.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinCoreOwnershipTransferred)
				if err := _PuffinCore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PuffinCore *PuffinCoreFilterer) ParseOwnershipTransferred(log types.Log) (*PuffinCoreOwnershipTransferred, error) {
	event := new(PuffinCoreOwnershipTransferred)
	if err := _PuffinCore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
