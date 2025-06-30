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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getBytes\",\"inputs\":[{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"}],\"outputs\":[{\"name\":\"randomness\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIntArray\",\"inputs\":[{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"}],\"outputs\":[{\"name\":\"randomInts\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getIntRange\",\"inputs\":[{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"}],\"outputs\":[{\"name\":\"randomInts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBytes\",\"inputs\":[{\"name\":\"randomness\",\"type\":\"string[]\",\"internalType\":\"string[]\"},{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIntArray\",\"inputs\":[{\"name\":\"randomInts\",\"type\":\"int256[]\",\"internalType\":\"int256[]\"},{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIntRange\",\"inputs\":[{\"name\":\"randomInts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"round\",\"type\":\"int64\",\"internalType\":\"int64\"},{\"name\":\"seed\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"updateOracleUpdaterAddress\",\"inputs\":[{\"name\":\"newOracleUpdaterAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"BytesSet\",\"inputs\":[{\"name\":\"round\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IntArraySet\",\"inputs\":[{\"name\":\"round\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"IntRangeSet\",\"inputs\":[{\"name\":\"round\",\"type\":\"int256\",\"indexed\":true,\"internalType\":\"int256\"},{\"name\":\"seed\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"signature\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UpdaterAddressChange\",\"inputs\":[{\"name\":\"newUpdater\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561001057600080fd5b5033600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611f36806100616000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80636aa45efc1161005b5780636aa45efc14610102578063809fcd511461011e578063d9082a9b1461013a578063f8f0838e1461016c5761007d565b806345459cd01461008257806353759265146100b45780635d52ea10146100e6575b600080fd5b61009c60048036038101906100979190610fca565b610188565b6040516100ab9392919061114f565b60405180910390f35b6100ce60048036038101906100c99190610fca565b610391565b6040516100dd939291906112a7565b60405180910390f35b61010060048036038101906100fb91906113ae565b610619565b005b61011c600480360381019061011791906114d5565b61079c565b005b61013860048036038101906101339190611558565b610871565b005b610154600480360381019061014f9190610fca565b610a2a565b604051610163939291906116e9565b60405180910390f35b6101866004803603810190610181919061178b565b610c33565b005b6060806060600260008560070b815260200190815260200160002060030160009054906101000a900460ff166101f3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ea906118a0565b60405180910390fd5b6000600260008660070b815260200190815260200160002090508060000181600101826002018280548060200260200160405190810160405280929190818152602001828054801561026457602002820191906000526020600020905b815481526020019060010190808311610250575b50505050509250818054610277906118ef565b80601f01602080910402602001604051908101604052809291908181526020018280546102a3906118ef565b80156102f05780601f106102c5576101008083540402835291602001916102f0565b820191906000526020600020905b8154815290600101906020018083116102d357829003601f168201915b50505050509150808054610303906118ef565b80601f016020809104026020016040519081016040528092919081815260200182805461032f906118ef565b801561037c5780601f106103515761010080835404028352916020019161037c565b820191906000526020600020905b81548152906001019060200180831161035f57829003601f168201915b50505050509050935093509350509193909250565b60608060606000808560070b815260200190815260200160002060030160009054906101000a900460ff166103fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f29061196c565b60405180910390fd5b60008060008660070b8152602001908152602001600020905080600001816001018260020182805480602002602001604051908101604052809291908181526020016000905b828210156104ed578382906000526020600020018054610460906118ef565b80601f016020809104026020016040519081016040528092919081815260200182805461048c906118ef565b80156104d95780601f106104ae576101008083540402835291602001916104d9565b820191906000526020600020905b8154815290600101906020018083116104bc57829003601f168201915b505050505081526020019060010190610441565b5050505092508180546104ff906118ef565b80601f016020809104026020016040519081016040528092919081815260200182805461052b906118ef565b80156105785780601f1061054d57610100808354040283529160200191610578565b820191906000526020600020905b81548152906001019060200180831161055b57829003601f168201915b5050505050915080805461058b906118ef565b80601f01602080910402602001604051908101604052809291908181526020018280546105b7906118ef565b80156106045780601f106105d957610100808354040283529160200191610604565b820191906000526020600020905b8154815290600101906020018083116105e757829003601f168201915b50505050509050935093509350509193909250565b60405180608001604052808888906106319190611b6f565b815260200185858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200183838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508152602001600115158152506000808760070b81526020019081526020016000206000820151816000019080519060200190610702929190610dec565b5060208201518160010190816107189190611d30565b50604082015181600201908161072e9190611d30565b5060608201518160030160006101000a81548160ff0219169083151502179055509050508460070b7ff08ab236a050614700c1f2c02caaf5de6f6ee744d522ab71eb23633fac1909c28585858560405161078b9493929190611e2f565b60405180910390a250505050505050565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107f657600080fd5b80600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f121e958a4cadf7f8dadefa22cc019700365240223668418faebed197da07089f816040516108669190611e79565b60405180910390a150565b6040518060800160405280888880806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050815260200185858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200183838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200160011515815250600160008760070b81526020019081526020016000206000820151816000019080519060200190610990929190610e45565b5060208201518160010190816109a69190611d30565b5060408201518160020190816109bc9190611d30565b5060608201518160030160006101000a81548160ff0219169083151502179055509050508460070b7fc43db9672cb481f936463442b4cefa9165d1388dfa3804b7585f84bb05f3210785858585604051610a199493929190611e2f565b60405180910390a250505050505050565b6060806060600160008560070b815260200190815260200160002060030160009054906101000a900460ff16610a95576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8c90611ee0565b60405180910390fd5b6000600160008660070b8152602001908152602001600020905080600001816001018260020182805480602002602001604051908101604052809291908181526020018280548015610b0657602002820191906000526020600020905b815481526020019060010190808311610af2575b50505050509250818054610b19906118ef565b80601f0160208091040260200160405190810160405280929190818152602001828054610b45906118ef565b8015610b925780601f10610b6757610100808354040283529160200191610b92565b820191906000526020600020905b815481529060010190602001808311610b7557829003601f168201915b50505050509150808054610ba5906118ef565b80601f0160208091040260200160405190810160405280929190818152602001828054610bd1906118ef565b8015610c1e5780601f10610bf357610100808354040283529160200191610c1e565b820191906000526020600020905b815481529060010190602001808311610c0157829003601f168201915b50505050509050935093509350509193909250565b6040518060800160405280888880806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050815260200185858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200183838080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200160011515815250600260008760070b81526020019081526020016000206000820151816000019080519060200190610d52929190610e92565b506020820151816001019081610d689190611d30565b506040820151816002019081610d7e9190611d30565b5060608201518160030160006101000a81548160ff0219169083151502179055509050508460070b7f7b6265887873467f1c48ffdb20746ee624f83e95fcda282cf9623f5179e87aae85858585604051610ddb9493929190611e2f565b60405180910390a250505050505050565b828054828255906000526020600020908101928215610e34579160200282015b82811115610e33578251829081610e239190611d30565b5091602001919060010190610e0c565b5b509050610e419190610edf565b5090565b828054828255906000526020600020908101928215610e81579160200282015b82811115610e80578251825591602001919060010190610e65565b5b509050610e8e9190610f03565b5090565b828054828255906000526020600020908101928215610ece579160200282015b82811115610ecd578251825591602001919060010190610eb2565b5b509050610edb9190610f20565b5090565b5b80821115610eff5760008181610ef69190610f3d565b50600101610ee0565b5090565b5b80821115610f1c576000816000905550600101610f04565b5090565b5b80821115610f39576000816000905550600101610f21565b5090565b508054610f49906118ef565b6000825580601f10610f5b5750610f7a565b601f016020900490600052602060002090810190610f799190610f03565b5b50565b6000604051905090565b600080fd5b600080fd5b60008160070b9050919050565b610fa781610f91565b8114610fb257600080fd5b50565b600081359050610fc481610f9e565b92915050565b600060208284031215610fe057610fdf610f87565b5b6000610fee84828501610fb5565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b61103681611023565b82525050565b6000611048838361102d565b60208301905092915050565b6000602082019050919050565b600061106c82610ff7565b6110768185611002565b935061108183611013565b8060005b838110156110b2578151611099888261103c565b97506110a483611054565b925050600181019050611085565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156110f95780820151818401526020810190506110de565b60008484015250505050565b6000601f19601f8301169050919050565b6000611121826110bf565b61112b81856110ca565b935061113b8185602086016110db565b61114481611105565b840191505092915050565b600060608201905081810360008301526111698186611061565b9050818103602083015261117d8185611116565b905081810360408301526111918184611116565b9050949350505050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b600082825260208201905092915050565b60006111e3826110bf565b6111ed81856111c7565b93506111fd8185602086016110db565b61120681611105565b840191505092915050565b600061121d83836111d8565b905092915050565b6000602082019050919050565b600061123d8261119b565b61124781856111a6565b935083602082028501611259856111b7565b8060005b8581101561129557848403895281516112768582611211565b945061128183611225565b925060208a0199505060018101905061125d565b50829750879550505050505092915050565b600060608201905081810360008301526112c18186611232565b905081810360208301526112d58185611116565b905081810360408301526112e98184611116565b9050949350505050565b600080fd5b600080fd5b600080fd5b60008083601f840112611318576113176112f3565b5b8235905067ffffffffffffffff811115611335576113346112f8565b5b602083019150836020820283011115611351576113506112fd565b5b9250929050565b60008083601f84011261136e5761136d6112f3565b5b8235905067ffffffffffffffff81111561138b5761138a6112f8565b5b6020830191508360018202830111156113a7576113a66112fd565b5b9250929050565b60008060008060008060006080888a0312156113cd576113cc610f87565b5b600088013567ffffffffffffffff8111156113eb576113ea610f8c565b5b6113f78a828b01611302565b9750975050602061140a8a828b01610fb5565b955050604088013567ffffffffffffffff81111561142b5761142a610f8c565b5b6114378a828b01611358565b9450945050606088013567ffffffffffffffff81111561145a57611459610f8c565b5b6114668a828b01611358565b925092505092959891949750929550565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006114a282611477565b9050919050565b6114b281611497565b81146114bd57600080fd5b50565b6000813590506114cf816114a9565b92915050565b6000602082840312156114eb576114ea610f87565b5b60006114f9848285016114c0565b91505092915050565b60008083601f840112611518576115176112f3565b5b8235905067ffffffffffffffff811115611535576115346112f8565b5b602083019150836020820283011115611551576115506112fd565b5b9250929050565b60008060008060008060006080888a03121561157757611576610f87565b5b600088013567ffffffffffffffff81111561159557611594610f8c565b5b6115a18a828b01611502565b975097505060206115b48a828b01610fb5565b955050604088013567ffffffffffffffff8111156115d5576115d4610f8c565b5b6115e18a828b01611358565b9450945050606088013567ffffffffffffffff81111561160457611603610f8c565b5b6116108a828b01611358565b925092505092959891949750929550565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b6116608161164d565b82525050565b60006116728383611657565b60208301905092915050565b6000602082019050919050565b600061169682611621565b6116a0818561162c565b93506116ab8361163d565b8060005b838110156116dc5781516116c38882611666565b97506116ce8361167e565b9250506001810190506116af565b5085935050505092915050565b60006060820190508181036000830152611703818661168b565b905081810360208301526117178185611116565b9050818103604083015261172b8184611116565b9050949350505050565b60008083601f84011261174b5761174a6112f3565b5b8235905067ffffffffffffffff811115611768576117676112f8565b5b602083019150836020820283011115611784576117836112fd565b5b9250929050565b60008060008060008060006080888a0312156117aa576117a9610f87565b5b600088013567ffffffffffffffff8111156117c8576117c7610f8c565b5b6117d48a828b01611735565b975097505060206117e78a828b01610fb5565b955050604088013567ffffffffffffffff81111561180857611807610f8c565b5b6118148a828b01611358565b9450945050606088013567ffffffffffffffff81111561183757611836610f8c565b5b6118438a828b01611358565b925092505092959891949750929550565b7f4e6f20496e744172726179456e74727920666f72207468697320726f756e6400600082015250565b600061188a601f836110ca565b915061189582611854565b602082019050919050565b600060208201905081810360008301526118b98161187d565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061190757607f821691505b60208210810361191a576119196118c0565b5b50919050565b7f4e6f204279746573456e74727920666f72207468697320726f756e6400000000600082015250565b6000611956601c836110ca565b915061196182611920565b602082019050919050565b6000602082019050818103600083015261198581611949565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6119c482611105565b810181811067ffffffffffffffff821117156119e3576119e261198c565b5b80604052505050565b60006119f6610f7d565b9050611a0282826119bb565b919050565b600067ffffffffffffffff821115611a2257611a2161198c565b5b602082029050602081019050919050565b600080fd5b600067ffffffffffffffff821115611a5357611a5261198c565b5b611a5c82611105565b9050602081019050919050565b82818337600083830152505050565b6000611a8b611a8684611a38565b6119ec565b905082815260208101848484011115611aa757611aa6611a33565b5b611ab2848285611a69565b509392505050565b600082601f830112611acf57611ace6112f3565b5b8135611adf848260208601611a78565b91505092915050565b6000611afb611af684611a07565b6119ec565b90508083825260208201905060208402830185811115611b1e57611b1d6112fd565b5b835b81811015611b6557803567ffffffffffffffff811115611b4357611b426112f3565b5b808601611b508982611aba565b85526020850194505050602081019050611b20565b5050509392505050565b6000611b7c368484611ae8565b905092915050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302611be67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611ba9565b611bf08683611ba9565b95508019841693508086168417925050509392505050565b6000819050919050565b6000611c2d611c28611c238461164d565b611c08565b61164d565b9050919050565b6000819050919050565b611c4783611c12565b611c5b611c5382611c34565b848454611bb6565b825550505050565b600090565b611c70611c63565b611c7b818484611c3e565b505050565b5b81811015611c9f57611c94600082611c68565b600181019050611c81565b5050565b601f821115611ce457611cb581611b84565b611cbe84611b99565b81016020851015611ccd578190505b611ce1611cd985611b99565b830182611c80565b50505b505050565b600082821c905092915050565b6000611d0760001984600802611ce9565b1980831691505092915050565b6000611d208383611cf6565b9150826002028217905092915050565b611d39826110bf565b67ffffffffffffffff811115611d5257611d5161198c565b5b611d5c82546118ef565b611d67828285611ca3565b600060209050601f831160018114611d9a5760008415611d88578287015190505b611d928582611d14565b865550611dfa565b601f198416611da886611b84565b60005b82811015611dd057848901518255600182019150602085019450602081019050611dab565b86831015611ded5784890151611de9601f891682611cf6565b8355505b6001600288020188555050505b505050505050565b6000611e0e83856110ca565b9350611e1b838584611a69565b611e2483611105565b840190509392505050565b60006040820190508181036000830152611e4a818688611e02565b90508181036020830152611e5f818486611e02565b905095945050505050565b611e7381611497565b82525050565b6000602082019050611e8e6000830184611e6a565b92915050565b7f4e6f20496e7452616e6765456e74727920666f72207468697320726f756e6400600082015250565b6000611eca601f836110ca565b9150611ed582611e94565b602082019050919050565b60006020820190508181036000830152611ef981611ebd565b905091905056fea2646970667358221220e65ee4a124b683fa7edaf9e48022f38690c0df2e9dcf340b26b30a0c7439781864736f6c63430008130033",
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
// Solidity: function getBytes(int64 round) view returns(string[] randomness, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCaller) GetBytes(opts *bind.CallOpts, round int64) (struct {
	Randomness []string
	Seed       string
	Signature  string
}, error) {
	var out []interface{}
	err := _DIAOracleRandomness.contract.Call(opts, &out, "getBytes", round)

	outstruct := new(struct {
		Randomness []string
		Seed       string
		Signature  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Randomness = *abi.ConvertType(out[0], new([]string)).(*[]string)
	outstruct.Seed = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Signature = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// GetBytes is a free data retrieval call binding the contract method 0x53759265.
//
// Solidity: function getBytes(int64 round) view returns(string[] randomness, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessSession) GetBytes(round int64) (struct {
	Randomness []string
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetBytes(&_DIAOracleRandomness.CallOpts, round)
}

// GetBytes is a free data retrieval call binding the contract method 0x53759265.
//
// Solidity: function getBytes(int64 round) view returns(string[] randomness, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCallerSession) GetBytes(round int64) (struct {
	Randomness []string
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetBytes(&_DIAOracleRandomness.CallOpts, round)
}

// GetIntArray is a free data retrieval call binding the contract method 0x45459cd0.
//
// Solidity: function getIntArray(int64 round) view returns(int256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCaller) GetIntArray(opts *bind.CallOpts, round int64) (struct {
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	var out []interface{}
	err := _DIAOracleRandomness.contract.Call(opts, &out, "getIntArray", round)

	outstruct := new(struct {
		RandomInts []*big.Int
		Seed       string
		Signature  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RandomInts = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Seed = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Signature = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// GetIntArray is a free data retrieval call binding the contract method 0x45459cd0.
//
// Solidity: function getIntArray(int64 round) view returns(int256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessSession) GetIntArray(round int64) (struct {
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetIntArray(&_DIAOracleRandomness.CallOpts, round)
}

// GetIntArray is a free data retrieval call binding the contract method 0x45459cd0.
//
// Solidity: function getIntArray(int64 round) view returns(int256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCallerSession) GetIntArray(round int64) (struct {
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetIntArray(&_DIAOracleRandomness.CallOpts, round)
}

// GetIntRange is a free data retrieval call binding the contract method 0xd9082a9b.
//
// Solidity: function getIntRange(int64 round) view returns(uint256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCaller) GetIntRange(opts *bind.CallOpts, round int64) (struct {
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	var out []interface{}
	err := _DIAOracleRandomness.contract.Call(opts, &out, "getIntRange", round)

	outstruct := new(struct {
		RandomInts []*big.Int
		Seed       string
		Signature  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RandomInts = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Seed = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Signature = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// GetIntRange is a free data retrieval call binding the contract method 0xd9082a9b.
//
// Solidity: function getIntRange(int64 round) view returns(uint256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessSession) GetIntRange(round int64) (struct {
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetIntRange(&_DIAOracleRandomness.CallOpts, round)
}

// GetIntRange is a free data retrieval call binding the contract method 0xd9082a9b.
//
// Solidity: function getIntRange(int64 round) view returns(uint256[] randomInts, string seed, string signature)
func (_DIAOracleRandomness *DIAOracleRandomnessCallerSession) GetIntRange(round int64) (struct {
	RandomInts []*big.Int
	Seed       string
	Signature  string
}, error) {
	return _DIAOracleRandomness.Contract.GetIntRange(&_DIAOracleRandomness.CallOpts, round)
}

// SetBytes is a paid mutator transaction binding the contract method 0x5d52ea10.
//
// Solidity: function setBytes(string[] randomness, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactor) SetBytes(opts *bind.TransactOpts, randomness []string, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.contract.Transact(opts, "setBytes", randomness, round, seed, signature)
}

// SetBytes is a paid mutator transaction binding the contract method 0x5d52ea10.
//
// Solidity: function setBytes(string[] randomness, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessSession) SetBytes(randomness []string, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.SetBytes(&_DIAOracleRandomness.TransactOpts, randomness, round, seed, signature)
}

