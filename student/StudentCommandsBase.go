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
     RegisterAttendanceCommand eventhorizon.CommandType = "RegisterAttendance"
     CreateAttendanceCommand eventhorizon.CommandType = "CreateAttendance"
     DeleteAttendanceCommand eventhorizon.CommandType = "DeleteAttendance"
     ConfirmAttendanceCommand eventhorizon.CommandType = "ConfirmAttendance"
     CancelAttendanceCommand eventhorizon.CommandType = "CancelAttendance"
     UpdateAttendanceCommand eventhorizon.CommandType = "UpdateAttendance"
)


const (
     CreateCourseCommand eventhorizon.CommandType = "CreateCourse"
     DeleteCourseCommand eventhorizon.CommandType = "DeleteCourse"
     UpdateCourseCommand eventhorizon.CommandType = "UpdateCourse"
)


const (
     CreateGradeCommand eventhorizon.CommandType = "CreateGrade"
     DeleteGradeCommand eventhorizon.CommandType = "DeleteGrade"
     UpdateGradeCommand eventhorizon.CommandType = "UpdateGrade"
)


const (
     CreateGroupCommand eventhorizon.CommandType = "CreateGroup"
     DeleteGroupCommand eventhorizon.CommandType = "DeleteGroup"
     UpdateGroupCommand eventhorizon.CommandType = "UpdateGroup"
)


const (
     CreateSchoolApplicationCommand eventhorizon.CommandType = "CreateSchoolApplication"
     DeleteSchoolApplicationCommand eventhorizon.CommandType = "DeleteSchoolApplication"
     UpdateSchoolApplicationCommand eventhorizon.CommandType = "UpdateSchoolApplication"
)


const (
     CreateSchoolYearCommand eventhorizon.CommandType = "CreateSchoolYear"
     DeleteSchoolYearCommand eventhorizon.CommandType = "DeleteSchoolYear"
     UpdateSchoolYearCommand eventhorizon.CommandType = "UpdateSchoolYear"
)




        
type RegisterAttendance struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *RegisterAttendance) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *RegisterAttendance) AggregateType() eventhorizon.AggregateType  { return AttendanceAggregateType }
func (o *RegisterAttendance) CommandType() eventhorizon.CommandType      { return RegisterAttendanceCommand }



        
type CreateAttendance struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Hours int `json:"hours" eh:"optional"`
    State *AttendanceState `json:"state" eh:"optional"`
    Token string `json:"token" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *CreateAttendance) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CreateAttendance) AggregateType() eventhorizon.AggregateType  { return AttendanceAggregateType }
func (o *CreateAttendance) CommandType() eventhorizon.CommandType      { return CreateAttendanceCommand }



        
type DeleteAttendance struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *DeleteAttendance) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DeleteAttendance) AggregateType() eventhorizon.AggregateType  { return AttendanceAggregateType }
func (o *DeleteAttendance) CommandType() eventhorizon.CommandType      { return DeleteAttendanceCommand }



        
type ConfirmAttendance struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *ConfirmAttendance) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *ConfirmAttendance) AggregateType() eventhorizon.AggregateType  { return AttendanceAggregateType }
func (o *ConfirmAttendance) CommandType() eventhorizon.CommandType      { return ConfirmAttendanceCommand }



        
type CancelAttendance struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *CancelAttendance) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CancelAttendance) AggregateType() eventhorizon.AggregateType  { return AttendanceAggregateType }
func (o *CancelAttendance) CommandType() eventhorizon.CommandType      { return CancelAttendanceCommand }



        
type UpdateAttendance struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Hours int `json:"hours" eh:"optional"`
    State *AttendanceState `json:"state" eh:"optional"`
    Token string `json:"token" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *UpdateAttendance) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *UpdateAttendance) AggregateType() eventhorizon.AggregateType  { return AttendanceAggregateType }
