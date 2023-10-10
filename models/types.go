package models

type User struct {
	Id       int
	UserName string
}

type Server struct {
	Users map[int]*User
	Cache map[int]*User
	DBHit int
}
