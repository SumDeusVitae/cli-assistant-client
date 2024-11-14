package assistant

import "time"

type RegistrationForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    struct {
		Email string `json:"email"`
		Valid bool   `json:"valid"`
	} `json:"email"`
}

type UserRespond struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Login     string    `json:"login"`
	Email     struct {
		String string `json:"String"`
		Valid  bool   `json:"Valid"`
	} `json:"email"`
	APIKey string `json:"api_key"`
}

type LoginForm struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
