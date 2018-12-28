package student

import (
    "ee/schkola/person"
    "ee/schkola/shared"
    "encoding/json"
    "fmt"
    "github.com/eugeis/gee/enum"
    "github.com/looplab/eventhorizon"
    "gopkg.in/mgo.v2/bson"
    "time"
)
        
type Attendance struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Hours int `json:"hours" eh:"optional"`
    State *AttendanceState `json:"state" eh:"optional"`
    Token string `json:"token" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func NewAttendance() (ret *Attendance) {
    ret = &Attendance{}
    return
}
func (o *Attendance) EntityID() eventhorizon.UUID { return o.Id }



        
type Course struct {
    Name string `json:"name" eh:"optional"`
    Begin *time.Time `json:"begin" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Teacher *shared.PersonName `json:"teacher" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Fee float64 `json:"fee" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func NewCourse() (ret *Course) {
    ret = &Course{}
    return
}
func (o *Course) EntityID() eventhorizon.UUID { return o.Id }



        
type Grade struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Grade float64 `json:"grade" eh:"optional"`
    Comment string `json:"comment" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func NewGrade() (ret *Grade) {
    ret = &Grade{}
    return
}
func (o *Grade) EntityID() eventhorizon.UUID { return o.Id }



        
type Group struct {
    Name string `json:"name" eh:"optional"`
    Category *GroupCategory `json:"category" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Representative *person.Profile `json:"representative" eh:"optional"`
    Students []*person.Profile `json:"students" eh:"optional"`
    Courses []*Course `json:"courses" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func NewGroup() (ret *Group) {
    ret = &Group{}
    return
}

func (o *Group) AddToStudents(item *person.Profile) *person.Profile {
    o.Students = append(o.Students, item)
    return item
}

func (o *Group) AddToCourses(item *Course) *Course {
    o.Courses = append(o.Courses, item)
    return item
}
func (o *Group) EntityID() eventhorizon.UUID { return o.Id }



        
type SchoolApplication struct {
    Profile *person.Profile `json:"profile" eh:"optional"`
    ChurchContactPerson *shared.PersonName `json:"churchContactPerson" eh:"optional"`
    ChurchContact *person.Contact `json:"churchContact" eh:"optional"`
    ChurchCommitment bool `json:"churchCommitment" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Group string `json:"group" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func NewSchoolApplication() (ret *SchoolApplication) {
    ret = &SchoolApplication{}
    return
}
func (o *SchoolApplication) EntityID() eventhorizon.UUID { return o.Id }



        
type SchoolYear struct {
    Name string `json:"name" eh:"optional"`
    Start *time.Time `json:"start" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Dates []*time.Time `json:"dates" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func NewSchoolYear() (ret *SchoolYear) {
    ret = &SchoolYear{}
    return
}

func (o *SchoolYear) AddToDates(item *time.Time) *time.Time {
    o.Dates = append(o.Dates, item)
    return item
}
func (o *SchoolYear) EntityID() eventhorizon.UUID { return o.Id }









type AttendanceState struct {
	name  string
	ordinal int
}

func (o *AttendanceState) Name() string {
    return o.name
}

func (o *AttendanceState) Ordinal() int {
    return o.ordinal
}

func (o AttendanceState) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *AttendanceState) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := AttendanceStates().ParseAttendanceState(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid AttendanceState %q", lit.Name)
        }
	}
	return
}

func (o AttendanceState) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *AttendanceState) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := AttendanceStates().ParseAttendanceState(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid AttendanceState %q", lit)
        }
    }
    return
}

func (o *AttendanceState) IsRegistered() bool {
    return o == _attendanceStates.Registered()
}

func (o *AttendanceState) IsConfirmed() bool {
    return o == _attendanceStates.Confirmed()
}

func (o *AttendanceState) IsCanceled() bool {
    return o == _attendanceStates.Canceled()
}

func (o *AttendanceState) IsPresent() bool {
    return o == _attendanceStates.Present()
}

type attendanceStates struct {
	values []*AttendanceState
    literals []enum.Literal
}

var _attendanceStates = &attendanceStates{values: []*AttendanceState{
    {name: "Registered", ordinal: 0},
    {name: "Confirmed", ordinal: 1},
    {name: "Canceled", ordinal: 2},
    {name: "Present", ordinal: 3}},
}

func AttendanceStates() *attendanceStates {
	return _attendanceStates
}

func (o *attendanceStates) Values() []*AttendanceState {
	return o.values
}

func (o *attendanceStates) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *attendanceStates) Registered() *AttendanceState {
    return _attendanceStates.values[0]
}

func (o *attendanceStates) Confirmed() *AttendanceState {
    return _attendanceStates.values[1]
}

func (o *attendanceStates) Canceled() *AttendanceState {
    return _attendanceStates.values[2]
}

func (o *attendanceStates) Present() *AttendanceState {
    return _attendanceStates.values[3]
}

func (o *attendanceStates) ParseAttendanceState(name string) (ret *AttendanceState, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*AttendanceState), ok
	}
	return
}


type GroupCategory struct {
	name  string
	ordinal int
}

func (o *GroupCategory) Name() string {
    return o.name
}

func (o *GroupCategory) Ordinal() int {
    return o.ordinal
}

func (o GroupCategory) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *GroupCategory) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := GroupCategorys().ParseGroupCategory(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GroupCategory %q", lit.Name)
        }
	}
	return
}

func (o GroupCategory) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *GroupCategory) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := GroupCategorys().ParseGroupCategory(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GroupCategory %q", lit)
        }
    }
    return
}

func (o *GroupCategory) IsCourseGroup() bool {
    return o == _groupCategorys.CourseGroup()
}

func (o *GroupCategory) IsYearGroup() bool {
    return o == _groupCategorys.YearGroup()
}

type groupCategorys struct {
	values []*GroupCategory
    literals []enum.Literal
}

var _groupCategorys = &groupCategorys{values: []*GroupCategory{
    {name: "CourseGroup", ordinal: 0},
    {name: "YearGroup", ordinal: 1}},
}

func GroupCategorys() *groupCategorys {
	return _groupCategorys
}

func (o *groupCategorys) Values() []*GroupCategory {
	return o.values
}

func (o *groupCategorys) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *groupCategorys) CourseGroup() *GroupCategory {
    return _groupCategorys.values[0]
}

func (o *groupCategorys) YearGroup() *GroupCategory {
    return _groupCategorys.values[1]
}

func (o *groupCategorys) ParseGroupCategory(name string) (ret *GroupCategory, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*GroupCategory), ok
	}
	return
}



