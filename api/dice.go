package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type Method string

const (
	CoinsBalance Method = "coins.balance"
	CoinsHistory Method = "coins.history"
	CoinsSend    Method = "coins.send"
	ApiCallback  Method = "api.callback"
	ApiRename    Method = "api.rename"
)

var ErrfailedErrorCheck = errors.New("failed to check error")

const apiURL = "https://api-dice.belle.dev/v2"

type Dice struct {
	token      string
	merchantID string
	httpClient *http.Client
}

func NewDice(token string, merchantID string) Dice {
	return Dice{
		token:      token,
		merchantID: merchantID,
		httpClient: &http.Client{},
	}
}

func (d *Dice) sendRequest(ctx context.Context, method Method, params any) ([]byte, error) {
	body := new(bytes.Buffer)

	if err := json.NewEncoder(body).Encode(params); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodGet, d.url(method), body)
	if err != nil {
		return nil, err
	}

	req.WithContext(ctx)

	req.Header.Set("dc-key", d.token)

	res, err := d.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err = d.checkError(res.StatusCode, resBody); err != nil {
		return nil, err
	}

	return resBody, nil
}

func (d *Dice) checkError(code int, res []byte) error {
	if code == http.StatusOK {
		return nil
	}

	var errBody Error

	if err := json.Unmarshal(res, &errBody); err != nil {
		return errors.Join(ErrfailedErrorCheck, err)
	}

	return errors.New(errBody.Desc)
}

func (d *Dice) url(method Method) string {
	return fmt.Sprintf("%s/%s?m=%s", apiURL, method, d.merchantID)
}
