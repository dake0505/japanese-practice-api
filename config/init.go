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
}

func NewInitialization(
	userRepo repository.UserRepository,
	userService service.UserService,
	userCtrl controller.UserController,
	roleRepo repository.RoleRepository,
	n2VocabularyRepo repository.N2VocabularyRepository,
	n2VocabularyService service.N2VocabularyService,
	n2VocabularyCtrl controller.N2VocabularyController,
) *Initialization {
	return &Initialization{
		userRepo:            userRepo,
		userSvc:             userService,
		UserCtrl:            userCtrl,
		RoleRepo:            roleRepo,
		n2VocabularyRepo:    n2VocabularyRepo,
		n2VocabularyService: n2VocabularyService,
		N2VocabularyCtrl:    n2VocabularyCtrl,
	}
}
