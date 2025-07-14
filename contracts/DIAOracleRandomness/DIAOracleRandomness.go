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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBytes\",\"inputs\":[{\"name\":\"requestId_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"randomness\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIntArray\",\"inputs\":[{\"name\":\"requestId_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"randomInts\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIntRange\",\"inputs\":[{\"name\":\"requestId_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"randomInts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBytes\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"randomness\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIntArray\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"randomInts\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIntRange\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"randomInts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOracleUpdaterAddress\",\"inputs\":[{\"name\":\"newOracleUpdaterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BytesSet\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"round\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IntArraySet\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"round\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IntRangeSet\",\"inputs\":[{\"name\":\"requestId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"round\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdaterAddressChange\",\"inputs\":[{\"name\":\"newUpdater\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b5033600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061201e806100616000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806375db54521161005b57806375db5452146100ed578063a6d414f414610120578063a8501ae41461013c578063cb71c13a1461016f5761007d565b806357bc2ef3146100825780636aa45efc146100b557806371c9a1d0146100d1575b600080fd5b61009c60048036038101906100979190611003565b61018b565b6040516100ac94939291906111db565b60405180910390f35b6100cf60048036038101906100ca9190611293565b610416565b005b6100eb60048036038101906100e691906113b4565b6104eb565b005b61010760048036038101906101029190611003565b61067e565b6040516101179493929190611558565b60405180910390f35b61013a60048036038101906101359190611608565b61088a565b005b61015660048036038101906101519190611003565b610a53565b60405161016694939291906117a2565b60405180910390f35b61018960048036038101906101849190611852565b610c5f565b005b6000606080606060008086815260200190815260200160002060040160009054906101000a900460ff166101f4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101eb9061197a565b60405180910390fd5b60008060008781526020019081526020016000209050806000015481600101826002018360030182805480602002602001604051908101604052809291908181526020016000905b828210156102e857838290600052602060002001805461025b906119c9565b80601f0160208091040260200160405190810160405280929190818152602001828054610287906119c9565b80156102d45780601f106102a9576101008083540402835291602001916102d4565b820191906000526020600020905b8154815290600101906020018083116102b757829003601f168201915b50505050508152602001906001019061023c565b5050505092508180546102fa906119c9565b80601f0160208091040260200160405190810160405280929190818152602001828054610326906119c9565b80156103735780601f1061034857610100808354040283529160200191610373565b820191906000526020600020905b81548152906001019060200180831161035657829003601f168201915b50505050509150808054610386906119c9565b80601f01602080910402602001604051908101604052809291908181526020018280546103b2906119c9565b80156103ff5780601f106103d4576101008083540402835291602001916103ff565b820191906000526020600020905b8154815290600101906020018083116103e257829003601f168201915b505050505090509450945094509450509193509193565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461047057600080fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f816040516104e09190611a09565b60405180910390a150565b6040518060a001604052808981526020018888906105099190611c07565b815260200185858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200183838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152602001600115158152506000808a81526020019081526020016000206000820151816000015560208201518160010190805190602001906105e1929190610e28565b5060408201518160020190816105f79190611dc8565b50606082015181600301908161060d9190611dc8565b5060808201518160040160006101000a81548160ff0219169083151502179055509050508460070b7f528e74ef5afddaab004103264af5c5f61653f9366f735afdb9ac2f86f780a88e898686868660405161066c959493929190611ec7565b60405180910390a25050505050505050565b600060608060606002600086815260200190815260200160002060040160009054906101000a900460ff166106e8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106df90611f5c565b60405180910390fd5b600060026000878152602001908152602001600020905080600001548160010182600201836003018280548060200260200160405190810160405280929190818152602001828054801561075b57602002820191906000526020600020905b815481526020019060010190808311610747575b5050505050925081805461076e906119c9565b80601f016020809104026020016040519081016040528092919081815260200182805461079a906119c9565b80156107e75780601f106107bc576101008083540402835291602001916107e7565b820191906000526020600020905b8154815290600101906020018083116107ca57829003601f168201915b505050505091508080546107fa906119c9565b80601f0160208091040260200160405190810160405280929190818152602001828054610826906119c9565b80156108735780601f1061084857610100808354040283529160200191610873565b820191906000526020600020905b81548152906001019060200180831161085657829003601f168201915b505050505090509450945094509450509193509193565b6040518060a00160405280898152602001888880806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050815260200185858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200183838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200160011515815250600260008a81526020019081526020016000206000820151816000015560208201518160010190805190602001906109b6929190610e81565b5060408201518160020190816109cc9190611dc8565b5060608201518160030190816109e29190611dc8565b5060808201518160040160006101000a81548160ff0219169083151502179055509050508460070b7f7cb19f70510748e0595cdd4e8923f152ab5eb90f39b2b7089c92fee973a3dbbf8986868686604051610a41959493929190611ec7565b60405180910390a25050505050505050565b600060608060606001600086815260200190815260200160002060040160009054906101000a900460ff16610abd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ab490611fc8565b60405180910390fd5b6000600160008781526020019081526020016000209050806000015481600101826002018360030182805480602002602001604051908101604052809291908181526020018280548015610b3057602002820191906000526020600020905b815481526020019060010190808311610b1c575b50505050509250818054610b43906119c9565b80601f0160208091040260200160405190810160405280929190818152602001828054610b6f906119c9565b8015610bbc5780601f10610b9157610100808354040283529160200191610bbc565b820191906000526020600020905b815481529060010190602001808311610b9f57829003601f168201915b50505050509150808054610bcf906119c9565b80601f0160208091040260200160405190810160405280929190818152602001828054610bfb906119c9565b8015610c485780601f10610c1d57610100808354040283529160200191610c48565b820191906000526020600020905b815481529060010190602001808311610c2b57829003601f168201915b505050505090509450945094509450509193509193565b6040518060a00160405280898152602001888880806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050815260200185858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200183838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200160011515815250600160008a8152602001908152602001600020600082015181600001556020820151816001019080519060200190610d8b929190610ece565b506040820151816002019081610da19190611dc8565b506060820151816003019081610db79190611dc8565b5060808201518160040160006101000a81548160ff0219169083151502179055509050508460070b7f515771067561fd383bd7fb9cf8f798dce875f0cd006815f77f5cebf737cd6bbb8986868686604051610e16959493929190611ec7565b60405180910390a25050505050505050565b828054828255906000526020600020908101928215610e70579160200282015b82811115610e6f578251829081610e5f9190611dc8565b5091602001919060010190610e48565b5b509050610e7d9190610f1b565b5090565b828054828255906000526020600020908101928215610ebd579160200282015b82811115610ebc578251825591602001919060010190610ea1565b5b509050610eca9190610f3f565b5090565b828054828255906000526020600020908101928215610f0a579160200282015b82811115610f09578251825591602001919060010190610eee565b5b509050610f179190610f5c565b5090565b5b80821115610f3b5760008181610f329190610f79565b50600101610f1c565b5090565b5b80821115610f58576000816000905550600101610f40565b5090565b5b80821115610f75576000816000905550600101610f5d565b5090565b508054610f85906119c9565b6000825580601f10610f975750610fb6565b601f016020900490600052602060002090810190610fb59190610f5c565b5b50565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610fe081610fcd565b8114610feb57600080fd5b50565b600081359050610ffd81610fd7565b92915050565b60006020828403121561101957611018610fc3565b5b600061102784828501610fee565b91505092915050565b61103981610fcd565b82525050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600081519050919050565b600082825260208201905092915050565b60005b838110156110a557808201518184015260208101905061108a565b60008484015250505050565b6000601f19601f8301169050919050565b60006110cd8261106b565b6110d78185611076565b93506110e7818560208601611087565b6110f0816110b1565b840191505092915050565b600061110783836110c2565b905092915050565b6000602082019050919050565b60006111278261103f565b611131818561104a565b9350836020820285016111438561105b565b8060005b8581101561117f578484038952815161116085826110fb565b945061116b8361110f565b925060208a01995050600181019050611147565b50829750879550505050505092915050565b600082825260208201905092915050565b60006111ad8261106b565b6111b78185611191565b93506111c7818560208601611087565b6111d0816110b1565b840191505092915050565b60006080820190506111f06000830187611030565b8181036020830152611202818661111c565b9050818103604083015261121681856111a2565b9050818103606083015261122a81846111a2565b905095945050505050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061126082611235565b9050919050565b61127081611255565b811461127b57600080fd5b50565b60008135905061128d81611267565b92915050565b6000602082840312156112a9576112a8610fc3565b5b60006112b78482850161127e565b91505092915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126112e5576112e46112c0565b5b8235905067ffffffffffffffff811115611302576113016112c5565b5b60208301915083602082028301111561131e5761131d6112ca565b5b9250929050565b60008160070b9050919050565b61133b81611325565b811461134657600080fd5b50565b60008135905061135881611332565b92915050565b60008083601f840112611374576113736112c0565b5b8235905067ffffffffffffffff811115611391576113906112c5565b5b6020830191508360018202830111156113ad576113ac6112ca565b5b9250929050565b60008060008060008060008060a0898b0312156113d4576113d3610fc3565b5b60006113e28b828c01610fee565b985050602089013567ffffffffffffffff81111561140357611402610fc8565b5b61140f8b828c016112cf565b975097505060406114228b828c01611349565b955050606089013567ffffffffffffffff81111561144357611442610fc8565b5b61144f8b828c0161135e565b9450945050608089013567ffffffffffffffff81111561147257611471610fc8565b5b61147e8b828c0161135e565b92509250509295985092959890939650565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b6114cf816114bc565b82525050565b60006114e183836114c6565b60208301905092915050565b6000602082019050919050565b600061150582611490565b61150f818561149b565b935061151a836114ac565b8060005b8381101561154b57815161153288826114d5565b975061153d836114ed565b92505060018101905061151e565b5085935050505092915050565b600060808201905061156d6000830187611030565b818103602083015261157f81866114fa565b9050818103604083015261159381856111a2565b905081810360608301526115a781846111a2565b905095945050505050565b60008083601f8401126115c8576115c76112c0565b5b8235905067ffffffffffffffff8111156115e5576115e46112c5565b5b602083019150836020820283011115611601576116006112ca565b5b9250929050565b60008060008060008060008060a0898b03121561162857611627610fc3565b5b60006116368b828c01610fee565b985050602089013567ffffffffffffffff81111561165757611656610fc8565b5b6116638b828c016115b2565b975097505060406116768b828c01611349565b955050606089013567ffffffffffffffff81111561169757611696610fc8565b5b6116a38b828c0161135e565b9450945050608089013567ffffffffffffffff8111156116c6576116c5610fc8565b5b6116d28b828c0161135e565b92509250509295985092959890939650565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61171981610fcd565b82525050565b600061172b8383611710565b60208301905092915050565b6000602082019050919050565b600061174f826116e4565b61175981856116ef565b935061176483611700565b8060005b8381101561179557815161177c888261171f565b975061178783611737565b925050600181019050611768565b5085935050505092915050565b60006080820190506117b76000830187611030565b81810360208301526117c98186611744565b905081810360408301526117dd81856111a2565b905081810360608301526117f181846111a2565b905095945050505050565b60008083601f840112611812576118116112c0565b5b8235905067ffffffffffffffff81111561182f5761182e6112c5565b5b60208301915083602082028301111561184b5761184a6112ca565b5b9250929050565b60008060008060008060008060a0898b03121561187257611871610fc3565b5b60006118808b828c01610fee565b985050602089013567ffffffffffffffff8111156118a1576118a0610fc8565b5b6118ad8b828c016117fc565b975097505060406118c08b828c01611349565b955050606089013567ffffffffffffffff8111156118e1576118e0610fc8565b5b6118ed8b828c0161135e565b9450945050608089013567ffffffffffffffff8111156119105761190f610fc8565b5b61191c8b828c0161135e565b92509250509295985092959890939650565b7f4e6f204279746573456e74727920666f72207468697320726f756e6400000000600082015250565b6000611964601c83611191565b915061196f8261192e565b602082019050919050565b6000602082019050818103600083015261199381611957565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806119e157607f821691505b6020821081036119f4576119f361199a565b5b50919050565b611a0381611255565b82525050565b6000602082019050611a1e60008301846119fa565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b611a5c826110b1565b810181811067ffffffffffffffff82111715611a7b57611a7a611a24565b5b80604052505050565b6000611a8e610fb9565b9050611a9a8282611a53565b919050565b600067ffffffffffffffff821115611aba57611ab9611a24565b5b602082029050602081019050919050565b600080fd5b600067ffffffffffffffff821115611aeb57611aea611a24565b5b611af4826110b1565b9050602081019050919050565b82818337600083830152505050565b6000611b23611b1e84611ad0565b611a84565b905082815260208101848484011115611b3f57611b3e611acb565b5b611b4a848285611b01565b509392505050565b600082601f830112611b6757611b666112c0565b5b8135611b77848260208601611b10565b91505092915050565b6000611b93611b8e84611a9f565b611a84565b90508083825260208201905060208402830185811115611bb657611bb56112ca565b5b835b81811015611bfd57803567ffffffffffffffff811115611bdb57611bda6112c0565b5b808601611be88982611b52565b85526020850194505050602081019050611bb8565b5050509392505050565b6000611c14368484611b80565b905092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302611c7e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611c41565b611c888683611c41565b95508019841693508086168417925050509392505050565b6000819050919050565b6000611cc5611cc0611cbb84610fcd565b611ca0565b610fcd565b9050919050565b6000819050919050565b611cdf83611caa565b611cf3611ceb82611ccc565b848454611c4e565b825550505050565b600090565b611d08611cfb565b611d13818484611cd6565b505050565b5b81811015611d3757611d2c600082611d00565b600181019050611d19565b5050565b601f821115611d7c57611d4d81611c1c565b611d5684611c31565b81016020851015611d65578190505b611d79611d7185611c31565b830182611d18565b50505b505050565b600082821c905092915050565b6000611d9f60001984600802611d81565b1980831691505092915050565b6000611db88383611d8e565b9150826002028217905092915050565b611dd18261106b565b67ffffffffffffffff811115611dea57611de9611a24565b5b611df482546119c9565b611dff828285611d3b565b600060209050601f831160018114611e325760008415611e20578287015190505b611e2a8582611dac565b865550611e92565b601f198416611e4086611c1c565b60005b82811015611e6857848901518255600182019150602085019450602081019050611e43565b86831015611e855784890151611e81601f891682611d8e565b8355505b6001600288020188555050505b505050505050565b6000611ea68385611191565b9350611eb3838584611b01565b611ebc836110b1565b840190509392505050565b6000606082019050611edc6000830188611030565b8181036020830152611eef818688611e9a565b90508181036040830152611f04818486611e9a565b90509695505050505050565b7f4e6f20496e744172726179456e74727920666f72207468697320726f756e6400600082015250565b6000611f46601f83611191565b9150611f5182611f10565b602082019050919050565b60006020820190508181036000830152611f7581611f39565b9050919050565b7f4e6f20496e7452616e6765456e74727920666f72207468697320726f756e6400600082015250565b6000611fb2601f83611191565b9150611fbd82611f7c565b602082019050919050565b60006020820190508181036000830152611fe181611fa5565b905091905056fea2646970667358221220a3df28adc2fbf2feffcd8472ad10c8ff903b55e23340659f138c06b3b801752364736f6c63430008130033",
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
// Solidity: function getIntArray(uint256 requestId_) view returns(uint256 requestId, int256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCaller) GetIntArray(opts *bind.CallOpts, requestId_ *big.Int) (struct {
	RequestId  *big.Int
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	var out []interface{}
	err := _DIAOracleRandomness.contract.Call(opts, &out, "getIntArray", requestId_)

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

// GetIntArray is a free data retrieval call binding the contract method 0x75db5452.
//
// Solidity: function getIntArray(uint256 requestId_) view returns(uint256 requestId, int256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessSession) GetIntArray(requestId_ *big.Int) (struct {
	RequestId  *big.Int
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetIntArray(&_DIAOracleRandomness.CallOpts, requestId_)
}

// GetIntArray is a free data retrieval call binding the contract method 0x75db5452.
//
// Solidity: function getIntArray(uint256 requestId_) view returns(uint256 requestId, int256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCallerSession) GetIntArray(requestId_ *big.Int) (struct {
	RequestId  *big.Int
	RandomInts []*big.Int
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
