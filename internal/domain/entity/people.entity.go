package entity

import "time"

type PeopleEntity struct {
	ID             string
	TypePeople     string
	FirstName      string
	SecondName     string
	Surname        string
	LastSurname    string
	TypeDocument   string
	DocumentNumber string
	Birthdate      time.Time
	Phone          string
	Address        string
	Email          string
	Company        string
	IsActive       bool
}
