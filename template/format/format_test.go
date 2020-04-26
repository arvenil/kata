package format

import (
	"testing"
)

func TestFormat_Parse(t *testing.T) {
	type fields struct {
		Text string
		JSON bool
	}
	type args struct {
		text string
		data interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{"json", fields{"", true}, args{"", &Format{}}, "{\"Text\":\"\",\"JSON\":false}"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Format{
				Text: tt.fields.Text,
				JSON: tt.fields.JSON,
			}
			if got := f.Parse(tt.args.text, tt.args.data); got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
