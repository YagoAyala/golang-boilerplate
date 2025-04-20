package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yagoayala/server/internal/controllers"
	"github.com/yagoayala/server/internal/repository"
	"github.com/yagoayala/server/internal/services"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	userRepo := repository.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("", userController.CreateUser)
		userRoutes.GET("", userController.GetAllUsers)
		userRoutes.GET("/:id", userController.GetUserByID)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}
}
