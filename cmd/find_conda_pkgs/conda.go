package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
)

func isCondaRoot(dirPath string) (bool, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return false, err
	}

	var fileNames []string
	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	if !slices.Contains(fileNames, "pkgs") {
		return false, nil
	}
	if !slices.Contains(fileNames, "envs") {
		return false, nil
	}
	if !slices.Contains(fileNames, "conda-meta") {
		return false, nil
	}
	return true, nil
}

func findAllCondaPaths(dirPath string) ([]string, error) {
	var condaPaths []string

	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "skipping %v: %v\n", path, err.Error())
			return nil
		}

		if !d.IsDir() {
			return nil
		}

		isConda, err := isCondaRoot(path)
		if err != nil || !isConda {
			return nil
		}

		condaPaths = append(condaPaths, path)

		return nil
	})

	if err != nil {
		return nil, err
	}

	return condaPaths, nil
}
