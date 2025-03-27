package blockchain

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

type Block struct {
	Index         int
	Previous_hash string
	Nonce         int
	Timestamp     int64
	Miner_address string
	Hash_val      string
	Tx_list       []Transaction
}

func (b *Block) AddTransaction(t Transaction) {
	b.Tx_list = append(b.Tx_list, t)
}

func (b *Block) Hash(nonce int) string {
	var buf strings.Builder
	fmt.Fprintf(&buf, "%d|%s|%d|%s|",
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
	nonce := 0
	hash := b.Hash(nonce)
	hash_val, ok := new(big.Int).SetString(hash, 16)
	if !ok {
		return 0, "", errors.New("failed to parse hash")
	}
	for hash_val.Cmp(target) == 1 {
		nonce++
		hash = b.Hash(nonce)
		hash_val, ok = new(big.Int).SetString(hash, 16)
		if !ok {
			return 0, "", errors.New("failed to parse hash")
		}
	}
	hash = "0x" + hash
	b.Nonce = nonce
	b.Hash_val = hash
	return nonce, hash, nil
}

func (b *Block) VerifyBlock() bool {
	nonce := b.Nonce
	hash := b.Hash_val
	hash_computed := b.Hash(nonce)
	if hash == hash_computed {
		return true
	} else {
		return false
	}
}

func GetTarget(difficulty int64) (*big.Int, error) {
	if difficulty <= 0 {
		return nil, errors.New("difficulty must be greater than 0")
	}
	difficulty_big := big.NewInt(difficulty)
	initial_target := "f0000000ffff0000000000000000000000000000000000000000000000000000"
	n_initial_target := new(big.Int)
	n_initial_target, ok := n_initial_target.SetString(initial_target, 16)
	if !ok {
		return nil, errors.New("failed to parse initial target")
	}
	target := new(big.Int).Div(n_initial_target, difficulty_big)
	return target, nil
}
