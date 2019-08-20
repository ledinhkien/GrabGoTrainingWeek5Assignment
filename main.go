package main

import (
	"GrabGoTrainingWeek5Assignment/services/comment"
	"GrabGoTrainingWeek5Assignment/services/post"
	"GrabGoTrainingWeek5Assignment/services/post_comment"
	"log"
	"net/http"
)


func main() {
	httpClient := &http.Client{}
	postHandler := post.New(httpClient)
	commentHandler := comment.New(httpClient)
	postWithCommentHandler := post_comment.New(*postHandler, *commentHandler)

	http.HandleFunc("/postWithComments", postWithCommentHandler.HandlePostWithCommentsRequest)
	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
