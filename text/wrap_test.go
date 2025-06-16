package text

import (
	"reflect"
	"testing"
)

func TestWrap(t *testing.T) {

	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		text    string
		options []WrappingOption
		want    string
	}{
		{
			name:    "simple",
			text:    "Я в нашей любимой закусочной, но скоро вернусь.",
			options: []WrappingOption{LineMaxWidth(30)},
			want:    "Я в нашей любимой закусочной, \nно скоро вернусь.",
		}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Wrap(tt.text, tt.options...)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("Wrap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_joinAfter(t *testing.T) {
	type args struct {
		slice []string
		n     int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "", args: args{slice: []string{"111", "222", "333", "444", "555"}, n: -1}, want: []string{"111", "222", "333", "444", "555"}},
		{name: "", args: args{slice: []string{"111", "222", "333", "444", "555"}, n: 0}, want: []string{"111", "222", "333", "444", "555"}},
		{name: "", args: args{slice: []string{"111", "222", "333", "444", "555"}, n: 1}, want: []string{"111222333444555"}},
		{name: "", args: args{slice: []string{"111", "222", "333", "444", "555"}, n: 2}, want: []string{"111", "222333444555"}},
		{name: "", args: args{slice: []string{"111", "222", "333", "444", "555"}, n: 3}, want: []string{"111", "222", "333444555"}},
		{name: "", args: args{slice: []string{"111", "222", "333", "444", "555"}, n: 4}, want: []string{"111", "222", "333", "444555"}},
		{name: "", args: args{slice: []string{"111", "222", "333", "444", "555"}, n: 5}, want: []string{"111", "222", "333", "444", "555"}},
		{name: "", args: args{slice: []string{"111", "222", "333", "444", "555"}, n: 6}, want: []string{"111", "222", "333", "444", "555"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := joinAfter(tt.args.slice, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("joinAfter(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
