package config

import (
	"gin-gonic-api/app/controller"
	"gin-gonic-api/app/repository"
	"gin-gonic-api/app/service"
)

type Initialization struct {
	userRepo repository.UserRepository
	userSvc  service.UserService
	UserCtrl controller.UserController

	authRepo repository.AuthRepository
	authSvc  service.AuthService
	AuthCtrl controller.AuthController

	typeRepo repository.TypeRepository
	typeSvc  service.TypeService
	TypeCtrl controller.TypeController

	itemRepo repository.ItemRepository
	itemSvc  service.ItemService
	ItemCtrl controller.ItemController

	answerRepo repository.AnswerRepository
	answerSvc  service.AnswerService
	AnswerCtrl controller.AnswerController

	recordRepo repository.RecordRepository
	recordSvc  service.RecordService
	RecordCtrl controller.RecordController
}

func NewInitialization(
	userRepo repository.UserRepository,
	userService service.UserService,
	userCtrl controller.UserController,

	authRepo repository.AuthRepository,
	authSvc service.AuthService,
	authCtrl controller.AuthController,

	typeRepo repository.TypeRepository,
	typeSvc service.TypeService,
	typeCtrl controller.TypeController,

	itemRepo repository.ItemRepository,
	itemSvc service.ItemService,
	itemCtrl controller.ItemController,

	answerRepo repository.AnswerRepository,
	answerSvc service.AnswerService,
	answerCtrl controller.AnswerController,

	recordRepo repository.RecordRepository,
	recordSvc service.RecordService,
	recordCtrl controller.RecordController,
) *Initialization {
	return &Initialization{
		userRepo:   userRepo,
		userSvc:    userService,
		UserCtrl:   userCtrl,
		authRepo:   authRepo,
		authSvc:    authSvc,
		AuthCtrl:   authCtrl,
		typeRepo:   typeRepo,
		typeSvc:    typeSvc,
		TypeCtrl:   typeCtrl,
		itemRepo:   itemRepo,
		itemSvc:    itemSvc,
		ItemCtrl:   itemCtrl,
		answerRepo: answerRepo,
		answerSvc:  answerSvc,
		AnswerCtrl: answerCtrl,
		recordRepo: recordRepo,
		recordSvc:  recordSvc,
		RecordCtrl: recordCtrl,
	}
}
