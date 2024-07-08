package handler

import (
	"github.com/gin-gonic/gin"
	"knapsackProblem/routes"
)
func RunBackend() {
	router := gin.Default()
    router.GET("/api/data", routes.GetData)
	router.POST("/api/data", routes.PostData)
	router.DELETE("/api/data/:id", routes.DeleteData)
	router.PATCH("/api/data/:id", routes.UpdateData)
    router.Run(":8080")
}



