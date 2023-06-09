// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dist

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

// DistMetaData contains all meta data concerning the Dist contract.
var DistMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20VotesUpgradeable\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_sweepReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_claimPeriodStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_claimPeriodEnd\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegateTo\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CanClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HasClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSweepReceiver\",\"type\":\"address\"}],\"name\":\"SweepReceiverSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Swept\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"claimAndDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimPeriodEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimPeriodStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimableTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_claimableAmount\",\"type\":\"uint256[]\"}],\"name\":\"setRecipients\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_sweepReceiver\",\"type\":\"address\"}],\"name\":\"setSweepReciever\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sweepReceiver\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20VotesUpgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalClaimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e06040523480156200001157600080fd5b5060405162001b8038038062001b80833981016040819052620000349162000436565b6200003f3362000315565b6001600160a01b038616620000a75760405162461bcd60e51b8152602060048201526024808201527f546f6b656e4469737472696275746f723a207a65726f20746f6b656e206164646044820152637265737360e01b60648201526084015b60405180910390fd5b6001600160a01b0385166200010b5760405162461bcd60e51b8152602060048201526024808201527f546f6b656e4469737472696275746f723a207a65726f207377656570206164646044820152637265737360e01b60648201526084016200009e565b6001600160a01b0384166200016f5760405162461bcd60e51b8152602060048201526024808201527f546f6b656e4469737472696275746f723a207a65726f206f776e6572206164646044820152637265737360e01b60648201526084016200009e565b438311620001c75760405162461bcd60e51b815260206004820152602f602482015260008051602062001b6083398151915260448201526e6520696e207468652066757475726560881b60648201526084016200009e565b8282116200021c5760405162461bcd60e51b815260206004820152602c602482015260008051602062001b6083398151915260448201526b19481899599bdc9948195b9960a21b60648201526084016200009e565b6001600160a01b0381166200027f5760405162461bcd60e51b815260206004820152602260248201527f546f6b656e4469737472696275746f723a207a65726f2064656c656761746520604482015261746f60f01b60648201526084016200009e565b6040516317066a5760e21b81526001600160a01b038281166004830152871690635c19a95c90602401600060405180830381600087803b158015620002c357600080fd5b505af1158015620002d8573d6000803e3d6000fd5b5050506001600160a01b03871660805250620002f48562000365565b60a083905260c0829052620003098462000315565b505050505050620004b3565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b038116620003d35760405162461bcd60e51b815260206004820152602d60248201527f546f6b656e4469737472696275746f723a207a65726f2073776565702072656360448201526c6569766572206164647265737360981b60648201526084016200009e565b600180546001600160a01b0319166001600160a01b0383169081179091556040517fbea8251f76064f657f2a744bf08a1b5700486d06eb94922162892eb022d95ef690600090a250565b6001600160a01b03811681146200043357600080fd5b50565b60008060008060008060c087890312156200045057600080fd5b86516200045d816200041d565b602088015190965062000470816200041d565b604088015190955062000483816200041d565b80945050606087015192506080870151915060a0870151620004a5816200041d565b809150509295509295509295565b60805160a05160c051611637620005296000396000818161011c0152818161040e01526107fd015260008181610167015261074d015260008181610261015281816102c0015281816104c7015281816106020152818161097f01528181610b3001528181610bd20152610fc601526116376000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806378e2b59411610097578063b438abde11610066578063b438abde14610216578063f2fde38b14610229578063f6e0df9f1461023c578063fc0c546a1461025c57600080fd5b806378e2b5941461019157806384d24226146101a45780638da5cb5b146101c4578063ae373c1b1461020357600080fd5b80634838ed19116100d35780634838ed19146101515780634e71d92d1461015a57806358c13b7e14610162578063715018a61461018957600080fd5b80632e1a7d4d146100fa57806335faa4161461010f5780633da082a014610117575b600080fd5b61010d6101083660046113ac565b610283565b005b61010d61040c565b61013e7f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020015b60405180910390f35b61013e60035481565b61010d61074b565b61013e7f000000000000000000000000000000000000000000000000000000000000000081565b61010d610aaa565b61010d61019f3660046113e7565b610abe565b61013e6101b236600461143f565b60026020526000908152604090205481565b60005473ffffffffffffffffffffffffffffffffffffffff165b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610148565b61010d6102113660046114af565b610ce8565b61010d61022436600461143f565b6110dc565b61010d61023736600461143f565b6110f0565b6001546101de9073ffffffffffffffffffffffffffffffffffffffff1681565b6101de7f000000000000000000000000000000000000000000000000000000000000000081565b61028b6111a4565b6040517fa9059cbb000000000000000000000000000000000000000000000000000000008152336004820152602481018290527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff169063a9059cbb906044016020604051808303816000875af115801561031e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610342919061151b565b6103d3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f546f6b656e4469737472696275746f723a206661696c207472616e736665722060448201527f746f6b656e00000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b60405181815233907f7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65906020015b60405180910390a250565b7f0000000000000000000000000000000000000000000000000000000000000000431015610496576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f546f6b656e4469737472696275746f723a206e6f7420656e646564000000000060448201526064016103ca565b6040517f70a082310000000000000000000000000000000000000000000000000000000081523060048201526000907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015610523573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610547919061153d565b9050806000036105b3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f546f6b656e4469737472696275746f723a206e6f206c6566746f76657273000060448201526064016103ca565b6001546040517fa9059cbb00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff9182166004820152602481018390527f00000000000000000000000000000000000000000000000000000000000000009091169063a9059cbb906044016020604051808303816000875af115801561064d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610671919061151b565b6106fd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f546f6b656e4469737472696275746f723a206661696c20746f6b656e2074726160448201527f6e7366657200000000000000000000000000000000000000000000000000000060648201526084016103ca565b6040518181527f7f221332ee403570bf4d61630b58189ea566ff1635269001e9df6a890f413dd89060200160405180910390a160015473ffffffffffffffffffffffffffffffffffffffff16ff5b7f00000000000000000000000000000000000000000000000000000000000000004310156107fb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f546f6b656e4469737472696275746f723a20636c61696d206e6f74207374617260448201527f746564000000000000000000000000000000000000000000000000000000000060648201526084016103ca565b7f00000000000000000000000000000000000000000000000000000000000000004310610884576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f546f6b656e4469737472696275746f723a20636c61696d20656e64656400000060448201526064016103ca565b3360009081526002602052604090205480610921576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602260248201527f546f6b656e4469737472696275746f723a206e6f7468696e6720746f20636c6160448201527f696d00000000000000000000000000000000000000000000000000000000000060648201526084016103ca565b3360008181526002602052604080822091909155517fa9059cbb00000000000000000000000000000000000000000000000000000000815260048101919091526024810182905273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb906044016020604051808303816000875af11580156109c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109ec919061151b565b610a78576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f546f6b656e4469737472696275746f723a206661696c20746f6b656e2074726160448201527f6e7366657200000000000000000000000000000000000000000000000000000060648201526084016103ca565b60405181815233907f8629b200ebe43db58ad688b85131d53251f3f3be4c14933b4641aeebacf1c08c90602001610401565b610ab26111a4565b610abc6000611225565b565b610ac661074b565b6040517fc3cda52000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8681166004830152600060248301526044820186905260ff851660648301526084820184905260a482018390527f0000000000000000000000000000000000000000000000000000000000000000169063c3cda5209060c401600060405180830381600087803b158015610b7457600080fd5b505af1158015610b88573d6000803e3d6000fd5b50506040517f587cde1e00000000000000000000000000000000000000000000000000000000815233600482015273ffffffffffffffffffffffffffffffffffffffff88811693507f000000000000000000000000000000000000000000000000000000000000000016915063587cde1e90602401602060405180830381865afa158015610c1a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c3e9190611556565b73ffffffffffffffffffffffffffffffffffffffff1614610ce1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f546f6b656e4469737472696275746f723a2064656c6567617465206661696c6560448201527f640000000000000000000000000000000000000000000000000000000000000060648201526084016103ca565b5050505050565b610cf06111a4565b828114610d7f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f546f6b656e4469737472696275746f723a20696e76616c69642061727261792060448201527f6c656e677468000000000000000000000000000000000000000000000000000060648201526084016103ca565b60035460005b84811015610f955760026000878784818110610da357610da3611573565b9050602002016020810190610db8919061143f565b73ffffffffffffffffffffffffffffffffffffffff16815260208101919091526040016000205415610e6c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602760248201527f546f6b656e4469737472696275746f723a20726563697069656e7420616c726560448201527f616479207365740000000000000000000000000000000000000000000000000060648201526084016103ca565b838382818110610e7e57610e7e611573565b9050602002013560026000888885818110610e9b57610e9b611573565b9050602002016020810190610eb0919061143f565b73ffffffffffffffffffffffffffffffffffffffff168152602081019190915260400160002055858582818110610ee957610ee9611573565b9050602002016020810190610efe919061143f565b73ffffffffffffffffffffffffffffffffffffffff167f87aeeb9eda09a064caef63d00f62c15063631980bfc422ad7dd30c8a79f0cbb7858584818110610f4757610f47611573565b90506020020135604051610f5d91815260200190565b60405180910390a2838382818110610f7757610f77611573565b90506020020135820191508080610f8d906115a2565b915050610d85565b506040517f70a0823100000000000000000000000000000000000000000000000000000000815230600482015281907f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16906370a0823190602401602060405180830381865afa158015611022573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611046919061153d565b10156110d3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f546f6b656e4469737472696275746f723a206e6f7420656e6f7567682062616c60448201527f616e63650000000000000000000000000000000000000000000000000000000060648201526084016103ca565b60035550505050565b6110e46111a4565b6110ed8161129a565b50565b6110f86111a4565b73ffffffffffffffffffffffffffffffffffffffff811661119b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016103ca565b6110ed81611225565b60005473ffffffffffffffffffffffffffffffffffffffff163314610abc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103ca565b6000805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b73ffffffffffffffffffffffffffffffffffffffff811661133d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602d60248201527f546f6b656e4469737472696275746f723a207a65726f2073776565702072656360448201527f656976657220616464726573730000000000000000000000000000000000000060648201526084016103ca565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081179091556040517fbea8251f76064f657f2a744bf08a1b5700486d06eb94922162892eb022d95ef690600090a250565b6000602082840312156113be57600080fd5b5035919050565b73ffffffffffffffffffffffffffffffffffffffff811681146110ed57600080fd5b600080600080600060a086880312156113ff57600080fd5b853561140a816113c5565b945060208601359350604086013560ff8116811461142757600080fd5b94979396509394606081013594506080013592915050565b60006020828403121561145157600080fd5b813561145c816113c5565b9392505050565b60008083601f84011261147557600080fd5b50813567ffffffffffffffff81111561148d57600080fd5b6020830191508360208260051b85010111156114a857600080fd5b9250929050565b600080600080604085870312156114c557600080fd5b843567ffffffffffffffff808211156114dd57600080fd5b6114e988838901611463565b9096509450602087013591508082111561150257600080fd5b5061150f87828801611463565b95989497509550505050565b60006020828403121561152d57600080fd5b8151801515811461145c57600080fd5b60006020828403121561154f57600080fd5b5051919050565b60006020828403121561156857600080fd5b815161145c816113c5565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036115fa577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea26469706673582212200f8c309fa553f24bef509e94696c8c86f4c28ecc8d5acb1810aa60a75a5c339064736f6c63430008100033546f6b656e4469737472696275746f723a2073746172742073686f756c642062000000000000000000000000912ce59144191c1204e64559fe8253a0e49e6548000000000000000000000000bfc1feca8b09a5c5d3effe7429ebe24b9c09ef580000000000000000000000002b9acfd85440b7828db8e54694ee07b2b056b30c000000000000000000000000000000000000000000000000000000000101ba20000000000000000000000000000000000000000000000000000000000115d50000000000000000000000000000000000000000000000000000000000000a4b86",
}

