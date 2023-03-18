// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package multisend

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

// MultisendMetaData contains all meta data concerning the Multisend contract.
var MultisendMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"multisend\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"multisendToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506110bc806100606000396000f3fe60806040526004361061003f5760003560e01c80630b66f3f514610044578063893d20e814610060578063a6f9dae11461008b578063aad41a41146100b4575b600080fd5b61005e60048036038101906100599190610956565b6100d0565b005b34801561006c57600080fd5b50610075610332565b60405161008291906109f0565b60405180910390f35b34801561009757600080fd5b506100b260048036038101906100ad9190610a0b565b61035b565b005b6100ce60048036038101906100c99190610b39565b61042c565b005b8051825114610114576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161010b90610c0e565b60405180910390fd5b61011d816105b1565b8373ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff16306040518363ffffffff1660e01b8152600401610178929190610c2e565b602060405180830381865afa158015610195573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101b99190610c6c565b10156101fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101f190610ce5565b60405180910390fd5b60005b825181101561032c578373ffffffffffffffffffffffffffffffffffffffff166323b872dd60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1685848151811061025657610255610d05565b5b602002602001015185858151811061027157610270610d05565b5b60200260200101516040518463ffffffff1660e01b815260040161029793929190610d43565b6020604051808303816000875af11580156102b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102da9190610db2565b610319576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031090610e2b565b60405180910390fd5b808061032490610e7a565b9150506101fd565b50505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103e090610f0e565b60405180910390fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b190610f0e565b60405180910390fd5b80518251146104fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104f590610fa0565b60405180910390fd5b6000610509826105b1565b90508034101561054e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161054590611032565b60405180910390fd5b60005b83518110156105ab576105988482815181106105705761056f610d05565b5b602002602001015184838151811061058b5761058a610d05565b5b6020026020010151610609565b80806105a390610e7a565b915050610551565b50505050565b6000806000905060005b83518110156105ff578381815181106105d7576105d6610d05565b5b6020026020010151826105ea9190611052565b915080806105f790610e7a565b9150506105bb565b5080915050919050565b8173ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f1935050505015801561064f573d6000803e3d6000fd5b505050565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061069382610668565b9050919050565b60006106a582610688565b9050919050565b6106b58161069a565b81146106c057600080fd5b50565b6000813590506106d2816106ac565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610726826106dd565b810181811067ffffffffffffffff82111715610745576107446106ee565b5b80604052505050565b6000610758610654565b9050610764828261071d565b919050565b600067ffffffffffffffff821115610784576107836106ee565b5b602082029050602081019050919050565b600080fd5b6107a381610688565b81146107ae57600080fd5b50565b6000813590506107c08161079a565b92915050565b60006107d96107d484610769565b61074e565b905080838252602082019050602084028301858111156107fc576107fb610795565b5b835b81811015610825578061081188826107b1565b8452602084019350506020810190506107fe565b5050509392505050565b600082601f830112610844576108436106d8565b5b81356108548482602086016107c6565b91505092915050565b600067ffffffffffffffff821115610878576108776106ee565b5b602082029050602081019050919050565b6000819050919050565b61089c81610889565b81146108a757600080fd5b50565b6000813590506108b981610893565b92915050565b60006108d26108cd8461085d565b61074e565b905080838252602082019050602084028301858111156108f5576108f4610795565b5b835b8181101561091e578061090a88826108aa565b8452602084019350506020810190506108f7565b5050509392505050565b600082601f83011261093d5761093c6106d8565b5b813561094d8482602086016108bf565b91505092915050565b60008060006060848603121561096f5761096e61065e565b5b600061097d868287016106c3565b935050602084013567ffffffffffffffff81111561099e5761099d610663565b5b6109aa8682870161082f565b925050604084013567ffffffffffffffff8111156109cb576109ca610663565b5b6109d786828701610928565b9150509250925092565b6109ea81610688565b82525050565b6000602082019050610a0560008301846109e1565b92915050565b600060208284031215610a2157610a2061065e565b5b6000610a2f848285016107b1565b91505092915050565b600067ffffffffffffffff821115610a5357610a526106ee565b5b602082029050602081019050919050565b6000610a6f82610668565b9050919050565b610a7f81610a64565b8114610a8a57600080fd5b50565b600081359050610a9c81610a76565b92915050565b6000610ab5610ab084610a38565b61074e565b90508083825260208201905060208402830185811115610ad857610ad7610795565b5b835b81811015610b015780610aed8882610a8d565b845260208401935050602081019050610ada565b5050509392505050565b600082601f830112610b2057610b1f6106d8565b5b8135610b30848260208601610aa2565b91505092915050565b60008060408385031215610b5057610b4f61065e565b5b600083013567ffffffffffffffff811115610b6e57610b6d610663565b5b610b7a85828601610b0b565b925050602083013567ffffffffffffffff811115610b9b57610b9a610663565b5b610ba785828601610928565b9150509250929050565b600082825260208201905092915050565b7f496e76616c696420696e707574206c656e677468730000000000000000000000600082015250565b6000610bf8601583610bb1565b9150610c0382610bc2565b602082019050919050565b60006020820190508181036000830152610c2781610beb565b9050919050565b6000604082019050610c4360008301856109e1565b610c5060208301846109e1565b9392505050565b600081519050610c6681610893565b92915050565b600060208284031215610c8257610c8161065e565b5b6000610c9084828501610c57565b91505092915050565b7f496e73756666696369656e7420616c6c6f77616e636500000000000000000000600082015250565b6000610ccf601683610bb1565b9150610cda82610c99565b602082019050919050565b60006020820190508181036000830152610cfe81610cc2565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b610d3d81610889565b82525050565b6000606082019050610d5860008301866109e1565b610d6560208301856109e1565b610d726040830184610d34565b949350505050565b60008115159050919050565b610d8f81610d7a565b8114610d9a57600080fd5b50565b600081519050610dac81610d86565b92915050565b600060208284031215610dc857610dc761065e565b5b6000610dd684828501610d9d565b91505092915050565b7f5472616e73666572206661696c65640000000000000000000000000000000000600082015250565b6000610e15600f83610bb1565b9150610e2082610ddf565b602082019050919050565b60006020820190508181036000830152610e4481610e08565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610e8582610889565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610eb757610eb6610e4b565b5b600182019050919050565b7f43616c6c6572206973206e6f74206f776e657200000000000000000000000000600082015250565b6000610ef8601383610bb1565b9150610f0382610ec2565b602082019050919050565b60006020820190508181036000830152610f2781610eeb565b9050919050565b7f546865206c656e677468206f662074776f2061727261792073686f756c64206260008201527f65207468652073616d6500000000000000000000000000000000000000000000602082015250565b6000610f8a602a83610bb1565b9150610f9582610f2e565b604082019050919050565b60006020820190508181036000830152610fb981610f7d565b9050919050565b7f5468652076616c7565206973206e6f742073756666696369656e74206f72206560008201527f7863656564000000000000000000000000000000000000000000000000000000602082015250565b600061101c602583610bb1565b915061102782610fc0565b604082019050919050565b6000602082019050818103600083015261104b8161100f565b9050919050565b600061105d82610889565b915061106883610889565b92508282019050808211156110805761107f610e4b565b5b9291505056fea2646970667358221220e5cbee381beb6c802d646007d73304266aec58411d281523671481fe50a1d2a964736f6c63430008110033",
}

