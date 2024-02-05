package getAllPosts

import (
	"context"
	"promova-news/.gen/news/public/model"
)

type (
	PostsRepository interface {
		GetAll(ctx context.Context, idOffset int, limit int) ([]*model.Posts, error)
	}
	getAllPostsHandler struct {
		repo PostsRepository
	}
)

func New(repo PostsRepository) *getAllPostsHandler {
	return &getAllPostsHandler{
		repo: repo,
	}
}

func (h *getAllPostsHandler) Handle(ctx context.Context, query GetAllPostsQuery) ([]*model.Posts, error) {
	if query.Limit == 0 {
		query.Limit = 50
	}
	return h.repo.GetAll(ctx, query.IdOffset, query.Limit)
}
