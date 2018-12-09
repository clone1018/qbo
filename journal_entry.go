package qbo

import (
	"context"
)

type JournalEntry struct {
	Adjustment  bool        `json:"Adjustment"`
	TotalAmt    float64     `json:"TotalAmt"`
	Domain      string      `json:"domain"`
	Sparse      bool        `json:"sparse"`
	ID          string      `json:"Id"`
	SyncToken   string      `json:"SyncToken"`
	MetaData    MetaData    `json:"MetaData"`
	DocNumber   string      `json:"DocNumber"`
	TxnDate     string      `json:"TxnDate"`
	CurrencyRef CurrencyRef `json:"CurrencyRef"`
	PrivateNote string      `json:"PrivateNote"`
	Line        []struct {
		ID                     string  `json:"Id"`
		Description            string  `json:"Description"`
		Amount                 float64 `json:"Amount"`
		DetailType             string  `json:"DetailType"`
		JournalEntryLineDetail struct {
			PostingType string `json:"PostingType"`
			Entity      struct {
				Type      string `json:"Type"`
				EntityRef struct {
					Value string `json:"value"`
					Name  string `json:"name"`
				} `json:"EntityRef"`
			} `json:"Entity"`
			AccountRef struct {
				Value string `json:"value"`
				Name  string `json:"name"`
			} `json:"AccountRef"`
			ClassRef struct {
				Value string `json:"value"`
				Name  string `json:"name"`
			} `json:"ClassRef"`
		} `json:"JournalEntryLineDetail"`
	} `json:"Line"`
	TxnTaxDetail struct {
	} `json:"TxnTaxDetail"`
}

type JournalEntryService interface {
	Create(context.Context, *Class) (*JournalEntry, *Response, error)
	Read(context.Context, string) (*JournalEntry, *Response, error)
	Update(context.Context, *Class) (*JournalEntry, *Response, error)
	Query(context.Context, string) ([]JournalEntry, *Response, error)
}

type JournalEntryServiceClient struct {
	client *Client
}

func (JournalEntryServiceClient) Create(context.Context, *Class) (*JournalEntry, *Response, error) {
	panic("implement me")
}

func (JournalEntryServiceClient) Read(context.Context, string) (*JournalEntry, *Response, error) {
	panic("implement me")
}

func (JournalEntryServiceClient) Update(context.Context, *Class) (*JournalEntry, *Response, error) {
	panic("implement me")
}

func (JournalEntryServiceClient) Query(context.Context, string) ([]JournalEntry, *Response, error) {
	panic("implement me")
}

var _ JournalEntryService = &JournalEntryServiceClient{}
