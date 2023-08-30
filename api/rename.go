package api

import (
	"context"
	"encoding/json"
)

type ApiRenameReq struct {
	Name string `json:"name"`
}

type ApiRenameRes struct {
}

func NewApiRenameReq(name string) ApiRenameReq {
	return ApiRenameReq{
		Name: name,
	}
}

// ApiRename Запрос для переименования маркета.
func (d *Dice) ApiRename(ctx context.Context, params ApiRenameReq) (ApiRenameRes, error) {
	res, err := d.sendRequest(ctx, ApiRename, params)
	if err != nil {
		return ApiRenameRes{}, err
	}

	var resBody ApiRenameRes

	if err = json.Unmarshal(res, &resBody); err != nil {
		return ApiRenameRes{}, err
	}

	return resBody, nil
}
