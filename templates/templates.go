package templates

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"
)

// basicFunctions are the set of initial
// functions provided to every template.
var basicFunctions = template.FuncMap{
	"json": func(v interface{}) string {
		buf := &bytes.Buffer{}
		enc := json.NewEncoder(buf)
		enc.SetEscapeHTML(false)
		err := enc.Encode(v)
		if err != nil {
		}
		// Remove the trailing new line added by the encoder
		return strings.TrimSpace(buf.String())
	},
	"split": strings.Split,
	"join":  strings.Join,
	"title": strings.Title,
	"lower": strings.ToLower,
	"upper": strings.ToUpper,
}

// New creates a new tagged template with the basic functions
// and parses the given format.
func New(tag, format string) (*template.Template, error) {
	return template.New(tag).Funcs(basicFunctions).Parse(format)
}
