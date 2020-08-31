package main

//DataMap to keep track of data where keys are org,paramName,paramVal & value is array of segmentIds
type DataMap struct {
	dataMap map[string]map[string]map[string][]string
}

//RequestParams contains request params
type RequestParams struct {
	OrgName   string `json:"orgName"`
	ParamName string `json:"paramName"`
	ParamVal  string `json:"paramVal"`
}

//ResponseParams contains response params
type ResponseParams struct {
	IDs []string `json:"ids"`
}
