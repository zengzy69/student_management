package service

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"student_management/model"
	"student_management/storage"
	"sync"
)

var mu sync.RWMutex

// AddStudent 添加学生
func AddStudent(student model.Student) {
	mu.Lock()
	storage.StudentMap[student.ID] = &student
	mu.Unlock()
}

// GetStudent 获取学生
func GetStudent(id string) (*model.Student, bool) {
	mu.RLock()
	student, exists := storage.StudentMap[id]
	mu.RUnlock()
	return student, exists
}

// UpdateStudent 更新学生
func UpdateStudent(id string, student model.Student) {
	mu.Lock()
	storage.StudentMap[id] = &student
	mu.Unlock()
}

// DeleteStudent 删除学生
func DeleteStudent(id string) {
	mu.Lock()
	delete(storage.StudentMap, id)
	mu.Unlock()
}

// ProcessCSV 处理CSV文件
func ProcessCSV(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Failed to open file:", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Failed to read CSV:", err)
	}

	var wg sync.WaitGroup
	dataCh := make(chan *model.Student, len(records))

	// 解析CSV并发送到channel
	go func() {
		for _, record := range records {
			wg.Add(1)
			go func(record []string) {
				defer wg.Done()
				id := record[1]
				grades := parseGrades(record[4:])
				dataCh <- &model.Student{Name: record[0], ID: id, Gender: record[2], Class: record[3], Grades: grades}
			}(record)
		}
		wg.Wait()
		close(dataCh)
	}()

	// 处理学生数据
	for student := range dataCh {
		mu.Lock()
		storage.StudentMap[student.ID] = student
		mu.Unlock()
	}
}

// parseGrades 解析成绩
func parseGrades(data []string) map[string]float64 {
	grades := make(map[string]float64)
	for _, entry := range data {
		parts := strings.Split(entry, ":")
		if len(parts) == 2 {
			score, err := strconv.ParseFloat(parts[1], 64)
			if err == nil {
				grades[parts[0]] = score
			}
		}
	}
	return grades
}
