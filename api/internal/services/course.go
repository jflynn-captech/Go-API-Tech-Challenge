package services

import (
	"context"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"jf.go.techchallenge/internal/apperror"
	"jf.go.techchallenge/internal/applog"
	"jf.go.techchallenge/internal/models"
	"jf.go.techchallenge/protodata"
)

type Course struct {
	logger       *applog.AppLogger
	courseClient protodata.CourseRepositoryClient
}

func NewCourse(logger *applog.AppLogger, courseClient protodata.CourseRepositoryClient) *Course {
	return &Course{
		logger:       logger,
		courseClient: courseClient,
	}
}

func (s Course) GetOneByGuid(guid string) (*protodata.Course, error) {
	return s.courseClient.GetByGuid(context.Background(), &protodata.Guid{Guid: guid})
}

func (s Course) Update(guid string, input models.CourseInput) (*protodata.Course, error) {
	course, err := s.GetOneByGuid(guid)
	if err != nil {
		return course, err
	}

	err = s.parse(input, course)

	if err != nil {
		return course, err
	}

	return s.courseClient.Save(context.Background(), course)
}

func (s Course) Delete(guid string) error {
	course, err := s.GetOneByGuid(guid)
	if err != nil {
		return err
	}

	_, err = s.courseClient.Delete(context.Background(), course)
	return err
}

func (s Course) Create(input models.CourseInput) (*protodata.Course, error) {
	newCourse := protodata.Course{}

	err := s.parse(input, &newCourse)

	if err != nil {
		return nil, err
	}
	newCourse.Guid = uuid.NewString()
	return s.courseClient.Save(context.Background(), &newCourse)
}

var courseFilters = MakeFilterColumns(ValidFilters{
	"Name",
})

func (s Course) GetAll(urlParams url.Values) ([]*protodata.Course, error) {
	filters, err := ParseURLFilters(urlParams, courseFilters)

	if err != nil {
		return nil, err
	}

	list, err := s.courseClient.GetAll(context.Background(), &protodata.Filters{Filters: filters})
	if err != nil {
		return nil, err
	}
	return list.Courses, err
}

func (s Course) parse(input models.CourseInput, course *protodata.Course) error {
	var errors []error

	if strings.Trim(input.Name, " ") == "" {
		errors = append(errors, apperror.BadRequest("Name must not be blank"))
	}

	if len(errors) > 0 {
		return apperror.Of(errors)
	}

	course.Name = input.Name

	return nil
}
