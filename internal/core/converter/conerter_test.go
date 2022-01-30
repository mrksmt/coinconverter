package converter

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestCalcResult(t *testing.T) {

	res, err := calculateResult(decimal.NewFromFloat(2.0), decimal.NewFromFloat(2.5))

	if err != nil {
		t.Errorf("not nil err: %s", err)
	}

	if res.Cmp(decimal.NewFromFloat(5.0)) != 0 {
		t.Errorf("wrong result: %s", res)
	}
}
