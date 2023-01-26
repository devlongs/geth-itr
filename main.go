package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func currentBlock() {
    block, err := client.BlockByNumber(ctx, nil)
    if err != nil {
        log.Println(err)
    }
    fmt.Println("Current block number is:", block.Number())
}

func main(){
	errr := godotenv.Load()
	if errr != nil {
		log.Fatal("Error loading .env file")
	}

	URL := os.Getenv("URL")

	var (
		ctx         = context.Background()
		url         = URL
		client, err = ethclient.DialContext(ctx, url)
	)
	// currentBlock()
	address := common.HexToAddress("0x8335659d19e46e720e7894294630436501407c3e")
	balance, err := client.BalanceAt(ctx, address, nil)
    if err != nil {
        log.Print("There was an error", err)
    }
    fmt.Println("The balance at the current block number is", balance)
	currentBlock()

}