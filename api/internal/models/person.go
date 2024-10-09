package models

type PersonInput struct {
	FirstName   string
	LastName    string
	Email       string
	Age         uint
	Type        string
	CourseGuids []string
}
