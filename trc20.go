package gotron

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

var TRC20Abi abi.ABI

func init() {
	var err error
	TRC20Abi, err = abi.JSON(strings.NewReader(TRC20ABI_JSON))
	if err != nil {
		panic(err)
	}
}
