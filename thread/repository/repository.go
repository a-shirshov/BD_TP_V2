package repository

import (
	"bd_tp_V2/models"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx"
)

type Repository struct {
	db *pgx.ConnPool
}

func NewThreadRepository(db *pgx.ConnPool) *Repository {
	return &Repository{
		db: db,
	}
}

func (tR *Repository) ThreadDetails(slug_or_id string) (*models.Thread, error) {
	var row *pgx.Row
	id, err := strconv.Atoi(slug_or_id)
	if err != nil {
		row = tR.db.QueryRow("FindThreadWithSlugQuery",slug_or_id)
	} else {
		row = tR.db.QueryRow("FindThreadWithIdQuery",id)
	}

	thread := &models.Thread{}
	err = row.Scan(&thread.ID,&thread.Title,&thread.Author,&thread.Forum,&thread.Message,&thread.Votes,&thread.Slug,&thread.Created)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorThreadNotFound
		} else {
			return nil, err
		}
	}
	return thread, nil
}

func (fR *Repository) CreatePostsNew(threadId int, forum string, posts *models.Posts) (*models.Posts, error) {
	newPosts := &models.Posts{}
	created := time.Now()
	for _, post := range posts.Posts {
		var author string
		err := fR.db.QueryRow("CheckNicknameQuery", post.Author).Scan(&author)
		if err == pgx.ErrNoRows {
			return nil, models.ErrorUserNotFound
		}

		if post.Parent != 0 {
			var id int
			err := fR.db.QueryRow("CheckThreadQuery", threadId, post.Parent).Scan(&id)
			if err == pgx.ErrNoRows {
				return nil, models.ErrorPostNotFound
			}
		}
		newPost := &models.Post{}

		err = fR.db.QueryRow("CreatePostQuery",post.Parent,post.Author,post.Message,forum,threadId,created).Scan(
			&newPost.ID,&newPost.Parent,&newPost.Author,&newPost.Message,&newPost.Edited,&newPost.Forum,&newPost.Thread,&newPost.Created)

		if err != nil {
			return nil, err
		}
		newPosts.Posts = append(newPosts.Posts, *newPost)
	}
	
	if len(newPosts.Posts) == 0 {
		return nil, models.ErrorNoPosts
	}
	return newPosts, nil
}

func (fR *Repository) ThreadVote(vote *models.Vote, threadId int) error {
	_, err := fR.db.Exec("InsertVoteQuery",threadId,vote.Nickname,vote.Voice)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			_, err := fR.db.Exec("UpdateVoteQuery",vote.Voice, threadId, vote.Nickname)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}

func (fR *Repository) ThreadGetPosts(id int, limit, since, sort, desc string) (*models.Posts, error) {
	if sort == "" || sort == "flat" {
		return fR.threadGetPostsFlat(id, limit, since, sort, desc)
	} else if sort == "tree" {
		return fR.threadGetPostsTree(id, limit, since, desc)
	} else if sort == "parent_tree" {
		return fR.threadGetPostsParentTree(id, limit, since, desc)
	}
	return nil, errors.New("недописал")
}

func (fR *Repository) threadGetPostsFlat(id int, limit, since, sort, desc string) (*models.Posts, error) {
	var rows *pgx.Rows
	var err error
	if desc == "true" {
		if since != "" {
			rows, err = fR.db.Query("GetPostsByThreadFlatDescSince", id, since, limit)
		} else {
			rows, err = fR.db.Query("GetPostsByThreadFlatDesc", id, limit)
		}
	} else {
		if since != "" {
			rows, err = fR.db.Query("GetPostsByThreadFlatAscSince", id, since, limit)
		} else {
			rows, err = fR.db.Query("GetPostsByThreadFlatAsc", id, limit)
		}
	}
	if err != nil {
		rows.Close()
		return nil, err
	}
	defer rows.Close()
	posts := &models.Posts{}
	for rows.Next() {
		post := &models.Post{}
		err := rows.Scan(&post.ID, &post.Parent, &post.Author, &post.Message, &post.Edited, &post.Forum, &post.Thread, &post.Created)
		if err != nil {
			return nil,err
		}
		posts.Posts = append(posts.Posts, *post)
	}
	return posts, nil
}

func (tR *Repository) threadGetPostsTree(id int, limit string, since string, desc string) (*models.Posts, error) {
	var err error
	var rows *pgx.Rows
	if desc == "true" {
		if since != "" {
			rows, err = tR.db.Query("GetPostsByThreadTreeDescSince", id, since, limit)
		} else {
			rows, err = tR.db.Query("GetPostsByThreadTreeDesc", id, limit)
		}
	} else {
		if since != "" {
			rows, err = tR.db.Query("GetPostsByThreadTreeAscSince", id, since, limit)
		} else {
			rows, err = tR.db.Query("GetPostsByThreadTreeAsc", id, limit)
		}
	}

	if err != nil {

		return nil, err
	}
	defer rows.Close()
	posts := &models.Posts{}
	for rows.Next() {
		post := &models.Post{}
		err := rows.Scan(&post.ID, &post.Parent, &post.Author, &post.Message, &post.Edited, &post.Forum, &post.Thread, &post.Created)
		if err != nil {
			return nil,err
		}
		posts.Posts = append(posts.Posts, *post)
	}
	return posts, nil
}

func (tR *Repository) threadGetPostsParentTree(id int, limit string, since string, desc string) (*models.Posts, error) {

	var rows *pgx.Rows
	var err error

	if desc == "true" {
		if since != "" {
			rows, err = tR.db.Query("QueryGetThreadPostsParentTreeDescSince", id, since, limit)
		} else {
			rows, err = tR.db.Query("QueryGetThreadPostsParentTreeDesc", id, limit)
		}
	} else {
		if since != "" {
			rows, err = tR.db.Query("QueryGetThreadPostsParentTreeAscSince", id, since, limit)
		} else {
			rows, err = tR.db.Query("QueryGetThreadPostsParentTreeAsc", id, limit)
		}
	}

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	posts := &models.Posts{}
	for rows.Next() {
		post := models.Post{}
		err := rows.Scan(&post.ID, &post.Parent, &post.Author, &post.Message, &post.Edited, &post.Forum, &post.Thread, &post.Created)
		if err != nil {
			return nil,err
		}
		posts.Posts = append(posts.Posts, post)
	}
	return posts, nil
}

func (tR *Repository) UpdateThreadDetails(thread *models.Thread) (*models.Thread, error) {
	row := tR.db.QueryRow("UpdateThreadDetailsQuery", thread.Title, thread.Message, thread.ID)
	err := row.Scan(&thread.ID,&thread.Title,&thread.Author,&thread.Forum,&thread.Message,&thread.Votes,&thread.Slug,&thread.Created)
	if err != nil {
		return nil, err
	}
	return thread, nil
}