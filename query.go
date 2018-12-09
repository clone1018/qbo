package qbo

import (
	"context"
	"net/http"
	"net/url"
)

type QueryService interface {
	Query(context.Context, string) (*QueryResponse, *Response, error)
}

type QueryServiceClient struct {
	client *Client
}

var _ QueryService = &QueryServiceClient{}

type QueryResponse struct {
	Account      []Account      `json:"Account"`
	Bill         []Bill         `json:"Bill"`
	Class        []Class        `json:"Class"`
	JournalEntry []JournalEntry `json:"JournalEntry"`

	StartPosition int `json:"startPosition"`
	MaxResults    int `json:"maxResults"`
}

type queryRoot struct {
	QueryResponse *QueryResponse `json:"QueryResponse"`
}

func (a *QueryServiceClient) Query(ctx context.Context, queryStr string) (*QueryResponse, *Response, error) {
	req, err := a.client.NewRequest(ctx, http.MethodGet, "/query?query="+url.QueryEscape(queryStr), nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(queryRoot)
	resp, err := a.client.Send(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.QueryResponse, resp, err
}
