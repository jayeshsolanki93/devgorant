package devgorant

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	API         = "https://www.devrant.io/api"
	APP_VERSION = 3
)

type Client struct {
}

// Fetches rants
func (c *Client) Rants(sort string, limit int, skip int) ([]RantModel, error) {
	if sort == "" {
		sort = "algo"
	} else if (sort != "algo") && (sort != "recent") && (sort != "top") {
		return nil, errors.New("Invalid string in sort method")
	}
	if limit <= 0 {
		limit = 50
	}

	url := fmt.Sprintf("%s/devrant/rants?sort=%s&limit=%d&skip=%d&app=%d",
		API, sort, limit, skip, APP_VERSION)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var data RantsResponse
	json.NewDecoder(res.Body).Decode(&data)
	return data.Rants, nil
}

// Fetches a rant and its comments given a valid rant id
func (c *Client) Rant(rantId int) (RantModel, []CommentModel, error) {
	url := fmt.Sprintf("%s/devrant/rants/%d?app=%d", API, rantId, APP_VERSION)
	res, err := http.Get(url)
	if err != nil {
		return RantModel{}, nil, err
	}
	var data RantResponse
	json.NewDecoder(res.Body).Decode(&data)
	return data.Rant, data.Comments, nil
}

// Fetches ranter's profile data
func (c *Client) Profile(username string) (UserModel, ContentModel, error) {
	userId, err := getUserId(username)
	if err != nil {
		return UserModel{}, ContentModel{}, err
	}
	url := fmt.Sprintf("%s/users/%d?app=%d", API, userId, APP_VERSION)
	res, err := http.Get(url)
	if err != nil {
		return UserModel{}, ContentModel{}, err
	}
	var data UserResponse
	json.NewDecoder(res.Body).Decode(&data)
	return data.Profile, data.Profile.Content.Content, nil
}

// Search for rants matching the search term
func (c *Client) Search(term string) ([]RantModel, error) {
	url := fmt.Sprintf("%s/devrant/search?term=%s&app=%d", API, term, APP_VERSION)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var data SearchResponse
	json.NewDecoder(res.Body).Decode(&data)
	return data.Rants, nil
}

// Returns a random rant
func (c *Client) Surprise() (RantModel, error) {
	url := fmt.Sprintf("%s/devrant/rants/surprise?app=%d", API, APP_VERSION)
	res, err := http.Get(url)
	if err != nil {
		return RantModel{}, err
	}
	var data RantResponse
	json.NewDecoder(res.Body).Decode(&data)
	return data.Rant, nil
}

// Returns the rants tagged for 'weekly'
func (c *Client) WeeklyRants() ([]RantModel, error) {
	url := fmt.Sprintf("%s/devrant/weekly-rants?app=%d", API, APP_VERSION)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	var data RantsResponse
	json.NewDecoder(res.Body).Decode(&data)
	return data.Rants, nil
}

// Fetches the userId given a valid username
func getUserId(username string) (int, error) {
	url := fmt.Sprintf("%s/get-user-id?username=%s&app=%d",
		API, username, APP_VERSION)
	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	var data GetUserIdResponse
	json.NewDecoder(res.Body).Decode(&data)
	return data.UserId, nil
}

func New() *Client {
	return new(Client)
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
