// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DIAOracleRandomness

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

// DIAOracleRandomnessMetaData contains all meta data concerning the DIAOracleRandomness contract.
var DIAOracleRandomnessMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBytes\",\"inputs\":[{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"randomness\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIntArray\",\"inputs\":[{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"randomInts\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIntRange\",\"inputs\":[{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"randomInts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBytes\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"randomness\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIntArray\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"randomInts\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIntRange\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"randomInts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOracleUpdaterAddress\",\"inputs\":[{\"name\":\"newOracleUpdaterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BytesSet\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"round\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IntArraySet\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"round\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IntRangeSet\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"round\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdaterAddressChange\",\"inputs\":[{\"name\":\"newUpdater\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f5ffd5b503360035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506122218061005c5f395ff3fe608060405234801561000f575f5ffd5b506004361061007b575f3560e01c8063537592651161005957806353759265146100ea5780636aa45efc1461011d578063ad29e1da14610139578063d9082a9b146101555761007b565b806307c827591461007f578063154122d11461009b57806345459cd0146100b7575b5f5ffd5b6100996004803603810190610094919061130a565b610188565b005b6100b560048036038101906100b09190611455565b61039e565b005b6100d160048036038101906100cc919061154b565b6105b4565b6040516100e194939291906116a6565b60405180910390f35b61010460048036038101906100ff919061154b565b610841565b6040516101149493929190611808565b60405180910390f35b610137600480360381019061013291906118c1565b610b4a565b005b610153600480360381019061014e9190611941565b610c1c565b005b61016f600480360381019061016a919061154b565b610dfd565b60405161017f9493929190611af7565b60405180910390f35b6040518060a001604052808a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018888808060200260200160405190810160405280939291908181526020018383602002808284375f81840152601f19601f82011690508083019250505050505050815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200183838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020016001151581525060025f8760070b81526020019081526020015f205f820151815f0190816102e29190611d80565b5060208201518160010190805190602001906102ff92919061108a565b5060408201518160020190816103159190611d80565b50606082015181600301908161032b9190611d80565b506080820151816004015f6101000a81548160ff0219169083151502179055509050508460070b7fdb8e3995a930d402dc2af8569cbaaa06b6be56a42d4014fb12c83c33232e9e648a8a8787878760405161038b96959493929190611e89565b60405180910390a2505050505050505050565b6040518060a001604052808a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018888808060200260200160405190810160405280939291908181526020018383602002808284375f81840152601f19601f82011690508083019250505050505050815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200183838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020016001151581525060015f8760070b81526020019081526020015f205f820151815f0190816104f89190611d80565b5060208201518160010190805190602001906105159291906110d5565b50604082015181600201908161052b9190611d80565b5060608201518160030190816105419190611d80565b506080820151816004015f6101000a81548160ff0219169083151502179055509050508460070b7f0143668be676fa7bd52d90cc60925a18d048062ff315d7542e640da4f2e355478a8a878787876040516105a196959493929190611e89565b60405180910390a2505050505050505050565b60608060608060025f8660070b81526020019081526020015f206004015f9054906101000a900460ff1661061d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061490611f23565b60405180910390fd5b5f60025f8760070b81526020019081526020015f209050805f0181600101826002018360030183805461064f90611bb0565b80601f016020809104026020016040519081016040528092919081815260200182805461067b90611bb0565b80156106c65780601f1061069d576101008083540402835291602001916106c6565b820191905f5260205f20905b8154815290600101906020018083116106a957829003601f168201915b505050505093508280548060200260200160405190810160405280929190818152602001828054801561071657602002820191905f5260205f20905b815481526020019060010190808311610702575b5050505050925081805461072990611bb0565b80601f016020809104026020016040519081016040528092919081815260200182805461075590611bb0565b80156107a05780601f10610777576101008083540402835291602001916107a0565b820191905f5260205f20905b81548152906001019060200180831161078357829003601f168201915b505050505091508080546107b390611bb0565b80601f01602080910402602001604051908101604052809291908181526020018280546107df90611bb0565b801561082a5780601f106108015761010080835404028352916020019161082a565b820191905f5260205f20905b81548152906001019060200180831161080d57829003601f168201915b505050505090509450945094509450509193509193565b6060806060805f5f8660070b81526020019081526020015f206004015f9054906101000a900460ff166108a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108a090611f8b565b60405180910390fd5b5f5f5f8760070b81526020019081526020015f209050805f018160010182600201836003018380546108da90611bb0565b80601f016020809104026020016040519081016040528092919081815260200182805461090690611bb0565b80156109515780601f1061092857610100808354040283529160200191610951565b820191905f5260205f20905b81548152906001019060200180831161093457829003601f168201915b5050505050935082805480602002602001604051908101604052809291908181526020015f905b82821015610a20578382905f5260205f2001805461099590611bb0565b80601f01602080910402602001604051908101604052809291908181526020018280546109c190611bb0565b8015610a0c5780601f106109e357610100808354040283529160200191610a0c565b820191905f5260205f20905b8154815290600101906020018083116109ef57829003601f168201915b505050505081526020019060010190610978565b505050509250818054610a3290611bb0565b80601f0160208091040260200160405190810160405280929190818152602001828054610a5e90611bb0565b8015610aa95780601f10610a8057610100808354040283529160200191610aa9565b820191905f5260205f20905b815481529060010190602001808311610a8c57829003601f168201915b50505050509150808054610abc90611bb0565b80601f0160208091040260200160405190810160405280929190818152602001828054610ae890611bb0565b8015610b335780601f10610b0a57610100808354040283529160200191610b33565b820191905f5260205f20905b815481529060010190602001808311610b1657829003601f168201915b505050505090509450945094509450509193509193565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610ba2575f5ffd5b8060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f81604051610c119190611fb8565b60405180910390a150565b6040518060a001604052808a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050508152602001888890610c7d919061216f565b815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200183838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050508152602001600115158152505f5f8760070b81526020019081526020015f205f820151815f019081610d419190611d80565b506020820151816001019080519060200190610d5e929190611120565b506040820151816002019081610d749190611d80565b506060820151816003019081610d8a9190611d80565b506080820151816004015f6101000a81548160ff0219169083151502179055509050508460070b7f257f2a31afdd06af9923520043e89ea1d5bf9cf877a0e5a97cbd81332e6c1d0b8a8a87878787604051610dea96959493929190611e89565b60405180910390a2505050505050505050565b60608060608060015f8660070b81526020019081526020015f206004015f9054906101000a900460ff16610e66576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e5d906121cd565b60405180910390fd5b5f60015f8760070b81526020019081526020015f209050805f01816001018260020183600301838054610e9890611bb0565b80601f0160208091040260200160405190810160405280929190818152602001828054610ec490611bb0565b8015610f0f5780601f10610ee657610100808354040283529160200191610f0f565b820191905f5260205f20905b815481529060010190602001808311610ef257829003601f168201915b5050505050935082805480602002602001604051908101604052809291908181526020018280548015610f5f57602002820191905f5260205f20905b815481526020019060010190808311610f4b575b50505050509250818054610f7290611bb0565b80601f0160208091040260200160405190810160405280929190818152602001828054610f9e90611bb0565b8015610fe95780601f10610fc057610100808354040283529160200191610fe9565b820191905f5260205f20905b815481529060010190602001808311610fcc57829003601f168201915b50505050509150808054610ffc90611bb0565b80601f016020809104026020016040519081016040528092919081815260200182805461102890611bb0565b80156110735780601f1061104a57610100808354040283529160200191611073565b820191905f5260205f20905b81548152906001019060200180831161105657829003601f168201915b505050505090509450945094509450509193509193565b828054828255905f5260205f209081019282156110c4579160200282015b828111156110c35782518255916020019190600101906110a8565b5b5090506110d19190611177565b5090565b828054828255905f5260205f2090810192821561110f579160200282015b8281111561110e5782518255916020019190600101906110f3565b5b50905061111c9190611192565b5090565b828054828255905f5260205f20908101928215611166579160200282015b828111156111655782518290816111559190611d80565b509160200191906001019061113e565b5b50905061117391906111ad565b5090565b5b8082111561118e575f815f905550600101611178565b5090565b5b808211156111a9575f815f905550600101611193565b5090565b5b808211156111cc575f81816111c391906111d0565b506001016111ae565b5090565b5080546111dc90611bb0565b5f825580601f106111ed575061120a565b601f0160209004905f5260205f20908101906112099190611192565b5b50565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f84011261123f5761123e61121e565b5b8235905067ffffffffffffffff81111561125c5761125b611222565b5b60208301915083600182028301111561127857611277611226565b5b9250929050565b5f5f83601f8401126112945761129361121e565b5b8235905067ffffffffffffffff8111156112b1576112b0611222565b5b6020830191508360208202830111156112cd576112cc611226565b5b9250929050565b5f8160070b9050919050565b6112e9816112d4565b81146112f3575f5ffd5b50565b5f81359050611304816112e0565b92915050565b5f5f5f5f5f5f5f5f5f60a08a8c03121561132757611326611216565b5b5f8a013567ffffffffffffffff8111156113445761134361121a565b5b6113508c828d0161122a565b995099505060208a013567ffffffffffffffff8111156113735761137261121a565b5b61137f8c828d0161127f565b975097505060406113928c828d016112f6565b95505060608a013567ffffffffffffffff8111156113b3576113b261121a565b5b6113bf8c828d0161122a565b945094505060808a013567ffffffffffffffff8111156113e2576113e161121a565b5b6113ee8c828d0161122a565b92509250509295985092959850929598565b5f5f83601f8401126114155761141461121e565b5b8235905067ffffffffffffffff81111561143257611431611222565b5b60208301915083602082028301111561144e5761144d611226565b5b9250929050565b5f5f5f5f5f5f5f5f5f60a08a8c03121561147257611471611216565b5b5f8a013567ffffffffffffffff81111561148f5761148e61121a565b5b61149b8c828d0161122a565b995099505060208a013567ffffffffffffffff8111156114be576114bd61121a565b5b6114ca8c828d01611400565b975097505060406114dd8c828d016112f6565b95505060608a013567ffffffffffffffff8111156114fe576114fd61121a565b5b61150a8c828d0161122a565b945094505060808a013567ffffffffffffffff81111561152d5761152c61121a565b5b6115398c828d0161122a565b92509250509295985092959850929598565b5f602082840312156115605761155f611216565b5b5f61156d848285016112f6565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6115b882611576565b6115c28185611580565b93506115d2818560208601611590565b6115db8161159e565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b6116218161160f565b82525050565b5f6116328383611618565b60208301905092915050565b5f602082019050919050565b5f611654826115e6565b61165e81856115f0565b935061166983611600565b805f5b838110156116995781516116808882611627565b975061168b8361163e565b92505060018101905061166c565b5085935050505092915050565b5f6080820190508181035f8301526116be81876115ae565b905081810360208301526116d2818661164a565b905081810360408301526116e681856115ae565b905081810360608301526116fa81846115ae565b905095945050505050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f82825260208201905092915050565b5f61174882611576565b611752818561172e565b9350611762818560208601611590565b61176b8161159e565b840191505092915050565b5f611781838361173e565b905092915050565b5f602082019050919050565b5f61179f82611705565b6117a9818561170f565b9350836020820285016117bb8561171f565b805f5b858110156117f657848403895281516117d78582611776565b94506117e283611789565b925060208a019950506001810190506117be565b50829750879550505050505092915050565b5f6080820190508181035f83015261182081876115ae565b905081810360208301526118348186611795565b9050818103604083015261184881856115ae565b9050818103606083015261185c81846115ae565b905095945050505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61189082611867565b9050919050565b6118a081611886565b81146118aa575f5ffd5b50565b5f813590506118bb81611897565b92915050565b5f602082840312156118d6576118d5611216565b5b5f6118e3848285016118ad565b91505092915050565b5f5f83601f8401126119015761190061121e565b5b8235905067ffffffffffffffff81111561191e5761191d611222565b5b60208301915083602082028301111561193a57611939611226565b5b9250929050565b5f5f5f5f5f5f5f5f5f60a08a8c03121561195e5761195d611216565b5b5f8a013567ffffffffffffffff81111561197b5761197a61121a565b5b6119878c828d0161122a565b995099505060208a013567ffffffffffffffff8111156119aa576119a961121a565b5b6119b68c828d016118ec565b975097505060406119c98c828d016112f6565b95505060608a013567ffffffffffffffff8111156119ea576119e961121a565b5b6119f68c828d0161122a565b945094505060808a013567ffffffffffffffff811115611a1957611a1861121a565b5b611a258c828d0161122a565b92509250509295985092959850929598565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b611a7281611a60565b82525050565b5f611a838383611a69565b60208301905092915050565b5f602082019050919050565b5f611aa582611a37565b611aaf8185611a41565b9350611aba83611a51565b805f5b83811015611aea578151611ad18882611a78565b9750611adc83611a8f565b925050600181019050611abd565b5085935050505092915050565b5f6080820190508181035f830152611b0f81876115ae565b90508181036020830152611b238186611a9b565b90508181036040830152611b3781856115ae565b90508181036060830152611b4b81846115ae565b905095945050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611bc757607f821691505b602082108103611bda57611bd9611b83565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611c3c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611c01565b611c468683611c01565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611c81611c7c611c7784611a60565b611c5e565b611a60565b9050919050565b5f819050919050565b611c9a83611c67565b611cae611ca682611c88565b848454611c0d565b825550505050565b5f5f905090565b611cc5611cb6565b611cd0818484611c91565b505050565b5b81811015611cf357611ce85f82611cbd565b600181019050611cd6565b5050565b601f821115611d3857611d0981611be0565b611d1284611bf2565b81016020851015611d21578190505b611d35611d2d85611bf2565b830182611cd5565b50505b505050565b5f82821c905092915050565b5f611d585f1984600802611d3d565b1980831691505092915050565b5f611d708383611d49565b9150826002028217905092915050565b611d8982611576565b67ffffffffffffffff811115611da257611da1611b56565b5b611dac8254611bb0565b611db7828285611cf7565b5f60209050601f831160018114611de8575f8415611dd6578287015190505b611de08582611d65565b865550611e47565b601f198416611df686611be0565b5f5b82811015611e1d57848901518255600182019150602085019450602081019050611df8565b86831015611e3a5784890151611e36601f891682611d49565b8355505b6001600288020188555050505b505050505050565b828183375f83830152505050565b5f611e688385611580565b9350611e75838584611e4f565b611e7e8361159e565b840190509392505050565b5f6060820190508181035f830152611ea281888a611e5d565b90508181036020830152611eb7818688611e5d565b90508181036040830152611ecc818486611e5d565b9050979650505050505050565b7f4e6f20496e744172726179456e74727920666f72207468697320726f756e64005f82015250565b5f611f0d601f83611580565b9150611f1882611ed9565b602082019050919050565b5f6020820190508181035f830152611f3a81611f01565b9050919050565b7f4e6f204279746573456e74727920666f72207468697320726f756e64000000005f82015250565b5f611f75601c83611580565b9150611f8082611f41565b602082019050919050565b5f6020820190508181035f830152611fa281611f69565b9050919050565b611fb281611886565b82525050565b5f602082019050611fcb5f830184611fa9565b92915050565b611fda8261159e565b810181811067ffffffffffffffff82111715611ff957611ff8611b56565b5b80604052505050565b5f61200b61120d565b90506120178282611fd1565b919050565b5f67ffffffffffffffff82111561203657612035611b56565b5b602082029050602081019050919050565b5f5ffd5b5f67ffffffffffffffff82111561206557612064611b56565b5b61206e8261159e565b9050602081019050919050565b5f61208d6120888461204b565b612002565b9050828152602081018484840111156120a9576120a8612047565b5b6120b4848285611e4f565b509392505050565b5f82601f8301126120d0576120cf61121e565b5b81356120e084826020860161207b565b91505092915050565b5f6120fb6120f68461201c565b612002565b9050808382526020820190506020840283018581111561211e5761211d611226565b5b835b8181101561216557803567ffffffffffffffff8111156121435761214261121e565b5b80860161215089826120bc565b85526020850194505050602081019050612120565b5050509392505050565b5f61217b3684846120e9565b905092915050565b7f4e6f20496e7452616e6765456e74727920666f72207468697320726f756e64005f82015250565b5f6121b7601f83611580565b91506121c282612183565b602082019050919050565b5f6020820190508181035f8301526121e4816121ab565b905091905056fea2646970667358221220ce31b27c53528f6d14c2ef45f87b8d535f5d423078d9e34f1fdc5fc2c40882ad64736f6c634300081c0033",
}

