package repository

import (
	"bd_tp_V2/models"
	"github.com/jackc/pgx"
	
)

type Repository struct {
	db *pgx.ConnPool
}

func NewForumRepository(db *pgx.ConnPool) *Repository {
	return &Repository{
		db: db,
	}
}

func (fR *Repository) ForumDetails(slug string) (*models.Forum, error) {
	forum := &models.Forum{}
	err := fR.db.QueryRow("ForumDetailsQuery", slug).Scan(&forum.Title, &forum.User, &forum.Slug, &forum.Posts, &forum.Threads)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorForumNotFound
		} else {
			return nil, err
		}
	}
	return forum, nil
}


func (fR *Repository) CreateForum(forum *models.Forum) (*models.Forum, error) {
	err := fR.db.QueryRow("CreateForumQuery",forum.Title,forum.User,forum.Slug).Scan(
		&forum.Title,&forum.User,&forum.Slug,&forum.Posts,&forum.Threads)
	if err != nil {
		return nil, err
	}
	return forum, nil
}

func (fR *Repository) ForumThreadCreate (thread *models.Thread) (*models.Thread, error) {
	err := fR.db.QueryRow("CreateForumThreadQuery",thread.Title,thread.Author,thread.Forum,thread.Message,thread.Slug,thread.Created).Scan(
		&thread.ID, &thread.Title, &thread.Author, &thread.Forum, &thread.Message, &thread.Votes, &thread.Slug, &thread.Created)
	if err != nil {
		return nil, err
	}
	return thread, nil
}

func (fR *Repository) GetThreadsByForum(slug, limit, since, desc string) (*models.Threads, error) {
	var rows *pgx.Rows
	var err error
	if desc == "true" {
		if since == "" {
			rows, err = fR.db.Query("GetForumThreadsDescQuery", slug, limit)
		} else {
			rows, err = fR.db.Query("GetForumThreadsDescSinceQuery", slug, since, limit)
		}
	} else {
		if since == "" {
			rows, err = fR.db.Query("GetForumThreadsAscQuery", slug, limit)
		} else {
			rows, err = fR.db.Query("GetForumThreadsAscSinceQuery", slug, since, limit)
		}
	}
	threads := &models.Threads{}

	if err != nil {

		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		thread := &models.Thread{}
		err = rows.Scan(&thread.ID, &thread.Title, &thread.Author, &thread.Forum, &thread.Message, &thread.Votes, &thread.Slug, &thread.Created)
		if err != nil {
			rows.Close()
			return nil, err
		}
		threads.Threads = append(threads.Threads, *thread)
	}
	rows.Close()
	return threads, nil
}

func (fR *Repository) GetForumUsers(slug string, limit, since, desc string) (*models.Users, error) {
	var rows *pgx.Rows
	var err error

	if desc == "true" {
		if since == "" {
			rows, err = fR.db.Query("GetUsersByForumDesc", slug, limit)
		} else {
			rows, err = fR.db.Query("GetUsersByForumDescSince", slug, since, limit)
		}
	} else {
		if since == "" {
			rows, err = fR.db.Query("GetUsersByForumAsc", slug, limit)
		} else {
			rows, err = fR.db.Query("GetUsersByForumAscSince", slug, since, limit)
		}
	}
	users := &models.Users{}

	if err != nil {
		rows.Close()
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.Nickname, &user.Fullname, &user.About, &user.Email)
		if err != nil {
			return nil, err
		}
		users.Users = append(users.Users, *user)
	}
	
	return users, nil
}