// MultisendABI is the input ABI used to generate the binding from.
// Deprecated: Use MultisendMetaData.ABI instead.
var MultisendABI = MultisendMetaData.ABI

// MultisendBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MultisendMetaData.Bin instead.
var MultisendBin = MultisendMetaData.Bin

// DeployMultisend deploys a new Ethereum contract, binding an instance of Multisend to it.
func DeployMultisend(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Multisend, error) {
	parsed, err := MultisendMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MultisendBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Multisend{MultisendCaller: MultisendCaller{contract: contract}, MultisendTransactor: MultisendTransactor{contract: contract}, MultisendFilterer: MultisendFilterer{contract: contract}}, nil
}

// Multisend is an auto generated Go binding around an Ethereum contract.
type Multisend struct {
	MultisendCaller     // Read-only binding to the contract
	MultisendTransactor // Write-only binding to the contract
	MultisendFilterer   // Log filterer for contract events
}

// MultisendCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultisendCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisendTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultisendTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisendFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultisendFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultisendSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultisendSession struct {
	Contract     *Multisend        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultisendCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultisendCallerSession struct {
	Contract *MultisendCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// MultisendTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultisendTransactorSession struct {
	Contract     *MultisendTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MultisendRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultisendRaw struct {
	Contract *Multisend // Generic contract binding to access the raw methods on
}

// MultisendCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultisendCallerRaw struct {
	Contract *MultisendCaller // Generic read-only contract binding to access the raw methods on
}

// MultisendTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultisendTransactorRaw struct {
	Contract *MultisendTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultisend creates a new instance of Multisend, bound to a specific deployed contract.
