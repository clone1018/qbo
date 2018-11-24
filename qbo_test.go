package qbo

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

var (
	mux *http.ServeMux

	ctx = context.TODO()

	client *Client

	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)

	client = NewClient(nil, "sandbox")
	parsedUrl, _ := url.Parse(server.URL)
	client.BaseURL = parsedUrl
}

func teardown() {
	server.Close()
}

func testMethod(t *testing.T, r *http.Request, expected string) {
	if expected != r.Method {
		t.Errorf("Request method = %v, expected %v", r.Method, expected)
	}
}

func TestNewClient(t *testing.T) {

}

func TestNewRequest(t *testing.T) {
	setup()
	defer teardown()

	inURL, outURL := "/v3/foo", client.BaseURL.String()+"/v3/foo"
	inBody, outBody := &Account{Name: "l"},
		`{"Name":"l","SubAccount":false,"ParentRef":{"value":""},"FullyQualifiedName":"","Active":false,"Classification":"","AccountType":"","AccountSubType":"","AcctNum":"","CurrentBalance":0,"CurrentBalanceWithSubAccounts":0,"CurrencyRef":{"value":"","name":""},"domain":"","sparse":false,"Id":"","SyncToken":"","MetaData":{"CreateTime":"","LastUpdatedTime":""}}`+"\n"
	req, _ := client.NewRequest(ctx, http.MethodGet, inURL, inBody)

	// test relative URL was expanded
	if req.URL.String() != outURL {
		t.Errorf("NewRequest(%v) URL = %v, expected %v", inURL, req.URL, outURL)
	}

	// test body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if string(body) != outBody {
		t.Errorf("NewRequest(%v)Body = %v, expected %v", inBody, string(body), outBody)
	}

	// test default user-agent is attached to the request
	userAgent := req.Header.Get("User-Agent")
	if userAgent != "github.com/clone1018/qbo" {
		t.Errorf("NewRequest() User-Agent = %v, expected %v", userAgent, "github.com/clone1018/qbo")
	}
}
