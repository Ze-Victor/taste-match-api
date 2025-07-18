package user

import (
	"errors"
	"fmt"
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

func (controller *UserController) Health(httpContext *gin.Context) {
	httpContext.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
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

func (controller *UserController) MatchUsers(httpContext *gin.Context) {
	userAID := httpContext.Query("userA")
	userBID := httpContext.Query("userB")

	if userAID == "" || userBID == "" {
		httpContext.JSON(http.StatusBadRequest, gin.H{
			"error": "Both 'userA' and 'userB' query parameters are required.",
		})
		return
	}

	idA, errA := strconv.Atoi(userAID)
	if errA != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid 'userA' ID. It must be a positive integer.",
		})
		return
	}

	idB, errB := strconv.Atoi(userBID)
	if errB != nil {
		httpContext.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid 'userB' ID. It must be a positive integer.",
		})
		return
	}

	matchPercentage, err := controller.UserBusiness.CalculateMatch(idA, idB)

	if err != nil {
		if err.Error() == "user not found" {
			httpContext.JSON(http.StatusNotFound, gin.H{
				"error": fmt.Sprintf("Could not find one or both users. Please check IDs: %s, %s", userAID, userBID),
			})
			return
		}

		httpContext.JSON(http.StatusInternalServerError, gin.H{
			"error": "An internal error occurred while calculating the match.",
		})
		return
	}

	httpContext.JSON(http.StatusOK, gin.H{
		"user_a_id":        userAID,
		"user_b_id":        userBID,
		"match_percentage": fmt.Sprintf("%.2f%%", matchPercentage),
	})
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
