package finance

import (
    "context"
    "github.com/eugeis/gee/eh"
    "github.com/eugeis/gee/net"
    "github.com/gorilla/mux"
    "github.com/looplab/eventhorizon"
    "github.com/looplab/eventhorizon/commandhandler/bus"
    "net/http"
)
type ExpenseHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *ExpenseQueryRepository `json:"queryRepository" eh:"optional"`
}

func NewExpenseHttpQueryHandler(queryRepository *ExpenseQueryRepository) (ret *ExpenseHttpQueryHandler) {
    httpQueryHandler := eh.NewHttpQueryHandler()
    ret = &ExpenseHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *ExpenseHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllExpense", w, r)
}

func (o *ExpenseHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByExpenseId", w, r)
}

func (o *ExpenseHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllExpense", w, r)
}

func (o *ExpenseHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByExpenseId", w, r)
}

func (o *ExpenseHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllExpense", w, r)
}

func (o *ExpenseHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByExpenseId", w, r)
}


type ExpenseHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewExpenseHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandHandler) (ret *ExpenseHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &ExpenseHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *ExpenseHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateExpense{Id: id}, w, r)
}

func (o *ExpenseHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateExpense{Id: id}, w, r)
}

func (o *ExpenseHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteExpense{Id: id}, w, r)
}


type ExpenseRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    QueryHandler *ExpenseHttpQueryHandler `json:"queryHandler" eh:"optional"`
    CommandHandler *ExpenseHttpCommandHandler `json:"commandHandler" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func NewExpenseRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *ExpenseRouter) {
    pathPrefix = pathPrefix + "/" + "expenses"
    entityFactory := func() eventhorizon.Entity { return NewExpense() }
    repo := readRepos(string(ExpenseAggregateType), entityFactory)
    queryRepository := NewExpenseQueryRepository(repo, context)
    queryHandler := NewExpenseHttpQueryHandler(queryRepository)
    commandHandler := NewExpenseHttpCommandHandler(context, commandBus)
    ret = &ExpenseRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *ExpenseRouter) Setup(router *mux.Router) (err error) {
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountExpenseById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("CountExpenseAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistExpenseById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("ExistExpenseAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindExpenseById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("FindExpenseAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateExpense").HandlerFunc(o.CommandHandler.Create)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateExpense").HandlerFunc(o.CommandHandler.Update)
    router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteExpense").HandlerFunc(o.CommandHandler.Delete)
    return
}


type ExpensePurposeHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *ExpensePurposeQueryRepository `json:"queryRepository" eh:"optional"`
}

func NewExpensePurposeHttpQueryHandler(queryRepository *ExpensePurposeQueryRepository) (ret *ExpensePurposeHttpQueryHandler) {
    httpQueryHandler := eh.NewHttpQueryHandler()
    ret = &ExpensePurposeHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *ExpensePurposeHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllExpensePurpose", w, r)
}

func (o *ExpensePurposeHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByExpensePurposeId", w, r)
}

func (o *ExpensePurposeHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllExpensePurpose", w, r)
}

func (o *ExpensePurposeHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByExpensePurposeId", w, r)
}

func (o *ExpensePurposeHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllExpensePurpose", w, r)
}

func (o *ExpensePurposeHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByExpensePurposeId", w, r)
}


type ExpensePurposeHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewExpensePurposeHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandHandler) (ret *ExpensePurposeHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &ExpensePurposeHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *ExpensePurposeHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateExpensePurpose{Id: id}, w, r)
}

func (o *ExpensePurposeHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateExpensePurpose{Id: id}, w, r)
}

func (o *ExpensePurposeHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteExpensePurpose{Id: id}, w, r)
}


type ExpensePurposeRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    QueryHandler *ExpensePurposeHttpQueryHandler `json:"queryHandler" eh:"optional"`
    CommandHandler *ExpensePurposeHttpCommandHandler `json:"commandHandler" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func NewExpensePurposeRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *ExpensePurposeRouter) {
    pathPrefix = pathPrefix + "/" + "expensePurposes"
    entityFactory := func() eventhorizon.Entity { return NewExpensePurpose() }
    repo := readRepos(string(ExpensePurposeAggregateType), entityFactory)
    queryRepository := NewExpensePurposeQueryRepository(repo, context)
    queryHandler := NewExpensePurposeHttpQueryHandler(queryRepository)
    commandHandler := NewExpensePurposeHttpCommandHandler(context, commandBus)
    ret = &ExpensePurposeRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *ExpensePurposeRouter) Setup(router *mux.Router) (err error) {
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountExpensePurposeById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("CountExpensePurposeAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistExpensePurposeById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("ExistExpensePurposeAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindExpensePurposeById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("FindExpensePurposeAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateExpensePurpose").HandlerFunc(o.CommandHandler.Create)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateExpensePurpose").HandlerFunc(o.CommandHandler.Update)
    router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteExpensePurpose").HandlerFunc(o.CommandHandler.Delete)
    return
}


type FeeHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *FeeQueryRepository `json:"queryRepository" eh:"optional"`
}

