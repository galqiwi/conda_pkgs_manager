package main

import (
	"os"
	"path/filepath"
)

func getDiskUsage(dirPath string) (uint64, error) {
	var output uint64

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			output += uint64(info.Size())
		}
		return nil
	})

	if err != nil {
		return 0, err
	}
	return output, nil
}
