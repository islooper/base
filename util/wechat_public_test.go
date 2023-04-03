package util

import (
	"fmt"
	"testing"
)

func TestSimpleWeChatPublicClient_CreateTextTask(t *testing.T) {
	type fields struct {
		AccessToken string
	}
	type args struct {
		touser  []string
		content string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Rsp
		wantErr bool
	}{
		{
			name:   "test1",
			fields: fields{"123"},
			args: args{
				touser:  []string{"123", "321"},
				content: "这是一条测试是否通行的测试",
			},
			want:    new(Rsp),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SimpleWeChatPublicClient{
				AccessToken: tt.fields.AccessToken,
			}
			got, err := s.CreateTextTask(tt.args.touser, tt.args.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTextTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}
