package directorytools

import (
	"os"
	"path/filepath"
	"strings"
)

func GetFilesInDirectory(directoryPath string) ([]string, error) {
	var files []string

	// Convert the directory path to an absolute path
	absPath, err := filepath.Abs(directoryPath)
	if err != nil {
		return nil, err
	}

	// Walk the directory tree
	//err = filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
	//	if err != nil {
	//		return err
	//	}
	//	// Check if it's a file (not a directory)
	//	if !info.IsDir() {
	//		files = append(files, strings.Replace(path, "\\", "/", -1))
	//	}
	//	return nil
	//})

	err = filepath.Walk(absPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		fi, err := os.Stat(path)
		if err != nil {
			return err
		}
		switch mode := fi.Mode(); {
		case mode.IsRegular():
			files = append(files, strings.Replace(path, "\\", "/", -1))
		}
		// Check if it's a file (not a directory)
		//if !info.IsDir() {
		//	files = append(files, strings.Replace(path, "\\", "/", -1))
		//}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
