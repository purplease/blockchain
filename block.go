package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"strings"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int // We need to use the Nonce value when calculating the hash of the block to include it in the proof-of-work mechanism
}

func calculateHash(block Block) string {
	//By including the Nonce value in the record string, the hash calculation will take into account the incrementing Nonce value and produce a different hash for each iteration
	record := string(block.Index) + block.Timestamp + block.Data + block.PrevHash + strconv.Itoa(block.Nonce)
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
	newBlock.Nonce = 0

	for !isBlockValid(newBlock) {
		newBlock.Nonce++
		newBlock.Hash = calculateHash(newBlock)
	}

	return newBlock
}

func isBlockValid(block Block) bool {
	prefix := strings.Repeat("0", Difficulty)
	return strings.HasPrefix(block.Hash, prefix)
}
