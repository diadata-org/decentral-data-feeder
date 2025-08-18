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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBytes\",\"inputs\":[{\"name\":\"requestId_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"randomness\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIntArray\",\"inputs\":[{\"name\":\"requestId_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"randomInts\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIntRange\",\"inputs\":[{\"name\":\"requestId_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"randomInts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBytes\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"randomness\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIntArray\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"randomInts\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIntRange\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"randomInts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOracleUpdaterAddress\",\"inputs\":[{\"name\":\"newOracleUpdaterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BytesSet\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"round\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IntArraySet\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"round\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IntRangeSet\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"round\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdaterAddressChange\",\"inputs\":[{\"name\":\"newUpdater\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x6080604052348015600e575f5ffd5b503360035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611fa18061005c5f395ff3fe608060405234801561000f575f5ffd5b506004361061007b575f3560e01c806375db54521161005957806375db5452146100ea578063a6d414f41461011e578063a8501ae41461013a578063cb71c13a1461016d5761007b565b806357bc2ef31461007f5780636aa45efc146100b257806371c9a1d0146100ce575b5f5ffd5b61009960048036038101906100949190610ff6565b610189565b6040516100a994939291906111a3565b60405180910390f35b6100cc60048036038101906100c79190611255565b610403565b005b6100e860048036038101906100e3919061136c565b6104d5565b005b61010460048036038101906100ff9190610ff6565b610661565b604051610115959493929190611512565b60405180910390f35b610138600480360381019061013391906115cd565b610875565b005b610154600480360381019061014f9190610ff6565b610a70565b604051610164949392919061175b565b60405180910390f35b61018760048036038101906101829190611808565b610c6e565b005b5f60608060605f5f8681526020019081526020015f206004015f9054906101000a900460ff166101ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101e590611929565b60405180910390fd5b5f5f5f8781526020019081526020015f209050805f015481600101826002018360030182805480602002602001604051908101604052809291908181526020015f905b828210156102d9578382905f5260205f2001805461024e90611974565b80601f016020809104026020016040519081016040528092919081815260200182805461027a90611974565b80156102c55780601f1061029c576101008083540402835291602001916102c5565b820191905f5260205f20905b8154815290600101906020018083116102a857829003601f168201915b505050505081526020019060010190610231565b5050505092508180546102eb90611974565b80601f016020809104026020016040519081016040528092919081815260200182805461031790611974565b80156103625780601f1061033957610100808354040283529160200191610362565b820191905f5260205f20905b81548152906001019060200180831161034557829003601f168201915b5050505050915080805461037590611974565b80601f01602080910402602001604051908101604052809291908181526020018280546103a190611974565b80156103ec5780601f106103c3576101008083540402835291602001916103ec565b820191905f5260205f20905b8154815290600101906020018083116103cf57829003601f168201915b505050505090509450945094509450509193509193565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461045b575f5ffd5b8060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f816040516104ca91906119b3565b60405180910390a150565b6040518060a001604052808981526020018888906104f39190611ba5565b815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200183838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050508152602001600115158152505f5f8a81526020019081526020015f205f820151815f015560208201518160010190805190602001906105c5929190610e2f565b5060408201518160020190816105db9190611d59565b5060608201518160030190816105f19190611d59565b506080820151816004015f6101000a81548160ff0219169083151502179055509050508460070b7f528e74ef5afddaab004103264af5c5f61653f9366f735afdb9ac2f86f780a88e898686868660405161064f959493929190611e54565b60405180910390a25050505050505050565b5f60605f60608060025f8781526020019081526020015f206005015f9054906101000a900460ff166106c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106bf90611ee5565b60405180910390fd5b5f60025f8881526020019081526020015f209050805f015481600101826002015f9054906101000a900460070b83600301846004018380548060200260200160405190810160405280929190818152602001828054801561074657602002820191905f5260205f20905b815481526020019060010190808311610732575b5050505050935081805461075990611974565b80601f016020809104026020016040519081016040528092919081815260200182805461078590611974565b80156107d05780601f106107a7576101008083540402835291602001916107d0565b820191905f5260205f20905b8154815290600101906020018083116107b357829003601f168201915b505050505091508080546107e390611974565b80601f016020809104026020016040519081016040528092919081815260200182805461080f90611974565b801561085a5780601f106108315761010080835404028352916020019161085a565b820191905f5260205f20905b81548152906001019060200180831161083d57829003601f168201915b50505050509050955095509550955095505091939590929450565b6040518060c001604052808981526020018888808060200260200160405190810160405280939291908181526020018383602002808284375f81840152601f19601f8201169050808301925050505050505081526020018660070b815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200183838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020016001151581525060025f8a81526020019081526020015f205f820151815f015560208201518160010190805190602001906109a3929190610e86565b506040820151816002015f6101000a81548167ffffffffffffffff021916908360070b67ffffffffffffffff16021790555060608201518160030190816109ea9190611d59565b506080820151816004019081610a009190611d59565b5060a0820151816005015f6101000a81548160ff0219169083151502179055509050508460070b7f7cb19f70510748e0595cdd4e8923f152ab5eb90f39b2b7089c92fee973a3dbbf8986868686604051610a5e959493929190611e54565b60405180910390a25050505050505050565b5f606080606060015f8681526020019081526020015f206004015f9054906101000a900460ff16610ad6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610acd90611f4d565b60405180910390fd5b5f60015f8781526020019081526020015f209050805f015481600101826002018360030182805480602002602001604051908101604052809291908181526020018280548015610b4357602002820191905f5260205f20905b815481526020019060010190808311610b2f575b50505050509250818054610b5690611974565b80601f0160208091040260200160405190810160405280929190818152602001828054610b8290611974565b8015610bcd5780601f10610ba457610100808354040283529160200191610bcd565b820191905f5260205f20905b815481529060010190602001808311610bb057829003601f168201915b50505050509150808054610be090611974565b80601f0160208091040260200160405190810160405280929190818152602001828054610c0c90611974565b8015610c575780601f10610c2e57610100808354040283529160200191610c57565b820191905f5260205f20905b815481529060010190602001808311610c3a57829003601f168201915b505050505090509450945094509450509193509193565b6040518060a001604052808981526020018888808060200260200160405190810160405280939291908181526020018383602002808284375f81840152601f19601f82011690508083019250505050505050815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f82011690508083019250505050505050815260200183838080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f8201169050808301925050505050505081526020016001151581525060015f8a81526020019081526020015f205f820151815f01556020820151816001019080519060200190610d93929190610ed1565b506040820151816002019081610da99190611d59565b506060820151816003019081610dbf9190611d59565b506080820151816004015f6101000a81548160ff0219169083151502179055509050508460070b7f515771067561fd383bd7fb9cf8f798dce875f0cd006815f77f5cebf737cd6bbb8986868686604051610e1d959493929190611e54565b60405180910390a25050505050505050565b828054828255905f5260205f20908101928215610e75579160200282015b82811115610e74578251829081610e649190611d59565b5091602001919060010190610e4d565b5b509050610e829190610f1c565b5090565b828054828255905f5260205f20908101928215610ec0579160200282015b82811115610ebf578251825591602001919060010190610ea4565b5b509050610ecd9190610f3f565b5090565b828054828255905f5260205f20908101928215610f0b579160200282015b82811115610f0a578251825591602001919060010190610eef565b5b509050610f189190610f5a565b5090565b5b80821115610f3b575f8181610f329190610f75565b50600101610f1d565b5090565b5b80821115610f56575f815f905550600101610f40565b5090565b5b80821115610f71575f815f905550600101610f5b565b5090565b508054610f8190611974565b5f825580601f10610f925750610faf565b601f0160209004905f5260205f2090810190610fae9190610f5a565b5b50565b5f604051905090565b5f5ffd5b5f5ffd5b5f819050919050565b610fd581610fc3565b8114610fdf575f5ffd5b50565b5f81359050610ff081610fcc565b92915050565b5f6020828403121561100b5761100a610fbb565b5b5f61101884828501610fe2565b91505092915050565b61102a81610fc3565b82525050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f61109b82611059565b6110a58185611063565b93506110b5818560208601611073565b6110be81611081565b840191505092915050565b5f6110d48383611091565b905092915050565b5f602082019050919050565b5f6110f282611030565b6110fc818561103a565b93508360208202850161110e8561104a565b805f5b85811015611149578484038952815161112a85826110c9565b9450611135836110dc565b925060208a01995050600181019050611111565b50829750879550505050505092915050565b5f82825260208201905092915050565b5f61117582611059565b61117f818561115b565b935061118f818560208601611073565b61119881611081565b840191505092915050565b5f6080820190506111b65f830187611021565b81810360208301526111c881866110e8565b905081810360408301526111dc818561116b565b905081810360608301526111f0818461116b565b905095945050505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611224826111fb565b9050919050565b6112348161121a565b811461123e575f5ffd5b50565b5f8135905061124f8161122b565b92915050565b5f6020828403121561126a57611269610fbb565b5b5f61127784828501611241565b91505092915050565b5f5ffd5b5f5ffd5b5f5ffd5b5f5f83601f8401126112a1576112a0611280565b5b8235905067ffffffffffffffff8111156112be576112bd611284565b5b6020830191508360208202830111156112da576112d9611288565b5b9250929050565b5f8160070b9050919050565b6112f6816112e1565b8114611300575f5ffd5b50565b5f81359050611311816112ed565b92915050565b5f5f83601f84011261132c5761132b611280565b5b8235905067ffffffffffffffff81111561134957611348611284565b5b60208301915083600182028301111561136557611364611288565b5b9250929050565b5f5f5f5f5f5f5f5f60a0898b03121561138857611387610fbb565b5b5f6113958b828c01610fe2565b985050602089013567ffffffffffffffff8111156113b6576113b5610fbf565b5b6113c28b828c0161128c565b975097505060406113d58b828c01611303565b955050606089013567ffffffffffffffff8111156113f6576113f5610fbf565b5b6114028b828c01611317565b9450945050608089013567ffffffffffffffff81111561142557611424610fbf565b5b6114318b828c01611317565b92509250509295985092959890939650565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b5f819050919050565b61147e8161146c565b82525050565b5f61148f8383611475565b60208301905092915050565b5f602082019050919050565b5f6114b182611443565b6114bb818561144d565b93506114c68361145d565b805f5b838110156114f65781516114dd8882611484565b97506114e88361149b565b9250506001810190506114c9565b5085935050505092915050565b61150c816112e1565b82525050565b5f60a0820190506115255f830188611021565b818103602083015261153781876114a7565b90506115466040830186611503565b8181036060830152611558818561116b565b9050818103608083015261156c818461116b565b90509695505050505050565b5f5f83601f84011261158d5761158c611280565b5b8235905067ffffffffffffffff8111156115aa576115a9611284565b5b6020830191508360208202830111156115c6576115c5611288565b5b9250929050565b5f5f5f5f5f5f5f5f60a0898b0312156115e9576115e8610fbb565b5b5f6115f68b828c01610fe2565b985050602089013567ffffffffffffffff81111561161757611616610fbf565b5b6116238b828c01611578565b975097505060406116368b828c01611303565b955050606089013567ffffffffffffffff81111561165757611656610fbf565b5b6116638b828c01611317565b9450945050608089013567ffffffffffffffff81111561168657611685610fbf565b5b6116928b828c01611317565b92509250509295985092959890939650565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6116d681610fc3565b82525050565b5f6116e783836116cd565b60208301905092915050565b5f602082019050919050565b5f611709826116a4565b61171381856116ae565b935061171e836116be565b805f5b8381101561174e57815161173588826116dc565b9750611740836116f3565b925050600181019050611721565b5085935050505092915050565b5f60808201905061176e5f830187611021565b818103602083015261178081866116ff565b90508181036040830152611794818561116b565b905081810360608301526117a8818461116b565b905095945050505050565b5f5f83601f8401126117c8576117c7611280565b5b8235905067ffffffffffffffff8111156117e5576117e4611284565b5b60208301915083602082028301111561180157611800611288565b5b9250929050565b5f5f5f5f5f5f5f5f60a0898b03121561182457611823610fbb565b5b5f6118318b828c01610fe2565b985050602089013567ffffffffffffffff81111561185257611851610fbf565b5b61185e8b828c016117b3565b975097505060406118718b828c01611303565b955050606089013567ffffffffffffffff81111561189257611891610fbf565b5b61189e8b828c01611317565b9450945050608089013567ffffffffffffffff8111156118c1576118c0610fbf565b5b6118cd8b828c01611317565b92509250509295985092959890939650565b7f4e6f204279746573456e74727920666f72207468697320726f756e64000000005f82015250565b5f611913601c8361115b565b915061191e826118df565b602082019050919050565b5f6020820190508181035f83015261194081611907565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061198b57607f821691505b60208210810361199e5761199d611947565b5b50919050565b6119ad8161121a565b82525050565b5f6020820190506119c65f8301846119a4565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b611a0282611081565b810181811067ffffffffffffffff82111715611a2157611a206119cc565b5b80604052505050565b5f611a33610fb2565b9050611a3f82826119f9565b919050565b5f67ffffffffffffffff821115611a5e57611a5d6119cc565b5b602082029050602081019050919050565b5f5ffd5b5f67ffffffffffffffff821115611a8d57611a8c6119cc565b5b611a9682611081565b9050602081019050919050565b828183375f83830152505050565b5f611ac3611abe84611a73565b611a2a565b905082815260208101848484011115611adf57611ade611a6f565b5b611aea848285611aa3565b509392505050565b5f82601f830112611b0657611b05611280565b5b8135611b16848260208601611ab1565b91505092915050565b5f611b31611b2c84611a44565b611a2a565b90508083825260208201905060208402830185811115611b5457611b53611288565b5b835b81811015611b9b57803567ffffffffffffffff811115611b7957611b78611280565b5b808601611b868982611af2565b85526020850194505050602081019050611b56565b5050509392505050565b5f611bb1368484611b1f565b905092915050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302611c157fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611bda565b611c1f8683611bda565b95508019841693508086168417925050509392505050565b5f819050919050565b5f611c5a611c55611c5084610fc3565b611c37565b610fc3565b9050919050565b5f819050919050565b611c7383611c40565b611c87611c7f82611c61565b848454611be6565b825550505050565b5f5f905090565b611c9e611c8f565b611ca9818484611c6a565b505050565b5b81811015611ccc57611cc15f82611c96565b600181019050611caf565b5050565b601f821115611d1157611ce281611bb9565b611ceb84611bcb565b81016020851015611cfa578190505b611d0e611d0685611bcb565b830182611cae565b50505b505050565b5f82821c905092915050565b5f611d315f1984600802611d16565b1980831691505092915050565b5f611d498383611d22565b9150826002028217905092915050565b611d6282611059565b67ffffffffffffffff811115611d7b57611d7a6119cc565b5b611d858254611974565b611d90828285611cd0565b5f60209050601f831160018114611dc1575f8415611daf578287015190505b611db98582611d3e565b865550611e20565b601f198416611dcf86611bb9565b5f5b82811015611df657848901518255600182019150602085019450602081019050611dd1565b86831015611e135784890151611e0f601f891682611d22565b8355505b6001600288020188555050505b505050505050565b5f611e33838561115b565b9350611e40838584611aa3565b611e4983611081565b840190509392505050565b5f606082019050611e675f830188611021565b8181036020830152611e7a818688611e28565b90508181036040830152611e8f818486611e28565b90509695505050505050565b7f4e6f20496e744172726179456e74727920666f72207468697320726f756e64005f82015250565b5f611ecf601f8361115b565b9150611eda82611e9b565b602082019050919050565b5f6020820190508181035f830152611efc81611ec3565b9050919050565b7f4e6f20496e7452616e6765456e74727920666f72207468697320726f756e64005f82015250565b5f611f37601f8361115b565b9150611f4282611f03565b602082019050919050565b5f6020820190508181035f830152611f6481611f2b565b905091905056fea264697066735822122082364d4def79e29c6419335f56c059ffc044ad2fb95091de6a6dcfd940113fbf64736f6c634300081e0033",
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

