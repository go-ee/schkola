package auth

import (
    "errors"
    "fmt"
    "github.com/go-ee/utils/eh"
    "github.com/looplab/eventhorizon"
)
type AccountConfirmationDisabledHandler struct {
    EnabledHandler func (*AccountEnabled, eventhorizon.Entity) (err error)  `json:"enabledHandler" eh:"optional"`
}

func NewAccountConfirmationDisabledHandler() (ret *AccountConfirmationDisabledHandler) {
    ret = &AccountConfirmationDisabledHandler{}
    return
}

func (o *AccountConfirmationDisabledHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case AccountEnabledEvent:
        err = o.EnabledHandler(event.Data().(*AccountEnabled), entity.(*Account))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AccountConfirmationDisabledHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(AccountEnabledEvent, func() eventhorizon.EventData {
		return &AccountEnabled{}
	})

    //default handler implementation
    o.EnabledHandler = func(event *AccountEnabled, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
        }
        return
    }
    return
}


type AccountConfirmationDisabledExecutor struct {
}

func NewAccountConfirmationDisabledExecutor() (ret *AccountConfirmationDisabledExecutor) {
    ret = &AccountConfirmationDisabledExecutor{}
    return
}


type AccountConfirmationEnabledHandler struct {
    DisabledHandler func (*AccountDisabled, eventhorizon.Entity) (err error)  `json:"disabledHandler" eh:"optional"`
}

func NewAccountConfirmationEnabledHandler() (ret *AccountConfirmationEnabledHandler) {
    ret = &AccountConfirmationEnabledHandler{}
    return
}

func (o *AccountConfirmationEnabledHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case AccountDisabledEvent:
        err = o.DisabledHandler(event.Data().(*AccountDisabled), entity.(*Account))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AccountConfirmationEnabledHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(AccountDisabledEvent, func() eventhorizon.EventData {
		return &AccountDisabled{}
	})

    //default handler implementation
    o.DisabledHandler = func(event *AccountDisabled, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
        }
        return
    }
    return
}


type AccountConfirmationEnabledExecutor struct {
}

func NewAccountConfirmationEnabledExecutor() (ret *AccountConfirmationEnabledExecutor) {
    ret = &AccountConfirmationEnabledExecutor{}
    return
}


type AccountConfirmationInitialHandler struct {
    CreatedHandler func (*AccountCreated, eventhorizon.Entity) (err error)  `json:"createdHandler" eh:"optional"`
    CreatedHandler func (*AccountCreated, eventhorizon.Entity) (err error)  `json:"createdHandler" eh:"optional"`
}

func NewAccountConfirmationInitialHandler() (ret *AccountConfirmationInitialHandler) {
    ret = &AccountConfirmationInitialHandler{}
    return
}

func (o *AccountConfirmationInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case AccountCreatedEvent:
        err = o.CreatedHandler(event.Data().(*AccountCreated), entity.(*Account))
    case AccountCreatedEvent:
        err = o.CreatedHandler(event.Data().(*AccountCreated), entity.(*Account))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AccountConfirmationInitialHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(AccountCreatedEvent, func() eventhorizon.EventData {
		return &AccountCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *AccountCreated, entity *Account) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, AccountAggregateType); err == nil {
            entity.Name = event.Name
            entity.Username = event.Username
            entity.Password = event.Password
            entity.Email = event.Email
            entity.Roles = event.Roles
            entity.Profile = event.Profile
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(AccountCreatedEvent, func() eventhorizon.EventData {
		return &AccountCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *AccountCreated, entity *Account) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, AccountAggregateType); err == nil {
            entity.Name = event.Name
            entity.Username = event.Username
            entity.Password = event.Password
            entity.Email = event.Email
            entity.Roles = event.Roles
            entity.Profile = event.Profile
            entity.Id = event.Id
        }
        return
    }
    return
}


type AccountConfirmationInitialExecutor struct {
}

func NewAccountConfirmationInitialExecutor() (ret *AccountConfirmationInitialExecutor) {
    ret = &AccountConfirmationInitialExecutor{}
    return
}


type AccountConfirmationHandlers struct {
    Disabled *AccountConfirmationDisabledHandler `json:"disabled" eh:"optional"`
    Enabled *AccountConfirmationEnabledHandler `json:"enabled" eh:"optional"`
    Initial *AccountConfirmationInitialHandler `json:"initial" eh:"optional"`
}

