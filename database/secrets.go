package database

import "gorm.io/gorm"

type Secret struct {
	gorm.Model
	NamespaceID uint `gorm:"uniqueIndex:idx_secret_namespace_key"`
	Namespace   Namespace
	Key         string `gorm:"uniqueIndex:idx_secret_namespace_key"`
	Value       string
	Description string
	Kind        string
}

func CreateSecret(db *gorm.DB, namespaceName string, key string, value string, desc string, kind string) (*Secret, *gorm.DB) {
	namespace, getNamespaceResult := GetNamespaceByName(db, namespaceName)
	if getNamespaceResult.Error != nil {
		return nil, getNamespaceResult
	}
	secret := Secret{
		NamespaceID: namespace.ID,
		Key:         key,
		Value:       value,
		Description: desc,
		Kind:        kind,
	}

	// TODO: Add Preloading associations
	result := db.Create(&secret)

	return &secret, result
}

func GetSecrets(db *gorm.DB, namespaceName string) (*[]Secret, *gorm.DB) {
	var secrets []Secret
	namespace, getNamespaceResult := GetNamespaceByName(db, namespaceName)
	if getNamespaceResult.Error != nil {
		return nil, getNamespaceResult
	}
	result := db.Where("namespace_id = ?", namespace.ID).Find(&secrets)

	return &secrets, result
}

func GetSecretByKey(db *gorm.DB, namespaceName string, secretKey string) (*Secret, *gorm.DB) {
	var secret Secret
	namespace, getNamespaceResult := GetNamespaceByName(db, namespaceName)
	if getNamespaceResult.Error != nil {
		return nil, getNamespaceResult
	}
	result := db.First(&secret, "namespace_id = ? AND key = ?", namespace.ID, secretKey)

	return &secret, result
}

func UpdateSecret(db *gorm.DB, namespaceName string, key string, value string, desc string, kind string) (*Secret, *gorm.DB) {
	var secret Secret
	namespace, getNamespaceResult := GetNamespaceByName(db, namespaceName)
	if getNamespaceResult.Error != nil {
		return nil, getNamespaceResult
	}
	result := db.First(&secret, "namespace_id = ? AND key = ?", namespace.ID, key)

	secret.Value = value
	secret.Description = desc
	secret.Kind = kind

	db.Save(&secret)

	return &secret, result
}
