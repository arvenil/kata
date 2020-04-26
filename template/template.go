package template

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

// New returns a new named template.Text with the basic functions and parses text as template.
func New(name, text string) (*template.Template, error) {
	return template.New(name).Funcs(basicFunctions).Parse(text)
}
