// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pair

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

// PancakeABI is the input ABI used to generate the binding from.
const PancakeABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve0\",\"type\":\"uint112\"},{\"indexed\":false,\"internalType\":\"uint112\",\"name\":\"reserve1\",\"type\":\"uint112\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MINIMUM_LIQUIDITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PERMIT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint112\",\"name\":\"_reserve0\",\"type\":\"uint112\"},{\"internalType\":\"uint112\",\"name\":\"_reserve1\",\"type\":\"uint112\"},{\"internalType\":\"uint32\",\"name\":\"_blockTimestampLast\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token1\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price0CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"price1CumulativeLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"skim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"sync\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Pancake is an auto generated Go binding around an Ethereum contract.
type Pancake struct {
	PancakeCaller     // Read-only binding to the contract
	PancakeTransactor // Write-only binding to the contract
	PancakeFilterer   // Log filterer for contract events
}

// PancakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PancakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PancakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PancakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PancakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PancakeSession struct {
	Contract     *Pancake          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PancakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PancakeCallerSession struct {
	Contract *PancakeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PancakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PancakeTransactorSession struct {
	Contract     *PancakeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PancakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PancakeRaw struct {
	Contract *Pancake // Generic contract binding to access the raw methods on
}

// PancakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PancakeCallerRaw struct {
	Contract *PancakeCaller // Generic read-only contract binding to access the raw methods on
}

// PancakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PancakeTransactorRaw struct {
	Contract *PancakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPancake creates a new instance of Pancake, bound to a specific deployed contract.
func NewPancake(address common.Address, backend bind.ContractBackend) (*Pancake, error) {
	contract, err := bindPancake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pancake{PancakeCaller: PancakeCaller{contract: contract}, PancakeTransactor: PancakeTransactor{contract: contract}, PancakeFilterer: PancakeFilterer{contract: contract}}, nil
}

// NewPancakeCaller creates a new read-only instance of Pancake, bound to a specific deployed contract.
func NewPancakeCaller(address common.Address, caller bind.ContractCaller) (*PancakeCaller, error) {
	contract, err := bindPancake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeCaller{contract: contract}, nil
}

// NewPancakeTransactor creates a new write-only instance of Pancake, bound to a specific deployed contract.
func NewPancakeTransactor(address common.Address, transactor bind.ContractTransactor) (*PancakeTransactor, error) {
	contract, err := bindPancake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PancakeTransactor{contract: contract}, nil
}

// NewPancakeFilterer creates a new log filterer instance of Pancake, bound to a specific deployed contract.
func NewPancakeFilterer(address common.Address, filterer bind.ContractFilterer) (*PancakeFilterer, error) {
	contract, err := bindPancake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PancakeFilterer{contract: contract}, nil
}

// bindPancake binds a generic wrapper to an already deployed contract.
func bindPancake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PancakeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pancake *PancakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pancake.Contract.PancakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pancake *PancakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pancake.Contract.PancakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pancake *PancakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pancake.Contract.PancakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pancake *PancakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pancake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pancake *PancakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pancake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pancake *PancakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pancake.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Pancake *PancakeCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Pancake *PancakeSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Pancake.Contract.DOMAINSEPARATOR(&_Pancake.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Pancake *PancakeCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Pancake.Contract.DOMAINSEPARATOR(&_Pancake.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Pancake *PancakeCaller) MINIMUMLIQUIDITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "MINIMUM_LIQUIDITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Pancake *PancakeSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _Pancake.Contract.MINIMUMLIQUIDITY(&_Pancake.CallOpts)
}

// MINIMUMLIQUIDITY is a free data retrieval call binding the contract method 0xba9a7a56.
//
// Solidity: function MINIMUM_LIQUIDITY() view returns(uint256)
func (_Pancake *PancakeCallerSession) MINIMUMLIQUIDITY() (*big.Int, error) {
	return _Pancake.Contract.MINIMUMLIQUIDITY(&_Pancake.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Pancake *PancakeCaller) PERMITTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "PERMIT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Pancake *PancakeSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Pancake.Contract.PERMITTYPEHASH(&_Pancake.CallOpts)
}

// PERMITTYPEHASH is a free data retrieval call binding the contract method 0x30adf81f.
//
// Solidity: function PERMIT_TYPEHASH() view returns(bytes32)
func (_Pancake *PancakeCallerSession) PERMITTYPEHASH() ([32]byte, error) {
	return _Pancake.Contract.PERMITTYPEHASH(&_Pancake.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Pancake *PancakeCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Pancake *PancakeSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Pancake.Contract.Allowance(&_Pancake.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_Pancake *PancakeCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Pancake.Contract.Allowance(&_Pancake.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Pancake *PancakeCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Pancake *PancakeSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Pancake.Contract.BalanceOf(&_Pancake.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_Pancake *PancakeCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Pancake.Contract.BalanceOf(&_Pancake.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pancake *PancakeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pancake *PancakeSession) Decimals() (uint8, error) {
	return _Pancake.Contract.Decimals(&_Pancake.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Pancake *PancakeCallerSession) Decimals() (uint8, error) {
	return _Pancake.Contract.Decimals(&_Pancake.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pancake *PancakeCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pancake *PancakeSession) Factory() (common.Address, error) {
	return _Pancake.Contract.Factory(&_Pancake.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pancake *PancakeCallerSession) Factory() (common.Address, error) {
	return _Pancake.Contract.Factory(&_Pancake.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_Pancake *PancakeCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0           *big.Int
		Reserve1           *big.Int
		BlockTimestampLast uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockTimestampLast = *abi.ConvertType(out[2], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_Pancake *PancakeSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _Pancake.Contract.GetReserves(&_Pancake.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint112 _reserve0, uint112 _reserve1, uint32 _blockTimestampLast)
func (_Pancake *PancakeCallerSession) GetReserves() (struct {
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}, error) {
	return _Pancake.Contract.GetReserves(&_Pancake.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Pancake *PancakeCaller) KLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "kLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Pancake *PancakeSession) KLast() (*big.Int, error) {
	return _Pancake.Contract.KLast(&_Pancake.CallOpts)
}

// KLast is a free data retrieval call binding the contract method 0x7464fc3d.
//
// Solidity: function kLast() view returns(uint256)
func (_Pancake *PancakeCallerSession) KLast() (*big.Int, error) {
	return _Pancake.Contract.KLast(&_Pancake.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pancake *PancakeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pancake *PancakeSession) Name() (string, error) {
	return _Pancake.Contract.Name(&_Pancake.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Pancake *PancakeCallerSession) Name() (string, error) {
	return _Pancake.Contract.Name(&_Pancake.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Pancake *PancakeCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Pancake *PancakeSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Pancake.Contract.Nonces(&_Pancake.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Pancake *PancakeCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Pancake.Contract.Nonces(&_Pancake.CallOpts, arg0)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Pancake *PancakeCaller) Price0CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "price0CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Pancake *PancakeSession) Price0CumulativeLast() (*big.Int, error) {
	return _Pancake.Contract.Price0CumulativeLast(&_Pancake.CallOpts)
}

// Price0CumulativeLast is a free data retrieval call binding the contract method 0x5909c0d5.
//
// Solidity: function price0CumulativeLast() view returns(uint256)
func (_Pancake *PancakeCallerSession) Price0CumulativeLast() (*big.Int, error) {
	return _Pancake.Contract.Price0CumulativeLast(&_Pancake.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Pancake *PancakeCaller) Price1CumulativeLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "price1CumulativeLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Pancake *PancakeSession) Price1CumulativeLast() (*big.Int, error) {
	return _Pancake.Contract.Price1CumulativeLast(&_Pancake.CallOpts)
}

// Price1CumulativeLast is a free data retrieval call binding the contract method 0x5a3d5493.
//
// Solidity: function price1CumulativeLast() view returns(uint256)
func (_Pancake *PancakeCallerSession) Price1CumulativeLast() (*big.Int, error) {
	return _Pancake.Contract.Price1CumulativeLast(&_Pancake.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pancake *PancakeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pancake *PancakeSession) Symbol() (string, error) {
	return _Pancake.Contract.Symbol(&_Pancake.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Pancake *PancakeCallerSession) Symbol() (string, error) {
	return _Pancake.Contract.Symbol(&_Pancake.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pancake *PancakeCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pancake *PancakeSession) Token0() (common.Address, error) {
	return _Pancake.Contract.Token0(&_Pancake.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_Pancake *PancakeCallerSession) Token0() (common.Address, error) {
	return _Pancake.Contract.Token0(&_Pancake.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pancake *PancakeCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pancake *PancakeSession) Token1() (common.Address, error) {
	return _Pancake.Contract.Token1(&_Pancake.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_Pancake *PancakeCallerSession) Token1() (common.Address, error) {
	return _Pancake.Contract.Token1(&_Pancake.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pancake *PancakeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pancake.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pancake *PancakeSession) TotalSupply() (*big.Int, error) {
	return _Pancake.Contract.TotalSupply(&_Pancake.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Pancake *PancakeCallerSession) TotalSupply() (*big.Int, error) {
	return _Pancake.Contract.TotalSupply(&_Pancake.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Pancake *PancakeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Pancake *PancakeSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.Approve(&_Pancake.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Pancake *PancakeTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.Approve(&_Pancake.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Pancake *PancakeTransactor) Burn(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "burn", to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Pancake *PancakeSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Pancake.Contract.Burn(&_Pancake.TransactOpts, to)
}

// Burn is a paid mutator transaction binding the contract method 0x89afcb44.
//
// Solidity: function burn(address to) returns(uint256 amount0, uint256 amount1)
func (_Pancake *PancakeTransactorSession) Burn(to common.Address) (*types.Transaction, error) {
	return _Pancake.Contract.Burn(&_Pancake.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_Pancake *PancakeTransactor) Initialize(opts *bind.TransactOpts, _token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "initialize", _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_Pancake *PancakeSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _Pancake.Contract.Initialize(&_Pancake.TransactOpts, _token0, _token1)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _token0, address _token1) returns()
func (_Pancake *PancakeTransactorSession) Initialize(_token0 common.Address, _token1 common.Address) (*types.Transaction, error) {
	return _Pancake.Contract.Initialize(&_Pancake.TransactOpts, _token0, _token1)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pancake *PancakeTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pancake *PancakeSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Pancake.Contract.Mint(&_Pancake.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_Pancake *PancakeTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _Pancake.Contract.Mint(&_Pancake.TransactOpts, to)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pancake *PancakeTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pancake *PancakeSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pancake.Contract.Permit(&_Pancake.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Pancake *PancakeTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Pancake.Contract.Permit(&_Pancake.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Pancake *PancakeTransactor) Skim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "skim", to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Pancake *PancakeSession) Skim(to common.Address) (*types.Transaction, error) {
	return _Pancake.Contract.Skim(&_Pancake.TransactOpts, to)
}

// Skim is a paid mutator transaction binding the contract method 0xbc25cf77.
//
// Solidity: function skim(address to) returns()
func (_Pancake *PancakeTransactorSession) Skim(to common.Address) (*types.Transaction, error) {
	return _Pancake.Contract.Skim(&_Pancake.TransactOpts, to)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Pancake *PancakeTransactor) Swap(opts *bind.TransactOpts, amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "swap", amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Pancake *PancakeSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Pancake.Contract.Swap(&_Pancake.TransactOpts, amount0Out, amount1Out, to, data)
}

// Swap is a paid mutator transaction binding the contract method 0x022c0d9f.
//
// Solidity: function swap(uint256 amount0Out, uint256 amount1Out, address to, bytes data) returns()
func (_Pancake *PancakeTransactorSession) Swap(amount0Out *big.Int, amount1Out *big.Int, to common.Address, data []byte) (*types.Transaction, error) {
	return _Pancake.Contract.Swap(&_Pancake.TransactOpts, amount0Out, amount1Out, to, data)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Pancake *PancakeTransactor) Sync(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "sync")
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Pancake *PancakeSession) Sync() (*types.Transaction, error) {
	return _Pancake.Contract.Sync(&_Pancake.TransactOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfff6cae9.
//
// Solidity: function sync() returns()
func (_Pancake *PancakeTransactorSession) Sync() (*types.Transaction, error) {
	return _Pancake.Contract.Sync(&_Pancake.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Pancake *PancakeTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Pancake *PancakeSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.Transfer(&_Pancake.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Pancake *PancakeTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.Transfer(&_Pancake.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Pancake *PancakeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pancake.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Pancake *PancakeSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.TransferFrom(&_Pancake.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Pancake *PancakeTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Pancake.Contract.TransferFrom(&_Pancake.TransactOpts, from, to, value)
}

// PancakeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Pancake contract.
type PancakeApprovalIterator struct {
	Event *PancakeApproval // Event containing the contract specifics and raw log

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
func (it *PancakeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeApproval)
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
		it.Event = new(PancakeApproval)
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
func (it *PancakeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeApproval represents a Approval event raised by the Pancake contract.
type PancakeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pancake *PancakeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PancakeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pancake.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PancakeApprovalIterator{contract: _Pancake.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pancake *PancakeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PancakeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Pancake.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeApproval)
				if err := _Pancake.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Pancake *PancakeFilterer) ParseApproval(log types.Log) (*PancakeApproval, error) {
	event := new(PancakeApproval)
	if err := _Pancake.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Pancake contract.
type PancakeBurnIterator struct {
	Event *PancakeBurn // Event containing the contract specifics and raw log

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
func (it *PancakeBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeBurn)
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
		it.Event = new(PancakeBurn)
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
func (it *PancakeBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeBurn represents a Burn event raised by the Pancake contract.
type PancakeBurn struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	To      common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Pancake *PancakeFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*PancakeBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pancake.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PancakeBurnIterator{contract: _Pancake.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Pancake *PancakeFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *PancakeBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pancake.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeBurn)
				if err := _Pancake.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xdccd412f0b1252819cb1fd330b93224ca42612892bb3f4f789976e6d81936496.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, address indexed to)
func (_Pancake *PancakeFilterer) ParseBurn(log types.Log) (*PancakeBurn, error) {
	event := new(PancakeBurn)
	if err := _Pancake.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the Pancake contract.
type PancakeMintIterator struct {
	Event *PancakeMint // Event containing the contract specifics and raw log

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
func (it *PancakeMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeMint)
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
		it.Event = new(PancakeMint)
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
func (it *PancakeMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeMint represents a Mint event raised by the Pancake contract.
type PancakeMint struct {
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Pancake *PancakeFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address) (*PancakeMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pancake.contract.FilterLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return &PancakeMintIterator{contract: _Pancake.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Pancake *PancakeFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PancakeMint, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pancake.contract.WatchLogs(opts, "Mint", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeMint)
				if err := _Pancake.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1)
func (_Pancake *PancakeFilterer) ParseMint(log types.Log) (*PancakeMint, error) {
	event := new(PancakeMint)
	if err := _Pancake.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Pancake contract.
type PancakeSwapIterator struct {
	Event *PancakeSwap // Event containing the contract specifics and raw log

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
func (it *PancakeSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeSwap)
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
		it.Event = new(PancakeSwap)
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
func (it *PancakeSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeSwap represents a Swap event raised by the Pancake contract.
type PancakeSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Pancake *PancakeFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*PancakeSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pancake.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PancakeSwapIterator{contract: _Pancake.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Pancake *PancakeFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *PancakeSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pancake.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeSwap)
				if err := _Pancake.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_Pancake *PancakeFilterer) ParseSwap(log types.Log) (*PancakeSwap, error) {
	event := new(PancakeSwap)
	if err := _Pancake.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the Pancake contract.
type PancakeSyncIterator struct {
	Event *PancakeSync // Event containing the contract specifics and raw log

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
func (it *PancakeSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeSync)
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
		it.Event = new(PancakeSync)
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
func (it *PancakeSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeSync represents a Sync event raised by the Pancake contract.
type PancakeSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Pancake *PancakeFilterer) FilterSync(opts *bind.FilterOpts) (*PancakeSyncIterator, error) {

	logs, sub, err := _Pancake.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &PancakeSyncIterator{contract: _Pancake.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Pancake *PancakeFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *PancakeSync) (event.Subscription, error) {

	logs, sub, err := _Pancake.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeSync)
				if err := _Pancake.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0x1c411e9a96e071241c2f21f7726b17ae89e3cab4c78be50e062b03a9fffbbad1.
//
// Solidity: event Sync(uint112 reserve0, uint112 reserve1)
func (_Pancake *PancakeFilterer) ParseSync(log types.Log) (*PancakeSync, error) {
	event := new(PancakeSync)
	if err := _Pancake.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PancakeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Pancake contract.
type PancakeTransferIterator struct {
	Event *PancakeTransfer // Event containing the contract specifics and raw log

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
func (it *PancakeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PancakeTransfer)
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
		it.Event = new(PancakeTransfer)
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
func (it *PancakeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PancakeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PancakeTransfer represents a Transfer event raised by the Pancake contract.
type PancakeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pancake *PancakeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PancakeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pancake.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PancakeTransferIterator{contract: _Pancake.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pancake *PancakeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PancakeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Pancake.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PancakeTransfer)
				if err := _Pancake.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Pancake *PancakeFilterer) ParseTransfer(log types.Log) (*PancakeTransfer, error) {
	event := new(PancakeTransfer)
	if err := _Pancake.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
