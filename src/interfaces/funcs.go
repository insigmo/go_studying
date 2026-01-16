package interfaces

import "fmt"

func Buy(p Payer, price int) {
	err := p.Pay(price)
	fmt.Println(err)
}
