package infra

import (
	"os"
	"strings"
)

func ReplaceInFile(filePath, searchStr, replaceStr string) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	newContent := strings.ReplaceAll(string(content), searchStr, replaceStr)
	return os.WriteFile(filePath, []byte(newContent), 0644)
}
