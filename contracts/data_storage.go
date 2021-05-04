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

// DataStorageABI is the input ABI used to generate the binding from.
const DataStorageABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_didChainAddr\",\"type\":\"string\"}],\"name\":\"LogSetDidChainAddr\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_didClaimHash\",\"type\":\"string\"}],\"name\":\"LogSetDidClaimHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_didDocumentHash\",\"type\":\"string\"}],\"name\":\"LogSetDidDocumentHash\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_didPublicKey\",\"type\":\"string\"}],\"name\":\"LogSetDidPublicKey\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fromAddr\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"toAddr\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogSetTransaction\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getDidChainAddr\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getDidClaimHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getDidDocumentHash\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getDidPublicKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txId\",\"type\":\"string\"}],\"name\":\"getTransaction\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_didChainAddr\",\"type\":\"string\"}],\"name\":\"setDidChainAddr\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_didClaimHash\",\"type\":\"string\"}],\"name\":\"setDidClaimHash\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_didDocumentHash\",\"type\":\"string\"}],\"name\":\"setDidDocumentHash\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_didPublicKey\",\"type\":\"string\"}],\"name\":\"setDidPublicKey\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"txId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_fromAddr\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_toAddr\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"setTransaction\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// DataStorageFuncSigs maps the 4-byte function signature to its string representation.
var DataStorageFuncSigs = map[string]string{
	"09b4e1a2": "getDidChainAddr(string)",
	"78ff3fe8": "getDidClaimHash(string)",
	"da1bc0b7": "getDidDocumentHash(string)",
	"d6ca0d57": "getDidPublicKey(string)",
	"ac2ccd07": "getTransaction(string)",
	"3e06d9e9": "setDidChainAddr(string,string)",
	"a8917e05": "setDidClaimHash(string,string)",
	"1dd5ec83": "setDidDocumentHash(string,string)",
	"46f8acd2": "setDidPublicKey(string,string)",
	"5a7f8c66": "setTransaction(string,string,string,uint256)",
}

