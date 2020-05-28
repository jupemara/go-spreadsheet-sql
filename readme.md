# go-spreadsheet-sql

A client library for [Google Sheets](https://www.google.com/sheets/about/) with SQL syntax.
As you are familiar, users can fetch and filter actual data on each Sheet by ["=QUERY()"](https://developers.google.com/chart/interactive/docs/querylanguage) .
This library allows you fetch and filter each sheet record with Golang.
This library supports public and private Sheet,
and also supports any authentication method.

## install

```bash
$ go get -u github.comjupemara/go-spreadsheet-sql
```

## simple usage case

```golang
import (
    "log"

	"google.golang.org/api/option"
    "github.com/jupemara/go-spreadsheet-sql" // actual package name is sheets
)

client := sheets.NewClient(
    "SHEET_KEY",
    "WORKSHEET_NAME",
    option.WithoutAuthentication(),
)
records, _ := client.Query(
    context.TODO(),
    `SELECT * WHERE A = "user001"`, // specify column name like "A" (sheet cell index not actual header name) in "WHERE" clause
)
log.Printf(`results: %+v`, records)
```

## client option

### required

### optional

## response object

## tips

## API(exported method) design