// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Agreement

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

// UserBorrower is an auto generated low-level Go binding around an user-defined struct.
type UserBorrower struct {
	Borrower common.Address
}

// UserABI is the input ABI used to generate the binding from.
const UserABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethPledge\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"goodsImgs\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"goodSign\",\"type\":\"string\"}],\"name\":\"addGoods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"img\",\"type\":\"string\"}],\"name\":\"addSticker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"people\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"headImg\",\"type\":\"string\"}],\"name\":\"addUserImg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"backs\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"over\",\"type\":\"string\"}],\"name\":\"agreeBack\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deal\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"since\",\"type\":\"string\"}],\"name\":\"agreeBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auctionStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowDays\",\"type\":\"uint256\"}],\"name\":\"borrowGoods\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"back\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"dealHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"delStick\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"backs\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"over\",\"type\":\"string\"}],\"name\":\"disagreeBack\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deal\",\"type\":\"uint256\"}],\"name\":\"disagreeBorrow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"doGoodsReturn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"back\",\"type\":\"uint256\"}],\"name\":\"getBackRec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"people\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deal\",\"type\":\"uint256\"}],\"name\":\"getDealRec\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getGoods\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethPledge\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"goodImg\",\"type\":\"string[]\"},{\"internalType\":\"string\",\"name\":\"goodSign\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGoodsId\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"oid\",\"type\":\"uint256\"}],\"name\":\"getOrders\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOrdersId\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getStick\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"img\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"people\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sign\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"goodsNum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserImg\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"img\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"goodsData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"internalType\":\"structUser.Borrower\",\"name\":\"borrowers\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethPledge\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isAgree\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"deal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"backs\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"available\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isBorrow\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"goodsId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"host\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"integralData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"password\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"people\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"headImg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sign\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"goodsNum\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLogin\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"people\",\"type\":\"address\"}],\"name\":\"login\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"people\",\"type\":\"address\"}],\"name\":\"logout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orderData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"orderOwner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"orderBorrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"oId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"back\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ordersId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"outGoods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"people\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"password\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"people\",\"type\":\"address\"}],\"name\":\"signIn\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"myAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"people\",\"type\":\"address\"}],\"name\":\"topUp\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethPledge\",\"type\":\"uint256\"}],\"name\":\"updGoods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"updateStick\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"people\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sign\",\"type\":\"string\"}],\"name\":\"updateUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// User is an auto generated Go binding around an Ethereum contract.
type User struct {
	UserCaller     // Read-only binding to the contract
	UserTransactor // Write-only binding to the contract
	UserFilterer   // Log filterer for contract events
}

// UserCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserSession struct {
	Contract     *User             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserCallerSession struct {
	Contract *UserCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserTransactorSession struct {
	Contract     *UserTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserRaw struct {
	Contract *User // Generic contract binding to access the raw methods on
}

// UserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserCallerRaw struct {
	Contract *UserCaller // Generic read-only contract binding to access the raw methods on
}

// UserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserTransactorRaw struct {
	Contract *UserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUser creates a new instance of User, bound to a specific deployed contract.
func NewUser(address common.Address, backend bind.ContractBackend) (*User, error) {
	contract, err := bindUser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &User{UserCaller: UserCaller{contract: contract}, UserTransactor: UserTransactor{contract: contract}, UserFilterer: UserFilterer{contract: contract}}, nil
}

// NewUserCaller creates a new read-only instance of User, bound to a specific deployed contract.
func NewUserCaller(address common.Address, caller bind.ContractCaller) (*UserCaller, error) {
	contract, err := bindUser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserCaller{contract: contract}, nil
}

// NewUserTransactor creates a new write-only instance of User, bound to a specific deployed contract.
func NewUserTransactor(address common.Address, transactor bind.ContractTransactor) (*UserTransactor, error) {
	contract, err := bindUser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserTransactor{contract: contract}, nil
}

// NewUserFilterer creates a new log filterer instance of User, bound to a specific deployed contract.
func NewUserFilterer(address common.Address, filterer bind.ContractFilterer) (*UserFilterer, error) {
	contract, err := bindUser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserFilterer{contract: contract}, nil
}

// bindUser binds a generic wrapper to an already deployed contract.
func bindUser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _User.Contract.UserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _User.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.contract.Transact(opts, method, params...)
}

// AuctionLimit is a free data retrieval call binding the contract method 0x44514a42.
//
// Solidity: function auctionLimit() view returns(uint256)
func (_User *UserCaller) AuctionLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "auctionLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionLimit is a free data retrieval call binding the contract method 0x44514a42.
//
// Solidity: function auctionLimit() view returns(uint256)
func (_User *UserSession) AuctionLimit() (*big.Int, error) {
	return _User.Contract.AuctionLimit(&_User.CallOpts)
}

// AuctionLimit is a free data retrieval call binding the contract method 0x44514a42.
//
// Solidity: function auctionLimit() view returns(uint256)
func (_User *UserCallerSession) AuctionLimit() (*big.Int, error) {
	return _User.Contract.AuctionLimit(&_User.CallOpts)
}

// AuctionStart is a free data retrieval call binding the contract method 0x4f245ef7.
//
// Solidity: function auctionStart() view returns(uint256)
func (_User *UserCaller) AuctionStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "auctionStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AuctionStart is a free data retrieval call binding the contract method 0x4f245ef7.
//
// Solidity: function auctionStart() view returns(uint256)
func (_User *UserSession) AuctionStart() (*big.Int, error) {
	return _User.Contract.AuctionStart(&_User.CallOpts)
}

// AuctionStart is a free data retrieval call binding the contract method 0x4f245ef7.
//
// Solidity: function auctionStart() view returns(uint256)
func (_User *UserCallerSession) AuctionStart() (*big.Int, error) {
	return _User.Contract.AuctionStart(&_User.CallOpts)
}

