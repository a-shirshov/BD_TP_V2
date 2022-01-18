package response

import (
	models "bd_tp_V2/models"
	"encoding/json"
	"io"
	"net/http"
)

func GetUserFromRequest(r io.Reader) (*models.User, error) {
	userInput := new(models.User)
	//err := json.UnmarshalFromReader(r, userInput)
	err := json.NewDecoder(r).Decode(userInput)
	if err != nil {
		return nil, err
	}
	return userInput, nil
}


func GetForumFromRequest(r io.Reader) (*models.Forum, error) {
	forumInput := new(models.Forum)
	//err := json.UnmarshalFromReader(r, forumInput)
	err := json.NewDecoder(r).Decode(forumInput)
	if err != nil {
		return nil, err
	}
	return forumInput, nil
}
/*

func GetThreadFromRequest(r io.Reader) (*models.Thread, error) {
	threadInput := new(ThreadResponse)
	//err := json.UnmarshalFromReader(r, forumInput)
	err := json.NewDecoder(r).Decode(threadInput)
	if err != nil {
		return nil, err
	}
	result := &models.Thread{
		Title:   threadInput.Title,
		Author:  threadInput.Author,
		Message: threadInput.Message,
		Created: threadInput.Created,
		Slug:    threadInput.Slug,
	}
	return result, nil
}

func GetThreadsQueryInfo(r io.Reader) (*models.ForumThreadsRequest, error) {
	infoInput := new(ForumThreadsRequest)
	//err := json.UnmarshalFromReader(r, forumInput)
	err := json.NewDecoder(r).Decode(infoInput)
	if err != nil {
		return nil, err
	}
	result := &models.ForumThreadsRequest{
		Limit: infoInput.Limit,
		Since: infoInput.Since,
		Desc:  infoInput.Desc,
	}
	return result, nil
}

func GetPostsFromRequest(r io.Reader) ([]models.Post, error) {
	var postsInput []PostRequest
	err := json.NewDecoder(r).Decode(&postsInput)
	if err != nil {

		return nil, err
	}
	var posts []models.Post
	for _, post := range postsInput {
		posts = append(posts, models.Post{
			Parent:  post.Parent,
			Author:  post.Author,
			Message: post.Message,
		})
	}
	return posts, nil
}

func GetThreadUpdateFromRequest(r io.Reader) (*models.Thread, error) {
	var thread ThreadResponse
	err := json.NewDecoder(r).Decode(&thread)
	if err != nil {

		return nil, err
	}
	result := &models.Thread{
		Title:   thread.Title,
		Message: thread.Message,
	}
	return result, nil
}

func GetVoteFromRequest(r io.Reader) (*models.Vote, error) {
	var vote VoteRequest
	err := json.NewDecoder(r).Decode(&vote)
	if err != nil {

		return nil, err
	}
	result := &models.Vote{
		Nickname: vote.Nickname,
		Voice:    *vote.Voice,
	}
	return result, nil
}

func GetPostRelatedFromRequest(r io.Reader) (*models.PostsRelated, error) {
	var related PostRelated
	err := json.NewDecoder(r).Decode(&related)
	if err != nil {

		return nil, err
	}
	result := &models.PostsRelated{
		Related: related.Related,
	}
	return result, nil
}

func GetPostFromRequest(r io.Reader) (*models.Post, error) {
	var post PostResponse
	err := json.NewDecoder(r).Decode(&post)
	if err != nil {

		return nil, err
	}
	result := &models.Post{
		Message: post.Message,
	}
	return result, nil
}
*/
func SendResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if response != nil {
		b, err := json.Marshal(response)
		if err != nil {
			return
		}
		_, err = w.Write(b)
		if err != nil {
			return
		}
	}
}
