package getPostApi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"promova-news/.gen/news/public/model"
	"promova-news/src/app/query/getPost"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type mockGetPostHandler struct {
	mock.Mock
}

func (m *mockGetPostHandler) Handle(ctx context.Context, command getPost.GetPostQuery) (*model.Posts, error) {
	args := m.Called(ctx, command)
	if args.Get(0) != nil {
		return args.Get(0).(*model.Posts), args.Error(1)
	}

	return nil, args.Error(1)
}

type getPostApiSuite struct {
	suite.Suite
	handler *mockGetPostHandler
	api     *GetPostApi
}

func (s *getPostApiSuite) SetupTest() {
	s.handler = new(mockGetPostHandler)
	s.api = New(s.handler)
}

func (suite *getPostApiSuite) TestGetPostSuccess() {
	postID := "123"
	expectedPost := &model.Posts{
		ID:      123,
		Title:   "Expected",
		Content: "Content",
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath("/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues(postID)
	suite.handler.On("Handle", c.Request().Context(), getPost.GetPostQuery{Id: 123}).Return(expectedPost, nil)
	err := suite.api.GetPost(c)

	suite.NoError(err)
	suite.Equal(http.StatusOK, rec.Code)

	returnedPost := &model.Posts{}
	err = json.Unmarshal(rec.Body.Bytes(), returnedPost)
	suite.NoError(err)
	suite.Equal(expectedPost, returnedPost)
}

func (suite *getPostApiSuite) TestGetPostNil() {
	postID := "123"

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath("/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues(postID)
	suite.handler.On("Handle", c.Request().Context(), getPost.GetPostQuery{Id: 123}).Return(nil, nil)
	err := suite.api.GetPost(c)

	suite.NoError(err)
	suite.Equal(http.StatusNotFound, rec.Code)
}

func (suite *getPostApiSuite) TestGetPostIdNotNumber() {
	postID := "xyz"

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	c.SetPath("/posts/:id")
	c.SetParamNames("id")
	c.SetParamValues(postID)
	err := suite.api.GetPost(c)

	suite.Error(err, errors.New("id should be an int"))
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(getPostApiSuite))
}