// DistABI is the input ABI used to generate the binding from.
// Deprecated: Use DistMetaData.ABI instead.
var DistABI = DistMetaData.ABI

// DistBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DistMetaData.Bin instead.
var DistBin = DistMetaData.Bin

// DeployDist deploys a new Ethereum contract, binding an instance of Dist to it.
func DeployDist(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _sweepReceiver common.Address, _owner common.Address, _claimPeriodStart *big.Int, _claimPeriodEnd *big.Int, delegateTo common.Address) (common.Address, *types.Transaction, *Dist, error) {
	parsed, err := DistMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DistBin), backend, _token, _sweepReceiver, _owner, _claimPeriodStart, _claimPeriodEnd, delegateTo)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Dist{DistCaller: DistCaller{contract: contract}, DistTransactor: DistTransactor{contract: contract}, DistFilterer: DistFilterer{contract: contract}}, nil
}

// Dist is an auto generated Go binding around an Ethereum contract.
type Dist struct {
	DistCaller     // Read-only binding to the contract
	DistTransactor // Write-only binding to the contract
	DistFilterer   // Log filterer for contract events
}

// DistCaller is an auto generated read-only Go binding around an Ethereum contract.
type DistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DistSession struct {
	Contract     *Dist             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DistCallerSession struct {
	Contract *DistCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DistTransactorSession struct {
	Contract     *DistTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DistRaw is an auto generated low-level Go binding around an Ethereum contract.
type DistRaw struct {
	Contract *Dist // Generic contract binding to access the raw methods on
}

// DistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DistCallerRaw struct {
	Contract *DistCaller // Generic read-only contract binding to access the raw methods on
}

// DistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DistTransactorRaw struct {
	Contract *DistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDist creates a new instance of Dist, bound to a specific deployed contract.
func NewDist(address common.Address, backend bind.ContractBackend) (*Dist, error) {
	contract, err := bindDist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Dist{DistCaller: DistCaller{contract: contract}, DistTransactor: DistTransactor{contract: contract}, DistFilterer: DistFilterer{contract: contract}}, nil
}

// NewDistCaller creates a new read-only instance of Dist, bound to a specific deployed contract.
func NewDistCaller(address common.Address, caller bind.ContractCaller) (*DistCaller, error) {
	contract, err := bindDist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DistCaller{contract: contract}, nil
}

// NewDistTransactor creates a new write-only instance of Dist, bound to a specific deployed contract.
func NewDistTransactor(address common.Address, transactor bind.ContractTransactor) (*DistTransactor, error) {
	contract, err := bindDist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DistTransactor{contract: contract}, nil
}

// NewDistFilterer creates a new log filterer instance of Dist, bound to a specific deployed contract.
func NewDistFilterer(address common.Address, filterer bind.ContractFilterer) (*DistFilterer, error) {
	contract, err := bindDist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DistFilterer{contract: contract}, nil
}

// bindDist binds a generic wrapper to an already deployed contract.
func bindDist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DistABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dist *DistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dist.Contract.DistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dist *DistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dist.Contract.DistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dist *DistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dist.Contract.DistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Dist *DistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Dist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Dist *DistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Dist *DistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Dist.Contract.contract.Transact(opts, method, params...)
}

// ClaimPeriodEnd is a free data retrieval call binding the contract method 0x3da082a0.
//
// Solidity: function claimPeriodEnd() view returns(uint256)
func (_Dist *DistCaller) ClaimPeriodEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dist.contract.Call(opts, &out, "claimPeriodEnd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimPeriodEnd is a free data retrieval call binding the contract method 0x3da082a0.
//
// Solidity: function claimPeriodEnd() view returns(uint256)
func (_Dist *DistSession) ClaimPeriodEnd() (*big.Int, error) {
	return _Dist.Contract.ClaimPeriodEnd(&_Dist.CallOpts)
}

// ClaimPeriodEnd is a free data retrieval call binding the contract method 0x3da082a0.
//
// Solidity: function claimPeriodEnd() view returns(uint256)
func (_Dist *DistCallerSession) ClaimPeriodEnd() (*big.Int, error) {
	return _Dist.Contract.ClaimPeriodEnd(&_Dist.CallOpts)
}

// ClaimPeriodStart is a free data retrieval call binding the contract method 0x58c13b7e.
//
// Solidity: function claimPeriodStart() view returns(uint256)
func (_Dist *DistCaller) ClaimPeriodStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dist.contract.Call(opts, &out, "claimPeriodStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimPeriodStart is a free data retrieval call binding the contract method 0x58c13b7e.
//
// Solidity: function claimPeriodStart() view returns(uint256)
func (_Dist *DistSession) ClaimPeriodStart() (*big.Int, error) {
	return _Dist.Contract.ClaimPeriodStart(&_Dist.CallOpts)
}

// ClaimPeriodStart is a free data retrieval call binding the contract method 0x58c13b7e.
//
// Solidity: function claimPeriodStart() view returns(uint256)
func (_Dist *DistCallerSession) ClaimPeriodStart() (*big.Int, error) {
	return _Dist.Contract.ClaimPeriodStart(&_Dist.CallOpts)
}

// ClaimableTokens is a free data retrieval call binding the contract method 0x84d24226.
//
// Solidity: function claimableTokens(address ) view returns(uint256)
func (_Dist *DistCaller) ClaimableTokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Dist.contract.Call(opts, &out, "claimableTokens", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableTokens is a free data retrieval call binding the contract method 0x84d24226.
//
// Solidity: function claimableTokens(address ) view returns(uint256)
func (_Dist *DistSession) ClaimableTokens(arg0 common.Address) (*big.Int, error) {
	return _Dist.Contract.ClaimableTokens(&_Dist.CallOpts, arg0)
}

// ClaimableTokens is a free data retrieval call binding the contract method 0x84d24226.
//
// Solidity: function claimableTokens(address ) view returns(uint256)
func (_Dist *DistCallerSession) ClaimableTokens(arg0 common.Address) (*big.Int, error) {
	return _Dist.Contract.ClaimableTokens(&_Dist.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dist *DistCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dist.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dist *DistSession) Owner() (common.Address, error) {
	return _Dist.Contract.Owner(&_Dist.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Dist *DistCallerSession) Owner() (common.Address, error) {
	return _Dist.Contract.Owner(&_Dist.CallOpts)
}

// SweepReceiver is a free data retrieval call binding the contract method 0xf6e0df9f.
//
// Solidity: function sweepReceiver() view returns(address)
func (_Dist *DistCaller) SweepReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dist.contract.Call(opts, &out, "sweepReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SweepReceiver is a free data retrieval call binding the contract method 0xf6e0df9f.
//
// Solidity: function sweepReceiver() view returns(address)
func (_Dist *DistSession) SweepReceiver() (common.Address, error) {
	return _Dist.Contract.SweepReceiver(&_Dist.CallOpts)
}

// SweepReceiver is a free data retrieval call binding the contract method 0xf6e0df9f.
//
// Solidity: function sweepReceiver() view returns(address)
func (_Dist *DistCallerSession) SweepReceiver() (common.Address, error) {
	return _Dist.Contract.SweepReceiver(&_Dist.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Dist *DistCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Dist.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Dist *DistSession) Token() (common.Address, error) {
	return _Dist.Contract.Token(&_Dist.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Dist *DistCallerSession) Token() (common.Address, error) {
	return _Dist.Contract.Token(&_Dist.CallOpts)
}

// TotalClaimable is a free data retrieval call binding the contract method 0x4838ed19.
//
// Solidity: function totalClaimable() view returns(uint256)
func (_Dist *DistCaller) TotalClaimable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Dist.contract.Call(opts, &out, "totalClaimable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimable is a free data retrieval call binding the contract method 0x4838ed19.
//
// Solidity: function totalClaimable() view returns(uint256)
func (_Dist *DistSession) TotalClaimable() (*big.Int, error) {
	return _Dist.Contract.TotalClaimable(&_Dist.CallOpts)
}

// TotalClaimable is a free data retrieval call binding the contract method 0x4838ed19.
//
// Solidity: function totalClaimable() view returns(uint256)
func (_Dist *DistCallerSession) TotalClaimable() (*big.Int, error) {
	return _Dist.Contract.TotalClaimable(&_Dist.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Dist *DistTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dist.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Dist *DistSession) Claim() (*types.Transaction, error) {
	return _Dist.Contract.Claim(&_Dist.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Dist *DistTransactorSession) Claim() (*types.Transaction, error) {
	return _Dist.Contract.Claim(&_Dist.TransactOpts)
}

// ClaimAndDelegate is a paid mutator transaction binding the contract method 0x78e2b594.
//
// Solidity: function claimAndDelegate(address delegatee, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dist *DistTransactor) ClaimAndDelegate(opts *bind.TransactOpts, delegatee common.Address, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dist.contract.Transact(opts, "claimAndDelegate", delegatee, expiry, v, r, s)
}

// ClaimAndDelegate is a paid mutator transaction binding the contract method 0x78e2b594.
//
// Solidity: function claimAndDelegate(address delegatee, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dist *DistSession) ClaimAndDelegate(delegatee common.Address, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dist.Contract.ClaimAndDelegate(&_Dist.TransactOpts, delegatee, expiry, v, r, s)
}

// ClaimAndDelegate is a paid mutator transaction binding the contract method 0x78e2b594.
//
// Solidity: function claimAndDelegate(address delegatee, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Dist *DistTransactorSession) ClaimAndDelegate(delegatee common.Address, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Dist.Contract.ClaimAndDelegate(&_Dist.TransactOpts, delegatee, expiry, v, r, s)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dist *DistTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dist.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dist *DistSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dist.Contract.RenounceOwnership(&_Dist.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Dist *DistTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Dist.Contract.RenounceOwnership(&_Dist.TransactOpts)
}

// SetRecipients is a paid mutator transaction binding the contract method 0xae373c1b.
//
// Solidity: function setRecipients(address[] _recipients, uint256[] _claimableAmount) returns()
func (_Dist *DistTransactor) SetRecipients(opts *bind.TransactOpts, _recipients []common.Address, _claimableAmount []*big.Int) (*types.Transaction, error) {
	return _Dist.contract.Transact(opts, "setRecipients", _recipients, _claimableAmount)
}

// SetRecipients is a paid mutator transaction binding the contract method 0xae373c1b.
//
// Solidity: function setRecipients(address[] _recipients, uint256[] _claimableAmount) returns()
func (_Dist *DistSession) SetRecipients(_recipients []common.Address, _claimableAmount []*big.Int) (*types.Transaction, error) {
	return _Dist.Contract.SetRecipients(&_Dist.TransactOpts, _recipients, _claimableAmount)
}

// SetRecipients is a paid mutator transaction binding the contract method 0xae373c1b.
//
// Solidity: function setRecipients(address[] _recipients, uint256[] _claimableAmount) returns()
func (_Dist *DistTransactorSession) SetRecipients(_recipients []common.Address, _claimableAmount []*big.Int) (*types.Transaction, error) {
	return _Dist.Contract.SetRecipients(&_Dist.TransactOpts, _recipients, _claimableAmount)
}

// SetSweepReciever is a paid mutator transaction binding the contract method 0xb438abde.
//
// Solidity: function setSweepReciever(address _sweepReceiver) returns()
func (_Dist *DistTransactor) SetSweepReciever(opts *bind.TransactOpts, _sweepReceiver common.Address) (*types.Transaction, error) {
	return _Dist.contract.Transact(opts, "setSweepReciever", _sweepReceiver)
}

// SetSweepReciever is a paid mutator transaction binding the contract method 0xb438abde.
//
// Solidity: function setSweepReciever(address _sweepReceiver) returns()
func (_Dist *DistSession) SetSweepReciever(_sweepReceiver common.Address) (*types.Transaction, error) {
	return _Dist.Contract.SetSweepReciever(&_Dist.TransactOpts, _sweepReceiver)
}

// SetSweepReciever is a paid mutator transaction binding the contract method 0xb438abde.
//
// Solidity: function setSweepReciever(address _sweepReceiver) returns()
func (_Dist *DistTransactorSession) SetSweepReciever(_sweepReceiver common.Address) (*types.Transaction, error) {
	return _Dist.Contract.SetSweepReciever(&_Dist.TransactOpts, _sweepReceiver)
}

// Sweep is a paid mutator transaction binding the contract method 0x35faa416.
//
// Solidity: function sweep() returns()
func (_Dist *DistTransactor) Sweep(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Dist.contract.Transact(opts, "sweep")
}

// Sweep is a paid mutator transaction binding the contract method 0x35faa416.
//
// Solidity: function sweep() returns()
func (_Dist *DistSession) Sweep() (*types.Transaction, error) {
	return _Dist.Contract.Sweep(&_Dist.TransactOpts)
}

// Sweep is a paid mutator transaction binding the contract method 0x35faa416.
//
// Solidity: function sweep() returns()
func (_Dist *DistTransactorSession) Sweep() (*types.Transaction, error) {
	return _Dist.Contract.Sweep(&_Dist.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dist *DistTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Dist.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dist *DistSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dist.Contract.TransferOwnership(&_Dist.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Dist *DistTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Dist.Contract.TransferOwnership(&_Dist.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Dist *DistTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Dist.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Dist *DistSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Dist.Contract.Withdraw(&_Dist.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Dist *DistTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Dist.Contract.Withdraw(&_Dist.TransactOpts, amount)
}

// DistCanClaimIterator is returned from FilterCanClaim and is used to iterate over the raw logs and unpacked data for CanClaim events raised by the Dist contract.
type DistCanClaimIterator struct {
	Event *DistCanClaim // Event containing the contract specifics and raw log

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
func (it *DistCanClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistCanClaim)
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
		it.Event = new(DistCanClaim)
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
func (it *DistCanClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistCanClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistCanClaim represents a CanClaim event raised by the Dist contract.
type DistCanClaim struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCanClaim is a free log retrieval operation binding the contract event 0x87aeeb9eda09a064caef63d00f62c15063631980bfc422ad7dd30c8a79f0cbb7.
//
// Solidity: event CanClaim(address indexed recipient, uint256 amount)
func (_Dist *DistFilterer) FilterCanClaim(opts *bind.FilterOpts, recipient []common.Address) (*DistCanClaimIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Dist.contract.FilterLogs(opts, "CanClaim", recipientRule)
	if err != nil {
		return nil, err
	}
	return &DistCanClaimIterator{contract: _Dist.contract, event: "CanClaim", logs: logs, sub: sub}, nil
}

// WatchCanClaim is a free log subscription operation binding the contract event 0x87aeeb9eda09a064caef63d00f62c15063631980bfc422ad7dd30c8a79f0cbb7.
//
// Solidity: event CanClaim(address indexed recipient, uint256 amount)
func (_Dist *DistFilterer) WatchCanClaim(opts *bind.WatchOpts, sink chan<- *DistCanClaim, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Dist.contract.WatchLogs(opts, "CanClaim", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistCanClaim)
				if err := _Dist.contract.UnpackLog(event, "CanClaim", log); err != nil {
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

// ParseCanClaim is a log parse operation binding the contract event 0x87aeeb9eda09a064caef63d00f62c15063631980bfc422ad7dd30c8a79f0cbb7.
//
// Solidity: event CanClaim(address indexed recipient, uint256 amount)
func (_Dist *DistFilterer) ParseCanClaim(log types.Log) (*DistCanClaim, error) {
	event := new(DistCanClaim)
	if err := _Dist.contract.UnpackLog(event, "CanClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistHasClaimedIterator is returned from FilterHasClaimed and is used to iterate over the raw logs and unpacked data for HasClaimed events raised by the Dist contract.
type DistHasClaimedIterator struct {
	Event *DistHasClaimed // Event containing the contract specifics and raw log

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
func (it *DistHasClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistHasClaimed)
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
		it.Event = new(DistHasClaimed)
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
func (it *DistHasClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistHasClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistHasClaimed represents a HasClaimed event raised by the Dist contract.
type DistHasClaimed struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHasClaimed is a free log retrieval operation binding the contract event 0x8629b200ebe43db58ad688b85131d53251f3f3be4c14933b4641aeebacf1c08c.
//
// Solidity: event HasClaimed(address indexed recipient, uint256 amount)
func (_Dist *DistFilterer) FilterHasClaimed(opts *bind.FilterOpts, recipient []common.Address) (*DistHasClaimedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Dist.contract.FilterLogs(opts, "HasClaimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &DistHasClaimedIterator{contract: _Dist.contract, event: "HasClaimed", logs: logs, sub: sub}, nil
}

// WatchHasClaimed is a free log subscription operation binding the contract event 0x8629b200ebe43db58ad688b85131d53251f3f3be4c14933b4641aeebacf1c08c.
//
// Solidity: event HasClaimed(address indexed recipient, uint256 amount)
func (_Dist *DistFilterer) WatchHasClaimed(opts *bind.WatchOpts, sink chan<- *DistHasClaimed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Dist.contract.WatchLogs(opts, "HasClaimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistHasClaimed)
				if err := _Dist.contract.UnpackLog(event, "HasClaimed", log); err != nil {
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

// ParseHasClaimed is a log parse operation binding the contract event 0x8629b200ebe43db58ad688b85131d53251f3f3be4c14933b4641aeebacf1c08c.
//
// Solidity: event HasClaimed(address indexed recipient, uint256 amount)
func (_Dist *DistFilterer) ParseHasClaimed(log types.Log) (*DistHasClaimed, error) {
	event := new(DistHasClaimed)
	if err := _Dist.contract.UnpackLog(event, "HasClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Dist contract.
type DistOwnershipTransferredIterator struct {
	Event *DistOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DistOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistOwnershipTransferred)
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
		it.Event = new(DistOwnershipTransferred)
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
func (it *DistOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistOwnershipTransferred represents a OwnershipTransferred event raised by the Dist contract.
type DistOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dist *DistFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DistOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dist.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DistOwnershipTransferredIterator{contract: _Dist.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dist *DistFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DistOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Dist.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistOwnershipTransferred)
				if err := _Dist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Dist *DistFilterer) ParseOwnershipTransferred(log types.Log) (*DistOwnershipTransferred, error) {
	event := new(DistOwnershipTransferred)
	if err := _Dist.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistSweepReceiverSetIterator is returned from FilterSweepReceiverSet and is used to iterate over the raw logs and unpacked data for SweepReceiverSet events raised by the Dist contract.
type DistSweepReceiverSetIterator struct {
	Event *DistSweepReceiverSet // Event containing the contract specifics and raw log

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
func (it *DistSweepReceiverSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistSweepReceiverSet)
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
		it.Event = new(DistSweepReceiverSet)
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
func (it *DistSweepReceiverSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistSweepReceiverSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistSweepReceiverSet represents a SweepReceiverSet event raised by the Dist contract.
type DistSweepReceiverSet struct {
	NewSweepReceiver common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSweepReceiverSet is a free log retrieval operation binding the contract event 0xbea8251f76064f657f2a744bf08a1b5700486d06eb94922162892eb022d95ef6.
//
// Solidity: event SweepReceiverSet(address indexed newSweepReceiver)
func (_Dist *DistFilterer) FilterSweepReceiverSet(opts *bind.FilterOpts, newSweepReceiver []common.Address) (*DistSweepReceiverSetIterator, error) {

	var newSweepReceiverRule []interface{}
	for _, newSweepReceiverItem := range newSweepReceiver {
		newSweepReceiverRule = append(newSweepReceiverRule, newSweepReceiverItem)
	}

	logs, sub, err := _Dist.contract.FilterLogs(opts, "SweepReceiverSet", newSweepReceiverRule)
	if err != nil {
		return nil, err
	}
	return &DistSweepReceiverSetIterator{contract: _Dist.contract, event: "SweepReceiverSet", logs: logs, sub: sub}, nil
}

// WatchSweepReceiverSet is a free log subscription operation binding the contract event 0xbea8251f76064f657f2a744bf08a1b5700486d06eb94922162892eb022d95ef6.
//
// Solidity: event SweepReceiverSet(address indexed newSweepReceiver)
func (_Dist *DistFilterer) WatchSweepReceiverSet(opts *bind.WatchOpts, sink chan<- *DistSweepReceiverSet, newSweepReceiver []common.Address) (event.Subscription, error) {

	var newSweepReceiverRule []interface{}
	for _, newSweepReceiverItem := range newSweepReceiver {
		newSweepReceiverRule = append(newSweepReceiverRule, newSweepReceiverItem)
	}

	logs, sub, err := _Dist.contract.WatchLogs(opts, "SweepReceiverSet", newSweepReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistSweepReceiverSet)
				if err := _Dist.contract.UnpackLog(event, "SweepReceiverSet", log); err != nil {
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

// ParseSweepReceiverSet is a log parse operation binding the contract event 0xbea8251f76064f657f2a744bf08a1b5700486d06eb94922162892eb022d95ef6.
//
// Solidity: event SweepReceiverSet(address indexed newSweepReceiver)
func (_Dist *DistFilterer) ParseSweepReceiverSet(log types.Log) (*DistSweepReceiverSet, error) {
	event := new(DistSweepReceiverSet)
	if err := _Dist.contract.UnpackLog(event, "SweepReceiverSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistSweptIterator is returned from FilterSwept and is used to iterate over the raw logs and unpacked data for Swept events raised by the Dist contract.
type DistSweptIterator struct {
	Event *DistSwept // Event containing the contract specifics and raw log

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
func (it *DistSweptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistSwept)
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
		it.Event = new(DistSwept)
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
func (it *DistSweptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistSweptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistSwept represents a Swept event raised by the Dist contract.
type DistSwept struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSwept is a free log retrieval operation binding the contract event 0x7f221332ee403570bf4d61630b58189ea566ff1635269001e9df6a890f413dd8.
//
// Solidity: event Swept(uint256 amount)
func (_Dist *DistFilterer) FilterSwept(opts *bind.FilterOpts) (*DistSweptIterator, error) {

	logs, sub, err := _Dist.contract.FilterLogs(opts, "Swept")
	if err != nil {
		return nil, err
	}
	return &DistSweptIterator{contract: _Dist.contract, event: "Swept", logs: logs, sub: sub}, nil
}

// WatchSwept is a free log subscription operation binding the contract event 0x7f221332ee403570bf4d61630b58189ea566ff1635269001e9df6a890f413dd8.
//
// Solidity: event Swept(uint256 amount)
func (_Dist *DistFilterer) WatchSwept(opts *bind.WatchOpts, sink chan<- *DistSwept) (event.Subscription, error) {

	logs, sub, err := _Dist.contract.WatchLogs(opts, "Swept")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistSwept)
				if err := _Dist.contract.UnpackLog(event, "Swept", log); err != nil {
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

// ParseSwept is a log parse operation binding the contract event 0x7f221332ee403570bf4d61630b58189ea566ff1635269001e9df6a890f413dd8.
//
// Solidity: event Swept(uint256 amount)
func (_Dist *DistFilterer) ParseSwept(log types.Log) (*DistSwept, error) {
	event := new(DistSwept)
	if err := _Dist.contract.UnpackLog(event, "Swept", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DistWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the Dist contract.
type DistWithdrawalIterator struct {
	Event *DistWithdrawal // Event containing the contract specifics and raw log

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
func (it *DistWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DistWithdrawal)
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
		it.Event = new(DistWithdrawal)
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
func (it *DistWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DistWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DistWithdrawal represents a Withdrawal event raised by the Dist contract.
type DistWithdrawal struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed recipient, uint256 amount)
func (_Dist *DistFilterer) FilterWithdrawal(opts *bind.FilterOpts, recipient []common.Address) (*DistWithdrawalIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Dist.contract.FilterLogs(opts, "Withdrawal", recipientRule)
	if err != nil {
		return nil, err
	}
	return &DistWithdrawalIterator{contract: _Dist.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed recipient, uint256 amount)
func (_Dist *DistFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *DistWithdrawal, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Dist.contract.WatchLogs(opts, "Withdrawal", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DistWithdrawal)
				if err := _Dist.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed recipient, uint256 amount)
func (_Dist *DistFilterer) ParseWithdrawal(log types.Log) (*DistWithdrawal, error) {
	event := new(DistWithdrawal)
	if err := _Dist.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
