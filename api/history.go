package api

import (
	"context"
	"encoding/json"
)

type CoinsHistoryReq struct {
	Limit      uint8  `json:"limit"`
	Offset     uint32 `json:"offset"`
	Pretty     uint8  `json:"pretty"`
	WithFailed uint8  `json:"with_failed"`
}

type CoinsHistoryRes struct {
}

func NewCoinsHistoryReq() CoinsBalanceReq {
	return CoinsBalanceReq{}
}

// WithLimit Будет возвращено только n транзакций.
func (c *CoinsHistoryReq) WithLimit(n uint8) {
	c.Limit = n
}

// WithOffset Устанавливает offset. При запросе будут пропущено n транзакций.
func (c *CoinsHistoryReq) WithOffset(n uint32) {
	c.Offset = n
}

// MakePretty Суммы переводов будут разделены пробелами.
func (c *CoinsHistoryReq) MakePretty() {
	c.Pretty = 1
}

// WithFailedTx Будут возвращены также неудавшиеся транзакции.
func (c *CoinsHistoryReq) WithFailedTx() {
	c.WithFailed = 1
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
