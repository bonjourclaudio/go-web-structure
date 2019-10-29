package user

type UserService interface {
	Find(id string) (*User, error)
	FindAll() ([]*User, error)
	Create(user User) error
	Delete(id string) error
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo,
	}
}

func (s *userService) Find(id string) (*User, error) {
	return s.repo.Find(id)
}

func (s *userService) FindAll() ([]*User, error) {
	return s.repo.FindAll()
}

func (s *userService) Create(user User) error {
	return s.repo.Create(user)
}

func (s *userService) Delete(id string) error {
	return s.repo.Delete(id)
}
