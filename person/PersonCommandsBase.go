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
     CreateChurchCommand eventhorizon.CommandType = "CreateChurch"
     DeleteChurchCommand eventhorizon.CommandType = "DeleteChurch"
     UpdateChurchCommand eventhorizon.CommandType = "UpdateChurch"
)


const (
     CreateGraduationCommand eventhorizon.CommandType = "CreateGraduation"
     DeleteGraduationCommand eventhorizon.CommandType = "DeleteGraduation"
     UpdateGraduationCommand eventhorizon.CommandType = "UpdateGraduation"
)


const (
     CreateProfileCommand eventhorizon.CommandType = "CreateProfile"
     DeleteProfileCommand eventhorizon.CommandType = "DeleteProfile"
     UpdateProfileCommand eventhorizon.CommandType = "UpdateProfile"
)




        
type CreateChurch struct {
    Name string `json:"name" eh:"optional"`
    Address *Address `json:"address" eh:"optional"`
    Pastor *shared.PersonName `json:"pastor" eh:"optional"`
    Contact *Contact `json:"contact" eh:"optional"`
    Association string `json:"association" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *CreateChurch) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CreateChurch) AggregateType() eventhorizon.AggregateType  { return ChurchAggregateType }
func (o *CreateChurch) CommandType() eventhorizon.CommandType      { return CreateChurchCommand }



        
type DeleteChurch struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *DeleteChurch) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DeleteChurch) AggregateType() eventhorizon.AggregateType  { return ChurchAggregateType }
func (o *DeleteChurch) CommandType() eventhorizon.CommandType      { return DeleteChurchCommand }



        
type UpdateChurch struct {
    Name string `json:"name" eh:"optional"`
    Address *Address `json:"address" eh:"optional"`
    Pastor *shared.PersonName `json:"pastor" eh:"optional"`
    Contact *Contact `json:"contact" eh:"optional"`
    Association string `json:"association" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *UpdateChurch) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *UpdateChurch) AggregateType() eventhorizon.AggregateType  { return ChurchAggregateType }
func (o *UpdateChurch) CommandType() eventhorizon.CommandType      { return UpdateChurchCommand }



        
type CreateGraduation struct {
    Name string `json:"name" eh:"optional"`
    Level *GraduationLevel `json:"level" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *CreateGraduation) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CreateGraduation) AggregateType() eventhorizon.AggregateType  { return GraduationAggregateType }
func (o *CreateGraduation) CommandType() eventhorizon.CommandType      { return CreateGraduationCommand }



        
type DeleteGraduation struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *DeleteGraduation) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DeleteGraduation) AggregateType() eventhorizon.AggregateType  { return GraduationAggregateType }
func (o *DeleteGraduation) CommandType() eventhorizon.CommandType      { return DeleteGraduationCommand }



        
type UpdateGraduation struct {
    Name string `json:"name" eh:"optional"`
    Level *GraduationLevel `json:"level" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *UpdateGraduation) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *UpdateGraduation) AggregateType() eventhorizon.AggregateType  { return GraduationAggregateType }
func (o *UpdateGraduation) CommandType() eventhorizon.CommandType      { return UpdateGraduationCommand }



        
type CreateProfile struct {
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
func (o *CreateProfile) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CreateProfile) AggregateType() eventhorizon.AggregateType  { return ProfileAggregateType }
func (o *CreateProfile) CommandType() eventhorizon.CommandType      { return CreateProfileCommand }



        
type DeleteProfile struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *DeleteProfile) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DeleteProfile) AggregateType() eventhorizon.AggregateType  { return ProfileAggregateType }
func (o *DeleteProfile) CommandType() eventhorizon.CommandType      { return DeleteProfileCommand }



        
type UpdateProfile struct {
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
func (o *UpdateProfile) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *UpdateProfile) AggregateType() eventhorizon.AggregateType  { return ProfileAggregateType }
func (o *UpdateProfile) CommandType() eventhorizon.CommandType      { return UpdateProfileCommand }





type ChurchCommandType struct {
	name  string
	ordinal int
}

func (o *ChurchCommandType) Name() string {
    return o.name
}

func (o *ChurchCommandType) Ordinal() int {
    return o.ordinal
}

func (o ChurchCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *ChurchCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := ChurchCommandTypes().ParseChurchCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ChurchCommandType %q", lit.Name)
        }
	}
	return
}

func (o ChurchCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *ChurchCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := ChurchCommandTypes().ParseChurchCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ChurchCommandType %q", lit)
        }
    }
    return
}

func (o *ChurchCommandType) IsCreateChurch() bool {
    return o == _churchCommandTypes.CreateChurch()
}

func (o *ChurchCommandType) IsDeleteChurch() bool {
    return o == _churchCommandTypes.DeleteChurch()
}

func (o *ChurchCommandType) IsUpdateChurch() bool {
    return o == _churchCommandTypes.UpdateChurch()
}

type churchCommandTypes struct {
	values []*ChurchCommandType
    literals []enum.Literal
}

