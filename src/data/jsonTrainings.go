package data

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Students struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Address    string
	Phone      string
	Rating     []int
}

type Journal struct {
	ID       int
	Number   string
	Year     int
	Students []Students
}

type TotalAverage struct {
	Average float64
}

type Human struct {
	Name   string `json:"name"`
	Age    int    `json:"age,omitempty"`
	Sex    string
	Status bool `json:""`
}

type AllGlobalID struct {
	ID int `json:"global_id"`
}

func JsonNKVD() {
	file, err := os.Open("/Users/ruabbb1/git/go_studying/src/data/data-20190514T0100.json")
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)
	if err != nil {
		log.Fatal(err)
	}
	data, _ := io.ReadAll(file)
	var res []AllGlobalID
	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Fatal(err)
	}
	var counter int
	for _, r := range res {
		counter += r.ID
	}
	fmt.Println(counter)
}

func JsonBuf() {
	data1 := Human{
		Name:   "Arsen",
		Age:    0,
		Sex:    "withmale",
		Status: false,
	}
	data2 := Human{}
	buf := bytes.Buffer{}

	enc := json.NewEncoder(&buf)
	dec := json.NewDecoder(&buf)

	enc.Encode(data1)
	dec.Decode(&data2)

	fmt.Println(data2)
}

func JsonMain() {
	//journal := Journal{}
	//file, err := os.OpenFile("/Users/ruabbb1/git/go_studying/src/data/students.json", os.O_RDONLY, os.ModePerm)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//data, _ := io.ReadAll(file)
	journal := Journal{}
	data, _ := io.ReadAll(os.Stdin)

	err := json.Unmarshal(data, &journal)
	if err != nil {
		fmt.Println(err)
	}

	data, _ = json.MarshalIndent(journal, "", "    ")
	fmt.Println(string(data))
	var ratingCounter int
	studentCount := len(journal.Students)

	for _, student := range journal.Students {
		ratingCounter += len(student.Rating)
	}
	res := TotalAverage{Average: float64(ratingCounter) / float64(studentCount)}
	out, _ := json.MarshalIndent(res, "", "    ")
	_, err = os.Stdout.Write(out)
	if err != nil {
		log.Fatal(err)
	}

}

func JsonExample() {
	data := Human{
		Name:   "Arsen",
		Age:    0,
		Sex:    "withmale",
		Status: false,
	}
	s, _ := json.MarshalIndent(data, "", "\t")
	if !json.Valid(s) {
		fmt.Printf("invalid JSON\n")
	}
	data.Status = true
	err := json.Unmarshal(s, &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(s))
}
