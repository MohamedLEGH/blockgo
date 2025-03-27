package main

import (
	"blockgo/blockchain"
	"fmt"
)

type Transaction = blockchain.Transaction
type Block = blockchain.Block
type Chain = blockchain.Chain

func main() {
	miner_address := "me"
	chain := Chain{Difficulty : 1}
	chain.CreateGenesisBlock(miner_address)
	// // tx := Transaction{Timestamp: 0, Sender: "me", Receiver: "him", Value: 10}
	// block := Block{Index: 0, Previous_hash: "", Timestamp: 0, Miner_address: "me", Tx_list: []Transaction{}}
	// fmt.Printf("%+v\n", block)
	// block.Mine(10)
	// hash0 := block.Hash_val
	// // hash0 := block.Hash(1)
	fmt.Printf("%+v\n", chain)
	tx := Transaction{Sender: "me", Receiver: "him", Value: 10}
	// fmt.Printf("%+v\n", tx)
	chain.AddTransaction(tx)
	fmt.Printf("%+v\n", chain)
	chain.MineNewBlock(miner_address)
	fmt.Printf("%+v\n", chain)
	// fmt.Printf("%+v\n", block.Hash(1))
}
