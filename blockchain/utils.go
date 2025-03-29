package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/txscript"
	"log"
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
	hash := sha256.Sum256([]byte(msg))
	signature, err := schnorr.Sign(privKey, hash[:])
	if err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(signature.Serialize())
}

func VerifySignature(taprootAddress string, msg string, sig string) {
	hash := sha256.Sum256([]byte(msg))

	addr, err := btcutil.DecodeAddress(taprootAddress, &chaincfg.MainNetParams)
	if err != nil {
		log.Fatal(err)
	}
	pubKeyBytes := addr.ScriptAddress()
	pubKey, err := schnorr.ParsePubKey(pubKeyBytes)
	if err != nil {
		log.Fatal(err)
	}

	sigBytes, err := hex.DecodeString(sig)
	if err != nil {
		log.Fatal(err)
	}
	signature, err := schnorr.ParseSignature(sigBytes)
	if err != nil {
		log.Fatal(err)
	}
	signature.Verify(hash[:], pubKey)
}
