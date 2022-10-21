package controller

import (
	"fmt"
	"hacktiv8_fp_1/common"
	"hacktiv8_fp_1/dto"

	// "hacktiv8_fp_1/entity"
	// "hacktiv8_fp_1/helpers"
	"hacktiv8_fp_1/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TodosController interface {
	GetTodos(ctx *gin.Context)
	GetTodoById(ctx *gin.Context)
	CreateNewTodo(ctx *gin.Context)
	UpdateTodo(ctx *gin.Context)
	DeleteTodo(ctx *gin.Context)
}

type todosController struct {
	todosService service.TodosService
}

func NewTodosController(ts service.TodosService) TodosController {
	return &todosController{
		todosService: ts,
	}
}

// @BasePath /todos

// @Summary Gets all todos
// @Description Gets all todos
// @Tags todos
// @Accept json
// @Produce json
// @Success 200 {array} entity.Todos
// @Router /todos [get]
func (c *todosController) GetTodos(ctx *gin.Context) {
	result, err := c.todosService.GetTodos(ctx.Request.Context())
	if err != nil {
		res := common.BuildErrorResponse("Failed to get todo list", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

// @Summary get a todo item by ID
// @ID get-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} entity.Todos
// @Failure 404 {object} common.Response
// @Router /todos/{id} [get]
func (c *todosController) GetTodoById(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := c.todosService.GetTodoById(ctx.Request.Context(), id)
	if err != nil {
		res := common.BuildErrorResponse("todo not found", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusNotFound, res)
		return
	}
	ctx.JSON(http.StatusOK, common.BuildResponse(true, "Success", result))
}


func (c *todosController) CreateNewTodo(ctx *gin.Context) {
	credential := ctx.MustGet("credential")
	fmt.Println(credential)

	var todosDTO dto.TodosCreateDTO
	errDTO := ctx.ShouldBind(&todosDTO)
	if errDTO != nil {
		response := common.BuildErrorResponse("Failed to process request", errDTO.Error(), common.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	result, err := c.todosService.CreateTodo(ctx.Request.Context(), todosDTO) // testing aja ini
	if err != nil {
		res := common.BuildErrorResponse("Bad create request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}
	ctx.JSON(http.StatusCreated, common.BuildResponse(true, "Success", result))
}


func (c *todosController) UpdateTodo(ctx *gin.Context) {
	var todoDTO dto.TodosUpdateDTO

	if err := ctx.ShouldBind(&todoDTO); err != nil {
		res := common.BuildErrorResponse("Failed to bind todos request", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	id := ctx.Param("id")

	result, err := c.todosService.UpdateTodo(ctx.Request.Context(), id, todoDTO)

	if err != nil {
		res := common.BuildErrorResponse("Failed to update todos", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", result)
	ctx.JSON(http.StatusOK, res)
}

// @Summary delete a todo item by ID
// @ID delete-todo-by-id
// @Produce json
// @Param id path string true "todo ID"
// @Success 200 {object} entity.Todos
// @Failure 404 {object} common.Response
// @Router /todos/{id} [delete]
func (c *todosController) DeleteTodo(ctx *gin.Context) {
	id := ctx.Param("id")
	err := c.todosService.DeleteTodoByID(ctx.Request.Context(), id)
	if err != nil {
		res := common.BuildErrorResponse("Failed to delete todos", err.Error(), common.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := common.BuildResponse(true, "OK", common.EmptyObj{})
	ctx.JSON(http.StatusOK, res)
}
