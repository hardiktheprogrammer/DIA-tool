package controller

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Controller() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/67e0b3b448f84921b3049e5336279397")
	if err != nil {
		log.Fatal(err)
	}

	// Define the Oracle Address and Chain ID
	oracleAddress := common.HexToAddress("0x123456789abcdef123456789abcdef12345678")
	chainID := int64(1) // Mainnet

	// Retrieve the latest block number
	blockNumber, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Loop through the most recent blocks and retrieve transactions
	for i := blockNumber; i > blockNumber-10; i-- {
		block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		if err != nil {
			log.Fatal(err)
		}

		for _, tx := range block.Transactions() {
			if tx.To() == &oracleAddress && tx.ChainId().Cmp(big.NewInt(chainID)) == 0 {
				// Process the transaction as needed
				fmt.Printf("Transaction hash: %s\n", tx.Hash().Hex())
			}
		}
	}
}
