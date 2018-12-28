package person

import (
    "ee/schkola/shared"
    "encoding/json"
    "fmt"
    "github.com/eugeis/gee/enum"
    "github.com/looplab/eventhorizon"
    "gopkg.in/mgo.v2/bson"
    "time"
)
const (
     ChurchCreatedEvent eventhorizon.EventType = "ChurchCreated"
     ChurchDeletedEvent eventhorizon.EventType = "ChurchDeleted"
     ChurchUpdatedEvent eventhorizon.EventType = "ChurchUpdated"
)


const (
     GraduationCreatedEvent eventhorizon.EventType = "GraduationCreated"
     GraduationDeletedEvent eventhorizon.EventType = "GraduationDeleted"
     GraduationUpdatedEvent eventhorizon.EventType = "GraduationUpdated"
)


const (
     ProfileCreatedEvent eventhorizon.EventType = "ProfileCreated"
     ProfileDeletedEvent eventhorizon.EventType = "ProfileDeleted"
     ProfileUpdatedEvent eventhorizon.EventType = "ProfileUpdated"
)




type ChurchCreated struct {
    Name string `json:"name" eh:"optional"`
    Address *Address `json:"address" eh:"optional"`
    Pastor *shared.PersonName `json:"pastor" eh:"optional"`
    Contact *Contact `json:"contact" eh:"optional"`
    Association string `json:"association" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type ChurchDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type ChurchUpdated struct {
    Name string `json:"name" eh:"optional"`
    Address *Address `json:"address" eh:"optional"`
    Pastor *shared.PersonName `json:"pastor" eh:"optional"`
    Contact *Contact `json:"contact" eh:"optional"`
    Association string `json:"association" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type GraduationCreated struct {
    Name string `json:"name" eh:"optional"`
    Level *GraduationLevel `json:"level" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type GraduationDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type GraduationUpdated struct {
    Name string `json:"name" eh:"optional"`
    Level *GraduationLevel `json:"level" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type ProfileCreated struct {
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


type ProfileDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type ProfileUpdated struct {
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




type ChurchEventType struct {
	name  string
	ordinal int
}

func (o *ChurchEventType) Name() string {
    return o.name
}

func (o *ChurchEventType) Ordinal() int {
    return o.ordinal
}

func (o ChurchEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *ChurchEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := ChurchEventTypes().ParseChurchEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ChurchEventType %q", lit.Name)
        }
	}
	return
}

func (o ChurchEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *ChurchEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := ChurchEventTypes().ParseChurchEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ChurchEventType %q", lit)
        }
    }
    return
}

func (o *ChurchEventType) IsChurchCreated() bool {
    return o == _churchEventTypes.ChurchCreated()
}

func (o *ChurchEventType) IsChurchDeleted() bool {
    return o == _churchEventTypes.ChurchDeleted()
}

func (o *ChurchEventType) IsChurchUpdated() bool {
    return o == _churchEventTypes.ChurchUpdated()
}

type churchEventTypes struct {
	values []*ChurchEventType
    literals []enum.Literal
}

var _churchEventTypes = &churchEventTypes{values: []*ChurchEventType{
    {name: "ChurchCreated", ordinal: 0},
    {name: "ChurchDeleted", ordinal: 1},
    {name: "ChurchUpdated", ordinal: 2}},
}

func ChurchEventTypes() *churchEventTypes {
	return _churchEventTypes
}

func (o *churchEventTypes) Values() []*ChurchEventType {
	return o.values
}

func (o *churchEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *churchEventTypes) ChurchCreated() *ChurchEventType {
    return _churchEventTypes.values[0]
}

func (o *churchEventTypes) ChurchDeleted() *ChurchEventType {
    return _churchEventTypes.values[1]
}

func (o *churchEventTypes) ChurchUpdated() *ChurchEventType {
    return _churchEventTypes.values[2]
}

func (o *churchEventTypes) ParseChurchEventType(name string) (ret *ChurchEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*ChurchEventType), ok
	}
	return
}


