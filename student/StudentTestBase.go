package student

import (
    "github.com/go-ee/schkola/person"
    "fmt"
    "github.com/go-ee/utils"
    "github.com/google/uuid"
    "time"
)
func NewAttendancesByPropNames(count int) []*Attendance {
	items := make([]*Attendance, count)
	for i := 0; i < count; i++ {
		items[i] = NewAttendanceByPropNames(i)
	}
	return items
}

func NewAttendanceByPropNames(intSalt int) (ret *Attendance)  {
    ret = NewAttendance()
    ret.Student = person.NewProfileByPropNames(intSalt)
    ret.Date = utils.PtrTime(time.Now())
    ret.Course = NewCourseByPropNames(intSalt)
    ret.Hours = intSalt
    ret.State = AttendanceStates().Registered()
    ret.Token = fmt.Sprintf("Token %v", intSalt)
    ret.Id = uuid.New()
    return
}


func NewCoursesByPropNames(count int) []*Course {
	items := make([]*Course, count)
	for i := 0; i < count; i++ {
		items[i] = NewCourseByPropNames(i)
	}
	return items
}

func NewCourseByPropNames(intSalt int) (ret *Course)  {
    ret = NewCourse()
    ret.Name = fmt.Sprintf("Name %v", intSalt)
    ret.Begin = utils.PtrTime(time.Now())
    ret.End = utils.PtrTime(time.Now())
    ret.Teacher = person.NewPersonNameByPropNames(intSalt)
    ret.SchoolYear = NewSchoolYearByPropNames(intSalt)
    ret.Fee = float64(intSalt)
    ret.Description = fmt.Sprintf("Description %v", intSalt)
    ret.Id = uuid.New()
    return
}


func NewGradesByPropNames(count int) []*Grade {
	items := make([]*Grade, count)
	for i := 0; i < count; i++ {
		items[i] = NewGradeByPropNames(i)
	}
	return items
}

func NewGradeByPropNames(intSalt int) (ret *Grade)  {
    ret = NewGrade()
    ret.Student = person.NewProfileByPropNames(intSalt)
    ret.Course = NewCourseByPropNames(intSalt)
    ret.Grade = float64(intSalt)
    ret.Comment = fmt.Sprintf("Comment %v", intSalt)
    ret.Id = uuid.New()
    return
}


func NewGroupsByPropNames(count int) []*Group {
	items := make([]*Group, count)
	for i := 0; i < count; i++ {
		items[i] = NewGroupByPropNames(i)
	}
	return items
}

func NewGroupByPropNames(intSalt int) (ret *Group)  {
    ret = NewGroup()
    ret.Name = fmt.Sprintf("Name %v", intSalt)
    ret.Category = GroupCategorys().CourseGroup()
    ret.SchoolYear = NewSchoolYearByPropNames(intSalt)
    ret.Representative = person.NewProfileByPropNames(intSalt)
    ret.Students = []*person.Profile{}
    ret.Courses = []*Course{}
    ret.Id = uuid.New()
    return
}


func NewSchoolApplicationsByPropNames(count int) []*SchoolApplication {
	items := make([]*SchoolApplication, count)
	for i := 0; i < count; i++ {
		items[i] = NewSchoolApplicationByPropNames(i)
	}
	return items
}

func NewSchoolApplicationByPropNames(intSalt int) (ret *SchoolApplication)  {
    ret = NewSchoolApplication()
    ret.Profile = person.NewProfileByPropNames(intSalt)
    ret.ChurchContactPerson = person.NewPersonNameByPropNames(intSalt)
    ret.ChurchContact = person.NewContactByPropNames(intSalt)
    ret.ChurchCommitment = false
    ret.SchoolYear = NewSchoolYearByPropNames(intSalt)
    ret.Group = fmt.Sprintf("Group %v", intSalt)
    ret.Id = uuid.New()
    return
}


func NewSchoolYearsByPropNames(count int) []*SchoolYear {
	items := make([]*SchoolYear, count)
	for i := 0; i < count; i++ {
		items[i] = NewSchoolYearByPropNames(i)
	}
	return items
}

func NewSchoolYearByPropNames(intSalt int) (ret *SchoolYear)  {
    ret = NewSchoolYear()
    ret.Name = fmt.Sprintf("Name %v", intSalt)
    ret.Start = utils.PtrTime(time.Now())
    ret.End = utils.PtrTime(time.Now())
    ret.Dates = []*time.Time{}
    ret.Id = uuid.New()
    return
}







