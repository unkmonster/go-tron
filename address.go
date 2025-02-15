package gotron

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type Address [AddressLength]byte

func (a *Address) SetBytes(b []byte) {
	copy(a[:], b)
}

func (a *Address) Hex() string {
	return hex.EncodeToString(a[:])
}

func (a *Address) Bytes() []byte {
	return (*a)[:]
}

func (a *Address) Base58Check() string {
	return Base58CheckEncode(a.Bytes())
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

func PubkeyToAddress(p ecdsa.PublicKey) Address {
	ethAddr := crypto.PubkeyToAddress(p)
	return EthAddressToAddress(ethAddr)
}

// 以太坊地址+0x41前缀
func EthAddressToAddress(eth common.Address) Address {
	return BytesToAddress(append([]byte{0x41}, eth.Bytes()...))
}
