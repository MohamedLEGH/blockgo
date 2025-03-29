package main

import (
	"blockgo/blockchain"
	"fmt"
)

type Transaction = blockchain.Transaction
type Block = blockchain.Block
type Chain = blockchain.Chain

func main() {
	privKey, address := blockchain.GenerateAccount()
	// privKey, _ := blockchain.GenerateAccount()
	fmt.Printf("Address:%s\n", address)
	// fmt.Printf("%s\n", privKey, address)
	msg := "test message"
	signature := blockchain.SignMessage(privKey, msg)
	blockchain.VerifySignature(address, msg, signature)
	fmt.Printf("Signature:%s\n", signature)
	miner_address := "me"
	chain := Chain{Difficulty: 1}
	chain.CreateGenesisBlock(miner_address)
	fmt.Printf("%+v\n", chain)
	tx := Transaction{Sender: address, Receiver: "him", Value: 10}
	tx.Sign(privKey)
	tx.Verify()
	chain.AddTransaction(tx)
	fmt.Printf("%+v\n", chain)
	chain.MineNewBlock(miner_address)
	fmt.Printf("%+v\n", chain)
}