// DataStorageBin is the compiled bytecode used for deploying new contracts.
var DataStorageBin = "0x608060405234801561001057600080fd5b50611773806100206000396000f3fe6080604052600436106100915760003560e01c806378ff3fe81161005957806378ff3fe8146106e9578063a8917e051461079a578063ac2ccd07146108c3578063d6ca0d5714610a59578063da1bc0b714610b0a57610091565b806309b4e1a2146100965780631dd5ec83146101bc5780633e06d9e9146102e757806346f8acd2146104105780635a7f8c6614610539575b600080fd5b3480156100a257600080fd5b50610147600480360360208110156100b957600080fd5b810190602081018135600160201b8111156100d357600080fd5b8201836020820111156100e557600080fd5b803590602001918460018302840111600160201b8311171561010657600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610bbb945050505050565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610181578181015183820152602001610169565b50505050905090810190601f1680156101ae5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6102e5600480360360408110156101d257600080fd5b810190602081018135600160201b8111156101ec57600080fd5b8201836020820111156101fe57600080fd5b803590602001918460018302840111600160201b8311171561021f57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561027157600080fd5b82018360208201111561028357600080fd5b803590602001918460018302840111600160201b831117156102a457600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610cbe945050505050565b005b6102e5600480360360408110156102fd57600080fd5b810190602081018135600160201b81111561031757600080fd5b82018360208201111561032957600080fd5b803590602001918460018302840111600160201b8311171561034a57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561039c57600080fd5b8201836020820111156103ae57600080fd5b803590602001918460018302840111600160201b831117156103cf57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610e4d945050505050565b6102e56004803603604081101561042657600080fd5b810190602081018135600160201b81111561044057600080fd5b82018360208201111561045257600080fd5b803590602001918460018302840111600160201b8311171561047357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156104c557600080fd5b8201836020820111156104d757600080fd5b803590602001918460018302840111600160201b831117156104f857600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550610f3b945050505050565b6102e56004803603608081101561054f57600080fd5b810190602081018135600160201b81111561056957600080fd5b82018360208201111561057b57600080fd5b803590602001918460018302840111600160201b8311171561059c57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b8111156105ee57600080fd5b82018360208201111561060057600080fd5b803590602001918460018302840111600160201b8311171561062157600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561067357600080fd5b82018360208201111561068557600080fd5b803590602001918460018302840111600160201b831117156106a657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295505091359250611029915050565b3480156106f557600080fd5b506101476004803603602081101561070c57600080fd5b810190602081018135600160201b81111561072657600080fd5b82018360208201111561073857600080fd5b803590602001918460018302840111600160201b8311171561075957600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061127e945050505050565b6102e5600480360360408110156107b057600080fd5b810190602081018135600160201b8111156107ca57600080fd5b8201836020820111156107dc57600080fd5b803590602001918460018302840111600160201b831117156107fd57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b81111561084f57600080fd5b82018360208201111561086157600080fd5b803590602001918460018302840111600160201b8311171561088257600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506112c0945050505050565b3480156108cf57600080fd5b50610974600480360360208110156108e657600080fd5b810190602081018135600160201b81111561090057600080fd5b82018360208201111561091257600080fd5b803590602001918460018302840111600160201b8311171561093357600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506113ae945050505050565b604051808060200180602001848152602001838103835286818151815260200191508051906020019080838360005b838110156109bb5781810151838201526020016109a3565b50505050905090810190601f1680156109e85780820380516001836020036101000a031916815260200191505b50838103825285518152855160209182019187019080838360005b83811015610a1b578181015183820152602001610a03565b50505050905090810190601f168015610a485780820380516001836020036101000a031916815260200191505b509550505050505060405180910390f35b348015610a6557600080fd5b5061014760048036036020811015610a7c57600080fd5b810190602081018135600160201b811115610a9657600080fd5b820183602082011115610aa857600080fd5b803590602001918460018302840111600160201b83111715610ac957600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929550611618945050505050565b348015610b1657600080fd5b5061014760048036036020811015610b2d57600080fd5b810190602081018135600160201b811115610b4757600080fd5b820183602082011115610b5957600080fd5b803590602001918460018302840111600160201b83111715610b7a57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955061165a945050505050565b6060815160001415610bcc57600080fd5b6000826040518082805190602001908083835b60208310610bfe5780518252601f199092019160209182019101610bdf565b518151600019602094850361010090810a820192831692199390931691909117909252949092019687526040805197889003820188208054601f6002600183161590980290950116959095049283018290048202880182019052818752929450925050830182828015610cb25780601f10610c8757610100808354040283529160200191610cb2565b820191906000526020600020905b815481529060010190602001808311610c9557829003601f168201915b50505050509050919050565b8151610cc957600080fd5b8051610cd457600080fd5b806002836040518082805190602001908083835b60208310610d075780518252601f199092019160209182019101610ce8565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381019093208451610d48959194919091019250905061169c565b507fecfe12f3841a41927be024305b874a40bc8f83a654e5792e1e6876dddb6ae8648282604051808060200180602001838103835285818151815260200191508051906020019080838360005b83811015610dad578181015183820152602001610d95565b50505050905090810190601f168015610dda5780820380516001836020036101000a031916815260200191505b50838103825284518152845160209182019186019080838360005b83811015610e0d578181015183820152602001610df5565b50505050905090810190601f168015610e3a5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a15050565b8151610e5857600080fd5b8051610e6357600080fd5b806000836040518082805190602001908083835b60208310610e965780518252601f199092019160209182019101610e77565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381019093208451610ed7959194919091019250905061169c565b507f40864443127272df4def7ce0f1a6e85eb291d7dd41320921a5f5b8cd39d2986482826040518080602001806020018381038352858181518152602001915080519060200190808383600083811015610dad578181015183820152602001610d95565b8151610f4657600080fd5b8051610f5157600080fd5b806003836040518082805190602001908083835b60208310610f845780518252601f199092019160209182019101610f65565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381019093208451610fc5959194919091019250905061169c565b507f0ee7f9696ca518c1a77dbaa136bbbcfbda3da315a4248a97e7ad096c6ee9023982826040518080602001806020018381038352858181518152602001915080519060200190808383600083811015610dad578181015183820152602001610d95565b835161103457600080fd5b825161103f57600080fd5b815161104a57600080fd5b8061105457600080fd5b6040518060600160405280848152602001838152602001828152506004856040518082805190602001908083835b602083106110a15780518252601f199092019160209182019101611082565b51815160209384036101000a60001901801990921691161790529201948552506040519384900381019093208451805191946110e29450859350019061169c565b5060208281015180516110fb926001850192019061169c565b50604082015181600201559050507f539fe4105f626a394482b39ad5f0bbe9ce36477451037d3504939674f0a50a378484848460405180806020018060200180602001858152602001848103845288818151815260200191508051906020019080838360005b83811015611179578181015183820152602001611161565b50505050905090810190601f1680156111a65780820380516001836020036101000a031916815260200191505b50848103835287518152875160209182019189019080838360005b838110156111d95781810151838201526020016111c1565b50505050905090810190601f1680156112065780820380516001836020036101000a031916815260200191505b50848103825286518152865160209182019188019080838360005b83811015611239578181015183820152602001611221565b50505050905090810190601f1680156112665780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a150505050565b606081516000141561128f57600080fd5b60018260405180828051906020019080838360208310610bfe5780518252601f199092019160209182019101610bdf565b81516112cb57600080fd5b80516112d657600080fd5b806001836040518082805190602001908083835b602083106113095780518252601f1990920191602091820191016112ea565b51815160209384036101000a6000190180199092169116179052920194855250604051938490038101909320845161134a959194919091019250905061169c565b507f98de1647be6ada8239cad39d6dd45ff1ec584ae98c0a679eb8218f6e79efd26982826040518080602001806020018381038352858181518152602001915080519060200190808383600083811015610dad578181015183820152602001610d95565b60608060008351600014156113c257600080fd5b60606004856040518082805190602001908083835b602083106113f65780518252601f1990920191602091820191016113d7565b518151600019602094850361010090810a820192831692199390931691909117909252949092019687526040805197889003820188208054601f60026001831615909802909501169590950492830182900482028801820190528187529294509250508301828280156114aa5780601f1061147f576101008083540402835291602001916114aa565b820191906000526020600020905b81548152906001019060200180831161148d57829003601f168201915b5050505050905060606004866040518082805190602001908083835b602083106114e55780518252601f1990920191602091820191016114c6565b518151600019602094850361010090810a8201928316921993909316919091179092529490920196875260408051978890038201882060019081018054601f6002938216159098029095019094160494850182900482028801820190528387529094509192505083018282801561159d5780601f106115725761010080835404028352916020019161159d565b820191906000526020600020905b81548152906001019060200180831161158057829003601f168201915b5050505050905060006004876040518082805190602001908083835b602083106115d85780518252601f1990920191602091820191016115b9565b51815160209384036101000a6000190180199092169116179052920194855250604051938490030190922060020154949993985093965091945050505050565b606081516000141561162957600080fd5b60038260405180828051906020019080838360208310610bfe5780518252601f199092019160209182019101610bdf565b606081516000141561166b57600080fd5b60028260405180828051906020019080838360208310610bfe5780518252601f199092019160209182019101610bdf565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826116d25760008555611718565b82601f106116eb57805160ff1916838001178555611718565b82800160010185558215611718579182015b828111156117185782518255916020019190600101906116fd565b50611724929150611728565b5090565b5b80821115611724576000815560010161172956fea26469706673582212205f233f86dc2ea138f69e0a68e7b4add8023b3e761876fcd13b11a77af2113df864736f6c63430007040033"

