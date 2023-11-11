package types

type User struct{
	ID int
	Username string
	Password string
}

func ValidateUser(user User) bool{
	return true
}