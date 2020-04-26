package templates

import (
	"bytes"
	"testing"
)

func TestNewParse(t *testing.T) {
	tm, err := New("foo", "this is a {{ . }}")
	if err != nil {
		t.Errorf("New() error = %v", err)
		return
	}

	var b bytes.Buffer
	err = tm.Execute(&b, "string")
	if err != nil {
		t.Errorf("New() error = %v", err)
		return
	}
	got := b.String()
	want := "this is a string"
	if got != want {
		t.Errorf("Execute() got = %v, want %v", got, want)
	}
}
