package io_scanner

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Main() {
	//print(adding("%^80", "hhhhh20&&&&nd"))
	//csvReader()
	//zamykanie()
	//operate(true, 2, "+")
	readWithBufio()
}

func readWithBufio() {
	//scanner := bufio.NewReader(os.Stdin)
	var sum int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		text := scanner.Text()
		if text == "" {
			break
		}
		n, _ := strconv.Atoi(text)
		sum += n
	}
	io.WriteString(os.Stdout, strconv.Itoa(sum))
}

func operate(a, b, op interface{}) {
	x, ok := a.(float64)
	if !ok {
		fmt.Printf("value=%v: %T", a, a)
		return
	}

	y, ok := b.(float64)
	if !ok {
		fmt.Printf("value=%v: %T", b, b)
		return
	}
	var res float64
	switch op {
	case "+":
		res = x + y
	case "-":
		res = x - y
	case "*":
		res = x * y
	case "/":
		res = x / y
	default:
		fmt.Println("неизвестная операция")
		return
	}
	fmt.Printf("%.4f", res)
}

func zamykanie() {
	fn := func(a uint) uint {
		var res string
		t := strconv.FormatUint(uint64(a), 10)

		if a == 0 {
			return uint(100)
		}
		for _, s := range t {
			if (s-'0')%2 == 0 && (s-'0') > 0 {
				res += string(s)
			}
		}
		if len(res) == 0 {
			res = "100"
		}
		i, _ := strconv.ParseUint(res, 10, 64)
		return uint(i)
	}
	fmt.Println(fn(10123))
}

func csvReader() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	s = strings.ReplaceAll(s, ",", ".")
	s = strings.ReplaceAll(s, " ", "")
	var nums []float64

	for _, num := range strings.Split(s, ";") {
		n, _ := strconv.ParseFloat(num, 64)
		nums = append(nums, n)
	}
	fmt.Printf("%.4f", nums[0]/nums[1])
}

func adding(first, second string) int64 {
	n1 := cleanStringNumber(first)
	n2 := cleanStringNumber(second)
	return n1 + n2
}

func cleanStringNumber(number string) int64 {
	var res string

	for _, s := range number {
		if unicode.IsDigit(s) {
			res += string(s)
		}
	}
	i, _ := strconv.Atoi(res)
	return int64(i)
}
func work(x int) int {
	return x
}

func sumInt(arr ...int) (int, int) {
	var sum, counter int

	for i := range arr {
		sum += i
		counter++
	}
	return counter, sum
}

func ExampleString() {
	// Создадим строковый литерал s, значение которого "Это строка".
	// Строка состоит из 10 символов, но её длина в байтах будет 19,
	// так как кириллические символы занимают 2 байта, а пробел — 1 байт.
	var s string = "Это строка"
	b := []rune(s)
	a := strings.Index("test", "e")
	с := strings.Index(s, "Э")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(с)
	strings.ToTitle(s)
	fmt.Printf("Длина строки: %d байт\n", len(s))

	// Получим подстроку строки
	fmt.Printf("Напечатаем только второе слово в кавычках: \"%v\"\n", s[7:])

	// Попробуем изменить строку (возникнет ошибка компиляции, так как строки неизменяемы):
	// s[3] = 12
	// Ошибка компиляции: cannot assign to s[3], потому что строки - неизменяемые последовательности.

	// "Изменим строку", создав новую строку
	s = s + " Новая строка"
	fmt.Printf("%v\n", s)

	// Проитерируемся по строке
	for _, b := range s {
		fmt.Printf("%v ", b)
	}
	fmt.Print("\n")

	// Output:
	// Длина строки: 19 байт
	// Напечатаем только второе слово в кавычках: "строка"
	// Это строка Новая строка
	// 1069 1090 1086 32 1089 1090 1088 1086 1082 1072 32 1053 1086 1074 1072 1103 32 1089 1090 1088 1086 1082 1072
}
