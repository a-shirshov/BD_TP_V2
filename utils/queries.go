package utils

type Queries struct {
	Name string
	Query string
}

const (
	FindUserByNicknameQueryV2 = `select nickname,fullname,about,email from "user" where nickname = $1`
	CreateUserQueryV2 = `insert into "user" (nickname, fullname, about, email) values ($1, $2, $3, $4)`
	FindUserByEmailQueryV2 = `select nickname,fullname,about,email from "user" where email = $1`
	UpdateUserQuery = `update "user" set fullname = $1, about = $2, email = $3 
	where nickname = $4 returning nickname, fullname, about, email`
)

var queries = []Queries {
	{
		Name: "FindUserByNicknameQueryV2",
		Query: FindUserByNicknameQueryV2,
	},
	{
		Name: "CreateUserQueryV2",
		Query: CreateUserQueryV2,
	},
	{
		Name: "FindUserByEmailQueryV2",
		Query: FindUserByEmailQueryV2,
	},
	{
		Name: "UpdateUserQuery",
		Query: UpdateUserQuery,
	},
}