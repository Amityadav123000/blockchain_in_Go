package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transaction  []*Trnsaction
}

// function returns new block i.e intializing new block
func NewBlock(nonce int, previousHash [32]byte, transaction []*Trnsaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transaction = transaction
	return b
}

// function to print block elements
func (b *Block) Print() {
	fmt.Printf("timestamp      %d\n", b.timestamp)
	fmt.Printf("nonce          %d\n", b.nonce)
	fmt.Printf("previous_hash  %s\n", b.previousHash)
	for _, t := range b.transaction {
		t.Print()
	}
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	// fmt.Println(string(m))
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64         `json:"timestmp"`
		Nonce        int           `json:"nonce"`
		PreviousHash [32]byte      `json:"previous_hash"`
		Trnsaction   []*Trnsaction `json:"transaction"`
	}{
		Timestamp:    b.timestamp,
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Trnsaction:   b.transaction,
	})
}

type Blockchain struct {
	transactionPool []*Trnsaction
	chain           []*Block
}

func NewBlockChain() *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) CreateBlock(nonce int, previosusHash [32]byte) *Block {
	b := NewBlock(nonce, previosusHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Trnsaction{}
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc *Blockchain) AddTransaction(sender string, reciepeint string, value float32) {
	t := NewTransaction(sender, reciepeint, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

type Trnsaction struct {
	senderBLockChainAddress   string
	recieverBlockChainAddress string
	value                     float32
}

func NewTransaction(sender string, recipient string, value float32) *Trnsaction {
	return &Trnsaction{sender, recipient, value}
}

func (t *Trnsaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf("senderBlockChainAddress    %s\n", t.senderBLockChainAddress)
	fmt.Printf("recieverBlockChainAddress  %s\n", t.recieverBlockChainAddress)
	fmt.Printf("value                      %.1f\n", t.value)
}

func (t *Trnsaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		SenderBLockChainAddress   string  `json:"sender_blockchain_address"`
		RecieverBlockChainAddress string  `json:"reciepient_blockchain_address"`
		Value                     float32 `json:"value"`
	}{
		SenderBLockChainAddress:   t.senderBLockChainAddress,
		RecieverBlockChainAddress: t.recieverBlockChainAddress,
		Value:                     t.value,
	})
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

	blockchain.AddTransaction("A", "B", 1.0)
	previousHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(5, previousHash)
	// blockchain.Print()

	blockchain.AddTransaction("C", "D", 2.0)
	blockchain.AddTransaction("X", "Y", 3.0)
	previousHash = blockchain.LastBlock().Hash()
	blockchain.CreateBlock(2, previousHash)
	blockchain.Print()
}
