package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	models "todo/models"

	"github.com/jmoiron/sqlx"
)

type UserPostgress struct {
	db *sqlx.DB
}

func NewUserPostgress(db *sqlx.DB) *UserPostgress {
	return &UserPostgress{db: db}
}

func (r *UserPostgress) CreateUser(user *models.CreateUserReq) (int, error) {
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

	row := r.db.QueryRow(query,
		user.FullName,
		user.Username,
		user.Password,
		user.Photo,
		user.Birthday,
		user.Location)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (r *UserPostgress) GetUserById(req *models.IdRequest) (*models.UserResponse, error) {
	var user models.UserResponse
	query := `SELECT
	                id ,
					fullname,
					username,
					photo,
					birthday,
					location					
				FROM users WHERE id=$1 AND deleted_at IS NULL`

	err := r.db.Get(&user, query, req.ID)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r UserPostgress) GetAllUsers(req *models.GetAllUserReq) (rep models.GetAllUser, err error) {
	params := make(map[string]interface{})
	offset := (req.Page - 1) * req.Limit
	params["limit"] = req.Limit
	params["offset"] = offset

	query := `
	  SELECT 
	         id,
			 fullname,
			 username,
			 photo,
			 birthday,
			 location
		FROM users WHERE deleted_at IS NULL`

	if req.Search != "" {
		filter := "AND username ILIKE '%' || :filter || '%'"
		params["filter"] = req.Search
		query += filter
	}

	// Execute the SQL query
	rows, err := r.db.NamedQuery(query, params)
	if err != nil {
		return rep, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the response
	for rows.Next() {
		var user models.CreateUserReq
		err := rows.StructScan(&user)
		if err != nil {
			return rep, err
		}
		rep.Users = append(rep.Users, user)
	}

	return rep, nil
}

func (r *UserPostgress) UpdateUser(req *models.UpdateUser) (int, error) {
	query := `
	   UPDATE  users SET
	            fullname=$2,
				username=$3,
				password_hash=$4,
				photo=$5,
				birthday=$6,
				location=$7,
				updated_at=NOW(),
				updated_by=$8
		WHERE id=$1;`

	resp, err := r.db.Exec(query,
		req.ID,
		req.FullName,
		req.Username,
		req.Password,
		req.Photo,
		req.Birthday,
		req.Location,
		req.UpdatedBy,
	)

	if err != nil {
		return 0, err
	}

	if res, _ := resp.RowsAffected(); res == 0 {
		return 0, sql.ErrNoRows
	}

	return req.ID, nil
}

func (r *UserPostgress) DeleteUser(req *models.IdRequest) (int, error) {
	query := `
		UPDATE users 
		SET 
			deleted_at = NOW(),
			deleted_by=$2 
		WHERE id=$1;`

	_, err := r.db.Exec(query, req.ID)
	if err != nil {
		return 0, err
	}

	return req.ID, nil
}

func GenerateInsertQuery(dataSlice []models.CreateUserReq, tableName string) string {
	query := fmt.Sprintf("INSERT INTO %s (fullname, username, password_hash, photo, birthday, location, created_by) VALUES", tableName)
	valueStrings := []string{}
	for _, data := range dataSlice {
		fullname := data.FullName
		username := data.Username
		password := data.Password
		photo := data.Photo
		birthday := data.Birthday
		location := data.Location
		created_by := data.CreatedBy
		values := fmt.Sprintf("('%s', '%s', '%s', '%s', '%s', '%s', %d)", fullname, username, password, photo, birthday, location, created_by)
		valueStrings = append(valueStrings, values)
	}
	query += "\n" + strings.Join(valueStrings, ",\n") + ";"
	return query
}
func (r *UserPostgress) CreateUsers(users []models.CreateUserReq) ([]int, error) {
	ids := make([]int, 0, len(users))

	query := GenerateInsertQuery(users, "users")

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}

		ids = append(ids, id)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ids, nil
}

func (r *UserPostgress) UpdateUsers(users []models.UpdateUser) (string, error) {
	query := `UPDATE users AS u
			  SET fullname = c.fullname,
				  username = c.username,
				  password_hash = c.password_hash,
				  photo = c.photo,
				  birthday = c.birthday::date,
				  location = c.location,
				  updated_by = c.updated_by::integer
			  FROM (VALUES `

	valueArgs := []interface{}{}

	for i, data := range users {
		id := data.ID
		fullname := data.FullName
		username := data.Username
		password := data.Password
		photo := data.Photo
		birthday := data.Birthday
		location := data.Location
		updatedBy := data.UpdatedBy

		query += fmt.Sprintf(`($%d::integer, $%d, $%d, $%d, $%d, $%d, $%d, $%d)`, (i*8)+1, (i*8)+2, (i*8)+3, (i*8)+4, (i*8)+5, (i*8)+6, (i*8)+7, (i*8)+8)
		valueArgs = append(valueArgs, id, fullname, username, password, birthday, photo, location, updatedBy)

		if i != len(users)-1 {
			query += `, `
		}
	}

	query += `) AS c(id, fullname, username, password_hash, birthday, photo, location, updated_by)
			  WHERE u.id = c.id::integer`
	fmt.Println(query)
	result, err := r.db.Exec(query, valueArgs...)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}
	fmt.Println(rowsAffected)
	if rowsAffected != int64(len(users)) {
		return "", errors.New("not all users were updated")
	}

	return "Updated all", nil
}
