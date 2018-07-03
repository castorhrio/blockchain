package main

import (
	"fmt"

	"./chain"
)

func main() {
	bc := chain.NewChain()
	bc.AddBlock("Send 1 BTC to henry")
	bc.AddBlock("Send 2 more BTC to henry")

	for _, block := range bc.Chains {
		fmt.Printf("Prev hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
