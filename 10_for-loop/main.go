package _0_for_loop

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {
		fmt.Println(i)
	}


	// For Condition While-ish
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	// Break / Continue
	j := 0
	for {

		if j%2 == 0 {
			continue
		}

		fmt.Println(j)
		if j >= 50 {
			break
		}

		j++
	}


	// Nested
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, " - ", j)
		}
	}
}
