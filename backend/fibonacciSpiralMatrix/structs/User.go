package structs

type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
}