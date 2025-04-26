package controllers

import "github.com/felipebs10/curso-go/db"

type Subscription struct {
	Name  string
	Email string
}

func Create(name string, email string) error {
	s := Subscription{Name: name, Email: email}

	return db.Insert("newsletter", s)
}
