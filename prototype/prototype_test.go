// @author cold bin
// @date 2023/8/31

package prototype

import (
	"reflect"
	"testing"
)

func TestKeywords_Clone(t *testing.T) {
	type args struct {
		updatedWords []*Keyword
	}
	tests := []struct {
		name  string
		words Keywords
		args  args
		want  Keywords
	}{
		{
			name:  "not update, but add a word",
			words: Keywords{"peking": &Keyword{"peking", 10000}},
			args:  args{updatedWords: []*Keyword{{Word: "pekingpeking", Visit: 1000000}}},
			want: Keywords{
				"peking":       &Keyword{"peking", 10000},
				"pekingpeking": &Keyword{Word: "pekingpeking", Visit: 1000000}},
		},
		{
			name:  "update a word, but not add a word",
			words: Keywords{"peking": &Keyword{"peking", 10000}},
			args:  args{updatedWords: []*Keyword{{Word: "peking", Visit: 1000000}}},
			want:  Keywords{"peking": &Keyword{"peking", 1000000}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.words.Clone(tt.args.updatedWords); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keywords.Clone() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
