package public_test

import (
	"testing"

	"github.com/jupemara/go-spreadsheet-sql/pkg/public"
)

func TestSheet(t *testing.T) {
	cases := map[string]struct {
		Query    string
		Length   int
		Contents []map[string]string
	}{
		"should return single row": {
			`SELECT * WHERE A = "spreadsheet-sql-public001"`,
			1,
			[]map[string]string{{
				"name":  "spreadsheet-sql-public001",
				"url":   "https://spreadsheet-sql-public001.example.com",
				"email": "spreadsheet-sql-public001@example.com",
			}},
		},
		"should return multiple rows": {
			`SELECT * WHERE A LIKE "spreadsheet-sql-public00%"`,
			2,
			[]map[string]string{{
				"name":  "spreadsheet-sql-public001",
				"url":   "https://spreadsheet-sql-public001.example.com",
				"email": "spreadsheet-sql-public001@example.com",
			}, {
				"name":  "spreadsheet-sql-public002",
				"url":   "https://spreadsheet-sql-public002.example.com",
				"email": "spreadsheet-sql-public002@example.com",
			}},
		},
		"should return 0 length slice": {
			`SELECT * WHERE A = "non-existent-user"`,
			0,
			[]map[string]string{},
		},
	}
	s, err := public.NewSheet(
		"14aayP76anHyRJyeVcTBJMTvqwyPeWZFFBpGffhko9HU",
		"test",
	)
	if err != nil {
		t.Errorf("unexpcted error: %s", err)
	}
	for _, c := range cases {
		records, err := s.Query(c.Query)
		if err != nil {
			t.Errorf("unexpcted error: %s", err)
			break
		}
		if len(records) != c.Length {
			t.Errorf("expected length: %d, but actual: %d", c.Length, len(records))
		}
		for i, record := range records {
			for k, v := range record {
				if v != c.Contents[i][k] {
					t.Errorf("expected content: %s, but actual: %s", c.Contents[i][k], v)
				}
			}
		}
	}
}
