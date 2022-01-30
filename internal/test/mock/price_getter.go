package mock

import (
	"context"
	"converter/internal/core/ports"
	"fmt"

	"github.com/shopspring/decimal"
)

type PriceGetterFix struct {
	Price decimal.Decimal
}

var _ ports.PriceGetter = (*PriceGetterFix)(nil)

func (pg *PriceGetterFix) GetPrice(ctx context.Context, symbolA, symbolB string) (decimal.Decimal, error) {
	return pg.Price, nil
}

type PriceGetterErr struct{}

var _ ports.PriceGetter = (*PriceGetterFix)(nil)

func (pg *PriceGetterErr) GetPrice(ctx context.Context, symbolA, symbolB string) (decimal.Decimal, error) {
	return decimal.Zero, fmt.Errorf("MOCK") // TODO: make special type
}

type PriceGetterNoResp struct{}

var _ ports.PriceGetter = (*PriceGetterNoResp)(nil)

func (pg *PriceGetterNoResp) GetPrice(ctx context.Context, symbolA, symbolB string) (decimal.Decimal, error) {
	<-ctx.Done()
	return decimal.Zero, context.Canceled
}
