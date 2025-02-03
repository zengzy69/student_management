package test

import (
	"encoding/csv"
	"github.com/stretchr/testify/assert"
	"os"
	"student_management/model"
	"student_management/service"
	"student_management/storage"
	"testing"
)

// 测试添加学生
func TestAddStudent(t *testing.T) {
	student := model.Student{
		Name:   "Tom",
		ID:     "2023001",
		Gender: "Male",
		Class:  "Class A",
		Grades: map[string]float64{"Math": 90, "English": 88},
	}

	// 添加学生
	service.AddStudent(student)

	// 检查学生是否添加成功
	storedStudent, exists := storage.StudentMap["2023001"]
	assert.True(t, exists)
	assert.Equal(t, "Tom", storedStudent.Name)
	assert.Equal(t, "2023001", storedStudent.ID)
	assert.Equal(t, "Male", storedStudent.Gender)
	assert.Equal(t, "Class A", storedStudent.Class)
	assert.Equal(t, 90.0, storedStudent.Grades["Math"])
	assert.Equal(t, 88.0, storedStudent.Grades["English"])
}

// 测试获取学生
func TestGetStudent(t *testing.T) {
	student := model.Student{
		Name:   "Tom",
		ID:     "2023002",
		Gender: "Male",
		Class:  "Class B",
		Grades: map[string]float64{"Math": 95, "English": 85},
	}

	// 添加学生
	service.AddStudent(student)

	// 获取学生
	storedStudent, exists := service.GetStudent("2023002")
	assert.True(t, exists)
	assert.Equal(t, "Tom", storedStudent.Name)
	assert.Equal(t, "2023002", storedStudent.ID)
	assert.Equal(t, "Male", storedStudent.Gender)
	assert.Equal(t, "Class B", storedStudent.Class)
	assert.Equal(t, 95.0, storedStudent.Grades["Math"])
	assert.Equal(t, 85.0, storedStudent.Grades["English"])
}

// 测试更新学生信息
func TestUpdateStudent(t *testing.T) {
	student := model.Student{
		Name:   "Tom",
		ID:     "2023003",
		Gender: "Male",
		Class:  "Class C",
		Grades: map[string]float64{"Math": 80, "English": 90},
	}

	// 添加学生
	service.AddStudent(student)

	// 更新学生信息
	updatedStudent := model.Student{
		Name:   "Tommy",
		ID:     "2023003",
		Gender: "Male",
		Class:  "Class D",
		Grades: map[string]float64{"Math": 92, "English": 85},
	}

	service.UpdateStudent("2023003", updatedStudent)

	// 获取更新后的学生信息
	storedStudent, exists := service.GetStudent("2023003")
	assert.True(t, exists)
	assert.Equal(t, "Tommy", storedStudent.Name)
	assert.Equal(t, "2023003", storedStudent.ID)
	assert.Equal(t, "Male", storedStudent.Gender)
	assert.Equal(t, "Class D", storedStudent.Class)
	assert.Equal(t, 92.0, storedStudent.Grades["Math"])
	assert.Equal(t, 85.0, storedStudent.Grades["English"])
}

// 测试删除学生
func TestDeleteStudent(t *testing.T) {
	student := model.Student{
		Name:   "Tom",
		ID:     "2023004",
		Gender: "Male",
		Class:  "Class E",
		Grades: map[string]float64{"Math": 70, "English": 75},
	}

	// 添加学生
	service.AddStudent(student)

	// 删除学生
	service.DeleteStudent("2023004")

	// 检查学生是否被删除
	storedStudent, exists := service.GetStudent("2023004")
	assert.False(t, exists)
	assert.Nil(t, storedStudent)
}

// 测试 CSV 处理功能
func TestProcessCSV(t *testing.T) {
	// 假设我们有一个简单的 CSV 文件内容（注意：在实际测试中，这应该是从临时文件中读取）
	csvData := [][]string{
		{"Name", "ID", "Gender", "Class", "Math:80", "English:75"},
		{"Alice", "2023005", "Female", "Class F", "Math:90", "English:85"},
		{"Bob", "2023006", "Male", "Class G", "Math:95", "English:88"},
	}

	// 临时存储 CSV 数据到文件
	tempFile := "./temp_students.csv"
	file, err := os.Create(tempFile)
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile)

	writer := csv.NewWriter(file)
	err = writer.WriteAll(csvData)
	if err != nil {
		t.Fatalf("Failed to write CSV data: %v", err)
	}
	file.Close()

	// 处理 CSV 文件
	service.ProcessCSV(tempFile)

	// 检查学生是否已正确添加到存储
	storedStudent1, exists1 := storage.StudentMap["2023005"]
	storedStudent2, exists2 := storage.StudentMap["2023006"]
	assert.True(t, exists1)
	assert.True(t, exists2)
	assert.Equal(t, "Alice", storedStudent1.Name)
	assert.Equal(t, "Bob", storedStudent2.Name)
	assert.Equal(t, 90.0, storedStudent1.Grades["Math"])
	assert.Equal(t, 85.0, storedStudent2.Grades["English"])
}
