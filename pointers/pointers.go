package pointers

import "fmt"

func main() {
	x := 10
	p := &x
	fmt.Println(x)
	fmt.Println(p)
	fmt.Println(*p)
	*p = 5
	increment(&x)
	k := *p
	fmt.Println(k)
}

func increment(p *int) {
	*p += 1
}