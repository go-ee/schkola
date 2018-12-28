package library

import (
    "errors"
    "fmt"
    "github.com/looplab/eventhorizon"
)
type BookInitialHandler struct {
}

func NewBookInitialHandler() (ret *BookInitialHandler) {
    ret = &BookInitialHandler{}
    return
}

func (o *BookInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    
    return
}

func (o *BookInitialHandler) SetupEventHandler() (err error) {
    return
}


type BookInitialExecutor struct {
}

func NewBookInitialExecutor() (ret *BookInitialExecutor) {
    ret = &BookInitialExecutor{}
    return
}


type BookHandlers struct {
    Initial *BookInitialHandler `json:"initial" eh:"optional"`
}

func NewBookHandlers() (ret *BookHandlers) {
    initial := NewBookInitialHandler()
    ret = &BookHandlers{
        Initial: initial,
    }
    return
}


type BookExecutors struct {
    Initial *BookInitialExecutor `json:"initial" eh:"optional"`
}

func NewBookExecutors() (ret *BookExecutors) {
    initial := NewBookInitialExecutor()
    ret = &BookExecutors{
        Initial: initial,
    }
    return
}









