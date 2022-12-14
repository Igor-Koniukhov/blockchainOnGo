package main

import (
	"fmt"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}
func (b *Block) Print() {
	fmt.Printf("timestamp %d\n", b.timestamp)
	fmt.Printf("nonce %d\n", b.nonce)
	fmt.Printf("previous_hash %s\n", b.previousHash)
	fmt.Printf("nonce %d\n", b.nonce)
	fmt.Printf("nonce %s\n", b.transactions)

}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init hash")
	return bc
}
func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s \n", strings.Repeat("==", 25), i, strings.Repeat("=", 25))
		block.Print()
	}

}

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func main() {
	blockChain := NewBlockchain()
	blockChain.Print()
	blockChain.CreateBlock(3, "hash1")
	blockChain.Print()
	blockChain.CreateBlock(4, "hash2")
	blockChain.Print()
}
