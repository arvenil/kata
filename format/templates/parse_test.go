package templates

import (
	"testing"
)

func TestParse(t *testing.T) {
	got := Parse("this is a {{ . }}", "string")
	want := "this is a string"
	if got != want {
		t.Errorf("Parse() got = %v, want %v", got, want)
	}
}
