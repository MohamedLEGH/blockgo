package main

import (
	"blockgo/blockchain"
	// "fmt"
	"github.com/davecgh/go-spew/spew"
)

type Transaction = blockchain.Transaction
type Block = blockchain.Block
type Chain = blockchain.Chain

func main() {
	privKey, address := blockchain.GenerateAccount()
	chain := Chain{Difficulty: 1}
	chain.CreateGenesisBlock(privKey)
	spew.Dump(chain)
	tx := Transaction{Sender: address, Receiver: "him", Value: 10}
	tx.Sign(privKey)
	chain.AddTransaction(tx)
	spew.Dump(chain)
	chain.MineNewBlock(privKey)
	spew.Dump(chain)
}