// DIAOracleRandomnessABI is the input ABI used to generate the binding from.
// Deprecated: Use DIAOracleRandomnessMetaData.ABI instead.
var DIAOracleRandomnessABI = DIAOracleRandomnessMetaData.ABI

// DIAOracleRandomnessBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DIAOracleRandomnessMetaData.Bin instead.
var DIAOracleRandomnessBin = DIAOracleRandomnessMetaData.Bin

// DeployDIAOracleRandomness deploys a new Ethereum contract, binding an instance of DIAOracleRandomness to it.
func DeployDIAOracleRandomness(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DIAOracleRandomness, error) {
	parsed, err := DIAOracleRandomnessMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DIAOracleRandomnessBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DIAOracleRandomness{DIAOracleRandomnessCaller: DIAOracleRandomnessCaller{contract: contract}, DIAOracleRandomnessTransactor: DIAOracleRandomnessTransactor{contract: contract}, DIAOracleRandomnessFilterer: DIAOracleRandomnessFilterer{contract: contract}}, nil
}

// DIAOracleRandomness is an auto generated Go binding around an Ethereum contract.
type DIAOracleRandomness struct {
	DIAOracleRandomnessCaller     // Read-only binding to the contract
	DIAOracleRandomnessTransactor // Write-only binding to the contract
	DIAOracleRandomnessFilterer   // Log filterer for contract events
}