// GetBytes is a free data retrieval call binding the contract method 0x57bc2ef3.
//
// Solidity: function getBytes(uint256 requestId_) view returns(uint256 requestId, string[] randomness, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCaller) GetBytes(opts *bind.CallOpts, requestId_ *big.Int) (struct {
	RequestId  *big.Int
	Randomness []string
	Seed       string
	Signature  string
}, error) {
	var out []interface{}
	err := _DIAOracleRandomness.contract.Call(opts, &out, "getBytes", requestId_)

	outstruct := new(struct {
		RequestId  *big.Int
		Randomness []string
		Seed       string
		Signature  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Randomness = *abi.ConvertType(out[1], new([]string)).(*[]string)
	outstruct.Seed = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Signature = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// GetBytes is a free data retrieval call binding the contract method 0x57bc2ef3.
//
// Solidity: function getBytes(uint256 requestId_) view returns(uint256 requestId, string[] randomness, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessSession) GetBytes(requestId_ *big.Int) (struct {
	RequestId  *big.Int
	Randomness []string
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetBytes(&_DIAOracleRandomness.CallOpts, requestId_)
}

// GetBytes is a free data retrieval call binding the contract method 0x57bc2ef3.
//
// Solidity: function getBytes(uint256 requestId_) view returns(uint256 requestId, string[] randomness, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCallerSession) GetBytes(requestId_ *big.Int) (struct {
	RequestId  *big.Int
	Randomness []string
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetBytes(&_DIAOracleRandomness.CallOpts, requestId_)
}

// GetIntArray is a free data retrieval call binding the contract method 0x75db5452.
//
// Solidity: function getIntArray(uint256 requestId_) view returns(uint256 requestId, int256[] randomInts, int64 round, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCaller) GetIntArray(opts *bind.CallOpts, requestId_ *big.Int) (struct {
	RequestId  *big.Int
	RandomInts []*big.Int
	Round      int64
	Seed       string
	Signature  string
}, error) {
	var out []interface{}
	err := _DIAOracleRandomness.contract.Call(opts, &out, "getIntArray", requestId_)

	outstruct := new(struct {
		RequestId  *big.Int
		RandomInts []*big.Int
		Round      int64
		Seed       string
		Signature  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RandomInts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Round = *abi.ConvertType(out[2], new(int64)).(*int64)
	outstruct.Seed = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Signature = *abi.ConvertType(out[4], new(string)).(*string)

	return *outstruct, err

}

// GetIntArray is a free data retrieval call binding the contract method 0x75db5452.
//
// Solidity: function getIntArray(uint256 requestId_) view returns(uint256 requestId, int256[] randomInts, int64 round, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessSession) GetIntArray(requestId_ *big.Int) (struct {
	RequestId  *big.Int
	RandomInts []*big.Int
	Round      int64
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetIntArray(&_DIAOracleRandomness.CallOpts, requestId_)
}

// GetIntArray is a free data retrieval call binding the contract method 0x75db5452.
//
// Solidity: function getIntArray(uint256 requestId_) view returns(uint256 requestId, int256[] randomInts, int64 round, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCallerSession) GetIntArray(requestId_ *big.Int) (struct {
	RequestId  *big.Int
	RandomInts []*big.Int
	Round      int64
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetIntArray(&_DIAOracleRandomness.CallOpts, requestId_)
}

