package getAllPostsApi

import (
	"context"
	"net/http"
	"promova-news/.gen/news/public/model"
	"promova-news/src/app/query/getAllPosts"

	"go.uber.org/fx"

	"github.com/labstack/echo/v4"
)

type (
	GetAllPostsHandler interface {
		Handle(ctx context.Context, command getAllPosts.GetAllPostsQuery) ([]*model.Posts, error)
	}
	GetAllPostsApi struct {
		handler GetAllPostsHandler
	}
)

func New(handler GetAllPostsHandler) *GetAllPostsApi {
	return &GetAllPostsApi{
		handler: handler,
	}
}

func (req *GetAllPostsApi) GetAllPosts(c echo.Context) error {
	command := getAllPosts.GetAllPostsQuery{}
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

var Module = fx.Module("getAllPostApi", fx.Provide(New))