// DIAOracleRandomnessCaller is an auto generated read-only Go binding around an Ethereum contract.
type DIAOracleRandomnessCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIAOracleRandomnessTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DIAOracleRandomnessTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIAOracleRandomnessFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DIAOracleRandomnessFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIAOracleRandomnessSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DIAOracleRandomnessSession struct {
	Contract     *DIAOracleRandomness // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DIAOracleRandomnessCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DIAOracleRandomnessCallerSession struct {
	Contract *DIAOracleRandomnessCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// DIAOracleRandomnessTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DIAOracleRandomnessTransactorSession struct {
	Contract     *DIAOracleRandomnessTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// DIAOracleRandomnessRaw is an auto generated low-level Go binding around an Ethereum contract.
type DIAOracleRandomnessRaw struct {
	Contract *DIAOracleRandomness // Generic contract binding to access the raw methods on
}

// DIAOracleRandomnessCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DIAOracleRandomnessCallerRaw struct {
	Contract *DIAOracleRandomnessCaller // Generic read-only contract binding to access the raw methods on
}

// DIAOracleRandomnessTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DIAOracleRandomnessTransactorRaw struct {
	Contract *DIAOracleRandomnessTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDIAOracleRandomness creates a new instance of DIAOracleRandomness, bound to a specific deployed contract.
