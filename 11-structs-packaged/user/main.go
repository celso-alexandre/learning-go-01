package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func (u *User) PrintData() {
	fmt.Println()
	fmt.Println("First name:", u.firstName)
	fmt.Println("Last name:", u.lastName)
	fmt.Println("Birth date:", u.birthDate)
	fmt.Println("Created at:", u.createdAt)
}

func (u *User) Clear() {
	u.firstName = ""
	u.lastName = ""
}

type Admin struct {
	email    string
	password string
	User     //inheritance
}

func (a *Admin) PrintData() {
	fmt.Println()
	a.User.PrintData()
	fmt.Println("Email:", a.email)
}

func New(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("first name, last name, and birth date are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(firstName, lastName, birthDate, email, password string) (*Admin, error) {
	if email == "" || password == "" {
		return nil, errors.New("email and password are required")
	}

	user, err := New(firstName, lastName, birthDate)
	if err != nil {
		return nil, err
	}

	return &Admin{
		email:    email,
		password: password,
		User:     *user,
	}, nil
}
