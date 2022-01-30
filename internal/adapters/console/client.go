package console

import (
	"context"
	"converter/internal/core/ports"
	"fmt"
	"os"
	"sync"

	"github.com/shopspring/decimal"
)

type ConsoleClient struct {
	ctx       context.Context
	cancel    context.CancelFunc
	converter ports.CoinConverter
}

func NewConsoleClient(converter ports.CoinConverter) *ConsoleClient {
	c := &ConsoleClient{
		converter: converter,
	}
	return c
}

func (c *ConsoleClient) Run(ctx context.Context, wg *sync.WaitGroup, errChan chan<- error) error {

	c.ctx, c.cancel = context.WithCancel(ctx)

	wg.Add(1)
	go c.mainLoop(c.ctx, wg)

	return nil
}

func (c *ConsoleClient) Close() error {
	c.cancel()
	return nil
}

func (c *ConsoleClient) mainLoop(ctx context.Context, wg *sync.WaitGroup) {

	defer wg.Done()

	a := os.Args[len(os.Args)-3]
	s1 := os.Args[len(os.Args)-2]
	s2 := os.Args[len(os.Args)-1]

	amount, err := decimal.NewFromString(a)
	if err != nil {
		handleInputErr(err)
		return
	}

	result, err := c.converter.Convert(amount, s1, s2)
	if err != nil {
		handleConvertErr(err)
		return
	}

	fmt.Printf("%s %s -> %s %s\n", s1, amount, s2, result)
}

func handleInputErr(err error) {
	fmt.Printf("input err: %s", err)
}

func handleConvertErr(err error) {
	fmt.Printf("convert err: %s", err)
}
