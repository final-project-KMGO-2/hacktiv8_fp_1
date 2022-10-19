package main

import (
	"hacktiv8_fp_1/controller"
	"hacktiv8_fp_1/repository"
	"hacktiv8_fp_1/route"
	"hacktiv8_fp_1/service"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	var (
		todosRepository repository.TodosRepository = repository.NewTodosRepository("./db/db.json")
		userRepository  repository.UserRepository  = repository.NewUserRepository("./db/user.json")
		todosService    service.TodosService       = service.NewTodosService(todosRepository)
		userService     service.UserService        = service.NewUserService(userRepository)
		authService     service.AuthService        = service.NewAuthService(userRepository)
		jwtService      service.JWTService         = service.NewJWTService()
		todosController controller.TodosController = controller.NewTodosController(todosService)
		authController  controller.AuthController  = controller.NewAuthController(userService, authService, jwtService)
	)
	server := gin.Default()

	// err := seeders.SeedUser()
	// if err != nil {
	// 	log.Fatal(err);
	// }

	route.SwaggerRoute(server)
	route.TodosRoute(server, todosController, jwtService)
	route.UserRoute(server, authController)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
