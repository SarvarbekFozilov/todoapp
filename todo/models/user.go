package todo

type CreateUser struct {
	Id            int    `json:"-" db:"id"`
	Username      string `json:"username" db:"username" binding:"required"`
	Password      string `json:"password"  db:"password_hash" binding:"required"`
	Photo         string `json:"photo" db:"photo" binding:"required" `
	Birthlocation string `json:"birth_location" db:"birth_location" binding:"required" `
}
