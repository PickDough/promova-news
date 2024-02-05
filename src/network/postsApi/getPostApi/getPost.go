package getPostApi

import (
	"context"
	"errors"
	"net/http"
	"promova-news/.gen/news/public/model"
	"promova-news/src/app/query/getPost"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type (
	GetPostHandler interface {
		Handle(ctx context.Context, command getPost.GetPostQuery) (*model.Posts, error)
	}
	GetPostApi struct {
		handler GetPostHandler
	}
)

func New(handler GetPostHandler) *GetPostApi {
	return &GetPostApi{
		handler: handler,
	}
}

// @Summary Get post
// @ID get-post
// @Accept  */*
// @Produce  json
// @Param id path int true "Post ID"
// @Success 200 {object} model.Posts
// @Failure 400 {object} echo.HTTPError
// @Failure 404 {object} echo.HTTPError
// @Router /posts/:id  [get]
func (req *GetPostApi) GetPost(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return errors.New("id should be an int")
	}
	res, err := req.handler.Handle(c.Request().Context(), getPost.GetPostQuery{
		Id: int(id),
	})
	if err != nil {
		return err
	}
	if res == nil {
		return c.JSON(http.StatusNotFound, echo.ErrNotFound)
	}

	return c.JSON(http.StatusOK, res)
}

var Module = fx.Module("getPostApi", fx.Provide(New))
