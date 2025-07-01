package entity

type PostResponse struct {
	PostID    string `json:"post_id"`
	UserID    string `json:"user_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Photo     string `json:"photo"`
	Video     string `json:"video"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
