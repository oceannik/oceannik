package database

import (
	"log"

	pb "github.com/oceannik/oceannik/proto"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Open(databasePath string) *gorm.DB {
	log.Printf("[Database] Opening database at %s", databasePath)

	db, err := gorm.Open(sqlite.Open(databasePath), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the Oceannik SQLite database.")
	}

	return db
}

func InitData(db *gorm.DB) {
	log.Printf("[Database] Initializing the database with default data")

	log.Printf("[Database] Creating the default namespace...")
	defaultNamespace := Namespace{
		Name:        "default",
		Description: "This is the default namespace for all your projects.",
	}
	db.Create(&defaultNamespace)

	log.Printf("[Database] Creating the default project...")
	defaultProject := Project{
		Name:             "example-project",
		Description:      "Example project that deploys `example-test-app`",
		RepositoryUrl:    "https://github.com/oceannik/examples",
		RepositoryBranch: "main",
		ConfigPath:       "example-project/oceannik.yml",
	}
	db.Create(&defaultProject)

	log.Printf("[Database] Creating namespace secrets...")
	namespaceSecrets := []Secret{
		{
			NamespaceID: defaultNamespace.ID,
			Key:         "OCEANNIK_HOSTS",
			Value:       "<replace_with_your_own>",
			Description: "File describing hosts and their roles. Make sure listed SSH keys have the exact same key names as in Oceannik's secrets.",
			Kind:        pb.SecretKind_FILE.String(),
		},
		{
			NamespaceID: defaultNamespace.ID,
			Key:         "GIT_REPO_SSH_PRIVATE_KEY",
			Value:       "<replace_with_your_own>",
			Description: "SSH Private key for accessing the project's repository. The example repository does not require one because it's a public repository.",
			Kind:        pb.SecretKind_FILE.String(),
		},
		{
			NamespaceID: defaultNamespace.ID,
			Key:         "INFRA_SSH_PRIVATE_KEY_1",
			Value:       "<replace_with_your_own>",
			Description: "SSH Private key for accessing the node the deployment will be performed against. More than one key can be configured.",
			Kind:        pb.SecretKind_FILE.String(),
		},
		{
			NamespaceID: defaultNamespace.ID,
			Key:         "INFRA_SSH_PRIVATE_KEY_2",
			Value:       "<replace_with_your_own>",
			Description: "SSH Private key for accessing the node the deployment will be performed against. This is a second key.",
			Kind:        pb.SecretKind_FILE.String(),
		},
	}
	db.Create(&namespaceSecrets)
}

func PerformAutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&Namespace{}, &Deployment{}, &Secret{}, &Project{})
}
