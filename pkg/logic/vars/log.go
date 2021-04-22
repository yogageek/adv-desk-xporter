package logic

import "time"

type Log struct {
	Database  string      `json:"database"`
	Mode      string      `json:"mode"`
	FileName  string      `json:"fileName"`
	Result    string      `json:"result"`
	Error     interface{} `json:"error"`
	CreatedAt time.Time   `json:"createdAt"`
	// UserName  string      `json:"userName"`
}
