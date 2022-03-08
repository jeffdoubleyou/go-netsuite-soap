package netsuite_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestAccountSearch(t *testing.T) {
	req := client.NewAccountSearchRequest()
	req.RequestBody().SearchRecord.Type = "listAcct:AccountSearch"
	req.RequestBody().SearchRecord.Basic.Number.Operator = "contains"
	req.RequestBody().SearchRecord.Basic.Number.SearchValue = "100"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
