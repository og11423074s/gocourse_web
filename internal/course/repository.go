package course

import (
	"fmt"
	"github.com/og11423074s/go_course_web/internal/domain"
	"gorm.io/gorm"
	"log"
	"strings"
	"time"
)

type (
	Repository interface {
		Create(course *domain.Course) error
		GetAll(filters Filters, offset, limit int) ([]domain.Course, error)
		Get(id string) (*domain.Course, error)
		DeleteById(id string) error
		Update(id string, name *string, startDate, endDate *time.Time) error
		Count(filters Filters) (int, error)
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

func (r *repo) Create(course *domain.Course) error {

	if err := r.db.Create(course).Error; err != nil {
		r.log.Printf("error: %v", err)
		return err
	}

	r.log.Println("course created ith id: ", course.ID)
	return nil
}

func (r *repo) GetAll(filters Filters, offset, limit int) ([]domain.Course, error) {
	var courses []domain.Course

	tx := r.db.Model(&courses)
	tx = applyFilters(tx, filters)
	tx = tx.Offset(offset).Limit(limit)

	result := tx.Order("created_at desc").Find(&courses)

	if result.Error != nil {
		return nil, result.Error
	}

	return courses, nil
}

func (r *repo) Get(id string) (*domain.Course, error) {
	course := domain.Course{ID: id}

	err := r.db.First(&course).Error
	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (r *repo) DeleteById(id string) error {
	user := domain.Course{ID: id}

	err := r.db.Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repo) Update(id string, name *string, startDate, endDate *time.Time) error {

	values := make(map[string]interface{})

	if name != nil {
		values["name"] = *name
	}

	if !startDate.IsZero() {
		values["start_date"] = startDate

	}

	if !endDate.IsZero() {
		values["end_date"] = endDate

	}

	if err := r.db.Model(&domain.Course{}).Where("id = ?", id).UpdateColumns(values).Error; err != nil {
		return err
	}

	return nil

}

func (r *repo) Count(filters Filters) (int, error) {
	var count int64
	tx := r.db.Model(&domain.Course{})
	tx = applyFilters(tx, filters)
	result := tx.Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return int(count), nil
}

func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {
	if filters.name != "" {
		filters.name = fmt.Sprintf("%%%s%%", strings.ToLower(filters.name))
		tx = tx.Where("lower(name) like ?", filters.name)
	}

	return tx
}