// DeployDataStorage deploys a new Ethereum contract, binding an instance of DataStorage to it.
func DeployDataStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DataStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(DataStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DataStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DataStorage{DataStorageCaller: DataStorageCaller{contract: contract}, DataStorageTransactor: DataStorageTransactor{contract: contract}, DataStorageFilterer: DataStorageFilterer{contract: contract}}, nil
}

// DataStorage is an auto generated Go binding around an Ethereum contract.
type DataStorage struct {
	DataStorageCaller     // Read-only binding to the contract
	DataStorageTransactor // Write-only binding to the contract
	DataStorageFilterer   // Log filterer for contract events
}

// DataStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataStorageSession struct {
	Contract     *DataStorage      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataStorageCallerSession struct {
	Contract *DataStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DataStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataStorageTransactorSession struct {
	Contract     *DataStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DataStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataStorageRaw struct {
	Contract *DataStorage // Generic contract binding to access the raw methods on
}

// DataStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataStorageCallerRaw struct {
	Contract *DataStorageCaller // Generic read-only contract binding to access the raw methods on
}

// DataStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataStorageTransactorRaw struct {
	Contract *DataStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataStorage creates a new instance of DataStorage, bound to a specific deployed contract.
func NewDataStorage(address common.Address, backend bind.ContractBackend) (*DataStorage, error) {
	contract, err := bindDataStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataStorage{DataStorageCaller: DataStorageCaller{contract: contract}, DataStorageTransactor: DataStorageTransactor{contract: contract}, DataStorageFilterer: DataStorageFilterer{contract: contract}}, nil
}

// NewDataStorageCaller creates a new read-only instance of DataStorage, bound to a specific deployed contract.
func NewDataStorageCaller(address common.Address, caller bind.ContractCaller) (*DataStorageCaller, error) {
	contract, err := bindDataStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataStorageCaller{contract: contract}, nil
}

// NewDataStorageTransactor creates a new write-only instance of DataStorage, bound to a specific deployed contract.
func NewDataStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*DataStorageTransactor, error) {
	contract, err := bindDataStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataStorageTransactor{contract: contract}, nil
}

// NewDataStorageFilterer creates a new log filterer instance of DataStorage, bound to a specific deployed contract.
func NewDataStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*DataStorageFilterer, error) {
	contract, err := bindDataStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataStorageFilterer{contract: contract}, nil
}

// bindDataStorage binds a generic wrapper to an already deployed contract.
func bindDataStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataStorage *DataStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataStorage.Contract.DataStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataStorage *DataStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataStorage.Contract.DataStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataStorage *DataStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataStorage.Contract.DataStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataStorage *DataStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataStorage *DataStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataStorage *DataStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataStorage.Contract.contract.Transact(opts, method, params...)
}

