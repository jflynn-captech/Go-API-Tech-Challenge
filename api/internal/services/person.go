package services

import (
	"context"
	"net/mail"
	"net/url"
	"slices"
	"strings"

	"github.com/google/uuid"
	"jf.go.techchallenge/internal/apperror"
	"jf.go.techchallenge/internal/applog"
	"jf.go.techchallenge/internal/models"
	"jf.go.techchallenge/protodata"
)

type Person struct {
	logger       *applog.AppLogger
	personClient protodata.PersonRepositoryClient
	courseClient protodata.CourseRepositoryClient
}

func NewPerson(logger *applog.AppLogger, personClient protodata.PersonRepositoryClient, courseClient protodata.CourseRepositoryClient) *Person {
	return &Person{
		logger:       logger,
		personClient: personClient,
		courseClient: courseClient,
	}
}

// Parse dont validate https://lexi-lambda.github.io/blog/2019/11/05/parse-don-t-validate/
func (s Person) parse(input models.PersonInput, person *protodata.Person) []error {
	var errors []error

	if strings.Trim(input.FirstName, " ") == "" {
		errors = append(errors, apperror.BadRequest("First Name must not be blank"))
	}

	if strings.Trim(input.LastName, " ") == "" {
		errors = append(errors, apperror.BadRequest("Last Name must not be blank"))
	}

	// Validate email address format
	_, emailErr := mail.ParseAddress(input.Email)
	if emailErr != nil {
		errors = append(errors, apperror.BadRequest("Email must be a valid email address"))
	}

	if input.Age < 10 {
		errors = append(errors, apperror.BadRequest("Must be at least 10 years old to enrol."))
	}

	if !slices.Contains([]string{"professor", "student"}, input.Type) {
		errors = append(errors, apperror.BadRequest("Invalid Person type, must be either 'professor' or 'student'"))
	}

	if len(errors) > 0 {
		return errors
	}

	person.FirstName = input.FirstName
	person.LastName = input.LastName
	person.Email = input.Email
	person.Type = string(input.Type)
	person.Age = uint32(input.Age)

	return nil
}

// Validates all CourseGuids are valid and exist.
func (s Person) parseCourses(input models.PersonInput) ([]*protodata.Course, []error) {
	// Check course guids for validity. whether they should be added or removed is left up to the repository.
	var courses []*protodata.Course
	var errors []error
	for _, courseGuid := range input.CourseGuids {
		course, err := s.courseClient.GetByGuid(context.Background(), &protodata.Guid{Guid: courseGuid})
		if err != nil {
			errors = append(errors, err)
		} else {
			courses = append(courses, course)
		}
	}
	return courses, errors
}

func (s Person) GetOneByGuid(guid string) (*protodata.Person, error) {
	return s.personClient.GetByGuid(context.Background(), &protodata.Guid{Guid: guid})
}

func (s Person) Update(guid string, input models.PersonInput) (*protodata.Person, error) {
	person, err := s.GetOneByGuid(guid)
	if err != nil {
		return person, err
	}

	errs := s.parse(input, person)

	courses, courseErrs := s.parseCourses(input)

	if errs != nil || courseErrs != nil {
		return person, apperror.Of(append(errs, courseErrs...))
	}
	person.Courses = courses
	return s.personClient.Save(context.Background(), person)
}

func (s Person) Delete(guid string) error {
	person, err := s.GetOneByGuid(guid)
	if err != nil {
		return err
	}
	_, err = s.personClient.Delete(context.Background(), person)

	return err
}

func (s Person) Create(input models.PersonInput) (*protodata.Person, error) {
	newPerson := protodata.Person{}

	errs := s.parse(input, &newPerson)

	courses, courseErrs := s.parseCourses(input)

	if errs != nil || courseErrs != nil {
		return nil, apperror.Of(append(errs, courseErrs...))
	}

	newPerson.Guid = uuid.NewString()
	newPerson.Courses = courses

	return s.personClient.Save(context.Background(), &newPerson)
}

var personFilters = MakeFilterColumns(ValidFilters{
	"FirstName",
	"LastName",
	"Email",
})

func (s Person) GetAll(urlParams url.Values) ([]*protodata.Person, error) {

	filters, err := ParseURLFilters(urlParams, personFilters)

	if err != nil {
		return nil, err
	}

	personList, err := s.personClient.GetAll(context.Background(), &protodata.Filters{Filters: filters})
	if err != nil {
		return nil, err
	}
	return personList.People, err
}
