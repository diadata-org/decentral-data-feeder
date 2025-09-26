// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package RandomRequestManagerV2

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

// RandomRequestManagerV2MetaData contains all meta data concerning the RandomRequestManagerV2 contract.
var RandomRequestManagerV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"baseRequest\",\"inputs\":[{\"name\":\"rType\",\"type\":\"uint8\",\"internalType\":\"enumRandomRequestManager.RequestType\"},{\"name\":\"inputParams\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bytesRequests\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"lenBytes\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"numBytes\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"counter\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fulfillRequest\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fulfillmentData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"intArrayRequests\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"bitSize\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"numInts\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"intRangeRequests\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"min\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"max\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"numInts\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"requestTypes\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumRandomRequestManager.RequestType\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"RequestFulfilled\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"requestType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumRandomRequestManager.RequestType\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RequestReceived\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"requestType\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumRandomRequestManager.RequestType\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"params\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f5ffd5b50611e248061001c5f395ff3fe608060405234801561000f575f5ffd5b506004361061007b575f3560e01c806361bc221a1161005957806361bc221a146100ff57806391e82dc91461011d5780639b0fbd3a14610152578063b35aaba5146101865761007b565b80630beff4fd1461007f5780631b08b354146100b35780634bb9816a146100e3575b5f5ffd5b61009960048036038101906100949190610f4a565b6101b6565b6040516100aa959493929190611029565b60405180910390f35b6100cd60048036038101906100c89190611105565b610293565b6040516100da9190611162565b60405180910390f35b6100fd60048036038101906100f8919061117b565b61080f565b005b610107610b60565b6040516101149190611162565b60405180910390f35b61013760048036038101906101329190610f4a565b610b65565b604051610149969594939291906111d8565b60405180910390f35b61016c60048036038101906101679190610f4a565b610c55565b60405161017d959493929190611029565b60405180910390f35b6101a0600480360381019061019b9190610f4a565b610d32565b6040516101ad91906112b1565b60405180910390f35b6002602052805f5260405f205f91509050805f0154908060010180546101db906112f7565b80601f0160208091040260200160405190810160405280929190818152602001828054610207906112f7565b80156102525780601f1061022957610100808354040283529160200191610252565b820191905f5260205f20905b81548152906001019060200180831161023557829003601f168201915b505050505090806002015f9054906101000a900460ff16908060020160019054906101000a900460ff1690806004015f9054906101000a900460ff16905085565b5f61029d33610d4f565b90505f5f8154809291906102b090611354565b91905055508360015f8381526020019081526020015f205f6101000a81548160ff021916908360028111156102e8576102e761123e565b5b02179055505f6002811115610300576102ff61123e565b5b8460028111156103135761031261123e565b5b0361046c575f5f5f858581019061032a91906114ed565b9250925092506040518060c001604052808581526020018481526020018360ff1681526020018260ff1681526020018260ff1667ffffffffffffffff8111156103765761037561139f565b5b6040519080825280602002602001820160405280156103a957816020015b60608152602001906001900390816103945790505b5081526020015f151581525060025f8681526020019081526020015f205f820151815f015560208201518160010190816103e391906116f9565b506040820151816002015f6101000a81548160ff021916908360ff16021790555060608201518160020160016101000a81548160ff021916908360ff1602179055506080820151816003019080519060200190610441929190610d83565b5060a0820151816004015f6101000a81548160ff0219169083151502179055509050505050506107cc565b600160028111156104805761047f61123e565b5b8460028111156104935761049261123e565b5b03610615575f5f5f5f86868101906104ab91906117c8565b93509350935093506040518060e001604052808681526020018581526020018460ff1681526020018360ff1681526020018260ff1681526020018260ff1667ffffffffffffffff8111156105025761050161139f565b5b6040519080825280602002602001820160405280156105305781602001602082028036833780820191505090505b5081526020015f151581525060035f8781526020019081526020015f205f820151815f0155602082015181600101908161056a91906116f9565b506040820151816002015f6101000a81548160ff021916908360ff16021790555060608201518160020160016101000a81548160ff021916908360ff16021790555060808201518160020160026101000a81548160ff021916908360ff16021790555060a08201518160030190805190602001906105e9929190610dda565b5060c0820151816004015f6101000a81548160ff021916908315150217905550905050505050506107cb565b6002808111156106285761062761123e565b5b84600281111561063b5761063a61123e565b5b0361078f575f5f5f858581019061065291906114ed565b9250925092506040518060c001604052808581526020018481526020018360ff1681526020018260ff1681526020018260ff1667ffffffffffffffff81111561069e5761069d61139f565b5b6040519080825280602002602001820160405280156106cc5781602001602082028036833780820191505090505b5081526020015f151581525060045f8681526020019081526020015f205f820151815f0155602082015181600101908161070691906116f9565b506040820151816002015f6101000a81548160ff021916908360ff16021790555060608201518160020160016101000a81548160ff021916908360ff1602179055506080820151816003019080519060200190610764929190610e25565b5060a0820151816004015f6101000a81548160ff0219169083151502179055509050505050506107ca565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107c190611892565b60405180910390fd5b5b5b807fb09e70d737d82238f152497f76a39087113b71590afc3e3be392a1a3d193e4ba8585856040516108009392919061190f565b60405180910390a29392505050565b5f60015f8581526020019081526020015f205f9054906101000a900460ff16905060028160028111156108455761084461123e565b5b60ff161115610889576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108809061199c565b60405180910390fd5b5f600281111561089c5761089b61123e565b5b8160028111156108af576108ae61123e565b5b03610967575f83838101906108c49190611a98565b90505f60025f8781526020019081526020015f209050806004015f9054906101000a900460ff161561092b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161092290611b29565b60405180910390fd5b81816003019080519060200190610943929190610d83565b506001816004015f6101000a81548160ff0219169083151502179055505050610b22565b6001600281111561097b5761097a61123e565b5b81600281111561098e5761098d61123e565b5b03610a46575f83838101906109a39190611c07565b90505f60035f8781526020019081526020015f209050806004015f9054906101000a900460ff1615610a0a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a0190611b29565b60405180910390fd5b81816003019080519060200190610a22929190610dda565b506001816004015f6101000a81548160ff0219169083151502179055505050610b21565b600280811115610a5957610a5861123e565b5b816002811115610a6c57610a6b61123e565b5b03610b20575f8383810190610a819190611d41565b90505f60045f8781526020019081526020015f209050806004015f9054906101000a900460ff1615610ae8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610adf90611b29565b60405180910390fd5b81816003019080519060200190610b00929190610e25565b506001816004015f6101000a81548160ff02191690831515021790555050505b5b5b837f39136673c8494d491b055ca4c9c8ea19dd9ea0be620ff4bedfa57998851b230982604051610b5291906112b1565b60405180910390a250505050565b5f5481565b6003602052805f5260405f205f91509050805f015490806001018054610b8a906112f7565b80601f0160208091040260200160405190810160405280929190818152602001828054610bb6906112f7565b8015610c015780601f10610bd857610100808354040283529160200191610c01565b820191905f5260205f20905b815481529060010190602001808311610be457829003601f168201915b505050505090806002015f9054906101000a900460ff16908060020160019054906101000a900460ff16908060020160029054906101000a900460ff1690806004015f9054906101000a900460ff16905086565b6004602052805f5260405f205f91509050805f015490806001018054610c7a906112f7565b80601f0160208091040260200160405190810160405280929190818152602001828054610ca6906112f7565b8015610cf15780601f10610cc857610100808354040283529160200191610cf1565b820191905f5260205f20905b815481529060010190602001808311610cd457829003601f168201915b505050505090806002015f9054906101000a900460ff16908060020160019054906101000a900460ff1690806004015f9054906101000a900460ff16905085565b6001602052805f5260405f205f915054906101000a900460ff1681565b5f815f54604051602001610d64929190611dc7565b604051602081830303815290604052805190602001205f1c9050919050565b828054828255905f5260205f20908101928215610dc9579160200282015b82811115610dc8578251829081610db891906116f9565b5091602001919060010190610da1565b5b509050610dd69190610e70565b5090565b828054828255905f5260205f20908101928215610e14579160200282015b82811115610e13578251825591602001919060010190610df8565b5b509050610e219190610e93565b5090565b828054828255905f5260205f20908101928215610e5f579160200282015b82811115610e5e578251825591602001919060010190610e43565b5b509050610e6c9190610eae565b5090565b5b80821115610e8f575f8181610e869190610ec9565b50600101610e71565b5090565b5b80821115610eaa575f815f905550600101610e94565b5090565b5b80821115610ec5575f815f905550600101610eaf565b5090565b508054610ed5906112f7565b5f825580601f10610ee65750610f03565b601f0160209004905f5260205f2090810190610f029190610e93565b5b50565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b610f2981610f17565b8114610f33575f5ffd5b50565b5f81359050610f4481610f20565b92915050565b5f60208284031215610f5f57610f5e610f0f565b5b5f610f6c84828501610f36565b91505092915050565b610f7e81610f17565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f610fc682610f84565b610fd08185610f8e565b9350610fe0818560208601610f9e565b610fe981610fac565b840191505092915050565b5f60ff82169050919050565b61100981610ff4565b82525050565b5f8115159050919050565b6110238161100f565b82525050565b5f60a08201905061103c5f830188610f75565b818103602083015261104e8187610fbc565b905061105d6040830186611000565b61106a6060830185611000565b611077608083018461101a565b9695505050505050565b6003811061108d575f5ffd5b50565b5f8135905061109e81611081565b92915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126110c5576110c46110a4565b5b8235905067ffffffffffffffff8111156110e2576110e16110a8565b5b6020830191508360018202830111156110fe576110fd6110ac565b5b9250929050565b5f5f5f6040848603121561111c5761111b610f0f565b5b5f61112986828701611090565b935050602084013567ffffffffffffffff81111561114a57611149610f13565b5b611156868287016110b0565b92509250509250925092565b5f6020820190506111755f830184610f75565b92915050565b5f5f5f6040848603121561119257611191610f0f565b5b5f61119f86828701610f36565b935050602084013567ffffffffffffffff8111156111c0576111bf610f13565b5b6111cc868287016110b0565b92509250509250925092565b5f60c0820190506111eb5f830189610f75565b81810360208301526111fd8188610fbc565b905061120c6040830187611000565b6112196060830186611000565b6112266080830185611000565b61123360a083018461101a565b979650505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b6003811061127c5761127b61123e565b5b50565b5f81905061128c8261126b565b919050565b5f61129b8261127f565b9050919050565b6112ab81611291565b82525050565b5f6020820190506112c45f8301846112a2565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061130e57607f821691505b602082108103611321576113206112ca565b5b50919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61135e82610f17565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036113905761138f611327565b5b600182019050919050565b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6113d582610fac565b810181811067ffffffffffffffff821117156113f4576113f361139f565b5b80604052505050565b5f611406610f06565b905061141282826113cc565b919050565b5f67ffffffffffffffff8211156114315761143061139f565b5b61143a82610fac565b9050602081019050919050565b828183375f83830152505050565b5f61146761146284611417565b6113fd565b9050828152602081018484840111156114835761148261139b565b5b61148e848285611447565b509392505050565b5f82601f8301126114aa576114a96110a4565b5b81356114ba848260208601611455565b91505092915050565b6114cc81610ff4565b81146114d6575f5ffd5b50565b5f813590506114e7816114c3565b92915050565b5f5f5f6060848603121561150457611503610f0f565b5b5f84013567ffffffffffffffff81111561152157611520610f13565b5b61152d86828701611496565b935050602061153e868287016114d9565b925050604061154f868287016114d9565b9150509250925092565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026115b57fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261157a565b6115bf868361157a565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6115fa6115f56115f084610f17565b6115d7565b610f17565b9050919050565b5f819050919050565b611613836115e0565b61162761161f82611601565b848454611586565b825550505050565b5f5f905090565b61163e61162f565b61164981848461160a565b505050565b5b8181101561166c576116615f82611636565b60018101905061164f565b5050565b601f8211156116b15761168281611559565b61168b8461156b565b8101602085101561169a578190505b6116ae6116a68561156b565b83018261164e565b50505b505050565b5f82821c905092915050565b5f6116d15f19846008026116b6565b1980831691505092915050565b5f6116e983836116c2565b9150826002028217905092915050565b61170282610f84565b67ffffffffffffffff81111561171b5761171a61139f565b5b61172582546112f7565b611730828285611670565b5f60209050601f831160018114611761575f841561174f578287015190505b61175985826116de565b8655506117c0565b601f19841661176f86611559565b5f5b8281101561179657848901518255600182019150602085019450602081019050611771565b868310156117b357848901516117af601f8916826116c2565b8355505b6001600288020188555050505b505050505050565b5f5f5f5f608085870312156117e0576117df610f0f565b5b5f85013567ffffffffffffffff8111156117fd576117fc610f13565b5b61180987828801611496565b945050602061181a878288016114d9565b935050604061182b878288016114d9565b925050606061183c878288016114d9565b91505092959194509250565b7f556e737570706f727465642072657175657374207479706500000000000000005f82015250565b5f61187c601883610f8e565b915061188782611848565b602082019050919050565b5f6020820190508181035f8301526118a981611870565b9050919050565b50565b5f6118be5f83610f8e565b91506118c9826118b0565b5f82019050919050565b5f82825260208201905092915050565b5f6118ee83856118d3565b93506118fb838584611447565b61190483610fac565b840190509392505050565b5f6060820190506119225f8301866112a2565b8181036020830152611933816118b3565b905081810360408301526119488184866118e3565b9050949350505050565b7f556e6b6e6f776e207265717565737420747970650000000000000000000000005f82015250565b5f611986601483610f8e565b915061199182611952565b602082019050919050565b5f6020820190508181035f8301526119b38161197a565b9050919050565b5f67ffffffffffffffff8211156119d4576119d361139f565b5b602082029050602081019050919050565b5f6119f76119f2846119ba565b6113fd565b90508083825260208201905060208402830185811115611a1a57611a196110ac565b5b835b81811015611a6157803567ffffffffffffffff811115611a3f57611a3e6110a4565b5b808601611a4c8982611496565b85526020850194505050602081019050611a1c565b5050509392505050565b5f82601f830112611a7f57611a7e6110a4565b5b8135611a8f8482602086016119e5565b91505092915050565b5f60208284031215611aad57611aac610f0f565b5b5f82013567ffffffffffffffff811115611aca57611ac9610f13565b5b611ad684828501611a6b565b91505092915050565b7f416c72656164792066756c66696c6c65640000000000000000000000000000005f82015250565b5f611b13601183610f8e565b9150611b1e82611adf565b602082019050919050565b5f6020820190508181035f830152611b4081611b07565b9050919050565b5f67ffffffffffffffff821115611b6157611b6061139f565b5b602082029050602081019050919050565b5f611b84611b7f84611b47565b6113fd565b90508083825260208201905060208402830185811115611ba757611ba66110ac565b5b835b81811015611bd05780611bbc8882610f36565b845260208401935050602081019050611ba9565b5050509392505050565b5f82601f830112611bee57611bed6110a4565b5b8135611bfe848260208601611b72565b91505092915050565b5f60208284031215611c1c57611c1b610f0f565b5b5f82013567ffffffffffffffff811115611c3957611c38610f13565b5b611c4584828501611bda565b91505092915050565b5f67ffffffffffffffff821115611c6857611c6761139f565b5b602082029050602081019050919050565b5f819050919050565b611c8b81611c79565b8114611c95575f5ffd5b50565b5f81359050611ca681611c82565b92915050565b5f611cbe611cb984611c4e565b6113fd565b90508083825260208201905060208402830185811115611ce157611ce06110ac565b5b835b81811015611d0a5780611cf68882611c98565b845260208401935050602081019050611ce3565b5050509392505050565b5f82601f830112611d2857611d276110a4565b5b8135611d38848260208601611cac565b91505092915050565b5f60208284031215611d5657611d55610f0f565b5b5f82013567ffffffffffffffff811115611d7357611d72610f13565b5b611d7f84828501611d14565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611db182611d88565b9050919050565b611dc181611da7565b82525050565b5f604082019050611dda5f830185611db8565b611de76020830184610f75565b939250505056fea26469706673582212206542e03840195a436c4d5736d35b1fcbef1207eb629500d13d2b7937837af07964736f6c634300081e0033",
}

