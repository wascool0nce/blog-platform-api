package request

type CreatePostRequest struct {
	Title    string `json:"title"    validate:"required,max=150"   example:"My post"`
	Content  string `json:"content"  validate:"required"           example:"My first post for software engineer."`
	Category string `json:"category" validate:"required,max=50"    example:"programming"`
}
