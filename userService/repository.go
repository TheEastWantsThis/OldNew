package userservice

import "gorm.io/gorm"

type MainUsersRepository interface {
	CreateUser(user *UsersOrm) error
	GetAllUsers() ([]UsersOrm, error)
	GetUserByID(id int) (UsersOrm, error)
	UpdateUser(user UsersOrm) error
	DeleteUser(id int) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) MainUsersRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *UsersOrm) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetAllUsers() ([]UsersOrm, error) {
	var user []UsersOrm
	err := r.db.Order("id asc").Find(&user).Error
	return user, err
}

func (r *userRepository) GetUserByID(id int) (UsersOrm, error) {
	var user UsersOrm
	err := r.db.First(&user, "ID = ?", id).Error
	return user, err
}

func (r *userRepository) UpdateUser(user UsersOrm) error {
	return r.db.Save(&user).Error
}

func (r *userRepository) DeleteUser(id int) error {
	return r.db.Delete(&UsersOrm{}, "ID = ?", id).Error
}
