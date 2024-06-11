package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

//go:embed template
var template embed.FS

func main() {

	if len(os.Args) != 3 {
		log.Fatal("Specify year and day for template")
	}

	var year, day int
	fmt.Sscanf(os.Args[1], "%d", &year)
	fmt.Sscanf(os.Args[2], "%d", &day)

	dir := fmt.Sprintf("%04d%cday%02d", year, os.PathSeparator, day)

	_, err := os.Stat(dir)
	if err == nil {
		fmt.Printf("directory %v already exists\n", dir)
		return
	}

	subfs, err := fs.Sub(template, "template")
	if err != nil {
		panic(err)
	}

	err = fs.WalkDir(subfs, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return os.MkdirAll(filepath.Join(dir, path), 0766)
		} else {
			b, err := fs.ReadFile(subfs, path)
			if err != nil {
				return err
			}
			err = os.WriteFile(filepath.Join(dir, path), b, 0766)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Template available in %v\n", dir)
}
