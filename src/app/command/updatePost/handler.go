package updatePost

import (
	"context"
	"promova-news/.gen/news/public/model"
)

type (
	PostsRepository interface {
		Update(ctx context.Context, post *model.Posts) (*model.Posts, error)
	}
	updatePostHandler struct {
		repo PostsRepository
	}
)

func New(repo PostsRepository) *updatePostHandler {
	return &updatePostHandler{
		repo: repo,
	}
}

func (h *updatePostHandler) Handle(ctx context.Context, command UpdatePostCommand) (*model.Posts, error) {
	return h.repo.Update(ctx, &model.Posts{
		ID:      int32(command.Id),
		Title:   command.Title,
		Content: command.Content,
	})
}
