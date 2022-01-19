package delivery

import (
	forumUsecase "bd_tp_V2/forum/usecase"
	"bd_tp_V2/models"
	"bd_tp_V2/response"
	"net/http"
	"strings"
)

type ForumDelivery struct {
	ForumUsecase *forumUsecase.Usecase
}

func NewForumDelivery(fU *forumUsecase.Usecase) *ForumDelivery {
	return &ForumDelivery{
		ForumUsecase: fU,
	}
}

func (fD *ForumDelivery) CreateForumV2(w http.ResponseWriter, r *http.Request) {
	f, err := response.GetForumFromRequest(r.Body)
	if err != nil {
		return
	}

	forum, err := fD.ForumUsecase.CreateForum(f)
	if err != nil {
		switch err {
		case models.ErrorForumExists:
			response.SendResponse(w,http.StatusConflict,forum)
			return

		case models.ErrorUserNotFound:
			response.SendResponse(w,http.StatusNotFound,models.Error{Message: err.Error()})
			return

		default:
			response.SendResponse(w, http.StatusInternalServerError,models.Error{Message: err.Error()})
			return
		}
	}
	response.SendResponse(w, http.StatusCreated, forum)
}

func (fD *ForumDelivery) ForumDetailsV2(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	split := strings.Split(path, "/")
	slug := split[len(split)-2]

	forum, err := fD.ForumUsecase.ForumDetails(slug)
	if err != nil {
		if err == models.ErrorForumNotFound {
			response.SendResponse(w, http.StatusNotFound,models.Error{Message: err.Error()})
			return
		} else {
			response.SendResponse(w,http.StatusInternalServerError, models.Error{Message: err.Error()})
		}
	}
	response.SendResponse(w,http.StatusOK,forum)
}

func (fD *ForumDelivery) ForumThreadCreateV2(w http.ResponseWriter, r *http.Request) {
	th, err := response.GetThreadFromRequest(r.Body)
	if err != nil {
		return
	}
	path := r.URL.Path
	split := strings.Split(path, "/")
	slug := split[len(split)-2]

	thread, err := fD.ForumUsecase.ForumThreadCreate(slug, th)
	if err != nil {
		if err == models.ErrorThreadExists {
			response.SendResponse(w, http.StatusConflict, thread)
			return
		} else if err == models.ErrorUserNotFound || err == models.ErrorForumNotFound {
			response.SendResponse(w, http.StatusNotFound, models.Error{Message: err.Error()})
			return 
		} else {
			response.SendResponse(w, http.StatusInternalServerError, models.Error{Message: err.Error()})
			return 
		}
	}
	response.SendResponse(w, http.StatusCreated, thread)
}

func (fD *ForumDelivery) GetThreadsByForum(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path
	split := strings.Split(path, "/")
	slug := split[len(split)-2]

	q := r.URL.Query()
	var limit string
	var since string
	var desc string
	if len(q["limit"]) > 0 {
		limit = q["limit"][0]
	}
	if len(q["since"]) > 0 {
		since = q["since"][0]
	}
	if len(q["desc"]) > 0 {
		desc = q["desc"][0]
	}

	threads, err := fD.ForumUsecase.GetThreadsByForum(slug, limit, since, desc)
	if err != nil {
		if err == models.ErrorForumNotFound {
			response.SendResponse(w, http.StatusNotFound, models.Error{Message: err.Error()})
			return 
		} else {
			response.SendResponse(w, http.StatusInternalServerError, models.Error{Message: err.Error()})
			return 
		}
	}
	if len(threads.Threads) == 0 {
		response.SendResponse(w, http.StatusOK, []models.Thread{})
		return 
	}
	response.SendResponse(w, http.StatusOK, threads.Threads)
}

func (fD *ForumDelivery) GetForumUsers(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	split := strings.Split(path, "/")
	slug := split[len(split)-2]

	q := r.URL.Query()
	var limit string
	var since string
	var desc string
	if len(q["limit"]) > 0 {
		limit = q["limit"][0]
	}
	if len(q["since"]) > 0 {
		since = q["since"][0]
	}
	if len(q["desc"]) > 0 {
		desc = q["desc"][0]
	}

	users, err := fD.ForumUsecase.GetForumUsers(slug, limit, since, desc)
	if err != nil {
		if err == models.ErrorForumNotFound {
			response.SendResponse(w,http.StatusNotFound,models.Error{Message: err.Error()})
			return
		} else {
			response.SendResponse(w, http.StatusInternalServerError, models.Error{Message: err.Error()})
			return 
		}
	}

	if len(users.Users) == 0 {
		response.SendResponse(w, http.StatusOK, []models.User{})
		return 
	}
	response.SendResponse(w, http.StatusOK, users.Users)
}