// GetIntRange is a free data retrieval call binding the contract method 0xa8501ae4.
//
// Solidity: function getIntRange(uint256 requestId_) view returns(uint256 requestId, uint256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCaller) GetIntRange(opts *bind.CallOpts, requestId_ *big.Int) (struct {
	RequestId  *big.Int
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	var out []interface{}
	err := _DIAOracleRandomness.contract.Call(opts, &out, "getIntRange", requestId_)

	outstruct := new(struct {
		RequestId  *big.Int
		RandomInts []*big.Int
		Seed       string
		Signature  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RequestId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RandomInts = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.Seed = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Signature = *abi.ConvertType(out[3], new(string)).(*string)

	return *outstruct, err

}

// GetIntRange is a free data retrieval call binding the contract method 0xa8501ae4.
//
// Solidity: function getIntRange(uint256 requestId_) view returns(uint256 requestId, uint256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessSession) GetIntRange(requestId_ *big.Int) (struct {
	RequestId  *big.Int
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetIntRange(&_DIAOracleRandomness.CallOpts, requestId_)
}

// GetIntRange is a free data retrieval call binding the contract method 0xa8501ae4.
//
// Solidity: function getIntRange(uint256 requestId_) view returns(uint256 requestId, uint256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCallerSession) GetIntRange(requestId_ *big.Int) (struct {
	RequestId  *big.Int
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetIntRange(&_DIAOracleRandomness.CallOpts, requestId_)
}

