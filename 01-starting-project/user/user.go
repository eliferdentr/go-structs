package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

type Admin struct {
	email    string
	password string
	User
}

func New(userFirstName string, userLastName string, userBirthdate string) (*User, error) {
	if userFirstName == "" || userLastName == "" || userBirthdate == "" {
		return nil, errors.New("first name, last name and birth date are required")
	}

	return &User{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}, nil

}

func NewAdmin(email string, password string) (*Admin, error){
	return &Admin{
		email: email,
		password: password,
		User : User{
			firstName: "ADMIN",
			lastName:  "ADMIN",
			birthdate: "---",
			createdAt: time.Now(),
		},
	}, nil

}

func (u *User) OutputUserDetails() {
	//(*User).firstName
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserNameAndSurname() {
	//(*User).firstName //is valid too
	u.firstName = ""
	u.lastName = ""
}
