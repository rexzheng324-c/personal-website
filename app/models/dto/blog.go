package dto

type Blog struct {
	// id
	// Required: true
	// Example: 4880bf21-0f61-4ad2-9073-abe15f51c41e
	Id string `json:"id"`
	// title
	// Required: true
	// Example: hello-world
	Title string `json:"title"`
	// content
	// Required: true
	// Example: hello-world-1
	Content string `json:"content"`
}

type ListBlogsBody struct {
	// blogs
	// Required: true
	Blogs []Blog `json:"blogs"`
	// offset
	// Required: true
	// Example: 10
	Offset int `json:"offset"`
	// page count
	// Required: true
	// Example: 10
	Count int `json:"count"`
}

// swagger:response ListBlog
type ListBlog struct {
	// In: body
	Body ListBlogsBody
}