// SetBytes is a paid mutator transaction binding the contract method 0x5d52ea10.
//
// Solidity: function setBytes(string[] randomness, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactorSession) SetBytes(randomness []string, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.SetBytes(&_DIAOracleRandomness.TransactOpts, randomness, round, seed, signature)
}

// SetIntArray is a paid mutator transaction binding the contract method 0xf8f0838e.
//
// Solidity: function setIntArray(int256[] randomInts, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactor) SetIntArray(opts *bind.TransactOpts, randomInts []*big.Int, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.contract.Transact(opts, "setIntArray", randomInts, round, seed, signature)
}

// SetIntArray is a paid mutator transaction binding the contract method 0xf8f0838e.
//
// Solidity: function setIntArray(int256[] randomInts, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessSession) SetIntArray(randomInts []*big.Int, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.SetIntArray(&_DIAOracleRandomness.TransactOpts, randomInts, round, seed, signature)
}

// SetIntArray is a paid mutator transaction binding the contract method 0xf8f0838e.
//
// Solidity: function setIntArray(int256[] randomInts, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactorSession) SetIntArray(randomInts []*big.Int, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.SetIntArray(&_DIAOracleRandomness.TransactOpts, randomInts, round, seed, signature)
}

// SetIntRange is a paid mutator transaction binding the contract method 0x809fcd51.
//
// Solidity: function setIntRange(uint256[] randomInts, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactor) SetIntRange(opts *bind.TransactOpts, randomInts []*big.Int, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.contract.Transact(opts, "setIntRange", randomInts, round, seed, signature)
}

