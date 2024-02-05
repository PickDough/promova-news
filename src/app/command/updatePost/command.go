package updatePost

type UpdatePostCommand struct {
	Id      int
	Title   string `json:"title"`
	Content string `json:"content"`
}
