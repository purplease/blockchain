package main

import "time"

const Difficulty = 2

type Blockchain struct {
	Chain []Block
}

func createGenesisBlock() Block {
	var genesisBlock Block
	genesisBlock.Index = 0
	genesisBlock.Timestamp = time.Now().String()
	genesisBlock.Data = "Genesis Block"
	genesisBlock.PrevHash = ""
	genesisBlock.Hash = calculateHash(genesisBlock)

	return genesisBlock
}

func createBlockchain() Blockchain {
	genesisBlock := createGenesisBlock()
	blockchain := Blockchain{}
	blockchain.Chain = append(blockchain.Chain, genesisBlock)

	return blockchain
}

func (bc *Blockchain) addBlock(data string) {
	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := generateBlock(prevBlock, data)
	bc.Chain = append(bc.Chain, newBlock)
}

func (bc *Blockchain) isValid() bool {
	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		prevBlock := bc.Chain[i-1]

		if currentBlock.Hash != calculateHash(currentBlock) {
			return false
		}

		if currentBlock.PrevHash != prevBlock.Hash {
			return false
		}
	}

	return true
}
