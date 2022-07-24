package utility

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"gorm.io/gorm"
)

func LoadSampleData(DB *gorm.DB) {
	log.Println("Fixtures Loading.....")

	pathToFile := getFixtureFilePath()

	q, err := os.ReadFile(pathToFile)
	if err != nil {
		log.Fatal("fixtures:", err)
	}
	DB.Exec(string(q))
	log.Println("Fixtures success !!!")
}

func getFixtureFilePath() string {
	_, currentDir, _, _ := runtime.Caller(0)
	rootProjectPath := path.Join(path.Dir(currentDir), "../../")
	fmt.Println("Relative", filepath.Join(rootProjectPath, "/fixtures.sql"))

	return filepath.Join(rootProjectPath, "/fixtures.sql")
}