func (o *UpdateAttendance) CommandType() eventhorizon.CommandType      { return UpdateAttendanceCommand }



        
type CreateCourse struct {
    Name string `json:"name" eh:"optional"`
    Begin *time.Time `json:"begin" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Teacher *shared.PersonName `json:"teacher" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Fee float64 `json:"fee" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *CreateCourse) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CreateCourse) AggregateType() eventhorizon.AggregateType  { return CourseAggregateType }
func (o *CreateCourse) CommandType() eventhorizon.CommandType      { return CreateCourseCommand }



        
type DeleteCourse struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *DeleteCourse) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DeleteCourse) AggregateType() eventhorizon.AggregateType  { return CourseAggregateType }
func (o *DeleteCourse) CommandType() eventhorizon.CommandType      { return DeleteCourseCommand }



        
type UpdateCourse struct {
    Name string `json:"name" eh:"optional"`
    Begin *time.Time `json:"begin" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Teacher *shared.PersonName `json:"teacher" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Fee float64 `json:"fee" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *UpdateCourse) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *UpdateCourse) AggregateType() eventhorizon.AggregateType  { return CourseAggregateType }
func (o *UpdateCourse) CommandType() eventhorizon.CommandType      { return UpdateCourseCommand }



        
type CreateGrade struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Grade float64 `json:"grade" eh:"optional"`
    Comment string `json:"comment" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *CreateGrade) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CreateGrade) AggregateType() eventhorizon.AggregateType  { return GradeAggregateType }
func (o *CreateGrade) CommandType() eventhorizon.CommandType      { return CreateGradeCommand }



        
type DeleteGrade struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *DeleteGrade) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DeleteGrade) AggregateType() eventhorizon.AggregateType  { return GradeAggregateType }
func (o *DeleteGrade) CommandType() eventhorizon.CommandType      { return DeleteGradeCommand }



        
type UpdateGrade struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Course *Course `json:"course" eh:"optional"`
    Grade float64 `json:"grade" eh:"optional"`
    Comment string `json:"comment" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *UpdateGrade) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *UpdateGrade) AggregateType() eventhorizon.AggregateType  { return GradeAggregateType }
func (o *UpdateGrade) CommandType() eventhorizon.CommandType      { return UpdateGradeCommand }



        
type CreateGroup struct {
    Name string `json:"name" eh:"optional"`
    Category *GroupCategory `json:"category" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Representative *person.Profile `json:"representative" eh:"optional"`
    Students []*person.Profile `json:"students" eh:"optional"`
    Courses []*Course `json:"courses" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *CreateGroup) AddToStudents(item *person.Profile) *person.Profile {
    o.Students = append(o.Students, item)
    return item
}

func (o *CreateGroup) AddToCourses(item *Course) *Course {
    o.Courses = append(o.Courses, item)
    return item
}
func (o *CreateGroup) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CreateGroup) AggregateType() eventhorizon.AggregateType  { return GroupAggregateType }
func (o *CreateGroup) CommandType() eventhorizon.CommandType      { return CreateGroupCommand }



        
type DeleteGroup struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *DeleteGroup) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DeleteGroup) AggregateType() eventhorizon.AggregateType  { return GroupAggregateType }
func (o *DeleteGroup) CommandType() eventhorizon.CommandType      { return DeleteGroupCommand }



        
type UpdateGroup struct {
    Name string `json:"name" eh:"optional"`
    Category *GroupCategory `json:"category" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Representative *person.Profile `json:"representative" eh:"optional"`
    Students []*person.Profile `json:"students" eh:"optional"`
    Courses []*Course `json:"courses" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *UpdateGroup) AddToStudents(item *person.Profile) *person.Profile {
    o.Students = append(o.Students, item)
    return item
}

func (o *UpdateGroup) AddToCourses(item *Course) *Course {
    o.Courses = append(o.Courses, item)
    return item
}
func (o *UpdateGroup) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *UpdateGroup) AggregateType() eventhorizon.AggregateType  { return GroupAggregateType }
func (o *UpdateGroup) CommandType() eventhorizon.CommandType      { return UpdateGroupCommand }



        
type CreateSchoolApplication struct {
    Profile *person.Profile `json:"profile" eh:"optional"`
    ChurchContactPerson *shared.PersonName `json:"churchContactPerson" eh:"optional"`
    ChurchContact *person.Contact `json:"churchContact" eh:"optional"`
    ChurchCommitment bool `json:"churchCommitment" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Group string `json:"group" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *CreateSchoolApplication) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CreateSchoolApplication) AggregateType() eventhorizon.AggregateType  { return SchoolApplicationAggregateType }
