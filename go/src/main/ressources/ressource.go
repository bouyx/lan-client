package ressources

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Quiche(c *gin.Context) {
	c.String(http.StatusOK, "lorraine")
}