func NewAccountConfirmationHandlers() (ret *AccountConfirmationHandlers) {
    disabled := NewAccountConfirmationDisabledHandler()
    enabled := NewAccountConfirmationEnabledHandler()
    initial := NewAccountConfirmationInitialHandler()
    ret = &AccountConfirmationHandlers{
        Disabled: disabled,
        Enabled: enabled,
        Initial: initial,
    }
    return
}


type AccountConfirmationExecutors struct {
    Disabled *AccountConfirmationDisabledExecutor `json:"disabled" eh:"optional"`
    Enabled *AccountConfirmationEnabledExecutor `json:"enabled" eh:"optional"`
    Initial *AccountConfirmationInitialExecutor `json:"initial" eh:"optional"`
}

func NewAccountConfirmationExecutors() (ret *AccountConfirmationExecutors) {
    disabled := NewAccountConfirmationDisabledExecutor()
    enabled := NewAccountConfirmationEnabledExecutor()
    initial := NewAccountConfirmationInitialExecutor()
    ret = &AccountConfirmationExecutors{
        Disabled: disabled,
        Enabled: enabled,
        Initial: initial,
    }
    return
}


type AccountDeletedHandler struct {
}

func NewAccountDeletedHandler() (ret *AccountDeletedHandler) {
    ret = &AccountDeletedHandler{}
    return
}

func (o *AccountDeletedHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    
    return
}

func (o *AccountDeletedHandler) SetupEventHandler() (err error) {
    return
}


type AccountDeletedExecutor struct {
}

func NewAccountDeletedExecutor() (ret *AccountDeletedExecutor) {
    ret = &AccountDeletedExecutor{}
    return
}


type AccountDisabledHandler struct {
    EnabledHandler func (*AccountEnabled, eventhorizon.Entity) (err error)  `json:"enabledHandler" eh:"optional"`
}

func NewAccountDisabledHandler() (ret *AccountDisabledHandler) {
    ret = &AccountDisabledHandler{}
    return
}

func (o *AccountDisabledHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case AccountEnabledEvent:
        err = o.EnabledHandler(event.Data().(*AccountEnabled), entity.(*Account))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AccountDisabledHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(AccountEnabledEvent, func() eventhorizon.EventData {
		return &AccountEnabled{}
	})

    //default handler implementation
    o.EnabledHandler = func(event *AccountEnabled, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
        }
        return
    }
    return
}


type AccountDisabledExecutor struct {
}

func NewAccountDisabledExecutor() (ret *AccountDisabledExecutor) {
    ret = &AccountDisabledExecutor{}
    return
}


type AccountEnabledHandler struct {
    DisabledHandler func (*AccountDisabled, eventhorizon.Entity) (err error)  `json:"disabledHandler" eh:"optional"`
}

func NewAccountEnabledHandler() (ret *AccountEnabledHandler) {
    ret = &AccountEnabledHandler{}
    return
}

func (o *AccountEnabledHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case AccountDisabledEvent:
        err = o.DisabledHandler(event.Data().(*AccountDisabled), entity.(*Account))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AccountEnabledHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(AccountDisabledEvent, func() eventhorizon.EventData {
		return &AccountDisabled{}
	})

    //default handler implementation
    o.DisabledHandler = func(event *AccountDisabled, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
        }
        return
    }
    return
}


type AccountEnabledExecutor struct {
}

func NewAccountEnabledExecutor() (ret *AccountEnabledExecutor) {
    ret = &AccountEnabledExecutor{}
    return
}


type AccountExistHandler struct {
    UpdatedHandler func (*AccountUpdated, eventhorizon.Entity) (err error)  `json:"updatedHandler" eh:"optional"`
    DeletedHandler func (*AccountDeleted, eventhorizon.Entity) (err error)  `json:"deletedHandler" eh:"optional"`
}

func NewAccountExistHandler() (ret *AccountExistHandler) {
    ret = &AccountExistHandler{}
    return
}

