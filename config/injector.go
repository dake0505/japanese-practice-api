//go:build wireinject
// +build wireinject

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

var db = wire.NewSet(ConnectToDB)

var paginatorSet = wire.NewSet(pkg.PaginatorHandler)

var firebaseSet = wire.NewSet(firebase.InitFirebase)

var userServiceSet = wire.NewSet(service.UserServiceInit,
	wire.Bind(new(service.UserService), new(*service.UserServiceImpl)),
)

var userRepoSet = wire.NewSet(repository.UserRepositoryInit,
	wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
)

var userCtrlSet = wire.NewSet(controller.UserControllerInit,
	wire.Bind(new(controller.UserController), new(*controller.UserControllerImpl)),
)

var roleRepoSet = wire.NewSet(repository.RoleRepositoryInit,
	wire.Bind(new(repository.RoleRepository), new(*repository.RoleRepositoryImpl)),
)

var n2VocabularyRepoSet = wire.NewSet(repository.N2VocabularyRepositoryInit,
	wire.Bind(new(repository.N2VocabularyRepository), new(*repository.N2VocabularyRepositoryImpl)),
)

var n2VocabularyServiceSet = wire.NewSet(service.N2VocabularyServiceInit,
	wire.Bind(new(service.N2VocabularyService), new(*service.N2VocabularyServiceImpl)),
)

var n2VocabularyCtrlSet = wire.NewSet(controller.N2VocabularyControllerInit,
	wire.Bind(new(controller.N2VocabularyController), new(*controller.N2VocabularyControllerImpl)),
)

var itemRepoSet = wire.NewSet(repository.ItemRepositoryInit,
	wire.Bind(new(repository.ItemRepository), new(*repository.ItemRepositoryImpl)),
)

var itemServiceSet = wire.NewSet(service.ItemServiceInit,
	wire.Bind(new(service.ItemService), new(*service.ItemServiceImpl)),
)

var itemCtrlSet = wire.NewSet(controller.ItemControllerInit,
	wire.Bind(new(controller.ItemController), new(*controller.ItemControllerImpl)),
)

var operationRepoSet = wire.NewSet(repository.OperationRepositoryInit,
	wire.Bind(new(repository.OperationRepository), new(*repository.OperationRepositoryImpl)),
)

var operationServiceSet = wire.NewSet(service.OperationServiceInit,
	wire.Bind(new(service.OperationService), new(*service.OperationServiceImpl)),
)

var operationCtrlSet = wire.NewSet(controller.OperationControllerInit,
	wire.Bind(new(controller.OperationController), new(*controller.OperationControllerImpl)),
)

var authRepoSet = wire.NewSet(repository.AuthRepositoryInit,
	wire.Bind(new(repository.AuthRepository), new(*repository.AuthRepositoryImpl)),
)

var authServiceSet = wire.NewSet(service.AuthServiceInit,
	wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)),
)

var authCtrlSet = wire.NewSet(controller.AuthControllerInit,
	wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)),
)

var typeRepoSet = wire.NewSet(repository.TypeRepositoryInit,
	wire.Bind(new(repository.TypeRepository), new(*repository.TypeRepositoryImpl)),
)

var typeServiceSet = wire.NewSet(service.TypeServiceInit,
	wire.Bind(new(service.TypeService), new(*service.TypeServiceImpl)),
)

var typeCtrlSet = wire.NewSet(controller.TypeControllerInit,
	wire.Bind(new(controller.TypeController), new(*controller.TypeControllerImpl)),
)

var answerRepoSet = wire.NewSet(repository.AnswerRepositoryInit,
	wire.Bind(new(repository.AnswerRepository), new(*repository.AnswerRepositoryImpl)),
)

var answerServiceSet = wire.NewSet(service.AnswerServiceInit,
	wire.Bind(new(service.AnswerService), new(*service.AnswerServiceImpl)),
)

var answerCtrlSet = wire.NewSet(controller.AnswerControllerInit,
	wire.Bind(new(controller.AnswerController), new(*controller.AnswerControllerImpl)),
)

func Init(ctx *gin.Context) *Initialization {
	wire.Build(NewInitialization, db, paginatorSet, firebaseSet, userCtrlSet, userServiceSet, userRepoSet, roleRepoSet, itemRepoSet, itemCtrlSet, itemServiceSet, n2VocabularyRepoSet, n2VocabularyServiceSet, n2VocabularyCtrlSet, operationRepoSet, operationServiceSet, operationCtrlSet, authRepoSet, authServiceSet, authCtrlSet, typeRepoSet, typeServiceSet, typeCtrlSet, answerServiceSet, answerRepoSet, answerCtrlSet)
	return nil
}
