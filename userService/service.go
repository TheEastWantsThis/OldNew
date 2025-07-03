package userservice

type MainUserService interface {
	CreateUser(user UsersOrm) (UsersOrm, error)
	GetAllUsers() ([]UsersOrm, error)
	GetUserByID(id int) (UsersOrm, error)
	UpdateUser(id int, user UsersOrm) (UsersOrm, error)
	DeleteUser(id int) error
}

type userService struct {
	repo MainUsersRepository
}

func NewUserService(repo MainUsersRepository) *userService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) CreateUser(user UsersOrm) (UsersOrm, error) {

	newUser := UsersOrm{
		Email:    user.Email,
		Password: user.Password,
	}

	if err := s.repo.CreateUser(&newUser); err != nil {
		return UsersOrm{}, err
	}

	return newUser, nil
}

func (s *userService) GetAllUsers() ([]UsersOrm, error) {
	return s.repo.GetAllUsers()
}

func (s *userService) GetUserByID(id int) (UsersOrm, error) {
	return s.repo.GetUserByID(id)
}

func (s *userService) UpdateUser(id int, user UsersOrm) (UsersOrm, error) {
	updateUser, err := s.repo.GetUserByID(id)
	if err != nil {
		return UsersOrm{}, err
	}

	updateUser.Email = user.Email
	updateUser.Password = user.Password

	if err := s.repo.UpdateUser(updateUser); err != nil {
		return UsersOrm{}, err
	}
	return updateUser, err
}

func (s *userService) DeleteUser(id int) error {
	return s.repo.DeleteUser(id)
}
