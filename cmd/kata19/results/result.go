package results

import (
	"fmt"
	"strings"

	"github.com/arvenil/kata/format"
	"github.com/arvenil/kata/format/flags"
)

func New() *results {
	return &results{
		format: format.Format{
			Template: flags.Template,
			Json:     flags.Json,
		},
	}
}

type Result struct {
	Start string   `json:",omitempty"`
	End   string   `json:",omitempty"`
	Words []string `json:",omitempty"`
	Err   string   `json:",omitempty"`
}

// results
type results struct {
	List   []Result
	format format.Format
}

func (r *results) Append(start, end string, words []string, err error) {
	var errString string
	if err != nil {
		errString = fmt.Sprintf("%s", err)
	}
	r.List = append(r.List, Result{
		Start: start,
		End:   end,
		Words: words,
		Err:   errString,
	})
}

// String representation of the list.
func (r *results) String() string {
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
	return r.format.Parse(f, r.List)
}
