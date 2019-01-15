package yoobi

import uuid "github.com/satori/go.uuid"

type Employees []Employee

type Employee struct {
	Lastname       string    `json:"lastname"`
	Employeeid     uuid.UUID `json:"employeeid"`
	Employeenumber string    `json:"employeenumber"`
	Self           URL       `json:"self"`
	State          string    `json:"state"`
	Firstname      string    `json:"firstname"`
	Infix          string    `json:"infix"`
}
