package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	if err := mainE(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
		// Not needed, but included for readability
		return
	}

	os.Exit(0)
}

func mainE() error {
	if len(os.Args) < 2 {
		return errors.New("prefix: please provide the path to the dist dir")
	}

	pathOfDistDir := filepath.Clean(os.Args[1])

	filepath.WalkDir(pathOfDistDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			binDir, binName := filepath.Split(path)

			prefixedBinName := fmt.Sprintf("f%s", binName)
			prefixedPath := filepath.Join(binDir, prefixedBinName)

			if err := os.Rename(path, prefixedPath); err != nil {
				return err
			}

			fmt.Printf("%s => %s\n", path, prefixedPath)
		}

		return nil
	})

	return nil
}