// RandomRequestManagerV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use RandomRequestManagerV2MetaData.ABI instead.
var RandomRequestManagerV2ABI = RandomRequestManagerV2MetaData.ABI

// RandomRequestManagerV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RandomRequestManagerV2MetaData.Bin instead.
var RandomRequestManagerV2Bin = RandomRequestManagerV2MetaData.Bin

// DeployRandomRequestManagerV2 deploys a new Ethereum contract, binding an instance of RandomRequestManagerV2 to it.
func DeployRandomRequestManagerV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RandomRequestManagerV2, error) {
	parsed, err := RandomRequestManagerV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RandomRequestManagerV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RandomRequestManagerV2{RandomRequestManagerV2Caller: RandomRequestManagerV2Caller{contract: contract}, RandomRequestManagerV2Transactor: RandomRequestManagerV2Transactor{contract: contract}, RandomRequestManagerV2Filterer: RandomRequestManagerV2Filterer{contract: contract}}, nil
}

// RandomRequestManagerV2 is an auto generated Go binding around an Ethereum contract.
type RandomRequestManagerV2 struct {
	RandomRequestManagerV2Caller     // Read-only binding to the contract
	RandomRequestManagerV2Transactor // Write-only binding to the contract
	RandomRequestManagerV2Filterer   // Log filterer for contract events
}

