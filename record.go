package sheet

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"io"
)

type Records struct {
	body io.Reader
}

func NewRecords(body io.Reader) *Records {
	return &Records{body: body}
}

// Data (without schema interface) returns string map array.
// keys of map will be set spreadsheet header (currently first row),
// values of map will be set each spreadsheet record.
// note: each map value will be always set as string, even if actual value is number.
func (r *Records) Data() ([]map[string]interface{}, error) {
	var records []map[string]interface{}
	x := io.TeeReader(r.body, bytes.NewBuffer(nil))
	reader := csv.NewReader(x)
	headers, err := reader.Read()
	if err == io.EOF {
		return []map[string]interface{}{}, nil
	}
	if err != nil {
		return nil, err
	}
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		record := map[string]interface{}{}
		for i, v := range row {
			record[headers[i]] = v
		}
		records = append(records, record)
	}
	return records, nil
}

func (r *Records) DataTo(p interface{}) error {
	vs, err := r.Data()
	if err != nil {
		return err
	}
	bs, err := json.Marshal(vs)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, p)
}
