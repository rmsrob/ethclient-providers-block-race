package ethclientProvidersBlockRace

import (
	"context"
	"fmt"
	"regexp"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Runner struct {
	url 	string
	laps 	time.Duration
}

func New(wsProviders []string, blockCount int) (best string, err error) {
	var (
		runnersStats []Runner
		ln = len(wsProviders)
	)

	switch {
	case blockCount < 1:
		err = fmt.Errorf("to start the race we need at least to run for 1 block")
		return "", err
	case ln <= 1:
		err = fmt.Errorf("to start the race we need at least 2 runners")
		return "", err
	case ln >= 2:
		for _, url := range wsProviders {
			re := regexp.MustCompile(`^(ws|wss):\/\/.*`)
			if !re.MatchString(url) {
				err = fmt.Errorf("url must be a websocket standard")
				return "", err
			}
		}
		fallthrough
	case ln >= 1 && blockCount >= 1:
		fmt.Printf("*** The race can beggin with %d runners over %d blocks ***\n", ln, blockCount)
	}

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
				// Disqualification rule if the provider stale for 48sec is out of the race
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
			from := runnersStats[0].laps - r.laps
			fmt.Printf("#%d > %s with %f seconds >> %f\n", i+1, r.url[:31]+"...", r.laps.Seconds(), from.Seconds())
		} else {
			fmt.Printf("#%d > %s with %f seconds >> 0\n", i+1, r.url[:31]+"...", r.laps.Seconds())
		}
	}

	return runnersStats[0].url, err
}
