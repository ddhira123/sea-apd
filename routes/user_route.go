package routes

import (
	"github.com/labstack/echo"
	user_ctl "github.com/williamchang80/sea-apd/controller/http/user"
	domain "github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	user_repo "github.com/williamchang80/sea-apd/repository/postgres/user"
	use_case "github.com/williamchang80/sea-apd/usecase/user"
)

type UserRoute struct {
	controller domain.UserController
	usecase    domain.UserUsecase
	repository domain.UserRepository
}

func NewUserRoute(e *echo.Echo) UserRoute {
	db := db.Postgres()
	if db != nil {
		db.AutoMigrate(&domain.User{})
	}
	repo := user_repo.NewUserRepository(db)
	usecase := use_case.NewUserUsecase(repo)
	controller := user_ctl.NewUserController(e, usecase)
	return UserRoute{
		controller: controller,
		usecase:    usecase,
		repository: repo,
	}
}
