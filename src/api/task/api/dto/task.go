package dto

import (
	"proyectos/src/api/task/domain/model"
	"time"
)

// Task dto
type Task struct {
	ID             int64     `validate:"gt=0" json:"id"`
	Name           string    `validate:"required,min=2,max=50" json:"name"`
	Description    string    `validate:"required" json:"description"`
	StartDate      time.Time `validate:"required" json:"start_date"`
	WorkedHours    int       `validate:"required" json:"worked_hours"`
	EstimatedHours int       `validate:"required" json:"estimated_hours"`
	ProjectID      int64     `validate:"required" json:"project_id"`
	State          string    `validate:"required" json:"state"`
	CreationDate   time.Time `validate:"required" json:"creation_date"`
	AssignedTo     string    `validate:"required" json:"assigned_to"`
}

//ToModel Task to Model
func (t *Task) ToModel() *model.Tasks {

	return &model.Tasks{
		ID:             t.ID,
		Name:           t.Name,
		Description:    t.Description,
		StartDate:      t.StartDate,
		WorkedHours:    t.WorkedHours,
		EstimatedHours: t.EstimatedHours,
		ProjectID:      t.ProjectID,
		State:          t.State,
		CreationDate:   t.CreationDate, //TODO o ver si es la fecha de hoy en dia
		AssignedTo:     t.AssignedTo,
	}
}

//FromModel Model from model
func FromModel(dm *model.Tasks) *Task {
	return &Task{
		ID:             dm.ID,
		Name:           dm.Name,
		Description:    dm.Description,
		StartDate:      dm.StartDate,
		WorkedHours:    dm.WorkedHours,
		EstimatedHours: dm.EstimatedHours,
		ProjectID:      dm.ProjectID,
		State:          dm.State,
		CreationDate:   dm.CreationDate,
		AssignedTo:     dm.AssignedTo,
	}
}
