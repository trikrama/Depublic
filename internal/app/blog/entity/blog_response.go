package entity

type BlogResponse struct {
	ID    int64  `json:"id"`
	Image string `json:"image"`
	Date  string `json:"date"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func NewBlogRespose(blog Blog) BlogResponse {
	return BlogResponse{
		ID:    blog.ID,
		Image: blog.Image,
		Date:  blog.Date,
		Title: blog.Title,
		Body:  blog.Body,
	}
}
