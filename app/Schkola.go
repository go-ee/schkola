package app

import (
	"ee/schkola/auth"
	"ee/schkola/finance"
	"ee/schkola/library"
	"ee/schkola/person"
	"ee/schkola/student"
	"errors"

	"github.com/eugeis/gee/crypt"
	"github.com/eugeis/gee/eh/app"
	"github.com/eugeis/gee/net"
	"github.com/looplab/eventhorizon"
)

type Schkola struct {
	*app.AppBase
}

func NewSchkola(appBase *app.AppBase) *Schkola {
	appBase.ProductName = "Schkola"
	return &Schkola{AppBase: appBase}
}

func (o *Schkola) Start() {

	authEngine := auth.NewAuthEventhorizonInitializer(o.EventStore, o.EventBus, o.CommandBus, o.ReadRepos)
	authEngine.Setup()
	authEngine.ActivatePasswordEncryption()

	financeEngine := finance.NewFinanceEventhorizonInitializer(o.EventStore, o.EventBus, o.CommandBus, o.ReadRepos)
	financeEngine.Setup()

	libraryEngine := library.NewLibraryEventhorizonInitializer(o.EventStore, o.EventBus, o.CommandBus, o.ReadRepos)
	libraryEngine.Setup()

	personEngine := person.NewPersonEventhorizonInitializer(o.EventStore, o.EventBus, o.CommandBus, o.ReadRepos)
	personEngine.Setup()

	studentEngine := student.NewStudentEventhorizonInitializer(o.EventStore, o.EventBus, o.CommandBus, o.ReadRepos)
	studentEngine.Setup()

	authRouter := auth.NewAuthRouter("", o.Ctx, o.CommandBus, o.ReadRepos)
	authRouter.Setup(o.Router)

	financeRouter := finance.NewFinanceRouter("", o.Ctx, o.CommandBus, o.ReadRepos)
	financeRouter.Setup(o.Router)

	libraryRouter := library.NewLibraryRouter("", o.Ctx, o.CommandBus, o.ReadRepos)
	libraryRouter.Setup(o.Router)

	personRouter := person.NewPersonRouter("", o.Ctx, o.CommandBus, o.ReadRepos)
	personRouter.Setup(o.Router)

	studentRouter := student.NewStudentRouter("", o.Ctx, o.CommandBus, o.ReadRepos)
	studentRouter.Setup(o.Router)

	if o.Secure {
		o.Jwt = o.initJwtController(authRouter.AccountRouter.QueryHandler.QueryRepository)
	}

	o.StartServer()
}

func (o *Schkola) initJwtController(accounts *auth.AccountQueryRepository) (ret *net.JwtController) {
	//TODO use appName, provide help how to generate RSA files first
	return net.NewJwtControllerApp("app",
		func(credentials net.UserCredentials) (ret interface{}, err error) {
			var account *auth.Account
			if account, err = accounts.FindById(eventhorizon.UUID(credentials.Username)); err == nil {
				if !crypt.HashAndEquals(credentials.Password, account.Password) {
					err = errors.New("password mismatch")
				} else {
					ret = account
				}
			}
			return
		})
}
