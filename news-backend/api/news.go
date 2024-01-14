package news

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

/*
	http: points to the HTTP client that should be used to make requests
	apiKey: News API key
	PageSize: number of pages to return (maximum of 100)
*/
type Client struct {
	http     *http.Client
	key      string
	PageSize int
}

/*
	Go struct equivalent of the JSON response from the
 	News API /everything endpoint for an article
*/
type Article struct {
	Source struct {
		ID   interface{} `json:"id"`
		Name string      `json:"name"`
	} `json:"source"`
	Author      string    `json:"author"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	URLToImage  string    `json:"urlToImage"`
	PublishedAt time.Time `json:"publishedAt"`
	Content     string    `json:"content"`
}

/*
	Go struct that has all the articles obtained
	from the News API /everything endpoint
*/
type Results struct {
	Status       string    `json:"status"`
	TotalResults int       `json:"totalResults"`
	Articles     []Article `json:"articles"`
}

func NewClient(httpClient *http.Client, key string, pageSize int) *Client {
	if pageSize > 100 {
		pageSize = 100
	}

	return &Client{httpClient, key, pageSize}
}

/*
	Call the News API /everything endpoint
	Method, function with a reciever argument that allows the function to use the properties of the reciever
	argument, in this case the struct Client
*/
func (c *Client) FetchEverything(query string, page string) (*Results, error) {
	endpoint := fmt.Sprintf(
		"https://newsapi.org/v2/everything?q=%s&pageSize=%d&page=%s&apiKey=%s&sortBy=publishedAt&language=en",
		url.QueryEscape(query),
		c.PageSize,
		page,
		c.key)

	response, err := c.http.Get(endpoint)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Response body converted to byte slice with ReadAll
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(body))
	}

	// Response body decoding using Unmarshal into the Result struct
	results := &Results{}
	return results, json.Unmarshal(body, results)
}
