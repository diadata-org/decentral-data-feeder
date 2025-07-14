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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBytes\",\"inputs\":[{\"name\":\"requestId_\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"randomness\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIntArray\",\"inputs\":[{\"name\":\"requestId_\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"randomInts\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIntRange\",\"inputs\":[{\"name\":\"requestId_\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"randomInts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBytes\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"randomness\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIntArray\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"randomInts\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIntRange\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"randomInts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOracleUpdaterAddress\",\"inputs\":[{\"name\":\"newOracleUpdaterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BytesSet\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"round\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IntArraySet\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"round\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IntRangeSet\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"round\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdaterAddressChange\",\"inputs\":[{\"name\":\"newUpdater\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f5ffd5b503360035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506122f98061005c5f395ff3fe608060405234801561000f575f5ffd5b506004361061007b575f3560e01c80636aa45efc116100595780636aa45efc146100ea578063ad29e1da14610106578063be12f89614610122578063d8de899d146101555761007b565b806307c827591461007f5780630a6ca10f1461009b578063154122d1146100ce575b5f5ffd5b6100996004803603810190610094919061137c565b610188565b005b6100b560048036038101906100b09190611472565b6103aa565b6040516100c594939291906115ed565b60405180910390f35b6100e860048036038101906100e391906116a1565b610651565b005b61010460048036038101906100ff91906117f1565b610873565b005b610120600480360381019061011b9190611871565b610945565b005b61013c60048036038101906101379190611472565b610b32565b60405161014c9493929190611a27565b60405180910390f35b61016f600480360381019061016a9190611472565b610dd9565b60405161017f9493929190611b89565b60405180910390f35b6040518060a001604052808a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018888808060200260200160405190810160405280939291908181526020018383602002808284375f81840152601f19601f82011690508083019250505050505050815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200183838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020016001151581525060028a8a6040516102cd929190611c24565b90815260200160405180910390205f820151815f0190816102ee9190611e66565b50602082015181600101908051906020019061030b9291906110fc565b5060408201518160020190816103219190611e66565b5060608201518160030190816103379190611e66565b506080820151816004015f6101000a81548160ff0219169083151502179055509050508460070b7fdb8e3995a930d402dc2af8569cbaaa06b6be56a42d4014fb12c83c33232e9e648a8a8787878760405161039796959493929190611f61565b60405180910390a2505050505050505050565b606080606080600286866040516103c2929190611c24565b90815260200160405180910390206004015f9054906101000a900460ff1661041f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161041690611ffb565b60405180910390fd5b5f60028787604051610432929190611c24565b90815260200160405180910390209050805f0181600101826002018360030183805461045d90611c96565b80601f016020809104026020016040519081016040528092919081815260200182805461048990611c96565b80156104d45780601f106104ab576101008083540402835291602001916104d4565b820191905f5260205f20905b8154815290600101906020018083116104b757829003601f168201915b505050505093508280548060200260200160405190810160405280929190818152602001828054801561052457602002820191905f5260205f20905b815481526020019060010190808311610510575b5050505050925081805461053790611c96565b80601f016020809104026020016040519081016040528092919081815260200182805461056390611c96565b80156105ae5780601f10610585576101008083540402835291602001916105ae565b820191905f5260205f20905b81548152906001019060200180831161059157829003601f168201915b505050505091508080546105c190611c96565b80601f01602080910402602001604051908101604052809291908181526020018280546105ed90611c96565b80156106385780601f1061060f57610100808354040283529160200191610638565b820191905f5260205f20905b81548152906001019060200180831161061b57829003601f168201915b5050505050905094509450945094505092959194509250565b6040518060a001604052808a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018888808060200260200160405190810160405280939291908181526020018383602002808284375f81840152601f19601f82011690508083019250505050505050815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200183838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020016001151581525060018a8a604051610796929190611c24565b90815260200160405180910390205f820151815f0190816107b79190611e66565b5060208201518160010190805190602001906107d4929190611147565b5060408201518160020190816107ea9190611e66565b5060608201518160030190816108009190611e66565b506080820151816004015f6101000a81548160ff0219169083151502179055509050508460070b7f0143668be676fa7bd52d90cc60925a18d048062ff315d7542e640da4f2e355478a8a8787878760405161086096959493929190611f61565b60405180910390a2505050505050505050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146108cb575f5ffd5b8060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f8160405161093a9190612028565b60405180910390a150565b6040518060a001604052808a8a8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020018888906109a691906121df565b815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200183838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050508152602001600115158152505f8a8a604051610a55929190611c24565b90815260200160405180910390205f820151815f019081610a769190611e66565b506020820151816001019080519060200190610a93929190611192565b506040820151816002019081610aa99190611e66565b506060820151816003019081610abf9190611e66565b506080820151816004015f6101000a81548160ff0219169083151502179055509050508460070b7f257f2a31afdd06af9923520043e89ea1d5bf9cf877a0e5a97cbd81332e6c1d0b8a8a87878787604051610b1f96959493929190611f61565b60405180910390a2505050505050505050565b60608060608060018686604051610b4a929190611c24565b90815260200160405180910390206004015f9054906101000a900460ff16610ba7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b9e9061223d565b60405180910390fd5b5f60018787604051610bba929190611c24565b90815260200160405180910390209050805f01816001018260020183600301838054610be590611c96565b80601f0160208091040260200160405190810160405280929190818152602001828054610c1190611c96565b8015610c5c5780601f10610c3357610100808354040283529160200191610c5c565b820191905f5260205f20905b815481529060010190602001808311610c3f57829003601f168201915b5050505050935082805480602002602001604051908101604052809291908181526020018280548015610cac57602002820191905f5260205f20905b815481526020019060010190808311610c98575b50505050509250818054610cbf90611c96565b80601f0160208091040260200160405190810160405280929190818152602001828054610ceb90611c96565b8015610d365780601f10610d0d57610100808354040283529160200191610d36565b820191905f5260205f20905b815481529060010190602001808311610d1957829003601f168201915b50505050509150808054610d4990611c96565b80601f0160208091040260200160405190810160405280929190818152602001828054610d7590611c96565b8015610dc05780601f10610d9757610100808354040283529160200191610dc0565b820191905f5260205f20905b815481529060010190602001808311610da357829003601f168201915b5050505050905094509450945094505092959194509250565b6060806060805f8686604051610df0929190611c24565b90815260200160405180910390206004015f9054906101000a900460ff16610e4d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e44906122a5565b60405180910390fd5b5f5f8787604051610e5f929190611c24565b90815260200160405180910390209050805f01816001018260020183600301838054610e8a90611c96565b80601f0160208091040260200160405190810160405280929190818152602001828054610eb690611c96565b8015610f015780601f10610ed857610100808354040283529160200191610f01565b820191905f5260205f20905b815481529060010190602001808311610ee457829003601f168201915b5050505050935082805480602002602001604051908101604052809291908181526020015f905b82821015610fd0578382905f5260205f20018054610f4590611c96565b80601f0160208091040260200160405190810160405280929190818152602001828054610f7190611c96565b8015610fbc5780601f10610f9357610100808354040283529160200191610fbc565b820191905f5260205f20905b815481529060010190602001808311610f9f57829003601f168201915b505050505081526020019060010190610f28565b505050509250818054610fe290611c96565b80601f016020809104026020016040519081016040528092919081815260200182805461100e90611c96565b80156110595780601f1061103057610100808354040283529160200191611059565b820191905f5260205f20905b81548152906001019060200180831161103c57829003601f168201915b5050505050915080805461106c90611c96565b80601f016020809104026020016040519081016040528092919081815260200182805461109890611c96565b80156110e35780601f106110ba576101008083540402835291602001916110e3565b820191905f5260205f20905b8154815290600101906020018083116110c657829003601f168201915b5050505050905094509450945094505092959194509250565b828054828255905f5260205f20908101928215611136579160200282015b8281111561113557825182559160200191906001019061111a565b5b50905061114391906111e9565b5090565b828054828255905f5260205f20908101928215611181579160200282015b82811115611180578251825591602001919060010190611165565b5b50905061118e9190611204565b5090565b828054828255905f5260205f209081019282156111d8579160200282015b828111156111d75782518290816111c79190611e66565b50916020019190600101906111b0565b5b5090506111e5919061121f565b5090565b5b80821115611200575f815f9055506001016111ea565b5090565b5b8082111561121b575f815f905550600101611205565b5090565b5b8082111561123e575f81816112359190611242565b50600101611220565b5090565b50805461124e90611c96565b5f825580601f1061125f575061127c565b601f0160209004905f5260205f209081019061127b9190611204565b5b50565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126112b1576112b0611290565b5b8235905067ffffffffffffffff8111156112ce576112cd611294565b5b6020830191508360018202830111156112ea576112e9611298565b5b9250929050565b5f5f83601f84011261130657611305611290565b5b8235905067ffffffffffffffff81111561132357611322611294565b5b60208301915083602082028301111561133f5761133e611298565b5b9250929050565b5f8160070b9050919050565b61135b81611346565b8114611365575f5ffd5b50565b5f8135905061137681611352565b92915050565b5f5f5f5f5f5f5f5f5f60a08a8c03121561139957611398611288565b5b5f8a013567ffffffffffffffff8111156113b6576113b561128c565b5b6113c28c828d0161129c565b995099505060208a013567ffffffffffffffff8111156113e5576113e461128c565b5b6113f18c828d016112f1565b975097505060406114048c828d01611368565b95505060608a013567ffffffffffffffff8111156114255761142461128c565b5b6114318c828d0161129c565b945094505060808a013567ffffffffffffffff8111156114545761145361128c565b5b6114608c828d0161129c565b92509250509295985092959850929598565b5f5f6020838503121561148857611487611288565b5b5f83013567ffffffffffffffff8111156114a5576114a461128c565b5b6114b18582860161129c565b92509250509250929050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6114ff826114bd565b61150981856114c7565b93506115198185602086016114d7565b611522816114e5565b840191505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b61156881611556565b82525050565b5f611579838361155f565b60208301905092915050565b5f602082019050919050565b5f61159b8261152d565b6115a58185611537565b93506115b083611547565b805f5b838110156115e05781516115c7888261156e565b97506115d283611585565b9250506001810190506115b3565b5085935050505092915050565b5f6080820190508181035f83015261160581876114f5565b905081810360208301526116198186611591565b9050818103604083015261162d81856114f5565b9050818103606083015261164181846114f5565b905095945050505050565b5f5f83601f84011261166157611660611290565b5b8235905067ffffffffffffffff81111561167e5761167d611294565b5b60208301915083602082028301111561169a57611699611298565b5b9250929050565b5f5f5f5f5f5f5f5f5f60a08a8c0312156116be576116bd611288565b5b5f8a013567ffffffffffffffff8111156116db576116da61128c565b5b6116e78c828d0161129c565b995099505060208a013567ffffffffffffffff81111561170a5761170961128c565b5b6117168c828d0161164c565b975097505060406117298c828d01611368565b95505060608a013567ffffffffffffffff81111561174a5761174961128c565b5b6117568c828d0161129c565b945094505060808a013567ffffffffffffffff8111156117795761177861128c565b5b6117858c828d0161129c565b92509250509295985092959850929598565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6117c082611797565b9050919050565b6117d0816117b6565b81146117da575f5ffd5b50565b5f813590506117eb816117c7565b92915050565b5f6020828403121561180657611805611288565b5b5f611813848285016117dd565b91505092915050565b5f5f83601f84011261183157611830611290565b5b8235905067ffffffffffffffff81111561184e5761184d611294565b5b60208301915083602082028301111561186a57611869611298565b5b9250929050565b5f5f5f5f5f5f5f5f5f60a08a8c03121561188e5761188d611288565b5b5f8a013567ffffffffffffffff8111156118ab576118aa61128c565b5b6118b78c828d0161129c565b995099505060208a013567ffffffffffffffff8111156118da576118d961128c565b5b6118e68c828d0161181c565b975097505060406118f98c828d01611368565b95505060608a013567ffffffffffffffff81111561191a5761191961128c565b5b6119268c828d0161129c565b945094505060808a013567ffffffffffffffff8111156119495761194861128c565b5b6119558c828d0161129c565b92509250509295985092959850929598565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b6119a281611990565b82525050565b5f6119b38383611999565b60208301905092915050565b5f602082019050919050565b5f6119d582611967565b6119df8185611971565b93506119ea83611981565b805f5b83811015611a1a578151611a0188826119a8565b9750611a0c836119bf565b9250506001810190506119ed565b5085935050505092915050565b5f6080820190508181035f830152611a3f81876114f5565b90508181036020830152611a5381866119cb565b90508181036040830152611a6781856114f5565b90508181036060830152611a7b81846114f5565b905095945050505050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f82825260208201905092915050565b5f611ac9826114bd565b611ad38185611aaf565b9350611ae38185602086016114d7565b611aec816114e5565b840191505092915050565b5f611b028383611abf565b905092915050565b5f602082019050919050565b5f611b2082611a86565b611b2a8185611a90565b935083602082028501611b3c85611aa0565b805f5b85811015611b775784840389528151611b588582611af7565b9450611b6383611b0a565b925060208a01995050600181019050611b3f565b50829750879550505050505092915050565b5f6080820190508181035f830152611ba181876114f5565b90508181036020830152611bb58186611b16565b90508181036040830152611bc981856114f5565b90508181036060830152611bdd81846114f5565b905095945050505050565b5f81905092915050565b828183375f83830152505050565b5f611c0b8385611be8565b9350611c18838584611bf2565b82840190509392505050565b5f611c30828486611c00565b91508190509392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680611cad57607f821691505b602082108103611cc057611cbf611c69565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611d227fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611ce7565b611d2c8683611ce7565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611d67611d62611d5d84611990565b611d44565b611990565b9050919050565b5f819050919050565b611d8083611d4d565b611d94611d8c82611d6e565b848454611cf3565b825550505050565b5f5f905090565b611dab611d9c565b611db6818484611d77565b505050565b5b81811015611dd957611dce5f82611da3565b600181019050611dbc565b5050565b601f821115611e1e57611def81611cc6565b611df884611cd8565b81016020851015611e07578190505b611e1b611e1385611cd8565b830182611dbb565b50505b505050565b5f82821c905092915050565b5f611e3e5f1984600802611e23565b1980831691505092915050565b5f611e568383611e2f565b9150826002028217905092915050565b611e6f826114bd565b67ffffffffffffffff811115611e8857611e87611c3c565b5b611e928254611c96565b611e9d828285611ddd565b5f60209050601f831160018114611ece575f8415611ebc578287015190505b611ec68582611e4b565b865550611f2d565b601f198416611edc86611cc6565b5f5b82811015611f0357848901518255600182019150602085019450602081019050611ede565b86831015611f205784890151611f1c601f891682611e2f565b8355505b6001600288020188555050505b505050505050565b5f611f4083856114c7565b9350611f4d838584611bf2565b611f56836114e5565b840190509392505050565b5f6060820190508181035f830152611f7a81888a611f35565b90508181036020830152611f8f818688611f35565b90508181036040830152611fa4818486611f35565b9050979650505050505050565b7f4e6f20496e744172726179456e74727920666f72207468697320726f756e64005f82015250565b5f611fe5601f836114c7565b9150611ff082611fb1565b602082019050919050565b5f6020820190508181035f83015261201281611fd9565b9050919050565b612022816117b6565b82525050565b5f60208201905061203b5f830184612019565b92915050565b61204a826114e5565b810181811067ffffffffffffffff8211171561206957612068611c3c565b5b80604052505050565b5f61207b61127f565b90506120878282612041565b919050565b5f67ffffffffffffffff8211156120a6576120a5611c3c565b5b602082029050602081019050919050565b5f5ffd5b5f67ffffffffffffffff8211156120d5576120d4611c3c565b5b6120de826114e5565b9050602081019050919050565b5f6120fd6120f8846120bb565b612072565b905082815260208101848484011115612119576121186120b7565b5b612124848285611bf2565b509392505050565b5f82601f8301126121405761213f611290565b5b81356121508482602086016120eb565b91505092915050565b5f61216b6121668461208c565b612072565b9050808382526020820190506020840283018581111561218e5761218d611298565b5b835b818110156121d557803567ffffffffffffffff8111156121b3576121b2611290565b5b8086016121c0898261212c565b85526020850194505050602081019050612190565b5050509392505050565b5f6121eb368484612159565b905092915050565b7f4e6f20496e7452616e6765456e74727920666f72207468697320726f756e64005f82015250565b5f612227601f836114c7565b9150612232826121f3565b602082019050919050565b5f6020820190508181035f8301526122548161221b565b9050919050565b7f4e6f204279746573456e74727920666f722074686973207265717565737449645f82015250565b5f61228f6020836114c7565b915061229a8261225b565b602082019050919050565b5f6020820190508181035f8301526122bc81612283565b905091905056fea264697066735822122059e9800f0f0004b42bc78969d7cd0f894f12cb1543097dcb66973e59dd170b7564736f6c634300081c0033",
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

