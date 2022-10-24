package edgefuzz

import (
	"crypto/ecdsa"
	"github.com/0xPolygon/polygon-edge/crypto"
	"github.com/coinbase/kryptology/pkg/signatures/bls/bls_sig"
	"github.com/0xPolygon/polygon-edge/consensus/ibft/signer"
)
func newECDSAKey() (*ecdsa.PrivateKey, []byte) {
	key, keyEncoded, _ := crypto.GenerateAndEncodeECDSAPrivateKey()
	return key, keyEncoded
}

func newBLSKey() (*bls_sig.SecretKey, []byte) {
	key, keyEncoded, _ := crypto.GenerateAndEncodeBLSSecretKey()
	return key, keyEncoded
}

func Fuzz(data []byte) int {
	ecdsaKey, _ := newECDSAKey()
	blsKey, _ := newBLSKey()
	if ecdsaKey == nil || blsKey == nil {
		return 1
	}

	// Create BLSKeyManager
	blsKeyManager := signer.NewBLSKeyManagerFromKeys(ecdsaKey, blsKey)
	msg := crypto.Keccak256(data)

	proposerSeal, _ := blsKeyManager.SignIBFTMessage(msg)
	blsKeyManager.Ecrecover(proposerSeal, msg)
	return 0
}
