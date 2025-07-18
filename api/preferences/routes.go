package preferences

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(routerGroup *gin.RouterGroup, preferencesController *PreferencesController) {

	routerGroup.GET("/preferences", preferencesController.FindAllPreferences)

}
