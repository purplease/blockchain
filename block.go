package main

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Data + block.PrevHash
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

func generateBlock(prevBlock Block, data string) Block {
	var newBlock Block
	t := time.Now()

	newBlock.Index = prevBlock.Index + 1
	newBlock.Timestamp = t.String()
	newBlock.Data = data
	newBlock.PrevHash = prevBlock.Hash
	newBlock.Hash = calculateHash(newBlock)

	return newBlock
}
