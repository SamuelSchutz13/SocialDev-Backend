package entity

type UserResponse struct {
	UserID   string `json:"user_id"`
	GoogleID string `json:"google_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Bio      string `json:"bio"`
	Github   string `json:"github"`
	Linkedin string `json:"linkedin"`
	Website  string `json:"website"`
}
