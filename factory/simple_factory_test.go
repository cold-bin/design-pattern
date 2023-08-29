// @author cold bin
// @date 2023/8/28

package factory

import (
	"reflect"
	"testing"
)

func TestNewConfig(t *testing.T) {
	type args struct {
		name string
		typ  string
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{
			name: "test yaml",
			args: args{typ: "yaml"},
			want: NewConfig("test_yaml", "yaml"),
		}, {
			name: "test json",
			args: args{typ: "json"},
			want: NewConfig("test_json", "json"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewConfig(tt.args.name, tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
