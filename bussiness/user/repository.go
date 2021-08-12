package user

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{
		db,
	}
}

func (r *repository) RegisterUser(payload CreateUserSpec) (User, error) {
	newUser := User{Firstname: payload.Firstname, Lastname: payload.Lastname, Username: payload.Username}

	create := r.db.Create(&newUser)

	return newUser, create.Error
}
