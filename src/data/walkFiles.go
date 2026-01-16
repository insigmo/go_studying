package data

import (
	"fmt"
	"os"
	"path/filepath"
)

type FileInfo struct {
	Name string
	Size int64
	Path string
}

func Main() {
	const root_path = "./"
	files := []FileInfo{}

	filepath.Walk(root_path,
		func(path string, info os.FileInfo, err error) error {
			fileInfo := FileInfo{info.Name(), info.Size(), path}
			files = append(files, fileInfo)
			return nil
		},
	)
	for _, file := range files {
		fmt.Println(file)
	}
}
