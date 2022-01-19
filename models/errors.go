package models

import "errors"

var (
	ErrorUserExists = errors.New("user already exists")
	ErrorUserNotFound = errors.New("user not found")
	ErrorUserUpdateConflict = errors.New("can't update user cause of conflict")

	ErrorForumNotFound = errors.New("forum not found")
	ErrorForumExists = errors.New("forum already exists")

	ErrorThreadNotFound = errors.New("thread not found")
	ErrorThreadExists = errors.New("thread already exists")

	ErrorPostNotFound = errors.New("post not found")
	ErrorNoPosts = errors.New("no posts")
)