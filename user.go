package devgorant

type UserModel struct {
	Username    string `json:"username"`
	Score       int    `json:"score"`
	About       string `json:"about"`
	Location    string `json:"location"`
	CreatedTime int    `json:"created_time"`
	Skills      string `json:"skills"`
	Github      string `json:"github"`
	Content     struct {
		Content ContentModel `json:"content"`
	} `json:"content"`
}
