package main

import (
	"log"

	"github.com/benknoble/game-roulette/app"
	"github.com/benknoble/game-roulette/data"
	"github.com/benknoble/game-roulette/web"
)

func main() {
	// load configuration somehow
	// consider https://github.com/timshannon/config
	// or Viper https://github.com/spf13/viper

	webConf := web.DefaultConfig()
	appConf := app.DefaultConfig()
	dataConf := data.DefaultConfig()

	err := data.Init(dataConf)
	if err != nil {
		log.Fatalf("Error initializing data layer: %s", err.Error())
	}

	err = app.Init(appConf)
	if err != nil {
		log.Fatalf("Error initializing app layer: %s", err.Error())
	}

	err = web.StartServer(webConf)
	if err != nil {
		app.Halt("Error starting server %s", err.Error())
	}
}
