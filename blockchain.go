package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash string
	timestamp    int64
	transaction  []string
}

// function returns new block
func NewBlock(nonce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	return b
}

// function to print block elements
func (b *Block) Print() {
	fmt.Printf("timestamp      %d\n", b.timestamp)
	fmt.Printf("nonce          %d\n", b.nonce)
	fmt.Printf("previous_hash  %s\n", b.previousHash)
	fmt.Printf("transaction    %s\n", b.transaction)
}

type Blockchain struct {
	transactionPool []string
	chain           []*Block
}

func NewBlockChain() *Blockchain {
	bc := new(Blockchain)
	bc.CreateBlock(0, "Init Hash")
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previosusHash string) *Block {
	b := NewBlock(nonce, previosusHash)
	bc.chain = append(bc.chain, b)
	return b
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s \n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 60))
}

func init() {
	log.SetPrefix("Blockchain:")
}

func main() {
	blockchain := NewBlockChain()
	// blockchain.Print()
	blockchain.CreateBlock(5, "Hash 1")
	// blockchain.Print()
	blockchain.CreateBlock(2, "Hash 2")
	blockchain.Print()
}
