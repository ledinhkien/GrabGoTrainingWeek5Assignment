package post_comment

import (
	"GrabGoTrainingWeek5Assignment/renderer"
	"GrabGoTrainingWeek5Assignment/services/comment"
	"GrabGoTrainingWeek5Assignment/services/post"
	"log"
	"net/http"
)

const (
	getPostsEndpoint = "https://my-json-server.typicode.com/typicode/demo/posts"
)

type Handler struct {
	postHandler    post.Handler
	commentHandler comment.Handler
}

func New(postHandler post.Handler, commentHandler comment.Handler) *Handler {
	return &Handler{postHandler: postHandler, commentHandler: commentHandler}
}


func (h *Handler) GetPostWithComments() ([]PostWithComments, error) {
	posts, err := h.postHandler.GetPosts()
	if err != nil {
		return nil, err
	}
	comments, err := h.commentHandler.GetComments()
	if err != nil {
		return nil, err
	}
	return combinePostWithComments(posts, comments), nil
}

func (h *Handler) HandlePostWithCommentsRequest(writer http.ResponseWriter, request *http.Request) {
	postWithComments, err := h.GetPostWithComments()
	resp := PostWithCommentsResponse{Posts: postWithComments}
	renderFunc, contentType, err := renderer.GetRenderFunctionAndContentType(request)
	if err != nil {
		log.Println("unable to parse request: ", err)
		writer.WriteHeader(415)
		return
	}
	buf, err := renderFunc(resp)
	if err != nil {
		log.Println("unable to parse response: ", err)
		writer.WriteHeader(500)
		return
	}

	writer.Header().Set("Content-Type", contentType)
	_, err = writer.Write(buf)
}
//func (h *post.Handler)

func combinePostWithComments(posts []post.Post, comments []comment.Comment) []PostWithComments {
	commentsByPostID := map[int64][]comment.Comment{}
	for _, comment := range comments {
		commentsByPostID[comment.PostID] = append(commentsByPostID[comment.PostID], comment)
	}

	result := make([]PostWithComments, 0, len(posts))
	for _, post := range posts {
		result = append(result, PostWithComments{
			ID:       post.ID,
			Title:    post.Title,
			Comments: commentsByPostID[post.ID],
		})
	}

	return result
}
