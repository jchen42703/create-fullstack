package testutil

import "os"

// Reads the log file as a string
func GetLogs(logFilePath string) (string, error) {
	fileBytes, err := os.ReadFile(logFilePath)
	if err != nil {
		return "", err
	}

	return string(fileBytes), nil
}
