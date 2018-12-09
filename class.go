package qbo

import (
	"context"
)

type Class struct {
	Name               string   `json:"Name"`
	SubClass           bool     `json:"SubClass"`
	FullyQualifiedName string   `json:"FullyQualifiedName"`
	Active             bool     `json:"Active"`
	Domain             string   `json:"domain"`
	Sparse             bool     `json:"sparse"`
	ID                 string   `json:"Id"`
	SyncToken          string   `json:"SyncToken"`
	MetaData           MetaData `json:"MetaData"`
}

type ClassService interface {
	Create(context.Context, *Class) (*Class, *Response, error)
	Read(context.Context, string) (*Class, *Response, error)
	Update(context.Context, *Class) (*Class, *Response, error)
	Query(context.Context, string) ([]Class, *Response, error)
}

type ClassServiceClient struct {
	client *Client
}

func (ClassServiceClient) Create(context.Context, *Class) (*Class, *Response, error) {
	panic("implement me")
}

func (ClassServiceClient) Read(context.Context, string) (*Class, *Response, error) {
	panic("implement me")
}

func (ClassServiceClient) Update(context.Context, *Class) (*Class, *Response, error) {
	panic("implement me")
}

func (ClassServiceClient) Query(context.Context, string) ([]Class, *Response, error) {
	panic("implement me")
}

var _ ClassService = &ClassServiceClient{}