// SetBytes is a paid mutator transaction binding the contract method 0x71c9a1d0.
//
// Solidity: function setBytes(uint256 requestId, string[] randomness, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactor) SetBytes(opts *bind.TransactOpts, requestId *big.Int, randomness []string, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.contract.Transact(opts, "setBytes", requestId, randomness, round, seed, signature)
}

// SetBytes is a paid mutator transaction binding the contract method 0x71c9a1d0.
//
// Solidity: function setBytes(uint256 requestId, string[] randomness, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessSession) SetBytes(requestId *big.Int, randomness []string, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.SetBytes(&_DIAOracleRandomness.TransactOpts, requestId, randomness, round, seed, signature)
}

// SetBytes is a paid mutator transaction binding the contract method 0x71c9a1d0.
//
// Solidity: function setBytes(uint256 requestId, string[] randomness, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactorSession) SetBytes(requestId *big.Int, randomness []string, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.SetBytes(&_DIAOracleRandomness.TransactOpts, requestId, randomness, round, seed, signature)
}

// SetIntArray is a paid mutator transaction binding the contract method 0xa6d414f4.
//
// Solidity: function setIntArray(uint256 requestId, int256[] randomInts, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactor) SetIntArray(opts *bind.TransactOpts, requestId *big.Int, randomInts []*big.Int, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.contract.Transact(opts, "setIntArray", requestId, randomInts, round, seed, signature)
}

