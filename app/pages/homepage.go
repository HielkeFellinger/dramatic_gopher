package pages

import (
	"net/http"

	"github.com/HielkeFellinger/dramatic_gopher/app/models"
	"github.com/HielkeFellinger/dramatic_gopher/app/views"
	"github.com/gin-gonic/gin"
)

func Homepage() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(models.User)
		if err := render(c, http.StatusOK, views.Homepage(user)); err != nil {
			return
		}
	}
}
