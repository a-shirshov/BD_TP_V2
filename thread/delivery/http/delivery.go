package delivery

import (
	"bd_tp_V2/models"
	"bd_tp_V2/response"
	threadUsecase "bd_tp_V2/thread/usecase"
	"net/http"
	"strings"
)

type ThreadDelivery struct {
	threadU *threadUsecase.Usecase
}

func NewThreadDelivery(tU *threadUsecase.Usecase) *ThreadDelivery {
	return &ThreadDelivery{
		threadU: tU,
	}
}

func (tD *ThreadDelivery) ThreadDetails(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	split := strings.Split(path, "/")
	slug_or_id := split[len(split)-2]

	thread, err := tD.threadU.ThreadDetails(slug_or_id)
	if err != nil {
		if err == models.ErrorThreadNotFound {
			response.SendResponse(w, http.StatusNotFound,models.Error{Message: err.Error()})
			return 
		} else {
			response.SendResponse(w, http.StatusInternalServerError, models.Error{Message: err.Error()})
			return
		}
	}
	response.SendResponse(w,http.StatusOK,thread)
}

func (tD *ThreadDelivery) CreatePostsNew(w http.ResponseWriter, r *http.Request) {
	postsRequest, err := response.GetPostsFromRequest(r.Body)
	if err != nil {
		return
	}

	path := r.URL.Path
	split := strings.Split(path, "/")
	slug_or_id := split[len(split)-2]

	posts := &models.Posts{Posts: postsRequest}

	newPosts,err := tD.threadU.CreatePostsNew(posts, slug_or_id)
	if err != nil {
		switch err {

		case models.ErrorThreadNotFound, models.ErrorUserNotFound:
			response.SendResponse(w,http.StatusNotFound,models.Error{Message: err.Error()})
			return

		case models.ErrorPostNotFound:
			response.SendResponse(w,http.StatusConflict,models.Error{Message: err.Error()})
			return

		default:
			response.SendResponse(w, http.StatusInternalServerError,models.Error{Message: err.Error()})
			return
		}
	}
	if len(posts.Posts) == 0 {
		response.OldSendResponse(w, http.StatusCreated, []models.Post{})
		return 
	}
	response.OldSendResponse(w, http.StatusCreated, newPosts.Posts)
}

func (tD *ThreadDelivery) ThreadVote(w http.ResponseWriter, r *http.Request) {
	voteRequest, err := response.GetVoteFromRequest(r.Body)
	if err != nil {
		return
	}
	path := r.URL.Path
	split := strings.Split(path, "/")
	slug_or_id := split[len(split)-2]

	thread, err := tD.threadU.ThreadVote(voteRequest, slug_or_id)
	if err != nil {
		if err == models.ErrorThreadNotFound || err == models.ErrorUserNotFound {
			response.SendResponse(w,http.StatusNotFound,models.Error{Message: err.Error()})
			return
		} else {
			response.SendResponse(w, http.StatusInternalServerError,models.Error{Message: err.Error()})
			return
		}
	}
	response.SendResponse(w,http.StatusOK,thread)
}

func (tD *ThreadDelivery) ThreadGetPosts(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	split := strings.Split(path, "/")
	slug_or_id := split[len(split)-2]

	q := r.URL.Query()
	var limit string
	var since string
	var sort string
	var desc string
	if len(q["limit"]) > 0 {
		limit = q["limit"][0]
	}
	if len(q["since"]) > 0 {
		since = q["since"][0]
	}
	if len(q["sort"]) > 0 {
		sort = q["sort"][0]
	}
	if len(q["desc"]) > 0 {
		desc = q["desc"][0]
	}

	posts, err := tD.threadU.ThreadGetPosts(slug_or_id, limit, since, sort, desc)
	if err != nil {
		if err == models.ErrorThreadNotFound {
			response.SendResponse(w,http.StatusNotFound,models.Error{Message: err.Error()})
			return
		} else {
			response.SendResponse(w, http.StatusInternalServerError,models.Error{Message: err.Error()})
			return
		}
	}

	if len(posts.Posts) == 0 {
		response.OldSendResponse(w,http.StatusOK,[]models.Post{})
		return
	}
	response.OldSendResponse(w,http.StatusOK, posts.Posts)
}

func (tD *ThreadDelivery) ThreadDetailsUpdate(w http.ResponseWriter, r *http.Request) {
	threadRequest, err := response.GetThreadFromRequest(r.Body)
	if err != nil {
		return
	}
	path := r.URL.Path
	split := strings.Split(path, "/")
	slug_or_id := split[len(split)-2]

	thread, err := tD.threadU.UpdateThreadDetails(slug_or_id,threadRequest)
	if err != nil {
		if err == models.ErrorThreadNotFound {
			response.SendResponse(w,http.StatusNotFound,models.Error{Message: err.Error()})
			return
		} else {
			response.SendResponse(w, http.StatusInternalServerError,models.Error{Message: err.Error()})
			return
		}
	}
	response.SendResponse(w,http.StatusOK, thread)
}