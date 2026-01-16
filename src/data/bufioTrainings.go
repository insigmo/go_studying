package data

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CsvSequence() {
	csvFilePath := filepath.Join("/Users/ruabbb1/git/go_studying/src/data", "sequence.csv")
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	reader := bufio.NewReader(csvFile)
	for i := 1; ; i++ {
		data, err := reader.ReadString(';')
		if err != nil {
			break
		}

		if data == "0;" {
			fmt.Println(i)
		}
	}
}