// GetBytes is a free data retrieval call binding the contract method 0xd8de899d.
//
// Solidity: function getBytes(string requestId_) view returns(string requestId, string[] randomness, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCaller) GetBytes(opts *bind.CallOpts, requestId_ string) (struct {
	RequestId  string
	Randomness []string
	Seed       string
	Signature  string
}, error) {
	var out []interface{}
	err := _DIAOracleRandomness.contract.Call(opts, &out, "getBytes", requestId_)

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

// GetBytes is a free data retrieval call binding the contract method 0xd8de899d.
//
// Solidity: function getBytes(string requestId_) view returns(string requestId, string[] randomness, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessSession) GetBytes(requestId_ string) (struct {
	RequestId  string
	Randomness []string
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetBytes(&_DIAOracleRandomness.CallOpts, requestId_)
}

// GetBytes is a free data retrieval call binding the contract method 0xd8de899d.
//
// Solidity: function getBytes(string requestId_) view returns(string requestId, string[] randomness, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCallerSession) GetBytes(requestId_ string) (struct {
	RequestId  string
	Randomness []string
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetBytes(&_DIAOracleRandomness.CallOpts, requestId_)
}

// GetIntArray is a free data retrieval call binding the contract method 0x0a6ca10f.
//
// Solidity: function getIntArray(string requestId_) view returns(string requestId, int256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCaller) GetIntArray(opts *bind.CallOpts, requestId_ string) (struct {
	RequestId  string
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	var out []interface{}
	err := _DIAOracleRandomness.contract.Call(opts, &out, "getIntArray", requestId_)

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

