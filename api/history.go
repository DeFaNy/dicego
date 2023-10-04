package api

import (
	"context"
	"encoding/json"
	"github.com/defany/dicego/pkg/format"
)

type CoinsHistoryReq struct {
	Limit      uint8  `json:"limit"`
	Offset     uint32 `json:"offset"`
	WithFailed uint8  `json:"with_failed,omitempty"`
}

type HistoryPayment struct {
	ID       string  `json:"id"`
	FromUser int     `json:"from_user"`
	ToUser   int     `json:"to_user"`
	Amount   float64 `json:"amount"`
	Date     int     `json:"date"`
	Success  int     `json:"success,omitempty"`
	Payload  int     `json:"payload,omitempty"`
}

type CoinsHistoryRes struct {
	User    int              `json:"user"`
	Limit   int              `json:"limit,omitempty"`
	Offset  int              `json:"offset,omitempty"`
	History []HistoryPayment `json:"history"`
}

func NewCoinsHistoryReq() CoinsHistoryReq {
	return CoinsHistoryReq{}
}

// WithLimit Будет возвращено только n транзакций.
func (c *CoinsHistoryReq) WithLimit(n uint8) {
	c.Limit = n
}

// WithOffset Устанавливает offset. При запросе будут пропущено n транзакций.
func (c *CoinsHistoryReq) WithOffset(n uint32) {
	c.Offset = n
}

// WithFailedTx Будут возвращены также неудавшиеся транзакции.
func (c *CoinsHistoryReq) WithFailedTx() {
	c.WithFailed = 1
}

// PrettyAmount Получение суммы перевода в user friendly виде.
func (h *HistoryPayment) PrettyAmount() string {
	return format.NumWithSpaces(h.Amount)
}

// CoinsHistory Получение истории транзакций.
func (d *Dice) CoinsHistory(ctx context.Context, params CoinsHistoryReq) (CoinsHistoryRes, error) {
	res, err := d.sendRequest(ctx, CoinsHistory, params)
	if err != nil {
		return CoinsHistoryRes{}, err
	}

	var body CoinsHistoryRes

	if err = json.Unmarshal(res, &body); err != nil {
		return CoinsHistoryRes{}, err
	}

	return body, nil
}