func NewDIAOracleRandomness(address common.Address, backend bind.ContractBackend) (*DIAOracleRandomness, error) {
	contract, err := bindDIAOracleRandomness(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DIAOracleRandomness{DIAOracleRandomnessCaller: DIAOracleRandomnessCaller{contract: contract}, DIAOracleRandomnessTransactor: DIAOracleRandomnessTransactor{contract: contract}, DIAOracleRandomnessFilterer: DIAOracleRandomnessFilterer{contract: contract}}, nil
}

// NewDIAOracleRandomnessCaller creates a new read-only instance of DIAOracleRandomness, bound to a specific deployed contract.
func NewDIAOracleRandomnessCaller(address common.Address, caller bind.ContractCaller) (*DIAOracleRandomnessCaller, error) {
	contract, err := bindDIAOracleRandomness(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DIAOracleRandomnessCaller{contract: contract}, nil
}

// NewDIAOracleRandomnessTransactor creates a new write-only instance of DIAOracleRandomness, bound to a specific deployed contract.
func NewDIAOracleRandomnessTransactor(address common.Address, transactor bind.ContractTransactor) (*DIAOracleRandomnessTransactor, error) {
	contract, err := bindDIAOracleRandomness(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DIAOracleRandomnessTransactor{contract: contract}, nil
}

// NewDIAOracleRandomnessFilterer creates a new log filterer instance of DIAOracleRandomness, bound to a specific deployed contract.
func NewDIAOracleRandomnessFilterer(address common.Address, filterer bind.ContractFilterer) (*DIAOracleRandomnessFilterer, error) {
	contract, err := bindDIAOracleRandomness(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DIAOracleRandomnessFilterer{contract: contract}, nil
}

// bindDIAOracleRandomness binds a generic wrapper to an already deployed contract.
func bindDIAOracleRandomness(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DIAOracleRandomnessMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIAOracleRandomness *DIAOracleRandomnessRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DIAOracleRandomness.Contract.DIAOracleRandomnessCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIAOracleRandomness *DIAOracleRandomnessRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.DIAOracleRandomnessTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIAOracleRandomness *DIAOracleRandomnessRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.DIAOracleRandomnessTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIAOracleRandomness *DIAOracleRandomnessCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DIAOracleRandomness.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIAOracleRandomness *DIAOracleRandomnessTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIAOracleRandomness *DIAOracleRandomnessTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.contract.Transact(opts, method, params...)
}

// GetBytes is a free data retrieval call binding the contract method 0x53759265.
//
// Solidity: function getBytes(int64 round) view returns(string requestId, string[] randomness, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCaller) GetBytes(opts *bind.CallOpts, round int64) (struct {
	RequestId  string
	Randomness []string
	Seed       string
	Signature  string
}, error) {
	var out []interface{}
	err := _DIAOracleRandomness.contract.Call(opts, &out, "getBytes", round)

	outstruct := new(struct {
		RequestId  string
		Randomness []string
		Seed       string
		Signature  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Randomness = *abi.ConvertType(out[1], new([]string)).(*[]string)
	outstruct.Seed = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Signature = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// GetBytes is a free data retrieval call binding the contract method 0x53759265.
//
// Solidity: function getBytes(int64 round) view returns(string requestId, string[] randomness, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessSession) GetBytes(round int64) (struct {
	RequestId  string
	Randomness []string
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetBytes(&_DIAOracleRandomness.CallOpts, round)
}

// GetBytes is a free data retrieval call binding the contract method 0x53759265.
//
// Solidity: function getBytes(int64 round) view returns(string requestId, string[] randomness, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCallerSession) GetBytes(round int64) (struct {
	RequestId  string
	Randomness []string
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetBytes(&_DIAOracleRandomness.CallOpts, round)
}

// GetIntArray is a free data retrieval call binding the contract method 0x45459cd0.
//
// Solidity: function getIntArray(int64 round) view returns(string requestId, int256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCaller) GetIntArray(opts *bind.CallOpts, round int64) (struct {
	RequestId  string
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	var out []interface{}
	err := _DIAOracleRandomness.contract.Call(opts, &out, "getIntArray", round)

	outstruct := new(struct {
		RequestId  string
		RandomInts []*big.Int
		Seed       string
		Signature  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.RandomInts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Seed = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Signature = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// GetIntArray is a free data retrieval call binding the contract method 0x45459cd0.
//
// Solidity: function getIntArray(int64 round) view returns(string requestId, int256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessSession) GetIntArray(round int64) (struct {
	RequestId  string
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetIntArray(&_DIAOracleRandomness.CallOpts, round)
}

// GetIntArray is a free data retrieval call binding the contract method 0x45459cd0.
//
// Solidity: function getIntArray(int64 round) view returns(string requestId, int256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCallerSession) GetIntArray(round int64) (struct {
	RequestId  string
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetIntArray(&_DIAOracleRandomness.CallOpts, round)
}

// GetIntRange is a free data retrieval call binding the contract method 0xd9082a9b.
//
// Solidity: function getIntRange(int64 round) view returns(string requestId, uint256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCaller) GetIntRange(opts *bind.CallOpts, round int64) (struct {
	RequestId  string
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	var out []interface{}
	err := _DIAOracleRandomness.contract.Call(opts, &out, "getIntRange", round)

	outstruct := new(struct {
		RequestId  string
		RandomInts []*big.Int
		Seed       string
		Signature  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.RandomInts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Seed = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Signature = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// GetIntRange is a free data retrieval call binding the contract method 0xd9082a9b.
//
// Solidity: function getIntRange(int64 round) view returns(string requestId, uint256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessSession) GetIntRange(round int64) (struct {
	RequestId  string
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetIntRange(&_DIAOracleRandomness.CallOpts, round)
}

// GetIntRange is a free data retrieval call binding the contract method 0xd9082a9b.
//
// Solidity: function getIntRange(int64 round) view returns(string requestId, uint256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCallerSession) GetIntRange(round int64) (struct {
	RequestId  string
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetIntRange(&_DIAOracleRandomness.CallOpts, round)
}

// SetBytes is a paid mutator transaction binding the contract method 0xad29e1da.
//
// Solidity: function setBytes(string requestId, string[] randomness, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactor) SetBytes(opts *bind.TransactOpts, requestId string, randomness []string, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.contract.Transact(opts, "setBytes", requestId, randomness, round, seed, signature)
}

// SetBytes is a paid mutator transaction binding the contract method 0xad29e1da.
//
// Solidity: function setBytes(string requestId, string[] randomness, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessSession) SetBytes(requestId string, randomness []string, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.SetBytes(&_DIAOracleRandomness.TransactOpts, requestId, randomness, round, seed, signature)
}

// SetBytes is a paid mutator transaction binding the contract method 0xad29e1da.
//
// Solidity: function setBytes(string requestId, string[] randomness, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactorSession) SetBytes(requestId string, randomness []string, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.SetBytes(&_DIAOracleRandomness.TransactOpts, requestId, randomness, round, seed, signature)
}

