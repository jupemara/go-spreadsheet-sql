package sheet

type schema struct {
	Version string `json:"version"`
	ReqId   string `json:"reqId"`
	Status  string `json:"status"`
	Sig     string `json:"sig"`
	Table   table  `json:"table"`
}

type table struct {
	Cols             []col `json:"cols"`
	Rows             []row `json:"rows"`
	ParsedNumHeaders int   `json:"parsedNumHeaders"`
}

type col struct {
	Id    string `json:"id"`
	Label string `json:"label"`
	Type  string `json:"type"`
}

type row struct {
	C []v `json:"c"`
}

type v struct {
	V interface{} `json:"v"`
}
