package devgorant

type RantModel struct {
	Id            int        `json:"id"`
	Text          string     `json:"text"`
	Upvotes       int        `json:"num_upvotes"`
	Downvotes     int        `json:"num_downvotes"`
	Score         int        `json:"score"`
	CreatedTime   int        `json:"created_time"`
	AttachedImage ImageModel `json:"attached_image"`
	NumComments   int        `json:"num_comments"`
	Tags          []string   `json:"tags"`
	UserId        int        `json:"user_id"`
	UserUsername  string     `json:"user_username"`
	UserScore     int        `json:"user_score"`
}
