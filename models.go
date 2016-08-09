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

type ImageModel struct {
	Url    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

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

type ContentModel struct {
	Rants     []RantModel    `json:"rants"`
	Upvoted   []RantModel    `json:"upvoted"`
	Comments  []CommentModel `json:"comments"`
	Favorites []RantModel    `json:"favorites"`
}

type NewsModel struct {
	Id       int    `json:"id"`
	Type     string `json:"type"`
	Headline string `json:"headline"`
	Body     string `json:"body"`
	Footer   string `json:"footer"`
	Height   int    `json:"height"`
	Action   string `json:"action"`
}

type RantsResponse struct {
	Success  bool        `json:"success"`
	Rants    []RantModel `json:"rants"`
	Settings string      `json:"settings"`
	Set      string      `json:"set"`
	Wrw      int         `json:"wrw"`
	News     NewsModel   `json:"news"`
}

type RantResponse struct {
	Success  bool           `json:"success"`
	Rant     RantModel      `json:"rant"`
	Comments []CommentModel `json:"comments"`
}

type UserResponse struct {
	Success bool      `json:"success"`
	Profile UserModel `json:"profile"`
}

type SearchResponse struct {
	Success bool        `json:"success"`
	Rants   []RantModel `json:"results"`
}

type GetUserIdResponse struct {
	Success bool `json:"success"`
	UserId  int  `json:"user_id"`
}