// GetBackRec is a free data retrieval call binding the contract method 0x09986cfb.
//
// Solidity: function getBackRec(uint256 id, uint256 back) view returns(uint256, uint256, bytes32, uint256, bool, string)
func (_User *UserCaller) GetBackRec(opts *bind.CallOpts, id *big.Int, back *big.Int) (*big.Int, *big.Int, [32]byte, *big.Int, bool, string, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "getBackRec", id, back)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new([32]byte), *new(*big.Int), *new(bool), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(bool)).(*bool)
	out5 := *abi.ConvertType(out[5], new(string)).(*string)

	return out0, out1, out2, out3, out4, out5, err

}

// GetBackRec is a free data retrieval call binding the contract method 0x09986cfb.
//
// Solidity: function getBackRec(uint256 id, uint256 back) view returns(uint256, uint256, bytes32, uint256, bool, string)
func (_User *UserSession) GetBackRec(id *big.Int, back *big.Int) (*big.Int, *big.Int, [32]byte, *big.Int, bool, string, error) {
	return _User.Contract.GetBackRec(&_User.CallOpts, id, back)
}

// GetBackRec is a free data retrieval call binding the contract method 0x09986cfb.
//
// Solidity: function getBackRec(uint256 id, uint256 back) view returns(uint256, uint256, bytes32, uint256, bool, string)
func (_User *UserCallerSession) GetBackRec(id *big.Int, back *big.Int) (*big.Int, *big.Int, [32]byte, *big.Int, bool, string, error) {
	return _User.Contract.GetBackRec(&_User.CallOpts, id, back)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address people) view returns(uint256)
func (_User *UserCaller) GetBalance(opts *bind.CallOpts, people common.Address) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "getBalance", people)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address people) view returns(uint256)
func (_User *UserSession) GetBalance(people common.Address) (*big.Int, error) {
	return _User.Contract.GetBalance(&_User.CallOpts, people)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address people) view returns(uint256)
func (_User *UserCallerSession) GetBalance(people common.Address) (*big.Int, error) {
	return _User.Contract.GetBalance(&_User.CallOpts, people)
}

