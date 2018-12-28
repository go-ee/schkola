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
const (
     AttendanceRegisteredEvent eventhorizon.EventType = "AttendanceRegistered"
     AttendanceCreatedEvent eventhorizon.EventType = "AttendanceCreated"
     AttendanceDeletedEvent eventhorizon.EventType = "AttendanceDeleted"
     AttendanceConfirmedEvent eventhorizon.EventType = "AttendanceConfirmed"
     AttendanceCanceledEvent eventhorizon.EventType = "AttendanceCanceled"
     AttendanceUpdatedEvent eventhorizon.EventType = "AttendanceUpdated"
)


const (
     CourseCreatedEvent eventhorizon.EventType = "CourseCreated"
     CourseDeletedEvent eventhorizon.EventType = "CourseDeleted"
     CourseUpdatedEvent eventhorizon.EventType = "CourseUpdated"
)


const (
     GradeCreatedEvent eventhorizon.EventType = "GradeCreated"
     GradeDeletedEvent eventhorizon.EventType = "GradeDeleted"
     GradeUpdatedEvent eventhorizon.EventType = "GradeUpdated"
)


const (
     GroupCreatedEvent eventhorizon.EventType = "GroupCreated"
     GroupDeletedEvent eventhorizon.EventType = "GroupDeleted"
     GroupUpdatedEvent eventhorizon.EventType = "GroupUpdated"
)


const (
     SchoolApplicationCreatedEvent eventhorizon.EventType = "SchoolApplicationCreated"
     SchoolApplicationDeletedEvent eventhorizon.EventType = "SchoolApplicationDeleted"
     SchoolApplicationUpdatedEvent eventhorizon.EventType = "SchoolApplicationUpdated"
)


const (
     SchoolYearCreatedEvent eventhorizon.EventType = "SchoolYearCreated"
     SchoolYearDeletedEvent eventhorizon.EventType = "SchoolYearDeleted"
     SchoolYearUpdatedEvent eventhorizon.EventType = "SchoolYearUpdated"
)




type AttendanceRegistered struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type AttendanceCreated struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Hours int `json:"hours" eh:"optional"`
    State *AttendanceState `json:"state" eh:"optional"`
    Token string `json:"token" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type AttendanceDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type AttendanceConfirmed struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type AttendanceCanceled struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type AttendanceUpdated struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Hours int `json:"hours" eh:"optional"`
    State *AttendanceState `json:"state" eh:"optional"`
    Token string `json:"token" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type CourseCreated struct {
    Name string `json:"name" eh:"optional"`
    Begin *time.Time `json:"begin" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Teacher *shared.PersonName `json:"teacher" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Fee float64 `json:"fee" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type CourseDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type CourseUpdated struct {
    Name string `json:"name" eh:"optional"`
    Begin *time.Time `json:"begin" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Teacher *shared.PersonName `json:"teacher" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Fee float64 `json:"fee" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type GradeCreated struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Grade float64 `json:"grade" eh:"optional"`
    Comment string `json:"comment" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type GradeDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type GradeUpdated struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Grade float64 `json:"grade" eh:"optional"`
    Comment string `json:"comment" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type GroupCreated struct {
    Name string `json:"name" eh:"optional"`
    Category *GroupCategory `json:"category" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Representative *person.Profile `json:"representative" eh:"optional"`
    Students []*person.Profile `json:"students" eh:"optional"`
    Courses []*Course `json:"courses" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *GroupCreated) AddToStudents(item *person.Profile) *person.Profile {
    o.Students = append(o.Students, item)
    return item
}

func (o *GroupCreated) AddToCourses(item *Course) *Course {
    o.Courses = append(o.Courses, item)
    return item
}


type GroupDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type GroupUpdated struct {
    Name string `json:"name" eh:"optional"`
    Category *GroupCategory `json:"category" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Representative *person.Profile `json:"representative" eh:"optional"`
    Students []*person.Profile `json:"students" eh:"optional"`
    Courses []*Course `json:"courses" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *GroupUpdated) AddToStudents(item *person.Profile) *person.Profile {
    o.Students = append(o.Students, item)
    return item
}

func (o *GroupUpdated) AddToCourses(item *Course) *Course {
    o.Courses = append(o.Courses, item)
    return item
}


type SchoolApplicationCreated struct {
    Profile *person.Profile `json:"profile" eh:"optional"`
    ChurchContactPerson *shared.PersonName `json:"churchContactPerson" eh:"optional"`
    ChurchContact *person.Contact `json:"churchContact" eh:"optional"`
    ChurchCommitment bool `json:"churchCommitment" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Group string `json:"group" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SchoolApplicationDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SchoolApplicationUpdated struct {
    Profile *person.Profile `json:"profile" eh:"optional"`
    ChurchContactPerson *shared.PersonName `json:"churchContactPerson" eh:"optional"`
    ChurchContact *person.Contact `json:"churchContact" eh:"optional"`
    ChurchCommitment bool `json:"churchCommitment" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Group string `json:"group" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SchoolYearCreated struct {
    Name string `json:"name" eh:"optional"`
    Start *time.Time `json:"start" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Dates []*time.Time `json:"dates" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *SchoolYearCreated) AddToDates(item *time.Time) *time.Time {
    o.Dates = append(o.Dates, item)
    return item
}


type SchoolYearDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SchoolYearUpdated struct {
    Name string `json:"name" eh:"optional"`
    Start *time.Time `json:"start" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Dates []*time.Time `json:"dates" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *SchoolYearUpdated) AddToDates(item *time.Time) *time.Time {
    o.Dates = append(o.Dates, item)
    return item
}




type AttendanceEventType struct {
	name  string
	ordinal int
}

func (o *AttendanceEventType) Name() string {
    return o.name
}

func (o *AttendanceEventType) Ordinal() int {
    return o.ordinal
}

func (o AttendanceEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *AttendanceEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := AttendanceEventTypes().ParseAttendanceEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid AttendanceEventType %q", lit.Name)
        }
	}
	return
}

func (o AttendanceEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *AttendanceEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := AttendanceEventTypes().ParseAttendanceEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid AttendanceEventType %q", lit)
        }
    }
    return
}

