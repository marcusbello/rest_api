package post

type Getter interface {
	GetAll() []Post
}

type Setter interface {
	Add(post Post)
}

type Post struct {
	PostId    int    `json:"post_id"`
	PostTitle string `json:"post_title"`
	PostBody  string `json:"post_body"`
}

type PostList struct {
	Posts []Post
}

func New() *PostList {
	return &PostList{
		Posts: []Post{},
	}
}

func (r *PostList) Add(post Post) {
	r.Posts = append(r.Posts, post)
}

func (r *PostList) GetAll() []Post {
	return r.Posts
}
