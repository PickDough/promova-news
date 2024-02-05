package postsRepository

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	"promova-news/.gen/news/public/model"

	. "promova-news/.gen/news/public/table"

	. "github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"

	"github.com/stretchr/testify/suite"

	"context"

	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

type postsRepositoryTestSuite struct {
	suite.Suite
	repo        *PostsRepository
	ctx         context.Context
	pgContainer *postgres.PostgresContainer
}

func (suite *postsRepositoryTestSuite) SetupSuite() {
	suite.ctx = context.Background()

	dbName := "test"
	dbUser := "test"
	dbPassword := "test"

	pgContainer, err := postgres.RunContainer(suite.ctx,
		testcontainers.WithImage("docker.io/postgres:15.2-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithInitScripts("init_test.sql"),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPassword),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(5*time.Second),
		),
	)
	if err != nil {
		log.Fatal(err)
	}
	suite.pgContainer = pgContainer

	conn, _ := pgContainer.ConnectionString(suite.ctx)
	fmt.Printf("CONNECTION: %s\n", conn)
	db, err := sql.Open("postgres", fmt.Sprintf("%ssslmode=disable", conn))
	if err != nil {
		log.Fatal(err)
	}
	suite.repo = New(db)
}

func (suite *postsRepositoryTestSuite) TestInsert() {
	ctx := context.Background()

	post := model.Posts{
		Title:   "testing",
		Content: "tests",
	}
	res, err := suite.repo.Insert(ctx, &post)
	suite.Nil(err)

	stmt := Posts.SELECT(Posts.AllColumns).WHERE(Posts.Title.EQ(String("testing"))).LIMIT(1)
	postFromDb := model.Posts{}
	err = stmt.Query(suite.repo.db, &postFromDb)
	suite.Nil(err)

	suite.Equal(&postFromDb, res)
}

func (suite *postsRepositoryTestSuite) TestUpdate() {
	ctx := context.Background()

	post := model.Posts{
		ID:      9,
		Title:   "was",
		Content: "changed",
	}
	res, err := suite.repo.Update(ctx, &post)
	suite.Nil(err)

	stmt := Posts.SELECT(Posts.AllColumns).WHERE(Posts.ID.EQ(Int(9))).LIMIT(1)
	postFromDb := model.Posts{}
	err = stmt.Query(suite.repo.db, &postFromDb)
	suite.Nil(err)

	suite.Equal(&postFromDb, res)
}

func (suite *postsRepositoryTestSuite) TestGet() {
	ctx := context.Background()

	res, err := suite.repo.Get(ctx, 9)
	suite.Nil(err)

	stmt := Posts.SELECT(Posts.AllColumns).WHERE(Posts.ID.EQ(Int(9))).LIMIT(1)
	postFromDb := model.Posts{}
	err = stmt.Query(suite.repo.db, &postFromDb)
	suite.Nil(err)

	suite.Equal(&postFromDb, res)
}

func (suite *postsRepositoryTestSuite) TestGetNil() {
	ctx := context.Background()

	res, err := suite.repo.Get(ctx, 99)
	suite.Nil(err)
	suite.Empty(res)
}

func (suite *postsRepositoryTestSuite) TestGetAll() {
	ctx := context.Background()
	res, err := suite.repo.GetAll(ctx, 9, 10)
	suite.Nil(err)

	expected := []*model.Posts{
		{
			ID:      10,
			Title:   "already",
			Content: "present",
		},
		{
			ID:      11,
			Title:   "also",
			Content: "present",
		},
	}

	suite.Equal(len(expected), len(res))
	for i := range expected {
		suite.Equal(expected[i].ID, res[i].ID)
		suite.Equal(expected[i].Title, res[i].Title)
		suite.Equal(expected[i].Content, res[i].Content)
	}
}

func (suite *postsRepositoryTestSuite) TestGetAllPagination() {
	ctx := context.Background()
	res, err := suite.repo.GetAll(ctx, 10, 10)
	suite.Nil(err)

	expected := []*model.Posts{
		{
			ID:      11,
			Title:   "also",
			Content: "present",
		},
	}

	suite.Equal(len(expected), len(res))
	for i := range expected {
		suite.Equal(expected[i].ID, res[i].ID)
		suite.Equal(expected[i].Title, res[i].Title)
		suite.Equal(expected[i].Content, res[i].Content)
	}
}

func (suite *postsRepositoryTestSuite) TestGetAllEmpty() {
	ctx := context.Background()
	res, err := suite.repo.GetAll(ctx, 99, 10)
	suite.Nil(err)
	suite.Empty(res)
}

func (suite *postsRepositoryTestSuite) TestDelete() {
	ctx := context.Background()
	err := suite.repo.Delete(ctx, 8)
	suite.Nil(err)

	stmt := Posts.SELECT(Posts.AllColumns).WHERE(Posts.ID.EQ(Int(8))).LIMIT(1)
	postFromDb := model.Posts{}
	err = stmt.Query(suite.repo.db, &postFromDb)

	suite.Equal(err, qrm.ErrNoRows)
}

func (suite *postsRepositoryTestSuite) TearDownSuite() {
	if err := suite.pgContainer.Terminate(suite.ctx); err != nil {
		log.Fatalf("failed to terminate container: %s", err)
	}
}

func TestPostsRepositoryTestSuite(t *testing.T) {
	s := new(postsRepositoryTestSuite)
	suite.Run(t, s)
}