func (o *CreateSchoolApplication) CommandType() eventhorizon.CommandType      { return CreateSchoolApplicationCommand }



        
type DeleteSchoolApplication struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *DeleteSchoolApplication) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DeleteSchoolApplication) AggregateType() eventhorizon.AggregateType  { return SchoolApplicationAggregateType }
func (o *DeleteSchoolApplication) CommandType() eventhorizon.CommandType      { return DeleteSchoolApplicationCommand }



        
type UpdateSchoolApplication struct {
    Profile *person.Profile `json:"profile" eh:"optional"`
    ChurchContactPerson *shared.PersonName `json:"churchContactPerson" eh:"optional"`
    ChurchContact *person.Contact `json:"churchContact" eh:"optional"`
    ChurchCommitment bool `json:"churchCommitment" eh:"optional"`
    SchoolYear *SchoolYear `json:"schoolYear" eh:"optional"`
    Group string `json:"group" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *UpdateSchoolApplication) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *UpdateSchoolApplication) AggregateType() eventhorizon.AggregateType  { return SchoolApplicationAggregateType }
func (o *UpdateSchoolApplication) CommandType() eventhorizon.CommandType      { return UpdateSchoolApplicationCommand }



        
type CreateSchoolYear struct {
    Name string `json:"name" eh:"optional"`
    Start *time.Time `json:"start" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Dates []*time.Time `json:"dates" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *CreateSchoolYear) AddToDates(item *time.Time) *time.Time {
    o.Dates = append(o.Dates, item)
    return item
}
func (o *CreateSchoolYear) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CreateSchoolYear) AggregateType() eventhorizon.AggregateType  { return SchoolYearAggregateType }
func (o *CreateSchoolYear) CommandType() eventhorizon.CommandType      { return CreateSchoolYearCommand }



        
type DeleteSchoolYear struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *DeleteSchoolYear) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DeleteSchoolYear) AggregateType() eventhorizon.AggregateType  { return SchoolYearAggregateType }
func (o *DeleteSchoolYear) CommandType() eventhorizon.CommandType      { return DeleteSchoolYearCommand }



        
type UpdateSchoolYear struct {
    Name string `json:"name" eh:"optional"`
    Start *time.Time `json:"start" eh:"optional"`
    End *time.Time `json:"end" eh:"optional"`
    Dates []*time.Time `json:"dates" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *UpdateSchoolYear) AddToDates(item *time.Time) *time.Time {
    o.Dates = append(o.Dates, item)
    return item
}
func (o *UpdateSchoolYear) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *UpdateSchoolYear) AggregateType() eventhorizon.AggregateType  { return SchoolYearAggregateType }
func (o *UpdateSchoolYear) CommandType() eventhorizon.CommandType      { return UpdateSchoolYearCommand }





type AttendanceCommandType struct {
	name  string
	ordinal int
}

func (o *AttendanceCommandType) Name() string {
    return o.name
}

func (o *AttendanceCommandType) Ordinal() int {
    return o.ordinal
}

func (o AttendanceCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *AttendanceCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := AttendanceCommandTypes().ParseAttendanceCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid AttendanceCommandType %q", lit.Name)
        }
	}
	return
}

func (o AttendanceCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *AttendanceCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := AttendanceCommandTypes().ParseAttendanceCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid AttendanceCommandType %q", lit)
        }
    }
    return
}

func (o *AttendanceCommandType) IsRegisterAttendance() bool {
    return o == _attendanceCommandTypes.RegisterAttendance()
}

func (o *AttendanceCommandType) IsCreateAttendance() bool {
    return o == _attendanceCommandTypes.CreateAttendance()
}

func (o *AttendanceCommandType) IsDeleteAttendance() bool {
    return o == _attendanceCommandTypes.DeleteAttendance()
}

func (o *AttendanceCommandType) IsConfirmAttendance() bool {
    return o == _attendanceCommandTypes.ConfirmAttendance()
}

func (o *AttendanceCommandType) IsCancelAttendance() bool {
    return o == _attendanceCommandTypes.CancelAttendance()
}

func (o *AttendanceCommandType) IsUpdateAttendance() bool {
    return o == _attendanceCommandTypes.UpdateAttendance()
}

