package library

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
     BookCreatedEvent eventhorizon.EventType = "BookCreated"
     BookDeletedEvent eventhorizon.EventType = "BookDeleted"
     BookUpdatedEvent eventhorizon.EventType = "BookUpdated"
)




type BookCreated struct {
    Title string `json:"title" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Language string `json:"language" eh:"optional"`
    ReleaseDate *time.Time `json:"releaseDate" eh:"optional"`
    Edition string `json:"edition" eh:"optional"`
    Category string `json:"category" eh:"optional"`
    Author *shared.PersonName `json:"author" eh:"optional"`
    Location *Location `json:"location" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type BookDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type BookUpdated struct {
    Title string `json:"title" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Language string `json:"language" eh:"optional"`
    ReleaseDate *time.Time `json:"releaseDate" eh:"optional"`
    Edition string `json:"edition" eh:"optional"`
    Category string `json:"category" eh:"optional"`
    Author *shared.PersonName `json:"author" eh:"optional"`
    Location *Location `json:"location" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}




type BookEventType struct {
	name  string
	ordinal int
}

func (o *BookEventType) Name() string {
    return o.name
}

func (o *BookEventType) Ordinal() int {
    return o.ordinal
}

func (o BookEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *BookEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := BookEventTypes().ParseBookEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid BookEventType %q", lit.Name)
        }
	}
	return
}

func (o BookEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *BookEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := BookEventTypes().ParseBookEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid BookEventType %q", lit)
        }
    }
    return
}

func (o *BookEventType) IsBookCreated() bool {
    return o == _bookEventTypes.BookCreated()
}

func (o *BookEventType) IsBookDeleted() bool {
    return o == _bookEventTypes.BookDeleted()
}

func (o *BookEventType) IsBookUpdated() bool {
    return o == _bookEventTypes.BookUpdated()
}

type bookEventTypes struct {
	values []*BookEventType
    literals []enum.Literal
}

var _bookEventTypes = &bookEventTypes{values: []*BookEventType{
    {name: "BookCreated", ordinal: 0},
    {name: "BookDeleted", ordinal: 1},
    {name: "BookUpdated", ordinal: 2}},
}

func BookEventTypes() *bookEventTypes {
	return _bookEventTypes
}

func (o *bookEventTypes) Values() []*BookEventType {
	return o.values
}

func (o *bookEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *bookEventTypes) BookCreated() *BookEventType {
    return _bookEventTypes.values[0]
}

func (o *bookEventTypes) BookDeleted() *BookEventType {
    return _bookEventTypes.values[1]
}

func (o *bookEventTypes) BookUpdated() *BookEventType {
    return _bookEventTypes.values[2]
}

func (o *bookEventTypes) ParseBookEventType(name string) (ret *BookEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*BookEventType), ok
	}
	return
}