func NewFeeHttpQueryHandler(queryRepository *FeeQueryRepository) (ret *FeeHttpQueryHandler) {
    httpQueryHandler := eh.NewHttpQueryHandler()
    ret = &FeeHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *FeeHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllFee", w, r)
}

func (o *FeeHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByFeeId", w, r)
}

func (o *FeeHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllFee", w, r)
}

func (o *FeeHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByFeeId", w, r)
}

func (o *FeeHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllFee", w, r)
}

func (o *FeeHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByFeeId", w, r)
}


type FeeHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewFeeHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandHandler) (ret *FeeHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &FeeHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *FeeHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateFee{Id: id}, w, r)
}

func (o *FeeHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateFee{Id: id}, w, r)
}

func (o *FeeHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteFee{Id: id}, w, r)
}


type FeeRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    QueryHandler *FeeHttpQueryHandler `json:"queryHandler" eh:"optional"`
    CommandHandler *FeeHttpCommandHandler `json:"commandHandler" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func NewFeeRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *FeeRouter) {
    pathPrefix = pathPrefix + "/" + "fees"
    entityFactory := func() eventhorizon.Entity { return NewFee() }
    repo := readRepos(string(FeeAggregateType), entityFactory)
    queryRepository := NewFeeQueryRepository(repo, context)
    queryHandler := NewFeeHttpQueryHandler(queryRepository)
    commandHandler := NewFeeHttpCommandHandler(context, commandBus)
    ret = &FeeRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *FeeRouter) Setup(router *mux.Router) (err error) {
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountFeeById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("CountFeeAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistFeeById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("ExistFeeAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindFeeById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("FindFeeAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateFee").HandlerFunc(o.CommandHandler.Create)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateFee").HandlerFunc(o.CommandHandler.Update)
    router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteFee").HandlerFunc(o.CommandHandler.Delete)
    return
}


type FeeKindHttpQueryHandler struct {
    *eh.HttpQueryHandler
    QueryRepository *FeeKindQueryRepository `json:"queryRepository" eh:"optional"`
}

func NewFeeKindHttpQueryHandler(queryRepository *FeeKindQueryRepository) (ret *FeeKindHttpQueryHandler) {
    httpQueryHandler := eh.NewHttpQueryHandler()
    ret = &FeeKindHttpQueryHandler{
        HttpQueryHandler: httpQueryHandler,
        QueryRepository: queryRepository,
    }
    return
}

func (o *FeeKindHttpQueryHandler) FindAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.FindAll()
    o.HandleResult(ret, err, "FindAllFeeKind", w, r)
}

func (o *FeeKindHttpQueryHandler) FindById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.FindById(id)
    o.HandleResult(ret, err, "FindByFeeKindId", w, r)
}

func (o *FeeKindHttpQueryHandler) CountAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.CountAll()
    o.HandleResult(ret, err, "CountAllFeeKind", w, r)
}

func (o *FeeKindHttpQueryHandler) CountById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.CountById(id)
    o.HandleResult(ret, err, "CountByFeeKindId", w, r)
}

func (o *FeeKindHttpQueryHandler) ExistAll(w http.ResponseWriter, r *http.Request) {
    ret, err := o.QueryRepository.ExistAll()
    o.HandleResult(ret, err, "ExistAllFeeKind", w, r)
}

func (o *FeeKindHttpQueryHandler) ExistById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    ret, err := o.QueryRepository.ExistById(id)
    o.HandleResult(ret, err, "ExistByFeeKindId", w, r)
}


type FeeKindHttpCommandHandler struct {
    *eh.HttpCommandHandler
}

func NewFeeKindHttpCommandHandler(context context.Context, commandBus eventhorizon.CommandHandler) (ret *FeeKindHttpCommandHandler) {
    httpCommandHandler := eh.NewHttpCommandHandler(context, commandBus)
    ret = &FeeKindHttpCommandHandler{
        HttpCommandHandler: httpCommandHandler,
    }
    return
}

func (o *FeeKindHttpCommandHandler) Create(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&CreateFeeKind{Id: id}, w, r)
}

func (o *FeeKindHttpCommandHandler) Update(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&UpdateFeeKind{Id: id}, w, r)
}

func (o *FeeKindHttpCommandHandler) Delete(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := eventhorizon.UUID(vars["id"])
    o.HandleCommand(&DeleteFeeKind{Id: id}, w, r)
}


type FeeKindRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    QueryHandler *FeeKindHttpQueryHandler `json:"queryHandler" eh:"optional"`
    CommandHandler *FeeKindHttpCommandHandler `json:"commandHandler" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func NewFeeKindRouter(pathPrefix string, context context.Context, commandBus eventhorizon.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *FeeKindRouter) {
    pathPrefix = pathPrefix + "/" + "feeKinds"
    entityFactory := func() eventhorizon.Entity { return NewFeeKind() }
    repo := readRepos(string(FeeKindAggregateType), entityFactory)
    queryRepository := NewFeeKindQueryRepository(repo, context)
    queryHandler := NewFeeKindHttpQueryHandler(queryRepository)
    commandHandler := NewFeeKindHttpCommandHandler(context, commandBus)
    ret = &FeeKindRouter{
        PathPrefix: pathPrefix,
        QueryHandler: queryHandler,
        CommandHandler: commandHandler,
    }
    return
}

