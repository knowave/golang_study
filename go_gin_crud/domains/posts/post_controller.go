package posts

type PostController struct {
	service *PostService
}

func NewPostController(service *PostService) *PostController {
	return &PostController{service: service}
}