func (o *AttendanceEventType) IsAttendanceRegistered() bool {
    return o == _attendanceEventTypes.AttendanceRegistered()
}

func (o *AttendanceEventType) IsAttendanceCreated() bool {
    return o == _attendanceEventTypes.AttendanceCreated()
}

func (o *AttendanceEventType) IsAttendanceDeleted() bool {
    return o == _attendanceEventTypes.AttendanceDeleted()
}

func (o *AttendanceEventType) IsAttendanceConfirmed() bool {
    return o == _attendanceEventTypes.AttendanceConfirmed()
}

func (o *AttendanceEventType) IsAttendanceCanceled() bool {
    return o == _attendanceEventTypes.AttendanceCanceled()
}

func (o *AttendanceEventType) IsAttendanceUpdated() bool {
    return o == _attendanceEventTypes.AttendanceUpdated()
}

type attendanceEventTypes struct {
	values []*AttendanceEventType
    literals []enum.Literal
}

var _attendanceEventTypes = &attendanceEventTypes{values: []*AttendanceEventType{
    {name: "AttendanceRegistered", ordinal: 0},
    {name: "AttendanceCreated", ordinal: 1},
    {name: "AttendanceDeleted", ordinal: 2},
    {name: "AttendanceConfirmed", ordinal: 3},
    {name: "AttendanceCanceled", ordinal: 4},
    {name: "AttendanceUpdated", ordinal: 5}},
}

func AttendanceEventTypes() *attendanceEventTypes {
	return _attendanceEventTypes
}

func (o *attendanceEventTypes) Values() []*AttendanceEventType {
	return o.values
}

func (o *attendanceEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *attendanceEventTypes) AttendanceRegistered() *AttendanceEventType {
    return _attendanceEventTypes.values[0]
}

func (o *attendanceEventTypes) AttendanceCreated() *AttendanceEventType {
    return _attendanceEventTypes.values[1]
}

func (o *attendanceEventTypes) AttendanceDeleted() *AttendanceEventType {
    return _attendanceEventTypes.values[2]
}

func (o *attendanceEventTypes) AttendanceConfirmed() *AttendanceEventType {
    return _attendanceEventTypes.values[3]
}

func (o *attendanceEventTypes) AttendanceCanceled() *AttendanceEventType {
    return _attendanceEventTypes.values[4]
}

func (o *attendanceEventTypes) AttendanceUpdated() *AttendanceEventType {
    return _attendanceEventTypes.values[5]
}

func (o *attendanceEventTypes) ParseAttendanceEventType(name string) (ret *AttendanceEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*AttendanceEventType), ok
	}
	return
}


type CourseEventType struct {
	name  string
	ordinal int
}

func (o *CourseEventType) Name() string {
    return o.name
}

func (o *CourseEventType) Ordinal() int {
    return o.ordinal
}

func (o CourseEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *CourseEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := CourseEventTypes().ParseCourseEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid CourseEventType %q", lit.Name)
        }
	}
	return
}

func (o CourseEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *CourseEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := CourseEventTypes().ParseCourseEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid CourseEventType %q", lit)
        }
    }
    return
}

func (o *CourseEventType) IsCourseCreated() bool {
    return o == _courseEventTypes.CourseCreated()
}

func (o *CourseEventType) IsCourseDeleted() bool {
    return o == _courseEventTypes.CourseDeleted()
}

