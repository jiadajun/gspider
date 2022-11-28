package common

type DbStruct struct {
	User           string `json:"user"`
	Password       string `json:"password"`
	Host           string `json:"host"`
	Port           int64  `json:"port"`
	Database       string `json:"database"`
	ConnectTimeout string `json:"connect_timeout"`
	Charset        string `json:"charset"`
}