// SetIntArray is a paid mutator transaction binding the contract method 0xa6d414f4.
//
// Solidity: function setIntArray(uint256 requestId, int256[] randomInts, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessSession) SetIntArray(requestId *big.Int, randomInts []*big.Int, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.SetIntArray(&_DIAOracleRandomness.TransactOpts, requestId, randomInts, round, seed, signature)
}

// SetIntArray is a paid mutator transaction binding the contract method 0xa6d414f4.
//
// Solidity: function setIntArray(uint256 requestId, int256[] randomInts, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactorSession) SetIntArray(requestId *big.Int, randomInts []*big.Int, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.SetIntArray(&_DIAOracleRandomness.TransactOpts, requestId, randomInts, round, seed, signature)
}

// SetIntRange is a paid mutator transaction binding the contract method 0xcb71c13a.
//
// Solidity: function setIntRange(uint256 requestId, uint256[] randomInts, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactor) SetIntRange(opts *bind.TransactOpts, requestId *big.Int, randomInts []*big.Int, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.contract.Transact(opts, "setIntRange", requestId, randomInts, round, seed, signature)
}

// SetIntRange is a paid mutator transaction binding the contract method 0xcb71c13a.
//
// Solidity: function setIntRange(uint256 requestId, uint256[] randomInts, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessSession) SetIntRange(requestId *big.Int, randomInts []*big.Int, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.SetIntRange(&_DIAOracleRandomness.TransactOpts, requestId, randomInts, round, seed, signature)
}

