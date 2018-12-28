package student

import (
    "context"
    "github.com/eugeis/gee/eh"
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "github.com/looplab/eventhorizon"
    "github.com/looplab/eventhorizon/commandhandler/bus"
    "net/http"
)
type AttendanceHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *AttendanceQueryRepository `json:"queryRepository" eh:"optional"`
}

func NewAttendanceHttpQueryHandler(queryRepository *AttendanceQueryRepository) (ret *AttendanceHttpQueryHandler) {
    httpQueryHandler := eh.NewHttpQueryHandler()
    ret = &AttendanceHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *AttendanceHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllAttendance", w, r)
}

func (o *AttendanceHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByAttendanceId", w, r)
}

func (o *AttendanceHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllAttendance", w, r)
}

func (o *AttendanceHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByAttendanceId", w, r)
}

func (o *AttendanceHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllAttendance", w, r)
}

func (o *AttendanceHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByAttendanceId", w, r)
}


type AttendanceHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewAttendanceHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandHandler) (ret *AttendanceHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &AttendanceHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *AttendanceHttpCommandHandler) Register(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&RegisterAttendance{Id: id}, w, r)
}

func (o *AttendanceHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateAttendance{Id: id}, w, r)
}

func (o *AttendanceHttpCommandHandler) Confirm(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&ConfirmAttendance{Id: id}, w, r)
}

func (o *AttendanceHttpCommandHandler) Cancel(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CancelAttendance{Id: id}, w, r)
}

func (o *AttendanceHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateAttendance{Id: id}, w, r)
}

func (o *AttendanceHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteAttendance{Id: id}, w, r)
}


type AttendanceRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    QueryHandler *AttendanceHttpQueryHandler `json:"queryHandler" eh:"optional"`
    CommandHandler *AttendanceHttpCommandHandler `json:"commandHandler" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func NewAttendanceRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AttendanceRouter) {
    pathPrefix = pathPrefix + "/" + "attendances"
    entityFactory := func() eventhorizon.Entity { return NewAttendance() }
    repo := readRepos(string(AttendanceAggregateType), entityFactory)
    queryRepository := NewAttendanceQueryRepository(repo, context)
    queryHandler := NewAttendanceHttpQueryHandler(queryRepository)
    commandHandler := NewAttendanceHttpCommandHandler(context, commandBus)
    ret = &AttendanceRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *AttendanceRouter) Setup(router *mux.Router) (err error) {
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountAttendanceById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("CountAttendanceAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistAttendanceById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("ExistAttendanceAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindAttendanceById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("FindAttendanceAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Queries(net.Command, "register").
        Name("RegisterAttendance").HandlerFunc(o.CommandHandler.Register)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateAttendance").HandlerFunc(o.CommandHandler.Create)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Queries(net.Command, "confirm").
        Name("ConfirmAttendance").HandlerFunc(o.CommandHandler.Confirm)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Queries(net.Command, "cancel").
        Name("CancelAttendance").HandlerFunc(o.CommandHandler.Cancel)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateAttendance").HandlerFunc(o.CommandHandler.Update)
    router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteAttendance").HandlerFunc(o.CommandHandler.Delete)
    return
}


type CourseHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *CourseQueryRepository `json:"queryRepository" eh:"optional"`
}

func NewCourseHttpQueryHandler(queryRepository *CourseQueryRepository) (ret *CourseHttpQueryHandler) {
    httpQueryHandler := eh.NewHttpQueryHandler()
    ret = &CourseHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *CourseHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllCourse", w, r)
}

func (o *CourseHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByCourseId", w, r)
}

func (o *CourseHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllCourse", w, r)
}

func (o *CourseHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByCourseId", w, r)
}

func (o *CourseHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllCourse", w, r)
}

func (o *CourseHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByCourseId", w, r)
}


type CourseHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewCourseHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandHandler) (ret *CourseHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &CourseHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *CourseHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateCourse{Id: id}, w, r)
}

func (o *CourseHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateCourse{Id: id}, w, r)
}

func (o *CourseHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteCourse{Id: id}, w, r)
}


type CourseRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    QueryHandler *CourseHttpQueryHandler `json:"queryHandler" eh:"optional"`
    CommandHandler *CourseHttpCommandHandler `json:"commandHandler" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func NewCourseRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *CourseRouter) {
    pathPrefix = pathPrefix + "/" + "courses"
    entityFactory := func() eventhorizon.Entity { return NewCourse() }
    repo := readRepos(string(CourseAggregateType), entityFactory)
    queryRepository := NewCourseQueryRepository(repo, context)
    queryHandler := NewCourseHttpQueryHandler(queryRepository)
    commandHandler := NewCourseHttpCommandHandler(context, commandBus)
    ret = &CourseRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *CourseRouter) Setup(router *mux.Router) (err error) {
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountCourseById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("CountCourseAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistCourseById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("ExistCourseAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindCourseById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("FindCourseAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateCourse").HandlerFunc(o.CommandHandler.Create)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateCourse").HandlerFunc(o.CommandHandler.Update)
    router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteCourse").HandlerFunc(o.CommandHandler.Delete)
    return
}


type GradeHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *GradeQueryRepository `json:"queryRepository" eh:"optional"`
}

