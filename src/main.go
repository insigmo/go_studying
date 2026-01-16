package main

import (
	"fmt"
	"main/interfaces"
	"main/list"
	"main/person"
	"strings"
)

const (
	a = iota + 1
	_
	b
	c
	d = c + 2
	t
	i
	i2 = iota + 2
)

func main() {
	//Slices()
	//Maps()
	//ChrevoVeshatel("Arsen")
	//Loops()
	//Panics()
	//Structs()
	//Methods()
	//Packages()
	//Interfaces()
	//io_scanner.ScanIO()
	//io_scanner.Main()
	//data.CsvSequence()
	//data.JsonMain()
	//data.JsonExample()
	//data.JsonBuf()
	//data.JsonNKVD()
	//timing.TimeCompare()
	//timing.TimeChanger()
	//timing.TimeDelta()
	//timing.TimeDelta2()
	//goroutings.Goroutings()
	//goroutings.WorkGroup()
	//goroutings.Ticker()
	//goroutings.Timer()
	//goroutings.DownloadAccounts()
	//goroutings.GoMain()
	//list.Main()
	list.ReversedMain()
	//data.CsvMain()
	//data.Main()
	//io_scanner.ReadFile()
	//Constants()
}

type Battery struct {
	level string
}

func (b Battery) String() string {
	count := strings.Count(b.level, "1")
	empty := strings.Repeat(" ", 10-count)
	full := strings.Repeat("X", count)
	return fmt.Sprintf("[%s%s]", empty, full)
}

func Interfaces() {
	price := 100
	card := interfaces.Card{Balance: 100}
	wallet := interfaces.Wallet{Balance: 100}
	applePay := interfaces.ApplePay{Balance: 90}
	interfaces.Buy(&card, price)
	interfaces.Buy(&wallet, price)
	interfaces.Buy(&applePay, price)
}

func Packages() {
	p := person.NewPerson{Name: "Arsen", Age: 35}
	p.UpdateAge(10)
	fmt.Println(p.Name, p.Age)
}

func Methods() {
	acc := Account{Name: "Arsen"}
	acc.SetName("Betal")
	fmt.Println(acc)
}

func (acc Account) SetName(name string) {
	acc.Name = name
	fmt.Println(acc)
}

func Structs() {
	acc := Account{
		Name:    "Arsen",
		Balance: -1000,
		Person:  Person{Name: "Aaaarsseeeen", Age: 35},
	}
	fmt.Printf("%+v\n", acc)
	fmt.Printf("%+v\n", acc.Age)
}

type Account struct {
	Name    string
	Balance int
	Person
}

type Person struct {
	Name string
	Age  int
}

func Slices() {
	//a := make([]int, 4, 6)
	a := []int{1, 2, 3}
	Print(a)
	//b := a[:]
	//b[0] = 1
	//b = append(b, 2)
	//Print(a)
	//
	//c := make([]int, len(b))
	//copy(c, b)
	//Print(c)
}
func Panics() {
	//defer func() {
	//	if err := recover(); err != nil {
	//		fmt.Println("Recovered in f", err)
	//	}
	//}()
	panic("top level panic")
}

func Loops() {
	str := "Hello"
	for pos, char := range str {
		fmt.Printf("%#U %d\n", char, pos)
	}
}

func Maps() {
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Printf("%s=%d\n", k, v)
	}

	key, keyExists := m["c"]
	fmt.Printf("%v %v\n", key, keyExists)

	key, keyExists = m["b"]
	fmt.Printf("%v %v\n", key, keyExists)
	delete(m, "b")
	m["c"] = 3
	//fmt.Printf("%v\n", m)

	if _, exists := m["b"]; exists != true {
		m["b"] = 2
	}

	for k, v := range m {
		fmt.Printf("%v %d\n", k, v)
	}
}

func ChrevoVeshatel(value string) {
	switch value {
	case "Arsen":
		fmt.Printf("%v pituh", value)
	case "Betal":
		fmt.Printf("%v Berbekov", value)
	default:
		fmt.Println("end")
	}
}

func Print(a []int) {
	fmt.Printf("len=%v cap=%v\n", len(a), cap(a))
	fmt.Printf("%v\n", a)
	fmt.Printf("%#v\n", a)
}

func Constants() {
	fmt.Println(i2)
}
