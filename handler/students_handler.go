package handler

import (
	"net/http"
	"student_management/model"
	"student_management/service"

	"github.com/gin-gonic/gin"
)

// AddStudent godoc
// @Summary 添加学生
// @Description 添加新的学生信息
// @Tags Student
// @Accept  json
// @Produce  json
// @Param student body model.Student true "学生信息"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /students [post]
func AddStudent(c *gin.Context) {
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	service.AddStudent(student)
	c.JSON(http.StatusOK, gin.H{"message": "Student added"})
}

// GetStudent godoc
// @Summary 获取学生信息
// @Description 根据学号查询学生信息
// @Tags Student
// @Produce json
// @Param id path string true "学生学号"
// @Success 200 {object} model.Student
// @Failure 404 {object} map[string]string
// @Router /students/{id} [get]
func GetStudent(c *gin.Context) {
	id := c.Param("id")
	student, exists := service.GetStudent(id)
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}
	c.JSON(http.StatusOK, student)
}

// UpdateStudent godoc
// @Summary 更新学生信息
// @Description 根据学号修改学生信息
// @Tags Student
// @Accept  json
// @Produce  json
// @Param id path string true "学生学号"
// @Param student body model.Student true "学生信息"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /students/{id} [put]
func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student model.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	service.UpdateStudent(id, student)
	c.JSON(http.StatusOK, gin.H{"message": "Student updated"})
}

// DeleteStudent godoc
// @Summary 删除学生信息
// @Description 根据学号删除学生信息
// @Tags Student
// @Param id path string true "学生学号"
// @Success 200 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /students/{id} [delete]
func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	service.DeleteStudent(id)
	c.JSON(http.StatusOK, gin.H{"message": "Student deleted"})
}

// UploadCSV godoc
// @Summary 上传 CSV 文件
// @Description 通过 CSV 文件批量导入学生数据
// @Tags CSV
// @Accept multipart/form-data
// @Param file formData file true "CSV 文件"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /upload [post]
func UploadCSV(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File upload error"})
		return
	}

	dst := "./" + file.Filename
	if err := c.SaveUploadedFile(file, dst); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	go service.ProcessCSV(dst) // 异步处理
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded and processing in background"})
}
