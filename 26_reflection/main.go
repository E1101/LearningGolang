package main

import "strconv"

type stringer interface {
	String() string
}


func main() {

}

// ..

func Sprint(x interface{}) string {
	switch x := x.(type) {
	case stringer:
		return x.String()
	case string:
		return x
	case int:
		return strconv.Itoa(x)
		// ...similar cases for int16, uint32, and so on...
	case bool:
		if x {
			return "true"
		}
		return "false"
	default:
		// array, chan, func, map, pointer, slice, struct
		return "???"
	}
}
