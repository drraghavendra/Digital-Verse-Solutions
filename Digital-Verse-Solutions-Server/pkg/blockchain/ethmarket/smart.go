// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethmarket

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

// DvNftMarketMarketItem is an auto generated low-level Go binding around an user-defined struct.
type DvNftMarketMarketItem struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	Sold        bool
}

// EthNftMarketABI is the input ABI used to generate the binding from.
const EthNftMarketABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"name\":\"MarketItemCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"createMarketItem\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"}],\"name\":\"createMarketSale\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchItemsCreated\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"internalType\":\"structDvNftMarket.MarketItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchMarketItems\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"internalType\":\"structDvNftMarket.MarketItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fetchMyNFTs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"itemId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"internalType\":\"structDvNftMarket.MarketItem[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getListingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EthNftMarket is an auto generated Go binding around an Ethereum contract.
type EthNftMarket struct {
	EthNftMarketCaller     // Read-only binding to the contract
	EthNftMarketTransactor // Write-only binding to the contract
	EthNftMarketFilterer   // Log filterer for contract events
}

// EthNftMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthNftMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthNftMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthNftMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthNftMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthNftMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthNftMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthNftMarketSession struct {
	Contract     *EthNftMarket     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthNftMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthNftMarketCallerSession struct {
	Contract *EthNftMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EthNftMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthNftMarketTransactorSession struct {
	Contract     *EthNftMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EthNftMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthNftMarketRaw struct {
	Contract *EthNftMarket // Generic contract binding to access the raw methods on
}

// EthNftMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthNftMarketCallerRaw struct {
	Contract *EthNftMarketCaller // Generic read-only contract binding to access the raw methods on
}

// EthNftMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthNftMarketTransactorRaw struct {
	Contract *EthNftMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthNftMarket creates a new instance of EthNftMarket, bound to a specific deployed contract.
func NewEthNftMarket(address common.Address, backend bind.ContractBackend) (*EthNftMarket, error) {
	contract, err := bindEthNftMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthNftMarket{EthNftMarketCaller: EthNftMarketCaller{contract: contract}, EthNftMarketTransactor: EthNftMarketTransactor{contract: contract}, EthNftMarketFilterer: EthNftMarketFilterer{contract: contract}}, nil
}

// NewEthNftMarketCaller creates a new read-only instance of EthNftMarket, bound to a specific deployed contract.
func NewEthNftMarketCaller(address common.Address, caller bind.ContractCaller) (*EthNftMarketCaller, error) {
	contract, err := bindEthNftMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthNftMarketCaller{contract: contract}, nil
}

// NewEthNftMarketTransactor creates a new write-only instance of EthNftMarket, bound to a specific deployed contract.
func NewEthNftMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*EthNftMarketTransactor, error) {
	contract, err := bindEthNftMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthNftMarketTransactor{contract: contract}, nil
}

// NewEthNftMarketFilterer creates a new log filterer instance of EthNftMarket, bound to a specific deployed contract.
func NewEthNftMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*EthNftMarketFilterer, error) {
	contract, err := bindEthNftMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthNftMarketFilterer{contract: contract}, nil
}

// bindEthNftMarket binds a generic wrapper to an already deployed contract.
func bindEthNftMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthNftMarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthNftMarket *EthNftMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthNftMarket.Contract.EthNftMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthNftMarket *EthNftMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthNftMarket.Contract.EthNftMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthNftMarket *EthNftMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthNftMarket.Contract.EthNftMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthNftMarket *EthNftMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthNftMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthNftMarket *EthNftMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthNftMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthNftMarket *EthNftMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthNftMarket.Contract.contract.Transact(opts, method, params...)
}

// FetchItemsCreated is a free data retrieval call binding the contract method 0xf064c32e.
//
// Solidity: function fetchItemsCreated() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_EthNftMarket *EthNftMarketCaller) FetchItemsCreated(opts *bind.CallOpts) ([]DvNftMarketMarketItem, error) {
	var out []interface{}
	err := _EthNftMarket.contract.Call(opts, &out, "fetchItemsCreated")

	if err != nil {
		return *new([]DvNftMarketMarketItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]DvNftMarketMarketItem)).(*[]DvNftMarketMarketItem)

	return out0, err

}

// FetchItemsCreated is a free data retrieval call binding the contract method 0xf064c32e.
//
// Solidity: function fetchItemsCreated() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_EthNftMarket *EthNftMarketSession) FetchItemsCreated() ([]DvNftMarketMarketItem, error) {
	return _EthNftMarket.Contract.FetchItemsCreated(&_EthNftMarket.CallOpts)
}

// FetchItemsCreated is a free data retrieval call binding the contract method 0xf064c32e.
//
// Solidity: function fetchItemsCreated() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_EthNftMarket *EthNftMarketCallerSession) FetchItemsCreated() ([]DvNftMarketMarketItem, error) {
	return _EthNftMarket.Contract.FetchItemsCreated(&_EthNftMarket.CallOpts)
}

