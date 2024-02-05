package getPost

import (
	"context"
	"testing"

	"promova-news/.gen/news/public/model"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type getPostHandlerSuite struct {
	suite.Suite
}

type postsRepository struct {
	mock.Mock
}

func (r *postsRepository) Get(ctx context.Context, id int) (*model.Posts, error) {
	args := r.Called(ctx, id)
	return args.Get(0).(*model.Posts), args.Error(1)
}

func (suite *getPostHandlerSuite) TestHandle() {
	repo := new(postsRepository)
	handler := New(repo)

	repo.On("Get", context.Background(), 1).Return(&model.Posts{ID: 1, Title: "Title", Content: "Content"}, nil)

	post, err := handler.Handle(context.Background(), GetPostQuery{Id: 1})

	suite.NoError(err)
	suite.EqualValues(1, post.ID)
	suite.Equal("Title", post.Title)
	suite.Equal("Content", post.Content)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(getPostHandlerSuite))
}
