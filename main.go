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
	chain := Chain{Difficulty: 1}
	chain.CreateGenesisBlock(miner_address)
	fmt.Printf("%+v\n", chain)
	tx := Transaction{Sender: "me", Receiver: "him", Value: 10}
	chain.AddTransaction(tx)
	fmt.Printf("%+v\n", chain)
	chain.MineNewBlock(miner_address)
	fmt.Printf("%+v\n", chain)
}
