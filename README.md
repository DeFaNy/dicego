# Библиотека для работы с VK Dice API
## Официальная группа игры: https://vk.com/vkdice
## По вопросам работы библиотеки: https://vk.com/defany

## Установка
```
go get github.com/defany/dicego
```

## Пример кода
```golang
package main

import (
	"context"
	"fmt"
	"github.com/defany/dicego/api"
	"log"
)

func example() {
	// Не подставляйте данные напрямую, берите их при помощи os.Getenv
	dice := api.NewDice("your-api-token", "your-merchant-id")

	ctx := context.Background()

	coinsBalanceParams := api.NewCoinsBalanceReq(222856843)

	res, err := dice.CoinsBalance(ctx, coinsBalanceParams)
	if err != nil {
		log.Println(fmt.Sprintf("failed to sent payment: %s", err.Error()))

		return
	}

	log.Println(fmt.Sprintf("Got user balance: %s", res.PrettyBalance()))

	sendPaymentParams := api.NewCoinsSendReq(297789589, 1_000)

	sendPaymentRes, err := dice.CoinsSend(ctx, sendPaymentParams)
	if err != nil {
		if err.Error() == api.ErrWriteRequestsLimit {
			log.Println("Слишком частые запросы переводов. Или записывающих реквестов в целом")
		} else {
			log.Println(fmt.Sprintf("failed to send payment: %s", err.Error()))
		}

		return
	}

	log.Println(fmt.Sprintf("succesfully sent payment: %+v", sendPaymentRes))

	paymentsHistoryParams := api.NewCoinsHistoryReq()
	paymentsHistoryParams.WithLimit(3)

	paymentsHistoryRes, err := dice.CoinsHistory(ctx, paymentsHistoryParams)
	if err != nil {
		log.Println(fmt.Sprintf("failed to get payments history: %s", err.Error()))

		return
	}

	log.Println("successfully got history, user id: ", paymentsHistoryRes.User)

	if len(paymentsHistoryRes.History) == 0 {
		return
	}

	prettyPaymentAmount := paymentsHistoryRes.History[0]

	log.Println("pretty payment amount: ", prettyPaymentAmount.PrettyAmount())
}

func main() {
	example()
}
```

# Уточнение.
Данная библиотека **не** добавляет возможности отлова уведомлений от игры.
Это сделано намерено, так как разработчик считает, что каждый другой человек сам выбирает что ему использовать для отлова событий.
И поэтому, дабы не тянуть лишние зависимости это добавлено не было. Однако, в будущем разработчик, возможно, по запросу, добавит поддержку дефолтного `net/http`.

Также библиотека не добавляет возможности получения `pretty` чисел от API. Для этого реализованы внутренние методы, которые, впрочем, точно также форматируют числа.
