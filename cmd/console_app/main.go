package main

import (
	"converter/internal/adapters/coinmarket"
	"converter/internal/adapters/console"
	"converter/internal/core/converter"
	"converter/pkg/application"
	"log"
	"time"

	"github.com/pkg/errors"
)

var AppName = "Coinconv"
var AppVersion = "developer"

func main() {

	// configure your logger
	log.SetFlags(log.Lshortfile)

	app := application.New(AppName, AppVersion)
	makeApplicationComponents(app)

	// run the service
	if err := app.Run(); err != nil {
		app.Stop(time.Second * 2)
		log.Fatal(errors.Wrap(err, "app.Run"))
	}

	// waiting for shutdown signal
	app.Wait()

	// shutdown all running components
	app.Stop(time.Second * 5)
}

// make all app components and add some of them to app.Run
func makeApplicationComponents(app *application.Application) {

	priceGetter := coinmarket.NewCoinMarketAdapter()
	_ = priceGetter

	converter := converter.NewCoinConverterV1(priceGetter)
	app.Add(converter)

	client := console.NewConsoleClient(converter)
	app.Add(client)
}
