package database

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name             string `gorm:"uniqueIndex"`
	Description      string
	RepositoryUrl    string
	RepositoryBranch string
	ConfigPath       string
}

func CreateProject(db *gorm.DB, name string, desc string, repo_url string, repo_branch string, config_path string) (*Project, *gorm.DB) {
	project := Project{
		Name:             name,
		Description:      desc,
		RepositoryUrl:    repo_url,
		RepositoryBranch: repo_branch,
		ConfigPath:       config_path,
	}
	result := db.Create(&project)

	return &project, result
}

func GetProjects(db *gorm.DB) (*[]Project, *gorm.DB) {
	var projects []Project
	result := db.Find(&projects)

	return &projects, result
}

func GetProjectByID(db *gorm.DB, id int) (*Project, *gorm.DB) {
	var project Project
	result := db.First(&project, id)

	return &project, result
}

func GetProjectByName(db *gorm.DB, name string) (*Project, *gorm.DB) {
	var project Project
	result := db.First(&project, "name = ?", name)

	return &project, result
}