func NewGradeHttpQueryHandler(queryRepository *GradeQueryRepository) (ret *GradeHttpQueryHandler) {
    httpQueryHandler := eh.NewHttpQueryHandler()
    ret = &GradeHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *GradeHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllGrade", w, r)
}

func (o *GradeHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByGradeId", w, r)
}

func (o *GradeHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllGrade", w, r)
}

func (o *GradeHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByGradeId", w, r)
}

func (o *GradeHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllGrade", w, r)
}

func (o *GradeHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByGradeId", w, r)
}


type GradeHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewGradeHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandHandler) (ret *GradeHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &GradeHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *GradeHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateGrade{Id: id}, w, r)
}

func (o *GradeHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateGrade{Id: id}, w, r)
}

func (o *GradeHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteGrade{Id: id}, w, r)
}


type GradeRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    QueryHandler *GradeHttpQueryHandler `json:"queryHandler" eh:"optional"`
    CommandHandler *GradeHttpCommandHandler `json:"commandHandler" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func NewGradeRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *GradeRouter) {
    pathPrefix = pathPrefix + "/" + "grades"
    entityFactory := func() eventhorizon.Entity { return NewGrade() }
    repo := readRepos(string(GradeAggregateType), entityFactory)
    queryRepository := NewGradeQueryRepository(repo, context)
    queryHandler := NewGradeHttpQueryHandler(queryRepository)
    commandHandler := NewGradeHttpCommandHandler(context, commandBus)
    ret = &GradeRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *GradeRouter) Setup(router *mux.Router) (err error) {
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountGradeById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("CountGradeAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistGradeById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("ExistGradeAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindGradeById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("FindGradeAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateGrade").HandlerFunc(o.CommandHandler.Create)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateGrade").HandlerFunc(o.CommandHandler.Update)
    router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteGrade").HandlerFunc(o.CommandHandler.Delete)
    return
}


type GroupHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *GroupQueryRepository `json:"queryRepository" eh:"optional"`
}

func NewGroupHttpQueryHandler(queryRepository *GroupQueryRepository) (ret *GroupHttpQueryHandler) {
    httpQueryHandler := eh.NewHttpQueryHandler()
    ret = &GroupHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *GroupHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllGroup", w, r)
}

func (o *GroupHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByGroupId", w, r)
}

func (o *GroupHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllGroup", w, r)
}

func (o *GroupHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByGroupId", w, r)
}

func (o *GroupHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllGroup", w, r)
}

func (o *GroupHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByGroupId", w, r)
}


type GroupHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewGroupHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandHandler) (ret *GroupHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &GroupHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *GroupHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateGroup{Id: id}, w, r)
}

func (o *GroupHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateGroup{Id: id}, w, r)
}

func (o *GroupHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteGroup{Id: id}, w, r)
}


type GroupRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    QueryHandler *GroupHttpQueryHandler `json:"queryHandler" eh:"optional"`
    CommandHandler *GroupHttpCommandHandler `json:"commandHandler" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func NewGroupRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *GroupRouter) {
    pathPrefix = pathPrefix + "/" + "groups"
    entityFactory := func() eventhorizon.Entity { return NewGroup() }
    repo := readRepos(string(GroupAggregateType), entityFactory)
    queryRepository := NewGroupQueryRepository(repo, context)
    queryHandler := NewGroupHttpQueryHandler(queryRepository)
    commandHandler := NewGroupHttpCommandHandler(context, commandBus)
    ret = &GroupRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *GroupRouter) Setup(router *mux.Router) (err error) {
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountGroupById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("CountGroupAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistGroupById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("ExistGroupAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindGroupById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("FindGroupAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateGroup").HandlerFunc(o.CommandHandler.Create)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateGroup").HandlerFunc(o.CommandHandler.Update)
    router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteGroup").HandlerFunc(o.CommandHandler.Delete)
    return
}


type SchoolApplicationHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *SchoolApplicationQueryRepository `json:"queryRepository" eh:"optional"`
}