// SetIntArray is a paid mutator transaction binding the contract method 0x07c82759.
//
// Solidity: function setIntArray(string requestId, int256[] randomInts, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactor) SetIntArray(opts *bind.TransactOpts, requestId string, randomInts []*big.Int, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.contract.Transact(opts, "setIntArray", requestId, randomInts, round, seed, signature)
}

// SetIntArray is a paid mutator transaction binding the contract method 0x07c82759.
//
// Solidity: function setIntArray(string requestId, int256[] randomInts, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessSession) SetIntArray(requestId string, randomInts []*big.Int, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.SetIntArray(&_DIAOracleRandomness.TransactOpts, requestId, randomInts, round, seed, signature)
}

// SetIntArray is a paid mutator transaction binding the contract method 0x07c82759.
//
// Solidity: function setIntArray(string requestId, int256[] randomInts, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactorSession) SetIntArray(requestId string, randomInts []*big.Int, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.SetIntArray(&_DIAOracleRandomness.TransactOpts, requestId, randomInts, round, seed, signature)
}

// SetIntRange is a paid mutator transaction binding the contract method 0x154122d1.
//
// Solidity: function setIntRange(string requestId, uint256[] randomInts, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactor) SetIntRange(opts *bind.TransactOpts, requestId string, randomInts []*big.Int, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.contract.Transact(opts, "setIntRange", requestId, randomInts, round, seed, signature)
}