// SetIntRange is a paid mutator transaction binding the contract method 0x809fcd51.
//
// Solidity: function setIntRange(uint256[] randomInts, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessSession) SetIntRange(randomInts []*big.Int, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.SetIntRange(&_DIAOracleRandomness.TransactOpts, randomInts, round, seed, signature)
}

// SetIntRange is a paid mutator transaction binding the contract method 0x809fcd51.
//
// Solidity: function setIntRange(uint256[] randomInts, int64 round, string seed, string signature) returns()
func (_DIAOracleRandomness *DIAOracleRandomnessTransactorSession) SetIntRange(randomInts []*big.Int, round int64, seed string, signature string) (*types.Transaction, error) {
	return _DIAOracleRandomness.Contract.SetIntRange(&_DIAOracleRandomness.TransactOpts, randomInts, round, seed, signature)
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
	Round     *big.Int
	Seed      string
	Signature string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBytesSet is a free log retrieval operation binding the contract event 0xf08ab236a050614700c1f2c02caaf5de6f6ee744d522ab71eb23633fac1909c2.
//
// Solidity: event BytesSet(int256 indexed round, string seed, string signature)
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

// WatchBytesSet is a free log subscription operation binding the contract event 0xf08ab236a050614700c1f2c02caaf5de6f6ee744d522ab71eb23633fac1909c2.
//
// Solidity: event BytesSet(int256 indexed round, string seed, string signature)
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

// ParseBytesSet is a log parse operation binding the contract event 0xf08ab236a050614700c1f2c02caaf5de6f6ee744d522ab71eb23633fac1909c2.
//
// Solidity: event BytesSet(int256 indexed round, string seed, string signature)
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
	Round     *big.Int
	Seed      string
	Signature string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIntArraySet is a free log retrieval operation binding the contract event 0x7b6265887873467f1c48ffdb20746ee624f83e95fcda282cf9623f5179e87aae.
//
// Solidity: event IntArraySet(int256 indexed round, string seed, string signature)
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

// WatchIntArraySet is a free log subscription operation binding the contract event 0x7b6265887873467f1c48ffdb20746ee624f83e95fcda282cf9623f5179e87aae.
//
// Solidity: event IntArraySet(int256 indexed round, string seed, string signature)
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

// ParseIntArraySet is a log parse operation binding the contract event 0x7b6265887873467f1c48ffdb20746ee624f83e95fcda282cf9623f5179e87aae.
//
// Solidity: event IntArraySet(int256 indexed round, string seed, string signature)
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
	Round     *big.Int
	Seed      string
	Signature string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterIntRangeSet is a free log retrieval operation binding the contract event 0xc43db9672cb481f936463442b4cefa9165d1388dfa3804b7585f84bb05f32107.
//
// Solidity: event IntRangeSet(int256 indexed round, string seed, string signature)
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

// WatchIntRangeSet is a free log subscription operation binding the contract event 0xc43db9672cb481f936463442b4cefa9165d1388dfa3804b7585f84bb05f32107.
//
// Solidity: event IntRangeSet(int256 indexed round, string seed, string signature)
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

// ParseIntRangeSet is a log parse operation binding the contract event 0xc43db9672cb481f936463442b4cefa9165d1388dfa3804b7585f84bb05f32107.
//
// Solidity: event IntRangeSet(int256 indexed round, string seed, string signature)
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
