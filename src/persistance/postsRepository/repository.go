package postsRepository

import (
	"context"
	"database/sql"
	"promova-news/.gen/news/public/model"
	. "promova-news/.gen/news/public/table"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type PostsRepository struct {
	db *sql.DB
}

func (d *PostsRepository) Insert(ctx context.Context, post *model.Posts) (*model.Posts, error) {
	stmt := Posts.INSERT(Posts.Title, Posts.Content).
		VALUES(post.Title, post.Content).
		RETURNING(Posts.AllColumns)

	dest := model.Posts{}
	err := stmt.Query(d.db, &dest)

	return &dest, err
}

func (d *PostsRepository) Update(ctx context.Context, post *model.Posts) (*model.Posts, error) {
	setFields := make([]any, 0)
	if post.Title != "" {
		setFields = append(setFields, Posts.Title.SET(String(post.Title)))
	}
	if post.Content != "" {
		setFields = append(setFields, Posts.Content.SET(String(post.Content)))
	}
	if len(setFields) > 0 {
		stmt := Posts.UPDATE().
			SET(Posts.UpdatedAt.SET(CAST(CURRENT_TIMESTAMP()).AS_TIMESTAMP()), setFields...).
			WHERE(Posts.ID.EQ(Int(int64(post.ID)))).
			RETURNING(Posts.AllColumns)

		dest := model.Posts{}
		err := stmt.Query(d.db, &dest)

		return &dest, err
	} else {
		return d.Get(ctx, int(post.ID))
	}
}

func (d *PostsRepository) Get(ctx context.Context, id int) (*model.Posts, error) {
	stmt := Posts.SELECT(Posts.AllColumns).WHERE(Posts.ID.EQ(Int(int64(id)))).LIMIT(1)

	dest := model.Posts{}
	err := stmt.Query(d.db, &dest)
	if err == qrm.ErrNoRows {
		return nil, nil
	}

	return &dest, err
}

func (d *PostsRepository) GetAll(ctx context.Context, idOffset int, limit int) ([]*model.Posts, error) {
	stmt := Posts.SELECT(Posts.AllColumns).
		WHERE(Posts.ID.GT(Int(int64(idOffset)))).
		LIMIT(int64(limit))

	dest := []*model.Posts{}
	err := stmt.Query(d.db, &dest)

	return dest, err
}

func (d *PostsRepository) Delete(ctx context.Context, id int) error {
	stmt := Posts.DELETE().WHERE(Posts.ID.EQ(Int(int64(id))))

	_, err := stmt.Exec(d.db)
	if err != nil {
		return err
	}

	return nil
}

func New(db *sql.DB) *PostsRepository {
	return &PostsRepository{
		db: db,
	}
}
