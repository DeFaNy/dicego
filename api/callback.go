package api

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
)

type ApiCallbackReq struct {
	Callback string `json:"callback"`
}

type ApiCallbackRes struct {
	Desc string `json:"desc"`
	Url  string `json:"url"`
}

type CallbackNotification struct {
	ID          string  `json:"id"`
	FromUser    int     `json:"from_user"`
	Amount      float64 `json:"amount"`
	Date        int     `json:"date"`
	Payload     int     `json:"payload,omitempty"`
	RequestHash string  `json:"request_hash,omitempty"`
}

func NewApiCallbackReq(callback string) ApiCallbackReq {
	return ApiCallbackReq{
		Callback: callback,
	}
}

// ApiCallback Установка новой ссылки для отправки уведомлений на ваш сервер.
func (d *Dice) ApiCallback(ctx context.Context, params ApiCallbackReq) (ApiCallbackRes, error) {
	res, err := d.sendRequest(ctx, ApiCallback, params)
	if err != nil {
		return ApiCallbackRes{}, err
	}

	var resBody ApiCallbackRes

	if err = json.Unmarshal(res, &resBody); err != nil {
		return ApiCallbackRes{}, err
	}

	return resBody, nil
}

// IsNotificationValid Проверка уведомления, которое прислал вам сервер.
func (d *Dice) IsNotificationValid(notif CallbackNotification) bool {
	requestHash := notif.RequestHash

	notif.RequestHash = ""

	res, err := json.Marshal(notif)
	if err != nil {
		return false
	}

	hasher := md5.New()
	hasher.Write([]byte(d.token))
	hasher.Write([]byte("|"))
	hasher.Write(res)

	return hex.EncodeToString(hasher.Sum(nil)) == requestHash
}
