package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/benknoble/game-roulette/app"
	"github.com/benknoble/game-roulette/data"
	"github.com/benknoble/game-roulette/web"
)

var (
	hostname string
)

func init() {
	go func() {
		//Capture program shutdown, to make sure
		//everything shuts down nicely
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		for sig := range c {
			if sig == os.Interrupt {
				app.Halt("Roulette Web Server %s shutting down", hostname)
			}
		}
	}()
}

func main() {
	var err error
	hostname, err = os.Hostname()
	if err != nil {
		app.Halt("Error getting hostname %s", err.Error())
	}

	// load configuration somehow
	// consider https://github.com/timshannon/config
	// or Viper https://github.com/spf13/viper

	webConf := web.DefaultConfig()
	appConf := app.DefaultConfig()
	dataConf := data.DefaultConfig()

	err = data.Init(dataConf)
	if err != nil {
		app.Halt("Error initializing data layer: %s", err.Error())
	}

	err = app.Init(appConf)
	if err != nil {
		app.Halt("Error initializing app layer: %s", err.Error())
	}

	err = web.StartServer(webConf)
	if err != nil {
		app.Halt("Error starting server %s", err.Error())
	}
}
