package user

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Address  string `json:"address"`
}

type Log struct {
	Addr      string `json:"address"`
	Namespace string `json:"namespace"`
	Cause     string `json:"cause"`
}
