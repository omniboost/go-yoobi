package yoobi_test

import (
	"encoding/json"
	"log"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	yoobi "github.com/omniboost/go-yoobi"
)

func TestCustomersPostTest(t *testing.T) {
	username := os.Getenv("YOOBI_USERNAME")
	password := os.Getenv("YOOBI_PASSWORD")
	organisation := os.Getenv("YOOBI_ORGANISATION")

	client := yoobi.NewClient(nil, organisation, username, password)
	client.SetDebug(true)
	client.SetDisallowUnknownFields(true)

	req := client.NewCustomersPostRequest()
	req.SetRequestBody(yoobi.CustomersPostRequestBody{
		Name: "sagrosync test",
		Code: "12",
	})

	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