func (o *CourseEventType) IsCourseUpdated() bool {
    return o == _courseEventTypes.CourseUpdated()
}

type courseEventTypes struct {
	values []*CourseEventType
    literals []enum.Literal
}

var _courseEventTypes = &courseEventTypes{values: []*CourseEventType{
    {name: "CourseCreated", ordinal: 0},
    {name: "CourseDeleted", ordinal: 1},
    {name: "CourseUpdated", ordinal: 2}},
}

func CourseEventTypes() *courseEventTypes {
	return _courseEventTypes
}

func (o *courseEventTypes) Values() []*CourseEventType {
	return o.values
}

func (o *courseEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *courseEventTypes) CourseCreated() *CourseEventType {
    return _courseEventTypes.values[0]
}

func (o *courseEventTypes) CourseDeleted() *CourseEventType {
    return _courseEventTypes.values[1]
}

func (o *courseEventTypes) CourseUpdated() *CourseEventType {
    return _courseEventTypes.values[2]
}

func (o *courseEventTypes) ParseCourseEventType(name string) (ret *CourseEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*CourseEventType), ok
	}
	return
}


type GradeEventType struct {
	name  string
	ordinal int
}

func (o *GradeEventType) Name() string {
    return o.name
}

func (o *GradeEventType) Ordinal() int {
    return o.ordinal
}

func (o GradeEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *GradeEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := GradeEventTypes().ParseGradeEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GradeEventType %q", lit.Name)
        }
	}
	return
}

func (o GradeEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *GradeEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := GradeEventTypes().ParseGradeEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GradeEventType %q", lit)
        }
    }
    return
}

func (o *GradeEventType) IsGradeCreated() bool {
    return o == _gradeEventTypes.GradeCreated()
}

func (o *GradeEventType) IsGradeDeleted() bool {
    return o == _gradeEventTypes.GradeDeleted()
}

func (o *GradeEventType) IsGradeUpdated() bool {
    return o == _gradeEventTypes.GradeUpdated()
}

type gradeEventTypes struct {
	values []*GradeEventType
    literals []enum.Literal
}

var _gradeEventTypes = &gradeEventTypes{values: []*GradeEventType{
    {name: "GradeCreated", ordinal: 0},
    {name: "GradeDeleted", ordinal: 1},
    {name: "GradeUpdated", ordinal: 2}},
}

func GradeEventTypes() *gradeEventTypes {
	return _gradeEventTypes
}

func (o *gradeEventTypes) Values() []*GradeEventType {
	return o.values
}

func (o *gradeEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *gradeEventTypes) GradeCreated() *GradeEventType {
    return _gradeEventTypes.values[0]
}

func (o *gradeEventTypes) GradeDeleted() *GradeEventType {
    return _gradeEventTypes.values[1]
}

func (o *gradeEventTypes) GradeUpdated() *GradeEventType {
    return _gradeEventTypes.values[2]
}

func (o *gradeEventTypes) ParseGradeEventType(name string) (ret *GradeEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*GradeEventType), ok
	}
	return
}


type GroupEventType struct {
	name  string
	ordinal int
}

func (o *GroupEventType) Name() string {
    return o.name
}

func (o *GroupEventType) Ordinal() int {
    return o.ordinal
}

func (o GroupEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *GroupEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := GroupEventTypes().ParseGroupEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GroupEventType %q", lit.Name)
        }
	}
	return
}

func (o GroupEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *GroupEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := GroupEventTypes().ParseGroupEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GroupEventType %q", lit)
        }
    }
    return
}

func (o *GroupEventType) IsGroupCreated() bool {
    return o == _groupEventTypes.GroupCreated()
}

func (o *GroupEventType) IsGroupDeleted() bool {
    return o == _groupEventTypes.GroupDeleted()
}

func (o *GroupEventType) IsGroupUpdated() bool {
    return o == _groupEventTypes.GroupUpdated()
}

type groupEventTypes struct {
	values []*GroupEventType
    literals []enum.Literal
}

var _groupEventTypes = &groupEventTypes{values: []*GroupEventType{
    {name: "GroupCreated", ordinal: 0},
    {name: "GroupDeleted", ordinal: 1},
    {name: "GroupUpdated", ordinal: 2}},
}

func GroupEventTypes() *groupEventTypes {
	return _groupEventTypes
}

func (o *groupEventTypes) Values() []*GroupEventType {
	return o.values
}

func (o *groupEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *groupEventTypes) GroupCreated() *GroupEventType {
    return _groupEventTypes.values[0]
}

func (o *groupEventTypes) GroupDeleted() *GroupEventType {
    return _groupEventTypes.values[1]
}

func (o *groupEventTypes) GroupUpdated() *GroupEventType {
    return _groupEventTypes.values[2]
}

