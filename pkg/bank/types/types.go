package types
type Category string
type Payment struct {
	ID  int
	Amount Money
	Category Category
}
type Money int64
