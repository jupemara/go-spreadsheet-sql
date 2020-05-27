package sheet_test

import (
	"context"
	"testing"

	sheet "github.com/jupemara/go-spreadsheet-sql"
)

func TestClient_Query(t *testing.T) {
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
	client, err := sheet.NewClient(
		context.Background(),
		"14aayP76anHyRJyeVcTBJMTvqwyPeWZFFBpGffhko9HU",
		"test",
	)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	for _, c := range cases {
		_, err := client.Query(context.Background(), c.Query)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		// results, err := records.Data()
		// if err != nil {
		// 	t.Errorf("unexpected error: %s", err)
		// }
		// if len(results) != c.Length {
		// 	t.Errorf("expected length: %d, but actual: %d", c.Length, len(results))
		// }
		// for i, record := range results {
		// 	for k, v := range record {
		// 		if v != c.Contents[i][k] {
		// 			t.Errorf("expected content: %s, but actual: %s", c.Contents[i][k], v)
		// 		}
		// 	}
		// }
	}
}