// SetIntRange is a paid mutator transaction binding the contract method 0xcb71c13a.
//
// Solidity: function setIntRange(uint256 requestId, uint256[] randomInts, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactorSession) SetIntRange(requestId *big.Int, randomInts []*big.Int, round int64, seed string, signature string) (*types.Transaction, error) {
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
	RequestId *big.Int
	Round     *big.Int
	Seed      string
	Signature string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBytesSet is a free log retrieval operation binding the contract event 0x528e74ef5afddaab004103264af5c5f61653f9366f735afdb9ac2f86f780a88e.
//
// Solidity: event BytesSet(uint256 requestId, int256 indexed round, string seed, string signature)
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

// WatchBytesSet is a free log subscription operation binding the contract event 0x528e74ef5afddaab004103264af5c5f61653f9366f735afdb9ac2f86f780a88e.
//
// Solidity: event BytesSet(uint256 requestId, int256 indexed round, string seed, string signature)
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

// ParseBytesSet is a log parse operation binding the contract event 0x528e74ef5afddaab004103264af5c5f61653f9366f735afdb9ac2f86f780a88e.
//
// Solidity: event BytesSet(uint256 requestId, int256 indexed round, string seed, string signature)
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
	RequestId *big.Int
	Round     *big.Int
	Seed      string
	Signature string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIntArraySet is a free log retrieval operation binding the contract event 0x7cb19f70510748e0595cdd4e8923f152ab5eb90f39b2b7089c92fee973a3dbbf.
