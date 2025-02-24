package user

import "log"

type Service interface {
	Create(firstName, lastName, email, phone string) (*User, error)
	Get(id string) (*User, error)
	GetAll() ([]User, error)
	Delete(id string) error
}

type service struct {
	log  *log.Logger
	repo Repository
}

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s service) Create(firstName, lastName, email, phone string) (*User, error) {
	s.log.Println("Create user service")

	user := User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}

	if err := s.repo.Create(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (s service) GetAll() ([]User, error) {
	s.log.Println("Get all user service")

	users, err := s.repo.GetAll()

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s service) Get(id string) (*User, error) {
	s.log.Println("Get user service")

	user, err := s.repo.Get(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s service) Delete(id string) error {

	_, err := s.repo.Get(id)

	if err != nil {
		return err
	}

	return s.repo.DeleteById(id)
}
