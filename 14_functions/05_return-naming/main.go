package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe"))
}

func greet(fname string, lname string) (s string) {
	// If you name your return arguments you are creating local variables just like
	// with your function parameters. This time when I set the id variable,
	// I remove the colon (:) from the short variable declaration and convert it to an assignment operation.

	// When you specify the named return values, you can assign the return values to the named variables
	// and can exit from functions by simply specifying the return keyword, without providing the return values
	// along with the return statement.
	s = fmt.Sprint(fname, lname)
	return
}

/*
IMPORTANT
Avoid using named returns.


Occasionally named returns are useful. Read this article for more information:
https://www.goinggo.net/2013/10/functions-and-naked-returns-in-go.html
*/
