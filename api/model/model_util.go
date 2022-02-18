package model

import (
	"bytes"
	"encoding/json"
)

func (d *DatumArr) UnmarshalJSON(b []byte) error {
	x := bytes.TrimLeft(b, " \t\r\n")

	isArray := len(x) > 0 && x[0] == '['

	if isArray {
		var data []Datum
		err := json.Unmarshal(b, &data)

		*d = data
		return err
	}

	var data Datum

	err := json.Unmarshal(b, &data)
	*d = DatumArr{data}
	return err
}
