package db

import (
	"pishikot/pkg/models"

	"gorm.io/gorm"
)

type TaskRepo struct {
	db *gorm.DB
}

func (t *TaskRepo) GetTask(id uint) (*models.Task, bool) {
	task := &models.Task{
		ID: id,
	}
	result := t.db.First(&task, "id = ?", id)
	if result.RowsAffected > 0 {
		return task, true
	}
	return task, false
}

func (t *TaskRepo) GetTasks() *[]models.Task {
	tasks := &[]models.Task{}
	t.db.Limit(50).Find(&tasks)
	return tasks
}