func (o *AccountExistHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case AccountUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*AccountUpdated), entity.(*Account))
    case AccountDeletedEvent:
        err = o.DeletedHandler(event.Data().(*AccountDeleted), entity.(*Account))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AccountExistHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(AccountUpdatedEvent, func() eventhorizon.EventData {
		return &AccountUpdated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *AccountUpdated, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
            entity.Name = event.Name
            entity.Username = event.Username
            entity.Password = event.Password
            entity.Email = event.Email
            entity.Roles = event.Roles
            entity.Profile = event.Profile
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(AccountDeletedEvent, func() eventhorizon.EventData {
		return &AccountDeleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *AccountDeleted, entity *Account) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, AccountAggregateType); err == nil {
            *entity = *NewAccount()
        }
        return
    }
    return
}


type AccountExistExecutor struct {
}

func NewAccountExistExecutor() (ret *AccountExistExecutor) {
    ret = &AccountExistExecutor{}
    return
}


type AccountInitialHandler struct {
    CreatedHandler func (*AccountCreated, eventhorizon.Entity) (err error)  `json:"createdHandler" eh:"optional"`
    CreatedHandler func (*AccountCreated, eventhorizon.Entity) (err error)  `json:"createdHandler" eh:"optional"`
}

func NewAccountInitialHandler() (ret *AccountInitialHandler) {
    ret = &AccountInitialHandler{}
    return
}

func (o *AccountInitialHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case AccountCreatedEvent:
        err = o.CreatedHandler(event.Data().(*AccountCreated), entity.(*Account))
    case AccountCreatedEvent:
        err = o.CreatedHandler(event.Data().(*AccountCreated), entity.(*Account))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *AccountInitialHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(AccountCreatedEvent, func() eventhorizon.EventData {
		return &AccountCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *AccountCreated, entity *Account) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, AccountAggregateType); err == nil {
            entity.Name = event.Name
            entity.Username = event.Username
            entity.Password = event.Password
            entity.Email = event.Email
            entity.Roles = event.Roles
            entity.Profile = event.Profile
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(AccountCreatedEvent, func() eventhorizon.EventData {
		return &AccountCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *AccountCreated, entity *Account) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, AccountAggregateType); err == nil {
            entity.Name = event.Name
            entity.Username = event.Username
            entity.Password = event.Password
            entity.Email = event.Email
            entity.Roles = event.Roles
            entity.Profile = event.Profile
            entity.Id = event.Id
        }
        return
    }
    return
}


type AccountInitialExecutor struct {
}

func NewAccountInitialExecutor() (ret *AccountInitialExecutor) {
    ret = &AccountInitialExecutor{}
    return
}


type AccountHandlers struct {
    Deleted *AccountDeletedHandler `json:"deleted" eh:"optional"`
    Disabled *AccountDisabledHandler `json:"disabled" eh:"optional"`
    Enabled *AccountEnabledHandler `json:"enabled" eh:"optional"`
    Exist *AccountExistHandler `json:"exist" eh:"optional"`
    Initial *AccountInitialHandler `json:"initial" eh:"optional"`
}

func NewAccountHandlers() (ret *AccountHandlers) {
    deleted := NewAccountDeletedHandler()
    disabled := NewAccountDisabledHandler()
    enabled := NewAccountEnabledHandler()
    exist := NewAccountExistHandler()
    initial := NewAccountInitialHandler()
    ret = &AccountHandlers{
        Deleted: deleted,
        Disabled: disabled,
        Enabled: enabled,
        Exist: exist,
        Initial: initial,
    }
    return
}


type AccountExecutors struct {
    Deleted *AccountDeletedExecutor `json:"deleted" eh:"optional"`
    Disabled *AccountDisabledExecutor `json:"disabled" eh:"optional"`
    Enabled *AccountEnabledExecutor `json:"enabled" eh:"optional"`
    Exist *AccountExistExecutor `json:"exist" eh:"optional"`
    Initial *AccountInitialExecutor `json:"initial" eh:"optional"`
}

func NewAccountExecutors() (ret *AccountExecutors) {
    deleted := NewAccountDeletedExecutor()
    disabled := NewAccountDisabledExecutor()
    enabled := NewAccountEnabledExecutor()
    exist := NewAccountExistExecutor()
    initial := NewAccountInitialExecutor()
    ret = &AccountExecutors{
        Deleted: deleted,
        Disabled: disabled,
        Enabled: enabled,
        Exist: exist,
        Initial: initial,
    }
    return
}









