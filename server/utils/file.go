package utils

import (
	"log"
	"os"
)

// ReadFile func
func ReadFile(path string, bytes int32) []byte {
	fp, err := os.OpenFile(path, os.O_RDONLY, 0755)
	defer fp.Close()
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, bytes)
	n, err := fp.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	return data[:n]
}

// WriteFile func
func WriteFile(path string, data []byte) {
	fp, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()
	_, err = fp.Write(data)
	if err != nil {
		log.Fatal(err)
	}
}

// RemoveFile func
func RemoveFile(path string) {
	err := os.Remove(path)

	if err != nil {
		log.Fatal(err)
		return
	}
}
