package usecase

import (
	"bd_tp_V2/models"
	userRepo "bd_tp_V2/user/repository"
)

type Usecase struct {
	userRepo *userRepo.Repository
}

func NewUserUsecase(uR *userRepo.Repository) *Usecase {
	return &Usecase{
		userRepo: uR,
	}
}
/*
func (uR *Usecase) CreateUser(u *models.User) ([]models.User, bool, error) {
	users, isNew, err := uR.userRepo.CreateUser(u)
	if err != nil {

		return nil, isNew, err
	}
	return users, isNew, nil
}

func (uR *Usecase) ProfileInfo(nickname string) (*models.User, error) {
	resultUser, err := uR.userRepo.ProfileInfo(nickname)
	if err != nil {
		return nil, err
	}
	return resultUser, nil
}

func (uR *Usecase) UpdateProfile(u *models.User) (*models.User, bool, error) {
	oldProfile, err := uR.userRepo.ProfileInfo(u.Nickname)
	if err != nil {
		return nil, false, err
	}
	if u.Fullname == "" {
		u.Fullname = oldProfile.Fullname
	}
	if u.Email == "" {
		u.Email = oldProfile.Email
	}
	if u.About == "" {
		u.About = oldProfile.About
	}
	resultUser, isFound, err := uR.userRepo.UpdateProfile(u)
	if err != nil {
		return nil, isFound, err
	}
	return resultUser, isFound, nil
}
*/

func (uR *Usecase) CreateUserV2 (u *models.User) ([]models.User, error) {
	var users []models.User

	oldNicknameUser, errNickname := uR.userRepo.GetUserByNickname(u.Nickname)
	if errNickname == nil {
		users = append(users, *oldNicknameUser)
	}
	
	oldEmailUser, errEmail := uR.userRepo.GetUserByEmail(u.Email)
	if errEmail == nil {
		if oldNicknameUser != nil {
			if oldNicknameUser.Nickname != oldEmailUser.Nickname && oldEmailUser.Email != oldNicknameUser.Email {
				users = append(users, *oldEmailUser)
			}
		} else {
			users = append(users, *oldEmailUser)
		}
	}

	if errNickname == nil || errEmail == nil {
		return users, models.ErrorUserExists
	}

	newUser, err := uR.userRepo.CreateUser(u)
	if err != nil {
		return nil, err
	}
	users = append(users, *newUser)
	return users, nil
}

func (uR *Usecase) ProfileInfoV2 (nickname string) (*models.User, error) {
	user,err := uR.userRepo.GetUserByNickname(nickname)
	if err != nil {
		return nil,err
	}
	return user, nil
}

func (uR *Usecase) UpdateUserV2(user *models.User) (*models.User, error) {
	oldUser, err := uR.userRepo.GetUserByNickname(user.Nickname)
	if err != nil {
		return nil, err
	}
	if user.Fullname == "" {
		user.Fullname = oldUser.Fullname
	}
	if user.About == "" {
		user.About = oldUser.About
	}
	if user.Email == "" {
		user.Email = oldUser.Email
	}
	newUser, err :=  uR.userRepo.UpdateUser(user) 
	if err != nil {
		return nil, err
	}
	return newUser, nil
}