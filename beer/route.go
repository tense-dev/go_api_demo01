package beer

import "github.com/gin-gonic/gin"

func NewRoutes(r *gin.Engine) {
	b := r.Group("/beer")
	b.GET("/", getAllBeer)
}

// Handlers

// getAllBeer  show product all beer
// @Summary Retrieves users from mongodb
// @Description Get Beer
// @Produce json
// @Success 200
// @Router /beer [get]
func getAllBeer(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GET ALL BEER",
	})
}
