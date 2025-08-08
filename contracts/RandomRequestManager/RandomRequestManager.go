// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RandomRequestManager

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

// RandomRequestManagerMetaData contains all meta data concerning the RandomRequestManager contract.
var RandomRequestManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"counter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fulfillRandomInt\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"randomInts\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"randomness\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"request\",\"inputs\":[{\"name\":\"numOfValues\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"request\",\"inputs\":[{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"numOfValues\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requests\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"numOfValues\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"RequestReceived\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"numOfValues\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610c718061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c80633ccb1e381461006457806361bc221a1461008057806381d12c581461009e578063b4844fc7146100cf578063d845a4b3146100ff578063dc6ad19e1461012f575b5f5ffd5b61007e60048036038101906100799190610540565b61015f565b005b610088610185565b60405161009591906105ac565b60405180910390f35b6100b860048036038101906100b391906105c5565b61018a565b6040516100c6929190610660565b60405180910390f35b6100e960048036038101906100e4919061068e565b610230565b6040516100f691906106e4565b60405180910390f35b610119600480360381019061011491906105c5565b61025b565b60405161012691906105ac565b60405180910390f35b61014960048036038101906101449190610752565b61032f565b60405161015691906105ac565b60405180910390f35b818160025f8681526020019081526020015f20919061017f92919061043e565b50505050565b5f5481565b6001602052805f5260405f205f91509050805f0180546101a9906107dc565b80601f01602080910402602001604051908101604052809291908181526020018280546101d5906107dc565b80156102205780601f106101f757610100808354040283529160200191610220565b820191905f5260205f20905b81548152906001019060200180831161020357829003601f168201915b5050505050908060010154905082565b6002602052815f5260405f208181548110610249575f80fd5b905f5260205f20015f91509150505481565b5f335f5f81548092919061026e90610839565b919050556040516020016102839291906108bf565b604051602081830303815290604052805190602001205f1c9050604051806040016040528060405180602001604052805f81525081526020018381525060015f8381526020019081526020015f205f820151815f0190816102e49190610ab3565b5060208201518160010155905050807f8769c91727bbf50934fb8b59dd2f7df4d307389cd4a7af7f79ec11cd47f7fad7836040516103229190610ba5565b60405180910390a2919050565b5f335f5f81548092919061034290610839565b919050556040516020016103579291906108bf565b604051602081830303815290604052805190602001205f1c9050604051806040016040528085858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018381525060015f8381526020019081526020015f205f820151815f0190816103ed9190610ab3565b5060208201518160010155905050807f8769c91727bbf50934fb8b59dd2f7df4d307389cd4a7af7f79ec11cd47f7fad785858560405161042f93929190610c0b565b60405180910390a29392505050565b828054828255905f5260205f20908101928215610478579160200282015b8281111561047757823582559160200191906001019061045c565b5b5090506104859190610489565b5090565b5b808211156104a0575f815f90555060010161048a565b5090565b5f5ffd5b5f5ffd5b5f819050919050565b6104be816104ac565b81146104c8575f5ffd5b50565b5f813590506104d9816104b5565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f840112610500576104ff6104df565b5b8235905067ffffffffffffffff81111561051d5761051c6104e3565b5b602083019150836020820283011115610539576105386104e7565b5b9250929050565b5f5f5f60408486031215610557576105566104a4565b5b5f610564868287016104cb565b935050602084013567ffffffffffffffff811115610585576105846104a8565b5b610591868287016104eb565b92509250509250925092565b6105a6816104ac565b82525050565b5f6020820190506105bf5f83018461059d565b92915050565b5f602082840312156105da576105d96104a4565b5b5f6105e7848285016104cb565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610632826105f0565b61063c81856105fa565b935061064c81856020860161060a565b61065581610618565b840191505092915050565b5f6040820190508181035f8301526106788185610628565b9050610687602083018461059d565b9392505050565b5f5f604083850312156106a4576106a36104a4565b5b5f6106b1858286016104cb565b92505060206106c2858286016104cb565b9150509250929050565b5f819050919050565b6106de816106cc565b82525050565b5f6020820190506106f75f8301846106d5565b92915050565b5f5f83601f840112610712576107116104df565b5b8235905067ffffffffffffffff81111561072f5761072e6104e3565b5b60208301915083600182028301111561074b5761074a6104e7565b5b9250929050565b5f5f5f60408486031215610769576107686104a4565b5b5f84013567ffffffffffffffff811115610786576107856104a8565b5b610792868287016106fd565b935093505060206107a5868287016104cb565b9150509250925092565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806107f357607f821691505b602082108103610806576108056107af565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610843826104ac565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036108755761087461080c565b5b600182019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6108a982610880565b9050919050565b6108b98161089f565b82525050565b5f6040820190506108d25f8301856108b0565b6108df602083018461059d565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261096f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610934565b6109798683610934565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6109b46109af6109aa846104ac565b610991565b6104ac565b9050919050565b5f819050919050565b6109cd8361099a565b6109e16109d9826109bb565b848454610940565b825550505050565b5f5f905090565b6109f86109e9565b610a038184846109c4565b505050565b5b81811015610a2657610a1b5f826109f0565b600181019050610a09565b5050565b601f821115610a6b57610a3c81610913565b610a4584610925565b81016020851015610a54578190505b610a68610a6085610925565b830182610a08565b50505b505050565b5f82821c905092915050565b5f610a8b5f1984600802610a70565b1980831691505092915050565b5f610aa38383610a7c565b9150826002028217905092915050565b610abc826105f0565b67ffffffffffffffff811115610ad557610ad46108e6565b5b610adf82546107dc565b610aea828285610a2a565b5f60209050601f831160018114610b1b575f8415610b09578287015190505b610b138582610a98565b865550610b7a565b601f198416610b2986610913565b5f5b82811015610b5057848901518255600182019150602085019450602081019050610b2b565b86831015610b6d5784890151610b69601f891682610a7c565b8355505b6001600288020188555050505b505050505050565b50565b5f610b905f836105fa565b9150610b9b82610b82565b5f82019050919050565b5f6040820190508181035f830152610bbc81610b85565b9050610bcb602083018461059d565b92915050565b828183375f83830152505050565b5f610bea83856105fa565b9350610bf7838584610bd1565b610c0083610618565b840190509392505050565b5f6040820190508181035f830152610c24818587610bdf565b9050610c33602083018461059d565b94935050505056fea264697066735822122028a5012d7cc9ba0e294e11cdcb8fb9c5b0a0ffa8c4c2f50d8fb5d97cdaf32da564736f6c634300081c0033",
}

// RandomRequestManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use RandomRequestManagerMetaData.ABI instead.
var RandomRequestManagerABI = RandomRequestManagerMetaData.ABI

// RandomRequestManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RandomRequestManagerMetaData.Bin instead.
var RandomRequestManagerBin = RandomRequestManagerMetaData.Bin

// DeployRandomRequestManager deploys a new Ethereum contract, binding an instance of RandomRequestManager to it.
func DeployRandomRequestManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RandomRequestManager, error) {
	parsed, err := RandomRequestManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RandomRequestManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RandomRequestManager{RandomRequestManagerCaller: RandomRequestManagerCaller{contract: contract}, RandomRequestManagerTransactor: RandomRequestManagerTransactor{contract: contract}, RandomRequestManagerFilterer: RandomRequestManagerFilterer{contract: contract}}, nil
}

// RandomRequestManager is an auto generated Go binding around an Ethereum contract.
type RandomRequestManager struct {
	RandomRequestManagerCaller     // Read-only binding to the contract
	RandomRequestManagerTransactor // Write-only binding to the contract
	RandomRequestManagerFilterer   // Log filterer for contract events
}

// RandomRequestManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RandomRequestManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomRequestManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RandomRequestManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomRequestManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RandomRequestManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomRequestManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RandomRequestManagerSession struct {
	Contract     *RandomRequestManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RandomRequestManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RandomRequestManagerCallerSession struct {
	Contract *RandomRequestManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// RandomRequestManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RandomRequestManagerTransactorSession struct {
	Contract     *RandomRequestManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// RandomRequestManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RandomRequestManagerRaw struct {
	Contract *RandomRequestManager // Generic contract binding to access the raw methods on
}

// RandomRequestManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RandomRequestManagerCallerRaw struct {
	Contract *RandomRequestManagerCaller // Generic read-only contract binding to access the raw methods on
}

// RandomRequestManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RandomRequestManagerTransactorRaw struct {
	Contract *RandomRequestManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRandomRequestManager creates a new instance of RandomRequestManager, bound to a specific deployed contract.
func NewRandomRequestManager(address common.Address, backend bind.ContractBackend) (*RandomRequestManager, error) {
	contract, err := bindRandomRequestManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RandomRequestManager{RandomRequestManagerCaller: RandomRequestManagerCaller{contract: contract}, RandomRequestManagerTransactor: RandomRequestManagerTransactor{contract: contract}, RandomRequestManagerFilterer: RandomRequestManagerFilterer{contract: contract}}, nil
}

// NewRandomRequestManagerCaller creates a new read-only instance of RandomRequestManager, bound to a specific deployed contract.
func NewRandomRequestManagerCaller(address common.Address, caller bind.ContractCaller) (*RandomRequestManagerCaller, error) {
	contract, err := bindRandomRequestManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RandomRequestManagerCaller{contract: contract}, nil
}

// NewRandomRequestManagerTransactor creates a new write-only instance of RandomRequestManager, bound to a specific deployed contract.
func NewRandomRequestManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*RandomRequestManagerTransactor, error) {
	contract, err := bindRandomRequestManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RandomRequestManagerTransactor{contract: contract}, nil
}

// NewRandomRequestManagerFilterer creates a new log filterer instance of RandomRequestManager, bound to a specific deployed contract.
func NewRandomRequestManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*RandomRequestManagerFilterer, error) {
	contract, err := bindRandomRequestManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RandomRequestManagerFilterer{contract: contract}, nil
}

