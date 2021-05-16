// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// VoteABI is the input ABI used to generate the binding from.
const VoteABI = "[{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"proposalNames\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"funds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"chairperson\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"giveRightToVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"name\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"funds\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokeVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposal\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"voted\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"vote\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winnerName\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"winnerName_\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"winningProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"winningProposal_\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// VoteFuncSigs maps the 4-byte function signature to its string representation.
var VoteFuncSigs = map[string]string{
	"2e4176cf": "chairperson()",
	"5c19a95c": "delegate(address)",
	"ce50aa7d": "giveRightToVote(address,uint256)",
	"013cf08b": "proposals(uint256)",
	"43c14b22": "revokeVote()",
	"0121b93f": "vote(uint256)",
	"a3ec138d": "voters(address)",
	"e2ba53f0": "winnerName()",
	"609ff1bd": "winningProposal()",
}

// VoteBin is the compiled bytecode used for deploying new contracts.
var VoteBin = "0x608060405234801561001057600080fd5b506040516109b13803806109b18339818101604052604081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825186602082028301116401000000008211171561008557600080fd5b82525081516020918201928201910280838360005b838110156100b257818101518382015260200161009a565b50505050905001604052602001805160405193929190846401000000008211156100db57600080fd5b9083019060208201858111156100f057600080fd5b825186602082028301116401000000008211171561010d57600080fd5b82525081516020918201928201910280838360005b8381101561013a578181015183820152602001610122565b50505050919091016040908152600080546001600160a01b03191633178082556001600160a01b03168152600160208190529181209190915593505050505b8251811015610201576002604051806060016040528085848151811061019b57fe5b60200260200101518152602001600081526020018484815181106101bb57fe5b6020908102919091018101519091528254600181810185556000948552938290208351600390920201908155908201518184015560409091015160029091015501610179565b50505061079e806102136000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80635c19a95c116100665780635c19a95c1461011e578063609ff1bd14610144578063a3ec138d1461015e578063ce50aa7d146101b2578063e2ba53f0146101de57610093565b80630121b93f14610098578063013cf08b146100b75780632e4176cf146100f257806343c14b2214610116575b600080fd5b6100b5600480360360208110156100ae57600080fd5b50356101e6565b005b6100d4600480360360208110156100cd57600080fd5b5035610289565b60408051938452602084019290925282820152519081900360600190f35b6100fa6102bc565b604080516001600160a01b039092168252519081900360200190f35b6100b56102cb565b6100b56004803603602081101561013457600080fd5b50356001600160a01b03166103b2565b61014c6105a1565b60408051918252519081900360200190f35b6101846004803603602081101561017457600080fd5b50356001600160a01b0316610608565b6040805194855292151560208501526001600160a01b03909116838301526060830152519081900360800190f35b6100b5600480360360408110156101c857600080fd5b506001600160a01b03813516906020013561063c565b61014c610713565b3360009081526001602081905260409091209081015460ff1615610242576040805162461bcd60e51b815260206004820152600e60248201526d20b63932b0b23c903b37ba32b21760911b604482015290519081900360640190fd5b6001818101805460ff191690911790556002808201839055815481549091908490811061026b57fe5b60009182526020909120600160039092020101805490910190555050565b6002818154811061029957600080fd5b600091825260209091206003909102018054600182015460029092015490925083565b6000546001600160a01b031681565b3360009081526001602081905260408083209182015461010081046001600160a01b031684529220909160ff16156103ae57600182015461010090046001600160a01b03166103465781546002808401548154811061032657fe5b600091825260209091206001600390920201018054919091039055610394565b600181015460ff161561038c5781546002828101548154811061036557fe5b60009182526020909120600160039092020101805491909103905581548154038155610394565b815481540381555b6001820180546001600160a81b0319169055600060028301555b5050565b3360009081526001602081905260409091209081015460ff1615610412576040805162461bcd60e51b81526020600482015260126024820152712cb7ba9030b63932b0b23c903b37ba32b21760711b604482015290519081900360640190fd5b6001600160a01b038216331415610470576040805162461bcd60e51b815260206004820152601e60248201527f53656c662d64656c65676174696f6e20697320646973616c6c6f7765642e0000604482015290519081900360640190fd5b6001600160a01b03828116600090815260016020819052604090912001546101009004161561051a576001600160a01b039182166000908152600160208190526040909120015461010090049091169033821415610515576040805162461bcd60e51b815260206004820152601960248201527f466f756e64206c6f6f7020696e2064656c65676174696f6e2e00000000000000604482015290519081900360640190fd5b610470565b6001818101805460ff19168217610100600160a81b0319166101006001600160a01b0386169081029190911790915560009081526020829052604090209081015460ff16156105945781546002828101548154811061057557fe5b600091825260209091206001600390920201018054909101905561059c565b815481540181555b505050565b600080805b6002548110156106035781600282815481106105be57fe5b90600052602060002090600302016001015411156105fb57600281815481106105e357fe5b90600052602060002090600302016001015491508092505b6001016105a6565b505090565b600160208190526000918252604090912080549181015460029091015460ff82169161010090046001600160a01b03169084565b6000546001600160a01b031633146106855760405162461bcd60e51b81526004018080602001828103825260288152602001806107416028913960400191505060405180910390fd5b6001600160a01b0382166000908152600160208190526040909120015460ff16156106f7576040805162461bcd60e51b815260206004820152601860248201527f54686520766f74657220616c726561647920766f7465642e0000000000000000604482015290519081900360640190fd5b6001600160a01b03909116600090815260016020526040902055565b6000600261071f6105a1565b8154811061072957fe5b90600052602060002090600302016000015490509056fe4f6e6c79206368616972706572736f6e2063616e206769766520726967687420746f20766f74652ea2646970667358221220672944a87705ba9ecb0daf080fa4f4805fe74f8d2db1a1fd0a0b7de233281a6964736f6c63430007040033"

