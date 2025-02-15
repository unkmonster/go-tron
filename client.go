package gotron

import (
	"context"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	c *resty.Client
}

func New(httpApi string, apiKey string) *Client {
	c := resty.New()
	c.SetBaseURL(httpApi)
	if apiKey != "" {
		c.SetHeader("TRON-PRO-API-KEY", apiKey)
	}

	return &Client{
		c: c,
	}
}

func (c *Client) Close() {

}

func (c *Client) Ping(ctx context.Context) error {
	_, err := c.c.R().SetContext(ctx).Get("")
	return err
}

func (c *Client) CreateTransaction(ctx context.Context, params *CreateTransactionParams) (*Transaction, error) {
	resp, err := c.c.R().
		SetContext(ctx).
		SetBody(params).
		SetResult(&Transaction{}).
		Post("/wallet/createtransaction")
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Transaction), nil
}

func (c *Client) BroadcastTransaction(ctx context.Context, txn *Transaction) (*BroadcastTransactionResult, error) {
	resp, err := c.c.R().
		SetContext(ctx).
		SetBody(txn).
		SetResult(&BroadcastTransactionResult{}).
		Post("/wallet/broadcasttransaction")
	if err != nil {
		return nil, err
	}
	return resp.Result().(*BroadcastTransactionResult), nil
}

// 获取区块 如果区块不存在返回空对象
func (c *Client) GetBlock(ctx context.Context, solid bool, params *GetBlockParams) (*Block, error) {
	var path string
	if solid {
		path = "/walletsolidity/getblock"
	} else {
		path = "/wallet/getblock"
	}

	resp, err := c.c.R().
		SetContext(ctx).
		SetBody(params).
		SetResult(&Block{}).
		Post(path)
	if err != nil {
		return nil, err
	}
	return resp.Result().(*Block), err
}

func (c *Client) TriggerSmartContract(ctx context.Context, params *TriggerSmartContractParams) (*TriggerSmartContractResult, error) {
	resp, err := c.c.R().
		SetContext(ctx).
		SetBody(params).
		SetResult(&TriggerSmartContractResult{}).
		Post("/wallet/triggersmartcontract")
	if err != nil {
		return nil, err
	}
	return resp.Result().(*TriggerSmartContractResult), nil
}
