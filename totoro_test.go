package totoro

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"os"
	"testing"
	"time"
)

var ethCli *ethclient.Client
var query ethereum.FilterQuery
var subscriber *Subscriber
var ch chan types.Log
var ctx context.Context
var privateKey string
var testContract *EventTest
var transactOpt *bind.TransactOpts

const (
	testContractAddr  = "0x4668e06aba7AAd5379Aa54c2cE9924B8A465ffC1"
	bscTestNetRPC     = "https://data-seed-prebsc-1-s1.binance.org:8545"
	bscTestNetChainID = 97
)

func init() {
	var err error
	privateKey = os.Getenv("PRIVATE_KEY_TOTORO")
	ctx = context.Background()
	query := ethereum.FilterQuery{
        FromBlock: big.NewInt(20837786),
        ToBlock:   nil,
        Addresses: []common.Address{common.HexToAddress(testContractAddr)},
    }
	ethCli, err = ethclient.Dial(bscTestNetRPC)
	if err != nil {
		panic(err)
	}
	ch = make(chan types.Log)
	subscriber = NewSubscriber(ctx, ethCli, query, ch)
	testContract, err = NewEventTest(common.HexToAddress(testContractAddr), ethCli)
	if err != nil {
		panic(err)
	}
	priKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err)
	}
	transactOpt, err = bind.NewKeyedTransactorWithChainID(priKey, big.NewInt(97))
	if err != nil {
		panic(err)
	}
}

func triggerEvtLoop() {
	go func() {
		for {
			if txn, err := testContract.TriggerEvt(transactOpt); err != nil {
				panic(err)
			} else {
				fmt.Printf("trigger evt success:%v\n", txn.Hash())
			}
			time.Sleep(1 * time.Second)
		}
	}()
}

func TestTotoroSubscribe(t *testing.T) {
	t.Run("totoro subscriber", func(t *testing.T) {
		triggerEvtLoop()
		for {
			select {
			case log := <-ch:
				t.Logf("log: %+v", log)
			case err := <-subscriber.Err():
				t.Logf("subscriber error: %v", err)
			}
		}
	})
}