// GetDidChainAddr is a free data retrieval call binding the contract method 0x09b4e1a2.
//
// Solidity: function getDidChainAddr(string did) view returns(string)
func (_DataStorage *DataStorageCaller) GetDidChainAddr(opts *bind.CallOpts, did string) (string, error) {
	var out []interface{}
	err := _DataStorage.contract.Call(opts, &out, "getDidChainAddr", did)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetDidChainAddr is a free data retrieval call binding the contract method 0x09b4e1a2.
//
// Solidity: function getDidChainAddr(string did) view returns(string)
func (_DataStorage *DataStorageSession) GetDidChainAddr(did string) (string, error) {
	return _DataStorage.Contract.GetDidChainAddr(&_DataStorage.CallOpts, did)
}

// GetDidChainAddr is a free data retrieval call binding the contract method 0x09b4e1a2.
//
// Solidity: function getDidChainAddr(string did) view returns(string)
func (_DataStorage *DataStorageCallerSession) GetDidChainAddr(did string) (string, error) {
	return _DataStorage.Contract.GetDidChainAddr(&_DataStorage.CallOpts, did)
}

// GetDidClaimHash is a free data retrieval call binding the contract method 0x78ff3fe8.
//
// Solidity: function getDidClaimHash(string did) view returns(string)
func (_DataStorage *DataStorageCaller) GetDidClaimHash(opts *bind.CallOpts, did string) (string, error) {
	var out []interface{}
	err := _DataStorage.contract.Call(opts, &out, "getDidClaimHash", did)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetDidClaimHash is a free data retrieval call binding the contract method 0x78ff3fe8.
//
// Solidity: function getDidClaimHash(string did) view returns(string)
func (_DataStorage *DataStorageSession) GetDidClaimHash(did string) (string, error) {
	return _DataStorage.Contract.GetDidClaimHash(&_DataStorage.CallOpts, did)
}

// GetDidClaimHash is a free data retrieval call binding the contract method 0x78ff3fe8.
//
// Solidity: function getDidClaimHash(string did) view returns(string)
func (_DataStorage *DataStorageCallerSession) GetDidClaimHash(did string) (string, error) {
	return _DataStorage.Contract.GetDidClaimHash(&_DataStorage.CallOpts, did)
}

// GetDidDocumentHash is a free data retrieval call binding the contract method 0xda1bc0b7.
//
// Solidity: function getDidDocumentHash(string did) view returns(string)
func (_DataStorage *DataStorageCaller) GetDidDocumentHash(opts *bind.CallOpts, did string) (string, error) {
	var out []interface{}
	err := _DataStorage.contract.Call(opts, &out, "getDidDocumentHash", did)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetDidDocumentHash is a free data retrieval call binding the contract method 0xda1bc0b7.
//
// Solidity: function getDidDocumentHash(string did) view returns(string)
func (_DataStorage *DataStorageSession) GetDidDocumentHash(did string) (string, error) {
	return _DataStorage.Contract.GetDidDocumentHash(&_DataStorage.CallOpts, did)
}

// GetDidDocumentHash is a free data retrieval call binding the contract method 0xda1bc0b7.
//
// Solidity: function getDidDocumentHash(string did) view returns(string)
func (_DataStorage *DataStorageCallerSession) GetDidDocumentHash(did string) (string, error) {
	return _DataStorage.Contract.GetDidDocumentHash(&_DataStorage.CallOpts, did)
}

// GetDidPublicKey is a free data retrieval call binding the contract method 0xd6ca0d57.
//
// Solidity: function getDidPublicKey(string did) view returns(string)
func (_DataStorage *DataStorageCaller) GetDidPublicKey(opts *bind.CallOpts, did string) (string, error) {
	var out []interface{}
	err := _DataStorage.contract.Call(opts, &out, "getDidPublicKey", did)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetDidPublicKey is a free data retrieval call binding the contract method 0xd6ca0d57.
//
// Solidity: function getDidPublicKey(string did) view returns(string)
func (_DataStorage *DataStorageSession) GetDidPublicKey(did string) (string, error) {
	return _DataStorage.Contract.GetDidPublicKey(&_DataStorage.CallOpts, did)
}

// GetDidPublicKey is a free data retrieval call binding the contract method 0xd6ca0d57.
//
// Solidity: function getDidPublicKey(string did) view returns(string)
func (_DataStorage *DataStorageCallerSession) GetDidPublicKey(did string) (string, error) {
	return _DataStorage.Contract.GetDidPublicKey(&_DataStorage.CallOpts, did)
}

// GetTransaction is a free data retrieval call binding the contract method 0xac2ccd07.
//
// Solidity: function getTransaction(string txId) view returns(string, string, uint256)
func (_DataStorage *DataStorageCaller) GetTransaction(opts *bind.CallOpts, txId string) (string, string, *big.Int, error) {
	var out []interface{}
	err := _DataStorage.contract.Call(opts, &out, "getTransaction", txId)

	if err != nil {
		return *new(string), *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetTransaction is a free data retrieval call binding the contract method 0xac2ccd07.
//
// Solidity: function getTransaction(string txId) view returns(string, string, uint256)
func (_DataStorage *DataStorageSession) GetTransaction(txId string) (string, string, *big.Int, error) {
	return _DataStorage.Contract.GetTransaction(&_DataStorage.CallOpts, txId)
}

// GetTransaction is a free data retrieval call binding the contract method 0xac2ccd07.
//
// Solidity: function getTransaction(string txId) view returns(string, string, uint256)
func (_DataStorage *DataStorageCallerSession) GetTransaction(txId string) (string, string, *big.Int, error) {
	return _DataStorage.Contract.GetTransaction(&_DataStorage.CallOpts, txId)
}

// SetDidChainAddr is a paid mutator transaction binding the contract method 0x3e06d9e9.
//
// Solidity: function setDidChainAddr(string did, string _didChainAddr) payable returns()
func (_DataStorage *DataStorageTransactor) SetDidChainAddr(opts *bind.TransactOpts, did string, _didChainAddr string) (*types.Transaction, error) {
	return _DataStorage.contract.Transact(opts, "setDidChainAddr", did, _didChainAddr)
}

// SetDidChainAddr is a paid mutator transaction binding the contract method 0x3e06d9e9.
//
// Solidity: function setDidChainAddr(string did, string _didChainAddr) payable returns()
func (_DataStorage *DataStorageSession) SetDidChainAddr(did string, _didChainAddr string) (*types.Transaction, error) {
	return _DataStorage.Contract.SetDidChainAddr(&_DataStorage.TransactOpts, did, _didChainAddr)
}

// SetDidChainAddr is a paid mutator transaction binding the contract method 0x3e06d9e9.
//
// Solidity: function setDidChainAddr(string did, string _didChainAddr) payable returns()
func (_DataStorage *DataStorageTransactorSession) SetDidChainAddr(did string, _didChainAddr string) (*types.Transaction, error) {
	return _DataStorage.Contract.SetDidChainAddr(&_DataStorage.TransactOpts, did, _didChainAddr)
}

// SetDidClaimHash is a paid mutator transaction binding the contract method 0xa8917e05.
//
// Solidity: function setDidClaimHash(string did, string _didClaimHash) payable returns()
func (_DataStorage *DataStorageTransactor) SetDidClaimHash(opts *bind.TransactOpts, did string, _didClaimHash string) (*types.Transaction, error) {
	return _DataStorage.contract.Transact(opts, "setDidClaimHash", did, _didClaimHash)
}

// SetDidClaimHash is a paid mutator transaction binding the contract method 0xa8917e05.
//
// Solidity: function setDidClaimHash(string did, string _didClaimHash) payable returns()
func (_DataStorage *DataStorageSession) SetDidClaimHash(did string, _didClaimHash string) (*types.Transaction, error) {
	return _DataStorage.Contract.SetDidClaimHash(&_DataStorage.TransactOpts, did, _didClaimHash)
}

// SetDidClaimHash is a paid mutator transaction binding the contract method 0xa8917e05.
//
// Solidity: function setDidClaimHash(string did, string _didClaimHash) payable returns()
func (_DataStorage *DataStorageTransactorSession) SetDidClaimHash(did string, _didClaimHash string) (*types.Transaction, error) {
	return _DataStorage.Contract.SetDidClaimHash(&_DataStorage.TransactOpts, did, _didClaimHash)
}

// SetDidDocumentHash is a paid mutator transaction binding the contract method 0x1dd5ec83.
//
// Solidity: function setDidDocumentHash(string did, string _didDocumentHash) payable returns()
func (_DataStorage *DataStorageTransactor) SetDidDocumentHash(opts *bind.TransactOpts, did string, _didDocumentHash string) (*types.Transaction, error) {
	return _DataStorage.contract.Transact(opts, "setDidDocumentHash", did, _didDocumentHash)
}

// SetDidDocumentHash is a paid mutator transaction binding the contract method 0x1dd5ec83.
//
// Solidity: function setDidDocumentHash(string did, string _didDocumentHash) payable returns()
func (_DataStorage *DataStorageSession) SetDidDocumentHash(did string, _didDocumentHash string) (*types.Transaction, error) {
	return _DataStorage.Contract.SetDidDocumentHash(&_DataStorage.TransactOpts, did, _didDocumentHash)
}

// SetDidDocumentHash is a paid mutator transaction binding the contract method 0x1dd5ec83.
//
// Solidity: function setDidDocumentHash(string did, string _didDocumentHash) payable returns()
func (_DataStorage *DataStorageTransactorSession) SetDidDocumentHash(did string, _didDocumentHash string) (*types.Transaction, error) {
	return _DataStorage.Contract.SetDidDocumentHash(&_DataStorage.TransactOpts, did, _didDocumentHash)
}

// SetDidPublicKey is a paid mutator transaction binding the contract method 0x46f8acd2.
//
// Solidity: function setDidPublicKey(string did, string _didPublicKey) payable returns()
func (_DataStorage *DataStorageTransactor) SetDidPublicKey(opts *bind.TransactOpts, did string, _didPublicKey string) (*types.Transaction, error) {
	return _DataStorage.contract.Transact(opts, "setDidPublicKey", did, _didPublicKey)
}

// SetDidPublicKey is a paid mutator transaction binding the contract method 0x46f8acd2.
//
// Solidity: function setDidPublicKey(string did, string _didPublicKey) payable returns()
func (_DataStorage *DataStorageSession) SetDidPublicKey(did string, _didPublicKey string) (*types.Transaction, error) {
	return _DataStorage.Contract.SetDidPublicKey(&_DataStorage.TransactOpts, did, _didPublicKey)
}

// SetDidPublicKey is a paid mutator transaction binding the contract method 0x46f8acd2.
//
// Solidity: function setDidPublicKey(string did, string _didPublicKey) payable returns()
func (_DataStorage *DataStorageTransactorSession) SetDidPublicKey(did string, _didPublicKey string) (*types.Transaction, error) {
	return _DataStorage.Contract.SetDidPublicKey(&_DataStorage.TransactOpts, did, _didPublicKey)
}

// SetTransaction is a paid mutator transaction binding the contract method 0x5a7f8c66.
//
// Solidity: function setTransaction(string txId, string _fromAddr, string _toAddr, uint256 _amount) payable returns()
func (_DataStorage *DataStorageTransactor) SetTransaction(opts *bind.TransactOpts, txId string, _fromAddr string, _toAddr string, _amount *big.Int) (*types.Transaction, error) {
	return _DataStorage.contract.Transact(opts, "setTransaction", txId, _fromAddr, _toAddr, _amount)
}

// SetTransaction is a paid mutator transaction binding the contract method 0x5a7f8c66.
//
// Solidity: function setTransaction(string txId, string _fromAddr, string _toAddr, uint256 _amount) payable returns()
func (_DataStorage *DataStorageSession) SetTransaction(txId string, _fromAddr string, _toAddr string, _amount *big.Int) (*types.Transaction, error) {
	return _DataStorage.Contract.SetTransaction(&_DataStorage.TransactOpts, txId, _fromAddr, _toAddr, _amount)
}

// SetTransaction is a paid mutator transaction binding the contract method 0x5a7f8c66.
//
// Solidity: function setTransaction(string txId, string _fromAddr, string _toAddr, uint256 _amount) payable returns()
func (_DataStorage *DataStorageTransactorSession) SetTransaction(txId string, _fromAddr string, _toAddr string, _amount *big.Int) (*types.Transaction, error) {
	return _DataStorage.Contract.SetTransaction(&_DataStorage.TransactOpts, txId, _fromAddr, _toAddr, _amount)
}

// DataStorageLogSetDidChainAddrIterator is returned from FilterLogSetDidChainAddr and is used to iterate over the raw logs and unpacked data for LogSetDidChainAddr events raised by the DataStorage contract.
type DataStorageLogSetDidChainAddrIterator struct {
	Event *DataStorageLogSetDidChainAddr // Event containing the contract specifics and raw log

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
func (it *DataStorageLogSetDidChainAddrIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataStorageLogSetDidChainAddr)
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
		it.Event = new(DataStorageLogSetDidChainAddr)
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
func (it *DataStorageLogSetDidChainAddrIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataStorageLogSetDidChainAddrIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataStorageLogSetDidChainAddr represents a LogSetDidChainAddr event raised by the DataStorage contract.
type DataStorageLogSetDidChainAddr struct {
	Did          string
	DidChainAddr string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogSetDidChainAddr is a free log retrieval operation binding the contract event 0x40864443127272df4def7ce0f1a6e85eb291d7dd41320921a5f5b8cd39d29864.
//
// Solidity: event LogSetDidChainAddr(string did, string _didChainAddr)
func (_DataStorage *DataStorageFilterer) FilterLogSetDidChainAddr(opts *bind.FilterOpts) (*DataStorageLogSetDidChainAddrIterator, error) {

	logs, sub, err := _DataStorage.contract.FilterLogs(opts, "LogSetDidChainAddr")
	if err != nil {
		return nil, err
	}
	return &DataStorageLogSetDidChainAddrIterator{contract: _DataStorage.contract, event: "LogSetDidChainAddr", logs: logs, sub: sub}, nil
}

// WatchLogSetDidChainAddr is a free log subscription operation binding the contract event 0x40864443127272df4def7ce0f1a6e85eb291d7dd41320921a5f5b8cd39d29864.
//
// Solidity: event LogSetDidChainAddr(string did, string _didChainAddr)
func (_DataStorage *DataStorageFilterer) WatchLogSetDidChainAddr(opts *bind.WatchOpts, sink chan<- *DataStorageLogSetDidChainAddr) (event.Subscription, error) {

	logs, sub, err := _DataStorage.contract.WatchLogs(opts, "LogSetDidChainAddr")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataStorageLogSetDidChainAddr)
				if err := _DataStorage.contract.UnpackLog(event, "LogSetDidChainAddr", log); err != nil {
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

// ParseLogSetDidChainAddr is a log parse operation binding the contract event 0x40864443127272df4def7ce0f1a6e85eb291d7dd41320921a5f5b8cd39d29864.
//
// Solidity: event LogSetDidChainAddr(string did, string _didChainAddr)
func (_DataStorage *DataStorageFilterer) ParseLogSetDidChainAddr(log types.Log) (*DataStorageLogSetDidChainAddr, error) {
	event := new(DataStorageLogSetDidChainAddr)
	if err := _DataStorage.contract.UnpackLog(event, "LogSetDidChainAddr", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataStorageLogSetDidClaimHashIterator is returned from FilterLogSetDidClaimHash and is used to iterate over the raw logs and unpacked data for LogSetDidClaimHash events raised by the DataStorage contract.
type DataStorageLogSetDidClaimHashIterator struct {
	Event *DataStorageLogSetDidClaimHash // Event containing the contract specifics and raw log

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
func (it *DataStorageLogSetDidClaimHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataStorageLogSetDidClaimHash)
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
		it.Event = new(DataStorageLogSetDidClaimHash)
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
func (it *DataStorageLogSetDidClaimHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataStorageLogSetDidClaimHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataStorageLogSetDidClaimHash represents a LogSetDidClaimHash event raised by the DataStorage contract.
type DataStorageLogSetDidClaimHash struct {
	Did          string
	DidClaimHash string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogSetDidClaimHash is a free log retrieval operation binding the contract event 0x98de1647be6ada8239cad39d6dd45ff1ec584ae98c0a679eb8218f6e79efd269.
//
// Solidity: event LogSetDidClaimHash(string did, string _didClaimHash)
func (_DataStorage *DataStorageFilterer) FilterLogSetDidClaimHash(opts *bind.FilterOpts) (*DataStorageLogSetDidClaimHashIterator, error) {

	logs, sub, err := _DataStorage.contract.FilterLogs(opts, "LogSetDidClaimHash")
	if err != nil {
		return nil, err
	}
	return &DataStorageLogSetDidClaimHashIterator{contract: _DataStorage.contract, event: "LogSetDidClaimHash", logs: logs, sub: sub}, nil
}

// WatchLogSetDidClaimHash is a free log subscription operation binding the contract event 0x98de1647be6ada8239cad39d6dd45ff1ec584ae98c0a679eb8218f6e79efd269.
//
// Solidity: event LogSetDidClaimHash(string did, string _didClaimHash)
func (_DataStorage *DataStorageFilterer) WatchLogSetDidClaimHash(opts *bind.WatchOpts, sink chan<- *DataStorageLogSetDidClaimHash) (event.Subscription, error) {

	logs, sub, err := _DataStorage.contract.WatchLogs(opts, "LogSetDidClaimHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataStorageLogSetDidClaimHash)
				if err := _DataStorage.contract.UnpackLog(event, "LogSetDidClaimHash", log); err != nil {
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

// ParseLogSetDidClaimHash is a log parse operation binding the contract event 0x98de1647be6ada8239cad39d6dd45ff1ec584ae98c0a679eb8218f6e79efd269.
//
// Solidity: event LogSetDidClaimHash(string did, string _didClaimHash)
func (_DataStorage *DataStorageFilterer) ParseLogSetDidClaimHash(log types.Log) (*DataStorageLogSetDidClaimHash, error) {
	event := new(DataStorageLogSetDidClaimHash)
	if err := _DataStorage.contract.UnpackLog(event, "LogSetDidClaimHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataStorageLogSetDidDocumentHashIterator is returned from FilterLogSetDidDocumentHash and is used to iterate over the raw logs and unpacked data for LogSetDidDocumentHash events raised by the DataStorage contract.
type DataStorageLogSetDidDocumentHashIterator struct {
	Event *DataStorageLogSetDidDocumentHash // Event containing the contract specifics and raw log

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
func (it *DataStorageLogSetDidDocumentHashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataStorageLogSetDidDocumentHash)
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
		it.Event = new(DataStorageLogSetDidDocumentHash)
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
func (it *DataStorageLogSetDidDocumentHashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataStorageLogSetDidDocumentHashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataStorageLogSetDidDocumentHash represents a LogSetDidDocumentHash event raised by the DataStorage contract.
type DataStorageLogSetDidDocumentHash struct {
	Did             string
	DidDocumentHash string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogSetDidDocumentHash is a free log retrieval operation binding the contract event 0xecfe12f3841a41927be024305b874a40bc8f83a654e5792e1e6876dddb6ae864.
//
// Solidity: event LogSetDidDocumentHash(string did, string _didDocumentHash)
func (_DataStorage *DataStorageFilterer) FilterLogSetDidDocumentHash(opts *bind.FilterOpts) (*DataStorageLogSetDidDocumentHashIterator, error) {

	logs, sub, err := _DataStorage.contract.FilterLogs(opts, "LogSetDidDocumentHash")
	if err != nil {
		return nil, err
	}
	return &DataStorageLogSetDidDocumentHashIterator{contract: _DataStorage.contract, event: "LogSetDidDocumentHash", logs: logs, sub: sub}, nil
}

// WatchLogSetDidDocumentHash is a free log subscription operation binding the contract event 0xecfe12f3841a41927be024305b874a40bc8f83a654e5792e1e6876dddb6ae864.
//
// Solidity: event LogSetDidDocumentHash(string did, string _didDocumentHash)
func (_DataStorage *DataStorageFilterer) WatchLogSetDidDocumentHash(opts *bind.WatchOpts, sink chan<- *DataStorageLogSetDidDocumentHash) (event.Subscription, error) {

	logs, sub, err := _DataStorage.contract.WatchLogs(opts, "LogSetDidDocumentHash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataStorageLogSetDidDocumentHash)
				if err := _DataStorage.contract.UnpackLog(event, "LogSetDidDocumentHash", log); err != nil {
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

// ParseLogSetDidDocumentHash is a log parse operation binding the contract event 0xecfe12f3841a41927be024305b874a40bc8f83a654e5792e1e6876dddb6ae864.
//
// Solidity: event LogSetDidDocumentHash(string did, string _didDocumentHash)
func (_DataStorage *DataStorageFilterer) ParseLogSetDidDocumentHash(log types.Log) (*DataStorageLogSetDidDocumentHash, error) {
	event := new(DataStorageLogSetDidDocumentHash)
	if err := _DataStorage.contract.UnpackLog(event, "LogSetDidDocumentHash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataStorageLogSetDidPublicKeyIterator is returned from FilterLogSetDidPublicKey and is used to iterate over the raw logs and unpacked data for LogSetDidPublicKey events raised by the DataStorage contract.
type DataStorageLogSetDidPublicKeyIterator struct {
	Event *DataStorageLogSetDidPublicKey // Event containing the contract specifics and raw log

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
func (it *DataStorageLogSetDidPublicKeyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataStorageLogSetDidPublicKey)
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
		it.Event = new(DataStorageLogSetDidPublicKey)
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
func (it *DataStorageLogSetDidPublicKeyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataStorageLogSetDidPublicKeyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataStorageLogSetDidPublicKey represents a LogSetDidPublicKey event raised by the DataStorage contract.
type DataStorageLogSetDidPublicKey struct {
	Did          string
	DidPublicKey string
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogSetDidPublicKey is a free log retrieval operation binding the contract event 0x0ee7f9696ca518c1a77dbaa136bbbcfbda3da315a4248a97e7ad096c6ee90239.
//
// Solidity: event LogSetDidPublicKey(string did, string _didPublicKey)
func (_DataStorage *DataStorageFilterer) FilterLogSetDidPublicKey(opts *bind.FilterOpts) (*DataStorageLogSetDidPublicKeyIterator, error) {

	logs, sub, err := _DataStorage.contract.FilterLogs(opts, "LogSetDidPublicKey")
	if err != nil {
		return nil, err
	}
	return &DataStorageLogSetDidPublicKeyIterator{contract: _DataStorage.contract, event: "LogSetDidPublicKey", logs: logs, sub: sub}, nil
}

// WatchLogSetDidPublicKey is a free log subscription operation binding the contract event 0x0ee7f9696ca518c1a77dbaa136bbbcfbda3da315a4248a97e7ad096c6ee90239.
//
// Solidity: event LogSetDidPublicKey(string did, string _didPublicKey)
func (_DataStorage *DataStorageFilterer) WatchLogSetDidPublicKey(opts *bind.WatchOpts, sink chan<- *DataStorageLogSetDidPublicKey) (event.Subscription, error) {

	logs, sub, err := _DataStorage.contract.WatchLogs(opts, "LogSetDidPublicKey")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataStorageLogSetDidPublicKey)
				if err := _DataStorage.contract.UnpackLog(event, "LogSetDidPublicKey", log); err != nil {
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

// ParseLogSetDidPublicKey is a log parse operation binding the contract event 0x0ee7f9696ca518c1a77dbaa136bbbcfbda3da315a4248a97e7ad096c6ee90239.
//
// Solidity: event LogSetDidPublicKey(string did, string _didPublicKey)
func (_DataStorage *DataStorageFilterer) ParseLogSetDidPublicKey(log types.Log) (*DataStorageLogSetDidPublicKey, error) {
	event := new(DataStorageLogSetDidPublicKey)
	if err := _DataStorage.contract.UnpackLog(event, "LogSetDidPublicKey", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DataStorageLogSetTransactionIterator is returned from FilterLogSetTransaction and is used to iterate over the raw logs and unpacked data for LogSetTransaction events raised by the DataStorage contract.
type DataStorageLogSetTransactionIterator struct {
	Event *DataStorageLogSetTransaction // Event containing the contract specifics and raw log

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
func (it *DataStorageLogSetTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DataStorageLogSetTransaction)
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
		it.Event = new(DataStorageLogSetTransaction)
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
func (it *DataStorageLogSetTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DataStorageLogSetTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DataStorageLogSetTransaction represents a LogSetTransaction event raised by the DataStorage contract.
type DataStorageLogSetTransaction struct {
	TxId     string
	FromAddr string
	ToAddr   string
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogSetTransaction is a free log retrieval operation binding the contract event 0x539fe4105f626a394482b39ad5f0bbe9ce36477451037d3504939674f0a50a37.
//
// Solidity: event LogSetTransaction(string txId, string fromAddr, string toAddr, uint256 amount)
func (_DataStorage *DataStorageFilterer) FilterLogSetTransaction(opts *bind.FilterOpts) (*DataStorageLogSetTransactionIterator, error) {

	logs, sub, err := _DataStorage.contract.FilterLogs(opts, "LogSetTransaction")
	if err != nil {
		return nil, err
	}
	return &DataStorageLogSetTransactionIterator{contract: _DataStorage.contract, event: "LogSetTransaction", logs: logs, sub: sub}, nil
}

// WatchLogSetTransaction is a free log subscription operation binding the contract event 0x539fe4105f626a394482b39ad5f0bbe9ce36477451037d3504939674f0a50a37.
//
// Solidity: event LogSetTransaction(string txId, string fromAddr, string toAddr, uint256 amount)
func (_DataStorage *DataStorageFilterer) WatchLogSetTransaction(opts *bind.WatchOpts, sink chan<- *DataStorageLogSetTransaction) (event.Subscription, error) {

	logs, sub, err := _DataStorage.contract.WatchLogs(opts, "LogSetTransaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DataStorageLogSetTransaction)
				if err := _DataStorage.contract.UnpackLog(event, "LogSetTransaction", log); err != nil {
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

// ParseLogSetTransaction is a log parse operation binding the contract event 0x539fe4105f626a394482b39ad5f0bbe9ce36477451037d3504939674f0a50a37.
//
// Solidity: event LogSetTransaction(string txId, string fromAddr, string toAddr, uint256 amount)
func (_DataStorage *DataStorageFilterer) ParseLogSetTransaction(log types.Log) (*DataStorageLogSetTransaction, error) {
	event := new(DataStorageLogSetTransaction)
	if err := _DataStorage.contract.UnpackLog(event, "LogSetTransaction", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
