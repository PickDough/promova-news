package deletePostApi

import (
	"context"
	"errors"
	"net/http"
	"promova-news/src/app/command/deletePost"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type (
	DeletePostHandler interface {
		Handle(ctx context.Context, command deletePost.DeletePostCommand) error
	}
	DeletePostApi struct {
		handler DeletePostHandler
	}
)

func New(handler DeletePostHandler) *DeletePostApi {
	return &DeletePostApi{
		handler: handler,
	}
}

// @Summary Delete post
// @ID delete-post
// @Accept  */*
// @Param id path int true "Post ID"
// @Success 200
// @Router /posts/:id [delete]
func (req *DeletePostApi) DeletePost(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return errors.New("id should be an int")
	}
	err = req.handler.Handle(c.Request().Context(), deletePost.DeletePostCommand{
		Id: id,
	})
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

var Module = fx.Module("deletePostApi", fx.Provide(New))