func NewSchoolApplicationHttpQueryHandler(queryRepository *SchoolApplicationQueryRepository) (ret *SchoolApplicationHttpQueryHandler) {
    httpQueryHandler := eh.NewHttpQueryHandler()
    ret = &SchoolApplicationHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *SchoolApplicationHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllSchoolApplication", w, r)
}

func (o *SchoolApplicationHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindBySchoolApplicationId", w, r)
}

func (o *SchoolApplicationHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllSchoolApplication", w, r)
}

func (o *SchoolApplicationHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountBySchoolApplicationId", w, r)
}

func (o *SchoolApplicationHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllSchoolApplication", w, r)
}

func (o *SchoolApplicationHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistBySchoolApplicationId", w, r)
}


type SchoolApplicationHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewSchoolApplicationHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandHandler) (ret *SchoolApplicationHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &SchoolApplicationHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *SchoolApplicationHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateSchoolApplication{Id: id}, w, r)
}

func (o *SchoolApplicationHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateSchoolApplication{Id: id}, w, r)
}

func (o *SchoolApplicationHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteSchoolApplication{Id: id}, w, r)
}


type SchoolApplicationRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    QueryHandler *SchoolApplicationHttpQueryHandler `json:"queryHandler" eh:"optional"`
    CommandHandler *SchoolApplicationHttpCommandHandler `json:"commandHandler" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func NewSchoolApplicationRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *SchoolApplicationRouter) {
    pathPrefix = pathPrefix + "/" + "schoolApplications"
    entityFactory := func() eventhorizon.Entity { return NewSchoolApplication() }
    repo := readRepos(string(SchoolApplicationAggregateType), entityFactory)
    queryRepository := NewSchoolApplicationQueryRepository(repo, context)
    queryHandler := NewSchoolApplicationHttpQueryHandler(queryRepository)
    commandHandler := NewSchoolApplicationHttpCommandHandler(context, commandBus)
    ret = &SchoolApplicationRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *SchoolApplicationRouter) Setup(router *mux.Router) (err error) {
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountSchoolApplicationById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("CountSchoolApplicationAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistSchoolApplicationById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("ExistSchoolApplicationAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindSchoolApplicationById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("FindSchoolApplicationAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateSchoolApplication").HandlerFunc(o.CommandHandler.Create)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateSchoolApplication").HandlerFunc(o.CommandHandler.Update)
    router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteSchoolApplication").HandlerFunc(o.CommandHandler.Delete)
    return
}


type SchoolYearHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *SchoolYearQueryRepository `json:"queryRepository" eh:"optional"`
}

func NewSchoolYearHttpQueryHandler(queryRepository *SchoolYearQueryRepository) (ret *SchoolYearHttpQueryHandler) {
    httpQueryHandler := eh.NewHttpQueryHandler()
    ret = &SchoolYearHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *SchoolYearHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllSchoolYear", w, r)
}

func (o *SchoolYearHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindBySchoolYearId", w, r)
}

func (o *SchoolYearHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllSchoolYear", w, r)
}

func (o *SchoolYearHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountBySchoolYearId", w, r)
}

func (o *SchoolYearHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllSchoolYear", w, r)
}

func (o *SchoolYearHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistBySchoolYearId", w, r)
}


type SchoolYearHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewSchoolYearHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandHandler) (ret *SchoolYearHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &SchoolYearHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *SchoolYearHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateSchoolYear{Id: id}, w, r)
}

func (o *SchoolYearHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateSchoolYear{Id: id}, w, r)
}

func (o *SchoolYearHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteSchoolYear{Id: id}, w, r)
}


type SchoolYearRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    QueryHandler *SchoolYearHttpQueryHandler `json:"queryHandler" eh:"optional"`
    CommandHandler *SchoolYearHttpCommandHandler `json:"commandHandler" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func NewSchoolYearRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *SchoolYearRouter) {
    pathPrefix = pathPrefix + "/" + "schoolYears"
    entityFactory := func() eventhorizon.Entity { return NewSchoolYear() }
    repo := readRepos(string(SchoolYearAggregateType), entityFactory)
    queryRepository := NewSchoolYearQueryRepository(repo, context)
    queryHandler := NewSchoolYearHttpQueryHandler(queryRepository)
    commandHandler := NewSchoolYearHttpCommandHandler(context, commandBus)
    ret = &SchoolYearRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *SchoolYearRouter) Setup(router *mux.Router) (err error) {
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountSchoolYearById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("CountSchoolYearAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistSchoolYearById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("ExistSchoolYearAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindSchoolYearById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("FindSchoolYearAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateSchoolYear").HandlerFunc(o.CommandHandler.Create)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateSchoolYear").HandlerFunc(o.CommandHandler.Update)
    router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteSchoolYear").HandlerFunc(o.CommandHandler.Delete)
    return
}


type StudentRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    AttendanceRouter *AttendanceRouter `json:"attendanceRouter" eh:"optional"`
    CourseRouter *CourseRouter `json:"courseRouter" eh:"optional"`
    GradeRouter *GradeRouter `json:"gradeRouter" eh:"optional"`
    GroupRouter *GroupRouter `json:"groupRouter" eh:"optional"`
    SchoolApplicationRouter *SchoolApplicationRouter `json:"schoolApplicationRouter" eh:"optional"`
    SchoolYearRouter *SchoolYearRouter `json:"schoolYearRouter" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func NewStudentRouter(pathPrefix string, context context.Context, commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *StudentRouter) {
    pathPrefix = pathPrefix + "/" + "student"
    attendanceRouter := NewAttendanceRouter(pathPrefix, context, commandBus, readRepos)
    courseRouter := NewCourseRouter(pathPrefix, context, commandBus, readRepos)
    gradeRouter := NewGradeRouter(pathPrefix, context, commandBus, readRepos)
    groupRouter := NewGroupRouter(pathPrefix, context, commandBus, readRepos)
    schoolApplicationRouter := NewSchoolApplicationRouter(pathPrefix, context, commandBus, readRepos)
    schoolYearRouter := NewSchoolYearRouter(pathPrefix, context, commandBus, readRepos)
    ret = &StudentRouter{
        PathPrefix: pathPrefix,
        AttendanceRouter: attendanceRouter,
        CourseRouter: courseRouter,
        GradeRouter: gradeRouter,
        GroupRouter: groupRouter,
        SchoolApplicationRouter: schoolApplicationRouter,
        SchoolYearRouter: schoolYearRouter,
    }
    return
}

func (o *StudentRouter) Setup(router *mux.Router) (err error) {
    if err = o.AttendanceRouter.Setup(router); err != nil {
        return
    }
    if err = o.CourseRouter.Setup(router); err != nil {
        return
    }
    if err = o.GradeRouter.Setup(router); err != nil {
        return
    }
    if err = o.GroupRouter.Setup(router); err != nil {
        return
    }
    if err = o.SchoolApplicationRouter.Setup(router); err != nil {
        return
    }
    if err = o.SchoolYearRouter.Setup(router); err != nil {
        return
    }
    return
}









