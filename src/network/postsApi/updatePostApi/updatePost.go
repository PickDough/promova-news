package updatePostApi

import (
	"context"
	"errors"
	"net/http"
	"promova-news/.gen/news/public/model"
	"promova-news/src/app/command/updatePost"
	"strconv"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

type (
	UpdatePostHandler interface {
		Handle(ctx context.Context, command updatePost.UpdatePostCommand) (*model.Posts, error)
	}
	UpdatePostApi struct {
		handler UpdatePostHandler
	}
)

func New(handler UpdatePostHandler) *UpdatePostApi {
	return &UpdatePostApi{
		handler: handler,
	}
}

func (req *UpdatePostApi) UpdatePost(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		return errors.New("id should be an int")
	}
	command := updatePost.UpdatePostCommand{Id: int(id)}
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

var Module = fx.Module("updatePostApi", fx.Provide(New))
