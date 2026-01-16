package data

import (
	"archive/zip"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CsvMain() {
	rootPath := "/Users/ruabbb1/git/go_studying/src/data"
	zipFile := filepath.Join(rootPath, "task.zip")

	zr, err := zip.OpenReader(zipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer zr.Close()

	dstFolder := filepath.Join(filepath.Dir(zipFile), "unpacked")
	err = os.MkdirAll(dstFolder, 0755)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range zr.File {
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		reader := csv.NewReader(rc)
		data, err := reader.ReadAll()
		if err != nil || len(data) == 0 || strings.Contains(data[0][0], "\u0000") {
			continue
		}
		dst := filepath.Join(dstFolder, f.Name)
		fmt.Printf("unzipping %s to %s\n", dst, data[4][2])

	}
}
