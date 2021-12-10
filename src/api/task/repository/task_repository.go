package repository

import (
	"context"
	"gorm.io/gorm"
	"log"
	"proyectos/src/api/task/domain"
	"proyectos/src/api/task/domain/model"
)

//TaskRepository type
type TaskRepository struct {
	DB *gorm.DB
}

//NewTaskRepository builder
func NewTaskRepository(db *gorm.DB) domain.Repository {
	return &TaskRepository{
		DB: db,
	}
}

func (t *TaskRepository) Create(_ context.Context, task *model.Tasks) (*model.Tasks, error) {
	err := t.DB.Create(&task).Error
	if err != nil {
		log.Printf("Error creating Tasks %v", err)
		return nil, err
	}

	return task, nil
}

func (t *TaskRepository) Delete(_ context.Context, task *model.Tasks) (*model.Tasks, error) {
	err := t.DB.Delete(&task).Error
	if err != nil {
		log.Printf("Error deleting Tasks %v", err)
		return nil, err
	}

	return task, nil
}
