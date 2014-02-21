package tweetsave

import (
	"encoding/json"
	"io"
)

func EncodePretty(v interface{}, dest io.Writer) error {
	data, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return err
	}
	_, err = dest.Write(data)
	return err
}
