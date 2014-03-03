package tweetsaver

import (
	"encoding/json"
	"io"
)

// EncodePretty will encode a value v into pretty printed/indented
// JSON to the Writer given in dest.
func EncodePretty(dest io.Writer, v interface{}) error {
	data, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return err
	}
	_, err = dest.Write(data)
	return err
}
