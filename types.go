package yoobi

import (
	"encoding/json"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Employees []Employee

type Employee struct {
	Lastname       string    `json:"lastname"`
	EmployeeID     uuid.UUID `json:"employeeid"`
	EmployeeNumber string    `json:"employeenumber"`
	Self           URL       `json:"self"`
	State          string    `json:"state"`
	FirstName      string    `json:"firstname"`
	Infix          string    `json:"infix"`
}

type Bool bool

func (b Bool) MarshalJSON() ([]byte, error) {
	if b == true {
		return json.Marshal("1")
	}
	return json.Marshal("0")
}

type Department struct {
	Name string `json:"name"`
}

type User struct {
	State string `json:"state"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CustomFields []CustomField

type CustomField struct{}

type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format("02-01-2006"))
}

func (d *Date) UnmarshalJSON(data []byte) (err error) {
	var value string
	err = json.Unmarshal(data, &value)
	if err != nil {
		return err
	}

	if value == "" {
		return nil
	}

	// first try standard date
	d.Time, err = time.Parse(time.RFC3339, value)
	if err == nil {
		return nil
	}

	// try iso8601 date format
	d.Time, err = time.Parse("2006-01-02", value)
	if err == nil {
		return nil
	}

	// try yoobi date format
	d.Time, err = time.Parse("02-01-2006", value)
	return err
}
