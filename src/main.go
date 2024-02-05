package main

import (
	"net/http"
	"promova-news/src/app"

	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		app.App,
		fx.Invoke(func(s *http.Server) {}),
	)
	app.Run()
}