type attendanceCommandTypes struct {
	values []*AttendanceCommandType
    literals []enum.Literal
}

var _attendanceCommandTypes = &attendanceCommandTypes{values: []*AttendanceCommandType{
    {name: "RegisterAttendance", ordinal: 0},
    {name: "CreateAttendance", ordinal: 1},
    {name: "DeleteAttendance", ordinal: 2},
    {name: "ConfirmAttendance", ordinal: 3},
    {name: "CancelAttendance", ordinal: 4},
    {name: "UpdateAttendance", ordinal: 5}},
}

func AttendanceCommandTypes() *attendanceCommandTypes {
	return _attendanceCommandTypes
}

func (o *attendanceCommandTypes) Values() []*AttendanceCommandType {
	return o.values
}

func (o *attendanceCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *attendanceCommandTypes) RegisterAttendance() *AttendanceCommandType {
    return _attendanceCommandTypes.values[0]
}

func (o *attendanceCommandTypes) CreateAttendance() *AttendanceCommandType {
    return _attendanceCommandTypes.values[1]
}

func (o *attendanceCommandTypes) DeleteAttendance() *AttendanceCommandType {
    return _attendanceCommandTypes.values[2]
}

func (o *attendanceCommandTypes) ConfirmAttendance() *AttendanceCommandType {
    return _attendanceCommandTypes.values[3]
}

func (o *attendanceCommandTypes) CancelAttendance() *AttendanceCommandType {
    return _attendanceCommandTypes.values[4]
}

func (o *attendanceCommandTypes) UpdateAttendance() *AttendanceCommandType {
    return _attendanceCommandTypes.values[5]
}

func (o *attendanceCommandTypes) ParseAttendanceCommandType(name string) (ret *AttendanceCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*AttendanceCommandType), ok
	}
	return
}


type CourseCommandType struct {
	name  string
	ordinal int
}

func (o *CourseCommandType) Name() string {
    return o.name
}

func (o *CourseCommandType) Ordinal() int {
    return o.ordinal
}

func (o CourseCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *CourseCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := CourseCommandTypes().ParseCourseCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid CourseCommandType %q", lit.Name)
        }
	}
	return
}

func (o CourseCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *CourseCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := CourseCommandTypes().ParseCourseCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid CourseCommandType %q", lit)
        }
    }
    return
}

func (o *CourseCommandType) IsCreateCourse() bool {
    return o == _courseCommandTypes.CreateCourse()
}

func (o *CourseCommandType) IsDeleteCourse() bool {
    return o == _courseCommandTypes.DeleteCourse()
}

func (o *CourseCommandType) IsUpdateCourse() bool {
    return o == _courseCommandTypes.UpdateCourse()
}

type courseCommandTypes struct {
	values []*CourseCommandType
    literals []enum.Literal
}

var _courseCommandTypes = &courseCommandTypes{values: []*CourseCommandType{
    {name: "CreateCourse", ordinal: 0},
    {name: "DeleteCourse", ordinal: 1},
    {name: "UpdateCourse", ordinal: 2}},
}

func CourseCommandTypes() *courseCommandTypes {
	return _courseCommandTypes
}

func (o *courseCommandTypes) Values() []*CourseCommandType {
	return o.values
}

func (o *courseCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *courseCommandTypes) CreateCourse() *CourseCommandType {
    return _courseCommandTypes.values[0]
}

func (o *courseCommandTypes) DeleteCourse() *CourseCommandType {
    return _courseCommandTypes.values[1]
}

func (o *courseCommandTypes) UpdateCourse() *CourseCommandType {
    return _courseCommandTypes.values[2]
}

func (o *courseCommandTypes) ParseCourseCommandType(name string) (ret *CourseCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*CourseCommandType), ok
	}
	return
}


type GradeCommandType struct {
	name  string
	ordinal int
}

func (o *GradeCommandType) Name() string {
    return o.name
}

func (o *GradeCommandType) Ordinal() int {
    return o.ordinal
}

func (o GradeCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *GradeCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := GradeCommandTypes().ParseGradeCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GradeCommandType %q", lit.Name)
        }
	}
	return
}

func (o GradeCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *GradeCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := GradeCommandTypes().ParseGradeCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GradeCommandType %q", lit)
        }
    }
    return
}

