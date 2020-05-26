package sheet

import (
	"encoding/csv"
	"io"
)

type Records struct {
	body io.Reader
}

// Data (without schema interface) returns string map array.
// keys of map will be set spreadsheet header (currently first row),
// values of map will be set each spreadsheet record.
// note: each map value will be always set as string, even if actual value is number.
func (r *Records) Data() ([]map[string]string, error) {
	var records []map[string]string
	reader := csv.NewReader(r.body)
	headers, err := reader.Read()
	if err == io.EOF {
		return []map[string]string{}, nil
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
		record := map[string]string{}
		for i, v := range row {
			record[headers[i]] = v
		}
		records = append(records, record)
	}
	return records, nil
}

// func (r *Records) DataTo(p interface{}) error {}
