// @author cold bin
// @date 2023/9/1

// Package proxy 来自 https://juejin.cn/post/6962375619774513183
package proxy

import (
	"context"
	"testing"
)

func TestUploaderProxy_UploadSingle(t *testing.T) {
	type fields struct {
		Up Uploader
	}
	type args struct {
		ctx   context.Context
		key   string
		value []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "qiniu oss",
			fields: fields{Up: &QiniuOss{}},
			args: args{
				ctx:   context.Background(),
				key:   "test.jpg",
				value: nil,
			},
		},
		{
			name:   "ali oss",
			fields: fields{Up: &AliOss{}},
			args: args{
				ctx:   context.Background(),
				key:   "test.jpg",
				value: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UploaderProxy{
				Up: tt.fields.Up,
			}
			u.UploadSingle(tt.args.ctx, tt.args.key, tt.args.value)
		})
	}
}
