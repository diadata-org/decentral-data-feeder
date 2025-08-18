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
	ABI: "[{\"type\":\"function\",\"name\":\"counter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fulfillRandomInt\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"randomInts\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"randomness\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"int256\",\"internalType\":\"int256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"request\",\"inputs\":[{\"name\":\"numOfValues\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"request\",\"inputs\":[{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"numOfValues\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"requests\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"numOfValues\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"RequestFulfilled\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"randomInts\",\"type\":\"int256[]\",\"indexed\":false,\"internalType\":\"int256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestReceived\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"numOfValues\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f5ffd5b50610db38061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610060575f3560e01c80633ccb1e381461006457806361bc221a1461008057806381d12c581461009e578063b4844fc7146100cf578063d845a4b3146100ff578063dc6ad19e1461012f575b5f5ffd5b61007e6004803603810190610079919061057a565b61015f565b005b6100886101bf565b60405161009591906105e6565b60405180910390f35b6100b860048036038101906100b391906105ff565b6101c4565b6040516100c692919061069a565b60405180910390f35b6100e960048036038101906100e491906106c8565b61026a565b6040516100f6919061071e565b60405180910390f35b610119600480360381019061011491906105ff565b610295565b60405161012691906105e6565b60405180910390f35b6101496004803603810190610144919061078c565b610369565b60405161015691906105e6565b60405180910390f35b818160025f8681526020019081526020015f20919061017f929190610478565b50827f970a87ecf33833398c3f9a8a833063237777b7afc09272e2143910ce09c07f1b83836040516101b29291906108cf565b60405180910390a2505050565b5f5481565b6001602052805f5260405f205f91509050805f0180546101e39061091e565b80601f016020809104026020016040519081016040528092919081815260200182805461020f9061091e565b801561025a5780601f106102315761010080835404028352916020019161025a565b820191905f5260205f20905b81548152906001019060200180831161023d57829003601f168201915b5050505050908060010154905082565b6002602052815f5260405f208181548110610283575f80fd5b905f5260205f20015f91509150505481565b5f335f5f8154809291906102a89061097b565b919050556040516020016102bd929190610a01565b604051602081830303815290604052805190602001205f1c9050604051806040016040528060405180602001604052805f81525081526020018381525060015f8381526020019081526020015f205f820151815f01908161031e9190610bf5565b5060208201518160010155905050807f8769c91727bbf50934fb8b59dd2f7df4d307389cd4a7af7f79ec11cd47f7fad78360405161035c9190610ce7565b60405180910390a2919050565b5f335f5f81548092919061037c9061097b565b91905055604051602001610391929190610a01565b604051602081830303815290604052805190602001205f1c9050604051806040016040528085858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018381525060015f8381526020019081526020015f205f820151815f0190816104279190610bf5565b5060208201518160010155905050807f8769c91727bbf50934fb8b59dd2f7df4d307389cd4a7af7f79ec11cd47f7fad785858560405161046993929190610d4d565b60405180910390a29392505050565b828054828255905f5260205f209081019282156104b2579160200282015b828111156104b1578235825591602001919060010190610496565b5b5090506104bf91906104c3565b5090565b5b808211156104da575f815f9055506001016104c4565b5090565b5f5ffd5b5f5ffd5b5f819050919050565b6104f8816104e6565b8114610502575f5ffd5b50565b5f81359050610513816104ef565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261053a57610539610519565b5b8235905067ffffffffffffffff8111156105575761055661051d565b5b60208301915083602082028301111561057357610572610521565b5b9250929050565b5f5f5f60408486031215610591576105906104de565b5b5f61059e86828701610505565b935050602084013567ffffffffffffffff8111156105bf576105be6104e2565b5b6105cb86828701610525565b92509250509250925092565b6105e0816104e6565b82525050565b5f6020820190506105f95f8301846105d7565b92915050565b5f60208284031215610614576106136104de565b5b5f61062184828501610505565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61066c8261062a565b6106768185610634565b9350610686818560208601610644565b61068f81610652565b840191505092915050565b5f6040820190508181035f8301526106b28185610662565b90506106c160208301846105d7565b9392505050565b5f5f604083850312156106de576106dd6104de565b5b5f6106eb85828601610505565b92505060206106fc85828601610505565b9150509250929050565b5f819050919050565b61071881610706565b82525050565b5f6020820190506107315f83018461070f565b92915050565b5f5f83601f84011261074c5761074b610519565b5b8235905067ffffffffffffffff8111156107695761076861051d565b5b60208301915083600182028301111561078557610784610521565b5b9250929050565b5f5f5f604084860312156107a3576107a26104de565b5b5f84013567ffffffffffffffff8111156107c0576107bf6104e2565b5b6107cc86828701610737565b935093505060206107df86828701610505565b9150509250925092565b5f82825260208201905092915050565b5f819050919050565b61080b81610706565b82525050565b5f61081c8383610802565b60208301905092915050565b61083181610706565b811461083b575f5ffd5b50565b5f8135905061084c81610828565b92915050565b5f610860602084018461083e565b905092915050565b5f602082019050919050565b5f61087f83856107e9565b935061088a826107f9565b805f5b858110156108c25761089f8284610852565b6108a98882610811565b97506108b483610868565b92505060018101905061088d565b5085925050509392505050565b5f6020820190508181035f8301526108e8818486610874565b90509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061093557607f821691505b602082108103610948576109476108f1565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610985826104e6565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036109b7576109b661094e565b5b600182019050919050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6109eb826109c2565b9050919050565b6109fb816109e1565b82525050565b5f604082019050610a145f8301856109f2565b610a2160208301846105d7565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610ab17fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610a76565b610abb8683610a76565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610af6610af1610aec846104e6565b610ad3565b6104e6565b9050919050565b5f819050919050565b610b0f83610adc565b610b23610b1b82610afd565b848454610a82565b825550505050565b5f5f905090565b610b3a610b2b565b610b45818484610b06565b505050565b5b81811015610b6857610b5d5f82610b32565b600181019050610b4b565b5050565b601f821115610bad57610b7e81610a55565b610b8784610a67565b81016020851015610b96578190505b610baa610ba285610a67565b830182610b4a565b50505b505050565b5f82821c905092915050565b5f610bcd5f1984600802610bb2565b1980831691505092915050565b5f610be58383610bbe565b9150826002028217905092915050565b610bfe8261062a565b67ffffffffffffffff811115610c1757610c16610a28565b5b610c21825461091e565b610c2c828285610b6c565b5f60209050601f831160018114610c5d575f8415610c4b578287015190505b610c558582610bda565b865550610cbc565b601f198416610c6b86610a55565b5f5b82811015610c9257848901518255600182019150602085019450602081019050610c6d565b86831015610caf5784890151610cab601f891682610bbe565b8355505b6001600288020188555050505b505050505050565b50565b5f610cd25f83610634565b9150610cdd82610cc4565b5f82019050919050565b5f6040820190508181035f830152610cfe81610cc7565b9050610d0d60208301846105d7565b92915050565b828183375f83830152505050565b5f610d2c8385610634565b9350610d39838584610d13565b610d4283610652565b840190509392505050565b5f6040820190508181035f830152610d66818587610d21565b9050610d7560208301846105d7565b94935050505056fea264697066735822122094830e683f3f2b47b8e523b96da3966f5ad13498685a3f535450ad4d9347ad9264736f6c634300081d0033",
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

