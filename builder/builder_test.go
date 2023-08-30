// @author cold bin
// @date 2023/8/30

package builder

import (
	"reflect"
	"testing"
)

func TestBuild_Build(t *testing.T) {
	type fields struct {
		name     string
		maxTotal int
		maxIdle  int
		minIdle  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Pool
		wantErr bool
	}{
		{
			name:    "maxIdle is less than minIdle",
			fields:  fields{maxIdle: 1, minIdle: 9},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "maxIdle is more than maxTotal",
			fields:  fields{maxIdle: 9, maxTotal: 1},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "good case",
			fields:  fields{name: "good", maxIdle: 9, minIdle: 1, maxTotal: 10},
			want:    &Pool{name: "good", maxIdle: 9, minIdle: 1, maxTotal: 10},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &Build{
				name:     tt.fields.name,
				maxTotal: tt.fields.maxTotal,
				maxIdle:  tt.fields.maxIdle,
				minIdle:  tt.fields.minIdle,
			}
			got, err := b.Build()
			if (err != nil) != tt.wantErr {
				t.Errorf("Build.Build() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Build.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