// FetchMarketItems is a free data retrieval call binding the contract method 0x0f08efe0.
//
// Solidity: function fetchMarketItems() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_EthNftMarket *EthNftMarketCaller) FetchMarketItems(opts *bind.CallOpts) ([]DvNftMarketMarketItem, error) {
	var out []interface{}
	err := _EthNftMarket.contract.Call(opts, &out, "fetchMarketItems")

	if err != nil {
		return *new([]DvNftMarketMarketItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]DvNftMarketMarketItem)).(*[]DvNftMarketMarketItem)

	return out0, err

}

// FetchMarketItems is a free data retrieval call binding the contract method 0x0f08efe0.
//
// Solidity: function fetchMarketItems() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_EthNftMarket *EthNftMarketSession) FetchMarketItems() ([]DvNftMarketMarketItem, error) {
	return _EthNftMarket.Contract.FetchMarketItems(&_EthNftMarket.CallOpts)
}

// FetchMarketItems is a free data retrieval call binding the contract method 0x0f08efe0.
//
// Solidity: function fetchMarketItems() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_EthNftMarket *EthNftMarketCallerSession) FetchMarketItems() ([]DvNftMarketMarketItem, error) {
	return _EthNftMarket.Contract.FetchMarketItems(&_EthNftMarket.CallOpts)
}

// FetchMyNFTs is a free data retrieval call binding the contract method 0x202e3740.
//
// Solidity: function fetchMyNFTs() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_EthNftMarket *EthNftMarketCaller) FetchMyNFTs(opts *bind.CallOpts) ([]DvNftMarketMarketItem, error) {
	var out []interface{}
	err := _EthNftMarket.contract.Call(opts, &out, "fetchMyNFTs")

	if err != nil {
		return *new([]DvNftMarketMarketItem), err
	}

	out0 := *abi.ConvertType(out[0], new([]DvNftMarketMarketItem)).(*[]DvNftMarketMarketItem)

	return out0, err

}

// FetchMyNFTs is a free data retrieval call binding the contract method 0x202e3740.
//
// Solidity: function fetchMyNFTs() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_EthNftMarket *EthNftMarketSession) FetchMyNFTs() ([]DvNftMarketMarketItem, error) {
	return _EthNftMarket.Contract.FetchMyNFTs(&_EthNftMarket.CallOpts)
}

// FetchMyNFTs is a free data retrieval call binding the contract method 0x202e3740.
//
// Solidity: function fetchMyNFTs() view returns((uint256,address,uint256,address,address,uint256,bool)[])
func (_EthNftMarket *EthNftMarketCallerSession) FetchMyNFTs() ([]DvNftMarketMarketItem, error) {
	return _EthNftMarket.Contract.FetchMyNFTs(&_EthNftMarket.CallOpts)
}

