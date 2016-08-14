package output

import (
	"encoding/json"
	"io"
)

func WriteFeatures(data *Stats, out io.Writer, oconf *OutputConfig) error {
	by, err := json.Marshal(data)
	if err != nil {
		return err
	}
	out.Write([]byte("["))
	out.Write(by)
	out.Write([]byte("]\n"))
	return nil
}
