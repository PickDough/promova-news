package createPostApi

import (
	"context"
	"net/http"
	"promova-news/.gen/news/public/model"
	"promova-news/src/app/command/createPost"

	"go.uber.org/fx"

	"github.com/labstack/echo/v4"
)

type (
	CreatePostHandler interface {
		Handle(ctx context.Context, command createPost.CreatePostCommand) (*model.Posts, error)
	}
	CreatePostApi struct {
		handler CreatePostHandler
	}
)

func New(handler CreatePostHandler) *CreatePostApi {
	return &CreatePostApi{
		handler: handler,
	}
}

func (req *CreatePostApi) CreatePost(c echo.Context) error {
	command := createPost.CreatePostCommand{}
	if err := c.Bind(&command); err != nil {
		return err
	}
	if err := c.Validate(command); err != nil {
		return err
	}

	res, err := req.handler.Handle(c.Request().Context(), command)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, res)
}

var Module = fx.Module("createPostApi", fx.Provide(New))
