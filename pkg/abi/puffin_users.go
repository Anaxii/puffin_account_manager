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

// PuffinUsersMetaData contains all meta data concerning the PuffinUsers contract.
var PuffinUsersMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"NewUser\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"UserRemoved\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"canView\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"count\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"}],\"name\":\"setCanView\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_status\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_val\",\"type\":\"uint256\"}],\"name\":\"setStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"status\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PuffinUsersABI is the input ABI used to generate the binding from.
// Deprecated: Use PuffinUsersMetaData.ABI instead.
var PuffinUsersABI = PuffinUsersMetaData.ABI

// PuffinUsers is an auto generated Go binding around an Ethereum contract.
type PuffinUsers struct {
	PuffinUsersCaller     // Read-only binding to the contract
	PuffinUsersTransactor // Write-only binding to the contract
	PuffinUsersFilterer   // Log filterer for contract events
}

// PuffinUsersCaller is an auto generated read-only Go binding around an Ethereum contract.
type PuffinUsersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinUsersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PuffinUsersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinUsersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PuffinUsersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PuffinUsersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PuffinUsersSession struct {
	Contract     *PuffinUsers      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PuffinUsersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PuffinUsersCallerSession struct {
	Contract *PuffinUsersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PuffinUsersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PuffinUsersTransactorSession struct {
	Contract     *PuffinUsersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PuffinUsersRaw is an auto generated low-level Go binding around an Ethereum contract.
type PuffinUsersRaw struct {
	Contract *PuffinUsers // Generic contract binding to access the raw methods on
}

// PuffinUsersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PuffinUsersCallerRaw struct {
	Contract *PuffinUsersCaller // Generic read-only contract binding to access the raw methods on
}

// PuffinUsersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PuffinUsersTransactorRaw struct {
	Contract *PuffinUsersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPuffinUsers creates a new instance of PuffinUsers, bound to a specific deployed contract.
func NewPuffinUsers(address common.Address, backend bind.ContractBackend) (*PuffinUsers, error) {
	contract, err := bindPuffinUsers(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PuffinUsers{PuffinUsersCaller: PuffinUsersCaller{contract: contract}, PuffinUsersTransactor: PuffinUsersTransactor{contract: contract}, PuffinUsersFilterer: PuffinUsersFilterer{contract: contract}}, nil
}

// NewPuffinUsersCaller creates a new read-only instance of PuffinUsers, bound to a specific deployed contract.
func NewPuffinUsersCaller(address common.Address, caller bind.ContractCaller) (*PuffinUsersCaller, error) {
	contract, err := bindPuffinUsers(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PuffinUsersCaller{contract: contract}, nil
}

// NewPuffinUsersTransactor creates a new write-only instance of PuffinUsers, bound to a specific deployed contract.
func NewPuffinUsersTransactor(address common.Address, transactor bind.ContractTransactor) (*PuffinUsersTransactor, error) {
	contract, err := bindPuffinUsers(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PuffinUsersTransactor{contract: contract}, nil
}

// NewPuffinUsersFilterer creates a new log filterer instance of PuffinUsers, bound to a specific deployed contract.
func NewPuffinUsersFilterer(address common.Address, filterer bind.ContractFilterer) (*PuffinUsersFilterer, error) {
	contract, err := bindPuffinUsers(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PuffinUsersFilterer{contract: contract}, nil
}

// bindPuffinUsers binds a generic wrapper to an already deployed contract.
func bindPuffinUsers(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PuffinUsersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PuffinUsers *PuffinUsersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PuffinUsers.Contract.PuffinUsersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PuffinUsers *PuffinUsersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinUsers.Contract.PuffinUsersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PuffinUsers *PuffinUsersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PuffinUsers.Contract.PuffinUsersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PuffinUsers *PuffinUsersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PuffinUsers.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PuffinUsers *PuffinUsersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinUsers.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PuffinUsers *PuffinUsersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PuffinUsers.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PuffinUsers *PuffinUsersCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinUsers.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PuffinUsers *PuffinUsersSession) Admin() (common.Address, error) {
	return _PuffinUsers.Contract.Admin(&_PuffinUsers.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_PuffinUsers *PuffinUsersCallerSession) Admin() (common.Address, error) {
	return _PuffinUsers.Contract.Admin(&_PuffinUsers.CallOpts)
}

// CanView is a free data retrieval call binding the contract method 0x50ef672d.
//
// Solidity: function canView(address ) view returns(bool)
func (_PuffinUsers *PuffinUsersCaller) CanView(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _PuffinUsers.contract.Call(opts, &out, "canView", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanView is a free data retrieval call binding the contract method 0x50ef672d.
//
// Solidity: function canView(address ) view returns(bool)
func (_PuffinUsers *PuffinUsersSession) CanView(arg0 common.Address) (bool, error) {
	return _PuffinUsers.Contract.CanView(&_PuffinUsers.CallOpts, arg0)
}

// CanView is a free data retrieval call binding the contract method 0x50ef672d.
//
// Solidity: function canView(address ) view returns(bool)
func (_PuffinUsers *PuffinUsersCallerSession) CanView(arg0 common.Address) (bool, error) {
	return _PuffinUsers.Contract.CanView(&_PuffinUsers.CallOpts, arg0)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_PuffinUsers *PuffinUsersCaller) Count(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PuffinUsers.contract.Call(opts, &out, "count")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_PuffinUsers *PuffinUsersSession) Count() (*big.Int, error) {
	return _PuffinUsers.Contract.Count(&_PuffinUsers.CallOpts)
}

// Count is a free data retrieval call binding the contract method 0x06661abd.
//
// Solidity: function count() view returns(uint256)
func (_PuffinUsers *PuffinUsersCallerSession) Count() (*big.Int, error) {
	return _PuffinUsers.Contract.Count(&_PuffinUsers.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinUsers *PuffinUsersCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PuffinUsers.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinUsers *PuffinUsersSession) Owner() (common.Address, error) {
	return _PuffinUsers.Contract.Owner(&_PuffinUsers.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PuffinUsers *PuffinUsersCallerSession) Owner() (common.Address, error) {
	return _PuffinUsers.Contract.Owner(&_PuffinUsers.CallOpts)
}

// Status is a free data retrieval call binding the contract method 0x645b8b1b.
//
// Solidity: function status(address _user) view returns(uint256, bool)
func (_PuffinUsers *PuffinUsersCaller) Status(opts *bind.CallOpts, _user common.Address) (*big.Int, bool, error) {
	var out []interface{}
	err := _PuffinUsers.contract.Call(opts, &out, "status", _user)

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
func (_PuffinUsers *PuffinUsersSession) Status(_user common.Address) (*big.Int, bool, error) {
	return _PuffinUsers.Contract.Status(&_PuffinUsers.CallOpts, _user)
}

// Status is a free data retrieval call binding the contract method 0x645b8b1b.
//
// Solidity: function status(address _user) view returns(uint256, bool)
func (_PuffinUsers *PuffinUsersCallerSession) Status(_user common.Address) (*big.Int, bool, error) {
	return _PuffinUsers.Contract.Status(&_PuffinUsers.CallOpts, _user)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinUsers *PuffinUsersTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PuffinUsers.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinUsers *PuffinUsersSession) RenounceOwnership() (*types.Transaction, error) {
	return _PuffinUsers.Contract.RenounceOwnership(&_PuffinUsers.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PuffinUsers *PuffinUsersTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PuffinUsers.Contract.RenounceOwnership(&_PuffinUsers.TransactOpts)
}

// SetCanView is a paid mutator transaction binding the contract method 0xd7246dc5.
//
// Solidity: function setCanView(address _user, bool _status) returns()
func (_PuffinUsers *PuffinUsersTransactor) SetCanView(opts *bind.TransactOpts, _user common.Address, _status bool) (*types.Transaction, error) {
	return _PuffinUsers.contract.Transact(opts, "setCanView", _user, _status)
}

// SetCanView is a paid mutator transaction binding the contract method 0xd7246dc5.
//
// Solidity: function setCanView(address _user, bool _status) returns()
func (_PuffinUsers *PuffinUsersSession) SetCanView(_user common.Address, _status bool) (*types.Transaction, error) {
	return _PuffinUsers.Contract.SetCanView(&_PuffinUsers.TransactOpts, _user, _status)
}

// SetCanView is a paid mutator transaction binding the contract method 0xd7246dc5.
//
// Solidity: function setCanView(address _user, bool _status) returns()
func (_PuffinUsers *PuffinUsersTransactorSession) SetCanView(_user common.Address, _status bool) (*types.Transaction, error) {
	return _PuffinUsers.Contract.SetCanView(&_PuffinUsers.TransactOpts, _user, _status)
}

// SetStatus is a paid mutator transaction binding the contract method 0x4f2ef3b3.
//
// Solidity: function setStatus(address _user, bool _status, uint256 _val) returns()
func (_PuffinUsers *PuffinUsersTransactor) SetStatus(opts *bind.TransactOpts, _user common.Address, _status bool, _val *big.Int) (*types.Transaction, error) {
	return _PuffinUsers.contract.Transact(opts, "setStatus", _user, _status, _val)
}

// SetStatus is a paid mutator transaction binding the contract method 0x4f2ef3b3.
//
// Solidity: function setStatus(address _user, bool _status, uint256 _val) returns()
func (_PuffinUsers *PuffinUsersSession) SetStatus(_user common.Address, _status bool, _val *big.Int) (*types.Transaction, error) {
	return _PuffinUsers.Contract.SetStatus(&_PuffinUsers.TransactOpts, _user, _status, _val)
}

// SetStatus is a paid mutator transaction binding the contract method 0x4f2ef3b3.
//
// Solidity: function setStatus(address _user, bool _status, uint256 _val) returns()
func (_PuffinUsers *PuffinUsersTransactorSession) SetStatus(_user common.Address, _status bool, _val *big.Int) (*types.Transaction, error) {
	return _PuffinUsers.Contract.SetStatus(&_PuffinUsers.TransactOpts, _user, _status, _val)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinUsers *PuffinUsersTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PuffinUsers.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinUsers *PuffinUsersSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PuffinUsers.Contract.TransferOwnership(&_PuffinUsers.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PuffinUsers *PuffinUsersTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PuffinUsers.Contract.TransferOwnership(&_PuffinUsers.TransactOpts, newOwner)
}

// PuffinUsersNewUserIterator is returned from FilterNewUser and is used to iterate over the raw logs and unpacked data for NewUser events raised by the PuffinUsers contract.
type PuffinUsersNewUserIterator struct {
	Event *PuffinUsersNewUser // Event containing the contract specifics and raw log

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
func (it *PuffinUsersNewUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinUsersNewUser)
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
		it.Event = new(PuffinUsersNewUser)
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
func (it *PuffinUsersNewUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinUsersNewUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinUsersNewUser represents a NewUser event raised by the PuffinUsers contract.
type PuffinUsersNewUser struct {
	User  common.Address
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNewUser is a free log retrieval operation binding the contract event 0xd75b88a9203b17f9356ca063241beac16a25d5a46f485b378c2c229e864bdd4d.
//
// Solidity: event NewUser(address indexed user, uint256 indexed count)
func (_PuffinUsers *PuffinUsersFilterer) FilterNewUser(opts *bind.FilterOpts, user []common.Address, count []*big.Int) (*PuffinUsersNewUserIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var countRule []interface{}
	for _, countItem := range count {
		countRule = append(countRule, countItem)
	}

	logs, sub, err := _PuffinUsers.contract.FilterLogs(opts, "NewUser", userRule, countRule)
	if err != nil {
		return nil, err
	}
	return &PuffinUsersNewUserIterator{contract: _PuffinUsers.contract, event: "NewUser", logs: logs, sub: sub}, nil
}

// WatchNewUser is a free log subscription operation binding the contract event 0xd75b88a9203b17f9356ca063241beac16a25d5a46f485b378c2c229e864bdd4d.
//
// Solidity: event NewUser(address indexed user, uint256 indexed count)
func (_PuffinUsers *PuffinUsersFilterer) WatchNewUser(opts *bind.WatchOpts, sink chan<- *PuffinUsersNewUser, user []common.Address, count []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var countRule []interface{}
	for _, countItem := range count {
		countRule = append(countRule, countItem)
	}

	logs, sub, err := _PuffinUsers.contract.WatchLogs(opts, "NewUser", userRule, countRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinUsersNewUser)
				if err := _PuffinUsers.contract.UnpackLog(event, "NewUser", log); err != nil {
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
func (_PuffinUsers *PuffinUsersFilterer) ParseNewUser(log types.Log) (*PuffinUsersNewUser, error) {
	event := new(PuffinUsersNewUser)
	if err := _PuffinUsers.contract.UnpackLog(event, "NewUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinUsersOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PuffinUsers contract.
type PuffinUsersOwnershipTransferredIterator struct {
	Event *PuffinUsersOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PuffinUsersOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinUsersOwnershipTransferred)
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
		it.Event = new(PuffinUsersOwnershipTransferred)
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
func (it *PuffinUsersOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinUsersOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinUsersOwnershipTransferred represents a OwnershipTransferred event raised by the PuffinUsers contract.
type PuffinUsersOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PuffinUsers *PuffinUsersFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PuffinUsersOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PuffinUsers.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PuffinUsersOwnershipTransferredIterator{contract: _PuffinUsers.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PuffinUsers *PuffinUsersFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PuffinUsersOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PuffinUsers.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinUsersOwnershipTransferred)
				if err := _PuffinUsers.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PuffinUsers *PuffinUsersFilterer) ParseOwnershipTransferred(log types.Log) (*PuffinUsersOwnershipTransferred, error) {
	event := new(PuffinUsersOwnershipTransferred)
	if err := _PuffinUsers.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PuffinUsersUserRemovedIterator is returned from FilterUserRemoved and is used to iterate over the raw logs and unpacked data for UserRemoved events raised by the PuffinUsers contract.
type PuffinUsersUserRemovedIterator struct {
	Event *PuffinUsersUserRemoved // Event containing the contract specifics and raw log

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
func (it *PuffinUsersUserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PuffinUsersUserRemoved)
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
		it.Event = new(PuffinUsersUserRemoved)
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
func (it *PuffinUsersUserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PuffinUsersUserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PuffinUsersUserRemoved represents a UserRemoved event raised by the PuffinUsers contract.
type PuffinUsersUserRemoved struct {
	User  common.Address
	Count *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterUserRemoved is a free log retrieval operation binding the contract event 0xf62c903ca15668006710c4411b74e3acb9efabe0d10bece1a2912050499f7efd.
//
// Solidity: event UserRemoved(address indexed user, uint256 indexed count)
func (_PuffinUsers *PuffinUsersFilterer) FilterUserRemoved(opts *bind.FilterOpts, user []common.Address, count []*big.Int) (*PuffinUsersUserRemovedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var countRule []interface{}
	for _, countItem := range count {
		countRule = append(countRule, countItem)
	}

	logs, sub, err := _PuffinUsers.contract.FilterLogs(opts, "UserRemoved", userRule, countRule)
	if err != nil {
		return nil, err
	}
	return &PuffinUsersUserRemovedIterator{contract: _PuffinUsers.contract, event: "UserRemoved", logs: logs, sub: sub}, nil
}

// WatchUserRemoved is a free log subscription operation binding the contract event 0xf62c903ca15668006710c4411b74e3acb9efabe0d10bece1a2912050499f7efd.
//
// Solidity: event UserRemoved(address indexed user, uint256 indexed count)
func (_PuffinUsers *PuffinUsersFilterer) WatchUserRemoved(opts *bind.WatchOpts, sink chan<- *PuffinUsersUserRemoved, user []common.Address, count []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var countRule []interface{}
	for _, countItem := range count {
		countRule = append(countRule, countItem)
	}

	logs, sub, err := _PuffinUsers.contract.WatchLogs(opts, "UserRemoved", userRule, countRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PuffinUsersUserRemoved)
				if err := _PuffinUsers.contract.UnpackLog(event, "UserRemoved", log); err != nil {
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
func (_PuffinUsers *PuffinUsersFilterer) ParseUserRemoved(log types.Log) (*PuffinUsersUserRemoved, error) {
	event := new(PuffinUsersUserRemoved)
	if err := _PuffinUsers.contract.UnpackLog(event, "UserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