func (o *groupEventTypes) ParseGroupEventType(name string) (ret *GroupEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*GroupEventType), ok
	}
	return
}


type SchoolApplicationEventType struct {
	name  string
	ordinal int
}

func (o *SchoolApplicationEventType) Name() string {
    return o.name
}

func (o *SchoolApplicationEventType) Ordinal() int {
    return o.ordinal
}

func (o SchoolApplicationEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *SchoolApplicationEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := SchoolApplicationEventTypes().ParseSchoolApplicationEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid SchoolApplicationEventType %q", lit.Name)
        }
	}
	return
}

func (o SchoolApplicationEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *SchoolApplicationEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := SchoolApplicationEventTypes().ParseSchoolApplicationEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid SchoolApplicationEventType %q", lit)
        }
    }
    return
}

func (o *SchoolApplicationEventType) IsSchoolApplicationCreated() bool {
    return o == _schoolApplicationEventTypes.SchoolApplicationCreated()
}

func (o *SchoolApplicationEventType) IsSchoolApplicationDeleted() bool {
    return o == _schoolApplicationEventTypes.SchoolApplicationDeleted()
}

func (o *SchoolApplicationEventType) IsSchoolApplicationUpdated() bool {
    return o == _schoolApplicationEventTypes.SchoolApplicationUpdated()
}

type schoolApplicationEventTypes struct {
	values []*SchoolApplicationEventType
    literals []enum.Literal
}

var _schoolApplicationEventTypes = &schoolApplicationEventTypes{values: []*SchoolApplicationEventType{
    {name: "SchoolApplicationCreated", ordinal: 0},
    {name: "SchoolApplicationDeleted", ordinal: 1},
    {name: "SchoolApplicationUpdated", ordinal: 2}},
}

func SchoolApplicationEventTypes() *schoolApplicationEventTypes {
	return _schoolApplicationEventTypes
}

func (o *schoolApplicationEventTypes) Values() []*SchoolApplicationEventType {
	return o.values
}

func (o *schoolApplicationEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *schoolApplicationEventTypes) SchoolApplicationCreated() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[0]
}

func (o *schoolApplicationEventTypes) SchoolApplicationDeleted() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[1]
}

func (o *schoolApplicationEventTypes) SchoolApplicationUpdated() *SchoolApplicationEventType {
    return _schoolApplicationEventTypes.values[2]
}

func (o *schoolApplicationEventTypes) ParseSchoolApplicationEventType(name string) (ret *SchoolApplicationEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*SchoolApplicationEventType), ok
	}
	return
}


type SchoolYearEventType struct {
	name  string
	ordinal int
}

func (o *SchoolYearEventType) Name() string {
    return o.name
}

func (o *SchoolYearEventType) Ordinal() int {
    return o.ordinal
}

func (o SchoolYearEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *SchoolYearEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := SchoolYearEventTypes().ParseSchoolYearEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid SchoolYearEventType %q", lit.Name)
        }
	}
	return
}

func (o SchoolYearEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *SchoolYearEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := SchoolYearEventTypes().ParseSchoolYearEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid SchoolYearEventType %q", lit)
        }
    }
    return
}

func (o *SchoolYearEventType) IsSchoolYearCreated() bool {
    return o == _schoolYearEventTypes.SchoolYearCreated()
}

func (o *SchoolYearEventType) IsSchoolYearDeleted() bool {
    return o == _schoolYearEventTypes.SchoolYearDeleted()
}

func (o *SchoolYearEventType) IsSchoolYearUpdated() bool {
    return o == _schoolYearEventTypes.SchoolYearUpdated()
}

type schoolYearEventTypes struct {
	values []*SchoolYearEventType
    literals []enum.Literal
}

var _schoolYearEventTypes = &schoolYearEventTypes{values: []*SchoolYearEventType{
    {name: "SchoolYearCreated", ordinal: 0},
    {name: "SchoolYearDeleted", ordinal: 1},
    {name: "SchoolYearUpdated", ordinal: 2}},
}

func SchoolYearEventTypes() *schoolYearEventTypes {
	return _schoolYearEventTypes
}

func (o *schoolYearEventTypes) Values() []*SchoolYearEventType {
	return o.values
}

func (o *schoolYearEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *schoolYearEventTypes) SchoolYearCreated() *SchoolYearEventType {
    return _schoolYearEventTypes.values[0]
}

func (o *schoolYearEventTypes) SchoolYearDeleted() *SchoolYearEventType {
    return _schoolYearEventTypes.values[1]
}

func (o *schoolYearEventTypes) SchoolYearUpdated() *SchoolYearEventType {
    return _schoolYearEventTypes.values[2]
}

func (o *schoolYearEventTypes) ParseSchoolYearEventType(name string) (ret *SchoolYearEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*SchoolYearEventType), ok
	}
	return
}



