package preferences

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PreferencesController struct {
	PreferencesBusiness PreferencesBusiness
}

func NewPreferencesController(preferencesBusiness PreferencesBusiness) *PreferencesController {
	return &PreferencesController{
		PreferencesBusiness: preferencesBusiness,
	}
}

func (controller *PreferencesController) FindAllPreferences(httpContext *gin.Context) {
	preferencess, err := controller.PreferencesBusiness.FindAll()

	if err != nil {
		httpContext.JSON(http.StatusInternalServerError, gin.H{"error": "an internal error occurred"})
		return
	}

	httpContext.JSON(http.StatusOK, preferencess)
}
