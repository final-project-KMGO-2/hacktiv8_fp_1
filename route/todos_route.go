package route

import (
	"hacktiv8_fp_1/controller"
	"hacktiv8_fp_1/middlewares"
	"hacktiv8_fp_1/service"

	"github.com/gin-gonic/gin"
)

func TodosRoute(router *gin.Engine, todosController controller.TodosController, jwtSvc service.JWTService) {
	todosRoutes := router.Group("/todos", middlewares.AuthMiddleware(jwtSvc))
	{
		todosRoutes.GET("", todosController.GetTodos)
		todosRoutes.POST("", todosController.CreateNewTodo)
	}
}
