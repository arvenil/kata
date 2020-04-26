/*
Package format provides configuration for output formatting.
*/
package format

import (
	"github.com/arvenil/kata/templates"
)

// Format contains configuration required to Template data.
type Format struct {
	Template string
	JSON     bool
}

// Parse Template and fill it with data.
func (f Format) Parse(template string, data interface{}) string {
	if f.Template != "" {
		template = f.Template
	}
	if f.JSON {
		template = "{{ json . }}"
	}

	return templates.Parse(template, data)
}