func (o *GradeCommandType) IsCreateGrade() bool {
    return o == _gradeCommandTypes.CreateGrade()
}

func (o *GradeCommandType) IsDeleteGrade() bool {
    return o == _gradeCommandTypes.DeleteGrade()
}

func (o *GradeCommandType) IsUpdateGrade() bool {
    return o == _gradeCommandTypes.UpdateGrade()
}

type gradeCommandTypes struct {
	values []*GradeCommandType
    literals []enum.Literal
}

var _gradeCommandTypes = &gradeCommandTypes{values: []*GradeCommandType{
    {name: "CreateGrade", ordinal: 0},
    {name: "DeleteGrade", ordinal: 1},
    {name: "UpdateGrade", ordinal: 2}},
}

func GradeCommandTypes() *gradeCommandTypes {
	return _gradeCommandTypes
}

func (o *gradeCommandTypes) Values() []*GradeCommandType {
	return o.values
}

func (o *gradeCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *gradeCommandTypes) CreateGrade() *GradeCommandType {
    return _gradeCommandTypes.values[0]
}

func (o *gradeCommandTypes) DeleteGrade() *GradeCommandType {
    return _gradeCommandTypes.values[1]
}

func (o *gradeCommandTypes) UpdateGrade() *GradeCommandType {
    return _gradeCommandTypes.values[2]
}

func (o *gradeCommandTypes) ParseGradeCommandType(name string) (ret *GradeCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*GradeCommandType), ok
	}
	return
}


type GroupCommandType struct {
	name  string
	ordinal int
}

func (o *GroupCommandType) Name() string {
    return o.name
}

func (o *GroupCommandType) Ordinal() int {
    return o.ordinal
}

func (o GroupCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *GroupCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := GroupCommandTypes().ParseGroupCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GroupCommandType %q", lit.Name)
        }
	}
	return
}

func (o GroupCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *GroupCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := GroupCommandTypes().ParseGroupCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GroupCommandType %q", lit)
        }
    }
    return
}

func (o *GroupCommandType) IsCreateGroup() bool {
    return o == _groupCommandTypes.CreateGroup()
}

func (o *GroupCommandType) IsDeleteGroup() bool {
    return o == _groupCommandTypes.DeleteGroup()
}

func (o *GroupCommandType) IsUpdateGroup() bool {
    return o == _groupCommandTypes.UpdateGroup()
}

type groupCommandTypes struct {
	values []*GroupCommandType
    literals []enum.Literal
}

var _groupCommandTypes = &groupCommandTypes{values: []*GroupCommandType{
    {name: "CreateGroup", ordinal: 0},
    {name: "DeleteGroup", ordinal: 1},
    {name: "UpdateGroup", ordinal: 2}},
}

func GroupCommandTypes() *groupCommandTypes {
	return _groupCommandTypes
}

func (o *groupCommandTypes) Values() []*GroupCommandType {
	return o.values
}

func (o *groupCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *groupCommandTypes) CreateGroup() *GroupCommandType {
    return _groupCommandTypes.values[0]
}

func (o *groupCommandTypes) DeleteGroup() *GroupCommandType {
    return _groupCommandTypes.values[1]
}

func (o *groupCommandTypes) UpdateGroup() *GroupCommandType {
    return _groupCommandTypes.values[2]
}

func (o *groupCommandTypes) ParseGroupCommandType(name string) (ret *GroupCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*GroupCommandType), ok
	}
	return
}


type SchoolApplicationCommandType struct {
	name  string
	ordinal int
}

func (o *SchoolApplicationCommandType) Name() string {
    return o.name
}

func (o *SchoolApplicationCommandType) Ordinal() int {
    return o.ordinal
}

func (o SchoolApplicationCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *SchoolApplicationCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := SchoolApplicationCommandTypes().ParseSchoolApplicationCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid SchoolApplicationCommandType %q", lit.Name)
        }
	}
	return
}

func (o SchoolApplicationCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *SchoolApplicationCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := SchoolApplicationCommandTypes().ParseSchoolApplicationCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid SchoolApplicationCommandType %q", lit)
        }
    }
    return
}

