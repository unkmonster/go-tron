package gotron

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"

	"github.com/ethereum/go-ethereum/crypto"
)

// 不同类型的交易的交易结构是相同的，但是raw_data中包含的内容不同，这主要反映在contract.parameter.value中。
type contract struct {
	Parameter struct {
		Value struct {
			OwnerAddress    string `json:"owner_address"`
			ToAddress       string `json:"to_address"`
			Amount          int64  `json:"amount"`
			Data            string `json:"data,omitempty"`             // only for smart contract
			ContractAddress string `json:"contract_address,omitempty"` // only for smart contract
		} `json:"value"`
		TypeUrl string `json:"type_url"`
	} `json:"parameter"`
	Type string `json:"type"`
}

// contract[0].parameter.value.data: 前4字节 function_selector,
type Transaction struct {
	TxId    string `json:"txID"`
	Visible bool   `json:"visible"`
	RawData struct {
		Contract      []contract `json:"contract"`
		Timestamp     int64      `json:"timestamp"`
		RefBlockBytes string     `json:"ref_block_bytes"`
		RefBlockHash  string     `json:"ref_block_hash"`
		Expiration    int64      `json:"expiration"`
		FeeLimit      int64      `json:"fee_limit"`
	} `json:"raw_data"`
	RawDataHex string   `json:"raw_data_hex"`
	Signature  []string `json:"signature,omitempty"`
	Error      string   `json:"Error,omitempty"`
}

type BroadResult struct {
	Result  bool   `json:"result"`
	Txid    string `json:"txid"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type BlockHeader struct {
	RawData struct {
		Number         int64  `json:"number"` // 区块高度
		TimeStamp      int64  `json:"timestamp"`
		WitnessAddress string `json:"witness_address"`
		ParentHash     string `json:"parentHash"`
		Version        int    `json:"version"`
	} `json:"raw_data"`
	WitnessSignature string `json:"witness_signature"`
}

type Block struct {
	BlockId      string         `json:"blockID"` // 区块哈希
	BlockHeader  BlockHeader    `json:"block_header"`
	Transactions []*Transaction `json:"transactions"`
}

func (txn *Transaction) Sign(prv *ecdsa.PrivateKey) error {
	data, err := hex.DecodeString(txn.RawDataHex)
	if err != nil {
		return err
	}

	hash := sha256.Sum256(data)
	sig, err := crypto.Sign(hash[:], prv)
	if err != nil {
		return err
	}

	txn.Signature = append(txn.Signature, hex.EncodeToString(sig))
	return nil
}
