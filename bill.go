package qbo

import (
	"context"
)

type BillService interface {
	Create(context.Context, *Bill) (*Bill, *Response, error)
	Read(context.Context, string) (*Bill, *Response, error)
	Update(context.Context, *Bill) (*Bill, *Response, error)
	Query(context.Context, string) ([]Bill, *Response, error)
}

type BillServiceClient struct {
	client *Client
}

var _ BillService = &BillServiceClient{}

type Bill struct {
	DueDate     string      `json:"DueDate"`
	Balance     int         `json:"Balance"`
	Domain      string      `json:"domain"`
	Sparse      bool        `json:"sparse"`
	ID          string      `json:"Id"`
	SyncToken   string      `json:"SyncToken"`
	MetaData    MetaData    `json:"MetaData"`
	TxnDate     string      `json:"TxnDate"`
	CurrencyRef CurrencyRef `json:"CurrencyRef"`
	PrivateNote string      `json:"PrivateNote"`
	LinkedTxn   []struct {
		TxnID   string `json:"TxnId"`
		TxnType string `json:"TxnType"`
	} `json:"LinkedTxn"`
	Line []struct {
		ID                            string  `json:"Id"`
		LineNum                       int     `json:"LineNum"`
		Amount                        float64 `json:"Amount"`
		DetailType                    string  `json:"DetailType"`
		AccountBasedExpenseLineDetail struct {
			AccountRef struct {
				Value string `json:"value"`
				Name  string `json:"name"`
			} `json:"AccountRef"`
			BillableStatus string `json:"BillableStatus"`
			TaxCodeRef     struct {
				Value string `json:"value"`
			} `json:"TaxCodeRef"`
		} `json:"AccountBasedExpenseLineDetail"`
	} `json:"Line"`
	VendorRef struct {
		Value string `json:"value"`
		Name  string `json:"name"`
	} `json:"VendorRef"`
	APAccountRef struct {
		Value string `json:"value"`
		Name  string `json:"name"`
	} `json:"APAccountRef"`
	TotalAmt float64 `json:"TotalAmt"`
}

func (a *BillServiceClient) Read(ctx context.Context, entityId string) (*Bill, *Response, error) {
	panic("implement me")
}

func (a *BillServiceClient) Create(ctx context.Context, account *Bill) (*Bill, *Response, error) {
	panic("implement me")
}

func (a *BillServiceClient) Update(context.Context, *Bill) (*Bill, *Response, error) {
	panic("implement me")
}

func (a *BillServiceClient) Query(context.Context, string) ([]Bill, *Response, error) {
	panic("implement me")
}
