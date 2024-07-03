// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package config

import (
	"gin-gonic-api/app/controller"
	"gin-gonic-api/app/firebase"
	"gin-gonic-api/app/pkg"
	"gin-gonic-api/app/repository"
	"gin-gonic-api/app/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// Injectors from injector.go:

func Init(ctx *gin.Context) *Initialization {
	gormDB := ConnectToDB()
	userRepositoryImpl := repository.UserRepositoryInit(gormDB)
	userServiceImpl := service.UserServiceInit(userRepositoryImpl)
	userControllerImpl := controller.UserControllerInit(userServiceImpl)
	app := firebase.InitFirebase()
	authRepositoryImpl := repository.AuthRepositoryInit(gormDB, app)
	authServiceImpl := service.AuthServiceInit(authRepositoryImpl, app)
	authControllerImpl := controller.AuthControllerInit(authServiceImpl)
	typeRepositoryImpl := repository.TypeRepositoryInit(gormDB)
	typeServiceImpl := service.TypeServiceInit(typeRepositoryImpl)
	typeControllerImpl := controller.TypeControllerInit(typeServiceImpl)
	itemRepositoryImpl := repository.ItemRepositoryInit(gormDB)
	answerRepositoryImpl := repository.AnswerRepositoryInit(gormDB)
	recordRepositoryImpl := repository.RecordRepositoryInit(gormDB)
	itemServiceImpl := service.ItemServiceInit(itemRepositoryImpl, answerRepositoryImpl, recordRepositoryImpl)
	itemControllerImpl := controller.ItemControllerInit(itemServiceImpl)
	answerServiceImpl := service.AnswerServiceInit(answerRepositoryImpl)
	answerControllerImpl := controller.AnswerControllerInit(answerServiceImpl)
	recordServiceImpl := service.RecordServiceInit(recordRepositoryImpl)
	recordControllerImpl := controller.RecordControllerInit(recordServiceImpl)
	initialization := NewInitialization(userRepositoryImpl, userServiceImpl, userControllerImpl, authRepositoryImpl, authServiceImpl, authControllerImpl, typeRepositoryImpl, typeServiceImpl, typeControllerImpl, itemRepositoryImpl, itemServiceImpl, itemControllerImpl, answerRepositoryImpl, answerServiceImpl, answerControllerImpl, recordRepositoryImpl, recordServiceImpl, recordControllerImpl)
	return initialization
}

// injector.go:

var db = wire.NewSet(ConnectToDB)

var paginatorSet = wire.NewSet(pkg.PaginatorHandler)

var firebaseSet = wire.NewSet(firebase.InitFirebase)

var userServiceSet = wire.NewSet(service.UserServiceInit, wire.Bind(new(service.UserService), new(*service.UserServiceImpl)))

var userRepoSet = wire.NewSet(repository.UserRepositoryInit, wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)))

var userCtrlSet = wire.NewSet(controller.UserControllerInit, wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)))


var itemRepoSet = wire.NewSet(repository.ItemRepositoryInit, wire.Bind(new(repository.ItemRepository), new(*repository.ItemRepositoryImpl)))

var itemServiceSet = wire.NewSet(service.ItemServiceInit, wire.Bind(new(service.ItemService), new(*service.ItemServiceImpl)))

var itemCtrlSet = wire.NewSet(controller.ItemControllerInit, wire.Bind(new(controller.ItemController), new(*controller.ItemControllerImpl)))


var authRepoSet = wire.NewSet(repository.AuthRepositoryInit, wire.Bind(new(repository.AuthRepository), new(*repository.AuthRepositoryImpl)))

var authServiceSet = wire.NewSet(service.AuthServiceInit, wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)))

var authCtrlSet = wire.NewSet(controller.AuthControllerInit, wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)))

var typeRepoSet = wire.NewSet(repository.TypeRepositoryInit, wire.Bind(new(repository.TypeRepository), new(*repository.TypeRepositoryImpl)))

var typeServiceSet = wire.NewSet(service.TypeServiceInit, wire.Bind(new(service.TypeService), new(*service.TypeServiceImpl)))

var typeCtrlSet = wire.NewSet(controller.TypeControllerInit, wire.Bind(new(controller.TypeController), new(*controller.TypeControllerImpl)))

var answerRepoSet = wire.NewSet(repository.AnswerRepositoryInit, wire.Bind(new(repository.AnswerRepository), new(*repository.AnswerRepositoryImpl)))

var answerServiceSet = wire.NewSet(service.AnswerServiceInit, wire.Bind(new(service.AnswerService), new(*service.AnswerServiceImpl)))

var answerCtrlSet = wire.NewSet(controller.AnswerControllerInit, wire.Bind(new(controller.AnswerController), new(*controller.AnswerControllerImpl)))

var recordRepoSet = wire.NewSet(repository.RecordRepositoryInit, wire.Bind(new(repository.RecordRepository), new(*repository.RecordRepositoryImpl)))

var recordServiceSet = wire.NewSet(service.RecordServiceInit, wire.Bind(new(service.RecordService), new(*service.RecordServiceImpl)))

var recordCtrlSet = wire.NewSet(controller.RecordControllerInit, wire.Bind(new(controller.RecordController), new(*controller.RecordControllerImpl)))
