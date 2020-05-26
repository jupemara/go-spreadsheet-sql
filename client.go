package sheet

import (
	"bytes"
	"io"
	"net/http"
	"net/url"

	"google.golang.org/api/option"
	ghttp "google.golang.org/api/transport/http"

	"context"
)

type Client struct {
	key        string
	sheet      string
	httpClient *http.Client
}

func NewClient(
	ctx context.Context,
	key, sheet string,
	opts ...option.ClientOption,
) (*Client, error) {
	httpClient, _, err := ghttp.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	return &Client{
		key:        key,
		sheet:      sheet,
		httpClient: httpClient,
	}, nil
}

// Query (without schema) returns map array.
// keys of map will be set spreadsheet header,
// ans values of map will be set each spreadsheet record.
// note: each record value will be always set as string, even if actual value is number.
func (c *Client) Query(ctx context.Context, q string) (*Records, error) {
	v := url.Values{}
	v.Add("headers", "1") // so far header line of each sheet is always 1
	v.Add("key", c.key)
	v.Add("sheet", c.sheet)
	v.Add("tq", q)
	v.Add("tqx", "out:csv")
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://spreadsheets.google.com/tq"+"?"+v.Encode(),
		nil,
	)
	if err != nil {
		return nil, err
	}
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	w := bytes.NewBuffer(nil)
	r := io.TeeReader(res.Body, w)
	return &Records{
		body: r,
	}, nil
}
