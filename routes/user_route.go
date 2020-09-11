package routes

import (
	"github.com/labstack/echo"
	controller "github.com/williamchang80/sea-apd/controller/http/user"
	domain "github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	"github.com/williamchang80/sea-apd/repository/postgres/user"
	usecase "github.com/williamchang80/sea-apd/usecase/user"
)

type UserRoute struct {
	controller domain.UserController
	usecase    domain.UserUsecase
	Repository domain.UserRepository
}

func NewUserRoute(e *echo.Echo) UserRoute {
	db := db.Postgres()
	authRoute := NewAuthRoute(e)
	repository := user.NewUserRepository(db)
	u := usecase.NewUserUsecase(repository, authRoute.usecase)
	controller := controller.NewUserController(e, u)
	if db != nil {
		db.AutoMigrate(&domain.User{})
	}

	return UserRoute{
		controller: controller,
		usecase:    u,
		Repository: repository,
	}
}
