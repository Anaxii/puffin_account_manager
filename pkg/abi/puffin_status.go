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

// PuffinStatusMetaData contains all meta data concerning the PuffinStatus contract.
var PuffinStatusMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"NewUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"UserRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"canView\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"removeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setCanView\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_val\",\"type\":\"uint256\"}],\"name\":\"setStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PuffinStatusABI is the input ABI used to generate the binding from.
// Deprecated: Use PuffinStatusMetaData.ABI instead.
var PuffinStatusABI = PuffinStatusMetaData.ABI

// PuffinStatus is an auto generated Go binding around an Ethereum contract.
type PuffinStatus struct {
	PuffinStatusCaller     // Read-only binding to the contract
	PuffinStatusTransactor // Write-only binding to the contract
	PuffinStatusFilterer   // Log filterer for contract events
}

// PuffinStatusCaller is an auto generated read-only Go binding around an Ethereum contract.
type PuffinStatusCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinStatusTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PuffinStatusTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinStatusFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PuffinStatusFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinStatusSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PuffinStatusSession struct {
	Contract     *PuffinStatus     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PuffinStatusCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PuffinStatusCallerSession struct {
	Contract *PuffinStatusCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PuffinStatusTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PuffinStatusTransactorSession struct {
	Contract     *PuffinStatusTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PuffinStatusRaw is an auto generated low-level Go binding around an Ethereum contract.
type PuffinStatusRaw struct {
	Contract *PuffinStatus // Generic contract binding to access the raw methods on
}

// PuffinStatusCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PuffinStatusCallerRaw struct {
	Contract *PuffinStatusCaller // Generic read-only contract binding to access the raw methods on
}

// PuffinStatusTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PuffinStatusTransactorRaw struct {
	Contract *PuffinStatusTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPuffinStatus creates a new instance of PuffinStatus, bound to a specific deployed contract.
func NewPuffinStatus(address common.Address, backend bind.ContractBackend) (*PuffinStatus, error) {
	contract, err := bindPuffinStatus(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PuffinStatus{PuffinStatusCaller: PuffinStatusCaller{contract: contract}, PuffinStatusTransactor: PuffinStatusTransactor{contract: contract}, PuffinStatusFilterer: PuffinStatusFilterer{contract: contract}}, nil
}

// NewPuffinStatusCaller creates a new read-only instance of PuffinStatus, bound to a specific deployed contract.
func NewPuffinStatusCaller(address common.Address, caller bind.ContractCaller) (*PuffinStatusCaller, error) {
	contract, err := bindPuffinStatus(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PuffinStatusCaller{contract: contract}, nil
}

// NewPuffinStatusTransactor creates a new write-only instance of PuffinStatus, bound to a specific deployed contract.
func NewPuffinStatusTransactor(address common.Address, transactor bind.ContractTransactor) (*PuffinStatusTransactor, error) {
	contract, err := bindPuffinStatus(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PuffinStatusTransactor{contract: contract}, nil
}

// NewPuffinStatusFilterer creates a new log filterer instance of PuffinStatus, bound to a specific deployed contract.
func NewPuffinStatusFilterer(address common.Address, filterer bind.ContractFilterer) (*PuffinStatusFilterer, error) {
	contract, err := bindPuffinStatus(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PuffinStatusFilterer{contract: contract}, nil
}

// bindPuffinStatus binds a generic wrapper to an already deployed contract.
func bindPuffinStatus(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PuffinStatusABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PuffinStatus *PuffinStatusRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PuffinStatus.Contract.PuffinStatusCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PuffinStatus *PuffinStatusRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinStatus.Contract.PuffinStatusTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PuffinStatus *PuffinStatusRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PuffinStatus.Contract.PuffinStatusTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PuffinStatus *PuffinStatusCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PuffinStatus.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PuffinStatus *PuffinStatusTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinStatus.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PuffinStatus *PuffinStatusTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PuffinStatus.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PuffinStatus *PuffinStatusCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinStatus.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PuffinStatus *PuffinStatusSession) Admin() (common.Address, error) {
	return _PuffinStatus.Contract.Admin(&_PuffinStatus.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PuffinStatus *PuffinStatusCallerSession) Admin() (common.Address, error) {
	return _PuffinStatus.Contract.Admin(&_PuffinStatus.CallOpts)
}

// CanView is a free data retrieval call binding the contract method 0x50ef672d.
//
// Solidity: function canView(address ) view returns(bool)
func (_PuffinStatus *PuffinStatusCaller) CanView(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PuffinStatus.contract.Call(opts, &out, "canView", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanView is a free data retrieval call binding the contract method 0x50ef672d.
//
// Solidity: function canView(address ) view returns(bool)
func (_PuffinStatus *PuffinStatusSession) CanView(arg0 common.Address) (bool, error) {
	return _PuffinStatus.Contract.CanView(&_PuffinStatus.CallOpts, arg0)
}

// CanView is a free data retrieval call binding the contract method 0x50ef672d.
//
// Solidity: function canView(address ) view returns(bool)
func (_PuffinStatus *PuffinStatusCallerSession) CanView(arg0 common.Address) (bool, error) {
	return _PuffinStatus.Contract.CanView(&_PuffinStatus.CallOpts, arg0)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_PuffinStatus *PuffinStatusCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PuffinStatus.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_PuffinStatus *PuffinStatusSession) Count() (*big.Int, error) {
	return _PuffinStatus.Contract.Count(&_PuffinStatus.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_PuffinStatus *PuffinStatusCallerSession) Count() (*big.Int, error) {
	return _PuffinStatus.Contract.Count(&_PuffinStatus.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PuffinStatus *PuffinStatusCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinStatus.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PuffinStatus *PuffinStatusSession) Factory() (common.Address, error) {
	return _PuffinStatus.Contract.Factory(&_PuffinStatus.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_PuffinStatus *PuffinStatusCallerSession) Factory() (common.Address, error) {
	return _PuffinStatus.Contract.Factory(&_PuffinStatus.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PuffinStatus *PuffinStatusCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PuffinStatus.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PuffinStatus *PuffinStatusSession) Name() (string, error) {
	return _PuffinStatus.Contract.Name(&_PuffinStatus.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_PuffinStatus *PuffinStatusCallerSession) Name() (string, error) {
	return _PuffinStatus.Contract.Name(&_PuffinStatus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinStatus *PuffinStatusCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinStatus.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinStatus *PuffinStatusSession) Owner() (common.Address, error) {
	return _PuffinStatus.Contract.Owner(&_PuffinStatus.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinStatus *PuffinStatusCallerSession) Owner() (common.Address, error) {
	return _PuffinStatus.Contract.Owner(&_PuffinStatus.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x645b8b1b.
//
// Solidity: function status(address _user) view returns(uint256, bool)
func (_PuffinStatus *PuffinStatusCaller) Status(opts *bind.CallOpts, _user common.Address) (*big.Int, bool, error) {
	var out []interface{}
	err := _PuffinStatus.contract.Call(opts, &out, "status", _user)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// Status is a free data retrieval call binding the contract method 0x645b8b1b.
//
// Solidity: function status(address _user) view returns(uint256, bool)
func (_PuffinStatus *PuffinStatusSession) Status(_user common.Address) (*big.Int, bool, error) {
	return _PuffinStatus.Contract.Status(&_PuffinStatus.CallOpts, _user)
}

// Status is a free data retrieval call binding the contract method 0x645b8b1b.
//
// Solidity: function status(address _user) view returns(uint256, bool)
func (_PuffinStatus *PuffinStatusCallerSession) Status(_user common.Address) (*big.Int, bool, error) {
	return _PuffinStatus.Contract.Status(&_PuffinStatus.CallOpts, _user)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(address _user) returns()
func (_PuffinStatus *PuffinStatusTransactor) RemoveUser(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _PuffinStatus.contract.Transact(opts, "removeUser", _user)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(address _user) returns()
func (_PuffinStatus *PuffinStatusSession) RemoveUser(_user common.Address) (*types.Transaction, error) {
	return _PuffinStatus.Contract.RemoveUser(&_PuffinStatus.TransactOpts, _user)
}

// RemoveUser is a paid mutator transaction binding the contract method 0x98575188.
//
// Solidity: function removeUser(address _user) returns()
func (_PuffinStatus *PuffinStatusTransactorSession) RemoveUser(_user common.Address) (*types.Transaction, error) {
	return _PuffinStatus.Contract.RemoveUser(&_PuffinStatus.TransactOpts, _user)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinStatus *PuffinStatusTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinStatus.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinStatus *PuffinStatusSession) RenounceOwnership() (*types.Transaction, error) {
	return _PuffinStatus.Contract.RenounceOwnership(&_PuffinStatus.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinStatus *PuffinStatusTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PuffinStatus.Contract.RenounceOwnership(&_PuffinStatus.TransactOpts)
}

// SetCanView is a paid mutator transaction binding the contract method 0xd7246dc5.
//
// Solidity: function setCanView(address _user, bool _status) returns()
func (_PuffinStatus *PuffinStatusTransactor) SetCanView(opts *bind.TransactOpts, _user common.Address, _status bool) (*types.Transaction, error) {
	return _PuffinStatus.contract.Transact(opts, "setCanView", _user, _status)
}

// SetCanView is a paid mutator transaction binding the contract method 0xd7246dc5.
//
// Solidity: function setCanView(address _user, bool _status) returns()
func (_PuffinStatus *PuffinStatusSession) SetCanView(_user common.Address, _status bool) (*types.Transaction, error) {
	return _PuffinStatus.Contract.SetCanView(&_PuffinStatus.TransactOpts, _user, _status)
}

// SetCanView is a paid mutator transaction binding the contract method 0xd7246dc5.
//
// Solidity: function setCanView(address _user, bool _status) returns()
func (_PuffinStatus *PuffinStatusTransactorSession) SetCanView(_user common.Address, _status bool) (*types.Transaction, error) {
	return _PuffinStatus.Contract.SetCanView(&_PuffinStatus.TransactOpts, _user, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x7e55d59c.
//
// Solidity: function setStatus(address _user, uint256 _val) returns()
func (_PuffinStatus *PuffinStatusTransactor) SetStatus(opts *bind.TransactOpts, _user common.Address, _val *big.Int) (*types.Transaction, error) {
	return _PuffinStatus.contract.Transact(opts, "setStatus", _user, _val)
}

// SetStatus is a paid mutator transaction binding the contract method 0x7e55d59c.
//
// Solidity: function setStatus(address _user, uint256 _val) returns()
func (_PuffinStatus *PuffinStatusSession) SetStatus(_user common.Address, _val *big.Int) (*types.Transaction, error) {
	return _PuffinStatus.Contract.SetStatus(&_PuffinStatus.TransactOpts, _user, _val)
}

// SetStatus is a paid mutator transaction binding the contract method 0x7e55d59c.
//
// Solidity: function setStatus(address _user, uint256 _val) returns()
func (_PuffinStatus *PuffinStatusTransactorSession) SetStatus(_user common.Address, _val *big.Int) (*types.Transaction, error) {
	return _PuffinStatus.Contract.SetStatus(&_PuffinStatus.TransactOpts, _user, _val)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinStatus *PuffinStatusTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PuffinStatus.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinStatus *PuffinStatusSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PuffinStatus.Contract.TransferOwnership(&_PuffinStatus.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinStatus *PuffinStatusTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PuffinStatus.Contract.TransferOwnership(&_PuffinStatus.TransactOpts, newOwner)
}

// PuffinStatusNewUserIterator is returned from FilterNewUser and is used to iterate over the raw logs and unpacked data for NewUser events raised by the PuffinStatus contract.
type PuffinStatusNewUserIterator struct {
	Event *PuffinStatusNewUser // Event containing the contract specifics and raw log

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
func (it *PuffinStatusNewUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinStatusNewUser)
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
		it.Event = new(PuffinStatusNewUser)
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
func (it *PuffinStatusNewUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinStatusNewUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinStatusNewUser represents a NewUser event raised by the PuffinStatus contract.
type PuffinStatusNewUser struct {
	User  common.Address
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewUser is a free log retrieval operation binding the contract event 0xd75b88a9203b17f9356ca063241beac16a25d5a46f485b378c2c229e864bdd4d.
//
// Solidity: event NewUser(address indexed user, uint256 indexed count)
func (_PuffinStatus *PuffinStatusFilterer) FilterNewUser(opts *bind.FilterOpts, user []common.Address, count []*big.Int) (*PuffinStatusNewUserIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var countRule []interface{}
	for _, countItem := range count {
		countRule = append(countRule, countItem)
	}

	logs, sub, err := _PuffinStatus.contract.FilterLogs(opts, "NewUser", userRule, countRule)
	if err != nil {
		return nil, err
	}
	return &PuffinStatusNewUserIterator{contract: _PuffinStatus.contract, event: "NewUser", logs: logs, sub: sub}, nil
}

// WatchNewUser is a free log subscription operation binding the contract event 0xd75b88a9203b17f9356ca063241beac16a25d5a46f485b378c2c229e864bdd4d.
//
// Solidity: event NewUser(address indexed user, uint256 indexed count)
func (_PuffinStatus *PuffinStatusFilterer) WatchNewUser(opts *bind.WatchOpts, sink chan<- *PuffinStatusNewUser, user []common.Address, count []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var countRule []interface{}
	for _, countItem := range count {
		countRule = append(countRule, countItem)
	}

	logs, sub, err := _PuffinStatus.contract.WatchLogs(opts, "NewUser", userRule, countRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinStatusNewUser)
				if err := _PuffinStatus.contract.UnpackLog(event, "NewUser", log); err != nil {
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

// ParseNewUser is a log parse operation binding the contract event 0xd75b88a9203b17f9356ca063241beac16a25d5a46f485b378c2c229e864bdd4d.
//
// Solidity: event NewUser(address indexed user, uint256 indexed count)
func (_PuffinStatus *PuffinStatusFilterer) ParseNewUser(log types.Log) (*PuffinStatusNewUser, error) {
	event := new(PuffinStatusNewUser)
	if err := _PuffinStatus.contract.UnpackLog(event, "NewUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinStatusOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PuffinStatus contract.
type PuffinStatusOwnershipTransferredIterator struct {
	Event *PuffinStatusOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PuffinStatusOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinStatusOwnershipTransferred)
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
		it.Event = new(PuffinStatusOwnershipTransferred)
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
func (it *PuffinStatusOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinStatusOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinStatusOwnershipTransferred represents a OwnershipTransferred event raised by the PuffinStatus contract.
type PuffinStatusOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PuffinStatus *PuffinStatusFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PuffinStatusOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PuffinStatus.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PuffinStatusOwnershipTransferredIterator{contract: _PuffinStatus.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PuffinStatus *PuffinStatusFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PuffinStatusOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PuffinStatus.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinStatusOwnershipTransferred)
				if err := _PuffinStatus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PuffinStatus *PuffinStatusFilterer) ParseOwnershipTransferred(log types.Log) (*PuffinStatusOwnershipTransferred, error) {
	event := new(PuffinStatusOwnershipTransferred)
	if err := _PuffinStatus.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinStatusUserRemovedIterator is returned from FilterUserRemoved and is used to iterate over the raw logs and unpacked data for UserRemoved events raised by the PuffinStatus contract.
type PuffinStatusUserRemovedIterator struct {
	Event *PuffinStatusUserRemoved // Event containing the contract specifics and raw log

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
func (it *PuffinStatusUserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinStatusUserRemoved)
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
		it.Event = new(PuffinStatusUserRemoved)
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
func (it *PuffinStatusUserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinStatusUserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinStatusUserRemoved represents a UserRemoved event raised by the PuffinStatus contract.
type PuffinStatusUserRemoved struct {
	User  common.Address
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUserRemoved is a free log retrieval operation binding the contract event 0xf62c903ca15668006710c4411b74e3acb9efabe0d10bece1a2912050499f7efd.
//
// Solidity: event UserRemoved(address indexed user, uint256 indexed count)
func (_PuffinStatus *PuffinStatusFilterer) FilterUserRemoved(opts *bind.FilterOpts, user []common.Address, count []*big.Int) (*PuffinStatusUserRemovedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var countRule []interface{}
	for _, countItem := range count {
		countRule = append(countRule, countItem)
	}

	logs, sub, err := _PuffinStatus.contract.FilterLogs(opts, "UserRemoved", userRule, countRule)
	if err != nil {
		return nil, err
	}
	return &PuffinStatusUserRemovedIterator{contract: _PuffinStatus.contract, event: "UserRemoved", logs: logs, sub: sub}, nil
}

// WatchUserRemoved is a free log subscription operation binding the contract event 0xf62c903ca15668006710c4411b74e3acb9efabe0d10bece1a2912050499f7efd.
//
// Solidity: event UserRemoved(address indexed user, uint256 indexed count)
func (_PuffinStatus *PuffinStatusFilterer) WatchUserRemoved(opts *bind.WatchOpts, sink chan<- *PuffinStatusUserRemoved, user []common.Address, count []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var countRule []interface{}
	for _, countItem := range count {
		countRule = append(countRule, countItem)
	}

	logs, sub, err := _PuffinStatus.contract.WatchLogs(opts, "UserRemoved", userRule, countRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinStatusUserRemoved)
				if err := _PuffinStatus.contract.UnpackLog(event, "UserRemoved", log); err != nil {
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

// ParseUserRemoved is a log parse operation binding the contract event 0xf62c903ca15668006710c4411b74e3acb9efabe0d10bece1a2912050499f7efd.
//
// Solidity: event UserRemoved(address indexed user, uint256 indexed count)
func (_PuffinStatus *PuffinStatusFilterer) ParseUserRemoved(log types.Log) (*PuffinStatusUserRemoved, error) {
	event := new(PuffinStatusUserRemoved)
	if err := _PuffinStatus.contract.UnpackLog(event, "UserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
