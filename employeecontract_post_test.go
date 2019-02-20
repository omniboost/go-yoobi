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

func TestEmployeeContractsPostTest(t *testing.T) {
	username := os.Getenv("YOOBI_USERNAME")
	password := os.Getenv("YOOBI_PASSWORD")
	organisation := os.Getenv("YOOBI_ORGANISATION")

	client := yoobi.NewClient(nil, organisation, username, password)
	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)

	req := client.NewEmployeeContractsPostRequest()

	req.PathParams().Number = "129990013"

	datestart := time.Date(2019, 2, 1, 0, 0, 0, 0, time.UTC)
	dateend := time.Date(2019, 2, 1, 0, 0, 0, 0, time.UTC)

	body := yoobi.EmployeeContractsPostRequestBody{
		PersNum:    "129990013",
		ContractNr: "TEST0123",
		StartDatum: &yoobi.Date{datestart},
		EindDatum:  &yoobi.Date{dateend},
		Type:       "permanentemployment",
		FTE:        1,
		Notitie:    "testnotitie",
		WeekAantal: 1,
		Maandag1:   8,
		Dinsdag1:   8,
		Woensdag1:  8,
		Donderdag1: 8,
		Vrijdag1:   8,
		Zaterdag1:  0,
		Zondag1:    0,
	}
	req.SetRequestBody(body)

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
