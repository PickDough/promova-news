package persistance

import (
	"database/sql"

	"promova-news/src/app/command/createPost"
	"promova-news/src/app/command/deletePost"
	"promova-news/src/app/command/updatePost"
	"promova-news/src/app/query/getAllPosts"
	"promova-news/src/app/query/getPost"
	"promova-news/src/config"
	"promova-news/src/persistance/postsRepository"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"go.uber.org/fx"
)

func NewDatabase(config *config.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", config.Db.Url)
	if err != nil {
		return nil, err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres", driver)
	if err != nil {
		return nil, err
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		return nil, err
	}

	return db, nil
}

var Module = fx.Module("persistance",
	fx.Provide(
		NewDatabase,
		fx.Annotate(
			postsRepository.New,
			fx.As(new(createPost.PostsRepository)),
		),
		fx.Annotate(
			postsRepository.New,
			fx.As(new(deletePost.PostsRepository)),
		),
		fx.Annotate(
			postsRepository.New,
			fx.As(new(updatePost.PostsRepository)),
		),
		fx.Annotate(
			postsRepository.New,
			fx.As(new(getAllPosts.PostsRepository)),
		),
		fx.Annotate(
			postsRepository.New,
			fx.As(new(getPost.PostsRepository)),
		),
	),
)
