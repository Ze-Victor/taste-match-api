package user

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	UserBusiness UserBusiness
}

func NewUserController(userBusiness UserBusiness) *UserController {
	return &UserController{
		UserBusiness: userBusiness,
	}
}

func (controller *UserController) Find(httpContext *gin.Context) {
	idStr := httpContext.Param("id")

	idInt, err := strconv.Atoi(idStr)

	if err != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuário inválido"})
		return
	}

	user, err := controller.UserBusiness.Find(idInt)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			httpContext.JSON(http.StatusNotFound, gin.H{"error": "usuário não encontrado"})
			return
		}
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": "ocorreu um erro interno"})
		return
	}

	httpContext.JSON(http.StatusOK, user)
}

func (controller *UserController) FindAllUsers(httpContext *gin.Context) {
	users, err := controller.UserBusiness.FindAll()

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": "an internal error occurred"})
		return
	}

	httpContext.JSON(http.StatusOK, users)
}

func (controller *UserController) Update(httpContext *gin.Context) {
	// TODO impl this
	return
}

func (controller *UserController) Create(httpContext *gin.Context) {
	// TODO impl this
	return
}

func (controller *UserController) Delete(httpContext *gin.Context) {
	// TODO impl this
	return
}
