package format_test

import (
	"testing"

	"github.com/arvenil/kata/template/format"
)

func TestFormat_Parse(t *testing.T) {
	t.Parallel()

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
		{"json", fields{Text: "", JSON: true}, args{"", &format.Format{}}, "{\"Text\":\"\",\"JSON\":false}"},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			f := format.Format{
				Text: tt.fields.Text,
				JSON: tt.fields.JSON,
			}
			if got := f.Parse(tt.args.text, tt.args.data); got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
