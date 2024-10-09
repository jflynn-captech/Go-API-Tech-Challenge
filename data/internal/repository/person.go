package repository

import (
	"context"
	"fmt"
	"strings"

	"gorm.io/gorm"
	"jf.go.techchallenge.data/internal/applog"
	"jf.go.techchallenge.data/internal/models"
	pb "jf.go.techchallenge.data/protodata"
)

func NewPerson(db *gorm.DB, logger *applog.AppLogger) pb.PersonRepositoryServer {
	return PersonRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

type PersonRepositoryImpl struct {
	pb.UnimplementedPersonRepositoryServer
	db     *gorm.DB
	logger *applog.AppLogger
}

func personFromModel(modelPerson models.Person) *pb.Person {
	return &pb.Person{
		ID:        uint64(modelPerson.ID),
		Guid:      modelPerson.Guid,
		Email:     modelPerson.Email,
		FirstName: modelPerson.FirstName,
		LastName:  modelPerson.LastName,
		Age:       uint32(modelPerson.Age),
	}
}

func (s PersonRepositoryImpl) GetAll(ctx context.Context, filters *pb.Filters) (*pb.PersonList, error) {
	var persons []models.Person

	var pbPersons []*pb.Person

	tx := s.db.Table("person")

	// process query parameteres.
	for key, value := range filters.Filters {
		tx.Where(fmt.Sprintf("%s like ?", key), strings.Join([]string{"%", value, "%"}, ""))
	}

	result := tx.Find(&persons)

	for _, mPerson := range persons {
		pbPersons = append(pbPersons, personFromModel(mPerson))
	}

	return &pb.PersonList{People: pbPersons}, LogDBErr(s.logger, result.Error, "Failed to Query Person Table")
}

func (s PersonRepositoryImpl) GetByGuid(ctx context.Context, guid *pb.Guid) (*pb.Person, error) {
	modelPers, err := s.selectByGuid(guid.Guid)
	return personFromModel(modelPers), LogDBErr(s.logger, err, "Failed to Query Person Table")
}

func (s PersonRepositoryImpl) selectByGuid(guid string) (models.Person, error) {
	var person models.Person

	result := s.db.Table("person").Preload("Courses").Find(&person, "guid = ?", guid)

	if result.RowsAffected == 0 {
		return person, fmt.Errorf("Person: %s Not Found", guid)
	}
	return person, result.Error
}

// Used for both update and insert.
// In order for courses to be processed correctly, the persons courses must be "hydrated" on the person input parameter.
func (s PersonRepositoryImpl) Save(ctx context.Context, pbPerson *pb.Person) (*pb.Person, error) {

	person, err := s.selectByGuid(pbPerson.Guid)

	if err != nil {
		return nil, err
	}

	person.FirstName = pbPerson.FirstName
	person.LastName = pbPerson.LastName
	person.Email = pbPerson.Email
	person.Age = uint(pbPerson.Age)
	person.Type = pbPerson.Type

	var coursesToDelete []models.PersonCourse
	var coursesShouldExist []models.Course

	// Tracks courses state requested by user
	courseMap := map[int]interface{}{}
	for _, value := range pbPerson.Courses {
		courseMap[int(value.ID)] = nil
	}

	// If any person course is not present in courses Map, it should be deleted.
	for _, course := range person.Courses {
		_, present := courseMap[int(course.ID)]
		if !present {
			coursesToDelete = append(coursesToDelete, models.PersonCourse{
				PersonID: person.ID,
				CourseID: course.ID,
			})
		} else {
			coursesShouldExist = append(coursesShouldExist, models.Course{
				ID:   uint(course.ID),
				Guid: course.Guid,
				Name: course.Name,
			})
		}
	}

	person.Courses = coursesShouldExist

	return personFromModel(person), LogDBErr(s.logger, s.db.Transaction(func(tx *gorm.DB) error {

		var coursesError error
		if len(coursesToDelete) > 0 {
			coursesError = tx.Delete(&coursesToDelete).Error
		}

		if coursesError != nil {
			return coursesError
		}

		return tx.Save(person).Error
	}), "Failed to Save Person!")
}

func (s PersonRepositoryImpl) Delete(ctx context.Context, person *pb.Person) (*pb.Guid, error) {

	// Handle deleting the courses the person is enrolled in as well as the person.
	return &pb.Guid{Guid: person.Guid}, LogDBErr(s.logger, s.db.Transaction(func(tx *gorm.DB) error {

		result := tx.Delete(&models.PersonCourse{}, "person_id = ?", person.ID)
		if err := LogDBErr(s.logger, result.Error, "Failed to delete person_course record"); err != nil {
			return err
		}

		return tx.Delete(person).Error
	}), "Failed to delete person record")
}
