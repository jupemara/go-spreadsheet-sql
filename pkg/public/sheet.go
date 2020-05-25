package public

import (
	"encoding/csv"
	"io"
	"net/http"

	"github.com/jupemara/go-spreadsheet-sql/pkg/util"

	"github.com/jupemara/go-spreadsheet-sql/pkg/sheet"
)

// TODO: fix name
type Sheet struct {
	spreadSheetKey string
	workSheetName  string
	client         *http.Client
}

func NewSheet(spreadSheetKey, workSheetName string) (*Sheet, error) {
	client, err := util.NewClient()
	if err != nil {
		return nil, err
	}
	return &Sheet{
		spreadSheetKey: spreadSheetKey,
		workSheetName:  workSheetName,
		client:         client,
	}, nil
}

// TODO: fix returnd type as Record
func (s *Sheet) Query(q string) ([]map[string]string, error) {
	req, err := sheet.NewRequest(
		s.spreadSheetKey,
		s.workSheetName,
		q,
	)
	if err != nil {
		return nil, err
	}
	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	// TODO: postprocessing
	defer res.Body.Close()

	var records []map[string]string // so far we provide just only key and values without converting
	r := csv.NewReader(res.Body)
	headers, err := r.Read()
	if err == io.EOF {
		return []map[string]string{}, nil
	}
	if err != nil {
		return nil, err
	}
	for {
		row, err := r.Read()
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
