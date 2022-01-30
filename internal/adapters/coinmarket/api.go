package coinmarket

import (
	"time"

	"github.com/shopspring/decimal"
)

type PriceConversionResponseSanBox struct {
	Status ResponseStatus        `json:"status"`
	Data   map[string]SymbolData `json:"data"`
}

type PriceConversionResponsePro struct {
	Status ResponseStatus `json:"status"`
	Data   SymbolData     `json:"data"`
}

type SymbolData struct {
	Symbol      string           `json:"symbol"`
	ID          string           `json:"id"`
	Name        string           `json:"name"`
	Amount      decimal.Decimal  `json:"amount"`
	LastUpdated time.Time        `json:"last_updated"`
	Quote       map[string]Quote `json:"quote"`
}

type Quote struct {
	Price       decimal.Decimal `json:"price"`
	LastUpdated time.Time       `json:"last_updated"`
}

type ResponseStatus struct {
	Time         time.Time `json:"timestamp"`
	ErrorCode    int       `json:"error_code"`
	ErrorMessage string    `json:"error_message"`
	Elapsed      int       `json:"elapsed"`
	CreditCount  int       `json:"credit_count"`
}
