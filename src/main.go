package main

import (
	"net/http"
	"promova-news/src/app"

	"go.uber.org/fx"

	_ "github.com/swaggo/echo-swagger/example/docs"
)

// @title Promova News API
// @version 1.0
// @description This is a test case for Promova.

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

func main() {
	app := fx.New(
		app.App,
		fx.Invoke(func(s *http.Server) {}),
	)
	app.Run()
}
