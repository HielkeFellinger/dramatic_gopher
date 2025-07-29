package pages

import (
	"github.com/HielkeFellinger/dramatic_gopher/app/views"
	"github.com/gin-gonic/gin"
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
