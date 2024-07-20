package routes

import (
	"github.com/gin-gonic/gin"
	"knapsackProblem/model"
	
)

func getRoutes(endpointGroup *gin.RouterGroup) {
	v2 := endpointGroup.Group("/knapsack")
	{
		v2.GET("/", GetData)
		v2.POST("/", PostData)
		v2.PATCH("/:ID", UpdateData)
		v2.DELETE("/:ID", DeleteData)
		// v2.GET("/calculate", CalculateKnapsack)
	}
}

func GetData(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Hello, World!"})
}

func PostData(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Data append inmto DB!"})
}

func UpdateData(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Data updated !"})
}

func DeleteData(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Data removed !"})
}

func CalculateKnapsack(ctx *gin.Context) ([] model.FinalKnapsack, error) {
	return []model.FinalKnapsack{}, nil
}

func GenerateCSVKnapsack(knapsack [] model.FinalKnapsack) {
	// handler.Write2CSV("result", knapsack)
}

func ReadCSV(sPath string) {
	// data, _ := handler.ReadCSVfile(sPath)
    // fmt.Println(data)
}