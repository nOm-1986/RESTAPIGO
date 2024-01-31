package database

import (
	"context"
	"database/sql"
	"rest-api-go/models"
)

type MysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository(url string) (*MysqlRepository, error) {
	db, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	return &MysqlRepository{db}, nil
}

func (repo *MysqlRepository) InsertUser(ctx context.Context, user *models.User) error {
	//
}
