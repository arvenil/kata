/*
Package format provides configuration for output formatting.
*/
package format

import (
	"github.com/arvenil/kata/template"
)

// Format contains configuration required to template data.
type Format struct {
	Text string
	JSON bool
}

// Parse text as template and fill it with data.
func (f Format) Parse(text string, data interface{}) string {
	if f.Text != "" {
		text = f.Text
	}
	if f.JSON {
		text = "{{ json . }}"
	}

	return template.Parse(text, data)
}
