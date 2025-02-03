package router

import (
	"student_management/handler"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/students", handler.AddStudent)
	r.GET("/students/:id", handler.GetStudent)
	r.PUT("/students/:id", handler.UpdateStudent)
	r.DELETE("/students/:id", handler.DeleteStudent)
	r.POST("/upload", handler.UploadCSV)
}
