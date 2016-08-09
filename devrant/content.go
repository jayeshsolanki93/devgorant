package devrant

type ContentModel struct {
	Rants     []RantModel    `json:"rants"`
	Upvoted   []RantModel    `json:"upvoted"`
	Comments  []CommentModel `json:"comments"`
	Favorites []RantModel    `json:"favorites"`
}