func (o *FeeKindRouter) Setup(router *mux.Router) (err error) {
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CountFeeKindById").HandlerFunc(o.QueryHandler.CountById).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("CountFeeKindAll").HandlerFunc(o.QueryHandler.CountAll).
        Queries(net.QueryType, net.QueryTypeCount)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("ExistFeeKindById").HandlerFunc(o.QueryHandler.ExistById).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("ExistFeeKindAll").HandlerFunc(o.QueryHandler.ExistAll).
        Queries(net.QueryType, net.QueryTypeExist)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("FindFeeKindById").HandlerFunc(o.QueryHandler.FindById)
    router.Methods(http.MethodGet).PathPrefix(o.PathPrefix).
        Name("FindFeeKindAll").HandlerFunc(o.QueryHandler.FindAll)
    router.Methods(http.MethodPost).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("CreateFeeKind").HandlerFunc(o.CommandHandler.Create)
    router.Methods(http.MethodPut).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("UpdateFeeKind").HandlerFunc(o.CommandHandler.Update)
    router.Methods(http.MethodDelete).PathPrefix(o.PathPrefix).Path("/{id}").
        Name("DeleteFeeKind").HandlerFunc(o.CommandHandler.Delete)
    return
}


type FinanceRouter struct {
    PathPrefix string `json:"pathPrefix" eh:"optional"`
    ExpenseRouter *ExpenseRouter `json:"expenseRouter" eh:"optional"`
    ExpensePurposeRouter *ExpensePurposeRouter `json:"expensePurposeRouter" eh:"optional"`
    FeeRouter *FeeRouter `json:"feeRouter" eh:"optional"`
    FeeKindRouter *FeeKindRouter `json:"feeKindRouter" eh:"optional"`
    Router *mux.Router `json:"router" eh:"optional"`
}

func NewFinanceRouter(pathPrefix string, context context.Context, commandBus *bus.CommandHandler, 
                readRepos func (string, func () (ret eventhorizon.Entity) ) (ret eventhorizon.ReadWriteRepo) ) (ret *FinanceRouter) {
    pathPrefix = pathPrefix + "/" + "finance"
    expenseRouter := NewExpenseRouter(pathPrefix, context, commandBus, readRepos)
    expensePurposeRouter := NewExpensePurposeRouter(pathPrefix, context, commandBus, readRepos)
    feeRouter := NewFeeRouter(pathPrefix, context, commandBus, readRepos)
    feeKindRouter := NewFeeKindRouter(pathPrefix, context, commandBus, readRepos)
    ret = &FinanceRouter{
        PathPrefix: pathPrefix,
        ExpenseRouter: expenseRouter,
        ExpensePurposeRouter: expensePurposeRouter,
        FeeRouter: feeRouter,
        FeeKindRouter: feeKindRouter,
    }
    return
}

func (o *FinanceRouter) Setup(router *mux.Router) (err error) {
    if err = o.ExpenseRouter.Setup(router); err != nil {
        return
    }
    if err = o.ExpensePurposeRouter.Setup(router); err != nil {
        return
    }
    if err = o.FeeRouter.Setup(router); err != nil {
        return
    }
    if err = o.FeeKindRouter.Setup(router); err != nil {
        return
    }
    return
}