// GetDealRec is a free data retrieval call binding the contract method 0xd05f9905.
//
// Solidity: function getDealRec(uint256 id, uint256 deal) view returns(uint256, string)
func (_User *UserCaller) GetDealRec(opts *bind.CallOpts, id *big.Int, deal *big.Int) (*big.Int, string, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "getDealRec", id, deal)

	if err != nil {
		return *new(*big.Int), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetDealRec is a free data retrieval call binding the contract method 0xd05f9905.
//
// Solidity: function getDealRec(uint256 id, uint256 deal) view returns(uint256, string)
func (_User *UserSession) GetDealRec(id *big.Int, deal *big.Int) (*big.Int, string, error) {
	return _User.Contract.GetDealRec(&_User.CallOpts, id, deal)
}

// GetDealRec is a free data retrieval call binding the contract method 0xd05f9905.
//
// Solidity: function getDealRec(uint256 id, uint256 deal) view returns(uint256, string)
func (_User *UserCallerSession) GetDealRec(id *big.Int, deal *big.Int) (*big.Int, string, error) {
	return _User.Contract.GetDealRec(&_User.CallOpts, id, deal)
}

// GetGoods is a free data retrieval call binding the contract method 0xb2590033.
//
// Solidity: function getGoods(uint256 id) view returns(address owner, string name, string species, uint256 rent, uint256 ethPledge, string[] goodImg, string goodSign)
func (_User *UserCaller) GetGoods(opts *bind.CallOpts, id *big.Int) (struct {
	Owner     common.Address
	Name      string
	Species   string
	Rent      *big.Int
	EthPledge *big.Int
	GoodImg   []string
	GoodSign  string
}, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "getGoods", id)

	outstruct := new(struct {
		Owner     common.Address
		Name      string
		Species   string
		Rent      *big.Int
		EthPledge *big.Int
		GoodImg   []string
		GoodSign  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Species = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Rent = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EthPledge = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.GoodImg = *abi.ConvertType(out[5], new([]string)).(*[]string)
	outstruct.GoodSign = *abi.ConvertType(out[6], new(string)).(*string)

	return *outstruct, err

}

// GetGoods is a free data retrieval call binding the contract method 0xb2590033.
//
// Solidity: function getGoods(uint256 id) view returns(address owner, string name, string species, uint256 rent, uint256 ethPledge, string[] goodImg, string goodSign)
func (_User *UserSession) GetGoods(id *big.Int) (struct {
	Owner     common.Address
	Name      string
	Species   string
	Rent      *big.Int
	EthPledge *big.Int
	GoodImg   []string
	GoodSign  string
}, error) {
	return _User.Contract.GetGoods(&_User.CallOpts, id)
}

// GetGoods is a free data retrieval call binding the contract method 0xb2590033.
//
// Solidity: function getGoods(uint256 id) view returns(address owner, string name, string species, uint256 rent, uint256 ethPledge, string[] goodImg, string goodSign)
func (_User *UserCallerSession) GetGoods(id *big.Int) (struct {
	Owner     common.Address
	Name      string
	Species   string
	Rent      *big.Int
	EthPledge *big.Int
	GoodImg   []string
	GoodSign  string
}, error) {
	return _User.Contract.GetGoods(&_User.CallOpts, id)
}

// GetGoodsId is a free data retrieval call binding the contract method 0x56bcb94c.
//
// Solidity: function getGoodsId() view returns(uint256[])
func (_User *UserCaller) GetGoodsId(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "getGoodsId")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetGoodsId is a free data retrieval call binding the contract method 0x56bcb94c.
//
// Solidity: function getGoodsId() view returns(uint256[])
func (_User *UserSession) GetGoodsId() ([]*big.Int, error) {
	return _User.Contract.GetGoodsId(&_User.CallOpts)
}

// GetGoodsId is a free data retrieval call binding the contract method 0x56bcb94c.
//
// Solidity: function getGoodsId() view returns(uint256[])
func (_User *UserCallerSession) GetGoodsId() ([]*big.Int, error) {
	return _User.Contract.GetGoodsId(&_User.CallOpts)
}

// GetOrders is a free data retrieval call binding the contract method 0x05e0141d.
//
// Solidity: function getOrders(uint256 oid) view returns(address, address, uint256, uint256, uint256)
func (_User *UserCaller) GetOrders(opts *bind.CallOpts, oid *big.Int) (common.Address, common.Address, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "getOrders", oid)

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// GetOrders is a free data retrieval call binding the contract method 0x05e0141d.
//
// Solidity: function getOrders(uint256 oid) view returns(address, address, uint256, uint256, uint256)
func (_User *UserSession) GetOrders(oid *big.Int) (common.Address, common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _User.Contract.GetOrders(&_User.CallOpts, oid)
}

// GetOrders is a free data retrieval call binding the contract method 0x05e0141d.
//
// Solidity: function getOrders(uint256 oid) view returns(address, address, uint256, uint256, uint256)
func (_User *UserCallerSession) GetOrders(oid *big.Int) (common.Address, common.Address, *big.Int, *big.Int, *big.Int, error) {
	return _User.Contract.GetOrders(&_User.CallOpts, oid)
}


// GetOrdersId is a free data retrieval call binding the contract method 0x5d1e0ec1.
//
// Solidity: function getOrdersId() view returns(uint256[])
func (_User *UserCaller) GetOrdersId(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "getOrdersId")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOrdersId is a free data retrieval call binding the contract method 0x5d1e0ec1.
//
// Solidity: function getOrdersId() view returns(uint256[])
func (_User *UserSession) GetOrdersId() ([]*big.Int, error) {
	return _User.Contract.GetOrdersId(&_User.CallOpts)
}

// GetOrdersId is a free data retrieval call binding the contract method 0x5d1e0ec1.
//
// Solidity: function getOrdersId() view returns(uint256[])
func (_User *UserCallerSession) GetOrdersId() ([]*big.Int, error) {
	return _User.Contract.GetOrdersId(&_User.CallOpts)
}

// GetStick is a free data retrieval call binding the contract method 0x92013acb.
//
// Solidity: function getStick(uint256 id) view returns(string name, string img)
func (_User *UserCaller) GetStick(opts *bind.CallOpts, id *big.Int) (struct {
	Name string
	Img  string
}, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "getStick", id)

	outstruct := new(struct {
		Name string
		Img  string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Img = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// GetStick is a free data retrieval call binding the contract method 0x92013acb.
//
// Solidity: function getStick(uint256 id) view returns(string name, string img)
func (_User *UserSession) GetStick(id *big.Int) (struct {
	Name string
	Img  string
}, error) {
	return _User.Contract.GetStick(&_User.CallOpts, id)
}

// GetStick is a free data retrieval call binding the contract method 0x92013acb.
//
// Solidity: function getStick(uint256 id) view returns(string name, string img)
func (_User *UserCallerSession) GetStick(id *big.Int) (struct {
	Name string
	Img  string
}, error) {
	return _User.Contract.GetStick(&_User.CallOpts, id)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address user) view returns(string name, address people, string email, string sign, uint256 goodsNum)
func (_User *UserCaller) GetUser(opts *bind.CallOpts, user common.Address) (struct {
	Name     string
	People   common.Address
	Email    string
	Sign     string
	GoodsNum *big.Int
}, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "getUser", user)

	outstruct := new(struct {
		Name     string
		People   common.Address
		Email    string
		Sign     string
		GoodsNum *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.People = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Email = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Sign = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.GoodsNum = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address user) view returns(string name, address people, string email, string sign, uint256 goodsNum)
func (_User *UserSession) GetUser(user common.Address) (struct {
	Name     string
	People   common.Address
	Email    string
	Sign     string
	GoodsNum *big.Int
}, error) {
	return _User.Contract.GetUser(&_User.CallOpts, user)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address user) view returns(string name, address people, string email, string sign, uint256 goodsNum)
func (_User *UserCallerSession) GetUser(user common.Address) (struct {
	Name     string
	People   common.Address
	Email    string
	Sign     string
	GoodsNum *big.Int
}, error) {
	return _User.Contract.GetUser(&_User.CallOpts, user)
}

// GetUserImg is a free data retrieval call binding the contract method 0x30447b78.
//
// Solidity: function getUserImg(address user) view returns(string img)
func (_User *UserCaller) GetUserImg(opts *bind.CallOpts, user common.Address) (string, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "getUserImg", user)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUserImg is a free data retrieval call binding the contract method 0x30447b78.
//
// Solidity: function getUserImg(address user) view returns(string img)
func (_User *UserSession) GetUserImg(user common.Address) (string, error) {
	return _User.Contract.GetUserImg(&_User.CallOpts, user)
}

// GetUserImg is a free data retrieval call binding the contract method 0x30447b78.
//
// Solidity: function getUserImg(address user) view returns(string img)
func (_User *UserCallerSession) GetUserImg(user common.Address) (string, error) {
	return _User.Contract.GetUserImg(&_User.CallOpts, user)
}

// GoodsData is a free data retrieval call binding the contract method 0xb7305eb2.
//
// Solidity: function goodsData(uint256 ) view returns(address owner, (address) borrowers, uint256 id, string name, string species, uint256 rent, uint256 ethPledge, bool isAgree, uint256 deal, uint256 backs, bool available, bool isBorrow)
func (_User *UserCaller) GoodsData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner     common.Address
	Borrowers UserBorrower
	Id        *big.Int
	Name      string
	Species   string
	Rent      *big.Int
	EthPledge *big.Int
	IsAgree   bool
	Deal      *big.Int
	Backs     *big.Int
	Available bool
	IsBorrow  bool
}, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "goodsData", arg0)

	outstruct := new(struct {
		Owner     common.Address
		Borrowers UserBorrower
		Id        *big.Int
		Name      string
		Species   string
		Rent      *big.Int
		EthPledge *big.Int
		IsAgree   bool
		Deal      *big.Int
		Backs     *big.Int
		Available bool
		IsBorrow  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Borrowers = *abi.ConvertType(out[1], new(UserBorrower)).(*UserBorrower)
	outstruct.Id = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Species = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Rent = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.EthPledge = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.IsAgree = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.Deal = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Backs = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.Available = *abi.ConvertType(out[10], new(bool)).(*bool)
	outstruct.IsBorrow = *abi.ConvertType(out[11], new(bool)).(*bool)

	return *outstruct, err

}

// GoodsData is a free data retrieval call binding the contract method 0xb7305eb2.
//
// Solidity: function goodsData(uint256 ) view returns(address owner, (address) borrowers, uint256 id, string name, string species, uint256 rent, uint256 ethPledge, bool isAgree, uint256 deal, uint256 backs, bool available, bool isBorrow)
func (_User *UserSession) GoodsData(arg0 *big.Int) (struct {
	Owner     common.Address
	Borrowers UserBorrower
	Id        *big.Int
	Name      string
	Species   string
	Rent      *big.Int
	EthPledge *big.Int
	IsAgree   bool
	Deal      *big.Int
	Backs     *big.Int
	Available bool
	IsBorrow  bool
}, error) {
	return _User.Contract.GoodsData(&_User.CallOpts, arg0)
}

// GoodsData is a free data retrieval call binding the contract method 0xb7305eb2.
//
// Solidity: function goodsData(uint256 ) view returns(address owner, (address) borrowers, uint256 id, string name, string species, uint256 rent, uint256 ethPledge, bool isAgree, uint256 deal, uint256 backs, bool available, bool isBorrow)
func (_User *UserCallerSession) GoodsData(arg0 *big.Int) (struct {
	Owner     common.Address
	Borrowers UserBorrower
	Id        *big.Int
	Name      string
	Species   string
	Rent      *big.Int
	EthPledge *big.Int
	IsAgree   bool
	Deal      *big.Int
	Backs     *big.Int
	Available bool
	IsBorrow  bool
}, error) {
	return _User.Contract.GoodsData(&_User.CallOpts, arg0)
}

// GoodsId is a free data retrieval call binding the contract method 0xdece54fb.
//
// Solidity: function goodsId(uint256 ) view returns(uint256)
func (_User *UserCaller) GoodsId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "goodsId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GoodsId is a free data retrieval call binding the contract method 0xdece54fb.
//
// Solidity: function goodsId(uint256 ) view returns(uint256)
func (_User *UserSession) GoodsId(arg0 *big.Int) (*big.Int, error) {
	return _User.Contract.GoodsId(&_User.CallOpts, arg0)
}

// GoodsId is a free data retrieval call binding the contract method 0xdece54fb.
//
// Solidity: function goodsId(uint256 ) view returns(uint256)
func (_User *UserCallerSession) GoodsId(arg0 *big.Int) (*big.Int, error) {
	return _User.Contract.GoodsId(&_User.CallOpts, arg0)
}

// Host is a free data retrieval call binding the contract method 0xf437bc59.
//
// Solidity: function host() view returns(address)
func (_User *UserCaller) Host(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "host")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Host is a free data retrieval call binding the contract method 0xf437bc59.
//
// Solidity: function host() view returns(address)
func (_User *UserSession) Host() (common.Address, error) {
	return _User.Contract.Host(&_User.CallOpts)
}

// Host is a free data retrieval call binding the contract method 0xf437bc59.
//
// Solidity: function host() view returns(address)
func (_User *UserCallerSession) Host() (common.Address, error) {
	return _User.Contract.Host(&_User.CallOpts)
}

// IntegralData is a free data retrieval call binding the contract method 0x89b21711.
//
// Solidity: function integralData(address ) view returns(string name, string password, address people, string email, string headImg, string sign, uint256 goodsNum, bool exist, bool isLogin)
func (_User *UserCaller) IntegralData(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name     string
	Password string
	People   common.Address
	Email    string
	HeadImg  string
	Sign     string
	GoodsNum *big.Int
	Exist    bool
	IsLogin  bool
}, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "integralData", arg0)

	outstruct := new(struct {
		Name     string
		Password string
		People   common.Address
		Email    string
		HeadImg  string
		Sign     string
		GoodsNum *big.Int
		Exist    bool
		IsLogin  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Password = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.People = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Email = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.HeadImg = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Sign = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.GoodsNum = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Exist = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.IsLogin = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// IntegralData is a free data retrieval call binding the contract method 0x89b21711.
//
// Solidity: function integralData(address ) view returns(string name, string password, address people, string email, string headImg, string sign, uint256 goodsNum, bool exist, bool isLogin)
func (_User *UserSession) IntegralData(arg0 common.Address) (struct {
	Name     string
	Password string
	People   common.Address
	Email    string
	HeadImg  string
	Sign     string
	GoodsNum *big.Int
	Exist    bool
	IsLogin  bool
}, error) {
	return _User.Contract.IntegralData(&_User.CallOpts, arg0)
}

// IntegralData is a free data retrieval call binding the contract method 0x89b21711.
//
// Solidity: function integralData(address ) view returns(string name, string password, address people, string email, string headImg, string sign, uint256 goodsNum, bool exist, bool isLogin)
func (_User *UserCallerSession) IntegralData(arg0 common.Address) (struct {
	Name     string
	Password string
	People   common.Address
	Email    string
	HeadImg  string
	Sign     string
	GoodsNum *big.Int
	Exist    bool
	IsLogin  bool
}, error) {
	return _User.Contract.IntegralData(&_User.CallOpts, arg0)
}


// OrderData is a free data retrieval call binding the contract method 0xcfb04022.
//
// Solidity: function orderData(uint256 ) view returns(address orderOwner, address orderBorrower, uint256 oId, uint256 id, uint256 back)
func (_User *UserCaller) OrderData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OrderOwner    common.Address
	OrderBorrower common.Address
	OId           *big.Int
	Id            *big.Int
	Back          *big.Int
}, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "orderData", arg0)

	outstruct := new(struct {
		OrderOwner    common.Address
		OrderBorrower common.Address
		OId           *big.Int
		Id            *big.Int
		Back          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderOwner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.OrderBorrower = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.OId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Id = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Back = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OrderData is a free data retrieval call binding the contract method 0xcfb04022.
//
// Solidity: function orderData(uint256 ) view returns(address orderOwner, address orderBorrower, uint256 oId, uint256 id, uint256 back)
func (_User *UserSession) OrderData(arg0 *big.Int) (struct {
	OrderOwner    common.Address
	OrderBorrower common.Address
	OId           *big.Int
	Id            *big.Int
	Back          *big.Int
}, error) {
	return _User.Contract.OrderData(&_User.CallOpts, arg0)
}

// OrderData is a free data retrieval call binding the contract method 0xcfb04022.
//
// Solidity: function orderData(uint256 ) view returns(address orderOwner, address orderBorrower, uint256 oId, uint256 id, uint256 back)
func (_User *UserCallerSession) OrderData(arg0 *big.Int) (struct {
	OrderOwner    common.Address
	OrderBorrower common.Address
	OId           *big.Int
	Id            *big.Int
	Back          *big.Int
}, error) {
	return _User.Contract.OrderData(&_User.CallOpts, arg0)
}

// OrdersId is a free data retrieval call binding the contract method 0xa2ea6015.
//
// Solidity: function ordersId(uint256 ) view returns(uint256)
func (_User *UserCaller) OrdersId(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "ordersId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrdersId is a free data retrieval call binding the contract method 0xa2ea6015.
//
// Solidity: function ordersId(uint256 ) view returns(uint256)
func (_User *UserSession) OrdersId(arg0 *big.Int) (*big.Int, error) {
	return _User.Contract.OrdersId(&_User.CallOpts, arg0)
}

// OrdersId is a free data retrieval call binding the contract method 0xa2ea6015.
//
// Solidity: function ordersId(uint256 ) view returns(uint256)
func (_User *UserCallerSession) OrdersId(arg0 *big.Int) (*big.Int, error) {
	return _User.Contract.OrdersId(&_User.CallOpts, arg0)
}

// SignIn is a free data retrieval call binding the contract method 0x8fa9e55c.
//
// Solidity: function signIn(address people) view returns(string)
func (_User *UserCaller) SignIn(opts *bind.CallOpts, people common.Address) (string, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "signIn", people)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SignIn is a free data retrieval call binding the contract method 0x8fa9e55c.
//
// Solidity: function signIn(address people) view returns(string)
func (_User *UserSession) SignIn(people common.Address) (string, error) {
	return _User.Contract.SignIn(&_User.CallOpts, people)
}

// SignIn is a free data retrieval call binding the contract method 0x8fa9e55c.
//
// Solidity: function signIn(address people) view returns(string)
func (_User *UserCallerSession) SignIn(people common.Address) (string, error) {
	return _User.Contract.SignIn(&_User.CallOpts, people)
}

// AddGoods is a paid mutator transaction binding the contract method 0xb199cd04.
//
// Solidity: function addGoods(address owner, string name, string species, uint256 rent, uint256 ethPledge, string[] goodsImgs, string goodSign) returns(uint256)
func (_User *UserTransactor) AddGoods(opts *bind.TransactOpts, owner common.Address, name string, species string, rent *big.Int, ethPledge *big.Int, goodsImgs []string, goodSign string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "addGoods", owner, name, species, rent, ethPledge, goodsImgs, goodSign)
}

// AddGoods is a paid mutator transaction binding the contract method 0xb199cd04.
//
// Solidity: function addGoods(address owner, string name, string species, uint256 rent, uint256 ethPledge, string[] goodsImgs, string goodSign) returns(uint256)
func (_User *UserSession) AddGoods(owner common.Address, name string, species string, rent *big.Int, ethPledge *big.Int, goodsImgs []string, goodSign string) (*types.Transaction, error) {
	return _User.Contract.AddGoods(&_User.TransactOpts, owner, name, species, rent, ethPledge, goodsImgs, goodSign)
}

// AddGoods is a paid mutator transaction binding the contract method 0xb199cd04.
//
// Solidity: function addGoods(address owner, string name, string species, uint256 rent, uint256 ethPledge, string[] goodsImgs, string goodSign) returns(uint256)
func (_User *UserTransactorSession) AddGoods(owner common.Address, name string, species string, rent *big.Int, ethPledge *big.Int, goodsImgs []string, goodSign string) (*types.Transaction, error) {
	return _User.Contract.AddGoods(&_User.TransactOpts, owner, name, species, rent, ethPledge, goodsImgs, goodSign)
}

// AddSticker is a paid mutator transaction binding the contract method 0x96ad77ff.
//
// Solidity: function addSticker(string species, string img) returns()
func (_User *UserTransactor) AddSticker(opts *bind.TransactOpts, species string, img string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "addSticker", species, img)
}

// AddSticker is a paid mutator transaction binding the contract method 0x96ad77ff.
//
// Solidity: function addSticker(string species, string img) returns()
func (_User *UserSession) AddSticker(species string, img string) (*types.Transaction, error) {
	return _User.Contract.AddSticker(&_User.TransactOpts, species, img)
}

// AddSticker is a paid mutator transaction binding the contract method 0x96ad77ff.
//
// Solidity: function addSticker(string species, string img) returns()
func (_User *UserTransactorSession) AddSticker(species string, img string) (*types.Transaction, error) {
	return _User.Contract.AddSticker(&_User.TransactOpts, species, img)
}

// AddUserImg is a paid mutator transaction binding the contract method 0x47a0a005.
//
// Solidity: function addUserImg(address people, string headImg) returns()
func (_User *UserTransactor) AddUserImg(opts *bind.TransactOpts, people common.Address, headImg string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "addUserImg", people, headImg)
}

// AddUserImg is a paid mutator transaction binding the contract method 0x47a0a005.
//
// Solidity: function addUserImg(address people, string headImg) returns()
func (_User *UserSession) AddUserImg(people common.Address, headImg string) (*types.Transaction, error) {
	return _User.Contract.AddUserImg(&_User.TransactOpts, people, headImg)
}

// AddUserImg is a paid mutator transaction binding the contract method 0x47a0a005.
//
// Solidity: function addUserImg(address people, string headImg) returns()
func (_User *UserTransactorSession) AddUserImg(people common.Address, headImg string) (*types.Transaction, error) {
	return _User.Contract.AddUserImg(&_User.TransactOpts, people, headImg)
}

// AgreeBack is a paid mutator transaction binding the contract method 0xf18dbb38.
//
// Solidity: function agreeBack(uint256 id, uint256 backs, string over) payable returns(uint256)
func (_User *UserTransactor) AgreeBack(opts *bind.TransactOpts, id *big.Int, backs *big.Int, over string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "agreeBack", id, backs, over)
}

// AgreeBack is a paid mutator transaction binding the contract method 0xf18dbb38.
//
// Solidity: function agreeBack(uint256 id, uint256 backs, string over) payable returns(uint256)
func (_User *UserSession) AgreeBack(id *big.Int, backs *big.Int, over string) (*types.Transaction, error) {
	return _User.Contract.AgreeBack(&_User.TransactOpts, id, backs, over)
}

// AgreeBack is a paid mutator transaction binding the contract method 0xf18dbb38.
//
// Solidity: function agreeBack(uint256 id, uint256 backs, string over) payable returns(uint256)
func (_User *UserTransactorSession) AgreeBack(id *big.Int, backs *big.Int, over string) (*types.Transaction, error) {
	return _User.Contract.AgreeBack(&_User.TransactOpts, id, backs, over)
}

// AgreeBorrow is a paid mutator transaction binding the contract method 0xa4507fb9.
//
// Solidity: function agreeBorrow(uint256 id, uint256 deal, string since) returns()
func (_User *UserTransactor) AgreeBorrow(opts *bind.TransactOpts, id *big.Int, deal *big.Int, since string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "agreeBorrow", id, deal, since)
}

// AgreeBorrow is a paid mutator transaction binding the contract method 0xa4507fb9.
//
// Solidity: function agreeBorrow(uint256 id, uint256 deal, string since) returns()
func (_User *UserSession) AgreeBorrow(id *big.Int, deal *big.Int, since string) (*types.Transaction, error) {
	return _User.Contract.AgreeBorrow(&_User.TransactOpts, id, deal, since)
}

// AgreeBorrow is a paid mutator transaction binding the contract method 0xa4507fb9.
//
// Solidity: function agreeBorrow(uint256 id, uint256 deal, string since) returns()
func (_User *UserTransactorSession) AgreeBorrow(id *big.Int, deal *big.Int, since string) (*types.Transaction, error) {
	return _User.Contract.AgreeBorrow(&_User.TransactOpts, id, deal, since)
}

// BorrowGoods is a paid mutator transaction binding the contract method 0xca66f9ae.
//
// Solidity: function borrowGoods(uint256 id, uint256 borrowDays) payable returns()
func (_User *UserTransactor) BorrowGoods(opts *bind.TransactOpts, id *big.Int, borrowDays *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "borrowGoods", id, borrowDays)
}

