package devrant

type NewsModel struct {
	Id       int    `json:"id"`
	Type     string `json:"type"`
	Headline string `json:"headline"`
	Body     string `json:"body"`
	Footer   string `json:"footer"`
	Height   int    `json:"height"`
	Action   string `json:"action"`
}
