# go-spreadsheet-sql

A client library for [Google Sheets](https://www.google.com/sheets/about/) with like SQL syntax.
As you are familiar, users can fetch and filter actual data on each Sheet by ["=QUERY()"](https://developers.google.com/chart/interactive/docs/querylanguage) .
This library allows you fetch and filter each Sheet record with Golang.
This library supports public and private Sheet,
and also supports any authentication method; oauth2, service account, credential json file, etc.

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

## client parameters

### required

You have to pass two required arguments; "spreadsheet key" and "worksheet name".

#### spreadsheet key

You can extract "spreadsheet key" from sheet url like following.
Each spreadsheet url is made up of https://docs.google.com/spreadsheets/d/${SPREADSHEET_KEY}/edit#gid=0 .

#### worksheet name

"worksheet name" is shown at the bottom of browser like a tab.

### optional

You can pass credential information by your own methods. For instance (now listing only popuplar options),

- credential file path as string
- service account file path as string
- credential json as `[]byte` (credential json structure must be same with https://github.com/golang/oauth2/blob/master/google/google.go#L99 )
- oauth2 token

For details please see https://google.golang.org/api/option .

## Response object

`sheets.Client.Query` method returns `sheets.Response` object.
`Response` object has two methods to convert `map` or `json`.
Designing those method signature is inspired by firestore client library.

firestore client library can return `DocumentSnapshot` as returnd value of `Get` method.
`DocumentSnapshot` has two methods; `Data` and `DataTo` .
Our `Response` object has also methods named samely.

`Response.Data` method simply returns `[]map[string]interface{}`.

When target sheet has following data structure,

| name    | email               | url                         |
|---------|---------------------|-----------------------------|
| user001 | user001@example.com | https://user001.example.com |

`Data` method returns data as below.

```golang
original, _ := res.Data()
// variable "original" is completely same with following map
sameWithOriginal := []map[string]interface{}{{
    "name": "user001",
    "email": "user001@example.com",
    "url": "https://user001.example.com",
}}
```

`Response.DataTo` method receive one argument as schema information.
You can use Golang struct similar with json annotated struct.

```golang
type Schema struct {
    Name string `json:"name"`
    Email string `json:"email"`
    Url string `json:"url"`
} 

res, _ := client.Query("SELECT *")
var result Schema
err := res.DataTo(&s)
sameWithResult := Schema{
    Name: "user001",
    Email: "user001@example.com",
    Url: "https://user001.example.com",
}
```

If you would be happy to be familiar with firestore exported method design, you can see https://godoc.org/cloud.google.com/go/firestore#DocumentSnapshot.DataTo .