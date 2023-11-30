package postgres

import (
	"context"

	"fmt"
	"user/models"
	"user/pkg/helper"

	"github.com/google/uuid"

	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (u *userRepo) CreateUser(ctx context.Context, req *models.CreateUser) (string, error) {
	id := uuid.NewString()

	query := `INSERT INTO users(
	                        id,
							username,
							password,
							photo,
						    location,
						     created_at)
							VALUES($1, $2, $3,$4,$5,now())`

	_, err := u.db.Exec(ctx, query,
		id,
		req.Username,
		req.Password,
		req.Photo,
		req.Blocation,
	)

	if err != nil {
		return "Error CreateUser", err
	}

	return id, nil
}

func (u userRepo) GetUser(ctx context.Context, req *models.IdRequest) (rep *models.User, err error) {
	query := `SELECT 
                    id,
                    username,
                    password,
					photo,
					location
				FROM
				    users
				WHERE
				    id = $1;`

	user := models.User{}

	err = u.db.QueryRow(context.Background(), query, req.Id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Photo,
		&user.Blocation,
	)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

func (u *userRepo) GetAllUser(ctx context.Context, req *models.GetAllUserRequest) (resp *models.GetAllUser, err error) {
	params := make(map[string]interface{})
	filter := " WHERE true "
	resp = &models.GetAllUser{}
	resp.Users = make([]models.User, 0)
	offset := (req.Page - 1) * req.Limit

	query := `SELECT 
	               COUNT(*) OVER(),
	               id,
				   username,
				   password,
				   photo,
				   location
			FROM users`

	params["limit"] = req.Limit
	params["offset"] = offset

	query = query + filter + "ORDER BY created_at desc  LIMIT :limit OFFSET :offset"
	resQuery, pArr := helper.ReplaceQueryParams(query, params)

	rows, err := u.db.Query(context.Background(), resQuery, pArr...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	count := 0
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&count,
			&user.ID,
			&user.Username,
			&user.Password,
			&user.Photo,
			&user.Blocation,
		)
		if err != nil {
			return nil, err
		}

		resp.Count = count
		resp.Users = append(resp.Users, user)
	}

	return resp, nil
}

func (u *userRepo) UpdateUser(ctx context.Context, req *models.UpdateUser) (string, error) {
	// Validate the UUID string
	_, err := uuid.Parse(req.ID)
	if err != nil {
		return "Error Update User", fmt.Errorf("invalid UUID: %v", err)
	}

	query := `UPDATE users SET 
		username = $1, 
		password = $2, 
		photo = $3,
		location = $4,
		updated_by = $5,
		updated_at = NOW() 
		WHERE id = $6 RETURNING id`

	result, err := u.db.Exec(context.Background(), query, req.Username, req.Password, req.Photo, req.Blocation, req.UpdatedBy, req.ID)
	if err != nil {
		return "Error Update User", err
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("user not found")
	}

	return req.ID, nil
}

func (b *userRepo) DeleteUser(c context.Context, req *models.IdRequest) (resp string, err error) {

	query := `
	 	UPDATE "users" 
		SET 
			"deleted_at" = NOW() 
		WHERE 
			"id" = $1
`
	result, err := b.db.Exec(
		context.Background(),
		query,
		req.Id,
	)
	if err != nil {
		return "", err
	}

	if result.RowsAffected() == 0 {
		return "", fmt.Errorf("user with ID %s not found", req.Id)

	}

	return req.Id, nil
}
