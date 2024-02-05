package updatePost

type UpdatePostCommand struct {
	Id      int    `swaggerignore:"true"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
