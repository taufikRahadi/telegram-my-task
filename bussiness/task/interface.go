package task

type Repository interface {
	CreateTask(payload Task) (Task, error)

	GetAllMyTask(username string) ([]Task, error)
}

type Service interface {
	CreateTask(payload *CreateTaskSpec) (Task, error)

	GetAllMyTask(username string) ([]Task, error)
}
