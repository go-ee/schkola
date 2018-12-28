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
     CreateBookCommand eventhorizon.CommandType = "CreateBook"
     DeleteBookCommand eventhorizon.CommandType = "DeleteBook"
     UpdateBookCommand eventhorizon.CommandType = "UpdateBook"
)




        
type CreateBook struct {
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
func (o *CreateBook) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CreateBook) AggregateType() eventhorizon.AggregateType  { return BookAggregateType }
func (o *CreateBook) CommandType() eventhorizon.CommandType      { return CreateBookCommand }



        
type DeleteBook struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *DeleteBook) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DeleteBook) AggregateType() eventhorizon.AggregateType  { return BookAggregateType }
func (o *DeleteBook) CommandType() eventhorizon.CommandType      { return DeleteBookCommand }



        
type UpdateBook struct {
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
func (o *UpdateBook) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *UpdateBook) AggregateType() eventhorizon.AggregateType  { return BookAggregateType }
func (o *UpdateBook) CommandType() eventhorizon.CommandType      { return UpdateBookCommand }





type BookCommandType struct {
	name  string
	ordinal int
}

func (o *BookCommandType) Name() string {
    return o.name
}

func (o *BookCommandType) Ordinal() int {
    return o.ordinal
}

func (o BookCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *BookCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := BookCommandTypes().ParseBookCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid BookCommandType %q", lit.Name)
        }
	}
	return
}

func (o BookCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *BookCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := BookCommandTypes().ParseBookCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid BookCommandType %q", lit)
        }
    }
    return
}

func (o *BookCommandType) IsCreateBook() bool {
    return o == _bookCommandTypes.CreateBook()
}

func (o *BookCommandType) IsDeleteBook() bool {
    return o == _bookCommandTypes.DeleteBook()
}

func (o *BookCommandType) IsUpdateBook() bool {
    return o == _bookCommandTypes.UpdateBook()
}

type bookCommandTypes struct {
	values []*BookCommandType
    literals []enum.Literal
}

var _bookCommandTypes = &bookCommandTypes{values: []*BookCommandType{
    {name: "CreateBook", ordinal: 0},
    {name: "DeleteBook", ordinal: 1},
    {name: "UpdateBook", ordinal: 2}},
}

func BookCommandTypes() *bookCommandTypes {
	return _bookCommandTypes
}

func (o *bookCommandTypes) Values() []*BookCommandType {
	return o.values
}

func (o *bookCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *bookCommandTypes) CreateBook() *BookCommandType {
    return _bookCommandTypes.values[0]
}

func (o *bookCommandTypes) DeleteBook() *BookCommandType {
    return _bookCommandTypes.values[1]
}

func (o *bookCommandTypes) UpdateBook() *BookCommandType {
    return _bookCommandTypes.values[2]
}

func (o *bookCommandTypes) ParseBookCommandType(name string) (ret *BookCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*BookCommandType), ok
	}
	return
}



