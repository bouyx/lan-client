package ressources

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Couille(c *gin.Context) {
	fmt.Print("saucisse")
	c.String(http.StatusOK, "success")
}