// RandomRequestManagerV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type RandomRequestManagerV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomRequestManagerV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type RandomRequestManagerV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomRequestManagerV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RandomRequestManagerV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RandomRequestManagerV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RandomRequestManagerV2Session struct {
	Contract     *RandomRequestManagerV2 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RandomRequestManagerV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RandomRequestManagerV2CallerSession struct {
	Contract *RandomRequestManagerV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// RandomRequestManagerV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RandomRequestManagerV2TransactorSession struct {
	Contract     *RandomRequestManagerV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// RandomRequestManagerV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type RandomRequestManagerV2Raw struct {
	Contract *RandomRequestManagerV2 // Generic contract binding to access the raw methods on
}

// RandomRequestManagerV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RandomRequestManagerV2CallerRaw struct {
	Contract *RandomRequestManagerV2Caller // Generic read-only contract binding to access the raw methods on
}

// RandomRequestManagerV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RandomRequestManagerV2TransactorRaw struct {
	Contract *RandomRequestManagerV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRandomRequestManagerV2 creates a new instance of RandomRequestManagerV2, bound to a specific deployed contract.
func NewRandomRequestManagerV2(address common.Address, backend bind.ContractBackend) (*RandomRequestManagerV2, error) {
	contract, err := bindRandomRequestManagerV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RandomRequestManagerV2{RandomRequestManagerV2Caller: RandomRequestManagerV2Caller{contract: contract}, RandomRequestManagerV2Transactor: RandomRequestManagerV2Transactor{contract: contract}, RandomRequestManagerV2Filterer: RandomRequestManagerV2Filterer{contract: contract}}, nil
}