// DeployVote deploys a new Ethereum contract, binding an instance of Vote to it.
func DeployVote(auth *bind.TransactOpts, backend bind.ContractBackend, proposalNames [][32]byte, funds []*big.Int) (common.Address, *types.Transaction, *Vote, error) {
	parsed, err := abi.JSON(strings.NewReader(VoteABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VoteBin), backend, proposalNames, funds)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// Vote is an auto generated Go binding around an Ethereum contract.
type Vote struct {
	VoteCaller     // Read-only binding to the contract
	VoteTransactor // Write-only binding to the contract
	VoteFilterer   // Log filterer for contract events
}

// VoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoteSession struct {
	Contract     *Vote             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoteCallerSession struct {
	Contract *VoteCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoteTransactorSession struct {
	Contract     *VoteTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoteRaw struct {
	Contract *Vote // Generic contract binding to access the raw methods on
}

// VoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoteCallerRaw struct {
	Contract *VoteCaller // Generic read-only contract binding to access the raw methods on
}

// VoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoteTransactorRaw struct {
	Contract *VoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVote creates a new instance of Vote, bound to a specific deployed contract.
func NewVote(address common.Address, backend bind.ContractBackend) (*Vote, error) {
	contract, err := bindVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// NewVoteCaller creates a new read-only instance of Vote, bound to a specific deployed contract.
func NewVoteCaller(address common.Address, caller bind.ContractCaller) (*VoteCaller, error) {
	contract, err := bindVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoteCaller{contract: contract}, nil
}

// NewVoteTransactor creates a new write-only instance of Vote, bound to a specific deployed contract.
func NewVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*VoteTransactor, error) {
	contract, err := bindVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoteTransactor{contract: contract}, nil
}

// NewVoteFilterer creates a new log filterer instance of Vote, bound to a specific deployed contract.
func NewVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*VoteFilterer, error) {
	contract, err := bindVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoteFilterer{contract: contract}, nil
}

// bindVote binds a generic wrapper to an already deployed contract.
func bindVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.VoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transact(opts, method, params...)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Vote *VoteCaller) Chairperson(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "chairperson")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Vote *VoteSession) Chairperson() (common.Address, error) {
	return _Vote.Contract.Chairperson(&_Vote.CallOpts)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Vote *VoteCallerSession) Chairperson() (common.Address, error) {
	return _Vote.Contract.Chairperson(&_Vote.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 voteCount, uint256 funds)
func (_Vote *VoteCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
	Funds     *big.Int
}, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Name      [32]byte
		VoteCount *big.Int
		Funds     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.VoteCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Funds = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 voteCount, uint256 funds)