// BorrowGoods is a paid mutator transaction binding the contract method 0xca66f9ae.
//
// Solidity: function borrowGoods(uint256 id, uint256 borrowDays) payable returns()
func (_User *UserSession) BorrowGoods(id *big.Int, borrowDays *big.Int) (*types.Transaction, error) {
	return _User.Contract.BorrowGoods(&_User.TransactOpts, id, borrowDays)
}

// BorrowGoods is a paid mutator transaction binding the contract method 0xca66f9ae.
//
// Solidity: function borrowGoods(uint256 id, uint256 borrowDays) payable returns()
func (_User *UserTransactorSession) BorrowGoods(id *big.Int, borrowDays *big.Int) (*types.Transaction, error) {
	return _User.Contract.BorrowGoods(&_User.TransactOpts, id, borrowDays)
}

// DealHash is a paid mutator transaction binding the contract method 0x94d8ad51.
//
// Solidity: function dealHash(uint256 id, uint256 back, uint256 blockNum) returns()
func (_User *UserTransactor) DealHash(opts *bind.TransactOpts, id *big.Int, back *big.Int, blockNum *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "dealHash", id, back, blockNum)
}

// DealHash is a paid mutator transaction binding the contract method 0x94d8ad51.
//
// Solidity: function dealHash(uint256 id, uint256 back, uint256 blockNum) returns()
func (_User *UserSession) DealHash(id *big.Int, back *big.Int, blockNum *big.Int) (*types.Transaction, error) {
	return _User.Contract.DealHash(&_User.TransactOpts, id, back, blockNum)
}

