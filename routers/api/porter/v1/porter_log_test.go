package v1

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestLogs(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "no input name",
			args: args{
				c: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Logs(tt.args.c)
		})
	}
}
