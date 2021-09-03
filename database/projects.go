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

func GetProjectByID(db *gorm.DB, id uint) (*Project, *gorm.DB) {
	var project Project
	result := db.First(&project, id)

	return &project, result
}

func GetProjectByName(db *gorm.DB, name string) (*Project, *gorm.DB) {
	var project Project
	result := db.First(&project, "name = ?", name)

	return &project, result
}

func UpdateProject(db *gorm.DB, name string, desc string, repo_url string, repo_branch string, config_path string) (*Project, *gorm.DB) {
	project, result := GetProjectByName(db, name)
	if result.Error != nil {
		return nil, result
	}

	// this is going to look ugly. I'm only slightly ashamed (not much)
	if name != "" {
		project.Name = name
	}
	if desc == "" {
		project.Description = desc
	}
	if repo_url == "" {
		project.RepositoryUrl = repo_url
	}
	if repo_branch == "" {
		project.RepositoryBranch = repo_branch
	}
	if config_path == "" {
		project.ConfigPath = config_path
	}

	result = db.Save(&project)

	return project, result
}
