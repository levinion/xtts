package xtts

type Client struct {
	authorization string
}

func C(auth string) *Client {
	return &Client{authorization: auth}
}
