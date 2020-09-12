package routes

import (
	"github.com/labstack/echo"
	admin2 "github.com/williamchang80/sea-apd/controller/http/admin"
	"github.com/williamchang80/sea-apd/domain/admin"
	domain "github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	repo "github.com/williamchang80/sea-apd/repository/postgres/user"
	user "github.com/williamchang80/sea-apd/usecase/admin"
)

type AdminRoute struct {
	controller admin.AdminController
	usecase    admin.AdminUsecase
	Repository domain.UserRepository
}

func NewAdminRoutes(e *echo.Echo) AdminRoute {
	db := db.Postgres()
	authRoute := NewAuthRoute(e)
	repo := repo.NewUserRepository(db)
	usecase := user.NewAdminUseCase(repo, authRoute.usecase)
	controller := admin2.NewAdminController(e, usecase)
	if db != nil {
		db.AutoMigrate(&domain.User{})
	}

	return AdminRoute{
		controller: controller,
		usecase:    usecase,
		Repository: repo,
	}
}
