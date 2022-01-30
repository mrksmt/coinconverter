package application

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"reflect"
	"sync"
	"syscall"
	"time"

	"github.com/pkg/errors"
)

type Application struct {
	name              string
	version           string
	rootContextCancel context.CancelFunc
	waitGroup         *sync.WaitGroup
	errChan           chan error
	components        []RunnableComponent
}

func New(appName string, appVersion string) *Application {

	s := &Application{
		name:    appName,
		version: appVersion,
	}

	return s
}

func (s *Application) Add(components ...RunnableComponent) *Application {
	s.components = append(s.components, components...)
	return s
}

func (s *Application) Run() error {

	ctx, cancel := context.WithCancel(context.Background())
	wg := new(sync.WaitGroup)

	s.rootContextCancel = cancel
	s.waitGroup = wg
	s.errChan = make(chan error, 42)

	for _, component := range s.components {
		if err := component.Run(ctx, wg, s.errChan); err != nil {
			s.rootContextCancel()
			componentType := reflect.TypeOf(component)
			msg := fmt.Sprintf("Run %s err", componentType)
			return errors.Wrap(err, msg)
		}
	}

	// log.Printf("Application %s %s RUN", s.name, s.version)
	return nil
}

func (s *Application) Wait() {

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	allComponentsDone := make(chan struct{})
	go func() {
		s.waitGroup.Wait()
		allComponentsDone <- struct{}{}
	}()

	select {
	case <-quit:
	case <-allComponentsDone:
	case err := <-s.errChan:
		log.Printf("component err: %s", err)
	}
}

func (s *Application) Stop(timeout time.Duration) {

	if waitTimeout(s.waitGroup, timeout) {
		log.Println("Timed out waiting for main wait group")
	} else {
		_ = 1
		// log.Println("Main wait group finished")
	}

	// log.Printf("Application %s %s DONE", s.name, s.version)
}

// waitTimeout waits for the waitgroup for the specified max timeout.
// Returns true if waiting timed out.
func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {

	c := make(chan struct{})

	go func() {
		defer close(c)
		wg.Wait()
	}()

	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
