package gotron

type CreateTransactionParams struct {
	OwnerAddr string `json:"owner_address"`
	ToAddr    string `json:"to_address"`
	Amount    int64  `json:"amount"`
	Visible   bool   `json:"visible"`
}

type BroadcastTransactionResult struct {
	Result  bool   `json:"result"`
	TxId    string `json:"txid"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

type GetBlockParams struct {
	Detail  bool   `json:"detail"`
	IdOrNum string `json:"id_or_num"`
}

type TriggerSmartContractParams struct {
	OwnerAddr    string `json:"owner_address"`
	ContractAddr string `json:"contract_address"`
	Data         string `json:"data"`
	FeeLimit     int64  `json:"fee_limit"`
}

type TriggerSmartContractResult struct {
	Result struct {
		Code    string `json:"code,omitempty"`
		Message string `json:"message,omitempty"`
		Result  bool   `json:"result,omitempty"`
	} `json:"result,omitempty"`
	Transaction *Transaction `json:"transaction"`
}
