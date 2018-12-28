package auth

import (
    "context"
    "github.com/eugeis/gee/eh"
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "github.com/looplab/eventhorizon"
    "github.com/looplab/eventhorizon/commandhandler/bus"
    "net/http"
)
type AccountHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *AccountQueryRepository `json:"queryRepository" eh:"optional"`
}

func NewAccountHttpQueryHandler(queryRepository *AccountQueryRepository) (ret *AccountHttpQueryHandler) {
    httpQueryHandler := eh.NewHttpQueryHandler()
    ret = &AccountHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *AccountHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllAccount", w, r)
}

func (o *AccountHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByAccountId", w, r)
}

func (o *AccountHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllAccount", w, r)
}

func (o *AccountHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByAccountId", w, r)
}

func (o *AccountHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllAccount", w, r)
}

func (o *AccountHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByAccountId", w, r)
}


type AccountHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewAccountHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandHandler) (ret *AccountHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &AccountHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *AccountHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateAccount{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) Enable(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&EnableAccount{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) Disable(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DisableAccount{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateAccount{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteAccount{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) SendEnabledConfirmation(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&SendAccountEnabledConfirmation{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) SendDisabledConfirmation(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&SendAccountDisabledConfirmation{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) Login(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&LoginAccount{Id: id}, w, r)
}

func (o *AccountHttpCommandHandler) SendCreatedConfirmation(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&SendAccountCreatedConfirmation{Id: id}, w, r)
}


type AccountRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    QueryHandler *AccountHttpQueryHandler `json:"queryHandler" eh:"optional"`
    CommandHandler *AccountHttpCommandHandler `json:"commandHandler" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func NewAccountRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AccountRouter) {
    pathPrefix = pathPrefix + "/" + "accounts"
    entityFactory := func() eventhorizon.Entity { return NewAccount() }
    repo := readRepos(string(AccountAggregateType), entityFactory)
    queryRepository := NewAccountQueryRepository(repo, context)
    queryHandler := NewAccountHttpQueryHandler(queryRepository)
    commandHandler := NewAccountHttpCommandHandler(context, commandBus)
    ret = &AccountRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *AccountRouter) Setup(router *mux.Router) (err error) {
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountAccountById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("CountAccountAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistAccountById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("ExistAccountAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindAccountById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("FindAccountAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Queries(net.Command, "login").
        Name("LoginAccount").HandlerFunc(o.CommandHandler.Login)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Queries(net.Command, "sendCreatedConfirmation").
        Name("SendAccountCreatedConfirmation").HandlerFunc(o.CommandHandler.SendCreatedConfirmation)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Queries(net.Command, "sendEnabledConfirmation").
        Name("SendAccountEnabledConfirmation").HandlerFunc(o.CommandHandler.SendEnabledConfirmation)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Queries(net.Command, "sendDisabledConfirmation").
        Name("SendAccountDisabledConfirmation").HandlerFunc(o.CommandHandler.SendDisabledConfirmation)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateAccount").HandlerFunc(o.CommandHandler.Create)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Queries(net.Command, "enable").
        Name("EnableAccount").HandlerFunc(o.CommandHandler.Enable)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Queries(net.Command, "disable").
        Name("DisableAccount").HandlerFunc(o.CommandHandler.Disable)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateAccount").HandlerFunc(o.CommandHandler.Update)
    router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteAccount").HandlerFunc(o.CommandHandler.Delete)
    return
}


type AuthRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    AccountRouter *AccountRouter `json:"accountRouter" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func NewAuthRouter(pathPrefix string, context context.Context, commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *AuthRouter) {
    pathPrefix = pathPrefix + "/" + "auth"
    accountRouter := NewAccountRouter(pathPrefix, context, commandBus, readRepos)
    ret = &AuthRouter{
        PathPrefix: pathPrefix,
        AccountRouter: accountRouter,
    }
    return
}

func (o *AuthRouter) Setup(router *mux.Router) (err error) {
    if err = o.AccountRouter.Setup(router); err != nil {
        return
    }
    return
}









