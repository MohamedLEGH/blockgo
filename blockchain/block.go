package blockchain

import (
	"fmt"
	"errors"
	"strings"
	"math/big"
)

type Block struct {
	Index         int
	Previous_hash string
	Nonce         int
	Timestamp     float32
	Miner_address string
	Hash_val      string
	Tx_list       []Transaction
}

func (b *Block) AddTransaction(t Transaction) {
	b.Tx_list = append(b.Tx_list, t)
}

func (b *Block) Hash(nonce int) string {
	var buf strings.Builder
    fmt.Fprintf(&buf, "%d|%s|%f|%s|", 
        b.Index, 
        b.Previous_hash, 
        b.Timestamp, 
        b.Miner_address)
	for _, tx := range b.Tx_list {
		fmt.Fprintf(&buf, "%v|", tx) 
	}
	fmt.Fprintf(&buf, "%d", nonce)
	return Hash(buf.String())
}

func (b *Block) Mine(difficulty int64) (int, string, error) {
	target, err := GetTarget(difficulty)
	if err != nil {
		return 0, "", err
	}
	// fmt.Println("Target:", target)
	nonce := 0
	hash := b.Hash(nonce)
	hash_val, ok := new(big.Int).SetString(hash, 16)
    if !ok {
        return 0, "", errors.New("failed to parse hash")
    }
	// hash > target
	// fmt.Println("Hash val:", hash_val)
	for hash_val.Cmp(target) == 1 {  
		nonce++
		hash = b.Hash(nonce)
		hash_val, ok = new(big.Int).SetString(hash, 16)
		// fmt.Println("Target:", target)
		// fmt.Println("Hash val:", hash_val)
		if !ok {
			return 0, "", errors.New("failed to parse hash")
		}
	}
	b.Nonce = nonce
	b.Hash_val = "0x"+hash
	return nonce, hash, nil
}

func GetTarget(difficulty int64) (*big.Int, error) {
	if difficulty <= 0 {
		return nil, errors.New("difficulty must be greater than 0")
	}
	difficulty_big := big.NewInt(difficulty)

	// initial_target := "00000000ffff0000000000000000000000000000000000000000000000000000"
	// initial target of Bitcoin is already a very small number, for test I need a bigger target
	initial_target := "f0000000ffff0000000000000000000000000000000000000000000000000000"
	n_initial_target := new(big.Int)
	n_initial_target, ok := n_initial_target.SetString(initial_target, 16)
    if !ok {
        return nil, errors.New("failed to parse initial target")
    }
	target := new(big.Int).Div(n_initial_target, difficulty_big)
	return target, nil
}