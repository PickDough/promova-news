package postsApi

import (
	"promova-news/src/network/postsApi/createPostApi"
	"promova-news/src/network/postsApi/deletePostApi"
	"promova-news/src/network/postsApi/getAllPostsApi"
	"promova-news/src/network/postsApi/getPostApi"
	"promova-news/src/network/postsApi/updatePostApi"

	"go.uber.org/fx"
)

type PostsApi struct {
	*createPostApi.CreatePostApi
	*updatePostApi.UpdatePostApi
	*deletePostApi.DeletePostApi
	*getPostApi.GetPostApi
	*getAllPostsApi.GetAllPostsApi
}

func New(
	create *createPostApi.CreatePostApi,
	update *updatePostApi.UpdatePostApi,
	delete *deletePostApi.DeletePostApi,
	get *getPostApi.GetPostApi,
	getAll *getAllPostsApi.GetAllPostsApi,
) *PostsApi {
	return &PostsApi{
		create,
		update,
		delete,
		get,
		getAll,
	}
}

var Module = fx.Module("postsApi",
	createPostApi.Module,
	updatePostApi.Module,
	deletePostApi.Module,
	getPostApi.Module,
	getAllPostsApi.Module,
	fx.Provide(
		New,
	),
)
