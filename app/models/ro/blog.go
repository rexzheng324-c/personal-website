package ro

type CreateBlogBody struct {
	// title
	// Required: true
	// Example: hello-world
	Title string `json:"title" validate:"required" label:"标题"`
	// content
	// Required: true
	// Example: 123456
	Content string `json:"content" validate:"required" label:"内容"`
}

// swagger:parameters CreateBlogBody
type CreateBlog struct {
	// In: body
	Body CreateBlogBody
}

type ListBlogsBody struct {
	// title
	// Required: false
	// Example: hello-world
	Title string `json:"title" label:"标题"`
	// offset
	// Required: false
	// Example: 10
	Offset int `json:"offset" label:"偏移量"`
	// page limit
	// Required: true
	// Example: 10
	Limit int `json:"limit" validate:"required" label:"每页个数"`
}

// swagger:parameters ListBlogsBody
type ListBlogs struct {
	// In: body
	Body ListBlogsBody
}

type DeleteBlogBody struct {
	// id
	// Required: true
	// Example: 4880bf21-0f61-4ad2-9073-abe15f51c41e
	Id string `json:"id" validate:"required" label:"id"`
}

// swagger:parameters DeleteBlogBody
type DeleteBlog struct {
	// In: body
	Body DeleteBlogBody
}
