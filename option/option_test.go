// @author cold bin
// @date 2023/8/30

package option

import (
	"reflect"
	"testing"
)

func TestNewPool(t *testing.T) {
	type args struct {
		name string
		opts []Option
	}
	tests := []struct {
		name string
		args args
		want *Pool
	}{
		{
			name: "no option",
			args: args{name: "db_pool1"},
			want: &Pool{name: "db_pool1"},
		},
		{
			name: "only maxIdle option",
			args: args{
				name: "db_pool2",
				opts: []Option{WithMaxIdle(10)},
			},
			want: &Pool{name: "db_pool2", maxIdle: 10},
		},
		{
			name: "all of options",
			args: args{
				name: "db_pool3",
				opts: []Option{
					WithMaxIdle(10),
					WithMinIdle(1),
					WithMaxTotal(11),
				},
			},
			want: &Pool{name: "db_pool", maxTotal: 11, maxIdle: 10, minIdle: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPool(tt.args.name, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPool() = %v, want %v", got, tt.want)
			}
		})
	}
}
