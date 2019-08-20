package post_comment

import "GrabGoTrainingWeek5Assignment/services/comment"


type PostWithComments struct {
	ID       int64     `json:"id"`
	Title    string    `json:"string"`
	Comments []comment.Comment `json:"comments,omitempty"`
}

type PostWithCommentsResponse struct {
	Posts []PostWithComments `json:"posts"`
}


type HandlerInterface interface {
	GetPostWithComments() ([]PostWithComments, error)
}
