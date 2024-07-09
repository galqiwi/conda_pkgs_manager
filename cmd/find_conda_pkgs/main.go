package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type CondaPathInfo struct {
	Path          string
	PkgsPath      string
	PkgsDiskUsage uint64
}

func Main() error {
	if len(os.Args) < 2 {
		fmt.Println("Usage: find_conda_pkgs [PATH]...")
	}

	var condaPaths []string

	for _, path := range os.Args[1:] {
		argCondaPaths, err := findAllCondaPaths(path)
		if err != nil {
			return err
		}

		condaPaths = append(condaPaths, argCondaPaths...)
	}

	var condaPathsEntries []CondaPathInfo

	for _, path := range condaPaths {
		PkgsPath := filepath.Join(path, "pkgs")
		pkgsDiskUsage, err := getDiskUsage(PkgsPath)
		if err != nil {
			return err
		}

		condaPathsEntries = append(condaPathsEntries, CondaPathInfo{
			Path:          path,
			PkgsPath:      PkgsPath,
			PkgsDiskUsage: pkgsDiskUsage,
		})
	}

	sort.Slice(condaPathsEntries, func(i, j int) bool {
		ei, ej := condaPathsEntries[i], condaPathsEntries[j]
		if ei.PkgsDiskUsage > ej.PkgsDiskUsage {
			return true
		}
		if ei.PkgsDiskUsage < ej.PkgsDiskUsage {
			return false
		}
		return ei.Path < ej.Path
	})

	displayCondaPathsEntries(condaPathsEntries)

	return nil
}

func main() {
	err := Main()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
