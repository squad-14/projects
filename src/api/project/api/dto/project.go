package dto

import (
	"proyectos/src/api/project/domain/model"
	"time"
)

type Project struct {
	Code   int  `validate:"gt=0"` // TODO: if this is a PK, it could be autogenerated
	Name string `validate:"required,min=2,max=50"`
	Description string
	StartDate time.Time
	EndDate time.Time
	WorkedHours int
}

func (project *Project) ToModel() *model.Project {

	return &model.Project{
		Code: project.Code,
		Name: project.Name,
		Description: project.Description,
		StartDate: project.StartDate,
		EndDate: project.EndDate,
		WorkedHours: project.WorkedHours,
	}
}

func (project *Project) FromModel(modelProject *model.Project) {
	project.Code = modelProject.Code
	project.Name = modelProject.Name
	project.Description = modelProject.Description
	project.StartDate = modelProject.StartDate
	project.EndDate = modelProject.EndDate
	project.WorkedHours = modelProject.WorkedHours
}
