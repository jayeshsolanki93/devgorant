package devrant

type CommentModel struct {
	Id           int    `json:"id"`
	RantId       int    `json:"rant_id"`
	Body         string `json:"body"`
	Upvotes      int    `json:"num_upvotes"`
	Downvotes    int    `json:"num_downvotes"`
	Score        int    `json:"score"`
	CreatedTime  int    `json:"created_time"`
	UserId       int    `json:"user_id"`
	UserUsername string `json:"user_username"`
	UserScore    int    `json:"user_score"`
}
