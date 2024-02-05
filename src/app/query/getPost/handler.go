package getPost

import (
	"context"
	"promova-news/.gen/news/public/model"
)

type (
	PostsRepository interface {
		Get(ctx context.Context, id int) (*model.Posts, error)
	}
	getPostHandler struct {
		repo PostsRepository
	}
)

func New(repo PostsRepository) *getPostHandler {
	return &getPostHandler{
		repo: repo,
	}
}

func (h *getPostHandler) Handle(ctx context.Context, query GetPostQuery) (*model.Posts, error) {
	return h.repo.Get(ctx, query.Id)
}
