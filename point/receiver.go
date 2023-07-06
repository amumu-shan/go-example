package point

import "fmt"

type data struct {
	name string
}
type printer interface {
	print()
}

func (p *data) print() {
	fmt.Println("name:", p.name)
}
