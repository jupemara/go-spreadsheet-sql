package sheet

import (
	"encoding/json"
	"io"
)

type Response struct {
	buf io.Reader
}

func NewResponse(buf io.Reader) *Response {
	return &Response{buf}
}

// Data converts from raw bytes http response body to map array.
// keys of map will be set spreadsheet header,
// and values of map will be set each spreadsheet record.
func (r *Response) Data() ([]map[string]interface{}, error) {
	var s schema
	if err := json.NewDecoder(r.buf).Decode(&s); err != nil {
		return nil, err
	}
	vs := []map[string]interface{}{}
	ks := []string{}
	for _, v := range s.Table.Cols {
		ks = append(ks, v.Label)
	}
	for _, row := range s.Table.Rows {
		v := map[string]interface{}{}
		for i, column := range row.C {
			v[ks[i]] = column.V
		}
		vs = append(vs, v)
	}
	return vs, nil
}

func (r *Response) DataTo(p interface{}) error {
	vs, err := r.Data()
	if err != nil {
		return err
	}
	bs, err := json.Marshal(vs)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(bs, p); err != nil {
		return err
	}
	return nil
}
