package storage

import "student_management/model"

// 用 map 存储学生数据
var StudentMap = make(map[string]*model.Student)
