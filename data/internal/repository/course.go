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

func NewCourse(db *gorm.DB, logger *applog.AppLogger) pb.CourseRepositoryServer {
	return CourseRepositoryImpl{
		db:     db,
		logger: logger,
	}
}

type CourseRepositoryImpl struct {
	pb.UnimplementedCourseRepositoryServer
	db     *gorm.DB
	logger *applog.AppLogger
}

func courseFromModel(modelCourse models.Course) *pb.Course {
	return &pb.Course{
		ID:   uint64(modelCourse.ID),
		Guid: modelCourse.Guid,
		Name: modelCourse.Name,
	}
}

func (s CourseRepositoryImpl) GetAll(ctx context.Context, filters *pb.Filters) (*pb.CourseList, error) {
	var courses []models.Course
	tx := s.db.Table("course")

	for key, value := range filters.Filters {
		tx.Where(fmt.Sprintf("%s like ?", key), strings.Join([]string{"%", value, "%"}, ""))
	}

	result := tx.Find(&courses)

	var pbCourses []*pb.Course
	for _, course := range courses {
		pbCourses = append(pbCourses, courseFromModel(course))
	}

	return &pb.CourseList{Courses: pbCourses}, LogDBErr(s.logger, result.Error, "Failed to query courses table")
}

func (s CourseRepositoryImpl) GetByGuid(ctx context.Context, guid *pb.Guid) (*pb.Course, error) {
	var course models.Course

	result := s.db.Table("course").Find(&course, "guid = ?", guid)

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Course: %s Not Found", guid)
	}

	return courseFromModel(course), LogDBErr(s.logger, result.Error, "Failed to Query Course Table")
}

func (s CourseRepositoryImpl) Save(ctx context.Context, course *pb.Course) (*pb.Course, error) {
	saveErr := LogDBErr(s.logger, s.db.Save(course).Error, "Failed to Save Course")
	if saveErr != nil {
		return nil, saveErr
	}
	return s.GetByGuid(ctx, &pb.Guid{Guid: course.Guid})
}

func (s CourseRepositoryImpl) Delete(ctx context.Context, course *pb.Course) (*pb.Guid, error) {
	// Handle deleting the courses
	return &pb.Guid{Guid: course.Guid}, s.db.Transaction(func(tx *gorm.DB) error {
		result := tx.Delete(&models.PersonCourse{}, "course_id = ?", course.ID)

		if err := LogDBErr(s.logger, result.Error, "Failed to delete person_course for course record"); err != nil {
			return err
		}

		return LogDBErr(s.logger, tx.Delete(course).Error, "Failed to delete course record")
	})
}
