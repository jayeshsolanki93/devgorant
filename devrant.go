package devgorant

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

const (
	API           = "https://www.devrant.io/api"
	APP_VERSION   = 3
	RANTS_PATH    = "%s/devrant/rants?sort=%s&limit=%d&skip=%d&app=%d"
	RANT_PATH     = "%s/devrant/rants/%d?app=%d"
	USER_PATH     = "%s/users/%d?app=%d"
	USER_ID_PATH  = "%s/get-user-id?username=%s&app=%d"
	SEARCH_PATH   = "%s/devrant/search?term=%s&app=%d"
	SURPRISE_PATH = "%s/devrant/rants/surprise?app=%d"
	WEEKLY_PATH   = "%s/devrant/weekly-rants?app=%d"
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

	url := fmt.Sprintf(RANTS_PATH, API, sort, limit, skip, APP_VERSION)
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
	url := fmt.Sprintf(RANT_PATH, API, rantId, APP_VERSION)
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
	url := fmt.Sprintf(USER_PATH, API, userId, APP_VERSION)
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
	url := fmt.Sprintf(SEARCH_PATH, API, term, APP_VERSION)
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
	url := fmt.Sprintf(SURPRISE_PATH, API, APP_VERSION)
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
	url := fmt.Sprintf(WEEKLY_PATH, API, APP_VERSION)
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
	url := fmt.Sprintf(USER_ID_PATH, API, username, APP_VERSION)
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
