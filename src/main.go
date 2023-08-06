package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

// Block represents a block in the BlockDAG.
type Block struct {
	Index        int64
	Timestamp    int64
	Transactions []string
	PrevHash     string
	Hash         string
}

// NewBlock creates a new block in the BlockDAG.
func NewBlock(index int64, transactions []string, prevHash string) *Block {
	block := &Block{
		Index:        index,
		Timestamp:    time.Now().Unix(),
		Transactions: transactions,
		PrevHash:     prevHash,
	}
	block.Hash = block.calculateHash()
	return block
}

// calculateHash calculates the hash of the block.
func (b *Block) calculateHash() string {
	record := fmt.Sprintf("%d%d%s%s", b.Index, b.Timestamp, b.Transactions, b.PrevHash)
	hash := sha256.Sum256([]byte(record))
	return hex.EncodeToString(hash[:])
}

// BlockDAG represents the blockchain using a directed acyclic graph (DAG) of blocks.
type BlockDAG struct {
	blocks []*Block
}

// NewBlockDAG creates a new BlockDAG instance.
func NewBlockDAG() *BlockDAG {
	genesisBlock := NewBlock(0, []string{"Genesis Transaction"}, "")
	return &BlockDAG{blocks: []*Block{genesisBlock}}
}

// AddBlock adds a new block to the BlockDAG.
func (bd *BlockDAG) AddBlock(transactions []string) {
	prevBlock := bd.blocks[len(bd.blocks)-1]
	newBlock := NewBlock(prevBlock.Index+1, transactions, prevBlock.Hash)
	bd.blocks = append(bd.blocks, newBlock)
}

// ValidateChain validates the integrity of the BlockDAG.
func (bd *BlockDAG) ValidateChain() bool {
	for i := 1; i < len(bd.blocks); i++ {
		currentBlock := bd.blocks[i]
		prevBlock := bd.blocks[i-1]

		if currentBlock.Hash != currentBlock.calculateHash() {
			return false
		}

		if currentBlock.PrevHash != prevBlock.Hash {
			return false
		}
	}

	return true
}

func main() {
	blockDAG := NewBlockDAG()

	blockDAG.AddBlock([]string{"Transaction 1"})
	blockDAG.AddBlock([]string{"Transaction 2"})
	blockDAG.AddBlock([]string{"Transaction 3"})

	fmt.Println("BlockDAG created and validated:", blockDAG.ValidateChain())
}
