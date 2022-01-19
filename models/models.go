package models

import "time"

type Error struct {
	Message string `json:"message,omitempty"`
}

type User struct {
	ID int `json:"id,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Fullname string `json:"fullname,omitempty"`
	About string  `json:"about,omitempty"`
	Email string `json:"email,omitempty"`
}

type Forum struct {
	ID int `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	User string	`json:"user,omitempty"`
	Slug string `json:"slug,omitempty"`
	Posts int `json:"posts,omitempty"`
	Threads int `json:"threads,omitempty"`
}

type Thread struct {
	ID int `json:"id,omitempty"`
	Title string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Forum string `json:"forum,omitempty"`
	Message string `json:"message,omitempty"`
	Votes int `json:"votes,omitempty"`
	Slug string `json:"slug,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

type Threads struct {
	Threads []Thread `json:"threads,omitempty"`
}

type Post struct {
	ID int `json:"id,omitempty"`
	Parent int `json:"parent,omitempty"` 
	Author string `json:"author,omitempty"` 
	Message string `json:"message,omitempty"`
	Edited bool `json:"isEdited,omitempty"`
	Forum string `json:"forum,omitempty"`
	Thread int `json:"thread,omitempty"`
	Created time.Time `json:"created,omitempty"`
}

type Posts struct {
	Posts []Post `json:"posts,omitempty"`
}

type Users struct {
	Users []User `json:"users,omitempty"`
}

type Vote struct {
	Nickname string `json:"nickname,omitempty"`
	Voice int 	`json:"voice,omitempty"`
}

type PostsRelated struct {
	Related []string `json:"related,omitempty"`
}

type FullPostInfo struct {
	Post *Post `json:"post"`
	Author *User `json:"author"`
	Thread *Thread `json:"thread"`
	Forum *Forum `json:"forum"`
}

type PostMessage struct {
	Message string `db:"message"`
}

type Status struct {
	User   int `json:"user"`
	Forum  int `json:"forum"`
	Thread int `json:"thread"`
	Post   int `json:"post"`
}
