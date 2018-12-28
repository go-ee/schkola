package auth

import (
    "ee/schkola/person"
    "ee/schkola/shared"
    "encoding/json"
    "fmt"
    "github.com/eugeis/gee/enum"
    "github.com/looplab/eventhorizon"
    "gopkg.in/mgo.v2/bson"
)
const (
     AccountCreatedEvent eventhorizon.EventType = "AccountCreated"
     AccountUpdatedEvent eventhorizon.EventType = "AccountUpdated"
     AccountDeletedEvent eventhorizon.EventType = "AccountDeleted"
     AccountEnabledEvent eventhorizon.EventType = "AccountEnabled"
     AccountDisabledEvent eventhorizon.EventType = "AccountDisabled"
     SendEnabledAccountConfirmationedEvent eventhorizon.EventType = "SendEnabledAccountConfirmationed"
     SendDisabledAccountConfirmationedEvent eventhorizon.EventType = "SendDisabledAccountConfirmationed"
     AccountLoggedEvent eventhorizon.EventType = "AccountLogged"
     SendCreatedAccountConfirmationedEvent eventhorizon.EventType = "SendCreatedAccountConfirmationed"
)




type AccountCreated struct {
    Name *shared.PersonName `json:"name" eh:"optional"`
    Username string `json:"username" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Roles []string `json:"roles" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *AccountCreated) AddToRoles(item string) string {
    o.Roles = append(o.Roles, item)
    return item
}


type AccountUpdated struct {
    Name *shared.PersonName `json:"name" eh:"optional"`
    Username string `json:"username" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Roles []string `json:"roles" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *AccountUpdated) AddToRoles(item string) string {
    o.Roles = append(o.Roles, item)
    return item
}


type AccountDeleted struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type AccountEnabled struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type AccountDisabled struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SendEnabledAccountConfirmationed struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SendDisabledAccountConfirmationed struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type AccountLogged struct {
    Username string `json:"username" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}


type SendCreatedAccountConfirmationed struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}




type AccountEventType struct {
	name  string
	ordinal int
}

func (o *AccountEventType) Name() string {
    return o.name
}

func (o *AccountEventType) Ordinal() int {
    return o.ordinal
}

func (o AccountEventType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *AccountEventType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := AccountEventTypes().ParseAccountEventType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid AccountEventType %q", lit.Name)
        }
	}
	return
}

func (o AccountEventType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *AccountEventType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := AccountEventTypes().ParseAccountEventType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid AccountEventType %q", lit)
        }
    }
    return
}

func (o *AccountEventType) IsAccountCreated() bool {
    return o == _accountEventTypes.AccountCreated()
}

func (o *AccountEventType) IsAccountUpdated() bool {
    return o == _accountEventTypes.AccountUpdated()
}

func (o *AccountEventType) IsAccountDeleted() bool {
    return o == _accountEventTypes.AccountDeleted()
}

func (o *AccountEventType) IsAccountEnabled() bool {
    return o == _accountEventTypes.AccountEnabled()
}

func (o *AccountEventType) IsAccountDisabled() bool {
    return o == _accountEventTypes.AccountDisabled()
}

func (o *AccountEventType) IsSendEnabledAccountConfirmationed() bool {
    return o == _accountEventTypes.SendEnabledAccountConfirmationed()
}

func (o *AccountEventType) IsSendDisabledAccountConfirmationed() bool {
    return o == _accountEventTypes.SendDisabledAccountConfirmationed()
}

func (o *AccountEventType) IsAccountLogged() bool {
    return o == _accountEventTypes.AccountLogged()
}

func (o *AccountEventType) IsSendCreatedAccountConfirmationed() bool {
    return o == _accountEventTypes.SendCreatedAccountConfirmationed()
}

type accountEventTypes struct {
	values []*AccountEventType
    literals []enum.Literal
}

var _accountEventTypes = &accountEventTypes{values: []*AccountEventType{
    {name: "AccountCreated", ordinal: 0},
    {name: "AccountUpdated", ordinal: 1},
    {name: "AccountDeleted", ordinal: 2},
    {name: "AccountEnabled", ordinal: 3},
    {name: "AccountDisabled", ordinal: 4},
    {name: "SendEnabledAccountConfirmationed", ordinal: 5},
    {name: "SendDisabledAccountConfirmationed", ordinal: 6},
    {name: "AccountLogged", ordinal: 7},
    {name: "SendCreatedAccountConfirmationed", ordinal: 8}},
}

func AccountEventTypes() *accountEventTypes {
	return _accountEventTypes
}

func (o *accountEventTypes) Values() []*AccountEventType {
	return o.values
}

func (o *accountEventTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *accountEventTypes) AccountCreated() *AccountEventType {
    return _accountEventTypes.values[0]
}

func (o *accountEventTypes) AccountUpdated() *AccountEventType {
    return _accountEventTypes.values[1]
}

func (o *accountEventTypes) AccountDeleted() *AccountEventType {
    return _accountEventTypes.values[2]
}

func (o *accountEventTypes) AccountEnabled() *AccountEventType {
    return _accountEventTypes.values[3]
}

func (o *accountEventTypes) AccountDisabled() *AccountEventType {
    return _accountEventTypes.values[4]
}

func (o *accountEventTypes) SendEnabledAccountConfirmationed() *AccountEventType {
    return _accountEventTypes.values[5]
}

func (o *accountEventTypes) SendDisabledAccountConfirmationed() *AccountEventType {
    return _accountEventTypes.values[6]
}

func (o *accountEventTypes) AccountLogged() *AccountEventType {
    return _accountEventTypes.values[7]
}

func (o *accountEventTypes) SendCreatedAccountConfirmationed() *AccountEventType {
    return _accountEventTypes.values[8]
}

func (o *accountEventTypes) ParseAccountEventType(name string) (ret *AccountEventType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*AccountEventType), ok
	}
	return
}



