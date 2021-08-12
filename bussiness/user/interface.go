package user

type Repository interface {
	RegisterUser(payload CreateUserSpec) (User, error)
}

type Service interface {
	RegisterUser(payload CreateUserSpec) (User, error)
}
