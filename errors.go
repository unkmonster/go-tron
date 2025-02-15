package gotron

import (
	"errors"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type HTTPError struct {
	Code int
	Body string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("http error, code: %d, body: %s", e.Code, e.Body)
}

func CheckRespStatus(resp *resty.Response) error {
	if resp.StatusCode() >= 400 {
		return &HTTPError{
			Code: resp.StatusCode(),
			Body: resp.String(),
		}
	}
	return nil
}

var ErrInvalidAmount = errors.New("amount must be greater than 0")
var ErrTransferTRXToYourself = errors.New("could not send TRX to yourself")
var ErrInvalidAddr = errors.New("invalid TRON primitive address")
