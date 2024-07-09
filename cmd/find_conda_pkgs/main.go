package main

import (
	"fmt"
	"os"
	"sort"
)

type CondaPathInfo struct {
	Path          string
	PkgsDiskUsage uint64
}

func Main() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("usage: find_conda_pkgs path")
	}

	rootPath := os.Args[1]

	condaPaths, err := findAllCondaPaths(rootPath)
	if err != nil {
		return err
	}

	var condaPathsEntries []CondaPathInfo

	for _, path := range condaPaths {
		diskUsage, err := getDiskUsage(path)
		if err != nil {
			return err
		}

		condaPathsEntries = append(condaPathsEntries, CondaPathInfo{
			Path:          path,
			PkgsDiskUsage: diskUsage,
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
