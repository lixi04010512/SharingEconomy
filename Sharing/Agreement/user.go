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
	Since    string
	Over     string
}

// UserABI is the input ABI used to generate the binding from.
const UserABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"CommunitySchedule\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"CommunitysData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"communityName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"communityIntroduce\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"communityAmounts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"communitySchedule\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiaryAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"communityName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"communityIntroduce\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"communityAmounts\",\"type\":\"uint256\"}],\"name\":\"addCommunity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethPledge\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"goodsImgs\",\"type\":\"string[]\"}],\"name\":\"addGoods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"}],\"name\":\"addSticker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"people\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"headImg\",\"type\":\"string\"}],\"name\":\"addUserImg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"backs\",\"type\":\"uint256\"}],\"name\":\"agreeBack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deal\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"since\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"}],\"name\":\"agreeBorrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"backs\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"over\",\"type\":\"string\"}],\"name\":\"backGoods\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deal\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowDays\",\"type\":\"uint256\"}],\"name\":\"borrowGoods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"delCommunity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"delGoods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"doGoodsReturn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"donator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"donatorName\",\"type\":\"string\"}],\"name\":\"donate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donatorData\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"donatorsAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getGoods\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethPledge\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getGoodsId\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUser\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"people\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"integral\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sign\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"goodsNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"goodsData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"since\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"over\",\"type\":\"string\"}],\"internalType\":\"structUser.Borrower\",\"name\":\"borrowers\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethPledge\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deal\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"backs\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"available\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isBorrow\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"goodsId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"integralData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"password\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"people\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"integral\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"headImg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sign\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"goodsNum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isLogin\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"isCommunityExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isGoodExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"isGoodsLend\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"}],\"name\":\"isStickExist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"people\",\"type\":\"address\"}],\"name\":\"login\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"people\",\"type\":\"address\"}],\"name\":\"logout\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"outGoods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"queryBorrower\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"people\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"password\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"species\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"rent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ethPledge\",\"type\":\"uint256\"}],\"name\":\"updGoods\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"people\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"sign\",\"type\":\"string\"}],\"name\":\"updateUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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

// CommunitySchedule is a free data retrieval call binding the contract method 0x4ed983fc.
//
// Solidity: function CommunitySchedule(uint256 number) view returns(uint256)
func (_User *UserCaller) CommunitySchedule(opts *bind.CallOpts, number *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "CommunitySchedule", number)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommunitySchedule is a free data retrieval call binding the contract method 0x4ed983fc.
//
// Solidity: function CommunitySchedule(uint256 number) view returns(uint256)
func (_User *UserSession) CommunitySchedule(number *big.Int) (*big.Int, error) {
	return _User.Contract.CommunitySchedule(&_User.CallOpts, number)
}

// CommunitySchedule is a free data retrieval call binding the contract method 0x4ed983fc.
//
// Solidity: function CommunitySchedule(uint256 number) view returns(uint256)
func (_User *UserCallerSession) CommunitySchedule(number *big.Int) (*big.Int, error) {
	return _User.Contract.CommunitySchedule(&_User.CallOpts, number)
}

