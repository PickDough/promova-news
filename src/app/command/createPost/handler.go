package createPost

import (
	"context"
	"promova-news/.gen/news/public/model"
)

type (
	PostsRepository interface {
		Insert(ctx context.Context, post *model.Posts) (*model.Posts, error)
	}
	createPostHandler struct {
		repo PostsRepository
	}
)

func New(repo PostsRepository) *createPostHandler {
	return &createPostHandler{
		repo: repo,
	}
}

func (c *createPostHandler) Handle(ctx context.Context, command CreatePostCommand) (*model.Posts, error) {
	return c.repo.Insert(
		ctx,
		&model.Posts{Title: command.Title, Content: command.Content},
	)
}