// RandomRequestManagerRequestFulfilledIterator is returned from FilterRequestFulfilled and is used to iterate over the raw logs and unpacked data for RequestFulfilled events raised by the RandomRequestManager contract.
type RandomRequestManagerRequestFulfilledIterator struct {
	Event *RandomRequestManagerRequestFulfilled // Event containing the contract specifics and raw log

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
func (it *RandomRequestManagerRequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomRequestManagerRequestFulfilled)
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
		it.Event = new(RandomRequestManagerRequestFulfilled)
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
func (it *RandomRequestManagerRequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomRequestManagerRequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomRequestManagerRequestFulfilled represents a RequestFulfilled event raised by the RandomRequestManager contract.
type RandomRequestManagerRequestFulfilled struct {
	RequestId  *big.Int
	RandomInts []*big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRequestFulfilled is a free log retrieval operation binding the contract event 0x970a87ecf33833398c3f9a8a833063237777b7afc09272e2143910ce09c07f1b.
//
// Solidity: event RequestFulfilled(uint256 indexed requestId, int256[] randomInts)
func (_RandomRequestManager *RandomRequestManagerFilterer) FilterRequestFulfilled(opts *bind.FilterOpts, requestId []*big.Int) (*RandomRequestManagerRequestFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _RandomRequestManager.contract.FilterLogs(opts, "RequestFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &RandomRequestManagerRequestFulfilledIterator{contract: _RandomRequestManager.contract, event: "RequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchRequestFulfilled is a free log subscription operation binding the contract event 0x970a87ecf33833398c3f9a8a833063237777b7afc09272e2143910ce09c07f1b.
//
// Solidity: event RequestFulfilled(uint256 indexed requestId, int256[] randomInts)
func (_RandomRequestManager *RandomRequestManagerFilterer) WatchRequestFulfilled(opts *bind.WatchOpts, sink chan<- *RandomRequestManagerRequestFulfilled, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _RandomRequestManager.contract.WatchLogs(opts, "RequestFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomRequestManagerRequestFulfilled)
				if err := _RandomRequestManager.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
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

// ParseRequestFulfilled is a log parse operation binding the contract event 0x970a87ecf33833398c3f9a8a833063237777b7afc09272e2143910ce09c07f1b.
//
// Solidity: event RequestFulfilled(uint256 indexed requestId, int256[] randomInts)
func (_RandomRequestManager *RandomRequestManagerFilterer) ParseRequestFulfilled(log types.Log) (*RandomRequestManagerRequestFulfilled, error) {
	event := new(RandomRequestManagerRequestFulfilled)
	if err := _RandomRequestManager.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
