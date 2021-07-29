package Types


type TranslateReq struct {
	Query	string	`json:"query"`
	From	string 	`json:"from"`
	To		string 	`json:"to"`
}
