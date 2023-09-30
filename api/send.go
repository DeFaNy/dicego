package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
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
	userID  *int
	payload string
}

func NewPaymentLink() PaymentLink {
	return PaymentLink{}
}

func (p *PaymentLink) WithUserID(id int) {
	p.userID = &id
}

func (p *PaymentLink) WithPayload(payload string) {
	p.payload = payload
}

func (p *PaymentLink) String() string {
	u := url.URL{
		Scheme: "https",
		Host:   "vk.com",
		Path:   fmt.Sprintf("/app%d", appID),
	}

	if p.userID != nil {
		u.Path += fmt.Sprintf("#%d", *p.userID)
	}

	if p.payload != "" {
		u.Path += fmt.Sprintf("/%s", p.payload)
	}

	return u.String()
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
