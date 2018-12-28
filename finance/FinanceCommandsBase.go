package finance

import (
    "ee/schkola/person"
    "encoding/json"
    "fmt"
    "github.com/eugeis/gee/enum"
    "github.com/looplab/eventhorizon"
    "gopkg.in/mgo.v2/bson"
    "time"
)
const (
     CreateExpenseCommand eventhorizon.CommandType = "CreateExpense"
     DeleteExpenseCommand eventhorizon.CommandType = "DeleteExpense"
     UpdateExpenseCommand eventhorizon.CommandType = "UpdateExpense"
)


const (
     CreateExpensePurposeCommand eventhorizon.CommandType = "CreateExpensePurpose"
     DeleteExpensePurposeCommand eventhorizon.CommandType = "DeleteExpensePurpose"
     UpdateExpensePurposeCommand eventhorizon.CommandType = "UpdateExpensePurpose"
)


const (
     CreateFeeCommand eventhorizon.CommandType = "CreateFee"
     DeleteFeeCommand eventhorizon.CommandType = "DeleteFee"
     UpdateFeeCommand eventhorizon.CommandType = "UpdateFee"
)


const (
     CreateFeeKindCommand eventhorizon.CommandType = "CreateFeeKind"
     DeleteFeeKindCommand eventhorizon.CommandType = "DeleteFeeKind"
     UpdateFeeKindCommand eventhorizon.CommandType = "UpdateFeeKind"
)




        
type CreateExpense struct {
    Purpose *ExpensePurpose `json:"purpose" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *CreateExpense) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CreateExpense) AggregateType() eventhorizon.AggregateType  { return ExpenseAggregateType }
func (o *CreateExpense) CommandType() eventhorizon.CommandType      { return CreateExpenseCommand }



        
type DeleteExpense struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *DeleteExpense) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DeleteExpense) AggregateType() eventhorizon.AggregateType  { return ExpenseAggregateType }
func (o *DeleteExpense) CommandType() eventhorizon.CommandType      { return DeleteExpenseCommand }



        
type UpdateExpense struct {
    Purpose *ExpensePurpose `json:"purpose" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Profile *person.Profile `json:"profile" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *UpdateExpense) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *UpdateExpense) AggregateType() eventhorizon.AggregateType  { return ExpenseAggregateType }
func (o *UpdateExpense) CommandType() eventhorizon.CommandType      { return UpdateExpenseCommand }



        
type CreateExpensePurpose struct {
    Name string `json:"name" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *CreateExpensePurpose) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CreateExpensePurpose) AggregateType() eventhorizon.AggregateType  { return ExpensePurposeAggregateType }
func (o *CreateExpensePurpose) CommandType() eventhorizon.CommandType      { return CreateExpensePurposeCommand }



        
type DeleteExpensePurpose struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *DeleteExpensePurpose) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DeleteExpensePurpose) AggregateType() eventhorizon.AggregateType  { return ExpensePurposeAggregateType }
func (o *DeleteExpensePurpose) CommandType() eventhorizon.CommandType      { return DeleteExpensePurposeCommand }



        
type UpdateExpensePurpose struct {
    Name string `json:"name" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *UpdateExpensePurpose) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *UpdateExpensePurpose) AggregateType() eventhorizon.AggregateType  { return ExpensePurposeAggregateType }
func (o *UpdateExpensePurpose) CommandType() eventhorizon.CommandType      { return UpdateExpensePurposeCommand }



        
type CreateFee struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Kind *FeeKind `json:"kind" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *CreateFee) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CreateFee) AggregateType() eventhorizon.AggregateType  { return FeeAggregateType }
func (o *CreateFee) CommandType() eventhorizon.CommandType      { return CreateFeeCommand }



        
type DeleteFee struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *DeleteFee) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DeleteFee) AggregateType() eventhorizon.AggregateType  { return FeeAggregateType }
func (o *DeleteFee) CommandType() eventhorizon.CommandType      { return DeleteFeeCommand }



        
type UpdateFee struct {
    Student *person.Profile `json:"student" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Kind *FeeKind `json:"kind" eh:"optional"`
    Date *time.Time `json:"date" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *UpdateFee) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *UpdateFee) AggregateType() eventhorizon.AggregateType  { return FeeAggregateType }
