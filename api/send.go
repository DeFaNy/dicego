package api

import (
	"context"
	"encoding/json"
)

type CoinsSendReq struct {
	User   int     `json:"user"`
	Amount float64 `json:"amount"`
}

type CoinsSendRes struct {
}

func NewCoinsSendReq(user int, amount float64) CoinsSendReq {
	return CoinsSendReq{
		User:   user,
		Amount: amount,
	}
}

// CoinsSend Метод для отправки перевода другому пользователю
func (d *Dice) CoinsSend(ctx context.Context, params CoinsSendReq) (CoinsSendRes, error) {
	res, err := d.sendRequest(ctx, CoinsSend, params)
	if err != nil {
		return CoinsSendRes{}, err
	}

	var resBody CoinsSendRes

	if err = json.Unmarshal(res, &resBody); err != nil {
		return CoinsSendRes{}, err
	}

	return resBody, err
}
