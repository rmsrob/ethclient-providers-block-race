package main

import (
	"ethclientProvidersBlockRace"
	"fmt"
	"log"
)


func main() {
	var (
		blockCount = 23
		wsProviders = []string{
			"wss://eth-mainnet.g.alchemy.com/v2/ALCHEMY_API_KEY",
			"wss://mainnet.infura.io/ws/v3/INFURA_API_KEY",
			"wss://rpc.ankr.com/eth",
		}
	)

	best, err := ethclientProvidersBlockRace.New(wsProviders, blockCount)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\nyou should use this provider: %s\n", best[:36]+"***")
}