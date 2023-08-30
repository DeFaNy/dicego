package main

import (
	"context"
	"fmt"
	"github.com/defany/dicego/api"
	"log"
)

func example() {
	// Не подставляйте данные напрямую, берите их при помощи os.GetEnv
	dice := api.NewDice("your_token", "your_merchant_id")

	ctx := context.Background()

	res, err := dice.ApiRename(ctx, api.NewApiRenameReq("my-awesome-market"))

	log.Println(fmt.Sprintf("%+v", res), err)
}

func main() {
	example()
}
