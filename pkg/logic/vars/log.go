package logic

import "time"

type Log struct {
	Database  string      `json:"database"`
	Mode      string      `json:"mode"`
	FileName  string      `json:"fileName" bson:"fileName"`
	Result    string      `json:"result"`
	Error     interface{} `json:"error,omitempty"`
	CreatedAt time.Time   `json:"createdAt" bson:"createdAt"`
	UserName  string      `json:"userName" bson:"userName"`
}
