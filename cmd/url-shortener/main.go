package main

import (
	"fmt"

	"github.com/israpilovsha/url-shortener/internal/config"
)

func main() {
	// TODO: init config:
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: init logger: slog import "log/slog"

	// TODO: init storage: sqlite

	// TODO: init router: chi, "chi render"

	// TODO: run server:
}
