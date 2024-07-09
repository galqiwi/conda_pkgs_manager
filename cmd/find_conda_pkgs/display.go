package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"math"
)

func PrettifyByteSize(b uint64) string {
	bf := float64(b)
	for _, unit := range []string{"", "Ki", "Mi", "Gi", "Ti", "Pi", "Ei", "Zi"} {
		if math.Abs(bf) < 1024.0 {
			return fmt.Sprintf("%3.1f %sB", bf, unit)
		}
		bf /= 1024.0
	}
	return fmt.Sprintf("%.1f YiB", bf)
}

func displayCondaPathsEntries(condaPathsEntries []CondaPathInfo) {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New("Pkgs path", "Size")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	for _, entry := range condaPathsEntries {
		tbl.AddRow(entry.PkgsPath, PrettifyByteSize(entry.PkgsDiskUsage))
	}
	tbl.Print()
}
