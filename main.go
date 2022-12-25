package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Runner struct {
	url 	string
	laps 	time.Duration
}

func main() {
	var runnersStats []Runner

	// Change here the full websocket providers url you want to enter into the race
	wsProviders := [3]string{
		"wss://eth-mainnet.g.alchemy.com/v2/YOUR-ALCHEMY-KEY",
		"wss://mainnet.infura.io/ws/v3/YOUR-INFURA-KEY",
		"wss://rpc.ankr.com/eth",
	}
	// Enter the numbers of blocks the race should last
	blockCount 	:= 3

	fmt.Printf("\n---WSS PROVIDERS RUN A SHORT %v BLOCKS RACE---\n", blockCount)
	var wg sync.WaitGroup
	results := make([]int, len(wsProviders))

	for i, url := range wsProviders {
		wg.Add(1)
		go func(i int, url string) {
			defer wg.Done()
			count := 0
			start := time.Now()

			client, err := ethclient.Dial(url)
			if err != nil {
				fmt.Printf("Error connecting to Ethereum client with url: %s > %v\n", url, err)
				return
			}
			defer client.Close()

			blocksChan := make(chan *types.Header)
			sub, err := client.SubscribeNewHead(context.Background(), blocksChan)
			if err != nil {
				fmt.Printf("Error subscribing to blocks with url: %s > %v\n", url, err)
				return
			}
			defer sub.Unsubscribe()

			for {
				select {
				case header := <-blocksChan:
					hash := header.Hash().Hex()
					shorterHash := hash[:4] + "..." + hash[len(hash)-5:]
					fmt.Printf("Received block with hash %s from url %s\n", shorterHash, url[:31]+"...")
					count++
					if count >= blockCount {
						elapsed := time.Since(start)
						results[i] = count
						runnersStats = append(runnersStats, Runner{
							url: url,
							laps: elapsed,
						})
						return
					}
				case <-time.After(time.Second*48):
					fmt.Printf("Timed out for url %s\n", url)
					return
				}
			}
		}(i, url)
	}
	wg.Wait()

	fmt.Printf("\n===WSS PROVIDERS PODIUM FOR RUNNING %d BLOCKS===\n", blockCount)
	for i, r := range runnersStats {
		if i != 0 {
			from := runnersStats[i-1].laps - r.laps
			fmt.Printf("#%d > %s with %f seconds >> %f\n", i+1, r.url[:31]+"...", r.laps.Seconds(), from.Seconds())
		} else {
			fmt.Printf("#%d > %s with %f seconds >> 0\n", i+1, r.url[:31]+"...", r.laps.Seconds())
		}
	}
}
