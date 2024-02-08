package helpers

import (
	"fmt"
	"path/filepath"
	"time"
)

func GenerateFileName(filename string) string {
	return fmt.Sprintf("%d |- %s", time.Now().Unix(), filename)
}

func IsPDF(filename string) bool {
	// Extract file extension
	fileExtension := filepath.Ext(filename)
	return fileExtension == ".pdf"
}
