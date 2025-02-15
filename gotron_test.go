package gotron

import (
	"context"
	"encoding/hex"
	"math/big"
	"testing"

	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

var cli = New("https://api.shasta.trongrid.io", "")
var ctx = context.Background()

var testAddr1 = "41c6e9be0a5dee6b995d47c111c1f01f7d896d51eb"
var testKey1 = "ac52aa609aa95b2c09094528a4981d2ac06c01b14a197240bf3167b20796fdf1"
var testBase58CheckAddr1 = "TU6xok89KFMZs4GEHCbJGhcdHYmt694EF7"
var testAddr2 = "41cbc56bee2abe672274e5f90db9aa91f7ceb94a6f"
var testKey2 = "1a1dcf54e92ded1ce8ec65bf6f7d47ef8ff3287ffba750609bc0b328cb710b64"

func TestPing(t *testing.T) {
	assert.NoError(t, cli.Ping(ctx))
}

func TestCreateTransaction(t *testing.T) {
	params := CreateTransactionParams{
		OwnerAddr: testAddr1,
		ToAddr:    testAddr2,
		Amount:    1,
		Visible:   false,
	}

	// should be success
	txn, err := cli.CreateTransaction(ctx, &params)
	assert.NoError(t, err)
	assert.Equal(t, params.Amount, txn.RawData.Contract[0].Parameter.Value.Amount)

	// should be fail
	params.ToAddr = params.OwnerAddr
	txn, err = cli.CreateTransaction(ctx, &params)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, txn.Error)
	}
}

func TestBroadcastTransaction(t *testing.T) {
	params := CreateTransactionParams{
		OwnerAddr: testAddr1,
		ToAddr:    testAddr2,
		Amount:    1,
		Visible:   false,
	}

	// 创建交易
	txn, err := cli.CreateTransaction(ctx, &params)
	if !assert.NoError(t, err) {
		return
	}

	// 签名交易
	priv, err := crypto.HexToECDSA(testKey1)
	if !assert.NoError(t, err) {
		return
	}
	assert.NoError(t, txn.Sign(priv))

	// 广播交易
	result, err := cli.BroadcastTransaction(ctx, txn)
	if !assert.NoError(t, err) {
		return
	}
	assert.True(t, result.Result)
	assert.NotEmpty(t, txn.TxId)
}

func TestGetBlock(t *testing.T) {
	params := GetBlockParams{
		Detail:  true,
		IdOrNum: "",
	}

	block, err := cli.GetBlock(ctx, false, &params)
	if !assert.NoError(t, err) {
		return
	}
	assert.NotEmpty(t, block.BlockId)

	blocksolidity, err := cli.GetBlock(ctx, true, &params)
	if !assert.NoError(t, err) {
		return
	}
	assert.NotEmpty(t, blocksolidity.BlockId)

	// 最新的固化的区块的高度一定小于最新的区块
	assert.Less(t, blocksolidity.BlockHeader.RawData.Number, block.BlockHeader.RawData.Number)

	// should be fail
	params.IdOrNum = "-1"
	block, err = cli.GetBlock(ctx, true, &params)
	if assert.NoError(t, err) {
		assert.Empty(t, block.BlockId)
	}
}

func TestTriggerSmartContract(t *testing.T) {
	shastaUsdt := "4142a1e39aefa49290f2b3f9ed688d7cecf86cd6e0"

	toAddr := ethcommon.BytesToAddress(ethcommon.Hex2Bytes(testAddr2)[1:])
	amount := big.NewInt(1e6)

	data, err := TRC20Abi.Pack("transfer", toAddr, amount)
	if !assert.NoError(t, err) {
		return
	}

	result, err := cli.TriggerSmartContract(ctx, &TriggerSmartContractParams{
		OwnerAddr:    testAddr1,
		ContractAddr: shastaUsdt,
		Data:         hex.EncodeToString(data),
		FeeLimit:     1000000000,
	})
	if !assert.NoError(t, err) {
		return
	}
	assert.True(t, result.Result.Result)
	assert.NotEmpty(t, result.Transaction.TxId)

	priv, err := crypto.HexToECDSA(testKey1)
	if !assert.NoError(t, err) {
		return
	}
	result.Transaction.Sign(priv)

	br, err := cli.BroadcastTransaction(ctx, result.Transaction)
	if !assert.NoError(t, err) {
		return
	}
	assert.True(t, br.Result)
}

func TestBase58Check(t *testing.T) {
	text := "hello world"
	encoded := Base58CheckEncode([]byte(text))
	assert.NotEqual(t, text, encoded)

	decoded, err := Base58CheckDecode(encoded)
	if assert.NoError(t, err) {
		assert.Equal(t, text, string(decoded))
	}
}

func TestPubkeyToAddress(t *testing.T) {
	key, err := crypto.GenerateKey()
	assert.NoError(t, err)
	address := PubkeyToAddress(key.PublicKey)
	assert.Equal(t, uint8(0x41), address[0])
}

func TestAddressBase58Check(t *testing.T) {
	addr := HexToAddress(testAddr1)
	base58Check := addr.Base58Check()
	assert.Equal(t, testBase58CheckAddr1, base58Check)
}
