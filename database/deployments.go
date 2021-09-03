package database

import (
	"time"

	pb "github.com/oceannik/oceannik/proto"
	"gorm.io/gorm"
)

type Deployment struct {
	gorm.Model
	NamespaceID uint
	Namespace   Namespace
	ProjectID   uint
	Project     Project
	Status      string
	ScheduledAt time.Time
	StartedAt   time.Time
	ExitedAt    time.Time
}

func CreateDeployment(db *gorm.DB, namespaceName string, projectName string) (*Deployment, *gorm.DB) {
	namespace, getNamespaceResult := GetNamespaceByName(db, namespaceName)
	if getNamespaceResult.Error != nil {
		return nil, getNamespaceResult
	}
	project, getProjectResult := GetProjectByName(db, projectName)
	if getProjectResult.Error != nil {
		return nil, getProjectResult
	}

	deployment := Deployment{
		NamespaceID: namespace.ID,
		ProjectID:   project.ID,
		Status:      pb.Deployment_SCHEDULED.String(),
		ScheduledAt: time.Now(),
	}

	result := db.Create(&deployment)
	if result.Error != nil {
		return nil, result
	}

	result = db.Joins("Namespace").Joins("Project").First(&deployment, &deployment.ID)
	if result.Error != nil {
		return nil, result
	}

	return &deployment, result
}

func GetDeployments(db *gorm.DB, namespaceName string) (*[]Deployment, *gorm.DB) {
	namespace, getNamespaceResult := GetNamespaceByName(db, namespaceName)
	var deployments []Deployment
	if getNamespaceResult.Error != nil {
		return nil, getNamespaceResult
	}
	result := db.Joins("Namespace").Joins("Project").Where("namespace_id = ?", namespace.ID).Order("deployments.scheduled_at desc").Find(&deployments)

	return &deployments, result
}

func GetDeploymentByID(db *gorm.DB, id uint) (*Deployment, *gorm.DB) {
	var deployment Deployment
	result := db.Joins("Namespace").Joins("Project").First(&deployment, id)

	return &deployment, result
}

func UpdateDeploymentStatus(db *gorm.DB, id uint, status string, startedAt time.Time, exitedAt time.Time) (*Deployment, *gorm.DB) {
	var deployment Deployment
	result := db.First(&deployment, id)

	deployment.Status = status
	deployment.StartedAt = startedAt
	deployment.ExitedAt = exitedAt

	db.Save(&deployment)

	return &deployment, result
}
