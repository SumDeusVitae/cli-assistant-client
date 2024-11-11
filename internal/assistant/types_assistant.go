package assistant

import "time"

type QuestionForm struct {
	Model   string `json:"model"`
	Request string `json:"request"`
}

type QuestionRespond struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Model     string    `json:"model"`
	Question  string    `json:"question"`
	Reply     struct {
		String string `json:"String"`
		Valid  bool   `json:"Valid"`
	} `json:"reply"`
	UserID string `json:"user_id"`
}