// CommunitysData is a free data retrieval call binding the contract method 0xfc711b83.
//
// Solidity: function CommunitysData(uint256 ) view returns(string communityName, string communityIntroduce, address beneficiaryAddr, uint256 communityAmounts, uint256 communitySchedule)
func (_User *UserCaller) CommunitysData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	CommunityName      string
	CommunityIntroduce string
	BeneficiaryAddr    common.Address
	CommunityAmounts   *big.Int
	CommunitySchedule  *big.Int
}, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "CommunitysData", arg0)

	outstruct := new(struct {
		CommunityName      string
		CommunityIntroduce string
		BeneficiaryAddr    common.Address
		CommunityAmounts   *big.Int
		CommunitySchedule  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CommunityName = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.CommunityIntroduce = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.BeneficiaryAddr = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.CommunityAmounts = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CommunitySchedule = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CommunitysData is a free data retrieval call binding the contract method 0xfc711b83.
//
// Solidity: function CommunitysData(uint256 ) view returns(string communityName, string communityIntroduce, address beneficiaryAddr, uint256 communityAmounts, uint256 communitySchedule)
func (_User *UserSession) CommunitysData(arg0 *big.Int) (struct {
	CommunityName      string
	CommunityIntroduce string
	BeneficiaryAddr    common.Address
	CommunityAmounts   *big.Int
	CommunitySchedule  *big.Int
}, error) {
	return _User.Contract.CommunitysData(&_User.CallOpts, arg0)
}

// CommunitysData is a free data retrieval call binding the contract method 0xfc711b83.
//
// Solidity: function CommunitysData(uint256 ) view returns(string communityName, string communityIntroduce, address beneficiaryAddr, uint256 communityAmounts, uint256 communitySchedule)
func (_User *UserCallerSession) CommunitysData(arg0 *big.Int) (struct {
	CommunityName      string
	CommunityIntroduce string
	BeneficiaryAddr    common.Address
	CommunityAmounts   *big.Int
	CommunitySchedule  *big.Int
}, error) {
	return _User.Contract.CommunitysData(&_User.CallOpts, arg0)
}

// DonatorData is a free data retrieval call binding the contract method 0x19dfc4d3.
//
// Solidity: function donatorData() view returns(address[])
func (_User *UserCaller) DonatorData(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "donatorData")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// DonatorData is a free data retrieval call binding the contract method 0x19dfc4d3.
//
// Solidity: function donatorData() view returns(address[])
func (_User *UserSession) DonatorData() ([]common.Address, error) {
	return _User.Contract.DonatorData(&_User.CallOpts)
}

// DonatorData is a free data retrieval call binding the contract method 0x19dfc4d3.
//
// Solidity: function donatorData() view returns(address[])
func (_User *UserCallerSession) DonatorData() ([]common.Address, error) {
	return _User.Contract.DonatorData(&_User.CallOpts)
}

// DonatorsAddr is a free data retrieval call binding the contract method 0xcbae0e49.
//
// Solidity: function donatorsAddr(uint256 ) view returns(address)
func (_User *UserCaller) DonatorsAddr(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "donatorsAddr", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DonatorsAddr is a free data retrieval call binding the contract method 0xcbae0e49.
//
// Solidity: function donatorsAddr(uint256 ) view returns(address)
func (_User *UserSession) DonatorsAddr(arg0 *big.Int) (common.Address, error) {
	return _User.Contract.DonatorsAddr(&_User.CallOpts, arg0)
}

// DonatorsAddr is a free data retrieval call binding the contract method 0xcbae0e49.
//
// Solidity: function donatorsAddr(uint256 ) view returns(address)
func (_User *UserCallerSession) DonatorsAddr(arg0 *big.Int) (common.Address, error) {
	return _User.Contract.DonatorsAddr(&_User.CallOpts, arg0)
}

// GetGoods is a free data retrieval call binding the contract method 0xb2590033.
//
// Solidity: function getGoods(uint256 id) view returns(address owner, string name, string species, uint256 rent, uint256 ethPledge)
func (_User *UserCaller) GetGoods(opts *bind.CallOpts, id *big.Int) (struct {
	Owner     common.Address
	Name      string
	Species   string
	Rent      *big.Int
	EthPledge *big.Int
}, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "getGoods", id)

	outstruct := new(struct {
		Owner     common.Address
		Name      string
		Species   string
		Rent      *big.Int
		EthPledge *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Species = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Rent = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.EthPledge = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetGoods is a free data retrieval call binding the contract method 0xb2590033.
//
// Solidity: function getGoods(uint256 id) view returns(address owner, string name, string species, uint256 rent, uint256 ethPledge)
func (_User *UserSession) GetGoods(id *big.Int) (struct {
	Owner     common.Address
	Name      string
	Species   string
	Rent      *big.Int
	EthPledge *big.Int
}, error) {
	return _User.Contract.GetGoods(&_User.CallOpts, id)
}

// GetGoods is a free data retrieval call binding the contract method 0xb2590033.
//
// Solidity: function getGoods(uint256 id) view returns(address owner, string name, string species, uint256 rent, uint256 ethPledge)
func (_User *UserCallerSession) GetGoods(id *big.Int) (struct {
	Owner     common.Address
	Name      string
	Species   string
	Rent      *big.Int
	EthPledge *big.Int
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

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address user) view returns(string name, address people, uint256 integral, string email, string sign, uint256 goodsNum, uint256 balance)
func (_User *UserCaller) GetUser(opts *bind.CallOpts, user common.Address) (struct {
	Name     string
	People   common.Address
	Integral *big.Int
	Email    string
	Sign     string
	GoodsNum *big.Int
	Balance  *big.Int
}, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "getUser", user)

	outstruct := new(struct {
		Name     string
		People   common.Address
		Integral *big.Int
		Email    string
		Sign     string
		GoodsNum *big.Int
		Balance  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.People = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Integral = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Email = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Sign = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.GoodsNum = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Balance = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address user) view returns(string name, address people, uint256 integral, string email, string sign, uint256 goodsNum, uint256 balance)
func (_User *UserSession) GetUser(user common.Address) (struct {
	Name     string
	People   common.Address
	Integral *big.Int
	Email    string
	Sign     string
	GoodsNum *big.Int
	Balance  *big.Int
}, error) {
	return _User.Contract.GetUser(&_User.CallOpts, user)
}

// GetUser is a free data retrieval call binding the contract method 0x6f77926b.
//
// Solidity: function getUser(address user) view returns(string name, address people, uint256 integral, string email, string sign, uint256 goodsNum, uint256 balance)
func (_User *UserCallerSession) GetUser(user common.Address) (struct {
	Name     string
	People   common.Address
	Integral *big.Int
	Email    string
	Sign     string
	GoodsNum *big.Int
	Balance  *big.Int
}, error) {
	return _User.Contract.GetUser(&_User.CallOpts, user)
}

// GoodsData is a free data retrieval call binding the contract method 0xb7305eb2.
//
// Solidity: function goodsData(uint256 ) view returns(address owner, (address,string,string) borrowers, uint256 id, string name, string species, uint256 rent, uint256 ethPledge, uint256 count, uint256 deal, uint256 backs, bool available, bool isBorrow)
func (_User *UserCaller) GoodsData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner     common.Address
	Borrowers UserBorrower
	Id        *big.Int
	Name      string
	Species   string
	Rent      *big.Int
	EthPledge *big.Int
	Count     *big.Int
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
		Count     *big.Int
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
	outstruct.Count = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Deal = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Backs = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.Available = *abi.ConvertType(out[10], new(bool)).(*bool)
	outstruct.IsBorrow = *abi.ConvertType(out[11], new(bool)).(*bool)

	return *outstruct, err

}

// GoodsData is a free data retrieval call binding the contract method 0xb7305eb2.
//
// Solidity: function goodsData(uint256 ) view returns(address owner, (address,string,string) borrowers, uint256 id, string name, string species, uint256 rent, uint256 ethPledge, uint256 count, uint256 deal, uint256 backs, bool available, bool isBorrow)
func (_User *UserSession) GoodsData(arg0 *big.Int) (struct {
	Owner     common.Address
	Borrowers UserBorrower
	Id        *big.Int
	Name      string
	Species   string
	Rent      *big.Int
	EthPledge *big.Int
	Count     *big.Int
	Deal      *big.Int
	Backs     *big.Int
	Available bool
	IsBorrow  bool
}, error) {
	return _User.Contract.GoodsData(&_User.CallOpts, arg0)
}

// GoodsData is a free data retrieval call binding the contract method 0xb7305eb2.
//
// Solidity: function goodsData(uint256 ) view returns(address owner, (address,string,string) borrowers, uint256 id, string name, string species, uint256 rent, uint256 ethPledge, uint256 count, uint256 deal, uint256 backs, bool available, bool isBorrow)
func (_User *UserCallerSession) GoodsData(arg0 *big.Int) (struct {
	Owner     common.Address
	Borrowers UserBorrower
	Id        *big.Int
	Name      string
	Species   string
	Rent      *big.Int
	EthPledge *big.Int
	Count     *big.Int
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

// IntegralData is a free data retrieval call binding the contract method 0x89b21711.
//
// Solidity: function integralData(address ) view returns(string name, string password, address people, uint256 integral, string email, string headImg, string sign, uint256 goodsNum, uint256 balance, bool exist, bool isLogin)
func (_User *UserCaller) IntegralData(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name     string
	Password string
	People   common.Address
	Integral *big.Int
	Email    string
	HeadImg  string
	Sign     string
	GoodsNum *big.Int
	Balance  *big.Int
	Exist    bool
	IsLogin  bool
}, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "integralData", arg0)

	outstruct := new(struct {
		Name     string
		Password string
		People   common.Address
		Integral *big.Int
		Email    string
		HeadImg  string
		Sign     string
		GoodsNum *big.Int
		Balance  *big.Int
		Exist    bool
		IsLogin  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Password = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.People = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Integral = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Email = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.HeadImg = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Sign = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.GoodsNum = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.Balance = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.Exist = *abi.ConvertType(out[9], new(bool)).(*bool)
	outstruct.IsLogin = *abi.ConvertType(out[10], new(bool)).(*bool)

	return *outstruct, err

}

// IntegralData is a free data retrieval call binding the contract method 0x89b21711.
//
// Solidity: function integralData(address ) view returns(string name, string password, address people, uint256 integral, string email, string headImg, string sign, uint256 goodsNum, uint256 balance, bool exist, bool isLogin)
func (_User *UserSession) IntegralData(arg0 common.Address) (struct {
	Name     string
	Password string
	People   common.Address
	Integral *big.Int
	Email    string
	HeadImg  string
	Sign     string
	GoodsNum *big.Int
	Balance  *big.Int
	Exist    bool
	IsLogin  bool
}, error) {
	return _User.Contract.IntegralData(&_User.CallOpts, arg0)
}

// IntegralData is a free data retrieval call binding the contract method 0x89b21711.
//
// Solidity: function integralData(address ) view returns(string name, string password, address people, uint256 integral, string email, string headImg, string sign, uint256 goodsNum, uint256 balance, bool exist, bool isLogin)
func (_User *UserCallerSession) IntegralData(arg0 common.Address) (struct {
	Name     string
	Password string
	People   common.Address
	Integral *big.Int
	Email    string
	HeadImg  string
	Sign     string
	GoodsNum *big.Int
	Balance  *big.Int
	Exist    bool
	IsLogin  bool
}, error) {
	return _User.Contract.IntegralData(&_User.CallOpts, arg0)
}

// IsCommunityExist is a free data retrieval call binding the contract method 0x23dae7f8.
//
// Solidity: function isCommunityExist(uint256 number) view returns(bool)
func (_User *UserCaller) IsCommunityExist(opts *bind.CallOpts, number *big.Int) (bool, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "isCommunityExist", number)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCommunityExist is a free data retrieval call binding the contract method 0x23dae7f8.
//
// Solidity: function isCommunityExist(uint256 number) view returns(bool)
func (_User *UserSession) IsCommunityExist(number *big.Int) (bool, error) {
	return _User.Contract.IsCommunityExist(&_User.CallOpts, number)
}

// IsCommunityExist is a free data retrieval call binding the contract method 0x23dae7f8.
//
// Solidity: function isCommunityExist(uint256 number) view returns(bool)
func (_User *UserCallerSession) IsCommunityExist(number *big.Int) (bool, error) {
	return _User.Contract.IsCommunityExist(&_User.CallOpts, number)
}

// IsGoodExist is a free data retrieval call binding the contract method 0x3827a1a2.
//
// Solidity: function isGoodExist(uint256 id) view returns(bool)
func (_User *UserCaller) IsGoodExist(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "isGoodExist", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGoodExist is a free data retrieval call binding the contract method 0x3827a1a2.
//
// Solidity: function isGoodExist(uint256 id) view returns(bool)
func (_User *UserSession) IsGoodExist(id *big.Int) (bool, error) {
	return _User.Contract.IsGoodExist(&_User.CallOpts, id)
}

// IsGoodExist is a free data retrieval call binding the contract method 0x3827a1a2.
//
// Solidity: function isGoodExist(uint256 id) view returns(bool)
func (_User *UserCallerSession) IsGoodExist(id *big.Int) (bool, error) {
	return _User.Contract.IsGoodExist(&_User.CallOpts, id)
}

// IsGoodsLend is a free data retrieval call binding the contract method 0x7c189689.
//
// Solidity: function isGoodsLend(uint256 id) view returns(bool)
func (_User *UserCaller) IsGoodsLend(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "isGoodsLend", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGoodsLend is a free data retrieval call binding the contract method 0x7c189689.
//
// Solidity: function isGoodsLend(uint256 id) view returns(bool)
func (_User *UserSession) IsGoodsLend(id *big.Int) (bool, error) {
	return _User.Contract.IsGoodsLend(&_User.CallOpts, id)
}

// IsGoodsLend is a free data retrieval call binding the contract method 0x7c189689.
//
// Solidity: function isGoodsLend(uint256 id) view returns(bool)
func (_User *UserCallerSession) IsGoodsLend(id *big.Int) (bool, error) {
	return _User.Contract.IsGoodsLend(&_User.CallOpts, id)
}

// IsStickExist is a free data retrieval call binding the contract method 0x96d89b3a.
//
// Solidity: function isStickExist(string species) view returns(bool)
func (_User *UserCaller) IsStickExist(opts *bind.CallOpts, species string) (bool, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "isStickExist", species)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStickExist is a free data retrieval call binding the contract method 0x96d89b3a.
//
// Solidity: function isStickExist(string species) view returns(bool)
func (_User *UserSession) IsStickExist(species string) (bool, error) {
	return _User.Contract.IsStickExist(&_User.CallOpts, species)
}

// IsStickExist is a free data retrieval call binding the contract method 0x96d89b3a.
//
// Solidity: function isStickExist(string species) view returns(bool)
func (_User *UserCallerSession) IsStickExist(species string) (bool, error) {
	return _User.Contract.IsStickExist(&_User.CallOpts, species)
}

// QueryBalance is a free data retrieval call binding the contract method 0x36f40c61.
//
// Solidity: function queryBalance() view returns(uint256)
func (_User *UserCaller) QueryBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "queryBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueryBalance is a free data retrieval call binding the contract method 0x36f40c61.
//
// Solidity: function queryBalance() view returns(uint256)
func (_User *UserSession) QueryBalance() (*big.Int, error) {
	return _User.Contract.QueryBalance(&_User.CallOpts)
}

// QueryBalance is a free data retrieval call binding the contract method 0x36f40c61.
//
// Solidity: function queryBalance() view returns(uint256)
func (_User *UserCallerSession) QueryBalance() (*big.Int, error) {
	return _User.Contract.QueryBalance(&_User.CallOpts)
}

// QueryBorrower is a free data retrieval call binding the contract method 0xaf3be6e6.
//
// Solidity: function queryBorrower(uint256 id) view returns(address)
func (_User *UserCaller) QueryBorrower(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "queryBorrower", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// QueryBorrower is a free data retrieval call binding the contract method 0xaf3be6e6.
//
// Solidity: function queryBorrower(uint256 id) view returns(address)
func (_User *UserSession) QueryBorrower(id *big.Int) (common.Address, error) {
	return _User.Contract.QueryBorrower(&_User.CallOpts, id)
}

// QueryBorrower is a free data retrieval call binding the contract method 0xaf3be6e6.
//
// Solidity: function queryBorrower(uint256 id) view returns(address)
func (_User *UserCallerSession) QueryBorrower(id *big.Int) (common.Address, error) {
	return _User.Contract.QueryBorrower(&_User.CallOpts, id)
}

// AddCommunity is a paid mutator transaction binding the contract method 0x3f104068.
//
// Solidity: function addCommunity(uint256 number, address beneficiaryAddr, string communityName, string communityIntroduce, uint256 communityAmounts) returns()
func (_User *UserTransactor) AddCommunity(opts *bind.TransactOpts, number *big.Int, beneficiaryAddr common.Address, communityName string, communityIntroduce string, communityAmounts *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "addCommunity", number, beneficiaryAddr, communityName, communityIntroduce, communityAmounts)
}

// AddCommunity is a paid mutator transaction binding the contract method 0x3f104068.
//
// Solidity: function addCommunity(uint256 number, address beneficiaryAddr, string communityName, string communityIntroduce, uint256 communityAmounts) returns()
func (_User *UserSession) AddCommunity(number *big.Int, beneficiaryAddr common.Address, communityName string, communityIntroduce string, communityAmounts *big.Int) (*types.Transaction, error) {
	return _User.Contract.AddCommunity(&_User.TransactOpts, number, beneficiaryAddr, communityName, communityIntroduce, communityAmounts)
}

// AddCommunity is a paid mutator transaction binding the contract method 0x3f104068.
//
// Solidity: function addCommunity(uint256 number, address beneficiaryAddr, string communityName, string communityIntroduce, uint256 communityAmounts) returns()
func (_User *UserTransactorSession) AddCommunity(number *big.Int, beneficiaryAddr common.Address, communityName string, communityIntroduce string, communityAmounts *big.Int) (*types.Transaction, error) {
	return _User.Contract.AddCommunity(&_User.TransactOpts, number, beneficiaryAddr, communityName, communityIntroduce, communityAmounts)
}

// AddGoods is a paid mutator transaction binding the contract method 0x5f4cf227.
//
// Solidity: function addGoods(address owner, string name, string species, uint256 rent, uint256 ethPledge, string[] goodsImgs) returns(uint256)
func (_User *UserTransactor) AddGoods(opts *bind.TransactOpts, owner common.Address, name string, species string, rent *big.Int, ethPledge *big.Int, goodsImgs []string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "addGoods", owner, name, species, rent, ethPledge, goodsImgs)
}

// AddGoods is a paid mutator transaction binding the contract method 0x5f4cf227.
//
// Solidity: function addGoods(address owner, string name, string species, uint256 rent, uint256 ethPledge, string[] goodsImgs) returns(uint256)
func (_User *UserSession) AddGoods(owner common.Address, name string, species string, rent *big.Int, ethPledge *big.Int, goodsImgs []string) (*types.Transaction, error) {
	return _User.Contract.AddGoods(&_User.TransactOpts, owner, name, species, rent, ethPledge, goodsImgs)
}

// AddGoods is a paid mutator transaction binding the contract method 0x5f4cf227.
//
// Solidity: function addGoods(address owner, string name, string species, uint256 rent, uint256 ethPledge, string[] goodsImgs) returns(uint256)
func (_User *UserTransactorSession) AddGoods(owner common.Address, name string, species string, rent *big.Int, ethPledge *big.Int, goodsImgs []string) (*types.Transaction, error) {
	return _User.Contract.AddGoods(&_User.TransactOpts, owner, name, species, rent, ethPledge, goodsImgs)
}

// AddSticker is a paid mutator transaction binding the contract method 0xd360ba7a.
//
// Solidity: function addSticker(string species) returns()
func (_User *UserTransactor) AddSticker(opts *bind.TransactOpts, species string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "addSticker", species)
}

// AddSticker is a paid mutator transaction binding the contract method 0xd360ba7a.
//
// Solidity: function addSticker(string species) returns()
func (_User *UserSession) AddSticker(species string) (*types.Transaction, error) {
	return _User.Contract.AddSticker(&_User.TransactOpts, species)
}

// AddSticker is a paid mutator transaction binding the contract method 0xd360ba7a.
//
// Solidity: function addSticker(string species) returns()
func (_User *UserTransactorSession) AddSticker(species string) (*types.Transaction, error) {
	return _User.Contract.AddSticker(&_User.TransactOpts, species)
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

// AgreeBack is a paid mutator transaction binding the contract method 0x7a865b3f.
//
// Solidity: function agreeBack(uint256 id, uint256 backs) returns()
func (_User *UserTransactor) AgreeBack(opts *bind.TransactOpts, id *big.Int, backs *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "agreeBack", id, backs)
}

// AgreeBack is a paid mutator transaction binding the contract method 0x7a865b3f.
//
// Solidity: function agreeBack(uint256 id, uint256 backs) returns()
func (_User *UserSession) AgreeBack(id *big.Int, backs *big.Int) (*types.Transaction, error) {
	return _User.Contract.AgreeBack(&_User.TransactOpts, id, backs)
}

// AgreeBack is a paid mutator transaction binding the contract method 0x7a865b3f.
//
// Solidity: function agreeBack(uint256 id, uint256 backs) returns()
func (_User *UserTransactorSession) AgreeBack(id *big.Int, backs *big.Int) (*types.Transaction, error) {
	return _User.Contract.AgreeBack(&_User.TransactOpts, id, backs)
}

// AgreeBorrow is a paid mutator transaction binding the contract method 0xaa403a13.
//
// Solidity: function agreeBorrow(uint256 id, uint256 deal, string since, address borrower) returns()
func (_User *UserTransactor) AgreeBorrow(opts *bind.TransactOpts, id *big.Int, deal *big.Int, since string, borrower common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "agreeBorrow", id, deal, since, borrower)
}

// AgreeBorrow is a paid mutator transaction binding the contract method 0xaa403a13.
//
// Solidity: function agreeBorrow(uint256 id, uint256 deal, string since, address borrower) returns()
func (_User *UserSession) AgreeBorrow(id *big.Int, deal *big.Int, since string, borrower common.Address) (*types.Transaction, error) {
	return _User.Contract.AgreeBorrow(&_User.TransactOpts, id, deal, since, borrower)
}

// AgreeBorrow is a paid mutator transaction binding the contract method 0xaa403a13.
//
// Solidity: function agreeBorrow(uint256 id, uint256 deal, string since, address borrower) returns()
func (_User *UserTransactorSession) AgreeBorrow(id *big.Int, deal *big.Int, since string, borrower common.Address) (*types.Transaction, error) {
	return _User.Contract.AgreeBorrow(&_User.TransactOpts, id, deal, since, borrower)
}

// BackGoods is a paid mutator transaction binding the contract method 0x912f3e0f.
//
// Solidity: function backGoods(uint256 id, uint256 backs, string over) payable returns()
func (_User *UserTransactor) BackGoods(opts *bind.TransactOpts, id *big.Int, backs *big.Int, over string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "backGoods", id, backs, over)
}

// BackGoods is a paid mutator transaction binding the contract method 0x912f3e0f.
//
// Solidity: function backGoods(uint256 id, uint256 backs, string over) payable returns()
func (_User *UserSession) BackGoods(id *big.Int, backs *big.Int, over string) (*types.Transaction, error) {
	return _User.Contract.BackGoods(&_User.TransactOpts, id, backs, over)
}

// BackGoods is a paid mutator transaction binding the contract method 0x912f3e0f.
//
// Solidity: function backGoods(uint256 id, uint256 backs, string over) payable returns()
func (_User *UserTransactorSession) BackGoods(id *big.Int, backs *big.Int, over string) (*types.Transaction, error) {
	return _User.Contract.BackGoods(&_User.TransactOpts, id, backs, over)
}

// Borrow is a paid mutator transaction binding the contract method 0x0ecbcdab.
//
// Solidity: function borrow(uint256 id, uint256 deal) payable returns()
func (_User *UserTransactor) Borrow(opts *bind.TransactOpts, id *big.Int, deal *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "borrow", id, deal)
}

// Borrow is a paid mutator transaction binding the contract method 0x0ecbcdab.
//
// Solidity: function borrow(uint256 id, uint256 deal) payable returns()
func (_User *UserSession) Borrow(id *big.Int, deal *big.Int) (*types.Transaction, error) {
	return _User.Contract.Borrow(&_User.TransactOpts, id, deal)
}

// Borrow is a paid mutator transaction binding the contract method 0x0ecbcdab.
//
// Solidity: function borrow(uint256 id, uint256 deal) payable returns()
func (_User *UserTransactorSession) Borrow(id *big.Int, deal *big.Int) (*types.Transaction, error) {
	return _User.Contract.Borrow(&_User.TransactOpts, id, deal)
}

// BorrowGoods is a paid mutator transaction binding the contract method 0xca66f9ae.
//
// Solidity: function borrowGoods(uint256 id, uint256 borrowDays) returns(uint256, address, uint256)
func (_User *UserTransactor) BorrowGoods(opts *bind.TransactOpts, id *big.Int, borrowDays *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "borrowGoods", id, borrowDays)
}

// BorrowGoods is a paid mutator transaction binding the contract method 0xca66f9ae.
//
// Solidity: function borrowGoods(uint256 id, uint256 borrowDays) returns(uint256, address, uint256)
func (_User *UserSession) BorrowGoods(id *big.Int, borrowDays *big.Int) (*types.Transaction, error) {
	return _User.Contract.BorrowGoods(&_User.TransactOpts, id, borrowDays)
}

// BorrowGoods is a paid mutator transaction binding the contract method 0xca66f9ae.
//
// Solidity: function borrowGoods(uint256 id, uint256 borrowDays) returns(uint256, address, uint256)
func (_User *UserTransactorSession) BorrowGoods(id *big.Int, borrowDays *big.Int) (*types.Transaction, error) {
	return _User.Contract.BorrowGoods(&_User.TransactOpts, id, borrowDays)
}

// DelCommunity is a paid mutator transaction binding the contract method 0xc39a2be8.
//
// Solidity: function delCommunity(uint256 number) returns()
func (_User *UserTransactor) DelCommunity(opts *bind.TransactOpts, number *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "delCommunity", number)
}

// DelCommunity is a paid mutator transaction binding the contract method 0xc39a2be8.
//
// Solidity: function delCommunity(uint256 number) returns()
func (_User *UserSession) DelCommunity(number *big.Int) (*types.Transaction, error) {
	return _User.Contract.DelCommunity(&_User.TransactOpts, number)
}

// DelCommunity is a paid mutator transaction binding the contract method 0xc39a2be8.
//
// Solidity: function delCommunity(uint256 number) returns()
func (_User *UserTransactorSession) DelCommunity(number *big.Int) (*types.Transaction, error) {
	return _User.Contract.DelCommunity(&_User.TransactOpts, number)
}

// DelGoods is a paid mutator transaction binding the contract method 0xeb8d6a96.
//
// Solidity: function delGoods(uint256 id) returns()
func (_User *UserTransactor) DelGoods(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "delGoods", id)
}

// DelGoods is a paid mutator transaction binding the contract method 0xeb8d6a96.
//
// Solidity: function delGoods(uint256 id) returns()
func (_User *UserSession) DelGoods(id *big.Int) (*types.Transaction, error) {
	return _User.Contract.DelGoods(&_User.TransactOpts, id)
}

// DelGoods is a paid mutator transaction binding the contract method 0xeb8d6a96.
//
// Solidity: function delGoods(uint256 id) returns()
func (_User *UserTransactorSession) DelGoods(id *big.Int) (*types.Transaction, error) {
	return _User.Contract.DelGoods(&_User.TransactOpts, id)
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

// Donate is a paid mutator transaction binding the contract method 0x7f15463e.
//
// Solidity: function donate(uint256 number, address donator, uint256 amount, string donatorName) payable returns(uint256)
func (_User *UserTransactor) Donate(opts *bind.TransactOpts, number *big.Int, donator common.Address, amount *big.Int, donatorName string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "donate", number, donator, amount, donatorName)
}

// Donate is a paid mutator transaction binding the contract method 0x7f15463e.
//
// Solidity: function donate(uint256 number, address donator, uint256 amount, string donatorName) payable returns(uint256)
func (_User *UserSession) Donate(number *big.Int, donator common.Address, amount *big.Int, donatorName string) (*types.Transaction, error) {
	return _User.Contract.Donate(&_User.TransactOpts, number, donator, amount, donatorName)
}

// Donate is a paid mutator transaction binding the contract method 0x7f15463e.
//
// Solidity: function donate(uint256 number, address donator, uint256 amount, string donatorName) payable returns(uint256)
func (_User *UserTransactorSession) Donate(number *big.Int, donator common.Address, amount *big.Int, donatorName string) (*types.Transaction, error) {
	return _User.Contract.Donate(&_User.TransactOpts, number, donator, amount, donatorName)
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
