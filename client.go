package huego

import (
	"net/http"
	"io/ioutil"
)

const bridgeURL string = "http://192.168.0.4/api/"

type Client struct {
	Username string
}

func NewClient(username string) *Client {
	return &Client{
		Username: username,
	}
}

func (c *Client) doRequest(req *http.Request) ([]byte, error){
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return body, nil
}