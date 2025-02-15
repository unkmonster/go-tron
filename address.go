package gotron

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

type Address [AddressLength]byte

func (a *Address) SetBytes(b []byte) {
	copy(a[:], b)
}

func (a *Address) Hex() string {
	return hex.EncodeToString(a[:])
}

func BytesToAddress(b []byte) Address {
	var a Address
	a.SetBytes(b)
	return a
}

func HexToAddress(h string) Address {
	b, _ := hex.DecodeString(h)
	return BytesToAddress(b)
}

// 以太坊地址+0x41前缀
func PubkeyToAddress(p ecdsa.PublicKey) Address {
	ethAddr := crypto.PubkeyToAddress(p)
	return BytesToAddress(append([]byte{0x41}, ethAddr.Bytes()...))
}
