// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package VRFRequestor

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

// VRFRequestorMetaData contains all meta data concerning the VRFRequestor contract.
var VRFRequestorMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"counter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"request\",\"inputs\":[{\"name\":\"numOfValues\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"request\",\"inputs\":[{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"numOfValues\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requests\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"numOfValues\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"RequestReceived\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"numOfValues\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610a378061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061004a575f3560e01c806361bc221a1461004e57806381d12c581461006c578063d845a4b31461009d578063dc6ad19e146100cd575b5f5ffd5b6100566100fd565b60405161006391906103a3565b60405180910390f35b610086600480360381019061008191906103ee565b610102565b604051610094929190610489565b60405180910390f35b6100b760048036038101906100b291906103ee565b6101a8565b6040516100c491906103a3565b60405180910390f35b6100e760048036038101906100e29190610518565b61027c565b6040516100f491906103a3565b60405180910390f35b5f5481565b6001602052805f5260405f205f91509050805f018054610121906105a2565b80601f016020809104026020016040519081016040528092919081815260200182805461014d906105a2565b80156101985780601f1061016f57610100808354040283529160200191610198565b820191905f5260205f20905b81548152906001019060200180831161017b57829003601f168201915b5050505050908060010154905082565b5f335f5f8154809291906101bb906105ff565b919050556040516020016101d0929190610685565b604051602081830303815290604052805190602001205f1c9050604051806040016040528060405180602001604052805f81525081526020018381525060015f8381526020019081526020015f205f820151815f0190816102319190610879565b5060208201518160010155905050807f8769c91727bbf50934fb8b59dd2f7df4d307389cd4a7af7f79ec11cd47f7fad78360405161026f919061096b565b60405180910390a2919050565b5f335f5f81548092919061028f906105ff565b919050556040516020016102a4929190610685565b604051602081830303815290604052805190602001205f1c9050604051806040016040528085858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018381525060015f8381526020019081526020015f205f820151815f01908161033a9190610879565b5060208201518160010155905050807f8769c91727bbf50934fb8b59dd2f7df4d307389cd4a7af7f79ec11cd47f7fad785858560405161037c939291906109d1565b60405180910390a29392505050565b5f819050919050565b61039d8161038b565b82525050565b5f6020820190506103b65f830184610394565b92915050565b5f5ffd5b5f5ffd5b6103cd8161038b565b81146103d7575f5ffd5b50565b5f813590506103e8816103c4565b92915050565b5f60208284031215610403576104026103bc565b5b5f610410848285016103da565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61045b82610419565b6104658185610423565b9350610475818560208601610433565b61047e81610441565b840191505092915050565b5f6040820190508181035f8301526104a18185610451565b90506104b06020830184610394565b9392505050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126104d8576104d76104b7565b5b8235905067ffffffffffffffff8111156104f5576104f46104bb565b5b602083019150836001820283011115610511576105106104bf565b5b9250929050565b5f5f5f6040848603121561052f5761052e6103bc565b5b5f84013567ffffffffffffffff81111561054c5761054b6103c0565b5b610558868287016104c3565b9350935050602061056b868287016103da565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806105b957607f821691505b6020821081036105cc576105cb610575565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6106098261038b565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361063b5761063a6105d2565b5b600182019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61066f82610646565b9050919050565b61067f81610665565b82525050565b5f6040820190506106985f830185610676565b6106a56020830184610394565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026107357fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826106fa565b61073f86836106fa565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61077a6107756107708461038b565b610757565b61038b565b9050919050565b5f819050919050565b61079383610760565b6107a761079f82610781565b848454610706565b825550505050565b5f5f905090565b6107be6107af565b6107c981848461078a565b505050565b5b818110156107ec576107e15f826107b6565b6001810190506107cf565b5050565b601f82111561083157610802816106d9565b61080b846106eb565b8101602085101561081a578190505b61082e610826856106eb565b8301826107ce565b50505b505050565b5f82821c905092915050565b5f6108515f1984600802610836565b1980831691505092915050565b5f6108698383610842565b9150826002028217905092915050565b61088282610419565b67ffffffffffffffff81111561089b5761089a6106ac565b5b6108a582546105a2565b6108b08282856107f0565b5f60209050601f8311600181146108e1575f84156108cf578287015190505b6108d9858261085e565b865550610940565b601f1984166108ef866106d9565b5f5b82811015610916578489015182556001820191506020850194506020810190506108f1565b86831015610933578489015161092f601f891682610842565b8355505b6001600288020188555050505b505050505050565b50565b5f6109565f83610423565b915061096182610948565b5f82019050919050565b5f6040820190508181035f8301526109828161094b565b90506109916020830184610394565b92915050565b828183375f83830152505050565b5f6109b08385610423565b93506109bd838584610997565b6109c683610441565b840190509392505050565b5f6040820190508181035f8301526109ea8185876109a5565b90506109f96020830184610394565b94935050505056fea264697066735822122033e1a8efc2f954406e36d9c6d3e4a35668b7f4464604c9ce2c07c27c0a8d180a64736f6c634300081c0033",
}

