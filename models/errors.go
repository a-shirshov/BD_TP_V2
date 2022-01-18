package models

import "errors"

var (
	ErrorUserExists = errors.New("user already exists")
	ErrorUserNotFound = errors.New("user not found")
	ErrorUserUpdateConflict = errors.New("can't update user cause of conflict")
)