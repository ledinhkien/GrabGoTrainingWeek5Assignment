package post


type Post struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type HandlerInterface interface {
	GetPosts() ([]Post, error)
}