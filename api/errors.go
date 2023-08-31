package api

type Error struct {
	Status string `json:"status"`
	Desc   string `json:"desc"`
	Code   int    `json:"code"`
}

const (
	ErrPaymentToMyself    = "Никогда в тебе не сомневался!"
	ErrWriteRequestsLimit = "Записывающие запросы ограничены: 1 запрос / секунду. Подожди 2 секунды."
)
