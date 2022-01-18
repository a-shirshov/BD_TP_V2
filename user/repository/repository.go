package repository

import (
	models "bd_tp_V2/models"
	"strings"
	"github.com/jackc/pgx"
)

/*
const (
	createUserQuery = `insert into "user" (nickname, fullname, about, email) values ($1, $2, $3, $4)`
	findUserByNicknameQueryV2 = `select nickname,fullname,about,email from "user" where nickname = $1`
	findUserByEmailQuery = `select * from "user" where email = $1`
	updateUserQuery = `update "user" set fullname = $1,about = $2,email = $3 where nickname = $4 `
	getUserID = `select id from "user" where nickname = $1`
)
*/

type Repository struct {
	db *pgx.ConnPool
}

func NewRepository (db *pgx.ConnPool) *Repository {
	return &Repository{
		db: db,
	}
}
/*
func (uR *Repository) CreateUser(u *models.User) ([]models.User, bool, error) {
	isNew := true
	query := createUserQuery
	var users []models.User
	_, err := uR.db.Exec(query,u.Nickname,u.Fullname,u.About,u.Email)
	if err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint`) {
			query := findUserByNicknameQuery
			userByNickname := models.User{}
			err := uR.db.Get(&userByNickname,query,u.Nickname)
			if err == nil {
				users = append(users, userByNickname)
			}
			query = findUserByEmailQuery
			userByEmail := models.User{}
			err = uR.db.Get(&userByEmail,query,u.Email)
			if err == nil {
				if userByNickname.Email != userByEmail.Email{
					users = append(users, userByEmail)
				}
			}
			isNew = false
			return users, isNew, nil
		}

	}
	users = append(users, *u)
	return users, isNew, nil
}

func (uR* Repository) ProfileInfo (nickname string) (*models.User, error) {
	query := findUserByNicknameQuery
	user := models.User{}
	err := uR.db.Get(&user,query,nickname)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (uR *Repository) UpdateProfile (u *models.User) (*models.User, bool, error) {
	query := updateUserQuery
	user := models.User{}
	isFound := true
	result,err := uR.db.Exec(query,u.Fullname,u.About,u.Email,u.Nickname)
	if err != nil {
		
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "user_email_key"`) {
			return nil,isFound,err
		}
		isFound := false
		return nil, isFound, err
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		isFound := false
		return nil, isFound, errors.New("no user with with nickname")
	}
	
	query = findUserByNicknameQuery
	err = uR.db.Get(&user,query,u.Nickname)
	if err != nil {
		
		return nil, isFound, err
	}
	return &user, isFound, nil
}

func (uR *Repository) GetIdByNickname (nickname string) (int, error) {
	query := getUserID
	var userID int
	err := uR.db.Get(&userID,query,nickname)
	if err != nil {
		return 0, err
	}
	return userID,nil
}
*/
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