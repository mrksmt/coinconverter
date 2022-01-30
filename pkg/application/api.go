package application

import (
	"context"
	"io"
	"sync"
)

type RunnableComponent interface {
	Run(ctx context.Context, wg *sync.WaitGroup, errChan chan<- error) error // for run a component
	io.Closer                                                                // for stop a component
}
