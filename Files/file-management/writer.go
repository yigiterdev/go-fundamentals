package fileutils

import "os"

func WriteToFile(filename string, content string) error {
	err := os.WriteFile(filename, []byte(content), 0o644)

	return err
}
