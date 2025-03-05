package enrollment

import (
	"errors"
	"github.com/og11423074s/go_course_web/internal/course"
	"github.com/og11423074s/go_course_web/internal/domain"
	"github.com/og11423074s/go_course_web/internal/user"
	"log"
)

type (
	Service interface {
		Create(userID, courseID string) (*domain.Enrollment, error)
	}

	service struct {
		log       *log.Logger
		repo      Repository
		userSrv   user.Service
		courseSrv course.Service
	}
)

func NewService(log *log.Logger, repo Repository, userSrv user.Service, courseSrv course.Service) Service {
	return &service{
		log:       log,
		repo:      repo,
		userSrv:   userSrv,
		courseSrv: courseSrv,
	}
}

func (s *service) Create(userID, courseID string) (*domain.Enrollment, error) {

	enroll := &domain.Enrollment{
		UserID:   userID,
		CourseID: courseID,
		Status:   "P",
	}

	// Check if user exists
	if _, err := s.userSrv.Get(userID); err != nil {
		return nil, errors.New("user id does not exist")
	}

	// Check if course exists
	if _, err := s.courseSrv.Get(courseID); err != nil {
		return nil, errors.New("course id does not exist")
	}

	if err := s.repo.Create(enroll); err != nil {
		s.log.Println(err)
		return nil, err
	}

	return enroll, nil
}
