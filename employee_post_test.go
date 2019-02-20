package yoobi_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/joho/godotenv/autoload"
	yoobi "github.com/omniboost/go-yoobi"
)

func TestEmployeesPostTest(t *testing.T) {
	username := os.Getenv("YOOBI_USERNAME")
	password := os.Getenv("YOOBI_PASSWORD")
	organisation := os.Getenv("YOOBI_ORGANISATION")

	client := yoobi.NewClient(nil, organisation, username, password)
	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)

	req := client.NewEmployeesPostRequest()

	dob := time.Date(19811, 2, 1, 0, 0, 0, 0, time.UTC)
	startreg := time.Date(2003, 5, 11, 0, 0, 0, 0, time.UTC)
	//startdate := time.Date(0000, 0, 0, 0, 0, 0, 0, time.UTC)

	// dept := yoobi.Department()
	// dept.Name = "Development"

	body := yoobi.EmployeesPostRequestBody{
		EmployeeNumber:        "129990013",
		BSN:                   "123456789",
		DateOfBirth:           &yoobi.Date{dob},
		FirstName:             "Firstnametest",
		Initials:              "FNT",
		Infix:                 "van",
		LastName:              "Lastnametest",
		Gender:                "male",
		EmployeeAbbr:          "TEST",
		JobName:               "SomeJobName",
		StartRegistrationDate: yoobi.Date{startreg},
		StartDate:             yoobi.Date{},
		EndDate:               yoobi.Date{},
		CostCenter:            "123",
		CostUnit:              "456",
		AllowOvertime:         true,
		//Department:            dept,
		State:        "active",
		User:         nil,
		CustomFields: nil,
	}
	req.SetRequestBody(body)

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
