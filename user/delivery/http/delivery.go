package delivery

import (
	"bd_tp_V2/models"
	"bd_tp_V2/response"
	userUseCase "bd_tp_V2/user/usecase"
	"net/http"
	"strings"
)

type UserDelivery struct {
	userUsecase *userUseCase.Usecase
}

func NewUserDelivery(uU *userUseCase.Usecase) *UserDelivery {
	return &UserDelivery{
		userUsecase: uU,
	}
}

func (uD *UserDelivery) CreateUserV2(w http.ResponseWriter, r *http.Request) {
	u, err := response.GetUserFromRequest(r.Body)

	path := r.URL.Path
	split := strings.Split(path, "/")
	nickname := split[len(split)-2]
	u.Nickname = nickname

	if err != nil {
		return
	}
	users,err := uD.userUsecase.CreateUserV2(u)

	if err != nil {
		if err == models.ErrorUserExists {
			response.SendResponse(w,http.StatusConflict,users)
			return
		} else {
			response.SendResponse(w, http.StatusInternalServerError, models.Error{Message: err.Error()})
			return
		}
	}
	response.SendResponse(w, http.StatusCreated,users[0])
}

func (uD *UserDelivery) ProfileInfoV2(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	split := strings.Split(path, "/")
	nickname := split[len(split)-2]

	user, err := uD.userUsecase.ProfileInfoV2(nickname)
	if err != nil {
		if err == models.ErrorUserNotFound {
			response.SendResponse(w,http.StatusNotFound,models.Error{Message: err.Error()})
			return
		} else {
			response.SendResponse(w, http.StatusInternalServerError,models.Error{Message: err.Error()})
			return
		}
	}
	response.SendResponse(w, http.StatusOK,user)
}

func (uD *UserDelivery) UpdateProfileV2(w http.ResponseWriter, r *http.Request) {
	user, err := response.GetUserFromRequest(r.Body)
	if err != nil {
		return
	}

	path := r.URL.Path
	split := strings.Split(path, "/")
	nickname := split[len(split)-2]
	user.Nickname = nickname
	
	newProfile, err := uD.userUsecase.UpdateUserV2(user)
	if err != nil {
		switch err {

		case models.ErrorUserNotFound:
			response.SendResponse(w,http.StatusNotFound,models.Error{Message: err.Error()})
			return

		case models.ErrorUserUpdateConflict:
			response.SendResponse(w,http.StatusConflict,models.Error{Message: err.Error()})
			return

		default:
			response.SendResponse(w, http.StatusInternalServerError,models.Error{Message: err.Error()})
			return
		}

	}
	response.SendResponse(w, http.StatusOK,newProfile)
}