package filemanagement

import "os"

func ReadFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		// panic(err)
		return "", err
	}

	return string(content), nil
}
