package sheets_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"google.golang.org/api/option"

	sheets "github.com/jupemara/go-spreadsheet-sql"
)

func TestClient_Query_WithPrivateSheet(t *testing.T) {
	cases := map[string]struct {
		Query    string
		Length   int
		Contents []map[string]interface{}
	}{
		"should return single row": {
			`SELECT * WHERE A = "spreadsheet-sql-private001"`,
			1,
			[]map[string]interface{}{{
				"name":  "spreadsheet-sql-private001",
				"url":   "https://spreadsheet-sql-private001.example.com",
				"email": "spreadsheet-sql-private001@example.com",
			}},
		},
		"should return multiple rows": {
			`SELECT * WHERE A LIKE "spreadsheet-sql-private%"`,
			2,
			[]map[string]interface{}{{
				"name":  "spreadsheet-sql-private001",
				"url":   "https://spreadsheet-sql-private001.example.com",
				"email": "spreadsheet-sql-private001@example.com",
			}, {
				"name":  "spreadsheet-sql-private002",
				"url":   "https://spreadsheet-sql-private002.example.com",
				"email": "spreadsheet-sql-private002@example.com",
			}},
		},
		"should return 0 length slice": {
			`SELECT * WHERE A = "non-existent-user"`,
			0,
			[]map[string]interface{}{},
		},
	}
	key := os.Getenv("GO_SPREADSHEET_SQL_PRIVATE_SHEET_KEY")
	// To clarify credential json structure, see https://github.com/golang/oauth2/blob/master/google/google.go#L99
	credential := fmt.Sprintf(`{
  "type": "authorized_user",
  "client_id": "%s",
  "client_secret": "%s",
  "refresh_token": "%s"
}`,
		os.Getenv("GO_SPREADSHEET_SQL_CLIENT_ID"),
		os.Getenv("GO_SPREADSHEET_SQL_CLIENT_SECRET"),
		os.Getenv("GO_SPREADSHEET_SQL_REFRESH_TOKEN"),
	)
	client, err := sheets.NewClient(
		context.Background(),
		key,
		"test",
		option.WithCredentialsJSON([]byte(credential)),
	)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	for _, c := range cases {
		res, err := client.Query(context.Background(), c.Query)
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		results, err := res.Data()
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}
		if len(results) != c.Length {
			t.Errorf("expected length: %d, but actual: %d", c.Length, len(results))
		}
		for i, record := range results {
			for k, v := range record {
				if v != c.Contents[i][k] {
					t.Errorf("expected content: %s, but actual: %s", c.Contents[i][k], v)
				}
			}
		}
	}
}
