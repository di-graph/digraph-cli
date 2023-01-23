package analytics

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func GetAPIKey(args []string) string {
	apiKey := ""
	for _, arg := range args {
		if strings.HasPrefix(arg, "--api-key") {
			splitStr := strings.Split(arg, "--api-key=")
			if len(splitStr) > 1 {
				apiKey = splitStr[1]
			}
		}
	}
	if len(apiKey) == 0 {
		err := godotenv.Load(".env")

		if err != nil {
			return ""
		}

		apiKey = os.Getenv("api-key")
	}

	return apiKey
}
