package domain

type User struct {
	Id        int64    `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	FirstName string   `json:"firstname"`
	LastName  string   `json:"lastname"`
	Role      []string `json:"role"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}
