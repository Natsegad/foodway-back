package domain

type Config struct {
	Port string
	IP   string
}

type UserInfo struct {
	ID       uint32 `json:"user_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}