// SetIntRange is a paid mutator transaction binding the contract method 0x154122d1.
//
// Solidity: function setIntRange(string requestId, uint256[] randomInts, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessSession) SetIntRange(requestId string, randomInts []*big.Int, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.SetIntRange(&_DIAOracleRandomness.TransactOpts, requestId, randomInts, round, seed, signature)
}

// SetIntRange is a paid mutator transaction binding the contract method 0x154122d1.
//
// Solidity: function setIntRange(string requestId, uint256[] randomInts, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactorSession) SetIntRange(requestId string, randomInts []*big.Int, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.SetIntRange(&_DIAOracleRandomness.TransactOpts, requestId, randomInts, round, seed, signature)
}

// UpdateOracleUpdaterAddress is a paid mutator transaction binding the contract method 0x6aa45efc.
//
// Solidity: function updateOracleUpdaterAddress(address newOracleUpdaterAddress) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactor) UpdateOracleUpdaterAddress(opts *bind.TransactOpts, newOracleUpdaterAddress common.Address) (*types.Transaction, error) {
	return _DIAOracleRandomness.contract.Transact(opts, "updateOracleUpdaterAddress", newOracleUpdaterAddress)
}

// UpdateOracleUpdaterAddress is a paid mutator transaction binding the contract method 0x6aa45efc.
//
// Solidity: function updateOracleUpdaterAddress(address newOracleUpdaterAddress) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessSession) UpdateOracleUpdaterAddress(newOracleUpdaterAddress common.Address) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.UpdateOracleUpdaterAddress(&_DIAOracleRandomness.TransactOpts, newOracleUpdaterAddress)
}

// UpdateOracleUpdaterAddress is a paid mutator transaction binding the contract method 0x6aa45efc.
//
// Solidity: function updateOracleUpdaterAddress(address newOracleUpdaterAddress) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactorSession) UpdateOracleUpdaterAddress(newOracleUpdaterAddress common.Address) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.UpdateOracleUpdaterAddress(&_DIAOracleRandomness.TransactOpts, newOracleUpdaterAddress)
}

