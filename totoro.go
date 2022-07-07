package totoro

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

const (
	defaultPollingInterval = 5 * time.Second
	blockNumStepBack       = 5
)

type Subscriber struct {
	ethCli          *ethclient.Client
	q               ethereum.FilterQuery
	logs            map[common.Hash]uint64
	pollingInternal time.Duration
	logCh           chan types.Log
	closeCh         chan struct{}
	errCh           chan error
	ctx             context.Context
	firstQuery      bool
}

func NewSubscriber(
	ctx context.Context,
	ethCli *ethclient.Client,
	q ethereum.FilterQuery,
	ch chan types.Log,
) *Subscriber {
	scrb := &Subscriber{
		ctx:             ctx,
		ethCli:          ethCli,
		q:               q,
		logs:            make(map[common.Hash]uint64),
		pollingInternal: defaultPollingInterval,
		logCh:           ch,
		firstQuery:      true,
	}
	scrb.subscribe()
	return scrb
}

func (scrb *Subscriber) subscribe() {
	scrb.errCh = make(chan error, 10)
	scrb.closeCh = make(chan struct{})
	go scrb.polling()
}

func (scrb *Subscriber) polling() {
	ticker := time.NewTicker(scrb.pollingInternal)
	for {
		select {
		case <-scrb.closeCh:
			ticker.Stop()
			return
		case <-ticker.C:
			q := scrb.q
			if !scrb.firstQuery {
				lastBlockTo := q.ToBlock
				if lastBlockTo == nil {
					blockNumNow, err := scrb.ethCli.BlockNumber(scrb.ctx)
					if err != nil {
						scrb.errCh <- err
						continue
					}
					lastBlockTo = big.NewInt(int64(blockNumNow))
				}
				q.FromBlock = lastBlockTo.Sub(lastBlockTo, big.NewInt(blockNumStepBack))
			}
			lgs, err := scrb.ethCli.FilterLogs(scrb.ctx, q)
			if err != nil {
				scrb.errCh <- err
				continue
			}
			scrb.firstQuery = false
			for _, lg := range lgs {
				lg := lg
				if _, ok := scrb.logs[lg.TxHash]; !ok {
					scrb.logs[lg.TxHash] = lg.BlockNumber
					scrb.logCh <- lg
				}
			}
			scrb.clearLogs(q.FromBlock.Uint64())
		}
	}
}

func (scrb *Subscriber) clearLogs(lastBlock uint64) {
	for k, blockNum := range scrb.logs {
		if blockNum < lastBlock-blockNumStepBack {
			delete(scrb.logs, k)
		}
	}
}

func (scrb *Subscriber) SetPollingInternal(duration time.Duration) {
	scrb.pollingInternal = duration
	scrb.Unsubscribe()
	scrb.subscribe()
}

func (scrb *Subscriber) Unsubscribe() {
	close(scrb.closeCh)
	close(scrb.errCh)
}

func (scrb *Subscriber) Err() <-chan error {
	return scrb.errCh
}