func (_Vote *VoteSession) Proposals(arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
	Funds     *big.Int
}, error) {
	return _Vote.Contract.Proposals(&_Vote.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(bytes32 name, uint256 voteCount, uint256 funds)
func (_Vote *VoteCallerSession) Proposals(arg0 *big.Int) (struct {
	Name      [32]byte
	VoteCount *big.Int
	Funds     *big.Int
}, error) {
	return _Vote.Contract.Proposals(&_Vote.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 weight, bool voted, address delegate, uint256 vote)
func (_Vote *VoteCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "voters", arg0)

	outstruct := new(struct {
		Weight   *big.Int
		Voted    bool
		Delegate common.Address
		Vote     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Weight = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Voted = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Delegate = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Vote = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 weight, bool voted, address delegate, uint256 vote)
func (_Vote *VoteSession) Voters(arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	return _Vote.Contract.Voters(&_Vote.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(uint256 weight, bool voted, address delegate, uint256 vote)
func (_Vote *VoteCallerSession) Voters(arg0 common.Address) (struct {
	Weight   *big.Int
	Voted    bool
	Delegate common.Address
	Vote     *big.Int
}, error) {
	return _Vote.Contract.Voters(&_Vote.CallOpts, arg0)
}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() view returns(bytes32 winnerName_)
func (_Vote *VoteCaller) WinnerName(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "winnerName")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() view returns(bytes32 winnerName_)
func (_Vote *VoteSession) WinnerName() ([32]byte, error) {
	return _Vote.Contract.WinnerName(&_Vote.CallOpts)
}

// WinnerName is a free data retrieval call binding the contract method 0xe2ba53f0.
//
// Solidity: function winnerName() view returns(bytes32 winnerName_)
func (_Vote *VoteCallerSession) WinnerName() ([32]byte, error) {
	return _Vote.Contract.WinnerName(&_Vote.CallOpts)
}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint256 winningProposal_)
func (_Vote *VoteCaller) WinningProposal(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "winningProposal")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint256 winningProposal_)
func (_Vote *VoteSession) WinningProposal() (*big.Int, error) {
	return _Vote.Contract.WinningProposal(&_Vote.CallOpts)
}

// WinningProposal is a free data retrieval call binding the contract method 0x609ff1bd.
//
// Solidity: function winningProposal() view returns(uint256 winningProposal_)
func (_Vote *VoteCallerSession) WinningProposal() (*big.Int, error) {
	return _Vote.Contract.WinningProposal(&_Vote.CallOpts)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_Vote *VoteTransactor) Delegate(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "delegate", to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_Vote *VoteSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _Vote.Contract.Delegate(&_Vote.TransactOpts, to)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address to) returns()
func (_Vote *VoteTransactorSession) Delegate(to common.Address) (*types.Transaction, error) {
	return _Vote.Contract.Delegate(&_Vote.TransactOpts, to)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0xce50aa7d.
//
// Solidity: function giveRightToVote(address voter, uint256 number) returns()
func (_Vote *VoteTransactor) GiveRightToVote(opts *bind.TransactOpts, voter common.Address, number *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "giveRightToVote", voter, number)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0xce50aa7d.
//
// Solidity: function giveRightToVote(address voter, uint256 number) returns()
func (_Vote *VoteSession) GiveRightToVote(voter common.Address, number *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.GiveRightToVote(&_Vote.TransactOpts, voter, number)
}

// GiveRightToVote is a paid mutator transaction binding the contract method 0xce50aa7d.
//
// Solidity: function giveRightToVote(address voter, uint256 number) returns()
func (_Vote *VoteTransactorSession) GiveRightToVote(voter common.Address, number *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.GiveRightToVote(&_Vote.TransactOpts, voter, number)
}

// RevokeVote is a paid mutator transaction binding the contract method 0x43c14b22.
//
// Solidity: function revokeVote() returns()
func (_Vote *VoteTransactor) RevokeVote(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "revokeVote")
}

// RevokeVote is a paid mutator transaction binding the contract method 0x43c14b22.
//
// Solidity: function revokeVote() returns()
func (_Vote *VoteSession) RevokeVote() (*types.Transaction, error) {
	return _Vote.Contract.RevokeVote(&_Vote.TransactOpts)
}

// RevokeVote is a paid mutator transaction binding the contract method 0x43c14b22.
//
// Solidity: function revokeVote() returns()
func (_Vote *VoteTransactorSession) RevokeVote() (*types.Transaction, error) {
	return _Vote.Contract.RevokeVote(&_Vote.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) returns()
func (_Vote *VoteTransactor) Vote(opts *bind.TransactOpts, proposal *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "vote", proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) returns()
func (_Vote *VoteSession) Vote(proposal *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.Vote(&_Vote.TransactOpts, proposal)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 proposal) returns()
func (_Vote *VoteTransactorSession) Vote(proposal *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.Vote(&_Vote.TransactOpts, proposal)
}
