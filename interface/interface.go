package _interface

type Shaper interface {
	Area() float32
}
type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}
func (sq *Square) Sub() float32 {
	return sq.side - sq.side
}
