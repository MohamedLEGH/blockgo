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
	c.Tx_pool = append(c.Tx_pool, t)
}

func (c *Chain) AddBlock(b Block) {
	c.Block_list = append(c.Block_list, b)
}

// func MineBlock(b Block) {

// }

func (c *Chain) CreateGenesisBlock(miner_address string) {
	// need to check that blocklist is empty
	index := 0
	timestamp := time.Now().Unix()
	previous_hash := ""
	tx_list := []Transaction{}
	block := Block{Index: index, Previous_hash: previous_hash, Timestamp: timestamp, Miner_address: miner_address, Tx_list: tx_list}
	_, _, err := block.Mine(c.Difficulty)
	if err != nil {
		fmt.Printf("Error mining block")
	}
	c.AddBlock(block)
}

func (c *Chain) MineNewBlock(miner_address string) {
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
	c.AddBlock(block)
}
