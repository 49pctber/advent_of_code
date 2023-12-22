package main

import (
	"fmt"
	"os"
	"path/filepath"
)

/*
initializes a directory structure for a year's worth of Advent of Code problems
to initialize 2024, run `go run . 2024`

This code was taken from ChatGPT
*/
func main() {

	// Create the main folder
	mainFolderPath := os.Args[1]
	err := os.Mkdir(mainFolderPath, 0755)
	if err != nil {
		fmt.Println("Error creating main folder:", err)
		return
	}

	// Insert 25 subfolders and create files & another folder in each subfolder
	for i := 1; i <= 25; i++ {
		day_label := fmt.Sprintf("%02d", i)
		subfolderName := fmt.Sprintf("day%s", day_label)
		subfolderPath := filepath.Join(mainFolderPath, subfolderName)

		// Create the subfolder
		err := os.Mkdir(subfolderPath, 0755)
		if err != nil {
			fmt.Printf("Error creating subfolder %s: %s\n", subfolderName, err)
			continue
		}

		// Create files in the subfolder
		file1Path := filepath.Join(subfolderPath, fmt.Sprintf("day%s.go", day_label))
		file2Path := filepath.Join(subfolderPath, fmt.Sprintf("day%s_test.go", day_label))

		err = os.WriteFile(file1Path, []byte("package main\n\nfunc main() {\n\n}"), os.ModeAppend)
		if err != nil {
			fmt.Printf("Error creating File1.txt in %s: %s\n", subfolderName, err)
		}

		err = os.WriteFile(file2Path, []byte("package main\n"), os.ModeAppend)
		if err != nil {
			fmt.Printf("Error creating File2.txt in %s: %s\n", subfolderName, err)
		}

		// Create another folder within the subfolder
		anotherFolderPath := filepath.Join(subfolderPath, "input")
		err = os.Mkdir(anotherFolderPath, 0755)
		if err != nil {
			fmt.Printf("Error creating AnotherFolder in %s: %s\n", subfolderName, err)
		}
	}
	fmt.Println("Folder initialization complete.")
}
