package miner

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ProxyABIString is the input ABI used to generate the binding from.
const ProxyABIString = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receivingAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"msgSender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FlashbotsPayment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receivingAddress\",\"type\":\"address\"}],\"name\":\"RecipientUpdate\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"getMinerReceivingAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"payMiner\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queueEther\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newReceivingAddress\",\"type\":\"address\"}],\"name\":\"setMinerReceivingAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ProxyFlashbotsPayment represents a FlashbotsPayment event raised by the Proxy contract.
type ProxyFlashbotsPayment struct {
	Coinbase         common.Address
	ReceivingAddress common.Address
	MsgSender        common.Address
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

var (
	proxyPaymentTopic = common.HexToHash("0x95b37ac100e4bf5cd43ff3a48e440813330509e42025bdc7e778b7fa4e2a0c18")
	proxyABI, _       = abi.JSON(strings.NewReader(ProxyABIString))
)
