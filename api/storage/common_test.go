package storage

import (
	"encoding/json"
	"reflect"
	"testing"
)

func Test_rawListSplitAndSort(t *testing.T) {
	type args struct {
		s json.RawMessage
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "empty",
			args: args{
				s: json.RawMessage(`""`),
			},
			want: []string{},
		},
		{
			name: "foo bar",
			args: args{
				s: json.RawMessage(`"foo,bar"`),
			},
			want: []string{"bar", "foo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rawListSplitAndSort(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rawListSplitAndSort() = %v, want %v", got, tt.want)
				t.Errorf("%T", got)
				t.Errorf("%T", tt.want)
			}
		})
	}
}

func Test_listJoin(t *testing.T) {
	type args struct {
		l         *[]string
		separator string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty",
			args: args{
				l:         &[]string{},
				separator: ",",
			},
			want: "",
		},
		{
			name: "foo bar",
			args: args{
				l:         &[]string{"foo", "bar"},
				separator: ",",
			},
			want: "foo,bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringSliceJoin(tt.args.l, tt.args.separator); got != tt.want {
				t.Errorf("listJoin() = %v, want %v", got, tt.want)
			}
		})
	}
}
