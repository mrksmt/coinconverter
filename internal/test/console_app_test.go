package test

import (
	"context"
	"converter/internal/adapters/console"
	"converter/internal/core/converter"
	"converter/internal/test/mock"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestConsoleAppPositive(t *testing.T) {

	// input example
	// 123.45 USD BTC

	split := strings.Split("123.45 USD BTC", " ")

	ctx := context.Background()
	wg := new(sync.WaitGroup)

	os.Args = append(os.Args, split[0])
	os.Args = append(os.Args, split[1])
	os.Args = append(os.Args, split[2])

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	price, _ := decimal.NewFromString("0.00002644511108159157")
	priceGetter := &mock.PriceGetterFix{
		Price: price,
	}
	converter := converter.NewCoinConverterV1(priceGetter)
	_ = converter.Run(ctx, wg, nil)
	client := console.NewConsoleClient(converter)
	_ = client.Run(ctx, wg, nil)

	wg.Wait()

	w.Close()
	os.Stdout = old
	out, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}

	assert.Contains(t, string(out), "0.0032646489630224793165")
}
