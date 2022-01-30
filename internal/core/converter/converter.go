package converter

import (
	"context"
	"converter/internal/core/ports"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

const TIME_OUT = time.Second

type CoinConverterV1 struct {
	ctx         context.Context
	cancel      context.CancelFunc
	priceGetter ports.PriceGetter
}

var _ ports.CoinConverter = (*CoinConverterV1)(nil)

func NewCoinConverterV1(priceGetter ports.PriceGetter) *CoinConverterV1 {
	c := &CoinConverterV1{
		priceGetter: priceGetter,
	}
	return c
}

func (c *CoinConverterV1) Run(ctx context.Context, wg *sync.WaitGroup, errChan chan<- error) error {
	c.ctx, c.cancel = context.WithCancel(ctx)
	return nil
}

func (c *CoinConverterV1) Close() error {
	c.cancel()
	return nil
}

func (c *CoinConverterV1) Convert(amount decimal.Decimal, symbolA, symbolB string) (decimal.Decimal, error) {

	// check input data before processing
	err := checkParams(amount, symbolA, symbolB)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "wrong param")
	}

	// request additional data
	ctx, cancel := context.WithTimeout(c.ctx, TIME_OUT)
	defer cancel()
	price, err := c.priceGetter.GetPrice(ctx, symbolA, symbolB)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "get price")
	}

	// main logic
	result, err := calculateResult(amount, price)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "get result")
	}

	return result, nil
}

func checkParams(amount decimal.Decimal, symbolA, symbolB string) error {
	// TODO: implement me
	return nil
}

func calculateResult(amount, price decimal.Decimal) (decimal.Decimal, error) {
	return amount.Mul(price), nil
}
