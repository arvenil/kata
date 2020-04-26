package template

import (
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		format string
		data   interface{}
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"template-ok", args{"this is a {{ . }}", "string"}, "this is a string"},
		{"template-err", args{"this is a {{ . }", "string"}, "template: :1: unexpected \"}\" in operand"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.format, tt.args.data); got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
