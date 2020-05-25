package sheet

import (
	"net/http"
	"net/url"
)

const (
	baseUrl = "https://spreadsheets.google.com/tq"
	format  = "out:csv"
)

func NewRequest(spreadsheetKey, worksheetName, q string) (*http.Request, error) {
	u := buildUrl(
		buildQueryStrings(spreadsheetKey, worksheetName, q),
	)
	return http.NewRequest(
		http.MethodGet,
		u,
		nil,
	)
}

func buildUrl(qs url.Values) string {
	return baseUrl + "?" + qs.Encode()
}

func buildQueryStrings(spreadsheetKey, worksheetName, q string) url.Values {
	v := url.Values{}
	v.Add("headers", "1") // so far header line of each sheet is always 1
	v.Add("key", spreadsheetKey)
	v.Add("sheet", worksheetName)
	v.Add("tq", q)
	v.Add("tqx", format)
	return v
}