// NewRandomRequestManagerV2Caller creates a new read-only instance of RandomRequestManagerV2, bound to a specific deployed contract.
func NewRandomRequestManagerV2Caller(address common.Address, caller bind.ContractCaller) (*RandomRequestManagerV2Caller, error) {
	contract, err := bindRandomRequestManagerV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RandomRequestManagerV2Caller{contract: contract}, nil
}

// NewRandomRequestManagerV2Transactor creates a new write-only instance of RandomRequestManagerV2, bound to a specific deployed contract.
func NewRandomRequestManagerV2Transactor(address common.Address, transactor bind.ContractTransactor) (*RandomRequestManagerV2Transactor, error) {
	contract, err := bindRandomRequestManagerV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RandomRequestManagerV2Transactor{contract: contract}, nil
}

// NewRandomRequestManagerV2Filterer creates a new log filterer instance of RandomRequestManagerV2, bound to a specific deployed contract.
func NewRandomRequestManagerV2Filterer(address common.Address, filterer bind.ContractFilterer) (*RandomRequestManagerV2Filterer, error) {
	contract, err := bindRandomRequestManagerV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RandomRequestManagerV2Filterer{contract: contract}, nil
}

// bindRandomRequestManagerV2 binds a generic wrapper to an already deployed contract.
func bindRandomRequestManagerV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RandomRequestManagerV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomRequestManagerV2 *RandomRequestManagerV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomRequestManagerV2.Contract.RandomRequestManagerV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomRequestManagerV2 *RandomRequestManagerV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomRequestManagerV2.Contract.RandomRequestManagerV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomRequestManagerV2 *RandomRequestManagerV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomRequestManagerV2.Contract.RandomRequestManagerV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RandomRequestManagerV2 *RandomRequestManagerV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RandomRequestManagerV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RandomRequestManagerV2 *RandomRequestManagerV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RandomRequestManagerV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RandomRequestManagerV2 *RandomRequestManagerV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RandomRequestManagerV2.Contract.contract.Transact(opts, method, params...)
}

