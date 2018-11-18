package qbo

type ParentRef struct {
	Value string `json:"value"`
}

type CurrencyRef struct {
	Value string `json:"value"`
	Name  string `json:"name"`
}

type MetaData struct {
	CreateTime      string `json:"CreateTime"`
	LastUpdatedTime string `json:"LastUpdatedTime"`
}
