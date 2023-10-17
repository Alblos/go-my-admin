package internalDbTypes

type Connection struct {
	Id           int    `json:"id"`
	CommonName   string `json:"common_name"`
	DatabaseName string `json:"database_name"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	SslMode      string `json:"ssl_mode"`
}
