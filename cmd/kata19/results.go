package main

import (
	"fmt"
	"strings"

	"github.com/arvenil/kata/format/flags"
)

// Results
type Results []result

type result struct {
	Start string   `json:",omitempty"`
	End   string   `json:",omitempty"`
	Words []string `json:",omitempty"`
	Err   string   `json:",omitempty"`
}

func (r *Results) Append(start, end string, words []string, err error) {
	var errString string
	if err != nil {
		errString = fmt.Sprintf("%s", err)
	}
	*r = append(*r, result{
		Start: start,
		End:   end,
		Words: words,
		Err:   errString,
	})
}

// String representation of the list.
func (r Results) String() string {
	header := []string{
		"Start",
		"End",
		"Words chain",
		"Error (if any)",
	}
	field := []string{
		"{{.Start}}",
		"{{.End}}",
		"{{.Words}}",
		"{{.Err}}",
	}
	headers := strings.Join(header, "\t")
	fields := strings.Join(field, "\t")
	f := headers + "\n{{range .}}" + fields + "\n{{end}}"
	return flags.Format.Parse(f, r)
}
