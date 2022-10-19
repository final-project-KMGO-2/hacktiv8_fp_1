package route

import (
	"hacktiv8_fp_1/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine, authController controller.AuthController) {
	todosRoutes := router.Group("/auth")
	{
		todosRoutes.POST("/sign-up", authController.Register)
		todosRoutes.POST("/sign-in", authController.Login)
	}
}