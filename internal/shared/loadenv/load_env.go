// loading env files is extracted, and inteded to be shared acrose multiple packages
// to prevent diferent behaviour across environments.
//
// LoadEnv should find the project root and use the .env that is loacated there
// TODO: should be expanded to include and inject filed for different environments

package loadenv

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

const projectDirName = "all-caps-backend" 

func LoadEnv() {
    projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
    currentWorkDirectory, _ := os.Getwd()
    rootPath := projectName.Find([]byte(currentWorkDirectory))

    err := godotenv.Load(string(rootPath) + `/.env`)

    if err != nil {
        log.Fatalf("Error loading .env file")
    }
}