type GraduationEventType struct {
	name  string
	ordinal int
}

func (o *GraduationEventType) Name() string {
    return o.name
}

func (o *GraduationEventType) Ordinal() int {
    return o.ordinal
}

func (o GraduationEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *GraduationEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := GraduationEventTypes().ParseGraduationEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GraduationEventType %q", lit.Name)
        }
	}
	return
}

func (o GraduationEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *GraduationEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := GraduationEventTypes().ParseGraduationEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GraduationEventType %q", lit)
        }
    }
    return
}

func (o *GraduationEventType) IsGraduationCreated() bool {
    return o == _graduationEventTypes.GraduationCreated()
}

func (o *GraduationEventType) IsGraduationDeleted() bool {
    return o == _graduationEventTypes.GraduationDeleted()
}

func (o *GraduationEventType) IsGraduationUpdated() bool {
    return o == _graduationEventTypes.GraduationUpdated()
}

type graduationEventTypes struct {
	values []*GraduationEventType
    literals []enum.Literal
}

var _graduationEventTypes = &graduationEventTypes{values: []*GraduationEventType{
    {name: "GraduationCreated", ordinal: 0},
    {name: "GraduationDeleted", ordinal: 1},
    {name: "GraduationUpdated", ordinal: 2}},
}

func GraduationEventTypes() *graduationEventTypes {
	return _graduationEventTypes
}

func (o *graduationEventTypes) Values() []*GraduationEventType {
	return o.values
}

func (o *graduationEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *graduationEventTypes) GraduationCreated() *GraduationEventType {
    return _graduationEventTypes.values[0]
}

func (o *graduationEventTypes) GraduationDeleted() *GraduationEventType {
    return _graduationEventTypes.values[1]
}

func (o *graduationEventTypes) GraduationUpdated() *GraduationEventType {
    return _graduationEventTypes.values[2]
}

func (o *graduationEventTypes) ParseGraduationEventType(name string) (ret *GraduationEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*GraduationEventType), ok
	}
	return
}


type ProfileEventType struct {
	name  string
	ordinal int
}

func (o *ProfileEventType) Name() string {
    return o.name
}

func (o *ProfileEventType) Ordinal() int {
    return o.ordinal
}

func (o ProfileEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *ProfileEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := ProfileEventTypes().ParseProfileEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ProfileEventType %q", lit.Name)
        }
	}
	return
}

func (o ProfileEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *ProfileEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := ProfileEventTypes().ParseProfileEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ProfileEventType %q", lit)
        }
    }
    return
}

func (o *ProfileEventType) IsProfileCreated() bool {
    return o == _profileEventTypes.ProfileCreated()
}

func (o *ProfileEventType) IsProfileDeleted() bool {
    return o == _profileEventTypes.ProfileDeleted()
}

func (o *ProfileEventType) IsProfileUpdated() bool {
    return o == _profileEventTypes.ProfileUpdated()
}

type profileEventTypes struct {
	values []*ProfileEventType
    literals []enum.Literal
}

var _profileEventTypes = &profileEventTypes{values: []*ProfileEventType{
    {name: "ProfileCreated", ordinal: 0},
    {name: "ProfileDeleted", ordinal: 1},
    {name: "ProfileUpdated", ordinal: 2}},
}

func ProfileEventTypes() *profileEventTypes {
	return _profileEventTypes
}

func (o *profileEventTypes) Values() []*ProfileEventType {
	return o.values
}

func (o *profileEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *profileEventTypes) ProfileCreated() *ProfileEventType {
    return _profileEventTypes.values[0]
}

func (o *profileEventTypes) ProfileDeleted() *ProfileEventType {
    return _profileEventTypes.values[1]
}

func (o *profileEventTypes) ProfileUpdated() *ProfileEventType {
    return _profileEventTypes.values[2]
}

func (o *profileEventTypes) ParseProfileEventType(name string) (ret *ProfileEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*ProfileEventType), ok
	}
	return
}