func (o *UpdateFee) CommandType() eventhorizon.CommandType      { return UpdateFeeCommand }



        
type CreateFeeKind struct {
    Name string `json:"name" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *CreateFeeKind) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *CreateFeeKind) AggregateType() eventhorizon.AggregateType  { return FeeKindAggregateType }
func (o *CreateFeeKind) CommandType() eventhorizon.CommandType      { return CreateFeeKindCommand }



        
type DeleteFeeKind struct {
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *DeleteFeeKind) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *DeleteFeeKind) AggregateType() eventhorizon.AggregateType  { return FeeKindAggregateType }
func (o *DeleteFeeKind) CommandType() eventhorizon.CommandType      { return DeleteFeeKindCommand }



        
type UpdateFeeKind struct {
    Name string `json:"name" eh:"optional"`
    Amount float64 `json:"amount" eh:"optional"`
    Description string `json:"description" eh:"optional"`
    Id eventhorizon.UUID `json:"id" eh:"optional"`
}
func (o *UpdateFeeKind) AggregateID() eventhorizon.UUID            { return o.Id }
func (o *UpdateFeeKind) AggregateType() eventhorizon.AggregateType  { return FeeKindAggregateType }
func (o *UpdateFeeKind) CommandType() eventhorizon.CommandType      { return UpdateFeeKindCommand }





type ExpenseCommandType struct {
	name  string
	ordinal int
}

func (o *ExpenseCommandType) Name() string {
    return o.name
}

func (o *ExpenseCommandType) Ordinal() int {
    return o.ordinal
}

func (o ExpenseCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *ExpenseCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := ExpenseCommandTypes().ParseExpenseCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ExpenseCommandType %q", lit.Name)
        }
	}
	return
}

func (o ExpenseCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *ExpenseCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := ExpenseCommandTypes().ParseExpenseCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ExpenseCommandType %q", lit)
        }
    }
    return
}

func (o *ExpenseCommandType) IsCreateExpense() bool {
    return o == _expenseCommandTypes.CreateExpense()
}

func (o *ExpenseCommandType) IsDeleteExpense() bool {
    return o == _expenseCommandTypes.DeleteExpense()
}

func (o *ExpenseCommandType) IsUpdateExpense() bool {
    return o == _expenseCommandTypes.UpdateExpense()
}

type expenseCommandTypes struct {
	values []*ExpenseCommandType
    literals []enum.Literal
}

var _expenseCommandTypes = &expenseCommandTypes{values: []*ExpenseCommandType{
    {name: "CreateExpense", ordinal: 0},
    {name: "DeleteExpense", ordinal: 1},
    {name: "UpdateExpense", ordinal: 2}},
}

func ExpenseCommandTypes() *expenseCommandTypes {
	return _expenseCommandTypes
}

func (o *expenseCommandTypes) Values() []*ExpenseCommandType {
	return o.values
}

func (o *expenseCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *expenseCommandTypes) CreateExpense() *ExpenseCommandType {
    return _expenseCommandTypes.values[0]
}

func (o *expenseCommandTypes) DeleteExpense() *ExpenseCommandType {
    return _expenseCommandTypes.values[1]
}

func (o *expenseCommandTypes) UpdateExpense() *ExpenseCommandType {
    return _expenseCommandTypes.values[2]
}

func (o *expenseCommandTypes) ParseExpenseCommandType(name string) (ret *ExpenseCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*ExpenseCommandType), ok
	}
	return
}


type ExpensePurposeCommandType struct {
	name  string
	ordinal int
}

func (o *ExpensePurposeCommandType) Name() string {
    return o.name
}

func (o *ExpensePurposeCommandType) Ordinal() int {
    return o.ordinal
}

func (o ExpensePurposeCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *ExpensePurposeCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := ExpensePurposeCommandTypes().ParseExpensePurposeCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ExpensePurposeCommandType %q", lit.Name)
        }
	}
	return
}

func (o ExpensePurposeCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *ExpensePurposeCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := ExpensePurposeCommandTypes().ParseExpensePurposeCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid ExpensePurposeCommandType %q", lit)
        }
    }
    return
}

func (o *ExpensePurposeCommandType) IsCreateExpensePurpose() bool {
    return o == _expensePurposeCommandTypes.CreateExpensePurpose()
}

func (o *ExpensePurposeCommandType) IsDeleteExpensePurpose() bool {
    return o == _expensePurposeCommandTypes.DeleteExpensePurpose()
}

func (o *ExpensePurposeCommandType) IsUpdateExpensePurpose() bool {
    return o == _expensePurposeCommandTypes.UpdateExpensePurpose()
}

type expensePurposeCommandTypes struct {
	values []*ExpensePurposeCommandType
    literals []enum.Literal
}

var _expensePurposeCommandTypes = &expensePurposeCommandTypes{values: []*ExpensePurposeCommandType{
    {name: "CreateExpensePurpose", ordinal: 0},
    {name: "DeleteExpensePurpose", ordinal: 1},
    {name: "UpdateExpensePurpose", ordinal: 2}},
}

func ExpensePurposeCommandTypes() *expensePurposeCommandTypes {
	return _expensePurposeCommandTypes
}

func (o *expensePurposeCommandTypes) Values() []*ExpensePurposeCommandType {
	return o.values
}

func (o *expensePurposeCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *expensePurposeCommandTypes) CreateExpensePurpose() *ExpensePurposeCommandType {
    return _expensePurposeCommandTypes.values[0]
}

func (o *expensePurposeCommandTypes) DeleteExpensePurpose() *ExpensePurposeCommandType {
    return _expensePurposeCommandTypes.values[1]
}

func (o *expensePurposeCommandTypes) UpdateExpensePurpose() *ExpensePurposeCommandType {
    return _expensePurposeCommandTypes.values[2]
}

func (o *expensePurposeCommandTypes) ParseExpensePurposeCommandType(name string) (ret *ExpensePurposeCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*ExpensePurposeCommandType), ok
	}
	return
}


type FeeCommandType struct {
	name  string
	ordinal int
}

func (o *FeeCommandType) Name() string {
    return o.name
}

func (o *FeeCommandType) Ordinal() int {
    return o.ordinal
}

func (o FeeCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *FeeCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := FeeCommandTypes().ParseFeeCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid FeeCommandType %q", lit.Name)
        }
	}
	return
}

func (o FeeCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *FeeCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := FeeCommandTypes().ParseFeeCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid FeeCommandType %q", lit)
        }
    }
    return
}

func (o *FeeCommandType) IsCreateFee() bool {
    return o == _feeCommandTypes.CreateFee()
}

func (o *FeeCommandType) IsDeleteFee() bool {
    return o == _feeCommandTypes.DeleteFee()
}

func (o *FeeCommandType) IsUpdateFee() bool {
    return o == _feeCommandTypes.UpdateFee()
}

type feeCommandTypes struct {
	values []*FeeCommandType
    literals []enum.Literal
}

var _feeCommandTypes = &feeCommandTypes{values: []*FeeCommandType{
    {name: "CreateFee", ordinal: 0},
    {name: "DeleteFee", ordinal: 1},
    {name: "UpdateFee", ordinal: 2}},
}

func FeeCommandTypes() *feeCommandTypes {
	return _feeCommandTypes
}

func (o *feeCommandTypes) Values() []*FeeCommandType {
	return o.values
}

func (o *feeCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *feeCommandTypes) CreateFee() *FeeCommandType {
    return _feeCommandTypes.values[0]
}

func (o *feeCommandTypes) DeleteFee() *FeeCommandType {
    return _feeCommandTypes.values[1]
}

func (o *feeCommandTypes) UpdateFee() *FeeCommandType {
    return _feeCommandTypes.values[2]
}

func (o *feeCommandTypes) ParseFeeCommandType(name string) (ret *FeeCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*FeeCommandType), ok
	}
	return
}


type FeeKindCommandType struct {
	name  string
	ordinal int
}

func (o *FeeKindCommandType) Name() string {
    return o.name
}

func (o *FeeKindCommandType) Ordinal() int {
    return o.ordinal
}

func (o FeeKindCommandType) MarshalJSON() (ret []byte, err error) {
	return json.Marshal(&enum.EnumBaseJson{Name: o.name})
}

func (o *FeeKindCommandType) UnmarshalJSON(data []byte) (err error) {
	lit := enum.EnumBaseJson{}
	if err = json.Unmarshal(data, &lit); err == nil {
		if v, ok := FeeKindCommandTypes().ParseFeeKindCommandType(lit.Name); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid FeeKindCommandType %q", lit.Name)
        }
	}
	return
}

func (o FeeKindCommandType) GetBSON() (ret interface{}, err error) {
	return o.name, nil
}

func (o *FeeKindCommandType) SetBSON(raw bson.Raw) (err error) {
	var lit string
    if err = raw.Unmarshal(&lit); err == nil {
		if v, ok := FeeKindCommandTypes().ParseFeeKindCommandType(lit); ok {
            *o = *v
        } else {
            err = fmt.Errorf("invalid FeeKindCommandType %q", lit)
        }
    }
    return
}

func (o *FeeKindCommandType) IsCreateFeeKind() bool {
    return o == _feeKindCommandTypes.CreateFeeKind()
}

func (o *FeeKindCommandType) IsDeleteFeeKind() bool {
    return o == _feeKindCommandTypes.DeleteFeeKind()
}

func (o *FeeKindCommandType) IsUpdateFeeKind() bool {
    return o == _feeKindCommandTypes.UpdateFeeKind()
}

type feeKindCommandTypes struct {
	values []*FeeKindCommandType
    literals []enum.Literal
}

var _feeKindCommandTypes = &feeKindCommandTypes{values: []*FeeKindCommandType{
    {name: "CreateFeeKind", ordinal: 0},
    {name: "DeleteFeeKind", ordinal: 1},
    {name: "UpdateFeeKind", ordinal: 2}},
}

func FeeKindCommandTypes() *feeKindCommandTypes {
	return _feeKindCommandTypes
}

func (o *feeKindCommandTypes) Values() []*FeeKindCommandType {
	return o.values
}

func (o *feeKindCommandTypes) Literals() []enum.Literal {
	if o.literals == nil {
		o.literals = make([]enum.Literal, len(o.values))
		for i, item := range o.values {
			o.literals[i] = item
		}
	}
	return o.literals
}

func (o *feeKindCommandTypes) CreateFeeKind() *FeeKindCommandType {
    return _feeKindCommandTypes.values[0]
}

func (o *feeKindCommandTypes) DeleteFeeKind() *FeeKindCommandType {
    return _feeKindCommandTypes.values[1]
}

func (o *feeKindCommandTypes) UpdateFeeKind() *FeeKindCommandType {
    return _feeKindCommandTypes.values[2]
}

func (o *feeKindCommandTypes) ParseFeeKindCommandType(name string) (ret *FeeKindCommandType, ok bool) {
	if item, ok := enum.Parse(name, o.Literals()); ok {
		return item.(*FeeKindCommandType), ok
	}
	return
}



