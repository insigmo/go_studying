package interfaces

type Payer interface {
	Pay(price int) error
}
