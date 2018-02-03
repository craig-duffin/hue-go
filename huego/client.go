package huego

import (
	"io/ioutil"
	"net/http"
)

type Client struct {
	username string
	bridge string
}

func NewClient(user, bridgeAddress string) *Client {
	return &Client{
		username: user,
		bridge: bridgeAddress,
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
