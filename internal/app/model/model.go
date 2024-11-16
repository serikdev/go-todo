package model

import "gorm.io/gorm"

type Task struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	DueData     string `json:"due_data"`
	OverData    bool   `json:"over_data" gorm:"default:false"`
	Completed   bool   `json:"completed" gorm:"default:false"`
}

func (t *Task) BeforeCreate(tx *gorm.DB) (err error) {
	if t.OverData {
		t.Completed = true
	}
	return
}
