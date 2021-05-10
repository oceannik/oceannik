package database

import "gorm.io/gorm"

type Namespace struct {
	gorm.Model
	Name        string `gorm:"uniqueIndex"`
	Description string
}

func CreateNamespace(db *gorm.DB, name string, description string) (*Namespace, *gorm.DB) {
	namespace := Namespace{Name: name, Description: description}
	result := db.Create(&namespace)

	return &namespace, result
}

// func UpdateOrCreateNamespace(db *gorm.DB, name string, description string) (*Namespace, *gorm.DB) {}

func GetNamespaces(db *gorm.DB) (*[]Namespace, *gorm.DB) {
	var namespaces []Namespace
	result := db.Find(&namespaces)

	return &namespaces, result
}

func GetNamespaceByID(db *gorm.DB, id int) (*Namespace, *gorm.DB) {
	var namespace Namespace
	result := db.First(&namespace, id)

	return &namespace, result
}

func GetNamespaceByName(db *gorm.DB, name string) (*Namespace, *gorm.DB) {
	var namespace Namespace
	result := db.First(&namespace, "name = ?", name)

	return &namespace, result
}
