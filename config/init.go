package config

import (
	"gin-gonic-api/app/controller"
	"gin-gonic-api/app/repository"
	"gin-gonic-api/app/service"
)

type Initialization struct {
	userRepo            repository.UserRepository
	userSvc             service.UserService
	UserCtrl            controller.UserController
	RoleRepo            repository.RoleRepository
	n2VocabularyRepo    repository.N2VocabularyRepository
	N2VocabularyCtrl    controller.N2VocabularyController
	n2VocabularyService service.N2VocabularyService
	operationRepo       repository.OperationRepository
	OperationCtrl       controller.OperationController
	operationSerivce    service.OperationService
	authRepo            repository.AuthRepository
	authSvc             service.AuthService
	AuthCtrl            controller.AuthController
	typeRepo            repository.TypeRepository
	typeSvc             service.TypeService
	TypeCtrl            controller.TypeContoller
}

func NewInitialization(
	userRepo repository.UserRepository,
	userService service.UserService,
	userCtrl controller.UserController,
	roleRepo repository.RoleRepository,
	n2VocabularyRepo repository.N2VocabularyRepository,
	n2VocabularyService service.N2VocabularyService,
	n2VocabularyCtrl controller.N2VocabularyController,

	operationRepo repository.OperationRepository,
	operationCtrl controller.OperationController,
	operationSerivce service.OperationService,

	authRepo repository.AuthRepository,
	authSvc service.AuthService,
	authCtrl controller.AuthController,

	typeRepo repository.TypeRepository,
	typeSvc service.TypeService,
	typeCtrl controller.TypeContoller,

) *Initialization {
	return &Initialization{
		userRepo:            userRepo,
		userSvc:             userService,
		UserCtrl:            userCtrl,
		RoleRepo:            roleRepo,
		n2VocabularyRepo:    n2VocabularyRepo,
		n2VocabularyService: n2VocabularyService,
		N2VocabularyCtrl:    n2VocabularyCtrl,
		operationRepo:       operationRepo,
		OperationCtrl:       operationCtrl,
		operationSerivce:    operationSerivce,
		authRepo:            authRepo,
		authSvc:             authSvc,
		AuthCtrl:            authCtrl,
		typeRepo:            typeRepo,
		typeSvc:             typeSvc,
		TypeCtrl:            typeCtrl,
	}
}
