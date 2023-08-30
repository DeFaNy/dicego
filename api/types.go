package api

type Method string

const (
	CoinsBalance Method = "coins.balance"
	CoinsHistory Method = "coins.history"
	CoinsSend    Method = "coins.send"
	ApiCallback  Method = "api.callback"
	ApiRename    Method = "api.rename"
)
