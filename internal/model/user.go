package model

import "time"

type UserInfo struct {
	Phone             string
	Email             string
	FirstName         string
	SecondName        string
	LastName          string
	PassportNumber    uint32
	PassportCode      uint8
	PassportIssueDate time.Time
	Birthday          time.Time
}

type User struct {
	ID   int
	UUID string
	Info *UserInfo
}
