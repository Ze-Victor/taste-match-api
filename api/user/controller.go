package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	id, valid := httpContext.Params.Get("id")

	if !valid {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	idInt, error := strconv.Atoi(id)

	if error != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	user, error := controller.UserBusiness.Find(idInt)

	if error != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": "an internal error occurred"})
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
