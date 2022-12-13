package hcl

import (
	"fmt"
	"os"
	"path/filepath"
)

func ParseTFFiles(tfHCLDirectory string, fileContents string) (string, error) {
	entities, err := os.ReadDir(tfHCLDirectory)
	path, _ := filepath.Abs(tfHCLDirectory)
	resultString := fileContents
	if err != nil {
		return resultString, fmt.Errorf(err.Error())
	}
	for _, entity := range entities {
		if entity.IsDir() {
			resultString, err = ParseTFFiles(filepath.Join(path, entity.Name()), fileContents)
			if err != nil {
				return resultString, fmt.Errorf(err.Error())
			}
		} else {
			newFilepath := filepath.Join(path, entity.Name())
			extension := filepath.Ext(newFilepath)
			if extension == ".tf" {
				file, err := os.ReadFile(newFilepath)
				if err != nil {
					return resultString, fmt.Errorf(err.Error())
				}
				resultString = resultString + string(file)
			}
		}
	}
	return resultString, nil
}