// DealHash is a paid mutator transaction binding the contract method 0x94d8ad51.
//
// Solidity: function dealHash(uint256 id, uint256 back, uint256 blockNum) returns()
func (_User *UserTransactorSession) DealHash(id *big.Int, back *big.Int, blockNum *big.Int) (*types.Transaction, error) {
	return _User.Contract.DealHash(&_User.TransactOpts, id, back, blockNum)
}

// DelStick is a paid mutator transaction binding the contract method 0x690ae647.
//
// Solidity: function delStick(uint256 id) returns()
func (_User *UserTransactor) DelStick(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "delStick", id)
}

// DelStick is a paid mutator transaction binding the contract method 0x690ae647.
//
// Solidity: function delStick(uint256 id) returns()
func (_User *UserSession) DelStick(id *big.Int) (*types.Transaction, error) {
	return _User.Contract.DelStick(&_User.TransactOpts, id)
}

// DelStick is a paid mutator transaction binding the contract method 0x690ae647.
//
// Solidity: function delStick(uint256 id) returns()
func (_User *UserTransactorSession) DelStick(id *big.Int) (*types.Transaction, error) {
	return _User.Contract.DelStick(&_User.TransactOpts, id)
}

// DisagreeBack is a paid mutator transaction binding the contract method 0xd2ed4be5.
//
// Solidity: function disagreeBack(uint256 id, uint256 backs, string over) returns(uint256)
func (_User *UserTransactor) DisagreeBack(opts *bind.TransactOpts, id *big.Int, backs *big.Int, over string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "disagreeBack", id, backs, over)
}