// bindRandomRequestManager binds a generic wrapper to an already deployed contract.
func bindRandomRequestManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RandomRequestManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomRequestManager *RandomRequestManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomRequestManager.Contract.RandomRequestManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomRequestManager *RandomRequestManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomRequestManager.Contract.RandomRequestManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomRequestManager *RandomRequestManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomRequestManager.Contract.RandomRequestManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomRequestManager *RandomRequestManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomRequestManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomRequestManager *RandomRequestManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomRequestManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomRequestManager *RandomRequestManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomRequestManager.Contract.contract.Transact(opts, method, params...)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_RandomRequestManager *RandomRequestManagerCaller) Counter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomRequestManager.contract.Call(opts, &out, "counter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_RandomRequestManager *RandomRequestManagerSession) Counter() (*big.Int, error) {
	return _RandomRequestManager.Contract.Counter(&_RandomRequestManager.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_RandomRequestManager *RandomRequestManagerCallerSession) Counter() (*big.Int, error) {
	return _RandomRequestManager.Contract.Counter(&_RandomRequestManager.CallOpts)
}

// Randomness is a free data retrieval call binding the contract method 0xb4844fc7.
//
// Solidity: function randomness(uint256 , uint256 ) view returns(int256)
func (_RandomRequestManager *RandomRequestManagerCaller) Randomness(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RandomRequestManager.contract.Call(opts, &out, "randomness", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Randomness is a free data retrieval call binding the contract method 0xb4844fc7.
//
// Solidity: function randomness(uint256 , uint256 ) view returns(int256)
func (_RandomRequestManager *RandomRequestManagerSession) Randomness(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _RandomRequestManager.Contract.Randomness(&_RandomRequestManager.CallOpts, arg0, arg1)
}

// Randomness is a free data retrieval call binding the contract method 0xb4844fc7.
//
// Solidity: function randomness(uint256 , uint256 ) view returns(int256)
func (_RandomRequestManager *RandomRequestManagerCallerSession) Randomness(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _RandomRequestManager.Contract.Randomness(&_RandomRequestManager.CallOpts, arg0, arg1)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(string seed, uint256 numOfValues)
func (_RandomRequestManager *RandomRequestManagerCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seed        string
	NumOfValues *big.Int
}, error) {
	var out []interface{}
	err := _RandomRequestManager.contract.Call(opts, &out, "requests", arg0)

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
func (_RandomRequestManager *RandomRequestManagerSession) Requests(arg0 *big.Int) (struct {
	Seed        string
	NumOfValues *big.Int
}, error) {
	return _RandomRequestManager.Contract.Requests(&_RandomRequestManager.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(string seed, uint256 numOfValues)
func (_RandomRequestManager *RandomRequestManagerCallerSession) Requests(arg0 *big.Int) (struct {
	Seed        string
	NumOfValues *big.Int
}, error) {
	return _RandomRequestManager.Contract.Requests(&_RandomRequestManager.CallOpts, arg0)
}

// FulfillRandomInt is a paid mutator transaction binding the contract method 0x3ccb1e38.
//
// Solidity: function fulfillRandomInt(uint256 requestId, int256[] randomInts) returns()
func (_RandomRequestManager *RandomRequestManagerTransactor) FulfillRandomInt(opts *bind.TransactOpts, requestId *big.Int, randomInts []*big.Int) (*types.Transaction, error) {
	return _RandomRequestManager.contract.Transact(opts, "fulfillRandomInt", requestId, randomInts)
}

// FulfillRandomInt is a paid mutator transaction binding the contract method 0x3ccb1e38.
//
// Solidity: function fulfillRandomInt(uint256 requestId, int256[] randomInts) returns()
func (_RandomRequestManager *RandomRequestManagerSession) FulfillRandomInt(requestId *big.Int, randomInts []*big.Int) (*types.Transaction, error) {
	return _RandomRequestManager.Contract.FulfillRandomInt(&_RandomRequestManager.TransactOpts, requestId, randomInts)
}

// FulfillRandomInt is a paid mutator transaction binding the contract method 0x3ccb1e38.
//
// Solidity: function fulfillRandomInt(uint256 requestId, int256[] randomInts) returns()
func (_RandomRequestManager *RandomRequestManagerTransactorSession) FulfillRandomInt(requestId *big.Int, randomInts []*big.Int) (*types.Transaction, error) {
	return _RandomRequestManager.Contract.FulfillRandomInt(&_RandomRequestManager.TransactOpts, requestId, randomInts)
}

// Request is a paid mutator transaction binding the contract method 0xd845a4b3.
//
// Solidity: function request(uint256 numOfValues) returns(uint256 requestId)
func (_RandomRequestManager *RandomRequestManagerTransactor) Request(opts *bind.TransactOpts, numOfValues *big.Int) (*types.Transaction, error) {
	return _RandomRequestManager.contract.Transact(opts, "request", numOfValues)
}

// Request is a paid mutator transaction binding the contract method 0xd845a4b3.
//
// Solidity: function request(uint256 numOfValues) returns(uint256 requestId)
func (_RandomRequestManager *RandomRequestManagerSession) Request(numOfValues *big.Int) (*types.Transaction, error) {
	return _RandomRequestManager.Contract.Request(&_RandomRequestManager.TransactOpts, numOfValues)
}

// Request is a paid mutator transaction binding the contract method 0xd845a4b3.
//
// Solidity: function request(uint256 numOfValues) returns(uint256 requestId)
func (_RandomRequestManager *RandomRequestManagerTransactorSession) Request(numOfValues *big.Int) (*types.Transaction, error) {
	return _RandomRequestManager.Contract.Request(&_RandomRequestManager.TransactOpts, numOfValues)
}

// Request0 is a paid mutator transaction binding the contract method 0xdc6ad19e.
//
// Solidity: function request(string seed, uint256 numOfValues) returns(uint256 requestId)
func (_RandomRequestManager *RandomRequestManagerTransactor) Request0(opts *bind.TransactOpts, seed string, numOfValues *big.Int) (*types.Transaction, error) {
	return _RandomRequestManager.contract.Transact(opts, "request0", seed, numOfValues)
}

// Request0 is a paid mutator transaction binding the contract method 0xdc6ad19e.
//
// Solidity: function request(string seed, uint256 numOfValues) returns(uint256 requestId)
func (_RandomRequestManager *RandomRequestManagerSession) Request0(seed string, numOfValues *big.Int) (*types.Transaction, error) {
	return _RandomRequestManager.Contract.Request0(&_RandomRequestManager.TransactOpts, seed, numOfValues)
}

// Request0 is a paid mutator transaction binding the contract method 0xdc6ad19e.
//
// Solidity: function request(string seed, uint256 numOfValues) returns(uint256 requestId)
func (_RandomRequestManager *RandomRequestManagerTransactorSession) Request0(seed string, numOfValues *big.Int) (*types.Transaction, error) {
	return _RandomRequestManager.Contract.Request0(&_RandomRequestManager.TransactOpts, seed, numOfValues)
}

// RandomRequestManagerRequestReceivedIterator is returned from FilterRequestReceived and is used to iterate over the raw logs and unpacked data for RequestReceived events raised by the RandomRequestManager contract.
type RandomRequestManagerRequestReceivedIterator struct {
	Event *RandomRequestManagerRequestReceived // Event containing the contract specifics and raw log

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
func (it *RandomRequestManagerRequestReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomRequestManagerRequestReceived)
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
		it.Event = new(RandomRequestManagerRequestReceived)
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
func (it *RandomRequestManagerRequestReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomRequestManagerRequestReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomRequestManagerRequestReceived represents a RequestReceived event raised by the RandomRequestManager contract.
type RandomRequestManagerRequestReceived struct {
	RequestId   *big.Int
	Seed        string
	NumOfValues *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRequestReceived is a free log retrieval operation binding the contract event 0x8769c91727bbf50934fb8b59dd2f7df4d307389cd4a7af7f79ec11cd47f7fad7.
//
// Solidity: event RequestReceived(uint256 indexed requestId, string seed, uint256 numOfValues)
func (_RandomRequestManager *RandomRequestManagerFilterer) FilterRequestReceived(opts *bind.FilterOpts, requestId []*big.Int) (*RandomRequestManagerRequestReceivedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _RandomRequestManager.contract.FilterLogs(opts, "RequestReceived", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &RandomRequestManagerRequestReceivedIterator{contract: _RandomRequestManager.contract, event: "RequestReceived", logs: logs, sub: sub}, nil
}

// WatchRequestReceived is a free log subscription operation binding the contract event 0x8769c91727bbf50934fb8b59dd2f7df4d307389cd4a7af7f79ec11cd47f7fad7.
//
// Solidity: event RequestReceived(uint256 indexed requestId, string seed, uint256 numOfValues)
func (_RandomRequestManager *RandomRequestManagerFilterer) WatchRequestReceived(opts *bind.WatchOpts, sink chan<- *RandomRequestManagerRequestReceived, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _RandomRequestManager.contract.WatchLogs(opts, "RequestReceived", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomRequestManagerRequestReceived)
				if err := _RandomRequestManager.contract.UnpackLog(event, "RequestReceived", log); err != nil {
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
func (_RandomRequestManager *RandomRequestManagerFilterer) ParseRequestReceived(log types.Log) (*RandomRequestManagerRequestReceived, error) {
	event := new(RandomRequestManagerRequestReceived)
	if err := _RandomRequestManager.contract.UnpackLog(event, "RequestReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
