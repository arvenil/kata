package format

import (
	"github.com/arvenil/kata/format/templates"
)

// format contains configuration required to Template data.
type Format struct {
	Template string
	Json     bool
}

// Parse Template and fill it with data.
func (f Format) Parse(template string, data interface{}) string {
	if f.Template != "" {
		template = f.Template
	}
	if f.Json {
		template = "{{ json . }}"
	}

	return templates.Parse(template, data)
}
