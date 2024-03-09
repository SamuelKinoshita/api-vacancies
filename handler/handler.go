package handler

import (
	"backend_go/config"

	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandle() {
	logger = config.GetLooger("handler")
	db = config.GetSQlite()
}

// func CreateOpeningHandler(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg": "Post Open",
// 	})
// }

// func ShowOpeningHandler(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg": "GET Open",
// 	})
// }

// func DeleteOpeningHandler(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg": "GET Open",
// 	})
// }

// func UpdateOpeningHandler(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg": "GET Open",
// 	})
// }

// func ListOpeningsHandler(ctx *gin.Context) {
// 	ctx.JSON(http.StatusOK, gin.H{
// 		"msg": "GET Open",
// 	})
// }