func (o *SchoolApplicationCommandType) IsCreateSchoolApplication() bool {
    return o == _schoolApplicationCommandTypes.CreateSchoolApplication()
}

func (o *SchoolApplicationCommandType) IsDeleteSchoolApplication() bool {
    return o == _schoolApplicationCommandTypes.DeleteSchoolApplication()
}

func (o *SchoolApplicationCommandType) IsUpdateSchoolApplication() bool {
    return o == _schoolApplicationCommandTypes.UpdateSchoolApplication()
}

type schoolApplicationCommandTypes struct {
	values []*SchoolApplicationCommandType
    literals []enum.Literal
}

var _schoolApplicationCommandTypes = &schoolApplicationCommandTypes{values: []*SchoolApplicationCommandType{
    {name: "CreateSchoolApplication", ordinal: 0},
    {name: "DeleteSchoolApplication", ordinal: 1},
    {name: "UpdateSchoolApplication", ordinal: 2}},
}

func SchoolApplicationCommandTypes() *schoolApplicationCommandTypes {
	return _schoolApplicationCommandTypes
}

func (o *schoolApplicationCommandTypes) Values() []*SchoolApplicationCommandType {
	return o.values
}

func (o *schoolApplicationCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *schoolApplicationCommandTypes) CreateSchoolApplication() *SchoolApplicationCommandType {
    return _schoolApplicationCommandTypes.values[0]
}

func (o *schoolApplicationCommandTypes) DeleteSchoolApplication() *SchoolApplicationCommandType {
    return _schoolApplicationCommandTypes.values[1]
}

func (o *schoolApplicationCommandTypes) UpdateSchoolApplication() *SchoolApplicationCommandType {
    return _schoolApplicationCommandTypes.values[2]
}

func (o *schoolApplicationCommandTypes) ParseSchoolApplicationCommandType(name string) (ret *SchoolApplicationCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*SchoolApplicationCommandType), ok
	}
	return
}


type SchoolYearCommandType struct {
	name  string
	ordinal int
}

func (o *SchoolYearCommandType) Name() string {
    return o.name
}

func (o *SchoolYearCommandType) Ordinal() int {
    return o.ordinal
}

func (o SchoolYearCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *SchoolYearCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := SchoolYearCommandTypes().ParseSchoolYearCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid SchoolYearCommandType %q", lit.Name)
        }
	}
	return
}

func (o SchoolYearCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *SchoolYearCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := SchoolYearCommandTypes().ParseSchoolYearCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid SchoolYearCommandType %q", lit)
        }
    }
    return
}

func (o *SchoolYearCommandType) IsCreateSchoolYear() bool {
    return o == _schoolYearCommandTypes.CreateSchoolYear()
}

func (o *SchoolYearCommandType) IsDeleteSchoolYear() bool {
    return o == _schoolYearCommandTypes.DeleteSchoolYear()
}

func (o *SchoolYearCommandType) IsUpdateSchoolYear() bool {
    return o == _schoolYearCommandTypes.UpdateSchoolYear()
}

type schoolYearCommandTypes struct {
	values []*SchoolYearCommandType
    literals []enum.Literal
}

var _schoolYearCommandTypes = &schoolYearCommandTypes{values: []*SchoolYearCommandType{
    {name: "CreateSchoolYear", ordinal: 0},
    {name: "DeleteSchoolYear", ordinal: 1},
    {name: "UpdateSchoolYear", ordinal: 2}},
}

func SchoolYearCommandTypes() *schoolYearCommandTypes {
	return _schoolYearCommandTypes
}

func (o *schoolYearCommandTypes) Values() []*SchoolYearCommandType {
	return o.values
}

func (o *schoolYearCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *schoolYearCommandTypes) CreateSchoolYear() *SchoolYearCommandType {
    return _schoolYearCommandTypes.values[0]
}

func (o *schoolYearCommandTypes) DeleteSchoolYear() *SchoolYearCommandType {
    return _schoolYearCommandTypes.values[1]
}

func (o *schoolYearCommandTypes) UpdateSchoolYear() *SchoolYearCommandType {
    return _schoolYearCommandTypes.values[2]
}

func (o *schoolYearCommandTypes) ParseSchoolYearCommandType(name string) (ret *SchoolYearCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*SchoolYearCommandType), ok
	}
	return
}