// DisagreeBack is a paid mutator transaction binding the contract method 0xd2ed4be5.
//
// Solidity: function disagreeBack(uint256 id, uint256 backs, string over) returns(uint256)
func (_User *UserSession) DisagreeBack(id *big.Int, backs *big.Int, over string) (*types.Transaction, error) {
	return _User.Contract.DisagreeBack(&_User.TransactOpts, id, backs, over)
}

// DisagreeBack is a paid mutator transaction binding the contract method 0xd2ed4be5.
//
// Solidity: function disagreeBack(uint256 id, uint256 backs, string over) returns(uint256)
func (_User *UserTransactorSession) DisagreeBack(id *big.Int, backs *big.Int, over string) (*types.Transaction, error) {
	return _User.Contract.DisagreeBack(&_User.TransactOpts, id, backs, over)
}

// DisagreeBorrow is a paid mutator transaction binding the contract method 0x710878db.
//
// Solidity: function disagreeBorrow(uint256 id, uint256 deal) payable returns()
func (_User *UserTransactor) DisagreeBorrow(opts *bind.TransactOpts, id *big.Int, deal *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "disagreeBorrow", id, deal)
}

// DisagreeBorrow is a paid mutator transaction binding the contract method 0x710878db.
//
// Solidity: function disagreeBorrow(uint256 id, uint256 deal) payable returns()
func (_User *UserSession) DisagreeBorrow(id *big.Int, deal *big.Int) (*types.Transaction, error) {
	return _User.Contract.DisagreeBorrow(&_User.TransactOpts, id, deal)
}

