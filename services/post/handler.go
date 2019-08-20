package post

import (
	"GrabGoTrainingWeek5Assignment/httpclient"
	"encoding/json"
	"io/ioutil"
)

const (
	getPostsEndpoint = "https://my-json-server.typicode.com/typicode/demo/posts"
)

type Handler struct {
	httpclient httpclient.HTTPClient
}

func New(httpclient httpclient.HTTPClient) *Handler {
	return &Handler{httpclient: httpclient}
}

func (h *Handler) GetPosts() ([]Post, error) {
	resp, err := h.httpclient.Get(getPostsEndpoint)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var posts []Post
	if err = json.Unmarshal(body, &posts); err != nil {
		return nil, err
	}

	return posts, nil
}
