package Types


type TranslateReq struct {
	Query	string	`json:"query"`
	From	string 	`json:"from"`
	To		string 	`json:"to"`
}

type TranslateResult struct {
	Src string `json:"src"`
	Dst string `json:"dst"`
}

type TranslateRsp struct {
	From	string 	`json:"from"`
	To		string 	`json:"to"`
	Result  []TranslateResult `json:"trans_result"`
}