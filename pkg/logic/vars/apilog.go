package logic

import "time"

func main() {

}

type apiLog struct {
	Database  string      `json:"database,omitempty"`
	Mode      string      `json:"mode,omitempty"`
	FileName  []byte      `json:"fileName,omitempty"`
	Result    string      `json:"result,omitempty"`
	Error     interface{} `json:"error,omitempty"`
	CreatedAt time.Time   `json:"createdAt,omitempty"`
	// UserName  string      `json:"userName,omitempty"`
}
