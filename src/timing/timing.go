package timing

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"
)

func TimeCompare() {
	start := time.Now()
	finish := start.Add(time.Hour)
	fmt.Println(start.Before(finish))
	fmt.Println(start.After(finish))
	a, _ := start.MarshalJSON()
	fmt.Printf("%s\n", string(a))

	fmt.Printf("%v", finish.Compare(start))
}

func TimeChanger() {
	t, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	dt, _ := time.Parse(time.DateTime, strings.TrimSpace(t))
	if dt.Hour() <= 13 {
		fmt.Println(dt.Format(time.DateTime))
	} else {
		dt = dt.Add(time.Hour * 24)
		fmt.Println(dt.Format(time.DateTime))

	}
}

func TimeDelta() {
	buf := bytes.Buffer{}
	buf.Write([]byte("13.03.2018 14:00:15,12.03.2018 14:00:15\n"))
	//reader := bufio.NewReader(&buf)
	reader := bufio.NewReader(os.Stdin)

	first, _ := reader.ReadString(',')
	second, _ := reader.ReadString(',')
	fmt.Println(time.Minute * 10)
	first = strings.Trim(first, ",")
	second = strings.TrimSpace(second)
	layout := "2.01.2006 15:04:05"
	start, _ := time.Parse(layout, first)
	finish, _ := time.Parse(layout, second)
	if start.After(finish) {
		start, finish = finish, start
	}
	fmt.Println(start.Sub(finish).Round(time.Second))
}

func TimeDelta2() {
	const now = 1589570165
	var minutes, seconds int64
	fmt.Scanf("%d мин. %d", &minutes, &seconds)
	rightNow := time.Unix(now, 0).UTC()

	rightNow = rightNow.Add(time.Duration(minutes)*time.Minute + time.Duration(seconds)*time.Second)
	fmt.Println(rightNow.Format(time.UnixDate))

}
