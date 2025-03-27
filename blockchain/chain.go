package blockchain

import (
	"fmt"
    "time"
)

type Chain struct {
	// Timestamp int
	// Nonce int
	// Timestamp float32
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

func (c *Chain) CreateGenesisBlock() {
	index := 0
	timestamp := time.Now().Unix()
	previous_hash := ""
	miner_address := "me"
	tx_list := []Transaction{}
	block := Block{Index: index, Previous_hash: previous_hash, Timestamp: timestamp, Miner_address: miner_address, Tx_list: tx_list}
	_, _, err := block.Mine(c.Difficulty)
	if err != nil {
		fmt.Printf("Error mining block")
	}
	c.AddBlock(block)
}

// func (t Block) String() string {
// 	return fmt.Sprintf("index=%d, previous hash=%s, nonce=%d, timestamp=%f, miner address=%s, hash value=%s", t.Sender, t.Receiver, t.Value)
// }