// BytesRequests is a free data retrieval call binding the contract method 0x0beff4fd.
//
// Solidity: function bytesRequests(uint256 ) view returns(uint256 requestId, string seed, uint8 lenBytes, uint8 numBytes, bool fulfilled)
func (_RandomRequestManagerV2 *RandomRequestManagerV2Caller) BytesRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RequestId *big.Int
	Seed      string
	LenBytes  uint8
	NumBytes  uint8
	Fulfilled bool
}, error) {
	var out []interface{}
	err := _RandomRequestManagerV2.contract.Call(opts, &out, "bytesRequests", arg0)

	outstruct := new(struct {
		RequestId *big.Int
		Seed      string
		LenBytes  uint8
		NumBytes  uint8
		Fulfilled bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Seed = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.LenBytes = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.NumBytes = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Fulfilled = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// BytesRequests is a free data retrieval call binding the contract method 0x0beff4fd.
//
// Solidity: function bytesRequests(uint256 ) view returns(uint256 requestId, string seed, uint8 lenBytes, uint8 numBytes, bool fulfilled)
func (_RandomRequestManagerV2 *RandomRequestManagerV2Session) BytesRequests(arg0 *big.Int) (struct {
	RequestId *big.Int
	Seed      string
	LenBytes  uint8
	NumBytes  uint8
	Fulfilled bool
}, error) {
	return _RandomRequestManagerV2.Contract.BytesRequests(&_RandomRequestManagerV2.CallOpts, arg0)
}

// BytesRequests is a free data retrieval call binding the contract method 0x0beff4fd.
//
// Solidity: function bytesRequests(uint256 ) view returns(uint256 requestId, string seed, uint8 lenBytes, uint8 numBytes, bool fulfilled)
func (_RandomRequestManagerV2 *RandomRequestManagerV2CallerSession) BytesRequests(arg0 *big.Int) (struct {
	RequestId *big.Int
	Seed      string
	LenBytes  uint8
	NumBytes  uint8
	Fulfilled bool
}, error) {
	return _RandomRequestManagerV2.Contract.BytesRequests(&_RandomRequestManagerV2.CallOpts, arg0)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_RandomRequestManagerV2 *RandomRequestManagerV2Caller) Counter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RandomRequestManagerV2.contract.Call(opts, &out, "counter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_RandomRequestManagerV2 *RandomRequestManagerV2Session) Counter() (*big.Int, error) {
	return _RandomRequestManagerV2.Contract.Counter(&_RandomRequestManagerV2.CallOpts)
}

// Counter is a free data retrieval call binding the contract method 0x61bc221a.
//
// Solidity: function counter() view returns(uint256)
func (_RandomRequestManagerV2 *RandomRequestManagerV2CallerSession) Counter() (*big.Int, error) {
	return _RandomRequestManagerV2.Contract.Counter(&_RandomRequestManagerV2.CallOpts)
}

// IntArrayRequests is a free data retrieval call binding the contract method 0x9b0fbd3a.
//
// Solidity: function intArrayRequests(uint256 ) view returns(uint256 requestId, string seed, uint8 bitSize, uint8 numInts, bool fulfilled)
func (_RandomRequestManagerV2 *RandomRequestManagerV2Caller) IntArrayRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RequestId *big.Int
	Seed      string
	BitSize   uint8
	NumInts   uint8
	Fulfilled bool
}, error) {
	var out []interface{}
	err := _RandomRequestManagerV2.contract.Call(opts, &out, "intArrayRequests", arg0)

	outstruct := new(struct {
		RequestId *big.Int
		Seed      string
		BitSize   uint8
		NumInts   uint8
		Fulfilled bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Seed = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.BitSize = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.NumInts = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Fulfilled = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// IntArrayRequests is a free data retrieval call binding the contract method 0x9b0fbd3a.
//
// Solidity: function intArrayRequests(uint256 ) view returns(uint256 requestId, string seed, uint8 bitSize, uint8 numInts, bool fulfilled)
func (_RandomRequestManagerV2 *RandomRequestManagerV2Session) IntArrayRequests(arg0 *big.Int) (struct {
	RequestId *big.Int
	Seed      string
	BitSize   uint8
	NumInts   uint8
	Fulfilled bool
}, error) {
	return _RandomRequestManagerV2.Contract.IntArrayRequests(&_RandomRequestManagerV2.CallOpts, arg0)
}

// IntArrayRequests is a free data retrieval call binding the contract method 0x9b0fbd3a.
//
// Solidity: function intArrayRequests(uint256 ) view returns(uint256 requestId, string seed, uint8 bitSize, uint8 numInts, bool fulfilled)
func (_RandomRequestManagerV2 *RandomRequestManagerV2CallerSession) IntArrayRequests(arg0 *big.Int) (struct {
	RequestId *big.Int
	Seed      string
	BitSize   uint8
	NumInts   uint8
	Fulfilled bool
}, error) {
	return _RandomRequestManagerV2.Contract.IntArrayRequests(&_RandomRequestManagerV2.CallOpts, arg0)
}

// IntRangeRequests is a free data retrieval call binding the contract method 0x91e82dc9.
//
// Solidity: function intRangeRequests(uint256 ) view returns(uint256 requestId, string seed, uint8 min, uint8 max, uint8 numInts, bool fulfilled)
func (_RandomRequestManagerV2 *RandomRequestManagerV2Caller) IntRangeRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	RequestId *big.Int
	Seed      string
	Min       uint8
	Max       uint8
	NumInts   uint8
	Fulfilled bool
}, error) {
	var out []interface{}
	err := _RandomRequestManagerV2.contract.Call(opts, &out, "intRangeRequests", arg0)

	outstruct := new(struct {
		RequestId *big.Int
		Seed      string
		Min       uint8
		Max       uint8
		NumInts   uint8
		Fulfilled bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Seed = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Min = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Max = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.NumInts = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.Fulfilled = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// IntRangeRequests is a free data retrieval call binding the contract method 0x91e82dc9.
//
// Solidity: function intRangeRequests(uint256 ) view returns(uint256 requestId, string seed, uint8 min, uint8 max, uint8 numInts, bool fulfilled)
func (_RandomRequestManagerV2 *RandomRequestManagerV2Session) IntRangeRequests(arg0 *big.Int) (struct {
	RequestId *big.Int
	Seed      string
	Min       uint8
	Max       uint8
	NumInts   uint8
	Fulfilled bool
}, error) {
	return _RandomRequestManagerV2.Contract.IntRangeRequests(&_RandomRequestManagerV2.CallOpts, arg0)
}

// IntRangeRequests is a free data retrieval call binding the contract method 0x91e82dc9.
//
// Solidity: function intRangeRequests(uint256 ) view returns(uint256 requestId, string seed, uint8 min, uint8 max, uint8 numInts, bool fulfilled)
func (_RandomRequestManagerV2 *RandomRequestManagerV2CallerSession) IntRangeRequests(arg0 *big.Int) (struct {
	RequestId *big.Int
	Seed      string
	Min       uint8
	Max       uint8
	NumInts   uint8
	Fulfilled bool
}, error) {
	return _RandomRequestManagerV2.Contract.IntRangeRequests(&_RandomRequestManagerV2.CallOpts, arg0)
}

// RequestTypes is a free data retrieval call binding the contract method 0xb35aaba5.
//
// Solidity: function requestTypes(uint256 ) view returns(uint8)
func (_RandomRequestManagerV2 *RandomRequestManagerV2Caller) RequestTypes(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _RandomRequestManagerV2.contract.Call(opts, &out, "requestTypes", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// RequestTypes is a free data retrieval call binding the contract method 0xb35aaba5.
//
// Solidity: function requestTypes(uint256 ) view returns(uint8)
func (_RandomRequestManagerV2 *RandomRequestManagerV2Session) RequestTypes(arg0 *big.Int) (uint8, error) {
	return _RandomRequestManagerV2.Contract.RequestTypes(&_RandomRequestManagerV2.CallOpts, arg0)
}

// RequestTypes is a free data retrieval call binding the contract method 0xb35aaba5.
//
// Solidity: function requestTypes(uint256 ) view returns(uint8)
func (_RandomRequestManagerV2 *RandomRequestManagerV2CallerSession) RequestTypes(arg0 *big.Int) (uint8, error) {
	return _RandomRequestManagerV2.Contract.RequestTypes(&_RandomRequestManagerV2.CallOpts, arg0)
}

// BaseRequest is a paid mutator transaction binding the contract method 0x1b08b354.
//
// Solidity: function baseRequest(uint8 rType, bytes inputParams) returns(uint256 requestId)
func (_RandomRequestManagerV2 *RandomRequestManagerV2Transactor) BaseRequest(opts *bind.TransactOpts, rType uint8, inputParams []byte) (*types.Transaction, error) {
	return _RandomRequestManagerV2.contract.Transact(opts, "baseRequest", rType, inputParams)
}

// BaseRequest is a paid mutator transaction binding the contract method 0x1b08b354.
//
// Solidity: function baseRequest(uint8 rType, bytes inputParams) returns(uint256 requestId)
func (_RandomRequestManagerV2 *RandomRequestManagerV2Session) BaseRequest(rType uint8, inputParams []byte) (*types.Transaction, error) {
	return _RandomRequestManagerV2.Contract.BaseRequest(&_RandomRequestManagerV2.TransactOpts, rType, inputParams)
}

// BaseRequest is a paid mutator transaction binding the contract method 0x1b08b354.
//
// Solidity: function baseRequest(uint8 rType, bytes inputParams) returns(uint256 requestId)
func (_RandomRequestManagerV2 *RandomRequestManagerV2TransactorSession) BaseRequest(rType uint8, inputParams []byte) (*types.Transaction, error) {
	return _RandomRequestManagerV2.Contract.BaseRequest(&_RandomRequestManagerV2.TransactOpts, rType, inputParams)
}

// FulfillRequest is a paid mutator transaction binding the contract method 0x4bb9816a.
//
// Solidity: function fulfillRequest(uint256 requestId, bytes fulfillmentData) returns()
func (_RandomRequestManagerV2 *RandomRequestManagerV2Transactor) FulfillRequest(opts *bind.TransactOpts, requestId *big.Int, fulfillmentData []byte) (*types.Transaction, error) {
	return _RandomRequestManagerV2.contract.Transact(opts, "fulfillRequest", requestId, fulfillmentData)
}

// FulfillRequest is a paid mutator transaction binding the contract method 0x4bb9816a.
//
// Solidity: function fulfillRequest(uint256 requestId, bytes fulfillmentData) returns()
func (_RandomRequestManagerV2 *RandomRequestManagerV2Session) FulfillRequest(requestId *big.Int, fulfillmentData []byte) (*types.Transaction, error) {
	return _RandomRequestManagerV2.Contract.FulfillRequest(&_RandomRequestManagerV2.TransactOpts, requestId, fulfillmentData)
}

// FulfillRequest is a paid mutator transaction binding the contract method 0x4bb9816a.
//
// Solidity: function fulfillRequest(uint256 requestId, bytes fulfillmentData) returns()
func (_RandomRequestManagerV2 *RandomRequestManagerV2TransactorSession) FulfillRequest(requestId *big.Int, fulfillmentData []byte) (*types.Transaction, error) {
	return _RandomRequestManagerV2.Contract.FulfillRequest(&_RandomRequestManagerV2.TransactOpts, requestId, fulfillmentData)
}

// RandomRequestManagerV2RequestFulfilledIterator is returned from FilterRequestFulfilled and is used to iterate over the raw logs and unpacked data for RequestFulfilled events raised by the RandomRequestManagerV2 contract.
type RandomRequestManagerV2RequestFulfilledIterator struct {
	Event *RandomRequestManagerV2RequestFulfilled // Event containing the contract specifics and raw log

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
func (it *RandomRequestManagerV2RequestFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomRequestManagerV2RequestFulfilled)
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
		it.Event = new(RandomRequestManagerV2RequestFulfilled)
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
func (it *RandomRequestManagerV2RequestFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomRequestManagerV2RequestFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomRequestManagerV2RequestFulfilled represents a RequestFulfilled event raised by the RandomRequestManagerV2 contract.
type RandomRequestManagerV2RequestFulfilled struct {
	RequestId   *big.Int
	RequestType uint8
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRequestFulfilled is a free log retrieval operation binding the contract event 0x39136673c8494d491b055ca4c9c8ea19dd9ea0be620ff4bedfa57998851b2309.
//
// Solidity: event RequestFulfilled(uint256 indexed requestId, uint8 requestType)
func (_RandomRequestManagerV2 *RandomRequestManagerV2Filterer) FilterRequestFulfilled(opts *bind.FilterOpts, requestId []*big.Int) (*RandomRequestManagerV2RequestFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _RandomRequestManagerV2.contract.FilterLogs(opts, "RequestFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &RandomRequestManagerV2RequestFulfilledIterator{contract: _RandomRequestManagerV2.contract, event: "RequestFulfilled", logs: logs, sub: sub}, nil
}

// WatchRequestFulfilled is a free log subscription operation binding the contract event 0x39136673c8494d491b055ca4c9c8ea19dd9ea0be620ff4bedfa57998851b2309.
//
// Solidity: event RequestFulfilled(uint256 indexed requestId, uint8 requestType)
func (_RandomRequestManagerV2 *RandomRequestManagerV2Filterer) WatchRequestFulfilled(opts *bind.WatchOpts, sink chan<- *RandomRequestManagerV2RequestFulfilled, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _RandomRequestManagerV2.contract.WatchLogs(opts, "RequestFulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomRequestManagerV2RequestFulfilled)
				if err := _RandomRequestManagerV2.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
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

// ParseRequestFulfilled is a log parse operation binding the contract event 0x39136673c8494d491b055ca4c9c8ea19dd9ea0be620ff4bedfa57998851b2309.
//
// Solidity: event RequestFulfilled(uint256 indexed requestId, uint8 requestType)
func (_RandomRequestManagerV2 *RandomRequestManagerV2Filterer) ParseRequestFulfilled(log types.Log) (*RandomRequestManagerV2RequestFulfilled, error) {
	event := new(RandomRequestManagerV2RequestFulfilled)
	if err := _RandomRequestManagerV2.contract.UnpackLog(event, "RequestFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RandomRequestManagerV2RequestReceivedIterator is returned from FilterRequestReceived and is used to iterate over the raw logs and unpacked data for RequestReceived events raised by the RandomRequestManagerV2 contract.
type RandomRequestManagerV2RequestReceivedIterator struct {
	Event *RandomRequestManagerV2RequestReceived // Event containing the contract specifics and raw log

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
func (it *RandomRequestManagerV2RequestReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RandomRequestManagerV2RequestReceived)
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
		it.Event = new(RandomRequestManagerV2RequestReceived)
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
func (it *RandomRequestManagerV2RequestReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RandomRequestManagerV2RequestReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RandomRequestManagerV2RequestReceived represents a RequestReceived event raised by the RandomRequestManagerV2 contract.
type RandomRequestManagerV2RequestReceived struct {
	RequestId   *big.Int
	RequestType uint8
	Seed        string
	Params      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRequestReceived is a free log retrieval operation binding the contract event 0xb09e70d737d82238f152497f76a39087113b71590afc3e3be392a1a3d193e4ba.
//
// Solidity: event RequestReceived(uint256 indexed requestId, uint8 requestType, string seed, bytes params)
func (_RandomRequestManagerV2 *RandomRequestManagerV2Filterer) FilterRequestReceived(opts *bind.FilterOpts, requestId []*big.Int) (*RandomRequestManagerV2RequestReceivedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _RandomRequestManagerV2.contract.FilterLogs(opts, "RequestReceived", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &RandomRequestManagerV2RequestReceivedIterator{contract: _RandomRequestManagerV2.contract, event: "RequestReceived", logs: logs, sub: sub}, nil
}

// WatchRequestReceived is a free log subscription operation binding the contract event 0xb09e70d737d82238f152497f76a39087113b71590afc3e3be392a1a3d193e4ba.
//
// Solidity: event RequestReceived(uint256 indexed requestId, uint8 requestType, string seed, bytes params)
func (_RandomRequestManagerV2 *RandomRequestManagerV2Filterer) WatchRequestReceived(opts *bind.WatchOpts, sink chan<- *RandomRequestManagerV2RequestReceived, requestId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _RandomRequestManagerV2.contract.WatchLogs(opts, "RequestReceived", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RandomRequestManagerV2RequestReceived)
				if err := _RandomRequestManagerV2.contract.UnpackLog(event, "RequestReceived", log); err != nil {
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

// ParseRequestReceived is a log parse operation binding the contract event 0xb09e70d737d82238f152497f76a39087113b71590afc3e3be392a1a3d193e4ba.
//
// Solidity: event RequestReceived(uint256 indexed requestId, uint8 requestType, string seed, bytes params)
func (_RandomRequestManagerV2 *RandomRequestManagerV2Filterer) ParseRequestReceived(log types.Log) (*RandomRequestManagerV2RequestReceived, error) {
	event := new(RandomRequestManagerV2RequestReceived)
	if err := _RandomRequestManagerV2.contract.UnpackLog(event, "RequestReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
