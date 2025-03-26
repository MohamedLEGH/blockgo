package blockchain

import (
	"crypto/sha256"
	"fmt"
)

func Hash(s string) string {
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}
