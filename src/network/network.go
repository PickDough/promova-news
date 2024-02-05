package network

import (
	"context"
	"fmt"
	"net/http"
	"promova-news/src/config"
	"promova-news/src/network/postsApi"

	"github.com/go-playground/validator"
	"go.uber.org/fx"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Validation error: %s", err.Error()))
	}
	return nil
}

func NewServer(lc fx.Lifecycle, config *config.Config, postsApi *postsApi.PostsApi) *http.Server {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Validator = &Validator{validator: validator.New()}

	e.POST("/posts", postsApi.CreatePost)
	e.PUT("/posts/:id", postsApi.UpdatePost)
	e.DELETE("/posts/:id", postsApi.DeletePost)
	e.GET("/posts/:id", postsApi.GetPost)
	e.GET("/posts", postsApi.GetAllPosts)

	s := &http.Server{
		Addr:    config.Http.Port,
		Handler: e,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go s.ListenAndServe()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return s.Shutdown(ctx)
		},
	})

	return s
}

var Module = fx.Module(
	"api",
	postsApi.Module,
	fx.Provide(
		NewServer,
	),
)
