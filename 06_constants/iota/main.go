package iota

import "fmt"

const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

const (
	d = iota // 0
	e        // 1
	f        // 2
)

const (
	_ = iota      // 0
	z = iota * 10 // 1 * 10
	x = iota * 10 // 2 * 10
)

const (
	Ldate = 1 << iota  // 1 << 0 = 000000001 =  1
	Ltime              // 1 << 1 = 000000010 =  2
	Lmicroseconds      // 1 << 2 = 000000100 =  4
	Llongfile          // 1 << 3 = 000001000 =  8
	Lshortfile         // 1 << 4 = 000010000 =  16

	LstdFlags = Ldate | Ltime //  = 00000011 = 3
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

}
