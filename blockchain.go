package main

import (
	"fmt"
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

func (bc *Blockchain) CreateBlock(nonce int, previousHash string) *Block {
	b := NewBlock(nonce, previousHash)
	bc.chain = append(bc.chain, b)
	return b
}

func main() {
	b := NewBlock(0, "init hash")
	fmt.Println(b)
	b.Print()

}
