package utility

import (
	"log"
	"os"

	"gorm.io/gorm"
)

func LoadSampleData(DB *gorm.DB) {
	log.Println("Fixtures Loading.....")
	pathToFile := "../../fixtures.sql"
	q, err := os.ReadFile(pathToFile)
	if err != nil {
		log.Fatal("fixtures:", err)
	}
	DB.Exec(string(q))
	log.Println("Fixtures success !!!")
}
