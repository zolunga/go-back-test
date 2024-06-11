package models

type Secret struct {
	Mongouri string `json:"mongouri"`
	Username string `json:"username"`
	Password string `json:"password"`
	Jwtsign  string `json:"jwtsign"`
	Database string `json:"database"`
}
