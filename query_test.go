package qbo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestQueryServiceClient_Query(t *testing.T) {
	setup()
	defer teardown()

	results := []Account{
		{
			Name:        "Accounts Payable",
			AccountType: "income",
		},
		{
			Name:        "Operations",
			AccountType: "income",
		},
	}

	mux.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		resp, _ := json.Marshal(results)
		fmt.Fprint(w, fmt.Sprintf(`{"QueryResponse": { "Account":%s} }`, string(resp)))
	})

	resp, _, err := client.Query.Query(ctx, "select * from Account")
	if err != nil {
		t.Errorf("Account.Read returned error: %v", err)
	}

	if !reflect.DeepEqual(resp, results) {
		t.Errorf("Account.Read returned %+v, expected %+v", resp, results)
	}
}
