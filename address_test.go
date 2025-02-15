package gotron

import (
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestPubkeyToAddress(t *testing.T) {
	key, err := crypto.GenerateKey()
	assert.NoError(t, err)
	address := PubkeyToAddress(key.PublicKey)
	assert.Equal(t, uint8(0x41), address[0])
}
