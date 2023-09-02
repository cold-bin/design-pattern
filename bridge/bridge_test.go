// @author cold bin
// @date 2023/9/2

package bridge

import "testing"

func TestAlert(t *testing.T) {
	type fields struct {
		method AlertMethod
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "email alert",
			fields: fields{method: &EmailAlert{}},
			args:   args{message: "服务器负载较高"},
		},
		{
			name:   "sms alert",
			fields: fields{method: &SMSAlert{}},
			args:   args{message: "数据库连接失败"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &ErrorAlert{
				method: tt.fields.method,
			}
			e.Alert(tt.args.message)
		})
	}
}