// GetIntArray is a free data retrieval call binding the contract method 0x0a6ca10f.
//
// Solidity: function getIntArray(string requestId_) view returns(string requestId, int256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessSession) GetIntArray(requestId_ string) (struct {
	RequestId  string
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetIntArray(&_DIAOracleRandomness.CallOpts, requestId_)
}

// GetIntArray is a free data retrieval call binding the contract method 0x0a6ca10f.
//
// Solidity: function getIntArray(string requestId_) view returns(string requestId, int256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCallerSession) GetIntArray(requestId_ string) (struct {
	RequestId  string
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetIntArray(&_DIAOracleRandomness.CallOpts, requestId_)
}

// GetIntRange is a free data retrieval call binding the contract method 0xbe12f896.
//
// Solidity: function getIntRange(string requestId_) view returns(string requestId, uint256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCaller) GetIntRange(opts *bind.CallOpts, requestId_ string) (struct {
	RequestId  string
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	var out []interface{}
	err := _DIAOracleRandomness.contract.Call(opts, &out, "getIntRange", requestId_)

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

// GetIntRange is a free data retrieval call binding the contract method 0xbe12f896.
//
// Solidity: function getIntRange(string requestId_) view returns(string requestId, uint256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessSession) GetIntRange(requestId_ string) (struct {
	RequestId  string
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetIntRange(&_DIAOracleRandomness.CallOpts, requestId_)
}

// GetIntRange is a free data retrieval call binding the contract method 0xbe12f896.
//
// Solidity: function getIntRange(string requestId_) view returns(string requestId, uint256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCallerSession) GetIntRange(requestId_ string) (struct {
	RequestId  string
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetIntRange(&_DIAOracleRandomness.CallOpts, requestId_)
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
