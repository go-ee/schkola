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
     SendAccountEnabledConfirmationCommand eventhorizon.CommandType = "SendAccountEnabledConfirmation"
     SendAccountDisabledConfirmationCommand eventhorizon.CommandType = "SendAccountDisabledConfirmation"
     LoginAccountCommand eventhorizon.CommandType = "LoginAccount"
     SendAccountCreatedConfirmationCommand eventhorizon.CommandType = "SendAccountCreatedConfirmation"
     CreateAccountCommand eventhorizon.CommandType = "CreateAccount"
     DeleteAccountCommand eventhorizon.CommandType = "DeleteAccount"
     EnableAccountCommand eventhorizon.CommandType = "EnableAccount"
     DisableAccountCommand eventhorizon.CommandType = "DisableAccount"
     UpdateAccountCommand eventhorizon.CommandType = "UpdateAccount"
)




        
type SendAccountEnabledConfirmation struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *SendAccountEnabledConfirmation) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *SendAccountEnabledConfirmation) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *SendAccountEnabledConfirmation) CommandType() eventhorizon.CommandType      { return SendAccountEnabledConfirmationCommand }



        
type SendAccountDisabledConfirmation struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *SendAccountDisabledConfirmation) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *SendAccountDisabledConfirmation) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *SendAccountDisabledConfirmation) CommandType() eventhorizon.CommandType      { return SendAccountDisabledConfirmationCommand }



        
type LoginAccount struct {
    Username string `json:"username" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *LoginAccount) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *LoginAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *LoginAccount) CommandType() eventhorizon.CommandType      { return LoginAccountCommand }



        
type SendAccountCreatedConfirmation struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *SendAccountCreatedConfirmation) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *SendAccountCreatedConfirmation) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *SendAccountCreatedConfirmation) CommandType() eventhorizon.CommandType      { return SendAccountCreatedConfirmationCommand }



        
type CreateAccount struct {
    Name *shared.PersonName `json:"name" eh:"optional"`
    Username string `json:"username" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Roles []string `json:"roles" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *CreateAccount) AddToRoles(item string) string {
    o.Roles = append(o.Roles, item)
    return item
}
func (o *CreateAccount) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CreateAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *CreateAccount) CommandType() eventhorizon.CommandType      { return CreateAccountCommand }



        
type DeleteAccount struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *DeleteAccount) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DeleteAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *DeleteAccount) CommandType() eventhorizon.CommandType      { return DeleteAccountCommand }



        
type EnableAccount struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *EnableAccount) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *EnableAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *EnableAccount) CommandType() eventhorizon.CommandType      { return EnableAccountCommand }



        
type DisableAccount struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *DisableAccount) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DisableAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *DisableAccount) CommandType() eventhorizon.CommandType      { return DisableAccountCommand }



        
type UpdateAccount struct {
    Name *shared.PersonName `json:"name" eh:"optional"`
    Username string `json:"username" eh:"optional"`
    Password string `json:"password" eh:"optional"`
    Email string `json:"email" eh:"optional"`
    Roles []string `json:"roles" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}

func (o *UpdateAccount) AddToRoles(item string) string {
    o.Roles = append(o.Roles, item)
    return item
}
func (o *UpdateAccount) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *UpdateAccount) AggregateType() eventhorizon.AggregateType  { return AccountAggregateType }
func (o *UpdateAccount) CommandType() eventhorizon.CommandType      { return UpdateAccountCommand }





type AccountCommandType struct {
	name  string
	ordinal int
}

func (o *AccountCommandType) Name() string {
    return o.name
}

func (o *AccountCommandType) Ordinal() int {
    return o.ordinal
}

func (o AccountCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *AccountCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := AccountCommandTypes().ParseAccountCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid AccountCommandType %q", lit.Name)
        }
	}
	return
}

func (o AccountCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *AccountCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := AccountCommandTypes().ParseAccountCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid AccountCommandType %q", lit)
        }
    }
    return
}

func (o *AccountCommandType) IsSendAccountEnabledConfirmation() bool {
    return o == _accountCommandTypes.SendAccountEnabledConfirmation()
}

func (o *AccountCommandType) IsSendAccountDisabledConfirmation() bool {
    return o == _accountCommandTypes.SendAccountDisabledConfirmation()
}

func (o *AccountCommandType) IsLoginAccount() bool {
    return o == _accountCommandTypes.LoginAccount()
}

func (o *AccountCommandType) IsSendAccountCreatedConfirmation() bool {
    return o == _accountCommandTypes.SendAccountCreatedConfirmation()
}

func (o *AccountCommandType) IsCreateAccount() bool {
    return o == _accountCommandTypes.CreateAccount()
}

func (o *AccountCommandType) IsDeleteAccount() bool {
    return o == _accountCommandTypes.DeleteAccount()
}

func (o *AccountCommandType) IsEnableAccount() bool {
    return o == _accountCommandTypes.EnableAccount()
}

func (o *AccountCommandType) IsDisableAccount() bool {
    return o == _accountCommandTypes.DisableAccount()
}

func (o *AccountCommandType) IsUpdateAccount() bool {
    return o == _accountCommandTypes.UpdateAccount()
}

type accountCommandTypes struct {
	values []*AccountCommandType
    literals []enum.Literal
}

var _accountCommandTypes = &accountCommandTypes{values: []*AccountCommandType{
    {name: "SendAccountEnabledConfirmation", ordinal: 0},
    {name: "SendAccountDisabledConfirmation", ordinal: 1},
    {name: "LoginAccount", ordinal: 2},
    {name: "SendAccountCreatedConfirmation", ordinal: 3},
    {name: "CreateAccount", ordinal: 4},
    {name: "DeleteAccount", ordinal: 5},
    {name: "EnableAccount", ordinal: 6},
    {name: "DisableAccount", ordinal: 7},
    {name: "UpdateAccount", ordinal: 8}},
}

func AccountCommandTypes() *accountCommandTypes {
	return _accountCommandTypes
}

func (o *accountCommandTypes) Values() []*AccountCommandType {
	return o.values
}

func (o *accountCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *accountCommandTypes) SendAccountEnabledConfirmation() *AccountCommandType {
    return _accountCommandTypes.values[0]
}

func (o *accountCommandTypes) SendAccountDisabledConfirmation() *AccountCommandType {
    return _accountCommandTypes.values[1]
}

func (o *accountCommandTypes) LoginAccount() *AccountCommandType {
    return _accountCommandTypes.values[2]
}

func (o *accountCommandTypes) SendAccountCreatedConfirmation() *AccountCommandType {
    return _accountCommandTypes.values[3]
}

func (o *accountCommandTypes) CreateAccount() *AccountCommandType {
    return _accountCommandTypes.values[4]
}

func (o *accountCommandTypes) DeleteAccount() *AccountCommandType {
    return _accountCommandTypes.values[5]
}

func (o *accountCommandTypes) EnableAccount() *AccountCommandType {
    return _accountCommandTypes.values[6]
}

func (o *accountCommandTypes) DisableAccount() *AccountCommandType {
    return _accountCommandTypes.values[7]
}

func (o *accountCommandTypes) UpdateAccount() *AccountCommandType {
    return _accountCommandTypes.values[8]
}

func (o *accountCommandTypes) ParseAccountCommandType(name string) (ret *AccountCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*AccountCommandType), ok
	}
	return
}



