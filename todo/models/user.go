package todo

type CreateUser struct {
	Id            int    `json:"-" db:"id"`
	Username      string `json:"username" db:"username" binding:"required"`
	Password      string `json:"password"  db:"password_hash" binding:"required"`
	Photo         string `json:"photo" db:"photo" binding:"required" `
	Birthlocation string `json:"birth_location" db:"birth_location" binding:"required" `
}

type IdRequest struct {
	Id string `json:"-" db:"id"`
}
type GetAllUserRequest struct {
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
	Name  string `json:"name"`
}

type GetAllUser struct {
	Users []CreateUser `json:"user"`
	Count int          `json:"count"`
}
type User struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	Photo         string `json:"photo"`
	BirthLocation string `json:"birth_location"`
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	UpdatedBy     string `json:"updated_by"`
}
