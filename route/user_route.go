package route

import (
	"hacktiv8_fp_1/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine, todosController controller.TodosController) {
	todosRoutes := router.Group("/auth")
	{
		todosRoutes.POST("/sign-up", )
		todosRoutes.POST("/sign-in", )
	}
}