// DisagreeBorrow is a paid mutator transaction binding the contract method 0x710878db.
//
// Solidity: function disagreeBorrow(uint256 id, uint256 deal) payable returns()
func (_User *UserTransactorSession) DisagreeBorrow(id *big.Int, deal *big.Int) (*types.Transaction, error) {
	return _User.Contract.DisagreeBorrow(&_User.TransactOpts, id, deal)
}

// DoGoodsReturn is a paid mutator transaction binding the contract method 0xbf90d352.
//
// Solidity: function doGoodsReturn(uint256 id) returns()
func (_User *UserTransactor) DoGoodsReturn(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "doGoodsReturn", id)
}

// DoGoodsReturn is a paid mutator transaction binding the contract method 0xbf90d352.
//
// Solidity: function doGoodsReturn(uint256 id) returns()
func (_User *UserSession) DoGoodsReturn(id *big.Int) (*types.Transaction, error) {
	return _User.Contract.DoGoodsReturn(&_User.TransactOpts, id)
}

// DoGoodsReturn is a paid mutator transaction binding the contract method 0xbf90d352.
//
// Solidity: function doGoodsReturn(uint256 id) returns()
func (_User *UserTransactorSession) DoGoodsReturn(id *big.Int) (*types.Transaction, error) {
	return _User.Contract.DoGoodsReturn(&_User.TransactOpts, id)
}

// Login is a paid mutator transaction binding the contract method 0x35a6861a.
//
// Solidity: function login(address people) returns()
func (_User *UserTransactor) Login(opts *bind.TransactOpts, people common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "login", people)
}

// Login is a paid mutator transaction binding the contract method 0x35a6861a.
//
// Solidity: function login(address people) returns()
func (_User *UserSession) Login(people common.Address) (*types.Transaction, error) {
	return _User.Contract.Login(&_User.TransactOpts, people)
}

// Login is a paid mutator transaction binding the contract method 0x35a6861a.
//
// Solidity: function login(address people) returns()
func (_User *UserTransactorSession) Login(people common.Address) (*types.Transaction, error) {
	return _User.Contract.Login(&_User.TransactOpts, people)
}

// Logout is a paid mutator transaction binding the contract method 0xcc51f9b9.
//
// Solidity: function logout(address people) returns()
func (_User *UserTransactor) Logout(opts *bind.TransactOpts, people common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "logout", people)
}

// Logout is a paid mutator transaction binding the contract method 0xcc51f9b9.
//
// Solidity: function logout(address people) returns()
func (_User *UserSession) Logout(people common.Address) (*types.Transaction, error) {
	return _User.Contract.Logout(&_User.TransactOpts, people)
}

// Logout is a paid mutator transaction binding the contract method 0xcc51f9b9.
//
// Solidity: function logout(address people) returns()
func (_User *UserTransactorSession) Logout(people common.Address) (*types.Transaction, error) {
	return _User.Contract.Logout(&_User.TransactOpts, people)
}

// OutGoods is a paid mutator transaction binding the contract method 0x12aa1794.
//
// Solidity: function outGoods(address owner, uint256 id) returns()
func (_User *UserTransactor) OutGoods(opts *bind.TransactOpts, owner common.Address, id *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "outGoods", owner, id)
}

// OutGoods is a paid mutator transaction binding the contract method 0x12aa1794.
//
// Solidity: function outGoods(address owner, uint256 id) returns()
func (_User *UserSession) OutGoods(owner common.Address, id *big.Int) (*types.Transaction, error) {
	return _User.Contract.OutGoods(&_User.TransactOpts, owner, id)
}

