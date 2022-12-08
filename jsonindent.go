package jsonindent

import (
	"bytes"
	"encoding/json"
)

func JsonIndent(j string) string {
	var buf bytes.Buffer
	err := json.Indent(&buf, []byte(j), "", "  ")
	if err != nil {
		return "json indent error"
	}
	return buf.String()
}
