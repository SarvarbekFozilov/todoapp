package todo

type CreateUserReq struct {
	ID        int    `json:"id" db:"id"`
	FullName  string `json:"fullname"`
	Username  string `json:"username" db:"username" binding:"required"`
	Password  string `json:"password"  db:"password_hash" binding:"required"`
	Photo     string `json:"photo" db:"photo" binding:"required" `
	Birthday  string `json:"birthday" db:"birthday" binding:"required" `
	Location  string `json:"location" db:"location" binding:"required" `
	CreatedBy int    `json:"created_by" db:"created_by" `
}

type IdRequest struct {
	ID int `json:"-" db:"id"`
}
type GetAllUserReq struct {
	Page   int      `json:"page"`
	Limit  int      `json:"limit"`
	Search string   `json:"search"`
	Fields []string `json:"fields"`
}

type GetAllUser struct {
	Users []User `json:"user"`
	Count int    `json:"count"`
}
type User struct {
	ID        int    `json:"id,omitempty"`
	FullName  string `json:"fullname,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	Photo     string `json:"photo,omitempty"`
	Birthday  string `json:"birthday,omitempty"`
	Location  string `json:"location,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
	CreatedBy int    `json:"created_by,omitempty"`
	UpdatedAt string `json:"updated_at,omitempty"`
	UpdatedBy int    `json:"updated_by,omitempty"`
	DeletedAt string `json:"deleted_at,omitempty"`
	DeletedBy int    `json:"deleted_by,omitempty"`
}

type UpdateUser struct {
	ID        int    `json:"id"`
	FullName  string `json:"fullname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Photo     string `json:"photo"`
	Birthday  string `json:"birthday"`
	Location  string `json:"location"`
	UpdatedBy int    `json:"updated_by"`
}
type UserResponse struct {
	ID       int    `json:"id"`
	FullName string `json:"fullname"`
	Username string `json:"username"`
	Photo    string `json:"photo"`
	Birthday string `json:"birthday"`
	Location string `json:"location"`
}
type SignIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
