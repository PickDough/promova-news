package app

import (
	"promova-news/src/app/command/createPost"
	"promova-news/src/app/command/deletePost"
	"promova-news/src/app/command/updatePost"
	"promova-news/src/app/query/getAllPosts"
	"promova-news/src/app/query/getPost"
	"promova-news/src/config"
	"promova-news/src/network"
	"promova-news/src/network/postsApi/createPostApi"
	"promova-news/src/network/postsApi/deletePostApi"
	"promova-news/src/network/postsApi/getAllPostsApi"
	"promova-news/src/network/postsApi/getPostApi"
	"promova-news/src/network/postsApi/updatePostApi"
	"promova-news/src/persistance"

	"go.uber.org/fx"
)

var App = fx.Module("app",
	config.Module,
	persistance.Module,
	network.Module,
	fx.Provide(
		fx.Annotate(
			createPost.New,
			fx.As(new(createPostApi.CreatePostHandler)),
		),
		fx.Annotate(
			updatePost.New,
			fx.As(new(updatePostApi.UpdatePostHandler)),
		),
		fx.Annotate(
			deletePost.New,
			fx.As(new(deletePostApi.DeletePostHandler)),
		),
		fx.Annotate(
			getPost.New,
			fx.As(new(getPostApi.GetPostHandler)),
		),
		fx.Annotate(
			getAllPosts.New,
			fx.As(new(getAllPostsApi.GetAllPostsHandler)),
		),
	),
)
