package tools

import (
	"encoding/json"
	"io"
)

func FromJSON(i interface{}, r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(i)
}
