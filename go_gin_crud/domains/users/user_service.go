package users

type UserService struct {
	repository *UserRepository
}

func NewUserService(repository *UserRepository) *UserService {
	return &UserService{repository: repository}
}