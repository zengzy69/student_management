package model

// Student 学生结构体
type Student struct {
	Name   string             `json:"name"`
	ID     string             `json:"id"`
	Gender string             `json:"gender"`
	Class  string             `json:"class"`
	Grades map[string]float64 `json:"grades"`
}
