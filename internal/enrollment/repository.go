package enrollment

import (
	"github.com/og11423074s/go_course_web/internal/domain"
	"gorm.io/gorm"
	"log"
)

type (
	Repository interface {
		Create(enroll *domain.Enrollment) error
	}

	repo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewRepo(log *log.Logger, db *gorm.DB) Repository {
	return &repo{
		db:  db,
		log: log,
	}
}

func (r *repo) Create(enroll *domain.Enrollment) error {
	if err := r.db.Create(enroll).Error; err != nil {
		r.log.Printf("error: %v", err)
		return err
	}
	r.log.Println("enrollment created ith id: ", enroll.ID)
	return nil
}
