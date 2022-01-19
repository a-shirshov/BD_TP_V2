package response

import (
	models "bd_tp_V2/models"
	oldjson "encoding/json"
	"io"
	"net/http"
	json "github.com/mailru/easyjson"
)

func GetUserFromRequest(r io.Reader) (*models.User, error) {
	userInput := new(models.User)
	err := json.UnmarshalFromReader(r, userInput)
	//err := json.NewDecoder(r).Decode(userInput)
	if err != nil {
		return nil, err
	}
	return userInput, nil
}


func GetForumFromRequest(r io.Reader) (*models.Forum, error) {
	forumInput := new(models.Forum)
	err := json.UnmarshalFromReader(r, forumInput)
	//err := json.NewDecoder(r).Decode(forumInput)
	if err != nil {
		return nil, err
	}
	return forumInput, nil
}


func GetThreadFromRequest(r io.Reader) (*models.Thread, error) {
	threadInput := new(models.Thread)
	err := json.UnmarshalFromReader(r, threadInput)
	//err := json.NewDecoder(r).Decode(threadInput)
	if err != nil {
		return nil, err
	}
	return threadInput, nil
}

func GetPostsFromRequest(r io.Reader) ([]models.Post, error) {
	var postsInput []models.Post
	//err := json.UnmarshalFromReader(r, &postsInput)
	err := oldjson.NewDecoder(r).Decode(&postsInput)
	if err != nil {
		return nil, err
	}
	return postsInput, nil
}

func GetVoteFromRequest(r io.Reader) (*models.Vote, error) {
	voteInput := new(models.Vote)
	err := json.UnmarshalFromReader(r, voteInput)
	//err := json.NewDecoder(r).Decode(&voteInput)
	if err != nil {
		return nil, err
	}
	return voteInput, nil
}

func GetPostFromRequest(r io.Reader) (*models.Post, error) {
	postInput := new(models.Post)
	err := json.UnmarshalFromReader(r, postInput)
	//err := json.NewDecoder(r).Decode(&postInput)
	if err != nil {
		return nil, err
	}
	return postInput, nil
}


func SendResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if response != nil {
		b, err := oldjson.Marshal(response)
		if err != nil {
			return
		}
		_, err = w.Write(b)
		if err != nil {
			return
		}
	}
}
