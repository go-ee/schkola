package person

import (
    "errors"
    "fmt"
    "github.com/eugeis/gee/eh"
    "github.com/looplab/eventhorizon"
    "github.com/looplab/eventhorizon/commandhandler/bus"
    "time"
)
type ChurchCommandHandler struct {
    CreateHandler func (*CreateChurch, *Church, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*DeleteChurch, *Church, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*UpdateChurch, *Church, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *ChurchCommandHandler) AddCreatePreparer(preparer func (*CreateChurch, *Church) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *CreateChurch, entity *Church, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *ChurchCommandHandler) AddDeletePreparer(preparer func (*DeleteChurch, *Church) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *DeleteChurch, entity *Church, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *ChurchCommandHandler) AddUpdatePreparer(preparer func (*UpdateChurch, *Church) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *UpdateChurch, entity *Church, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *ChurchCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateChurchCommand:
        err = o.CreateHandler(cmd.(*CreateChurch), entity.(*Church), store)
    case DeleteChurchCommand:
        err = o.DeleteHandler(cmd.(*DeleteChurch), entity.(*Church), store)
    case UpdateChurchCommand:
        err = o.UpdateHandler(cmd.(*UpdateChurch), entity.(*Church), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *ChurchCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *CreateChurch, entity *Church, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, ChurchAggregateType); err == nil {
            store.StoreEvent(ChurchCreatedEvent, &ChurchCreated{
                Name: command.Name,
                Address: command.Address,
                Pastor: command.Pastor,
                Contact: command.Contact,
                Association: command.Association,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *DeleteChurch, entity *Church, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, ChurchAggregateType); err == nil {
            store.StoreEvent(ChurchDeletedEvent, &ChurchDeleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *UpdateChurch, entity *Church, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, ChurchAggregateType); err == nil {
            store.StoreEvent(ChurchUpdatedEvent, &ChurchUpdated{
                Name: command.Name,
                Address: command.Address,
                Pastor: command.Pastor,
                Contact: command.Contact,
                Association: command.Association,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type ChurchEventHandler struct {
    CreatedHandler func (*ChurchCreated, *Church) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*ChurchDeleted, *Church) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*ChurchUpdated, *Church) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *ChurchEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case ChurchCreatedEvent:
        err = o.CreatedHandler(event.Data().(*ChurchCreated), entity.(*Church))
    case ChurchDeletedEvent:
        err = o.DeletedHandler(event.Data().(*ChurchDeleted), entity.(*Church))
    case ChurchUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*ChurchUpdated), entity.(*Church))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *ChurchEventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(ChurchCreatedEvent, func() eventhorizon.EventData {
		return &ChurchCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *ChurchCreated, entity *Church) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, ChurchAggregateType); err == nil {
            entity.Name = event.Name
            entity.Address = event.Address
            entity.Pastor = event.Pastor
            entity.Contact = event.Contact
            entity.Association = event.Association
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(ChurchDeletedEvent, func() eventhorizon.EventData {
		return &ChurchDeleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *ChurchDeleted, entity *Church) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, ChurchAggregateType); err == nil {
            *entity = *NewChurch()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(ChurchUpdatedEvent, func() eventhorizon.EventData {
		return &ChurchUpdated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *ChurchUpdated, entity *Church) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, ChurchAggregateType); err == nil {
            entity.Name = event.Name
            entity.Address = event.Address
            entity.Pastor = event.Pastor
            entity.Contact = event.Contact
            entity.Association = event.Association
        }
        return
    }
    return
}


const ChurchAggregateType eventhorizon.AggregateType = "Church"

type ChurchAggregateInitializer struct {
    *eh.AggregateInitializer
    *ChurchCommandHandler
    *ChurchEventHandler
    ProjectorHandler *ChurchEventHandler `json:"projectorHandler" eh:"optional"`
}



func NewChurchAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *ChurchAggregateInitializer) {
    
    commandHandler := &ChurchCommandHandler{}
    eventHandler := &ChurchEventHandler{}
    entityFactory := func() eventhorizon.Entity { return NewChurch() }
    ret = &ChurchAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ChurchAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(ChurchAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        ChurchCommandTypes().Literals(), ChurchEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, commandBus, readRepos), ChurchCommandHandler: commandHandler, ChurchEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type GraduationCommandHandler struct {
    CreateHandler func (*CreateGraduation, *Graduation, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*DeleteGraduation, *Graduation, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*UpdateGraduation, *Graduation, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *GraduationCommandHandler) AddCreatePreparer(preparer func (*CreateGraduation, *Graduation) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *CreateGraduation, entity *Graduation, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *GraduationCommandHandler) AddDeletePreparer(preparer func (*DeleteGraduation, *Graduation) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *DeleteGraduation, entity *Graduation, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *GraduationCommandHandler) AddUpdatePreparer(preparer func (*UpdateGraduation, *Graduation) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *UpdateGraduation, entity *Graduation, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *GraduationCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateGraduationCommand:
        err = o.CreateHandler(cmd.(*CreateGraduation), entity.(*Graduation), store)
    case DeleteGraduationCommand:
        err = o.DeleteHandler(cmd.(*DeleteGraduation), entity.(*Graduation), store)
    case UpdateGraduationCommand:
        err = o.UpdateHandler(cmd.(*UpdateGraduation), entity.(*Graduation), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *GraduationCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *CreateGraduation, entity *Graduation, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, GraduationAggregateType); err == nil {
            store.StoreEvent(GraduationCreatedEvent, &GraduationCreated{
                Name: command.Name,
                Level: command.Level,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *DeleteGraduation, entity *Graduation, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, GraduationAggregateType); err == nil {
            store.StoreEvent(GraduationDeletedEvent, &GraduationDeleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *UpdateGraduation, entity *Graduation, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, GraduationAggregateType); err == nil {
            store.StoreEvent(GraduationUpdatedEvent, &GraduationUpdated{
                Name: command.Name,
                Level: command.Level,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type GraduationEventHandler struct {
    CreatedHandler func (*GraduationCreated, *Graduation) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*GraduationDeleted, *Graduation) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*GraduationUpdated, *Graduation) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *GraduationEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case GraduationCreatedEvent:
        err = o.CreatedHandler(event.Data().(*GraduationCreated), entity.(*Graduation))
    case GraduationDeletedEvent:
        err = o.DeletedHandler(event.Data().(*GraduationDeleted), entity.(*Graduation))
    case GraduationUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*GraduationUpdated), entity.(*Graduation))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *GraduationEventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(GraduationCreatedEvent, func() eventhorizon.EventData {
		return &GraduationCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *GraduationCreated, entity *Graduation) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, GraduationAggregateType); err == nil {
            entity.Name = event.Name
            entity.Level = event.Level
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(GraduationDeletedEvent, func() eventhorizon.EventData {
		return &GraduationDeleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *GraduationDeleted, entity *Graduation) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, GraduationAggregateType); err == nil {
            *entity = *NewGraduation()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(GraduationUpdatedEvent, func() eventhorizon.EventData {
		return &GraduationUpdated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *GraduationUpdated, entity *Graduation) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, GraduationAggregateType); err == nil {
            entity.Name = event.Name
            entity.Level = event.Level
        }
        return
    }
    return
}


const GraduationAggregateType eventhorizon.AggregateType = "Graduation"

type GraduationAggregateInitializer struct {
    *eh.AggregateInitializer
    *GraduationCommandHandler
    *GraduationEventHandler
    ProjectorHandler *GraduationEventHandler `json:"projectorHandler" eh:"optional"`
}



func NewGraduationAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *GraduationAggregateInitializer) {
    
    commandHandler := &GraduationCommandHandler{}
    eventHandler := &GraduationEventHandler{}
    entityFactory := func() eventhorizon.Entity { return NewGraduation() }
    ret = &GraduationAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(GraduationAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(GraduationAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        GraduationCommandTypes().Literals(), GraduationEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, commandBus, readRepos), GraduationCommandHandler: commandHandler, GraduationEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type ProfileCommandHandler struct {
    CreateHandler func (*CreateProfile, *Profile, eh.AggregateStoreEvent) (err error)  `json:"createHandler" eh:"optional"`
    DeleteHandler func (*DeleteProfile, *Profile, eh.AggregateStoreEvent) (err error)  `json:"deleteHandler" eh:"optional"`
    UpdateHandler func (*UpdateProfile, *Profile, eh.AggregateStoreEvent) (err error)  `json:"updateHandler" eh:"optional"`
}

func (o *ProfileCommandHandler) AddCreatePreparer(preparer func (*CreateProfile, *Profile) (err error) ) {
    prevHandler := o.CreateHandler
	o.CreateHandler = func(command *CreateProfile, entity *Profile, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *ProfileCommandHandler) AddDeletePreparer(preparer func (*DeleteProfile, *Profile) (err error) ) {
    prevHandler := o.DeleteHandler
	o.DeleteHandler = func(command *DeleteProfile, entity *Profile, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *ProfileCommandHandler) AddUpdatePreparer(preparer func (*UpdateProfile, *Profile) (err error) ) {
    prevHandler := o.UpdateHandler
	o.UpdateHandler = func(command *UpdateProfile, entity *Profile, store eh.AggregateStoreEvent) (err error) {
		if err = preparer(command, entity); err == nil {
			err = prevHandler(command, entity, store)
		}
		return
	}
}

func (o *ProfileCommandHandler) Execute(cmd eventhorizon.Command, entity eventhorizon.Entity, store eh.AggregateStoreEvent) (err error) {
    switch cmd.CommandType() {
    case CreateProfileCommand:
        err = o.CreateHandler(cmd.(*CreateProfile), entity.(*Profile), store)
    case DeleteProfileCommand:
        err = o.DeleteHandler(cmd.(*DeleteProfile), entity.(*Profile), store)
    case UpdateProfileCommand:
        err = o.UpdateHandler(cmd.(*UpdateProfile), entity.(*Profile), store)
    default:
		err = errors.New(fmt.Sprintf("Not supported command type '%v' for entity '%v", cmd.CommandType(), entity))
	}
    return
}

func (o *ProfileCommandHandler) SetupCommandHandler() (err error) {
    o.CreateHandler = func(command *CreateProfile, entity *Profile, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateNewId(entity.Id, command.Id, ProfileAggregateType); err == nil {
            store.StoreEvent(ProfileCreatedEvent, &ProfileCreated{
                Gender: command.Gender,
                Name: command.Name,
                BirthName: command.BirthName,
                Birthday: command.Birthday,
                Address: command.Address,
                Contact: command.Contact,
                PhotoData: command.PhotoData,
                Photo: command.Photo,
                Family: command.Family,
                Church: command.Church,
                Education: command.Education,
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.DeleteHandler = func(command *DeleteProfile, entity *Profile, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, ProfileAggregateType); err == nil {
            store.StoreEvent(ProfileDeletedEvent, &ProfileDeleted{
                Id: command.Id,}, time.Now())
        }
        return
    }
    o.UpdateHandler = func(command *UpdateProfile, entity *Profile, store eh.AggregateStoreEvent) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, command.Id, ProfileAggregateType); err == nil {
            store.StoreEvent(ProfileUpdatedEvent, &ProfileUpdated{
                Gender: command.Gender,
                Name: command.Name,
                BirthName: command.BirthName,
                Birthday: command.Birthday,
                Address: command.Address,
                Contact: command.Contact,
                PhotoData: command.PhotoData,
                Photo: command.Photo,
                Family: command.Family,
                Church: command.Church,
                Education: command.Education,
                Id: command.Id,}, time.Now())
        }
        return
    }
    return
}


type ProfileEventHandler struct {
    CreatedHandler func (*ProfileCreated, *Profile) (err error)  `json:"createdHandler" eh:"optional"`
    DeletedHandler func (*ProfileDeleted, *Profile) (err error)  `json:"deletedHandler" eh:"optional"`
    UpdatedHandler func (*ProfileUpdated, *Profile) (err error)  `json:"updatedHandler" eh:"optional"`
}

func (o *ProfileEventHandler) Apply(event eventhorizon.Event, entity eventhorizon.Entity) (err error) {
    switch event.EventType() {
    case ProfileCreatedEvent:
        err = o.CreatedHandler(event.Data().(*ProfileCreated), entity.(*Profile))
    case ProfileDeletedEvent:
        err = o.DeletedHandler(event.Data().(*ProfileDeleted), entity.(*Profile))
    case ProfileUpdatedEvent:
        err = o.UpdatedHandler(event.Data().(*ProfileUpdated), entity.(*Profile))
    default:
		err = errors.New(fmt.Sprintf("Not supported event type '%v' for entity '%v", event.EventType(), entity))
	}
    return
}

func (o *ProfileEventHandler) SetupEventHandler() (err error) {

    //register event object factory
    eventhorizon.RegisterEventData(ProfileCreatedEvent, func() eventhorizon.EventData {
		return &ProfileCreated{}
	})

    //default handler implementation
    o.CreatedHandler = func(event *ProfileCreated, entity *Profile) (err error) {
        if err = eh.ValidateNewId(entity.Id, event.Id, ProfileAggregateType); err == nil {
            entity.Gender = event.Gender
            entity.Name = event.Name
            entity.BirthName = event.BirthName
            entity.Birthday = event.Birthday
            entity.Address = event.Address
            entity.Contact = event.Contact
            entity.PhotoData = event.PhotoData
            entity.Photo = event.Photo
            entity.Family = event.Family
            entity.Church = event.Church
            entity.Education = event.Education
            entity.Id = event.Id
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(ProfileDeletedEvent, func() eventhorizon.EventData {
		return &ProfileDeleted{}
	})

    //default handler implementation
    o.DeletedHandler = func(event *ProfileDeleted, entity *Profile) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, ProfileAggregateType); err == nil {
            *entity = *NewProfile()
        }
        return
    }

    //register event object factory
    eventhorizon.RegisterEventData(ProfileUpdatedEvent, func() eventhorizon.EventData {
		return &ProfileUpdated{}
	})

    //default handler implementation
    o.UpdatedHandler = func(event *ProfileUpdated, entity *Profile) (err error) {
        if err = eh.ValidateIdsMatch(entity.Id, event.Id, ProfileAggregateType); err == nil {
            entity.Gender = event.Gender
            entity.Name = event.Name
            entity.BirthName = event.BirthName
            entity.Birthday = event.Birthday
            entity.Address = event.Address
            entity.Contact = event.Contact
            entity.PhotoData = event.PhotoData
            entity.Photo = event.Photo
            entity.Family = event.Family
            entity.Church = event.Church
            entity.Education = event.Education
        }
        return
    }
    return
}


const ProfileAggregateType eventhorizon.AggregateType = "Profile"

type ProfileAggregateInitializer struct {
    *eh.AggregateInitializer
    *ProfileCommandHandler
    *ProfileEventHandler
    ProjectorHandler *ProfileEventHandler `json:"projectorHandler" eh:"optional"`
}



func NewProfileAggregateInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *ProfileAggregateInitializer) {
    
    commandHandler := &ProfileCommandHandler{}
    eventHandler := &ProfileEventHandler{}
    entityFactory := func() eventhorizon.Entity { return NewProfile() }
    ret = &ProfileAggregateInitializer{AggregateInitializer: eh.NewAggregateInitializer(ProfileAggregateType,
        func(id eventhorizon.UUID) eventhorizon.Aggregate {
            return eh.NewAggregateBase(ProfileAggregateType, id, commandHandler, eventHandler, entityFactory())
        }, entityFactory,
        ProfileCommandTypes().Literals(), ProfileEventTypes().Literals(), eventHandler,
        []func() error{commandHandler.SetupCommandHandler, eventHandler.SetupEventHandler},
        eventStore, eventBus, commandBus, readRepos), ProfileCommandHandler: commandHandler, ProfileEventHandler: eventHandler, ProjectorHandler: eventHandler,
    }

    return
}


type PersonEventhorizonInitializer struct {
    eventStore eventhorizon.EventStore `json:"eventStore" eh:"optional"`
    eventBus eventhorizon.EventBus `json:"eventBus" eh:"optional"`
    commandBus *bus.CommandHandler `json:"commandBus" eh:"optional"`
    ChurchAggregateInitializer *ChurchAggregateInitializer `json:"churchAggregateInitializer" eh:"optional"`
    GraduationAggregateInitializer *GraduationAggregateInitializer `json:"graduationAggregateInitializer" eh:"optional"`
    ProfileAggregateInitializer *ProfileAggregateInitializer `json:"profileAggregateInitializer" eh:"optional"`
}

func NewPersonEventhorizonInitializer(eventStore eventhorizon.EventStore, eventBus eventhorizon.EventBus, commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *PersonEventhorizonInitializer) {
    churchAggregateInitializer := NewChurchAggregateInitializer(eventStore, eventBus, commandBus, readRepos)
    graduationAggregateInitializer := NewGraduationAggregateInitializer(eventStore, eventBus, commandBus, readRepos)
    profileAggregateInitializer := NewProfileAggregateInitializer(eventStore, eventBus, commandBus, readRepos)
    ret = &PersonEventhorizonInitializer{
        eventStore: eventStore,
        eventBus: eventBus,
        commandBus: commandBus,
        ChurchAggregateInitializer: churchAggregateInitializer,
        GraduationAggregateInitializer: graduationAggregateInitializer,
        ProfileAggregateInitializer: profileAggregateInitializer,
    }
    return
}

func (o *PersonEventhorizonInitializer) Setup() (err error) {
    
    if err = o.ChurchAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.GraduationAggregateInitializer.Setup(); err != nil {
        return
    }
    
    if err = o.ProfileAggregateInitializer.Setup(); err != nil {
        return
    }

    return
}









