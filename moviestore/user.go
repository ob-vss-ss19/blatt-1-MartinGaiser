package moviestore

import "fmt"

type Age uint8
type UserID uint16

type User struct {
	Name string
	Age Age
	UserID UserID
}

func (user *User)String() string{
	return fmt.Sprintf("%4d. %s, %d",user.UserID,user.Name,user.Age)
}
