package student

import (
    "context"
    "github.com/looplab/eventhorizon"
)
type AttendanceQueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewAttendanceQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *AttendanceQueryRepository) {
    ret = &AttendanceQueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *AttendanceQueryRepository) FindAll() (ret []*Attendance, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Attendance, len(result))
		for i, e := range result {
            ret[i] = e.(*Attendance)
		}
    }
    return
}

func (o *AttendanceQueryRepository) FindById(id eventhorizon.UUID) (ret *Attendance, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Attendance)
    }
    return
}

func (o *AttendanceQueryRepository) CountAll() (ret int, err error) {
    var result []*Attendance
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *AttendanceQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Attendance
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *AttendanceQueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *AttendanceQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}


type CourseQueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewCourseQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *CourseQueryRepository) {
    ret = &CourseQueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *CourseQueryRepository) FindAll() (ret []*Course, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Course, len(result))
		for i, e := range result {
            ret[i] = e.(*Course)
		}
    }
    return
}

func (o *CourseQueryRepository) FindById(id eventhorizon.UUID) (ret *Course, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Course)
    }
    return
}

func (o *CourseQueryRepository) CountAll() (ret int, err error) {
    var result []*Course
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *CourseQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Course
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *CourseQueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *CourseQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}


type GradeQueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewGradeQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *GradeQueryRepository) {
    ret = &GradeQueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *GradeQueryRepository) FindAll() (ret []*Grade, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Grade, len(result))
		for i, e := range result {
            ret[i] = e.(*Grade)
		}
    }
    return
}

func (o *GradeQueryRepository) FindById(id eventhorizon.UUID) (ret *Grade, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Grade)
    }
    return
}

func (o *GradeQueryRepository) CountAll() (ret int, err error) {
    var result []*Grade
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *GradeQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Grade
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *GradeQueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *GradeQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}


type GroupQueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewGroupQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *GroupQueryRepository) {
    ret = &GroupQueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *GroupQueryRepository) FindAll() (ret []*Group, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*Group, len(result))
		for i, e := range result {
            ret[i] = e.(*Group)
		}
    }
    return
}

func (o *GroupQueryRepository) FindById(id eventhorizon.UUID) (ret *Group, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*Group)
    }
    return
}

func (o *GroupQueryRepository) CountAll() (ret int, err error) {
    var result []*Group
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *GroupQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *Group
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *GroupQueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *GroupQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}


type SchoolApplicationQueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewSchoolApplicationQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *SchoolApplicationQueryRepository) {
    ret = &SchoolApplicationQueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *SchoolApplicationQueryRepository) FindAll() (ret []*SchoolApplication, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*SchoolApplication, len(result))
		for i, e := range result {
            ret[i] = e.(*SchoolApplication)
		}
    }
    return
}

func (o *SchoolApplicationQueryRepository) FindById(id eventhorizon.UUID) (ret *SchoolApplication, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*SchoolApplication)
    }
    return
}

func (o *SchoolApplicationQueryRepository) CountAll() (ret int, err error) {
    var result []*SchoolApplication
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *SchoolApplicationQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *SchoolApplication
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *SchoolApplicationQueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *SchoolApplicationQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}


type SchoolYearQueryRepository struct {
    repo eventhorizon.ReadRepo `json:"repo" eh:"optional"`
    context context.Context `json:"context" eh:"optional"`
}

func NewSchoolYearQueryRepository(repo eventhorizon.ReadRepo, context context.Context) (ret *SchoolYearQueryRepository) {
    ret = &SchoolYearQueryRepository{
        repo: repo,
        context: context,
    }
    return
}

func (o *SchoolYearQueryRepository) FindAll() (ret []*SchoolYear, err error) {
    var result []eventhorizon.Entity
	if result, err = o.repo.FindAll(o.context); err == nil {
        ret = make([]*SchoolYear, len(result))
		for i, e := range result {
            ret[i] = e.(*SchoolYear)
		}
    }
    return
}

func (o *SchoolYearQueryRepository) FindById(id eventhorizon.UUID) (ret *SchoolYear, err error) {
    var result eventhorizon.Entity
	if result, err = o.repo.Find(o.context, id); err == nil {
        ret = result.(*SchoolYear)
    }
    return
}

func (o *SchoolYearQueryRepository) CountAll() (ret int, err error) {
    var result []*SchoolYear
	if result, err = o.FindAll(); err == nil {
        ret = len(result)
    }
    return
}

func (o *SchoolYearQueryRepository) CountById(id eventhorizon.UUID) (ret int, err error) {
    var result *SchoolYear
	if result, err = o.FindById(id); err == nil && result != nil {
        ret = 1
    }
    return
}

func (o *SchoolYearQueryRepository) ExistAll() (ret bool, err error) {
    var result int
	if result, err = o.CountAll(); err == nil {
        ret = result > 0
    }
    return
}

func (o *SchoolYearQueryRepository) ExistById(id eventhorizon.UUID) (ret bool, err error) {
    var result int
	if result, err = o.CountById(id); err == nil {
        ret = result > 0
    }
    return
}