func NewMultisend(address common.Address, backend bind.ContractBackend) (*Multisend, error) {
	contract, err := bindMultisend(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Multisend{MultisendCaller: MultisendCaller{contract: contract}, MultisendTransactor: MultisendTransactor{contract: contract}, MultisendFilterer: MultisendFilterer{contract: contract}}, nil
}

// NewMultisendCaller creates a new read-only instance of Multisend, bound to a specific deployed contract.
func NewMultisendCaller(address common.Address, caller bind.ContractCaller) (*MultisendCaller, error) {
	contract, err := bindMultisend(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultisendCaller{contract: contract}, nil
}

// NewMultisendTransactor creates a new write-only instance of Multisend, bound to a specific deployed contract.
func NewMultisendTransactor(address common.Address, transactor bind.ContractTransactor) (*MultisendTransactor, error) {
	contract, err := bindMultisend(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultisendTransactor{contract: contract}, nil
}

// NewMultisendFilterer creates a new log filterer instance of Multisend, bound to a specific deployed contract.
func NewMultisendFilterer(address common.Address, filterer bind.ContractFilterer) (*MultisendFilterer, error) {
	contract, err := bindMultisend(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultisendFilterer{contract: contract}, nil
}

// bindMultisend binds a generic wrapper to an already deployed contract.
func bindMultisend(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultisendABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisend *MultisendRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multisend.Contract.MultisendCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisend *MultisendRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisend.Contract.MultisendTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisend *MultisendRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisend.Contract.MultisendTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Multisend *MultisendCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Multisend.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Multisend *MultisendTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Multisend.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Multisend *MultisendTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Multisend.Contract.contract.Transact(opts, method, params...)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Multisend *MultisendCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Multisend.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Multisend *MultisendSession) GetOwner() (common.Address, error) {
	return _Multisend.Contract.GetOwner(&_Multisend.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Multisend *MultisendCallerSession) GetOwner() (common.Address, error) {
	return _Multisend.Contract.GetOwner(&_Multisend.CallOpts)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Multisend *MultisendTransactor) ChangeOwner(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "changeOwner", newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Multisend *MultisendSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Multisend.Contract.ChangeOwner(&_Multisend.TransactOpts, newOwner)
}

// ChangeOwner is a paid mutator transaction binding the contract method 0xa6f9dae1.
//
// Solidity: function changeOwner(address newOwner) returns()
func (_Multisend *MultisendTransactorSession) ChangeOwner(newOwner common.Address) (*types.Transaction, error) {
	return _Multisend.Contract.ChangeOwner(&_Multisend.TransactOpts, newOwner)
}

// Multisend is a paid mutator transaction binding the contract method 0xaad41a41.
//
// Solidity: function multisend(address[] addresses, uint256[] amounts) payable returns()
func (_Multisend *MultisendTransactor) Multisend(opts *bind.TransactOpts, addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "multisend", addresses, amounts)
}

// Multisend is a paid mutator transaction binding the contract method 0xaad41a41.
//
// Solidity: function multisend(address[] addresses, uint256[] amounts) payable returns()
func (_Multisend *MultisendSession) Multisend(addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.Multisend(&_Multisend.TransactOpts, addresses, amounts)
}

// Multisend is a paid mutator transaction binding the contract method 0xaad41a41.
//
// Solidity: function multisend(address[] addresses, uint256[] amounts) payable returns()
func (_Multisend *MultisendTransactorSession) Multisend(addresses []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.Multisend(&_Multisend.TransactOpts, addresses, amounts)
}

// MultisendToken is a paid mutator transaction binding the contract method 0x0b66f3f5.
//
// Solidity: function multisendToken(address token, address[] recipients, uint256[] amounts) payable returns()
func (_Multisend *MultisendTransactor) MultisendToken(opts *bind.TransactOpts, token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.contract.Transact(opts, "multisendToken", token, recipients, amounts)
}

// MultisendToken is a paid mutator transaction binding the contract method 0x0b66f3f5.
//
// Solidity: function multisendToken(address token, address[] recipients, uint256[] amounts) payable returns()
func (_Multisend *MultisendSession) MultisendToken(token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.MultisendToken(&_Multisend.TransactOpts, token, recipients, amounts)
}

// MultisendToken is a paid mutator transaction binding the contract method 0x0b66f3f5.
//
// Solidity: function multisendToken(address token, address[] recipients, uint256[] amounts) payable returns()
func (_Multisend *MultisendTransactorSession) MultisendToken(token common.Address, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Multisend.Contract.MultisendToken(&_Multisend.TransactOpts, token, recipients, amounts)
}
