package api

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/defany/dicego/pkg/format"
	"github.com/defany/dicego/pkg/parse"
)

var CoinsBalanceNotRegistered = errors.New("not registered")

type CoinsBalanceReq struct {
	User int `json:"user"`
}

type CoinsBalanceRes struct {
	Balance  float64 `json:"balance"`
	IsHiding int     `json:"is_hiding"`
}

type coinsBalanceRes struct {
	Balance  string `json:"balance"`
	IsHiding int    `json:"is_hiding"`
}

func NewCoinsBalanceReq(user int) CoinsBalanceReq {
	return CoinsBalanceReq{
		User: user,
	}
}

// PrettyBalance Баланс пользователя в user friendly виде.
func (c *CoinsBalanceRes) PrettyBalance() string {
	return format.NumWithSpaces(c.Balance)
}

// CoinsBalance Получение баланса пользователя по его айди.
func (d *Dice) CoinsBalance(ctx context.Context, params CoinsBalanceReq) (CoinsBalanceRes, error) {
	res, err := d.sendRequest(ctx, CoinsBalance, params)
	if err != nil {
		return CoinsBalanceRes{}, err
	}

	var privateBody coinsBalanceRes

	if err = json.Unmarshal(res, &privateBody); err != nil {
		return CoinsBalanceRes{}, err
	}

	if privateBody.Balance == CoinsBalanceNotRegistered.Error() {
		return CoinsBalanceRes{}, CoinsBalanceNotRegistered
	}

	publicBody := CoinsBalanceRes{
		IsHiding: privateBody.IsHiding,
	}

	publicBody.Balance, err = parse.FloatWithoutExtra(privateBody.Balance)
	if err != nil {
		return CoinsBalanceRes{}, err
	}

	return publicBody, nil
}
