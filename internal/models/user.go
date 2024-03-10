package models

// TODO: strip user data from session information.
type User struct {
	Email         string `json:"email"`
	PasswordHash  string `json:"password"`
	Name          string `json:"name"`
	Phone         string `json:"phone"`
	BirthDate     string `json:"bDate"`
	Cookie        string `json:"session"`
	CookieExpires int64
}
