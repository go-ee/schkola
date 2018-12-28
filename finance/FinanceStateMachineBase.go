package finance

import (
    "errors"
    "fmt"
    "github.com/looplab/eventhorizon"
)
type ExpenseInitialHandler struct {
}

func NewExpenseInitialHandler() (ret *ExpenseInitialHandler) {
    ret = &ExpenseInitialHandler{}
    return
}

func (o *ExpenseInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    
    return
}

func (o *ExpenseInitialHandler) SetupEventHandler() (err error) {
    return
}


type ExpenseInitialExecutor struct {
}

func NewExpenseInitialExecutor() (ret *ExpenseInitialExecutor) {
    ret = &ExpenseInitialExecutor{}
    return
}


type ExpenseHandlers struct {
    Initial *ExpenseInitialHandler `json:"initial" eh:"optional"`
}

func NewExpenseHandlers() (ret *ExpenseHandlers) {
    initial := NewExpenseInitialHandler()
    ret = &ExpenseHandlers{
        Initial: initial,
    }
    return
}


type ExpenseExecutors struct {
    Initial *ExpenseInitialExecutor `json:"initial" eh:"optional"`
}

func NewExpenseExecutors() (ret *ExpenseExecutors) {
    initial := NewExpenseInitialExecutor()
    ret = &ExpenseExecutors{
        Initial: initial,
    }
    return
}


type ExpensePurposeInitialHandler struct {
}

func NewExpensePurposeInitialHandler() (ret *ExpensePurposeInitialHandler) {
    ret = &ExpensePurposeInitialHandler{}
    return
}

func (o *ExpensePurposeInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    
    return
}

func (o *ExpensePurposeInitialHandler) SetupEventHandler() (err error) {
    return
}


type ExpensePurposeInitialExecutor struct {
}

func NewExpensePurposeInitialExecutor() (ret *ExpensePurposeInitialExecutor) {
    ret = &ExpensePurposeInitialExecutor{}
    return
}


type ExpensePurposeHandlers struct {
    Initial *ExpensePurposeInitialHandler `json:"initial" eh:"optional"`
}

func NewExpensePurposeHandlers() (ret *ExpensePurposeHandlers) {
    initial := NewExpensePurposeInitialHandler()
    ret = &ExpensePurposeHandlers{
        Initial: initial,
    }
    return
}


type ExpensePurposeExecutors struct {
    Initial *ExpensePurposeInitialExecutor `json:"initial" eh:"optional"`
}

func NewExpensePurposeExecutors() (ret *ExpensePurposeExecutors) {
    initial := NewExpensePurposeInitialExecutor()
    ret = &ExpensePurposeExecutors{
        Initial: initial,
    }
    return
}


type FeeInitialHandler struct {
}

func NewFeeInitialHandler() (ret *FeeInitialHandler) {
    ret = &FeeInitialHandler{}
    return
}

func (o *FeeInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    
    return
}

func (o *FeeInitialHandler) SetupEventHandler() (err error) {
    return
}


type FeeInitialExecutor struct {
}

func NewFeeInitialExecutor() (ret *FeeInitialExecutor) {
    ret = &FeeInitialExecutor{}
    return
}


type FeeHandlers struct {
    Initial *FeeInitialHandler `json:"initial" eh:"optional"`
}

func NewFeeHandlers() (ret *FeeHandlers) {
    initial := NewFeeInitialHandler()
    ret = &FeeHandlers{
        Initial: initial,
    }
    return
}


type FeeExecutors struct {
    Initial *FeeInitialExecutor `json:"initial" eh:"optional"`
}

func NewFeeExecutors() (ret *FeeExecutors) {
    initial := NewFeeInitialExecutor()
    ret = &FeeExecutors{
        Initial: initial,
    }
    return
}


type FeeKindInitialHandler struct {
}

func NewFeeKindInitialHandler() (ret *FeeKindInitialHandler) {
    ret = &FeeKindInitialHandler{}
    return
}

func (o *FeeKindInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    
    return
}

func (o *FeeKindInitialHandler) SetupEventHandler() (err error) {
    return
}


type FeeKindInitialExecutor struct {
}

func NewFeeKindInitialExecutor() (ret *FeeKindInitialExecutor) {
    ret = &FeeKindInitialExecutor{}
    return
}


type FeeKindHandlers struct {
    Initial *FeeKindInitialHandler `json:"initial" eh:"optional"`
}

func NewFeeKindHandlers() (ret *FeeKindHandlers) {
    initial := NewFeeKindInitialHandler()
    ret = &FeeKindHandlers{
        Initial: initial,
    }
    return
}


type FeeKindExecutors struct {
    Initial *FeeKindInitialExecutor `json:"initial" eh:"optional"`
}

func NewFeeKindExecutors() (ret *FeeKindExecutors) {
    initial := NewFeeKindInitialExecutor()
    ret = &FeeKindExecutors{
        Initial: initial,
    }
    return
}









