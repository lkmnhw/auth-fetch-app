package models

type CurrencyResponse struct {
	Error   ErrorCurrencyResponse `json:"error"`
	Date    string                `json:"date"`
	Info    InfoCurrencyResponse  `json:"info"`
	Query   QueryCurrencyResponse `json:"query"`
	Result  float64               `json:"result"`
	Success bool                  `json:"success"`
}

type ErrorCurrencyResponse struct {
	Code int32  `json:"code"`
	Info string `json:"info"`
	Type string `json:"type"`
}

type InfoCurrencyResponse struct {
	Rate       float64 `json:"rate"`
	Timestampp int64   `json:"timestampp"`
}

type QueryCurrencyResponse struct {
	Amount float64 `json:"amount"`
	From   string  `json:"from"`
	To     string  `json:"to"`
}
