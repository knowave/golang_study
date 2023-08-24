package posts

type PostService struct {
	repository *PostRepository
}

func NewPostService(repository *PostRepository) *PostService {
	return &PostService{repository: repository}
}