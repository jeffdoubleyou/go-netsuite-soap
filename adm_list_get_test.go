package netsuite_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestAdmListGet(t *testing.T) {
	req := client.NewAdmListGetRequest()
	req.QueryParams().PageSize = 20
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
