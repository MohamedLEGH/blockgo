package main

import (
	"blockgo/blockchain"
	"fmt"
)

type Transaction = blockchain.Transaction
type Block = blockchain.Block

func main() {
	// tx := Transaction{Timestamp: 0, Sender: "me", Receiver: "him", Value: 10}
	block := Block{Index: 0, Previous_hash: "", Timestamp: 0, Miner_address: "me", Tx_list: []Transaction{}}
	fmt.Printf("%+v\n", block)
	block.Mine(10000)
	hash0 := block.Hash_val
	// hash0 := block.Hash(1)
	fmt.Printf("%+v\n", hash0)
	tx := Transaction{Sender: "me", Receiver: "him", Value: 10}
	// fmt.Printf("%+v\n", tx)
	block.AddTransaction(tx)
	fmt.Printf("%+v\n", block)
	// fmt.Printf("%+v\n", block.Hash(1))
}
