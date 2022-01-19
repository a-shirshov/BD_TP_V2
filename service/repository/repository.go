package repository

import (
	"bd_tp_V2/models"

	"github.com/jackc/pgx"
)

type Repository struct {
	db *pgx.ConnPool
}

func NewServiceRepository(db *pgx.ConnPool) *Repository {
	return &Repository{
		db: db,
	}
}

func (sR *Repository) Clear() error {
	_, err := sR.db.Exec("ClearQuery")
	if err != nil {
		return err
	}
	return nil
}

func (sR *Repository) GetStatus() (*models.Status, error) {
	status := &models.Status{}
	err := sR.db.QueryRow("GetStatusUserQuery").Scan(&status.User)
	if err != nil {
		return nil, err
	}
	err = sR.db.QueryRow("GetStatusForumQuery").Scan(&status.Forum)
	if err != nil {
		return nil, err
	}
	err = sR.db.QueryRow("GetStatusThreadQuery").Scan(&status.Thread)
	if err != nil {
		return nil, err
	}
	err = sR.db.QueryRow("GetStatusPostQuery").Scan(&status.Post)
	if err != nil {
		return nil, err
	}
	return status, nil
}