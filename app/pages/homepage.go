package pages

import (
	"log"
	"net/http"

	"github.com/HielkeFellinger/dramatic_gopher/app/models"
	"github.com/HielkeFellinger/dramatic_gopher/app/views"
	"github.com/gin-gonic/gin"
)

func Homepage() gin.HandlerFunc {
	return func(c *gin.Context) {
		user := c.MustGet("user").(models.User)

		log.Println(user)

		err := render(c, http.StatusOK, views.Homepage(user))
		if err != nil {
			return
		}
	}
}
