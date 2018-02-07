package web

import (
	"fmt"
	"net/http"
	// "github.com/benknoble/game-roulette/app"
	// "github.com/benknoble/game-roulette/data"
)

type Config struct{}

func DefaultConfig() *Config {
	return &Config{}
}

func StartServer(c *Config) error {
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte(fmt.Sprintf("Hello, World!\nThe roulette server is still under construction; please check back later!")))
		return
	})
	err := http.ListenAndServe(":8080", nil)
	return err
}
