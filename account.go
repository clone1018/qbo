package qbo

import (
	"context"
	"net/http"
)

type AccountService interface {
	Create(context.Context, *Account) (*Account, *Response, error)
	Read(context.Context, string) (*Account, *Response, error)
	Update(context.Context, *Account) (*Account, *Response, error)
	Query(context.Context, string) ([]Account, *Response, error)
}

type AccountServiceClient struct {
	client *Client
}

var _ AccountService = &AccountServiceClient{}

type Account struct {
	Name                          string      `json:"Name"`
	SubAccount                    bool        `json:"SubAccount"`
	ParentRef                     ParentRef   `json:"ParentRef"`
	FullyQualifiedName            string      `json:"FullyQualifiedName"`
	Active                        bool        `json:"Active"`
	Classification                string      `json:"Classification"`
	AccountType                   string      `json:"AccountType"`
	AccountSubType                string      `json:"AccountSubType"`
	AcctNum                       string      `json:"AcctNum"`
	CurrentBalance                int         `json:"CurrentBalance"`
	CurrentBalanceWithSubAccounts int         `json:"CurrentBalanceWithSubAccounts"`
	CurrencyRef                   CurrencyRef `json:"CurrencyRef"`
	Domain                        string      `json:"domain"`
	Sparse                        bool        `json:"sparse"`
	ID                            string      `json:"Id"`
	SyncToken                     string      `json:"SyncToken"`
	MetaData                      MetaData    `json:"MetaData"`
}

type accountRoot struct {
	Account *Account `json:"Account"`
}

func (a *AccountServiceClient) Read(ctx context.Context, entityId string) (*Account, *Response, error) {
	req, err := a.client.NewRequest(ctx, http.MethodGet, "/account/"+entityId, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(accountRoot)
	resp, err := a.client.Send(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Account, resp, err
}

func (a *AccountServiceClient) Create(ctx context.Context, account *Account) (*Account, *Response, error) {
	panic("implement me")
}

func (a *AccountServiceClient) Update(context.Context, *Account) (*Account, *Response, error) {
	panic("implement me")
}

func (a *AccountServiceClient) Query(ctx context.Context, searchQuery string) ([]Account, *Response, error) {
	panic("implement me")
	//return a.client.Query.Query(ctx, searchQuery)
}
