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

func (uR *Usecase) CreateUserV2 (u *models.User) (*models.Users, error) {
	users := &models.Users{}

	oldNicknameUser, errNickname := uR.userRepo.GetUserByNickname(u.Nickname)
	if errNickname == nil {
		users.Users = append(users.Users, *oldNicknameUser)
	}
	
	oldEmailUser, errEmail := uR.userRepo.GetUserByEmail(u.Email)
	if errEmail == nil {
		if oldNicknameUser != nil {
			if oldNicknameUser.Nickname != oldEmailUser.Nickname && oldEmailUser.Email != oldNicknameUser.Email {
				users.Users = append(users.Users, *oldEmailUser)
			}
		} else {
			users.Users = append(users.Users, *oldEmailUser)
		}
	}

	if errNickname == nil || errEmail == nil {
		return users, models.ErrorUserExists
	}

	newUser, err := uR.userRepo.CreateUser(u)
	if err != nil {
		return nil, err
	}
	users.Users = append(users.Users, *newUser)
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