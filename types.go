package yoobi

import (
	"encoding/json"
	"strconv"
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

type EmployeeContracts []EmployeeContract

type EmployeeContract struct {
	PersNum    string `json:"pers_num"`
	ContractNr string `json:"contractnr"`
	StartDatum *Date  `json:"startdatum"`
	EindDatum  *Date  `json:"einddatum"`
	Type       string `json:"type"`
	FTE        int    `json:"fte"`
	Notitie    string `json:"notitie"`
	WeekAantal int    `json:"week_aantal"`
	Maandag1   int    `json:"maandag1"`
	Dinsdag1   int    `json:"dinsdag1"`
	Woensdag1  int    `json:"woensdag1"`
	Donderdag1 int    `json:"donderdag1"`
	Vrijdag1   int    `json:"vrijdag1"`
	Zaterdag1  int    `json:"zaterdag1"`
	Zondag1    int    `json:"zondag1"`
}

type Bool bool

func (b Bool) MarshalJSON() ([]byte, error) {
	if b == true {
		return json.Marshal("1")
	}
	return json.Marshal("0")
}

type Int int

func (i *Int) UnmarshalJSON(data []byte) error {
	realInt := 0
	err := json.Unmarshal(data, &realInt)
	if err == nil {
		*i = Int(realInt)
		return nil
	}

	// error, so maybe it isn't an int
	str := ""
	err = json.Unmarshal(data, &str)
	if err != nil {
		return err
	}

	if str == "" {
		*i = 0
		return nil
	}

	realInt, err = strconv.Atoi(str)
	if err != nil {
		return err
	}

	i2 := Int(realInt)
	*i = i2
	return nil
}

func (i Int) MarshalJSON() ([]byte, error) {
	return json.Marshal(int(i))
}

type Department struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
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
	if d.IsZero() {
		return json.Marshal("")
	}
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

type Projects []Project

type Project struct {
	ProjectID      uuid.UUID `json:"projectid"`
	ProjectCode    string    `json:"projectcode"`
	Name           string    `json:"name"`
	Label          string    `json:"label"`
	Classification string    `json:"classification"`
	State          string    `json:"state"`
	EndDate        Date      `json:"enddate"`
	StartDate      Date      `json:"startdate"`
	Self           URL       `json:"self"`
}

type Customers []Customer

type Customer struct {
	CustomerID   uuid.UUID `json:"customerid"`
	CustomerCode string    `json:"customercode"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	State        string    `json:"state"`
	Self         URL       `json:"self"`
}
