package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/event"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
	"github.com/koinotice/oneplus/meshdb"
	"github.com/koinotice/oneplus/watch"
	"github.com/koinotice/oneplus/watch/blockwatch"
	"github.com/koinotice/oneplus/watch/plugin"
	"github.com/koinotice/oneplus/watch/structs"
	 

	"github.com/sirupsen/logrus"
	"path/filepath"
	"time"
)

func block() {
	api := "https://mainnet.infura.io/v3/19d753b2600445e292d54b1ef58d4df4"
	w := watch.NewHttpBasedEthWatcher(context.Background(), api)

	w.RegisterBlockPlugin(plugin.NewBlockNumPlugin(func(i uint64, b bool) {
		fmt.Println(">>", i, b)
	}))

	w.RunTillExit()
}
func receipt() {
	logrus.SetLevel(logrus.DebugLevel)

	ctx := context.Background()
	api := "https://kovan.infura.io/v3/19d753b2600445e292d54b1ef58d4df4"
	startBlock := 12220000
	contract := "0xAc34923B2b8De9d441570376e6c811D5aA5ed72f"
	interestedTopics := []string{
		"0x23b872dd7302113369cda2901243429419bec145408fa8b352b3dd92b66c680b",
		"0x6bf96fcc2cec9e08b082506ebbc10114578a497ff1ea436628ba8996b750677c",
		"0x5a746ce5ce37fc996a6e682f4f84b6f90d3be79fd8ac9a8a11264345f3d29edd",
		"0x9c4e90320be51bb93d854d0ab9ba8aa249dabc21192529efcd76ae7c22c6bc0b",
		"0x0ce31a5f70780bb6770b52a6793403d856441ccb475715e8382a0525d35b0558",
	}

	handler := func(log structs.RemovableReceiptLog) {
		logrus.Infof("log from tx: %s", log.GetTransactionHash())
	}

	stepsForBigLag := 100

	highestProcessed := watch.ListenForReceiptLogTillExit(ctx, api, startBlock, contract, interestedTopics, handler, stepsForBigLag)
	logrus.Infof("highestProcessed: %d", highestProcessed)
}

const (
	blockWatcherRetentionLimit = 20
	ethereumRPCRequestTimeout  = 30 * time.Second
	ethWatcherPollingInterval  = 1 * time.Minute
	peerConnectTimeout         = 60 * time.Second
	checkNewAddrInterval       = 20 * time.Second
	expirationPollingInterval  = 50 * time.Millisecond
	BlockPollingInterval       = 5 * time.Second
	EthereumRPCURL             = "https://rinkeby.infura.io/v3/cabc724fb9534d1bb245582a74ccf3e7"
)

type Watch struct{
	blockWatcher  *blockwatch.Watcher
	blockSubscription          event.Subscription

}

func BlockWatch() (*Watch, error) {

	databasePath := filepath.Join("0vedb", "db")
	meshDB, err := meshdb.NewMeshDB(databasePath)
	if err != nil {
		return nil, err
	}

	// Initialize gateway watcher (but don't start it yet).
	blockWatcherClient, err := blockwatch.NewRpcClient(EthereumRPCURL, ethereumRPCRequestTimeout)
	if err != nil {
		return nil,err
	}
	//topics := orderwatch.GetRelevantTopics()
	blockWatcherConfig := blockwatch.Config{
		MeshDB:              meshDB,
		PollingInterval:     BlockPollingInterval,
		StartBlockDepth:     ethrpc.LatestBlockNumber,
		BlockRetentionLimit: blockWatcherRetentionLimit,
		WithLogs:            true,
		//Topics:              topics,
		Client: blockWatcherClient,
	}
	blockWatcher := blockwatch.New(blockWatcherConfig)
	go func() {
		for {
			err, isOpen := <-blockWatcher.Errors
			if isOpen {
				logrus.WithField("error", err).Error("BlockWatcher error encountered")
			} else {
				return // Exit when the error channel is closed
			}
		}
	}()


	 watch:=&Watch{
		 blockWatcher:blockWatcher,
	 }
	return watch,nil
}



func (w *Watch)  setupEventWatcher() {
	blockEvents := make(chan []*blockwatch.Event, 1)
	w.blockSubscription = w.blockWatcher.Subscribe(blockEvents)

	go func() {
		for {
			select {
			case err, isOpen := <-w.blockSubscription.Err():
				close(blockEvents)
				if !isOpen {
					// event.Subscription closes the Error channel on unsubscribe.
					// We therefore cleanup this goroutine on channel closure.
					return
				}
				logrus.WithFields(logrus.Fields{
					"error": err.Error(),
				}).Error("subscription error encountered")
				return

			case events := <-blockEvents:
				//hashToOrderWithTxHashes := map[common.Hash]*OrderWithTxHashes{}
				for _, event := range events {
					fmt.Printf("current block number,%s \n", event.BlockHeader.Number)
				}

				//fmt.Printf("orderr hash %s\n", hashToOrderWithTxHashes)
				//w.generateOrderEventsIfChanged(hashToOrderWithTxHashes)
			}
		}
	}()
}

func main() {
	watch,err1:=BlockWatch()

	if err1!=nil{
		fmt.Printf(err1.Error())
	}
	if err :=  watch.blockWatcher.Start(); err != nil {
		logrus.WithField("error", err).Error("BlockWatcher start")


	}
	select {}
}
