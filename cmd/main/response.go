package main

type Response struct {
	Status Status          `json:"status"`
	Data   map[string]Data `json:"data"`
}

type Status struct {
	Timestamp    string      `json:"timestamp"`
	ErrorCode    int64       `json:"error_code"`
	ErrorMessage interface{} `json:"error_message"`
	Elapsed      int64       `json:"elapsed"`
	CreditCount  int64       `json:"credit_count"`
	Notice       interface{} `json:"notice"`
}

type Data struct {
	Symbol      string           `json:"symbol"`
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Amount      float64          `json:"amount"`
	LastUpdated string           `json:"last_updated"`
	Quote       map[string]Quote `json:"quote"`
}

type Quote struct {
	Price       float64 `json:"price"`
	LastUpdated string  `json:"last_updated"`
}
