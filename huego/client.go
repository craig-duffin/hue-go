package huego

import (
	"io/ioutil"
	"net/http"
)

const bridgeURL string = "http://192.168.0.64/api/"

type Client struct {
	username string
}

func NewClient(user string) *Client {
	return &Client{
		username: user,
	}
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) setUsername(user string) {
	c.username = user
}
