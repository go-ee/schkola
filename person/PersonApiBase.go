package person

import (
    "ee/schkola/shared"
    "encoding/json"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/eugeis/gee/enum"
    "github.com/looplab/eventhorizon"
    "gopkg.in/mgo.v2/bson"
    "time"
)
        
type Church struct {
    Name string `json:"name" eh:"optional"`
    Address *Address `json:"address" eh:"optional"`
    Pastor *shared.PersonName `json:"pastor" eh:"optional"`
    Contact *Contact `json:"contact" eh:"optional"`
    Association string `json:"association" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func NewChurch() (ret *Church) {
    ret = &Church{}
    return
}
func (o *Church) EntityID() eventhorizon.UUID { return o.Id }



        
type Graduation struct {
    Name string `json:"name" eh:"optional"`
    Level *GraduationLevel `json:"level" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func NewGraduation() (ret *Graduation) {
    ret = &Graduation{}
    return
}
func (o *Graduation) EntityID() eventhorizon.UUID { return o.Id }



        
type Profile struct {
    Gender *Gender `json:"gender" eh:"optional"`
    Name *shared.PersonName `json:"name" eh:"optional"`
    BirthName string `json:"birthName" eh:"optional"`
    Birthday *time.Time `json:"birthday" eh:"optional"`
    Address *Address `json:"address" eh:"optional"`
    Contact *Contact `json:"contact" eh:"optional"`
    PhotoData []byte `json:"photoData" eh:"optional"`
    Photo string `json:"photo" eh:"optional"`
    Family *Family `json:"family" eh:"optional"`
    Church *ChurchInfo `json:"church" eh:"optional"`
    Education *Education `json:"education" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func NewProfile() (ret *Profile) {
    ret = &Profile{}
    return
}

func (o *Profile) FindByEmail(email string) (ret *Profile, err error) {
    err = eh.QueryNotImplemented("findProfileByEmail")
    return
}

func (o *Profile) FindByPhone(phone string) (ret *Profile, err error) {
    err = eh.QueryNotImplemented("findProfileByPhone")
    return
}
func (o *Profile) EntityID() eventhorizon.UUID { return o.Id }







type Address struct {
    Street string `json:"street" eh:"optional"`
    Suite string `json:"suite" eh:"optional"`
    City string `json:"city" eh:"optional"`
    Code string `json:"code" eh:"optional"`
    Country string `json:"country" eh:"optional"`
}

func NewAddress() (ret *Address) {
    ret = &Address{}
    return
}


type ChurchInfo struct {
    Church string `json:"church" eh:"optional"`
    Member bool `json:"member" eh:"optional"`
    Services string `json:"services" eh:"optional"`
}

func NewChurchInfo() (ret *ChurchInfo) {
    ret = &ChurchInfo{}
    return
}


type Contact struct {
    Phone string `json:"phone" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Cellphone string `json:"cellphone" eh:"optional"`
}

func NewContact() (ret *Contact) {
    ret = &Contact{}
    return
}


type Education struct {
    Graduation *Graduation `json:"graduation" eh:"optional"`
    Other string `json:"other" eh:"optional"`
    Profession string `json:"profession" eh:"optional"`
}

func NewEducation() (ret *Education) {
    ret = &Education{}
    return
}


type Family struct {
    MaritalState *MaritalState `json:"maritalState" eh:"optional"`
    ChildrenCount int `json:"childrenCount" eh:"optional"`
    Partner *shared.PersonName `json:"partner" eh:"optional"`
}

func NewFamily() (ret *Family) {
    ret = &Family{}
    return
}




type Gender struct {
	name  string
	ordinal int
}

func (o *Gender) Name() string {
    return o.name
}

func (o *Gender) Ordinal() int {
    return o.ordinal
}

func (o Gender) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *Gender) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := Genders().ParseGender(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid Gender %q", lit.Name)
        }
	}
	return
}

func (o Gender) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *Gender) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := Genders().ParseGender(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid Gender %q", lit)
        }
    }
    return
}

func (o *Gender) IsUnknown() bool {
    return o == _genders.Unknown()
}

func (o *Gender) IsMale() bool {
    return o == _genders.Male()
}

func (o *Gender) IsFemale() bool {
    return o == _genders.Female()
}

type genders struct {
	values []*Gender
    literals []enum.Literal
}

var _genders = &genders{values: []*Gender{
    {name: "Unknown", ordinal: 0},
    {name: "Male", ordinal: 1},
    {name: "Female", ordinal: 2}},
}

func Genders() *genders {
	return _genders
}

func (o *genders) Values() []*Gender {
	return o.values
}

func (o *genders) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *genders) Unknown() *Gender {
    return _genders.values[0]
}

func (o *genders) Male() *Gender {
    return _genders.values[1]
}

func (o *genders) Female() *Gender {
    return _genders.values[2]
}

func (o *genders) ParseGender(name string) (ret *Gender, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*Gender), ok
	}
	return
}


type GraduationLevel struct {
	name  string
	ordinal int
}

