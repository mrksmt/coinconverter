package coinmarket

import (
	"context"
	"log"
	"testing"
)

func TestGetPrice(t *testing.T) {

	log.SetFlags(log.Lshortfile)

	cma := NewCoinMarketAdapter()
	price, err := cma.GetPrice(context.TODO(), "USD", "BTC")
	if err != nil {
		t.Error(err)
	}

	if price.Sign() <= 0 {
		t.Errorf("wrong price: %s", price)
	}

	// log.Println(price)
	// t.Error("MOCK")
}
