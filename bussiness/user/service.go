package user

type CreateUserSpec struct {
	Firstname string
	Lastname  string
	Username  string
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) RegisterUser(payload CreateUserSpec) (User, error) {
	return s.repository.RegisterUser(payload)
}