func (o *GraduationLevel) Name() string {
    return o.name
}

func (o *GraduationLevel) Ordinal() int {
    return o.ordinal
}

func (o GraduationLevel) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *GraduationLevel) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := GraduationLevels().ParseGraduationLevel(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GraduationLevel %q", lit.Name)
        }
	}
	return
}

func (o GraduationLevel) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *GraduationLevel) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := GraduationLevels().ParseGraduationLevel(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GraduationLevel %q", lit)
        }
    }
    return
}

func (o *GraduationLevel) IsUnknown() bool {
    return o == _graduationLevels.Unknown()
}

func (o *GraduationLevel) IsMiddleSchool() bool {
    return o == _graduationLevels.MiddleSchool()
}

func (o *GraduationLevel) IsSecondarySchool() bool {
    return o == _graduationLevels.SecondarySchool()
}

func (o *GraduationLevel) IsHighSchool() bool {
    return o == _graduationLevels.HighSchool()
}

func (o *GraduationLevel) IsTechnicalCollege() bool {
    return o == _graduationLevels.TechnicalCollege()
}

func (o *GraduationLevel) IsCollege() bool {
    return o == _graduationLevels.College()
}

type graduationLevels struct {
	values []*GraduationLevel
    literals []enum.Literal
}

var _graduationLevels = &graduationLevels{values: []*GraduationLevel{
    {name: "Unknown", ordinal: 0},
    {name: "MiddleSchool", ordinal: 1},
    {name: "SecondarySchool", ordinal: 2},
    {name: "HighSchool", ordinal: 3},
    {name: "TechnicalCollege", ordinal: 4},
    {name: "College", ordinal: 5}},
}

func GraduationLevels() *graduationLevels {
	return _graduationLevels
}

func (o *graduationLevels) Values() []*GraduationLevel {
	return o.values
}

func (o *graduationLevels) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *graduationLevels) Unknown() *GraduationLevel {
    return _graduationLevels.values[0]
}

func (o *graduationLevels) MiddleSchool() *GraduationLevel {
    return _graduationLevels.values[1]
}

func (o *graduationLevels) SecondarySchool() *GraduationLevel {
    return _graduationLevels.values[2]
}

func (o *graduationLevels) HighSchool() *GraduationLevel {
    return _graduationLevels.values[3]
}

func (o *graduationLevels) TechnicalCollege() *GraduationLevel {
    return _graduationLevels.values[4]
}

func (o *graduationLevels) College() *GraduationLevel {
    return _graduationLevels.values[5]
}

func (o *graduationLevels) ParseGraduationLevel(name string) (ret *GraduationLevel, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*GraduationLevel), ok
	}
	return
}


type MaritalState struct {
	name  string
	ordinal int
}

func (o *MaritalState) Name() string {
    return o.name
}

func (o *MaritalState) Ordinal() int {
    return o.ordinal
}

func (o MaritalState) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *MaritalState) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := MaritalStates().ParseMaritalState(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid MaritalState %q", lit.Name)
        }
	}
	return
}

func (o MaritalState) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *MaritalState) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := MaritalStates().ParseMaritalState(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid MaritalState %q", lit)
        }
    }
    return
}

func (o *MaritalState) IsUnknown() bool {
    return o == _maritalStates.Unknown()
}

func (o *MaritalState) IsSingle() bool {
    return o == _maritalStates.Single()
}

func (o *MaritalState) IsMarried() bool {
    return o == _maritalStates.Married()
}

func (o *MaritalState) IsSeparated() bool {
    return o == _maritalStates.Separated()
}

func (o *MaritalState) IsDivorced() bool {
    return o == _maritalStates.Divorced()
}

func (o *MaritalState) IsWidowed() bool {
    return o == _maritalStates.Widowed()
}

type maritalStates struct {
	values []*MaritalState
    literals []enum.Literal
}

var _maritalStates = &maritalStates{values: []*MaritalState{
    {name: "Unknown", ordinal: 0},
    {name: "Single", ordinal: 1},
    {name: "Married", ordinal: 2},
    {name: "Separated", ordinal: 3},
    {name: "Divorced", ordinal: 4},
    {name: "Widowed", ordinal: 5}},
}

func MaritalStates() *maritalStates {
	return _maritalStates
}

func (o *maritalStates) Values() []*MaritalState {
	return o.values
}

func (o *maritalStates) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *maritalStates) Unknown() *MaritalState {
    return _maritalStates.values[0]
}

func (o *maritalStates) Single() *MaritalState {
    return _maritalStates.values[1]
}

func (o *maritalStates) Married() *MaritalState {
    return _maritalStates.values[2]
}

func (o *maritalStates) Separated() *MaritalState {
    return _maritalStates.values[3]
}

func (o *maritalStates) Divorced() *MaritalState {
    return _maritalStates.values[4]
}

func (o *maritalStates) Widowed() *MaritalState {
    return _maritalStates.values[5]
}

func (o *maritalStates) ParseMaritalState(name string) (ret *MaritalState, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*MaritalState), ok
	}
	return
}



