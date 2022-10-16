package fs

import (
	"log"
	"os"
)

func CreateAllDirectories() {
	dirs := []string{
		"./internal/adaptor/repo",
		"./internal/core/port",
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			log.Panicf("error cannot create diretory err = %v", err)
		}
	}
}
