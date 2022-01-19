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
	CheckNicknameQuery = `select nickname from "user" where nickname = $1`

	ForumDetailsQuery = `select title, "user", slug, posts, threads from "forum" where slug = $1`
	CreateForumQuery = `insert into "forum" (title, "user", slug) values ($1, $2, $3) 
	returning title, "user", slug, posts, threads`
	CreateForumThreadQuery = `insert into "thread" (title, author, forum, message, slug, created) values ($1, $2, $3, $4, $5, $6) 
	returning id, title, author, forum, message, votes, slug, created`

	GetForumThreadsAscSinceQuery = `select * from thread where forum = $1 and created >= $2 order by created limit $3`
	GetForumThreadsAscQuery = `select * from thread where forum = $1 order by created limit $2`
	GetForumThreadsDescQuery = `select * from thread where forum = $1 order by created desc limit $2`
	GetForumThreadsDescSinceQuery = `select * from thread where forum = $1 and created <= $2 order by created desc limit $3`

	FindThreadWithIdQuery = `select * from thread where id = $1`
	FindThreadWithSlugQuery = `select * from thread where slug = $1`
	CheckThreadQuery = `select id from post where thread = $1 and id = $2`
	InsertVoteQuery = `insert into vote (thread, "user", voice) values ($1, $2, $3)`
	UpdateVoteQuery = `update vote set voice = $1 where thread = $2 and "user" = $3`
	UpdateThreadDetailsQuery = `update thread set title = $1, message = $2 where id = $3 
	returning id, title, author, forum, message, votes, slug, created`

	CreatePostQuery = `insert into post (parent, author, message, forum, thread, created) values ($1, $2, $3, $4, $5, $6)
	returning id, parent, author, message, is_edited, forum, thread, created`
	FindPostByIDQuery = `select id, parent, author, message, is_edited, forum, thread, created from post where id = $1`
	UpdatePostByIdQuery = `update post set message = $1, is_edited = true where id = $2 
	returning id, parent, author, message, is_edited, forum, thread, created`

	GetUsersByForumDesc = `select nickname, fullname, about, email from "user" as u
	join forum_user as fu on u.nickname = fu."user" and fu.forum = $1 
	order by nickname desc limit $2`
	GetUsersByForumDescSince = `select nickname, fullname, about, email from "user" as u
	join forum_user as fu on u.nickname = fu."user" and fu.forum = $1 and nickname < $2 
	order by nickname desc limit $3`
	GetUsersByForumAsc = `select nickname, fullname, about, email from "user" as u
	join forum_user as fu on u.nickname = fu."user" and fu.forum = $1 
	order by nickname limit $2`
	GetUsersByForumAscSince = `select nickname, fullname, about, email from "user" as u
	join forum_user as fu on u.nickname = fu."user" and fu.forum = $1 and nickname > $2 
	order by nickname limit $3`

	GetPostsByThreadFlatDescSince = `select id, parent, author, message, is_edited, forum, thread, created from post 
	where thread = $1 and id < $2
	order by id desc limit $3`
	GetPostsByThreadFlatDesc = `select id, parent, author, message, is_edited, forum, thread, created from post 
	where thread = $1 
	order by id desc limit $2`
	GetPostsByThreadFlatAsc = `select id, parent, author, message, is_edited, forum, thread, created
	from post 
	where thread = $1 
	order by id limit $2`
	GetPostsByThreadFlatAscSince = `select id, parent, author, message, is_edited, forum, thread, created from post 
	where thread = $1 and id > $2 
	order by id limit $3`

	GetPostsByThreadTreeDescSince = `select id, parent, author, message, is_edited, forum, thread, created from post 
	where thread = $1 and path < (select path FROM post where id = $2) 
	order by path desc, id desc limit $3`
	GetPostsByThreadTreeDesc = `select id, parent, author, message, is_edited, forum, thread, created from post 
	where thread = $1 
	order by path desc, id desc limit $2`
	GetPostsByThreadTreeAsc = `select id, parent, author, message, is_edited, forum, thread, created from post 
	where thread = $1 
	order by path, id limit $2`
	GetPostsByThreadTreeAscSince = `select id, parent, author, message, is_edited, forum, thread, created from post 
	where thread = $1 and path > (select path FROM post where id = $2) 
	order by path, id limit $3`

	QueryGetThreadPostsParentTreeAscSince = `select id, parent, author, message, is_edited, forum, thread, created from post 
	where path[1] in 
		(select id from post where thread = $1 and parent = 0 and path[1] > (select path[1] from post where id = $2) 
		order by id limit $3)
	order by path, id`
	QueryGetThreadPostsParentTreeAsc = `select id, parent, author, message, is_edited, forum, thread, created from post 
	where path[1] in (select id from post where thread = $1 and parent = 0 order by id limit $2)
	order by path, id` 
	QueryGetThreadPostsParentTreeDescSince = `select id, parent, author, message, is_edited, forum, thread, created 
	from post where path[1] in 
		(select id from post where thread = $1 and parent = 0 and path[1] <
		(select path[1] from post where id = $2) order by id desc limit $3)
	order by path[1] desc, path, id`
	QueryGetThreadPostsParentTreeDesc = `select id, parent, author, message, is_edited, forum, thread, created 
	from post where path[1] in (select id from post where thread = $1 and parent = 0 order by id desc limit $2)
	order by path[1] desc, path, id`

	ClearQuery = `truncate forum_user, vote, post, thread, forum, "user"`
	GetStatusUserQuery   = `select count(*) from "user"`
	GetStatusForumQuery  = `select count(*) from forum`
	GetStatusThreadQuery = `select count(*) from thread`
	GetStatusPostQuery   = `select count(*) from post`
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
	{
		Name: "ForumDetailsQuery",
		Query: ForumDetailsQuery,
	},
	{
		Name: "CreateForumQuery",
		Query: CreateForumQuery,
	},
	{
		Name: "FindThreadWithIdQuery",
		Query: FindThreadWithIdQuery,
	},
	{
		Name: "FindThreadWithSlugQuery",
		Query: FindThreadWithSlugQuery,
	},
	{
		Name: "CreateForumThreadQuery",
		Query: CreateForumThreadQuery,
	},
	{
		Name: "GetForumThreadsDescSinceQuery",
		Query: GetForumThreadsDescSinceQuery,
	},
	{
		Name: "GetForumThreadsDescQuery",
		Query: GetForumThreadsDescQuery,
	},
	{
		Name: "GetForumThreadsAscQuery",
		Query: GetForumThreadsAscQuery,
	},
	{
		Name: "GetForumThreadsAscSinceQuery",
		Query: GetForumThreadsAscSinceQuery,
	},
	{
		Name: "CheckNicknameQuery",
		Query: CheckNicknameQuery,
	},
	{
		Name: "CheckThreadQuery",
		Query: CheckThreadQuery,
	},
	{
		Name: "CreatePostQuery",
		Query: CreatePostQuery,
	},
	{
		Name: "InsertVoteQuery",
		Query: InsertVoteQuery,
	},
	{
		Name: "UpdateVoteQuery",
		Query: UpdateVoteQuery,
	},
	{
		Name: "GetPostsByThreadFlatAscSince",
		Query: GetPostsByThreadFlatAscSince,
	},
	{
		Name: "GetPostsByThreadFlatAsc",
		Query: GetPostsByThreadFlatAsc,
	},
	{
		Name: "GetPostsByThreadFlatDesc",
		Query: GetPostsByThreadFlatDesc,
	},
	{
		Name: "GetPostsByThreadFlatDescSince",
		Query: GetPostsByThreadFlatDescSince,
	},
	{
		Name: "GetPostsByThreadTreeDescSince",
		Query: GetPostsByThreadTreeDescSince,
	},
	{
		Name: "GetPostsByThreadTreeDesc",
		Query: GetPostsByThreadTreeDesc,
	},
	{
		Name: "GetPostsByThreadTreeAscSince",
		Query: GetPostsByThreadTreeAscSince,
	},
	{
		Name: "GetPostsByThreadTreeAsc",
		Query: GetPostsByThreadTreeAsc,
	},
	{
		Name: "QueryGetThreadPostsParentTreeAscSince",
		Query: QueryGetThreadPostsParentTreeAscSince,
	},
	{
		Name: "QueryGetThreadPostsParentTreeAsc",
		Query: QueryGetThreadPostsParentTreeAsc,
	},
	{
		Name: "QueryGetThreadPostsParentTreeDescSince",
		Query: QueryGetThreadPostsParentTreeDescSince,
	},
	{
		Name: "QueryGetThreadPostsParentTreeDesc",
		Query: QueryGetThreadPostsParentTreeDesc,
	},
	{
		Name: "UpdateThreadDetailsQuery",
		Query: UpdateThreadDetailsQuery,
	},
	{
		Name: "GetUsersByForumDesc",
		Query: GetUsersByForumDesc,
	},
	{
		Name: "GetUsersByForumDescSince",
		Query: GetUsersByForumDescSince,
	},
	{
		Name: "GetUsersByForumAsc",
		Query: GetUsersByForumAsc,
	},
	{
		Name: "GetUsersByForumAscSince",
		Query: GetUsersByForumAscSince,
	},
	{
		Name: "FindPostByIDQuery",
		Query: FindPostByIDQuery,
	},
	{
		Name: "UpdatePostByIdQuery",
		Query: UpdatePostByIdQuery,
	},
	{
		Name: "ClearQuery",
		Query: ClearQuery,
	},
	{
		Name: "GetStatusUserQuery",
		Query: GetStatusUserQuery,
	},
	{
		Name: "GetStatusForumQuery",
		Query: GetStatusForumQuery,
	},
	{
		Name: "GetStatusThreadQuery",
		Query: GetStatusThreadQuery,
	},
	{
		Name: "GetStatusPostQuery",
		Query: GetStatusPostQuery,
	},
}