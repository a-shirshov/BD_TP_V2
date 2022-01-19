package delivery

import (
	postUsecase "bd_tp_V2/post/usecase"
	"bd_tp_V2/response"
	"net/http"
	"bd_tp_V2/models"
	"strings"
	"strconv"
)

type PostDelivery struct {
	PostUsecase *postUsecase.Usecase
}

func NewPostDelivery(pU *postUsecase.Usecase) *PostDelivery {
	return &PostDelivery{
		PostUsecase: pU,
	}
}

func (pD *PostDelivery) PostDetails(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	split := strings.Split(path, "/")
	id := split[len(split)-2]

	q := r.URL.Query()
	var related string
	if len(q["related"]) > 0 {
		related = q["related"][0]
	}

	fullPost, err := pD.PostUsecase.PostDetails(id, related)
	if err != nil {
		if err == models.ErrorPostNotFound {
			response.SendResponse(w, http.StatusNotFound, models.Error{Message: err.Error()})
			return 
		} else {
			response.SendResponse(w, http.StatusInternalServerError,models.Error{Message: err.Error()})
			return
		}
	}
	response.SendResponse(w, http.StatusOK, fullPost)
}

func (pD *PostDelivery) UpdatePost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	split := strings.Split(path, "/")
	id := split[len(split)-2]

	post, err := response.GetPostFromRequest(r.Body)
	if err != nil {
		return
	}

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	post.ID = idInt
	newPost, err := pD.PostUsecase.UpdatePost(post)
	if err != nil {
		if err == models.ErrorPostNotFound {
			response.SendResponse(w, http.StatusNotFound, models.Error{Message: err.Error()})
			return 
		} else {
			response.SendResponse(w, http.StatusInternalServerError,models.Error{Message: err.Error()})
			return
		}
	}
	response.SendResponse(w, http.StatusOK, newPost)
}