// GetListingPrice is a free data retrieval call binding the contract method 0x12e85585.
//
// Solidity: function getListingPrice() view returns(uint256)
func (_EthNftMarket *EthNftMarketCaller) GetListingPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthNftMarket.contract.Call(opts, &out, "getListingPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetListingPrice is a free data retrieval call binding the contract method 0x12e85585.
//
// Solidity: function getListingPrice() view returns(uint256)
func (_EthNftMarket *EthNftMarketSession) GetListingPrice() (*big.Int, error) {
	return _EthNftMarket.Contract.GetListingPrice(&_EthNftMarket.CallOpts)
}

// GetListingPrice is a free data retrieval call binding the contract method 0x12e85585.
//
// Solidity: function getListingPrice() view returns(uint256)
func (_EthNftMarket *EthNftMarketCallerSession) GetListingPrice() (*big.Int, error) {
	return _EthNftMarket.Contract.GetListingPrice(&_EthNftMarket.CallOpts)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) payable returns()
func (_EthNftMarket *EthNftMarketTransactor) CreateMarketItem(opts *bind.TransactOpts, nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _EthNftMarket.contract.Transact(opts, "createMarketItem", nftContract, tokenId, price)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) payable returns()
func (_EthNftMarket *EthNftMarketSession) CreateMarketItem(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _EthNftMarket.Contract.CreateMarketItem(&_EthNftMarket.TransactOpts, nftContract, tokenId, price)
}

// CreateMarketItem is a paid mutator transaction binding the contract method 0x58eb2df5.
//
// Solidity: function createMarketItem(address nftContract, uint256 tokenId, uint256 price) payable returns()
func (_EthNftMarket *EthNftMarketTransactorSession) CreateMarketItem(nftContract common.Address, tokenId *big.Int, price *big.Int) (*types.Transaction, error) {
	return _EthNftMarket.Contract.CreateMarketItem(&_EthNftMarket.TransactOpts, nftContract, tokenId, price)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xc23b139e.
//
// Solidity: function createMarketSale(address nftContract, uint256 itemId) payable returns()
func (_EthNftMarket *EthNftMarketTransactor) CreateMarketSale(opts *bind.TransactOpts, nftContract common.Address, itemId *big.Int) (*types.Transaction, error) {
	return _EthNftMarket.contract.Transact(opts, "createMarketSale", nftContract, itemId)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xc23b139e.
//
// Solidity: function createMarketSale(address nftContract, uint256 itemId) payable returns()
func (_EthNftMarket *EthNftMarketSession) CreateMarketSale(nftContract common.Address, itemId *big.Int) (*types.Transaction, error) {
	return _EthNftMarket.Contract.CreateMarketSale(&_EthNftMarket.TransactOpts, nftContract, itemId)
}

// CreateMarketSale is a paid mutator transaction binding the contract method 0xc23b139e.
//
// Solidity: function createMarketSale(address nftContract, uint256 itemId) payable returns()
func (_EthNftMarket *EthNftMarketTransactorSession) CreateMarketSale(nftContract common.Address, itemId *big.Int) (*types.Transaction, error) {
	return _EthNftMarket.Contract.CreateMarketSale(&_EthNftMarket.TransactOpts, nftContract, itemId)
}

// EthNftMarketMarketItemCreatedIterator is returned from FilterMarketItemCreated and is used to iterate over the raw logs and unpacked data for MarketItemCreated events raised by the EthNftMarket contract.
type EthNftMarketMarketItemCreatedIterator struct {
	Event *EthNftMarketMarketItemCreated // Event containing the contract specifics and raw log

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
func (it *EthNftMarketMarketItemCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthNftMarketMarketItemCreated)
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
		it.Event = new(EthNftMarketMarketItemCreated)
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
func (it *EthNftMarketMarketItemCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthNftMarketMarketItemCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthNftMarketMarketItemCreated represents a MarketItemCreated event raised by the EthNftMarket contract.
type EthNftMarketMarketItemCreated struct {
	ItemId      *big.Int
	NftContract common.Address
	TokenId     *big.Int
	Seller      common.Address
	Owner       common.Address
	Price       *big.Int
	Sold        bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketItemCreated is a free log retrieval operation binding the contract event 0x045dfa01dcba2b36aba1d3dc4a874f4b0c5d2fbeb8d2c4b34a7d88c8d8f929d1.
//
// Solidity: event MarketItemCreated(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address seller, address owner, uint256 price, bool sold)
func (_EthNftMarket *EthNftMarketFilterer) FilterMarketItemCreated(opts *bind.FilterOpts, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (*EthNftMarketMarketItemCreatedIterator, error) {

	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _EthNftMarket.contract.FilterLogs(opts, "MarketItemCreated", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &EthNftMarketMarketItemCreatedIterator{contract: _EthNftMarket.contract, event: "MarketItemCreated", logs: logs, sub: sub}, nil
}

// WatchMarketItemCreated is a free log subscription operation binding the contract event 0x045dfa01dcba2b36aba1d3dc4a874f4b0c5d2fbeb8d2c4b34a7d88c8d8f929d1.
//
// Solidity: event MarketItemCreated(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address seller, address owner, uint256 price, bool sold)
func (_EthNftMarket *EthNftMarketFilterer) WatchMarketItemCreated(opts *bind.WatchOpts, sink chan<- *EthNftMarketMarketItemCreated, itemId []*big.Int, nftContract []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var itemIdRule []interface{}
	for _, itemIdItem := range itemId {
		itemIdRule = append(itemIdRule, itemIdItem)
	}
	var nftContractRule []interface{}
	for _, nftContractItem := range nftContract {
		nftContractRule = append(nftContractRule, nftContractItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _EthNftMarket.contract.WatchLogs(opts, "MarketItemCreated", itemIdRule, nftContractRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthNftMarketMarketItemCreated)
				if err := _EthNftMarket.contract.UnpackLog(event, "MarketItemCreated", log); err != nil {
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

// ParseMarketItemCreated is a log parse operation binding the contract event 0x045dfa01dcba2b36aba1d3dc4a874f4b0c5d2fbeb8d2c4b34a7d88c8d8f929d1.
//
// Solidity: event MarketItemCreated(uint256 indexed itemId, address indexed nftContract, uint256 indexed tokenId, address seller, address owner, uint256 price, bool sold)
func (_EthNftMarket *EthNftMarketFilterer) ParseMarketItemCreated(log types.Log) (*EthNftMarketMarketItemCreated, error) {
	event := new(EthNftMarketMarketItemCreated)
	if err := _EthNftMarket.contract.UnpackLog(event, "MarketItemCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
