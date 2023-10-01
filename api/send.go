package api

import (
	"context"
	"encoding/json"
	"fmt"
)

type CoinsSendReq struct {
	User   int     `json:"user"`
	Amount float64 `json:"amount"`
}

type CoinsSendRes struct {
	User    int     `json:"user"`
	Balance float64 `json:"balance"`
	Amount  float64 `json:"amount"`
}

type PaymentLink struct {
	userID  *string
	payload *int
}

func NewPaymentLink() PaymentLink {
	return PaymentLink{}
}

func (p *PaymentLink) WithUserID(id string) {
	p.userID = &id
}

func (p *PaymentLink) WithPayload(payload int) {
	p.payload = &payload
}

func (p *PaymentLink) String() string {
	u := fmt.Sprintf("https://vk.com/app%d", appID)

	if p.userID != nil {
		u += fmt.Sprintf("#%s", *p.userID)
	}

	if p.payload != nil {
		u += fmt.Sprintf("/%d", *p.payload)
	}

	return u
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