var _churchCommandTypes = &churchCommandTypes{values: []*ChurchCommandType{
    {name: "CreateChurch", ordinal: 0},
    {name: "DeleteChurch", ordinal: 1},
    {name: "UpdateChurch", ordinal: 2}},
}

func ChurchCommandTypes() *churchCommandTypes {
	return _churchCommandTypes
}

func (o *churchCommandTypes) Values() []*ChurchCommandType {
	return o.values
}

func (o *churchCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *churchCommandTypes) CreateChurch() *ChurchCommandType {
    return _churchCommandTypes.values[0]
}

func (o *churchCommandTypes) DeleteChurch() *ChurchCommandType {
    return _churchCommandTypes.values[1]
}

func (o *churchCommandTypes) UpdateChurch() *ChurchCommandType {
    return _churchCommandTypes.values[2]
}

func (o *churchCommandTypes) ParseChurchCommandType(name string) (ret *ChurchCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*ChurchCommandType), ok
	}
	return
}


type GraduationCommandType struct {
	name  string
	ordinal int
}

func (o *GraduationCommandType) Name() string {
    return o.name
}

func (o *GraduationCommandType) Ordinal() int {
    return o.ordinal
}

func (o GraduationCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *GraduationCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := GraduationCommandTypes().ParseGraduationCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GraduationCommandType %q", lit.Name)
        }
	}
	return
}

func (o GraduationCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *GraduationCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := GraduationCommandTypes().ParseGraduationCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid GraduationCommandType %q", lit)
        }
    }
    return
}

func (o *GraduationCommandType) IsCreateGraduation() bool {
    return o == _graduationCommandTypes.CreateGraduation()
}

func (o *GraduationCommandType) IsDeleteGraduation() bool {
    return o == _graduationCommandTypes.DeleteGraduation()
}

func (o *GraduationCommandType) IsUpdateGraduation() bool {
    return o == _graduationCommandTypes.UpdateGraduation()
}

type graduationCommandTypes struct {
	values []*GraduationCommandType
    literals []enum.Literal
}

var _graduationCommandTypes = &graduationCommandTypes{values: []*GraduationCommandType{
    {name: "CreateGraduation", ordinal: 0},
    {name: "DeleteGraduation", ordinal: 1},
    {name: "UpdateGraduation", ordinal: 2}},
}

func GraduationCommandTypes() *graduationCommandTypes {
	return _graduationCommandTypes
}

func (o *graduationCommandTypes) Values() []*GraduationCommandType {
	return o.values
}

func (o *graduationCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *graduationCommandTypes) CreateGraduation() *GraduationCommandType {
    return _graduationCommandTypes.values[0]
}

func (o *graduationCommandTypes) DeleteGraduation() *GraduationCommandType {
    return _graduationCommandTypes.values[1]
}

func (o *graduationCommandTypes) UpdateGraduation() *GraduationCommandType {
    return _graduationCommandTypes.values[2]
}

func (o *graduationCommandTypes) ParseGraduationCommandType(name string) (ret *GraduationCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*GraduationCommandType), ok
	}
	return
}


type ProfileCommandType struct {
	name  string
	ordinal int
}

func (o *ProfileCommandType) Name() string {
    return o.name
}

func (o *ProfileCommandType) Ordinal() int {
    return o.ordinal
}

func (o ProfileCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *ProfileCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := ProfileCommandTypes().ParseProfileCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ProfileCommandType %q", lit.Name)
        }
	}
	return
}

func (o ProfileCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *ProfileCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := ProfileCommandTypes().ParseProfileCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ProfileCommandType %q", lit)
        }
    }
    return
}

func (o *ProfileCommandType) IsCreateProfile() bool {
    return o == _profileCommandTypes.CreateProfile()
}

func (o *ProfileCommandType) IsDeleteProfile() bool {
    return o == _profileCommandTypes.DeleteProfile()
}

func (o *ProfileCommandType) IsUpdateProfile() bool {
    return o == _profileCommandTypes.UpdateProfile()
}

type profileCommandTypes struct {
	values []*ProfileCommandType
    literals []enum.Literal
}

var _profileCommandTypes = &profileCommandTypes{values: []*ProfileCommandType{
    {name: "CreateProfile", ordinal: 0},
    {name: "DeleteProfile", ordinal: 1},
    {name: "UpdateProfile", ordinal: 2}},
}

func ProfileCommandTypes() *profileCommandTypes {
	return _profileCommandTypes
}

func (o *profileCommandTypes) Values() []*ProfileCommandType {
	return o.values
}

func (o *profileCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *profileCommandTypes) CreateProfile() *ProfileCommandType {
    return _profileCommandTypes.values[0]
}

func (o *profileCommandTypes) DeleteProfile() *ProfileCommandType {
    return _profileCommandTypes.values[1]
}

func (o *profileCommandTypes) UpdateProfile() *ProfileCommandType {
    return _profileCommandTypes.values[2]
}

func (o *profileCommandTypes) ParseProfileCommandType(name string) (ret *ProfileCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*ProfileCommandType), ok
	}
	return
}



