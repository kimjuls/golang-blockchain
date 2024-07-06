package main

import "fmt"

func main() {
	blockchain := NewBlockchain()
	blockchain.CreateBlock("0.4 BTC Transaction from A to B")
	blockchain.CreateBlock("0.52 BTC Transaction from B to C")
	for _, block := range blockchain.blocks {
		fmt.Printf("Previous hash: %x\n", block.PreviousHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}
}
