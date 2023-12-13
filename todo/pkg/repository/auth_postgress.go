package repository

import (
	models "todo/models"

	"github.com/jmoiron/sqlx"
)

type AuthPostgress struct {
	db *sqlx.DB
}

func NewAuthPostgress(db *sqlx.DB) *AuthPostgress {
	return &AuthPostgress{db: db}
}

func (r *AuthPostgress) CreateUser(user *models.CreateUserReq) (int, error) {
	var id int
	query := `
	INSERT INTO users (
		fullname,
		username,
		password_hash,
		photo,
		birthday,
		location
	) values ($1,$2,$3,$4,$5,$6) RETURNING id;`

	if err := r.db.QueryRow(query,
		user.FullName,
		user.Username,
		user.Password,
		user.Photo,
		user.Birthday,
		user.Location,
	).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (r *AuthPostgress) GetUser(username, password string) (*models.UserResponse, error) {
	var user models.UserResponse
	query := "SELECT id FROM users WHERE username=$1 AND password_hash=$2"

	if err := r.db.QueryRow(query, username, password).Scan(&user.ID); err != nil {
		return nil, err
	}

	return &user, nil
}
