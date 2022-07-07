// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package totoro

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

// EventTestMetaData contains all meta data concerning the EventTest contract.
var EventTestMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Event\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"triggerEvt\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// EventTestABI is the input ABI used to generate the binding from.
// Deprecated: Use EventTestMetaData.ABI instead.
var EventTestABI = EventTestMetaData.ABI

// EventTest is an auto generated Go binding around an Ethereum contract.
type EventTest struct {
	EventTestCaller     // Read-only binding to the contract
	EventTestTransactor // Write-only binding to the contract
	EventTestFilterer   // Log filterer for contract events
}

// EventTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventTestSession struct {
	Contract     *EventTest        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventTestCallerSession struct {
	Contract *EventTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EventTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventTestTransactorSession struct {
	Contract     *EventTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EventTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventTestRaw struct {
	Contract *EventTest // Generic contract binding to access the raw methods on
}

// EventTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventTestCallerRaw struct {
	Contract *EventTestCaller // Generic read-only contract binding to access the raw methods on
}

// EventTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventTestTransactorRaw struct {
	Contract *EventTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEventTest creates a new instance of EventTest, bound to a specific deployed contract.
func NewEventTest(address common.Address, backend bind.ContractBackend) (*EventTest, error) {
	contract, err := bindEventTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EventTest{EventTestCaller: EventTestCaller{contract: contract}, EventTestTransactor: EventTestTransactor{contract: contract}, EventTestFilterer: EventTestFilterer{contract: contract}}, nil
}

// NewEventTestCaller creates a new read-only instance of EventTest, bound to a specific deployed contract.
func NewEventTestCaller(address common.Address, caller bind.ContractCaller) (*EventTestCaller, error) {
	contract, err := bindEventTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventTestCaller{contract: contract}, nil
}

// NewEventTestTransactor creates a new write-only instance of EventTest, bound to a specific deployed contract.
func NewEventTestTransactor(address common.Address, transactor bind.ContractTransactor) (*EventTestTransactor, error) {
	contract, err := bindEventTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventTestTransactor{contract: contract}, nil
}

// NewEventTestFilterer creates a new log filterer instance of EventTest, bound to a specific deployed contract.
func NewEventTestFilterer(address common.Address, filterer bind.ContractFilterer) (*EventTestFilterer, error) {
	contract, err := bindEventTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventTestFilterer{contract: contract}, nil
}

// bindEventTest binds a generic wrapper to an already deployed contract.
func bindEventTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EventTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventTest *EventTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventTest.Contract.EventTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventTest *EventTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventTest.Contract.EventTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventTest *EventTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventTest.Contract.EventTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EventTest *EventTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EventTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EventTest *EventTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EventTest *EventTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EventTest.Contract.contract.Transact(opts, method, params...)
}

// TriggerEvt is a paid mutator transaction binding the contract method 0xc18998b1.
//
// Solidity: function triggerEvt() returns()
func (_EventTest *EventTestTransactor) TriggerEvt(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EventTest.contract.Transact(opts, "triggerEvt")
}

// TriggerEvt is a paid mutator transaction binding the contract method 0xc18998b1.
//
// Solidity: function triggerEvt() returns()
func (_EventTest *EventTestSession) TriggerEvt() (*types.Transaction, error) {
	return _EventTest.Contract.TriggerEvt(&_EventTest.TransactOpts)
}

// TriggerEvt is a paid mutator transaction binding the contract method 0xc18998b1.
//
// Solidity: function triggerEvt() returns()
func (_EventTest *EventTestTransactorSession) TriggerEvt() (*types.Transaction, error) {
	return _EventTest.Contract.TriggerEvt(&_EventTest.TransactOpts)
}

// EventTestEventIterator is returned from FilterEvent and is used to iterate over the raw logs and unpacked data for Event events raised by the EventTest contract.
type EventTestEventIterator struct {
	Event *EventTestEvent // Event containing the contract specifics and raw log

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
func (it *EventTestEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventTestEvent)
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
		it.Event = new(EventTestEvent)
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
func (it *EventTestEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventTestEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventTestEvent represents a Event event raised by the EventTest contract.
type EventTestEvent struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterEvent is a free log retrieval operation binding the contract event 0x510e730eb6600b4c67d51768c6996795863364461fee983d92d5e461f209c7cf.
//
// Solidity: event Event(uint256 arg0)
func (_EventTest *EventTestFilterer) FilterEvent(opts *bind.FilterOpts) (*EventTestEventIterator, error) {

	logs, sub, err := _EventTest.contract.FilterLogs(opts, "Event")
	if err != nil {
		return nil, err
	}
	return &EventTestEventIterator{contract: _EventTest.contract, event: "Event", logs: logs, sub: sub}, nil
}

// WatchEvent is a free log subscription operation binding the contract event 0x510e730eb6600b4c67d51768c6996795863364461fee983d92d5e461f209c7cf.
//
// Solidity: event Event(uint256 arg0)
func (_EventTest *EventTestFilterer) WatchEvent(opts *bind.WatchOpts, sink chan<- *EventTestEvent) (event.Subscription, error) {

	logs, sub, err := _EventTest.contract.WatchLogs(opts, "Event")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventTestEvent)
				if err := _EventTest.contract.UnpackLog(event, "Event", log); err != nil {
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

// ParseEvent is a log parse operation binding the contract event 0x510e730eb6600b4c67d51768c6996795863364461fee983d92d5e461f209c7cf.
//
// Solidity: event Event(uint256 arg0)
func (_EventTest *EventTestFilterer) ParseEvent(log types.Log) (*EventTestEvent, error) {
	event := new(EventTestEvent)
	if err := _EventTest.contract.UnpackLog(event, "Event", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
