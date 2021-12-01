package format_test

import (
	"testing"

	"github.com/arvenil/kata/template/format"
)

func TestFormat_Parse(t *testing.T) {
	t.Parallel()

	type args struct {
		text string
		data interface{}
	}

	tests := []struct {
		name   string
		format format.Format
		args   args
		want   string
	}{
		{
			"text",
			format.Format{Text: "", JSON: false},
			args{"{{.JSON}}", &format.Format{
				Text: "",
				JSON: false,
			}},
			"false",
		},
		{
			"json",
			format.Format{Text: "", JSON: true},
			args{"", &format.Format{
				Text: "",
				JSON: false,
			}},
			"{\"Text\":\"\",\"JSON\":false}",
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if got := tt.format.Parse(tt.args.text, tt.args.data); got != tt.want {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
