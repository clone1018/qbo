package qbo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestAccountServiceClient_Read(t *testing.T) {
	setup()
	defer teardown()

	account := &Account{
		Name:        "Accounts Payable",
		AccountType: "income",
	}

	mux.HandleFunc("/account/12", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, http.MethodGet)
		resp, _ := json.Marshal(account)
		fmt.Fprint(w, fmt.Sprintf(`{"Account":%s}`, string(resp)))
	})

	resp, _, err := client.Account.Read(ctx, "12")
	if err != nil {
		t.Errorf("Account.Read returned error: %v", err)
	}

	if !reflect.DeepEqual(resp, account) {
		t.Errorf("Account.Read returned %+v, expected %+v", resp, account)
	}
}
