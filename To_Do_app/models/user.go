package models

type CreateUser struct {
	Username  string `json:"username"`
	Password  string `json:"password"`
	Photo     string `json:"photo"`
	Blocation string `json:"blocation"`
}

type User struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Photo     string `json:"photo"`
	Blocation string `json:"blocation"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at"`
	UpdatedBy string `json:"updated_by"`
	DeletedBy string `json:"deleted_by"`
}

type UpdateUser struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Photo     string `json:"photo"`
	Blocation string `json:"blocation"`
	UpdatedBy string `json:"updated_by"`
}
type IdRequest struct {
	Id string `json:"id"`
}

type GetAllUserRequest struct {
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
	UserName string `json:"username"`
}

type GetAllUser struct {
	Users []User `json:"user"`
	Count int    `json:"count"`
}
