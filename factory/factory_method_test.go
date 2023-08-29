// @author cold bin
// @date 2023/8/29

package factory

import (
	"reflect"
	"testing"
)

func TestNewConfigFactory(t *testing.T) {
	type args struct {
		typ string
	}
	tests := []struct {
		name string
		args args
		want ConfigFactory
	}{
		{
			name: "test yaml",
			args: args{typ: "yaml"},
			want: NewConfigFactory("yaml"),
		},
		{
			name: "test json",
			args: args{typ: "json"},
			want: NewConfigFactory("json"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfigFactory(tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfigFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
