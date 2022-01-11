// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dfc

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

// DfcMetaData contains all meta data concerning the Dfc contract.
var DfcMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Harvest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Referral\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_term\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKINGRETURN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"farmSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"harvest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resumeStaking\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardReceived\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"referer\",\"type\":\"address\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalDistributed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalEarned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"withdrawAltcoinTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DfcABI is the input ABI used to generate the binding from.
// Deprecated: Use DfcMetaData.ABI instead.
var DfcABI = DfcMetaData.ABI

// Dfc is an auto generated Go binding around an Ethereum contract.
type Dfc struct {
	DfcCaller     // Read-only binding to the contract
	DfcTransactor // Write-only binding to the contract
	DfcFilterer   // Log filterer for contract events
}

// DfcCaller is an auto generated read-only Go binding around an Ethereum contract.
type DfcCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DfcTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DfcTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DfcFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DfcFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DfcSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DfcSession struct {
	Contract     *Dfc              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DfcCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DfcCallerSession struct {
	Contract *DfcCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DfcTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DfcTransactorSession struct {
	Contract     *DfcTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DfcRaw is an auto generated low-level Go binding around an Ethereum contract.
type DfcRaw struct {
	Contract *Dfc // Generic contract binding to access the raw methods on
}

// DfcCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DfcCallerRaw struct {
	Contract *DfcCaller // Generic read-only contract binding to access the raw methods on
}

// DfcTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DfcTransactorRaw struct {
	Contract *DfcTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDfc creates a new instance of Dfc, bound to a specific deployed contract.
func NewDfc(address common.Address, backend bind.ContractBackend) (*Dfc, error) {
	contract, err := bindDfc(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dfc{DfcCaller: DfcCaller{contract: contract}, DfcTransactor: DfcTransactor{contract: contract}, DfcFilterer: DfcFilterer{contract: contract}}, nil
}

// NewDfcCaller creates a new read-only instance of Dfc, bound to a specific deployed contract.
func NewDfcCaller(address common.Address, caller bind.ContractCaller) (*DfcCaller, error) {
	contract, err := bindDfc(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DfcCaller{contract: contract}, nil
}

// NewDfcTransactor creates a new write-only instance of Dfc, bound to a specific deployed contract.
func NewDfcTransactor(address common.Address, transactor bind.ContractTransactor) (*DfcTransactor, error) {
	contract, err := bindDfc(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DfcTransactor{contract: contract}, nil
}

// NewDfcFilterer creates a new log filterer instance of Dfc, bound to a specific deployed contract.
func NewDfcFilterer(address common.Address, filterer bind.ContractFilterer) (*DfcFilterer, error) {
	contract, err := bindDfc(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DfcFilterer{contract: contract}, nil
}

// bindDfc binds a generic wrapper to an already deployed contract.
func bindDfc(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DfcABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dfc *DfcRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dfc.Contract.DfcCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dfc *DfcRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dfc.Contract.DfcTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dfc *DfcRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dfc.Contract.DfcTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dfc *DfcCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dfc.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dfc *DfcTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dfc.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dfc *DfcTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dfc.Contract.contract.Transact(opts, method, params...)
}

// ADAY is a free data retrieval call binding the contract method 0x96db733a.
//
// Solidity: function ADAY() view returns(uint256)
func (_Dfc *DfcCaller) ADAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dfc.contract.Call(opts, &out, "ADAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ADAY is a free data retrieval call binding the contract method 0x96db733a.
//
// Solidity: function ADAY() view returns(uint256)
func (_Dfc *DfcSession) ADAY() (*big.Int, error) {
	return _Dfc.Contract.ADAY(&_Dfc.CallOpts)
}

// ADAY is a free data retrieval call binding the contract method 0x96db733a.
//
// Solidity: function ADAY() view returns(uint256)
func (_Dfc *DfcCallerSession) ADAY() (*big.Int, error) {
	return _Dfc.Contract.ADAY(&_Dfc.CallOpts)
}

// STAKINGRETURN is a free data retrieval call binding the contract method 0x51741a54.
//
// Solidity: function STAKINGRETURN() view returns(uint256)
func (_Dfc *DfcCaller) STAKINGRETURN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dfc.contract.Call(opts, &out, "STAKINGRETURN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STAKINGRETURN is a free data retrieval call binding the contract method 0x51741a54.
//
// Solidity: function STAKINGRETURN() view returns(uint256)
func (_Dfc *DfcSession) STAKINGRETURN() (*big.Int, error) {
	return _Dfc.Contract.STAKINGRETURN(&_Dfc.CallOpts)
}

// STAKINGRETURN is a free data retrieval call binding the contract method 0x51741a54.
//
// Solidity: function STAKINGRETURN() view returns(uint256)
func (_Dfc *DfcCallerSession) STAKINGRETURN() (*big.Int, error) {
	return _Dfc.Contract.STAKINGRETURN(&_Dfc.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Dfc *DfcCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dfc.contract.Call(opts, &out, "allowance", _owner, _spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Dfc *DfcSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Dfc.Contract.Allowance(&_Dfc.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_Dfc *DfcCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _Dfc.Contract.Allowance(&_Dfc.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_Dfc *DfcCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dfc.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_Dfc *DfcSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Dfc.Contract.BalanceOf(&_Dfc.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_Dfc *DfcCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Dfc.Contract.BalanceOf(&_Dfc.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Dfc *DfcCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dfc.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Dfc *DfcSession) Decimals() (*big.Int, error) {
	return _Dfc.Contract.Decimals(&_Dfc.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_Dfc *DfcCallerSession) Decimals() (*big.Int, error) {
	return _Dfc.Contract.Decimals(&_Dfc.CallOpts)
}

// FarmSize is a free data retrieval call binding the contract method 0xd4d6485d.
//
// Solidity: function farmSize() view returns(uint256)
func (_Dfc *DfcCaller) FarmSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dfc.contract.Call(opts, &out, "farmSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FarmSize is a free data retrieval call binding the contract method 0xd4d6485d.
//
// Solidity: function farmSize() view returns(uint256)
func (_Dfc *DfcSession) FarmSize() (*big.Int, error) {
	return _Dfc.Contract.FarmSize(&_Dfc.CallOpts)
}

// FarmSize is a free data retrieval call binding the contract method 0xd4d6485d.
//
// Solidity: function farmSize() view returns(uint256)
func (_Dfc *DfcCallerSession) FarmSize() (*big.Int, error) {
	return _Dfc.Contract.FarmSize(&_Dfc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dfc *DfcCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dfc.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dfc *DfcSession) Name() (string, error) {
	return _Dfc.Contract.Name(&_Dfc.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Dfc *DfcCallerSession) Name() (string, error) {
	return _Dfc.Contract.Name(&_Dfc.CallOpts)
}

// RewardReceived is a free data retrieval call binding the contract method 0xaed5be17.
//
// Solidity: function rewardReceived() view returns(uint256)
func (_Dfc *DfcCaller) RewardReceived(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dfc.contract.Call(opts, &out, "rewardReceived")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardReceived is a free data retrieval call binding the contract method 0xaed5be17.
//
// Solidity: function rewardReceived() view returns(uint256)
func (_Dfc *DfcSession) RewardReceived() (*big.Int, error) {
	return _Dfc.Contract.RewardReceived(&_Dfc.CallOpts)
}

// RewardReceived is a free data retrieval call binding the contract method 0xaed5be17.
//
// Solidity: function rewardReceived() view returns(uint256)
func (_Dfc *DfcCallerSession) RewardReceived() (*big.Int, error) {
	return _Dfc.Contract.RewardReceived(&_Dfc.CallOpts)
}

// StakingPaused is a free data retrieval call binding the contract method 0xbbb781cc.
//
// Solidity: function stakingPaused() view returns(bool)
func (_Dfc *DfcCaller) StakingPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Dfc.contract.Call(opts, &out, "stakingPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// StakingPaused is a free data retrieval call binding the contract method 0xbbb781cc.
//
// Solidity: function stakingPaused() view returns(bool)
func (_Dfc *DfcSession) StakingPaused() (bool, error) {
	return _Dfc.Contract.StakingPaused(&_Dfc.CallOpts)
}

// StakingPaused is a free data retrieval call binding the contract method 0xbbb781cc.
//
// Solidity: function stakingPaused() view returns(bool)
func (_Dfc *DfcCallerSession) StakingPaused() (bool, error) {
	return _Dfc.Contract.StakingPaused(&_Dfc.CallOpts)
}

// StakingPeriod is a free data retrieval call binding the contract method 0xc03d5b47.
//
// Solidity: function stakingPeriod() view returns(uint256)
func (_Dfc *DfcCaller) StakingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dfc.contract.Call(opts, &out, "stakingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingPeriod is a free data retrieval call binding the contract method 0xc03d5b47.
//
// Solidity: function stakingPeriod() view returns(uint256)
func (_Dfc *DfcSession) StakingPeriod() (*big.Int, error) {
	return _Dfc.Contract.StakingPeriod(&_Dfc.CallOpts)
}

// StakingPeriod is a free data retrieval call binding the contract method 0xc03d5b47.
//
// Solidity: function stakingPeriod() view returns(uint256)
func (_Dfc *DfcCallerSession) StakingPeriod() (*big.Int, error) {
	return _Dfc.Contract.StakingPeriod(&_Dfc.CallOpts)
}

// StakingReward is a free data retrieval call binding the contract method 0x042249bb.
//
// Solidity: function stakingReward() view returns(uint256)
func (_Dfc *DfcCaller) StakingReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dfc.contract.Call(opts, &out, "stakingReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakingReward is a free data retrieval call binding the contract method 0x042249bb.
//
// Solidity: function stakingReward() view returns(uint256)
func (_Dfc *DfcSession) StakingReward() (*big.Int, error) {
	return _Dfc.Contract.StakingReward(&_Dfc.CallOpts)
}

// StakingReward is a free data retrieval call binding the contract method 0x042249bb.
//
// Solidity: function stakingReward() view returns(uint256)
func (_Dfc *DfcCallerSession) StakingReward() (*big.Int, error) {
	return _Dfc.Contract.StakingReward(&_Dfc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dfc *DfcCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Dfc.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dfc *DfcSession) Symbol() (string, error) {
	return _Dfc.Contract.Symbol(&_Dfc.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Dfc *DfcCallerSession) Symbol() (string, error) {
	return _Dfc.Contract.Symbol(&_Dfc.CallOpts)
}

// TotalDistributed is a free data retrieval call binding the contract method 0xefca2eed.
//
// Solidity: function totalDistributed() view returns(uint256)
func (_Dfc *DfcCaller) TotalDistributed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dfc.contract.Call(opts, &out, "totalDistributed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalDistributed is a free data retrieval call binding the contract method 0xefca2eed.
//
// Solidity: function totalDistributed() view returns(uint256)
func (_Dfc *DfcSession) TotalDistributed() (*big.Int, error) {
	return _Dfc.Contract.TotalDistributed(&_Dfc.CallOpts)
}

// TotalDistributed is a free data retrieval call binding the contract method 0xefca2eed.
//
// Solidity: function totalDistributed() view returns(uint256)
func (_Dfc *DfcCallerSession) TotalDistributed() (*big.Int, error) {
	return _Dfc.Contract.TotalDistributed(&_Dfc.CallOpts)
}

// TotalEarned is a free data retrieval call binding the contract method 0x6dfa8d99.
//
// Solidity: function totalEarned() view returns(uint256)
func (_Dfc *DfcCaller) TotalEarned(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dfc.contract.Call(opts, &out, "totalEarned")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalEarned is a free data retrieval call binding the contract method 0x6dfa8d99.
//
// Solidity: function totalEarned() view returns(uint256)
func (_Dfc *DfcSession) TotalEarned() (*big.Int, error) {
	return _Dfc.Contract.TotalEarned(&_Dfc.CallOpts)
}

// TotalEarned is a free data retrieval call binding the contract method 0x6dfa8d99.
//
// Solidity: function totalEarned() view returns(uint256)
func (_Dfc *DfcCallerSession) TotalEarned() (*big.Int, error) {
	return _Dfc.Contract.TotalEarned(&_Dfc.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Dfc *DfcCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dfc.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Dfc *DfcSession) TotalStaked() (*big.Int, error) {
	return _Dfc.Contract.TotalStaked(&_Dfc.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_Dfc *DfcCallerSession) TotalStaked() (*big.Int, error) {
	return _Dfc.Contract.TotalStaked(&_Dfc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dfc *DfcCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dfc.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dfc *DfcSession) TotalSupply() (*big.Int, error) {
	return _Dfc.Contract.TotalSupply(&_Dfc.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Dfc *DfcCallerSession) TotalSupply() (*big.Int, error) {
	return _Dfc.Contract.TotalSupply(&_Dfc.CallOpts)
}

// Withdrawable is a free data retrieval call binding the contract method 0x50188301.
//
// Solidity: function withdrawable() view returns(uint256)
func (_Dfc *DfcCaller) Withdrawable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dfc.contract.Call(opts, &out, "withdrawable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Withdrawable is a free data retrieval call binding the contract method 0x50188301.
//
// Solidity: function withdrawable() view returns(uint256)
func (_Dfc *DfcSession) Withdrawable() (*big.Int, error) {
	return _Dfc.Contract.Withdrawable(&_Dfc.CallOpts)
}

// Withdrawable is a free data retrieval call binding the contract method 0x50188301.
//
// Solidity: function withdrawable() view returns(uint256)
func (_Dfc *DfcCallerSession) Withdrawable() (*big.Int, error) {
	return _Dfc.Contract.Withdrawable(&_Dfc.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Dfc *DfcTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Dfc.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Dfc *DfcSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Dfc.Contract.Approve(&_Dfc.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool success)
func (_Dfc *DfcTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Dfc.Contract.Approve(&_Dfc.TransactOpts, _spender, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns()
func (_Dfc *DfcTransactor) Burn(opts *bind.TransactOpts, _value *big.Int) (*types.Transaction, error) {
	return _Dfc.contract.Transact(opts, "burn", _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns()
func (_Dfc *DfcSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Dfc.Contract.Burn(&_Dfc.TransactOpts, _value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 _value) returns()
func (_Dfc *DfcTransactorSession) Burn(_value *big.Int) (*types.Transaction, error) {
	return _Dfc.Contract.Burn(&_Dfc.TransactOpts, _value)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns(bool)
func (_Dfc *DfcTransactor) Harvest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dfc.contract.Transact(opts, "harvest")
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns(bool)
func (_Dfc *DfcSession) Harvest() (*types.Transaction, error) {
	return _Dfc.Contract.Harvest(&_Dfc.TransactOpts)
}

// Harvest is a paid mutator transaction binding the contract method 0x4641257d.
//
// Solidity: function harvest() returns(bool)
func (_Dfc *DfcTransactorSession) Harvest() (*types.Transaction, error) {
	return _Dfc.Contract.Harvest(&_Dfc.TransactOpts)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_Dfc *DfcTransactor) PauseStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dfc.contract.Transact(opts, "pauseStaking")
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_Dfc *DfcSession) PauseStaking() (*types.Transaction, error) {
	return _Dfc.Contract.PauseStaking(&_Dfc.TransactOpts)
}

// PauseStaking is a paid mutator transaction binding the contract method 0xf999c506.
//
// Solidity: function pauseStaking() returns()
func (_Dfc *DfcTransactorSession) PauseStaking() (*types.Transaction, error) {
	return _Dfc.Contract.PauseStaking(&_Dfc.TransactOpts)
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_Dfc *DfcTransactor) ResumeStaking(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dfc.contract.Transact(opts, "resumeStaking")
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_Dfc *DfcSession) ResumeStaking() (*types.Transaction, error) {
	return _Dfc.Contract.ResumeStaking(&_Dfc.TransactOpts)
}

// ResumeStaking is a paid mutator transaction binding the contract method 0x7475f913.
//
// Solidity: function resumeStaking() returns()
func (_Dfc *DfcTransactorSession) ResumeStaking() (*types.Transaction, error) {
	return _Dfc.Contract.ResumeStaking(&_Dfc.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x7acb7757.
//
// Solidity: function stake(uint256 _amount, address referer) returns(bool)
func (_Dfc *DfcTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int, referer common.Address) (*types.Transaction, error) {
	return _Dfc.contract.Transact(opts, "stake", _amount, referer)
}

// Stake is a paid mutator transaction binding the contract method 0x7acb7757.
//
// Solidity: function stake(uint256 _amount, address referer) returns(bool)
func (_Dfc *DfcSession) Stake(_amount *big.Int, referer common.Address) (*types.Transaction, error) {
	return _Dfc.Contract.Stake(&_Dfc.TransactOpts, _amount, referer)
}

// Stake is a paid mutator transaction binding the contract method 0x7acb7757.
//
// Solidity: function stake(uint256 _amount, address referer) returns(bool)
func (_Dfc *DfcTransactorSession) Stake(_amount *big.Int, referer common.Address) (*types.Transaction, error) {
	return _Dfc.Contract.Stake(&_Dfc.TransactOpts, _amount, referer)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_Dfc *DfcTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Dfc.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_Dfc *DfcSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Dfc.Contract.Transfer(&_Dfc.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool success)
func (_Dfc *DfcTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Dfc.Contract.Transfer(&_Dfc.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_Dfc *DfcTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Dfc.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_Dfc *DfcSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Dfc.Contract.TransferFrom(&_Dfc.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool success)
func (_Dfc *DfcTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Dfc.Contract.TransferFrom(&_Dfc.TransactOpts, _from, _to, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dfc *DfcTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dfc.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dfc *DfcSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dfc.Contract.TransferOwnership(&_Dfc.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dfc *DfcTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dfc.Contract.TransferOwnership(&_Dfc.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Dfc *DfcTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dfc.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Dfc *DfcSession) Withdraw() (*types.Transaction, error) {
	return _Dfc.Contract.Withdraw(&_Dfc.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Dfc *DfcTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Dfc.Contract.Withdraw(&_Dfc.TransactOpts)
}

// WithdrawAltcoinTokens is a paid mutator transaction binding the contract method 0x2195845f.
//
// Solidity: function withdrawAltcoinTokens(address _tokenContract) returns(bool)
func (_Dfc *DfcTransactor) WithdrawAltcoinTokens(opts *bind.TransactOpts, _tokenContract common.Address) (*types.Transaction, error) {
	return _Dfc.contract.Transact(opts, "withdrawAltcoinTokens", _tokenContract)
}

// WithdrawAltcoinTokens is a paid mutator transaction binding the contract method 0x2195845f.
//
// Solidity: function withdrawAltcoinTokens(address _tokenContract) returns(bool)
func (_Dfc *DfcSession) WithdrawAltcoinTokens(_tokenContract common.Address) (*types.Transaction, error) {
	return _Dfc.Contract.WithdrawAltcoinTokens(&_Dfc.TransactOpts, _tokenContract)
}

// WithdrawAltcoinTokens is a paid mutator transaction binding the contract method 0x2195845f.
//
// Solidity: function withdrawAltcoinTokens(address _tokenContract) returns(bool)
func (_Dfc *DfcTransactorSession) WithdrawAltcoinTokens(_tokenContract common.Address) (*types.Transaction, error) {
	return _Dfc.Contract.WithdrawAltcoinTokens(&_Dfc.TransactOpts, _tokenContract)
}

// DfcApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Dfc contract.
type DfcApprovalIterator struct {
	Event *DfcApproval // Event containing the contract specifics and raw log

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
func (it *DfcApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DfcApproval)
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
		it.Event = new(DfcApproval)
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
func (it *DfcApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DfcApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DfcApproval represents a Approval event raised by the Dfc contract.
type DfcApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Dfc *DfcFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _spender []common.Address) (*DfcApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Dfc.contract.FilterLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return &DfcApprovalIterator{contract: _Dfc.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Dfc *DfcFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *DfcApproval, _owner []common.Address, _spender []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _spenderRule []interface{}
	for _, _spenderItem := range _spender {
		_spenderRule = append(_spenderRule, _spenderItem)
	}

	logs, sub, err := _Dfc.contract.WatchLogs(opts, "Approval", _ownerRule, _spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DfcApproval)
				if err := _Dfc.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed _owner, address indexed _spender, uint256 _value)
func (_Dfc *DfcFilterer) ParseApproval(log types.Log) (*DfcApproval, error) {
	event := new(DfcApproval)
	if err := _Dfc.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DfcBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Dfc contract.
type DfcBurnIterator struct {
	Event *DfcBurn // Event containing the contract specifics and raw log

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
func (it *DfcBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DfcBurn)
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
		it.Event = new(DfcBurn)
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
func (it *DfcBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DfcBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DfcBurn represents a Burn event raised by the Dfc contract.
type DfcBurn struct {
	Burner common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 value)
func (_Dfc *DfcFilterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*DfcBurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _Dfc.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &DfcBurnIterator{contract: _Dfc.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 value)
func (_Dfc *DfcFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *DfcBurn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _Dfc.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DfcBurn)
				if err := _Dfc.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca5.
//
// Solidity: event Burn(address indexed burner, uint256 value)
func (_Dfc *DfcFilterer) ParseBurn(log types.Log) (*DfcBurn, error) {
	event := new(DfcBurn)
	if err := _Dfc.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DfcHarvestIterator is returned from FilterHarvest and is used to iterate over the raw logs and unpacked data for Harvest events raised by the Dfc contract.
type DfcHarvestIterator struct {
	Event *DfcHarvest // Event containing the contract specifics and raw log

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
func (it *DfcHarvestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DfcHarvest)
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
		it.Event = new(DfcHarvest)
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
func (it *DfcHarvestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DfcHarvestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DfcHarvest represents a Harvest event raised by the Dfc contract.
type DfcHarvest struct {
	User  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterHarvest is a free log retrieval operation binding the contract event 0xc9695243a805adb74c91f28311176c65b417e842d5699893cef56d18bfa48cba.
//
// Solidity: event Harvest(address indexed _user, uint256 _value)
func (_Dfc *DfcFilterer) FilterHarvest(opts *bind.FilterOpts, _user []common.Address) (*DfcHarvestIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Dfc.contract.FilterLogs(opts, "Harvest", _userRule)
	if err != nil {
		return nil, err
	}
	return &DfcHarvestIterator{contract: _Dfc.contract, event: "Harvest", logs: logs, sub: sub}, nil
}

// WatchHarvest is a free log subscription operation binding the contract event 0xc9695243a805adb74c91f28311176c65b417e842d5699893cef56d18bfa48cba.
//
// Solidity: event Harvest(address indexed _user, uint256 _value)
func (_Dfc *DfcFilterer) WatchHarvest(opts *bind.WatchOpts, sink chan<- *DfcHarvest, _user []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}

	logs, sub, err := _Dfc.contract.WatchLogs(opts, "Harvest", _userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DfcHarvest)
				if err := _Dfc.contract.UnpackLog(event, "Harvest", log); err != nil {
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

// ParseHarvest is a log parse operation binding the contract event 0xc9695243a805adb74c91f28311176c65b417e842d5699893cef56d18bfa48cba.
//
// Solidity: event Harvest(address indexed _user, uint256 _value)
func (_Dfc *DfcFilterer) ParseHarvest(log types.Log) (*DfcHarvest, error) {
	event := new(DfcHarvest)
	if err := _Dfc.contract.UnpackLog(event, "Harvest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DfcReferralIterator is returned from FilterReferral and is used to iterate over the raw logs and unpacked data for Referral events raised by the Dfc contract.
type DfcReferralIterator struct {
	Event *DfcReferral // Event containing the contract specifics and raw log

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
func (it *DfcReferralIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DfcReferral)
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
		it.Event = new(DfcReferral)
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
func (it *DfcReferralIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DfcReferralIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DfcReferral represents a Referral event raised by the Dfc contract.
type DfcReferral struct {
	User  common.Address
	From  common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterReferral is a free log retrieval operation binding the contract event 0xaeecfcda1271d292db728294b8ae465871ec039d51404caf49a7eb0ade51770a.
//
// Solidity: event Referral(address indexed _user, address indexed _from, uint256 _value)
func (_Dfc *DfcFilterer) FilterReferral(opts *bind.FilterOpts, _user []common.Address, _from []common.Address) (*DfcReferralIterator, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Dfc.contract.FilterLogs(opts, "Referral", _userRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &DfcReferralIterator{contract: _Dfc.contract, event: "Referral", logs: logs, sub: sub}, nil
}

// WatchReferral is a free log subscription operation binding the contract event 0xaeecfcda1271d292db728294b8ae465871ec039d51404caf49a7eb0ade51770a.
//
// Solidity: event Referral(address indexed _user, address indexed _from, uint256 _value)
func (_Dfc *DfcFilterer) WatchReferral(opts *bind.WatchOpts, sink chan<- *DfcReferral, _user []common.Address, _from []common.Address) (event.Subscription, error) {

	var _userRule []interface{}
	for _, _userItem := range _user {
		_userRule = append(_userRule, _userItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Dfc.contract.WatchLogs(opts, "Referral", _userRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DfcReferral)
				if err := _Dfc.contract.UnpackLog(event, "Referral", log); err != nil {
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

// ParseReferral is a log parse operation binding the contract event 0xaeecfcda1271d292db728294b8ae465871ec039d51404caf49a7eb0ade51770a.
//
// Solidity: event Referral(address indexed _user, address indexed _from, uint256 _value)
func (_Dfc *DfcFilterer) ParseReferral(log types.Log) (*DfcReferral, error) {
	event := new(DfcReferral)
	if err := _Dfc.contract.UnpackLog(event, "Referral", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DfcStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Dfc contract.
type DfcStakedIterator struct {
	Event *DfcStaked // Event containing the contract specifics and raw log

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
func (it *DfcStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DfcStaked)
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
		it.Event = new(DfcStaked)
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
func (it *DfcStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DfcStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DfcStaked represents a Staked event raised by the Dfc contract.
type DfcStaked struct {
	From  common.Address
	Value *big.Int
	Term  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed _from, uint256 _value, uint256 _term)
func (_Dfc *DfcFilterer) FilterStaked(opts *bind.FilterOpts, _from []common.Address) (*DfcStakedIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Dfc.contract.FilterLogs(opts, "Staked", _fromRule)
	if err != nil {
		return nil, err
	}
	return &DfcStakedIterator{contract: _Dfc.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed _from, uint256 _value, uint256 _term)
func (_Dfc *DfcFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *DfcStaked, _from []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _Dfc.contract.WatchLogs(opts, "Staked", _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DfcStaked)
				if err := _Dfc.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed _from, uint256 _value, uint256 _term)
func (_Dfc *DfcFilterer) ParseStaked(log types.Log) (*DfcStaked, error) {
	event := new(DfcStaked)
	if err := _Dfc.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DfcTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Dfc contract.
type DfcTransferIterator struct {
	Event *DfcTransfer // Event containing the contract specifics and raw log

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
func (it *DfcTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DfcTransfer)
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
		it.Event = new(DfcTransfer)
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
func (it *DfcTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DfcTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DfcTransfer represents a Transfer event raised by the Dfc contract.
type DfcTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Dfc *DfcFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*DfcTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Dfc.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &DfcTransferIterator{contract: _Dfc.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Dfc *DfcFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *DfcTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Dfc.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DfcTransfer)
				if err := _Dfc.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_Dfc *DfcFilterer) ParseTransfer(log types.Log) (*DfcTransfer, error) {
	event := new(DfcTransfer)
	if err := _Dfc.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
