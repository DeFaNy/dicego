# Библиотека для работы с VK Dice API
## Официальная группа игры: https://vk.com/vkdice

## Установка
```
go get github.com/defany/dicego
```

## Пример кода
```golang
package main

import (
	"context"
	"github.com/defany/dicego/api"
	"fmt"
	"log"
)

func example() {
	// Не подставляйте данные напрямую, берите их при помощи os.GetEnv
	dice := api.NewDice("your_token", "your_merchant_id")

	ctx := context.Background()

	res, err := dice.ApiRename(ctx, internal.NewApiRenameReq("my-awesome-market"))

	log.Println(fmt.Sprintf("%+v", res), err)
}

func main() {
	example()
}

```

# Уточнение.
Данная библиотека **не** добавляет возможности отлова уведомлений от игры.
Это сделано намерено, так как разработчик считает, что каждый другой человек сам выбирает что ему использовать для отлова событий.
И поэтому, дабы не тянуть лишние зависимости это добавлено не было. Однако, в будущем разработчик, возможно, по запросу, добавит поддержку дефолтного `net/http`.
