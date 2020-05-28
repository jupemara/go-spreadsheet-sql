package sheets_test

import (
	"strings"
	"testing"

	sheets "github.com/jupemara/go-spreadsheet-sql"
)

type user2 struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Count int    `json:"some_count"`
}

func TestResponse_DataTo(t *testing.T) {
	raw := `
{
  "version": "0.6",
  "reqId": "0",
  "status": "ok",
  "sig": "1068771412",
  "table": {
    "cols": [
      {
        "id": "A",
        "label": "name",
        "type": "string"
      },
      {
        "id": "B",
        "label": "email",
        "type": "string"
      },
      {
        "id": "C",
        "label": "some_count",
        "type": "number",
        "pattern": "General"
      }
    ],
    "rows": [
      {
        "c": [
          {"v": "user001"},
          {"v": "user001@example.com"},
          {
            "v": 1,
            "f": "1"
          }
        ]
      },
      {
        "c": [
          {"v": "user002"},
          {"v": "user002@example.com"},
          {"v": 128}
        ]
      },
      {
        "c": [
          {"v": "user003"}
        ]
      }
    ],
    "parsedNumHeaders": 1
  }
}`
	r := sheets.NewResponse(strings.NewReader(raw))
	vs := []user2{}
	err := r.DataTo(&vs)
	if err != nil {
		t.Errorf("unexpected error: %s", err)
		return
	}
	cases := [][]interface{}{{
		vs[0].Name, "user001",
	}, {
		vs[0].Email, "user001@example.com",
	}, {
		vs[0].Count, 1,
	}, {
		vs[1].Name, "user002",
	}, {
		vs[1].Email, "user002@example.com",
	}, {
		vs[1].Count, 128,
	}, {
		vs[2].Name, "user003",
	}, {
		vs[2].Email, "",
	}, {
		vs[2].Count, 0,
	}}
	for _, c := range cases {
		e := c[0]
		a := c[1]
		if e != a {
			t.Errorf("expected: %v, actual: %v", e, a)
		}
	}
}
