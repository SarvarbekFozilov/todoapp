package repository

import (
	"fmt"
	models "todo/models"

	"github.com/jmoiron/sqlx"
)

type AuthPostgress struct {
	db *sqlx.DB
}

func NewAuthPostgress(db *sqlx.DB) *AuthPostgress {
	return &AuthPostgress{db: db}
}

func (r *AuthPostgress) CreateUser(user *models.CreateUser) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s ( username, password_hash,photo,birth_location) values ($1, $2, $3,$4) RETURNING id", usersTable)

	row := r.db.QueryRow(query, user.Username, user.Password, user.Photo, user.Birthlocation)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
