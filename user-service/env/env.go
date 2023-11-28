package env

type User struct {
	ID       string `json:"id"`
	LastName string `json:"Lastname"`
	UserName string `json:"Username"`
}

type Users = []User


