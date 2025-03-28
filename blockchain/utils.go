package blockchain

import (
	"encoding/hex"
	"crypto/sha256"
	"fmt"
	"log"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
)

func Hash(s string) string {
	hash := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", hash)
}

func GeneratePrivateKey() string {
	privKey, err := btcec.NewPrivateKey()
	if err != nil {
		log.Fatal(err)
	}
	hexKey := hex.EncodeToString(privKey.Serialize())
	return hexKey
}

func PublicFromPrivate(hexKey string) string {
	keyBytes, err := hex.DecodeString(hexKey)
	if err != nil {
		log.Fatal(err)
	}
	_, pubKey := btcec.PrivKeyFromBytes(keyBytes)
	compressedPubKey := hex.EncodeToString(pubKey.SerializeCompressed())
	return compressedPubKey
}

func TapRootAddressFromPublic(hexPubKey string) string {
	pubKeyBytes, err := hex.DecodeString(hexPubKey)
	if err != nil {
		log.Fatal(err)
	}
	pubKey, err := btcec.ParsePubKey(pubKeyBytes)
	if err != nil {
		log.Fatal(err)
	}
	taprootKey := txscript.ComputeTaprootKeyNoScript(pubKey)
	taprootPubKey := schnorr.SerializePubKey(taprootKey)
	address, err := btcutil.NewAddressTaproot(taprootPubKey, &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}
	return address.EncodeAddress()
}

func GenerateAccount() (string, string) {
	privKey := GeneratePrivateKey()
	pubKey := PublicFromPrivate(privKey)
	address := TapRootAddressFromPublic(pubKey)
	return privKey, address
}

func SignMessage(hexKey string, msg string) string {
	keyBytes, err := hex.DecodeString(hexKey)
	if err != nil {
		log.Fatal(err)
	}
	privKey, _ := btcec.PrivKeyFromBytes(keyBytes)
	hash := chainhash.DoubleHashB([]byte(msg))
	signature, err := schnorr.Sign(privKey, hash[:])
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(signature.Serialize())
}