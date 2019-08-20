package comment

import (
	"GrabGoTrainingWeek5Assignment/httpclient"
	"encoding/json"
	"io/ioutil"
)

const (
	getCommentsEndpoint = "https://my-json-server.typicode.com/typicode/demo/comments"
)

type Handler struct {
	httpclient httpclient.HTTPClient
}

func New(httpclient httpclient.HTTPClient) *Handler {
	return &Handler{httpclient: httpclient}
}

func (h *Handler) GetComments() ([]Comment, error) {
	resp, err := h.httpclient.Get(getCommentsEndpoint)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	defer func() {
		_ = resp.Body.Close()
	}()

	var comments []Comment
	if err = json.Unmarshal(body, &comments); err != nil {
		return nil, err
	}

	return comments, nil
}