// DIAOracleRandomnessBytesSetIterator is returned from FilterBytesSet and is used to iterate over the raw logs and unpacked data for BytesSet events raised by the DIAOracleRandomness contract.
type DIAOracleRandomnessBytesSetIterator struct {
	Event *DIAOracleRandomnessBytesSet // Event containing the contract specifics and raw log

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
func (it *DIAOracleRandomnessBytesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIAOracleRandomnessBytesSet)
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
		it.Event = new(DIAOracleRandomnessBytesSet)
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
func (it *DIAOracleRandomnessBytesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIAOracleRandomnessBytesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIAOracleRandomnessBytesSet represents a BytesSet event raised by the DIAOracleRandomness contract.
type DIAOracleRandomnessBytesSet struct {
	RequestId string
	Round     *big.Int
	Seed      string
	Signature string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBytesSet is a free log retrieval operation binding the contract event 0x257f2a31afdd06af9923520043e89ea1d5bf9cf877a0e5a97cbd81332e6c1d0b.
//
// Solidity: event BytesSet(string requestId, int256 indexed round, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessFilterer) FilterBytesSet(opts *bind.FilterOpts, round []*big.Int) (*DIAOracleRandomnessBytesSetIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _DIAOracleRandomness.contract.FilterLogs(opts, "BytesSet", roundRule)
	if err != nil {
		return nil, err
	}
	return &DIAOracleRandomnessBytesSetIterator{contract: _DIAOracleRandomness.contract, event: "BytesSet", logs: logs, sub: sub}, nil
}

// WatchBytesSet is a free log subscription operation binding the contract event 0x257f2a31afdd06af9923520043e89ea1d5bf9cf877a0e5a97cbd81332e6c1d0b.
//
// Solidity: event BytesSet(string requestId, int256 indexed round, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessFilterer) WatchBytesSet(opts *bind.WatchOpts, sink chan<- *DIAOracleRandomnessBytesSet, round []*big.Int) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _DIAOracleRandomness.contract.WatchLogs(opts, "BytesSet", roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIAOracleRandomnessBytesSet)
				if err := _DIAOracleRandomness.contract.UnpackLog(event, "BytesSet", log); err != nil {
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

// ParseBytesSet is a log parse operation binding the contract event 0x257f2a31afdd06af9923520043e89ea1d5bf9cf877a0e5a97cbd81332e6c1d0b.
//
// Solidity: event BytesSet(string requestId, int256 indexed round, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessFilterer) ParseBytesSet(log types.Log) (*DIAOracleRandomnessBytesSet, error) {
	event := new(DIAOracleRandomnessBytesSet)
	if err := _DIAOracleRandomness.contract.UnpackLog(event, "BytesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DIAOracleRandomnessIntArraySetIterator is returned from FilterIntArraySet and is used to iterate over the raw logs and unpacked data for IntArraySet events raised by the DIAOracleRandomness contract.
type DIAOracleRandomnessIntArraySetIterator struct {
	Event *DIAOracleRandomnessIntArraySet // Event containing the contract specifics and raw log

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
func (it *DIAOracleRandomnessIntArraySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIAOracleRandomnessIntArraySet)
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
		it.Event = new(DIAOracleRandomnessIntArraySet)
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
func (it *DIAOracleRandomnessIntArraySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIAOracleRandomnessIntArraySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIAOracleRandomnessIntArraySet represents a IntArraySet event raised by the DIAOracleRandomness contract.
type DIAOracleRandomnessIntArraySet struct {
	RequestId string
	Round     *big.Int
	Seed      string
	Signature string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIntArraySet is a free log retrieval operation binding the contract event 0xdb8e3995a930d402dc2af8569cbaaa06b6be56a42d4014fb12c83c33232e9e64.
//
// Solidity: event IntArraySet(string requestId, int256 indexed round, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessFilterer) FilterIntArraySet(opts *bind.FilterOpts, round []*big.Int) (*DIAOracleRandomnessIntArraySetIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _DIAOracleRandomness.contract.FilterLogs(opts, "IntArraySet", roundRule)
	if err != nil {
		return nil, err
	}
	return &DIAOracleRandomnessIntArraySetIterator{contract: _DIAOracleRandomness.contract, event: "IntArraySet", logs: logs, sub: sub}, nil
}

// WatchIntArraySet is a free log subscription operation binding the contract event 0xdb8e3995a930d402dc2af8569cbaaa06b6be56a42d4014fb12c83c33232e9e64.
//
// Solidity: event IntArraySet(string requestId, int256 indexed round, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessFilterer) WatchIntArraySet(opts *bind.WatchOpts, sink chan<- *DIAOracleRandomnessIntArraySet, round []*big.Int) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _DIAOracleRandomness.contract.WatchLogs(opts, "IntArraySet", roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIAOracleRandomnessIntArraySet)
				if err := _DIAOracleRandomness.contract.UnpackLog(event, "IntArraySet", log); err != nil {
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

// ParseIntArraySet is a log parse operation binding the contract event 0xdb8e3995a930d402dc2af8569cbaaa06b6be56a42d4014fb12c83c33232e9e64.
//
// Solidity: event IntArraySet(string requestId, int256 indexed round, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessFilterer) ParseIntArraySet(log types.Log) (*DIAOracleRandomnessIntArraySet, error) {
	event := new(DIAOracleRandomnessIntArraySet)
	if err := _DIAOracleRandomness.contract.UnpackLog(event, "IntArraySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DIAOracleRandomnessIntRangeSetIterator is returned from FilterIntRangeSet and is used to iterate over the raw logs and unpacked data for IntRangeSet events raised by the DIAOracleRandomness contract.
type DIAOracleRandomnessIntRangeSetIterator struct {
	Event *DIAOracleRandomnessIntRangeSet // Event containing the contract specifics and raw log

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
func (it *DIAOracleRandomnessIntRangeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIAOracleRandomnessIntRangeSet)
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
		it.Event = new(DIAOracleRandomnessIntRangeSet)
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
func (it *DIAOracleRandomnessIntRangeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIAOracleRandomnessIntRangeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIAOracleRandomnessIntRangeSet represents a IntRangeSet event raised by the DIAOracleRandomness contract.
type DIAOracleRandomnessIntRangeSet struct {
	RequestId string
	Round     *big.Int
	Seed      string
	Signature string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIntRangeSet is a free log retrieval operation binding the contract event 0x0143668be676fa7bd52d90cc60925a18d048062ff315d7542e640da4f2e35547.
//
// Solidity: event IntRangeSet(string requestId, int256 indexed round, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessFilterer) FilterIntRangeSet(opts *bind.FilterOpts, round []*big.Int) (*DIAOracleRandomnessIntRangeSetIterator, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _DIAOracleRandomness.contract.FilterLogs(opts, "IntRangeSet", roundRule)
	if err != nil {
		return nil, err
	}
	return &DIAOracleRandomnessIntRangeSetIterator{contract: _DIAOracleRandomness.contract, event: "IntRangeSet", logs: logs, sub: sub}, nil
}

// WatchIntRangeSet is a free log subscription operation binding the contract event 0x0143668be676fa7bd52d90cc60925a18d048062ff315d7542e640da4f2e35547.
//
// Solidity: event IntRangeSet(string requestId, int256 indexed round, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessFilterer) WatchIntRangeSet(opts *bind.WatchOpts, sink chan<- *DIAOracleRandomnessIntRangeSet, round []*big.Int) (event.Subscription, error) {

	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}

	logs, sub, err := _DIAOracleRandomness.contract.WatchLogs(opts, "IntRangeSet", roundRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIAOracleRandomnessIntRangeSet)
				if err := _DIAOracleRandomness.contract.UnpackLog(event, "IntRangeSet", log); err != nil {
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

// ParseIntRangeSet is a log parse operation binding the contract event 0x0143668be676fa7bd52d90cc60925a18d048062ff315d7542e640da4f2e35547.
//
// Solidity: event IntRangeSet(string requestId, int256 indexed round, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessFilterer) ParseIntRangeSet(log types.Log) (*DIAOracleRandomnessIntRangeSet, error) {
	event := new(DIAOracleRandomnessIntRangeSet)
	if err := _DIAOracleRandomness.contract.UnpackLog(event, "IntRangeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DIAOracleRandomnessUpdaterAddressChangeIterator is returned from FilterUpdaterAddressChange and is used to iterate over the raw logs and unpacked data for UpdaterAddressChange events raised by the DIAOracleRandomness contract.
type DIAOracleRandomnessUpdaterAddressChangeIterator struct {
	Event *DIAOracleRandomnessUpdaterAddressChange // Event containing the contract specifics and raw log

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
func (it *DIAOracleRandomnessUpdaterAddressChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIAOracleRandomnessUpdaterAddressChange)
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
		it.Event = new(DIAOracleRandomnessUpdaterAddressChange)
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
func (it *DIAOracleRandomnessUpdaterAddressChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIAOracleRandomnessUpdaterAddressChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIAOracleRandomnessUpdaterAddressChange represents a UpdaterAddressChange event raised by the DIAOracleRandomness contract.
type DIAOracleRandomnessUpdaterAddressChange struct {
	NewUpdater common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterUpdaterAddressChange is a free log retrieval operation binding the contract event 0x121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f.
//
// Solidity: event UpdaterAddressChange(address newUpdater)
func (_DIAOracleRandomness *DIAOracleRandomnessFilterer) FilterUpdaterAddressChange(opts *bind.FilterOpts) (*DIAOracleRandomnessUpdaterAddressChangeIterator, error) {

	logs, sub, err := _DIAOracleRandomness.contract.FilterLogs(opts, "UpdaterAddressChange")
	if err != nil {
		return nil, err
	}
	return &DIAOracleRandomnessUpdaterAddressChangeIterator{contract: _DIAOracleRandomness.contract, event: "UpdaterAddressChange", logs: logs, sub: sub}, nil
}

// WatchUpdaterAddressChange is a free log subscription operation binding the contract event 0x121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f.
//
// Solidity: event UpdaterAddressChange(address newUpdater)
func (_DIAOracleRandomness *DIAOracleRandomnessFilterer) WatchUpdaterAddressChange(opts *bind.WatchOpts, sink chan<- *DIAOracleRandomnessUpdaterAddressChange) (event.Subscription, error) {

	logs, sub, err := _DIAOracleRandomness.contract.WatchLogs(opts, "UpdaterAddressChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIAOracleRandomnessUpdaterAddressChange)
				if err := _DIAOracleRandomness.contract.UnpackLog(event, "UpdaterAddressChange", log); err != nil {
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

// ParseUpdaterAddressChange is a log parse operation binding the contract event 0x121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f.
//
// Solidity: event UpdaterAddressChange(address newUpdater)
func (_DIAOracleRandomness *DIAOracleRandomnessFilterer) ParseUpdaterAddressChange(log types.Log) (*DIAOracleRandomnessUpdaterAddressChange, error) {
	event := new(DIAOracleRandomnessUpdaterAddressChange)
	if err := _DIAOracleRandomness.contract.UnpackLog(event, "UpdaterAddressChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
