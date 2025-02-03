package main

import (
	_ "student_management/docs" // 引入 Swagger 生成的文档
	"student_management/router"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title 学生成绩管理系统 API
// @version 1.0
// @description 这是一个使用 Gin + Swagger 构建的学生管理系统 API。
// @host localhost:8080
// @BasePath /
func main() {
	r := gin.Default()

	// 配置 Swagger 路由
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 设置 API 路由
	router.SetupRoutes(r)

	// 启动服务
	r.Run(":8080")
}
