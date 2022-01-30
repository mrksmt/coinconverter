package ports

import (
	"context"

	"github.com/shopspring/decimal"
)

type PriceGetter interface {
	GetPrice(ctx context.Context, symbolA, symbolB string) (decimal.Decimal, error)
}

type CoinConverter interface {
	Convert(amount decimal.Decimal, symbolA, symbolB string) (decimal.Decimal, error)
}
