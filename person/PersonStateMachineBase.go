package person

import (
    "github.com/looplab/eventhorizon"
)
type ChurchInitialHandler struct {
}

func NewChurchInitialHandler() (ret *ChurchInitialHandler) {
    ret = &ChurchInitialHandler{}
    return
}

func (o *ChurchInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    
    return
}

func (o *ChurchInitialHandler) SetupEventHandler() (err error) {
    return
}


type ChurchInitialExecutor struct {
}

func NewChurchInitialExecutor() (ret *ChurchInitialExecutor) {
    ret = &ChurchInitialExecutor{}
    return
}


type ChurchHandlers struct {
    Initial *ChurchInitialHandler `json:"initial" eh:"optional"`
}

func NewChurchHandlers() (ret *ChurchHandlers) {
    initial := NewChurchInitialHandler()
    ret = &ChurchHandlers{
        Initial: initial,
    }
    return
}


type ChurchExecutors struct {
    Initial *ChurchInitialExecutor `json:"initial" eh:"optional"`
}

func NewChurchExecutors() (ret *ChurchExecutors) {
    initial := NewChurchInitialExecutor()
    ret = &ChurchExecutors{
        Initial: initial,
    }
    return
}


type GraduationInitialHandler struct {
}

func NewGraduationInitialHandler() (ret *GraduationInitialHandler) {
    ret = &GraduationInitialHandler{}
    return
}

func (o *GraduationInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    
    return
}

func (o *GraduationInitialHandler) SetupEventHandler() (err error) {
    return
}


type GraduationInitialExecutor struct {
}

func NewGraduationInitialExecutor() (ret *GraduationInitialExecutor) {
    ret = &GraduationInitialExecutor{}
    return
}


type GraduationHandlers struct {
    Initial *GraduationInitialHandler `json:"initial" eh:"optional"`
}

func NewGraduationHandlers() (ret *GraduationHandlers) {
    initial := NewGraduationInitialHandler()
    ret = &GraduationHandlers{
        Initial: initial,
    }
    return
}


type GraduationExecutors struct {
    Initial *GraduationInitialExecutor `json:"initial" eh:"optional"`
}

func NewGraduationExecutors() (ret *GraduationExecutors) {
    initial := NewGraduationInitialExecutor()
    ret = &GraduationExecutors{
        Initial: initial,
    }
    return
}


type ProfileInitialHandler struct {
}

func NewProfileInitialHandler() (ret *ProfileInitialHandler) {
    ret = &ProfileInitialHandler{}
    return
}

func (o *ProfileInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    
    return
}

func (o *ProfileInitialHandler) SetupEventHandler() (err error) {
    return
}


type ProfileInitialExecutor struct {
}

func NewProfileInitialExecutor() (ret *ProfileInitialExecutor) {
    ret = &ProfileInitialExecutor{}
    return
}


type ProfileHandlers struct {
    Initial *ProfileInitialHandler `json:"initial" eh:"optional"`
}

func NewProfileHandlers() (ret *ProfileHandlers) {
    initial := NewProfileInitialHandler()
    ret = &ProfileHandlers{
        Initial: initial,
    }
    return
}


type ProfileExecutors struct {
    Initial *ProfileInitialExecutor `json:"initial" eh:"optional"`
}

func NewProfileExecutors() (ret *ProfileExecutors) {
    initial := NewProfileInitialExecutor()
    ret = &ProfileExecutors{
        Initial: initial,
    }
    return
}









