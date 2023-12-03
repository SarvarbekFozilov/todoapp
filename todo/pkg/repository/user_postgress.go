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
func (r *AuthPostgress) GetUserById(req *models.IdRequest) (rep models.CreateUser, err error) {
	var user models.CreateUser

	query := fmt.Sprintf("SELECT username, password_hash, photo, birth_location FROM %s WHERE id=$1", usersTable)
	err = r.db.Get(&user, query, req.Id)

	return user, err
}
func (r *AuthPostgress) GetAllUsers(req *models.GetAllUserRequest) (rep models.GetAllUser, err error) {
	params := make(map[string]interface{})
	filter := ""
	offset := (req.Page - 1) * req.Limit

	// Set up the filter condition based on the request parameters
	if req.Name != "" {
		filter = "WHERE username ILIKE '%' || :filter || '%'"
		params["filter"] = req.Name
	}

	// Construct the SQL query
	query := fmt.Sprintf("SELECT id, username, password_hash, photo, birth_location FROM %s %s ORDER BY id LIMIT :limit OFFSET :offset", usersTable, filter)
	params["limit"] = req.Limit
	params["offset"] = offset

	// Execute the SQL query
	rows, err := r.db.NamedQuery(query, params)
	if err != nil {
		return rep, err
	}
	defer rows.Close()

	// Iterate over the result rows and populate the response
	for rows.Next() {
		var user models.CreateUser
		err := rows.StructScan(&user)
		if err != nil {
			return rep, err
		}
		rep.Users = append(rep.Users, user)
	}

	return rep, nil
}
func (r *AuthPostgress) UpdateUser(req *models.User) (string, error) {
	query := fmt.Sprintf("UPDATE %s SET username = $1, password_hash = $2, photo = $3, birth_location = $4 WHERE id = $5", usersTable)

	_, err := r.db.Exec(query, req.Username, req.Password, req.Photo, req.BirthLocation, req.ID)
	if err != nil {
		return "Error Update User", err
	}

	return req.ID, nil
}

func (r *AuthPostgress) DeleteUser(req *models.IdRequest) (string, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE id = $1", usersTable)

	_, err := r.db.Exec(query, req.Id)
	if err != nil {
		return "Error Delete User", err
	}

	return req.Id, nil
}
func (r *AuthPostgress) CreateUsers(users []models.CreateUser) ([]int, error) {
	ids := make([]int, len(users))

	stmt := fmt.Sprintf("INSERT INTO %s (username, password_hash, photo, birth_location) VALUES ", usersTable)

	values := []interface{}{}
	for i, user := range users {
		stmt += fmt.Sprintf("($%d, $%d, $%d, $%d),", i*4+1, i*4+2, i*4+3, i*4+4)
		values = append(values, user.Username, user.Password, user.Photo, user.Birthlocation)
	}

	stmt = stmt[:len(stmt)-1]

	query := fmt.Sprintf("%s RETURNING id", stmt)

	rows, err := r.db.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	index := 0
	for rows.Next() {
		err := rows.Scan(&ids[index])
		if err != nil {
			return nil, err
		}
		index++
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return ids, nil
}
func (r *AuthPostgress) UpdateUsers(users []models.User) (string, error) {
	stmt := fmt.Sprintf("UPDATE %s SET username = v.username, password_hash = v.password_hash, photo = v.photo, birth_location = v.birth_location FROM (VALUES ", usersTable)

	values := []interface{}{}
	for i, user := range users {
		stmt += fmt.Sprintf("($%d, $%d, $%d, $%d),", i*4+1, i*4+2, i*4+3, i*4+4)
		values = append(values, user.ID, user.Username, user.Password, user.Photo, user.BirthLocation)
	}

	stmt = stmt[:len(stmt)-1] // Remove the trailing comma

	stmt += ") AS v(id, username, password_hash, photo, birth_location) WHERE %s.id = v.id"

	query := fmt.Sprintf(stmt, usersTable)

	_, err := r.db.Exec(query, values...)
	if err != nil {
		return "Error Update Users", err
	}

	return "Users updated successfully", nil
}
