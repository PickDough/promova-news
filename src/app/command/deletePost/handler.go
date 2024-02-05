package deletePost

import "context"

type (
	PostsRepository interface {
		Delete(ctx context.Context, id int) error
	}

	deletePostHandler struct {
		repo PostsRepository
	}
)

func New(repo PostsRepository) *deletePostHandler {
	return &deletePostHandler{
		repo: repo,
	}
}

func (h *deletePostHandler) Handle(ctx context.Context, command DeletePostCommand) error {
	return h.repo.Delete(ctx, command.Id)
}
