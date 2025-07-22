package pages

import (
	"github.com/gin-gonic/gin"
	"hielkefellinger.nl/dramatic_gopher/app/views"
	"net/http"
)

func Homepage() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := render(c, http.StatusOK, views.Homepage())
		if err != nil {
			return
		}
	}
}
