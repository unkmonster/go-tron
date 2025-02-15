package gotron

import (
	"crypto/sha256"

	"github.com/mr-tron/base58"
)

func Base58CheckEncode(bin []byte) string {
	h1 := sha256.Sum256(bin)
	h2 := sha256.Sum256(h1[:])
	checksum := h2[:4]

	return base58.Encode(append(bin, checksum...))
}

func Base58CheckDecode(text string) (bin []byte, err error) {
	decoded, err := base58.Decode(text)
	if err != nil {
		return
	}

	return decoded[:len(decoded)-4], nil
}
