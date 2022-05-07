package helpers

import (
	"encoding/json"
	"io"
)

func DecodeJSON(r io.Reader, v any) {
	json.NewDecoder(r).Decode(v)
}
