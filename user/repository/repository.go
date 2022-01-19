package repository

import (
	models "bd_tp_V2/models"
	"strings"
	"github.com/jackc/pgx"
)

type Repository struct {
	db *pgx.ConnPool
}

func NewRepository (db *pgx.ConnPool) *Repository {
	return &Repository{
		db: db,
	}
}

func (uR *Repository) CreateUser(user *models.User) (*models.User, error) {
	_, err := uR.db.Exec("CreateUserQueryV2",&user.Nickname,&user.Fullname,&user.About,&user.Email)
	if err != nil {
		return nil,err
	}
	return user, nil
}

func (uR *Repository) GetUserByNickname(nickname string) (*models.User, error) {
	user := &models.User{}
	err := uR.db.QueryRow("FindUserByNicknameQueryV2",nickname).Scan(&user.Nickname,&user.Fullname,&user.About,&user.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorUserNotFound
		} else {
			return nil, err
		}
	}
	return user, nil
}

func (uR *Repository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	err := uR.db.QueryRow("FindUserByEmailQueryV2",email).Scan(&user.Nickname,&user.Fullname,&user.About,&user.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, models.ErrorUserNotFound
		} else {
			return nil, err
		}
	}
	return user, nil
}

func (uR *Repository) UpdateUser(user *models.User) (*models.User,error) {
	newUser := &models.User{}
	err := uR.db.QueryRow("UpdateUserQuery",user.Fullname,user.About,user.Email,user.Nickname).Scan(
		&newUser.Nickname,&newUser.Fullname,&newUser.About,&newUser.Email)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return nil, models.ErrorUserUpdateConflict
		} else {
			return nil, err
		}
	}
	return newUser,nil
}