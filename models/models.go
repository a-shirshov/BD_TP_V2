package models

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
	ID int `db:"id"` 
	Title string `db:"title"` 
	Author string `db:"author"` 
	Forum string `db:"forum"`
	Message string `db:"message"`
	Votes int `db:"votes"`
	Slug string `db:"slug"`
	Created string `db:"created"`
}

type ForumThreadsRequest struct {
	Slug string 
	Limit string 
	Since string 
	Desc string 
}

type Threads struct {
	Threads []Thread `json:"threads,omitempty"`
}

type Post struct {
	ID int `db:"id"`
	Parent int `db:"parent"` 
	Author string `db:"author"` 
	Message string `db:"message"`
	Edited bool `db:"edited"`
	Forum string `db:"forum"`
	Thread int `db:"thread"`
	Created string `db:"created"`
}

type Posts struct {
	Posts []Post `json:"posts,omitempty"`
}

type Vote struct {
	Nickname string `json:"nickname,omitempty"`
	Voice int 	`json:"voice,omitempty"`
}

type PostsRelated struct {
	Related []string `json:"related,omitempty"`
}

type FullPostInfo struct {
	Post *Post
	Author *User
	Thread *Thread
	Forum *Forum
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