// OutGoods is a paid mutator transaction binding the contract method 0x12aa1794.
//
// Solidity: function outGoods(address owner, uint256 id) returns()
func (_User *UserTransactorSession) OutGoods(owner common.Address, id *big.Int) (*types.Transaction, error) {
	return _User.Contract.OutGoods(&_User.TransactOpts, owner, id)
}

// Register is a paid mutator transaction binding the contract method 0xa9ecaad4.
//
// Solidity: function register(string name, address people, string email, string password) payable returns()
func (_User *UserTransactor) Register(opts *bind.TransactOpts, name string, people common.Address, email string, password string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "register", name, people, email, password)
}

// Register is a paid mutator transaction binding the contract method 0xa9ecaad4.
//
// Solidity: function register(string name, address people, string email, string password) payable returns()
func (_User *UserSession) Register(name string, people common.Address, email string, password string) (*types.Transaction, error) {
	return _User.Contract.Register(&_User.TransactOpts, name, people, email, password)
}

// Register is a paid mutator transaction binding the contract method 0xa9ecaad4.
//
// Solidity: function register(string name, address people, string email, string password) payable returns()
func (_User *UserTransactorSession) Register(name string, people common.Address, email string, password string) (*types.Transaction, error) {
	return _User.Contract.Register(&_User.TransactOpts, name, people, email, password)
}

// TopUp is a paid mutator transaction binding the contract method 0xab839ad2.
//
// Solidity: function topUp(uint256 myAmount, address people) payable returns()
func (_User *UserTransactor) TopUp(opts *bind.TransactOpts, myAmount *big.Int, people common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "topUp", myAmount, people)
}

// TopUp is a paid mutator transaction binding the contract method 0xab839ad2.
//
// Solidity: function topUp(uint256 myAmount, address people) payable returns()
func (_User *UserSession) TopUp(myAmount *big.Int, people common.Address) (*types.Transaction, error) {
	return _User.Contract.TopUp(&_User.TransactOpts, myAmount, people)
}

// TopUp is a paid mutator transaction binding the contract method 0xab839ad2.
//
// Solidity: function topUp(uint256 myAmount, address people) payable returns()
func (_User *UserTransactorSession) TopUp(myAmount *big.Int, people common.Address) (*types.Transaction, error) {
	return _User.Contract.TopUp(&_User.TransactOpts, myAmount, people)
}

// UpdGoods is a paid mutator transaction binding the contract method 0xf139ddd2.
//
// Solidity: function updGoods(uint256 id, string name, string species, uint256 rent, uint256 ethPledge) returns()
func (_User *UserTransactor) UpdGoods(opts *bind.TransactOpts, id *big.Int, name string, species string, rent *big.Int, ethPledge *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "updGoods", id, name, species, rent, ethPledge)
}

// UpdGoods is a paid mutator transaction binding the contract method 0xf139ddd2.
//
// Solidity: function updGoods(uint256 id, string name, string species, uint256 rent, uint256 ethPledge) returns()
func (_User *UserSession) UpdGoods(id *big.Int, name string, species string, rent *big.Int, ethPledge *big.Int) (*types.Transaction, error) {
	return _User.Contract.UpdGoods(&_User.TransactOpts, id, name, species, rent, ethPledge)
}

// UpdGoods is a paid mutator transaction binding the contract method 0xf139ddd2.
//
// Solidity: function updGoods(uint256 id, string name, string species, uint256 rent, uint256 ethPledge) returns()
func (_User *UserTransactorSession) UpdGoods(id *big.Int, name string, species string, rent *big.Int, ethPledge *big.Int) (*types.Transaction, error) {
	return _User.Contract.UpdGoods(&_User.TransactOpts, id, name, species, rent, ethPledge)
}

// UpdateStick is a paid mutator transaction binding the contract method 0x7e75620d.
//
// Solidity: function updateStick(uint256 id, string newName) returns()
func (_User *UserTransactor) UpdateStick(opts *bind.TransactOpts, id *big.Int, newName string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "updateStick", id, newName)
}

// UpdateStick is a paid mutator transaction binding the contract method 0x7e75620d.
//
// Solidity: function updateStick(uint256 id, string newName) returns()
func (_User *UserSession) UpdateStick(id *big.Int, newName string) (*types.Transaction, error) {
	return _User.Contract.UpdateStick(&_User.TransactOpts, id, newName)
}

// UpdateStick is a paid mutator transaction binding the contract method 0x7e75620d.
//
// Solidity: function updateStick(uint256 id, string newName) returns()
func (_User *UserTransactorSession) UpdateStick(id *big.Int, newName string) (*types.Transaction, error) {
	return _User.Contract.UpdateStick(&_User.TransactOpts, id, newName)
}

// UpdateUser is a paid mutator transaction binding the contract method 0x0da4578b.
//
// Solidity: function updateUser(address people, string name, string email, string sign) returns()
func (_User *UserTransactor) UpdateUser(opts *bind.TransactOpts, people common.Address, name string, email string, sign string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "updateUser", people, name, email, sign)
}

// UpdateUser is a paid mutator transaction binding the contract method 0x0da4578b.
//
// Solidity: function updateUser(address people, string name, string email, string sign) returns()
func (_User *UserSession) UpdateUser(people common.Address, name string, email string, sign string) (*types.Transaction, error) {
	return _User.Contract.UpdateUser(&_User.TransactOpts, people, name, email, sign)
}

// UpdateUser is a paid mutator transaction binding the contract method 0x0da4578b.
//
// Solidity: function updateUser(address people, string name, string email, string sign) returns()
func (_User *UserTransactorSession) UpdateUser(people common.Address, name string, email string, sign string) (*types.Transaction, error) {
	return _User.Contract.UpdateUser(&_User.TransactOpts, people, name, email, sign)
}
