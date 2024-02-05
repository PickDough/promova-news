package getAllPosts

type GetAllPostsQuery struct {
	IdOffset int `json:"idOffset"`
	Limit    int `json:"limit" validate:"max=50"`
}
