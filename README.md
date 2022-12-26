[![Go Reference](https://pkg.go.dev/badge/github.com/rrobrms/ethclient-providers-block-race.svg)](https://pkg.go.dev/github.com/rrobrms/ethclient-providers-block-race)
[![Go Report Card](https://goreportcard.com/badge/github.com/rrobrms/ethclient-providers-block-race)](https://goreportcard.com/report/github.com/rrobrms/ethclient-providers-block-race)
[![Coverage Status](https://coveralls.io/repos/github/rrobrms/ethclient-providers-block-race/badge.svg?branch=master)](https://coveralls.io/github/rrobrms/ethclient-providers-block-race?branch=master)

# Ethclient Providers Block Race

> Test your Ethclient Providers with a race to discover x blocks

![race](https://user-images.githubusercontent.com/93430216/209518983-7b9e1efd-b623-4b5c-b661-9372c272587a.gif)


## Prerequisite
Required:
- [Go](https://go.dev/doc/install)

## Usage

```sh
go get github.com/rrobrms/ethclient-providers-block-race.go
```

```go
func MyFunc() {
	var (
		blockCount = 23
		wsProviders = []string{
			"wss://eth-mainnet.g.alchemy.com/v2/ALCHEMY_API_KEY",
			"wss://mainnet.infura.io/ws/v3/INFURA_API_KEY",
			"wss://rpc.ankr.com/eth",
		}
	)

	best, err := ethClientRace(wsProviders, blockCount)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("\nyou should use this provider: %s\n", best)
}
```