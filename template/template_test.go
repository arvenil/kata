package template_test

import (
	"bytes"
	"testing"

	"github.com/arvenil/kata/template"
)

func TestNewParse(t *testing.T) {
	t.Parallel()

	tm, err := template.New("foo", "this is a {{ . }}")
	if err != nil {
		t.Errorf("New() error = %v", err)

		return
	}

	var got bytes.Buffer
	if err = tm.Execute(&got, "string"); err != nil {
		t.Errorf("New() error = %v", err)

		return
	}

	if want := "this is a string"; got.String() != want {
		t.Errorf("Execute() got = %v, want %v", got.String(), want)
	}
}
