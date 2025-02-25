package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type Repository interface {
	Create(user *User) error
	GetAll() ([]User, error)
	Get(id string) (*User, error)
	DeleteById(id string) error
	Update(id string, firstName *string, lastName *string, email *string, phone *string) error
}

type repo struct {
	log *log.Logger
	db  *gorm.DB
}

func NewRepo(log *log.Logger, db *gorm.DB) Repository {
	return &repo{
		log: log,
		db:  db,
	}
}

func (r *repo) Create(user *User) error {

	user.ID = uuid.New().String()

	if err := r.db.Create(user); err.Error != nil {
		r.log.Println(err.Error)
		return err.Error
	}
	r.log.Println("User created with id: ", user.ID)
	return nil
}

func (r *repo) GetAll() ([]User, error) {
	var users []User

	err := r.db.Model(&users).Order("created_at desc").Find(&users).Error
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *repo) Get(id string) (*User, error) {
	user := User{ID: id}

	err := r.db.First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *repo) DeleteById(id string) error {
	user := User{ID: id}

	err := r.db.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Update(id string, firstName *string, lastName *string, email *string, phone *string) error {

	values := make(map[string]interface{})

	if firstName != nil {
		values["first_name"] = *firstName
	}

	if lastName != nil {
		values["last_name"] = *lastName
	}

	if email != nil {
		values["email"] = *email
	}

	if phone != nil {
		values["phone"] = *phone
	}

	if err := r.db.Model(&User{}).Where("id = ?", id).Updates(values).Error; err != nil {
		return err
	}

	return nil

}
