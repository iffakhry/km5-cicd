package router

import (
	"fakhry/clean-arch/features/user/handler"
	"fakhry/clean-arch/features/user/repository"
	"fakhry/clean-arch/features/user/usecase"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	userRepository := repository.New(db)
	userUsecase := usecase.New(userRepository)
	userController := handler.New(userUsecase)

	e.POST("/users", userController.CreateUser)
	e.GET("/hello", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "hello world",
		})
	})
}
