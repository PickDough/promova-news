package createPost

type CreatePostCommand struct {
	Title   string `json:"title" validate:"required,min=3,max=100"`
	Content string `json:"content" validate:"min=15"`
}
