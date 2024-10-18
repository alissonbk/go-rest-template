package router

import (
	"com.github.alissonbk/go-rest-template/app/controller"
	"com.github.alissonbk/go-rest-template/app/repository"
	"com.github.alissonbk/go-rest-template/app/service"
	"com.github.alissonbk/go-rest-template/config"
	"github.com/jmoiron/sqlx"
)

// Injection is responsible for dependency injection for each route by returning a `Controller Object` ready to be used by the router

type Injection struct {
	db *sqlx.DB
}

func NewInjection() *Injection {
	return &Injection{db: config.ConnectDB()}
}

func (i *Injection) NewUserController() *controller.UserController {
	r := repository.NewUserRepository(i.db)
	s := service.NewUserService(r)
	return controller.NewUserController(s)
}
