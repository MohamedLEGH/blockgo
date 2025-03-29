package blockchain

import (
	"fmt"
	"time"
)

type Chain struct {
	Difficulty int64
	Block_list []Block
	Tx_pool    []Transaction
}

func (c *Chain) AddTransaction(t Transaction) {
	t.Verify()
	c.Tx_pool = append(c.Tx_pool, t)
}

func (c *Chain) AddBlock(b Block) {
	b.Verify()
	c.Block_list = append(c.Block_list, b)
}

// func MineBlock(b Block) {
// }

func (c *Chain) CreateGenesisBlock(hexKey string) {
	// need to check that blocklist is empty
	miner_address := TapRootAddressFromPrivateKey(hexKey)
	index := 0
	timestamp := time.Now().Unix()
	previous_hash := ""
	tx_list := []Transaction{}
	block := Block{Index: index, Previous_hash: previous_hash, Timestamp: timestamp, Miner_address: miner_address, Tx_list: tx_list}
	_, _, err := block.Mine(c.Difficulty)
	if err != nil {
		fmt.Printf("Error mining block")
	}
	block.Sign(hexKey)
	c.AddBlock(block)
}

func (c *Chain) MineNewBlock(hexKey string) {
	miner_address := TapRootAddressFromPrivateKey(hexKey)
	// need to check that block list is not empty
	last_block := c.Block_list[len(c.Block_list)-1]
	index := last_block.Index + 1
	timestamp := time.Now().Unix()
	previous_hash := last_block.Hash_val
	tx_list := []Transaction{}
	for _, tx := range c.Tx_pool {
		tx_list = append(tx_list, tx)
	}
	c.Tx_pool = []Transaction{}
	block := Block{Index: index, Previous_hash: previous_hash, Timestamp: timestamp, Miner_address: miner_address, Tx_list: tx_list}
	_, _, err := block.Mine(c.Difficulty)
	if err != nil {
		fmt.Printf("Error mining block")
	}
	block.Sign(hexKey)
	c.AddBlock(block)
}

func (c *Chain) VerifyChain() bool {
	for _, block := range c.Block_list {
		if !block.Verify() {
			return false
		}
	}
	return true
}
