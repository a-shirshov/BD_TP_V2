package repository

import (
	"bd_tp_V2/models"

	"github.com/jackc/pgx"
	//"strings"
	//""
)

type Repository struct {
	db *pgx.ConnPool
}

func NewPostRepository(db *pgx.ConnPool) *Repository {
	return &Repository{
		db: db,
	}
}

func (pR *Repository) GetPost(id int) (*models.Post, error) {
	var post models.Post
	err := pR.db.QueryRow("FindPostByIDQuery", id).Scan(
		&post.ID, &post.Parent, &post.Author, &post.Message, &post.Edited, &post.Forum, &post.Thread, &post.Created)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorPostNotFound
		} else {
			return nil, err
		}
	}
	return &post, err
}

func (pR *Repository) UpdatePost(post *models.Post) (*models.Post, error) {
	var newPost models.Post
	err := pR.db.QueryRow("UpdatePostByIdQuery", post.Message, post.ID).Scan(
		&newPost.ID, &newPost.Parent, &newPost.Author, &newPost.Message, &newPost.Edited, &newPost.Forum, &newPost.Thread, &newPost.Created)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorPostNotFound
		} else {
			return nil, err
		}
	}
	return &newPost, nil
}


