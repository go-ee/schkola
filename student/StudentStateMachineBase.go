package student

import (
    "errors"
    "fmt"
    "github.com/looplab/eventhorizon"
)
type AttendanceInitialHandler struct {
}

func NewAttendanceInitialHandler() (ret *AttendanceInitialHandler) {
    ret = &AttendanceInitialHandler{}
    return
}

func (o *AttendanceInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    
    return
}

func (o *AttendanceInitialHandler) SetupEventHandler() (err error) {
    return
}


type AttendanceInitialExecutor struct {
}

func NewAttendanceInitialExecutor() (ret *AttendanceInitialExecutor) {
    ret = &AttendanceInitialExecutor{}
    return
}


type AttendanceHandlers struct {
    Initial *AttendanceInitialHandler `json:"initial" eh:"optional"`
}

func NewAttendanceHandlers() (ret *AttendanceHandlers) {
    initial := NewAttendanceInitialHandler()
    ret = &AttendanceHandlers{
        Initial: initial,
    }
    return
}


type AttendanceExecutors struct {
    Initial *AttendanceInitialExecutor `json:"initial" eh:"optional"`
}

func NewAttendanceExecutors() (ret *AttendanceExecutors) {
    initial := NewAttendanceInitialExecutor()
    ret = &AttendanceExecutors{
        Initial: initial,
    }
    return
}


type CourseInitialHandler struct {
}

func NewCourseInitialHandler() (ret *CourseInitialHandler) {
    ret = &CourseInitialHandler{}
    return
}

func (o *CourseInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    
    return
}

func (o *CourseInitialHandler) SetupEventHandler() (err error) {
    return
}


type CourseInitialExecutor struct {
}

func NewCourseInitialExecutor() (ret *CourseInitialExecutor) {
    ret = &CourseInitialExecutor{}
    return
}


type CourseHandlers struct {
    Initial *CourseInitialHandler `json:"initial" eh:"optional"`
}

func NewCourseHandlers() (ret *CourseHandlers) {
    initial := NewCourseInitialHandler()
    ret = &CourseHandlers{
        Initial: initial,
    }
    return
}


type CourseExecutors struct {
    Initial *CourseInitialExecutor `json:"initial" eh:"optional"`
}

func NewCourseExecutors() (ret *CourseExecutors) {
    initial := NewCourseInitialExecutor()
    ret = &CourseExecutors{
        Initial: initial,
    }
    return
}


type GradeInitialHandler struct {
}

func NewGradeInitialHandler() (ret *GradeInitialHandler) {
    ret = &GradeInitialHandler{}
    return
}

func (o *GradeInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    
    return
}

func (o *GradeInitialHandler) SetupEventHandler() (err error) {
    return
}


type GradeInitialExecutor struct {
}

func NewGradeInitialExecutor() (ret *GradeInitialExecutor) {
    ret = &GradeInitialExecutor{}
    return
}


type GradeHandlers struct {
    Initial *GradeInitialHandler `json:"initial" eh:"optional"`
}

func NewGradeHandlers() (ret *GradeHandlers) {
    initial := NewGradeInitialHandler()
    ret = &GradeHandlers{
        Initial: initial,
    }
    return
}


type GradeExecutors struct {
    Initial *GradeInitialExecutor `json:"initial" eh:"optional"`
}

func NewGradeExecutors() (ret *GradeExecutors) {
    initial := NewGradeInitialExecutor()
    ret = &GradeExecutors{
        Initial: initial,
    }
    return
}


type GroupInitialHandler struct {
}

func NewGroupInitialHandler() (ret *GroupInitialHandler) {
    ret = &GroupInitialHandler{}
    return
}

func (o *GroupInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    
    return
}

func (o *GroupInitialHandler) SetupEventHandler() (err error) {
    return
}


type GroupInitialExecutor struct {
}

func NewGroupInitialExecutor() (ret *GroupInitialExecutor) {
    ret = &GroupInitialExecutor{}
    return
}


type GroupHandlers struct {
    Initial *GroupInitialHandler `json:"initial" eh:"optional"`
}

func NewGroupHandlers() (ret *GroupHandlers) {
    initial := NewGroupInitialHandler()
    ret = &GroupHandlers{
        Initial: initial,
    }
    return
}


type GroupExecutors struct {
    Initial *GroupInitialExecutor `json:"initial" eh:"optional"`
}

func NewGroupExecutors() (ret *GroupExecutors) {
    initial := NewGroupInitialExecutor()
    ret = &GroupExecutors{
        Initial: initial,
    }
    return
}


type SchoolApplicationInitialHandler struct {
}

func NewSchoolApplicationInitialHandler() (ret *SchoolApplicationInitialHandler) {
    ret = &SchoolApplicationInitialHandler{}
    return
}

func (o *SchoolApplicationInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    
    return
}

func (o *SchoolApplicationInitialHandler) SetupEventHandler() (err error) {
    return
}


type SchoolApplicationInitialExecutor struct {
}

func NewSchoolApplicationInitialExecutor() (ret *SchoolApplicationInitialExecutor) {
    ret = &SchoolApplicationInitialExecutor{}
    return
}


type SchoolApplicationHandlers struct {
    Initial *SchoolApplicationInitialHandler `json:"initial" eh:"optional"`
}

func NewSchoolApplicationHandlers() (ret *SchoolApplicationHandlers) {
    initial := NewSchoolApplicationInitialHandler()
    ret = &SchoolApplicationHandlers{
        Initial: initial,
    }
    return
}


type SchoolApplicationExecutors struct {
    Initial *SchoolApplicationInitialExecutor `json:"initial" eh:"optional"`
}

func NewSchoolApplicationExecutors() (ret *SchoolApplicationExecutors) {
    initial := NewSchoolApplicationInitialExecutor()
    ret = &SchoolApplicationExecutors{
        Initial: initial,
    }
    return
}


type SchoolYearInitialHandler struct {
}

func NewSchoolYearInitialHandler() (ret *SchoolYearInitialHandler) {
    ret = &SchoolYearInitialHandler{}
    return
}

func (o *SchoolYearInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    
    return
}

func (o *SchoolYearInitialHandler) SetupEventHandler() (err error) {
    return
}


type SchoolYearInitialExecutor struct {
}

func NewSchoolYearInitialExecutor() (ret *SchoolYearInitialExecutor) {
    ret = &SchoolYearInitialExecutor{}
    return
}


type SchoolYearHandlers struct {
    Initial *SchoolYearInitialHandler `json:"initial" eh:"optional"`
}

func NewSchoolYearHandlers() (ret *SchoolYearHandlers) {
    initial := NewSchoolYearInitialHandler()
    ret = &SchoolYearHandlers{
        Initial: initial,
    }
    return
}


type SchoolYearExecutors struct {
    Initial *SchoolYearInitialExecutor `json:"initial" eh:"optional"`
}

func NewSchoolYearExecutors() (ret *SchoolYearExecutors) {
    initial := NewSchoolYearInitialExecutor()
    ret = &SchoolYearExecutors{
        Initial: initial,
    }
    return
}









