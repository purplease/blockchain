package main

import (
	"fmt"
)

func main() {
	bc := createBlockchain()

	bc.addBlock("Block 1 Data")
	bc.addBlock("Block 2 Data")

	fmt.Println("Blockchain Valid:", bc.isValid())
	fmt.Println("Blockchain:")
	for _, block := range bc.Chain {
		fmt.Printf("Index: %d\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Println("-----------------------------------")
	}
}
