package coinmarket

import (
	"context"
	"converter/internal/core/ports"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

const (
	TEST_API_KEY = "b54bcf4d-1bca-4e8e-9a24-22ff2c3d462c"
)

type CoinMarketAdapter struct {
	client *http.Client
}

var _ ports.PriceGetter = (*CoinMarketAdapter)(nil)

func NewCoinMarketAdapter() *CoinMarketAdapter {
	cma := &CoinMarketAdapter{
		client: &http.Client{},
	}
	return cma
}

func (cma *CoinMarketAdapter) GetPrice(ctx context.Context, symbolA, symbolB string) (decimal.Decimal, error) {

	req, err := http.NewRequest("GET", "https://sandbox-api.coinmarketcap.com/v1/tools/price-conversion", nil)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "make http request")
	}
	req = req.WithContext(ctx)

	q := url.Values{}
	q.Add("amount", "1")
	q.Add("symbol", symbolA)
	q.Add("convert", symbolB)

	req.Header.Set("Accepts", "application/json")
	req.Header.Add("X-CMC_PRO_API_KEY", TEST_API_KEY)
	req.URL.RawQuery = q.Encode()

	resp, err := cma.client.Do(req)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "sending request to server")
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "read http res body")
	}

	res := PriceConversionResponseSanBox{}
	err = json.Unmarshal(respBody, &res)
	if err != nil {
		return decimal.Zero, errors.Wrap(err, "resp unmarshal")
	}

	if resp.StatusCode != http.StatusOK {
		return decimal.Zero, fmt.Errorf("%s", res.Status.ErrorMessage)
	}

	symbolAData, exist := res.Data[symbolA]
	if !exist {
		return decimal.Zero, errors.Wrap(err, "wrong response, data not exist")
	}

	symbolBData, exist := symbolAData.Quote[symbolB]
	if !exist {
		return decimal.Zero, errors.Wrap(err, "wrong response, price not exist")
	}

	return symbolBData.Price, nil
}