// VRFRequestorABI is the input ABI used to generate the binding from.
// Deprecated: Use VRFRequestorMetaData.ABI instead.
var VRFRequestorABI = VRFRequestorMetaData.ABI

// VRFRequestorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VRFRequestorMetaData.Bin instead.
var VRFRequestorBin = VRFRequestorMetaData.Bin

// DeployVRFRequestor deploys a new Ethereum contract, binding an instance of VRFRequestor to it.
func DeployVRFRequestor(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VRFRequestor, error) {
	parsed, err := VRFRequestorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VRFRequestorBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VRFRequestor{VRFRequestorCaller: VRFRequestorCaller{contract: contract}, VRFRequestorTransactor: VRFRequestorTransactor{contract: contract}, VRFRequestorFilterer: VRFRequestorFilterer{contract: contract}}, nil
}

// VRFRequestor is an auto generated Go binding around an Ethereum contract.
type VRFRequestor struct {
	VRFRequestorCaller     // Read-only binding to the contract
	VRFRequestorTransactor // Write-only binding to the contract
	VRFRequestorFilterer   // Log filterer for contract events
}

// VRFRequestorCaller is an auto generated read-only Go binding around an Ethereum contract.
type VRFRequestorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFRequestorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VRFRequestorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFRequestorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VRFRequestorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VRFRequestorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VRFRequestorSession struct {
	Contract     *VRFRequestor     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VRFRequestorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VRFRequestorCallerSession struct {
	Contract *VRFRequestorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// VRFRequestorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VRFRequestorTransactorSession struct {
	Contract     *VRFRequestorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// VRFRequestorRaw is an auto generated low-level Go binding around an Ethereum contract.
type VRFRequestorRaw struct {
	Contract *VRFRequestor // Generic contract binding to access the raw methods on
}

// VRFRequestorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VRFRequestorCallerRaw struct {
	Contract *VRFRequestorCaller // Generic read-only contract binding to access the raw methods on
}

// VRFRequestorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VRFRequestorTransactorRaw struct {
	Contract *VRFRequestorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVRFRequestor creates a new instance of VRFRequestor, bound to a specific deployed contract.
func NewVRFRequestor(address common.Address, backend bind.ContractBackend) (*VRFRequestor, error) {
	contract, err := bindVRFRequestor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VRFRequestor{VRFRequestorCaller: VRFRequestorCaller{contract: contract}, VRFRequestorTransactor: VRFRequestorTransactor{contract: contract}, VRFRequestorFilterer: VRFRequestorFilterer{contract: contract}}, nil
}

// NewVRFRequestorCaller creates a new read-only instance of VRFRequestor, bound to a specific deployed contract.
func NewVRFRequestorCaller(address common.Address, caller bind.ContractCaller) (*VRFRequestorCaller, error) {
	contract, err := bindVRFRequestor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VRFRequestorCaller{contract: contract}, nil
}

// NewVRFRequestorTransactor creates a new write-only instance of VRFRequestor, bound to a specific deployed contract.
func NewVRFRequestorTransactor(address common.Address, transactor bind.ContractTransactor) (*VRFRequestorTransactor, error) {
	contract, err := bindVRFRequestor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VRFRequestorTransactor{contract: contract}, nil
}

// NewVRFRequestorFilterer creates a new log filterer instance of VRFRequestor, bound to a specific deployed contract.
func NewVRFRequestorFilterer(address common.Address, filterer bind.ContractFilterer) (*VRFRequestorFilterer, error) {
	contract, err := bindVRFRequestor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VRFRequestorFilterer{contract: contract}, nil
}

// bindVRFRequestor binds a generic wrapper to an already deployed contract.
func bindVRFRequestor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VRFRequestorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFRequestor *VRFRequestorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFRequestor.Contract.VRFRequestorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFRequestor *VRFRequestorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFRequestor.Contract.VRFRequestorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFRequestor *VRFRequestorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFRequestor.Contract.VRFRequestorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VRFRequestor *VRFRequestorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VRFRequestor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VRFRequestor *VRFRequestorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VRFRequestor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VRFRequestor *VRFRequestorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VRFRequestor.Contract.contract.Transact(opts, method, params...)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_VRFRequestor *VRFRequestorCaller) Counter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VRFRequestor.contract.Call(opts, &out, "counter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_VRFRequestor *VRFRequestorSession) Counter() (*big.Int, error) {
	return _VRFRequestor.Contract.Counter(&_VRFRequestor.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_VRFRequestor *VRFRequestorCallerSession) Counter() (*big.Int, error) {
	return _VRFRequestor.Contract.Counter(&_VRFRequestor.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(string seed, uint256 numOfValues)
func (_VRFRequestor *VRFRequestorCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seed        string
	NumOfValues *big.Int
}, error) {
	var out []interface{}
	err := _VRFRequestor.contract.Call(opts, &out, "requests", arg0)

	outstruct := new(struct {
		Seed        string
		NumOfValues *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seed = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.NumOfValues = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(string seed, uint256 numOfValues)
func (_VRFRequestor *VRFRequestorSession) Requests(arg0 *big.Int) (struct {
	Seed        string
	NumOfValues *big.Int
}, error) {
	return _VRFRequestor.Contract.Requests(&_VRFRequestor.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(string seed, uint256 numOfValues)
func (_VRFRequestor *VRFRequestorCallerSession) Requests(arg0 *big.Int) (struct {
	Seed        string
	NumOfValues *big.Int
}, error) {
	return _VRFRequestor.Contract.Requests(&_VRFRequestor.CallOpts, arg0)
}

// Request is a paid mutator transaction binding the contract method 0xd845a4b3.
//
// Solidity: function request(uint256 numOfValues) returns(uint256 requestId)
func (_VRFRequestor *VRFRequestorTransactor) Request(opts *bind.TransactOpts, numOfValues *big.Int) (*types.Transaction, error) {
	return _VRFRequestor.contract.Transact(opts, "request", numOfValues)
}

// Request is a paid mutator transaction binding the contract method 0xd845a4b3.
//
// Solidity: function request(uint256 numOfValues) returns(uint256 requestId)
func (_VRFRequestor *VRFRequestorSession) Request(numOfValues *big.Int) (*types.Transaction, error) {
	return _VRFRequestor.Contract.Request(&_VRFRequestor.TransactOpts, numOfValues)
}

// Request is a paid mutator transaction binding the contract method 0xd845a4b3.
//
// Solidity: function request(uint256 numOfValues) returns(uint256 requestId)
func (_VRFRequestor *VRFRequestorTransactorSession) Request(numOfValues *big.Int) (*types.Transaction, error) {
	return _VRFRequestor.Contract.Request(&_VRFRequestor.TransactOpts, numOfValues)
}

// Request0 is a paid mutator transaction binding the contract method 0xdc6ad19e.
//
// Solidity: function request(string seed, uint256 numOfValues) returns(uint256 requestId)
func (_VRFRequestor *VRFRequestorTransactor) Request0(opts *bind.TransactOpts, seed string, numOfValues *big.Int) (*types.Transaction, error) {
	return _VRFRequestor.contract.Transact(opts, "request0", seed, numOfValues)
}

// Request0 is a paid mutator transaction binding the contract method 0xdc6ad19e.
//
// Solidity: function request(string seed, uint256 numOfValues) returns(uint256 requestId)
func (_VRFRequestor *VRFRequestorSession) Request0(seed string, numOfValues *big.Int) (*types.Transaction, error) {
	return _VRFRequestor.Contract.Request0(&_VRFRequestor.TransactOpts, seed, numOfValues)
}

// Request0 is a paid mutator transaction binding the contract method 0xdc6ad19e.
//
// Solidity: function request(string seed, uint256 numOfValues) returns(uint256 requestId)
func (_VRFRequestor *VRFRequestorTransactorSession) Request0(seed string, numOfValues *big.Int) (*types.Transaction, error) {
	return _VRFRequestor.Contract.Request0(&_VRFRequestor.TransactOpts, seed, numOfValues)
}

// VRFRequestorRequestReceivedIterator is returned from FilterRequestReceived and is used to iterate over the raw logs and unpacked data for RequestReceived events raised by the VRFRequestor contract.
type VRFRequestorRequestReceivedIterator struct {
	Event *VRFRequestorRequestReceived // Event containing the contract specifics and raw log

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
func (it *VRFRequestorRequestReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VRFRequestorRequestReceived)
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
		it.Event = new(VRFRequestorRequestReceived)
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
func (it *VRFRequestorRequestReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VRFRequestorRequestReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VRFRequestorRequestReceived represents a RequestReceived event raised by the VRFRequestor contract.
type VRFRequestorRequestReceived struct {
	RequestId   *big.Int
	Seed        string
	NumOfValues *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRequestReceived is a free log retrieval operation binding the contract event 0x8769c91727bbf50934fb8b59dd2f7df4d307389cd4a7af7f79ec11cd47f7fad7.
//
// Solidity: event RequestReceived(uint256 indexed requestId, string seed, uint256 numOfValues)
func (_VRFRequestor *VRFRequestorFilterer) FilterRequestReceived(opts *bind.FilterOpts, requestId []*big.Int) (*VRFRequestorRequestReceivedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRFRequestor.contract.FilterLogs(opts, "RequestReceived", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &VRFRequestorRequestReceivedIterator{contract: _VRFRequestor.contract, event: "RequestReceived", logs: logs, sub: sub}, nil
}

// WatchRequestReceived is a free log subscription operation binding the contract event 0x8769c91727bbf50934fb8b59dd2f7df4d307389cd4a7af7f79ec11cd47f7fad7.
//
// Solidity: event RequestReceived(uint256 indexed requestId, string seed, uint256 numOfValues)
func (_VRFRequestor *VRFRequestorFilterer) WatchRequestReceived(opts *bind.WatchOpts, sink chan<- *VRFRequestorRequestReceived, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _VRFRequestor.contract.WatchLogs(opts, "RequestReceived", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VRFRequestorRequestReceived)
				if err := _VRFRequestor.contract.UnpackLog(event, "RequestReceived", log); err != nil {
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

// ParseRequestReceived is a log parse operation binding the contract event 0x8769c91727bbf50934fb8b59dd2f7df4d307389cd4a7af7f79ec11cd47f7fad7.
//
// Solidity: event RequestReceived(uint256 indexed requestId, string seed, uint256 numOfValues)
func (_VRFRequestor *VRFRequestorFilterer) ParseRequestReceived(log types.Log) (*VRFRequestorRequestReceived, error) {
	event := new(VRFRequestorRequestReceived)
	if err := _VRFRequestor.contract.UnpackLog(event, "RequestReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
