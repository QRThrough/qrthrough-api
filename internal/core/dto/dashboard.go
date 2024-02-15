package dto

import "github.com/JMjirapat/qrthrough-api/internal/core/model"

type Configuration struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Desc  string `json:"desc"`
}

type AccountRequestBody struct {
	Firstname string     `json:"firstname"`
	Lastname  string     `json:"lastname"`
	Tel       string     `json:"tel"`
	IsActive  bool       `json:"is_active"`
	Role      model.Role `json:"role"`
}

type AllAccountsResponseBody struct {
	Count    int64           `json:"count"`
	Accounts []model.Account `json:"accounts"`
}

type AllLogsResponseBody struct {
	Count int64         `json:"count"`
	Logs  []model.Usage `json:"logs"`
}

type SignInResponseBody struct {
	Role model.Role `json:"role"`
}