//
// Solidity: event IntArraySet(uint256 requestId, int256 indexed round, string seed, string signature)
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

// WatchIntArraySet is a free log subscription operation binding the contract event 0x7cb19f70510748e0595cdd4e8923f152ab5eb90f39b2b7089c92fee973a3dbbf.
//
// Solidity: event IntArraySet(uint256 requestId, int256 indexed round, string seed, string signature)
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

// ParseIntArraySet is a log parse operation binding the contract event 0x7cb19f70510748e0595cdd4e8923f152ab5eb90f39b2b7089c92fee973a3dbbf.
//
// Solidity: event IntArraySet(uint256 requestId, int256 indexed round, string seed, string signature)
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
	RequestId *big.Int
	Round     *big.Int
	Seed      string
	Signature string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIntRangeSet is a free log retrieval operation binding the contract event 0x515771067561fd383bd7fb9cf8f798dce875f0cd006815f77f5cebf737cd6bbb.
//
// Solidity: event IntRangeSet(uint256 requestId, int256 indexed round, string seed, string signature)
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

// WatchIntRangeSet is a free log subscription operation binding the contract event 0x515771067561fd383bd7fb9cf8f798dce875f0cd006815f77f5cebf737cd6bbb.
//
// Solidity: event IntRangeSet(uint256 requestId, int256 indexed round, string seed, string signature)
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

// ParseIntRangeSet is a log parse operation binding the contract event 0x515771067561fd383bd7fb9cf8f798dce875f0cd006815f77f5cebf737cd6bbb.
//
// Solidity: event IntRangeSet(uint256 requestId, int256 indexed round, string seed, string signature)
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
