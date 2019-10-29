package user

type UserRepository interface {
	Find(id string) (*User, error)
	FindAll() ([]*User, error)
	Create(user User) error
	Delete(id string) error
}