// Cory Nelson
// Go program to test different variables and the fmt.Printf function

package main
import "fmt"

func main() {
	// Test variables
	a := 1138							// an int
	b := 7.0							// a double
	c := 'a'							// a char
	d := "Hello World!"		// a string
	e := true     				// a bool

	fmt.Printf("%T \n",a) // should show as int
	fmt.Printf("%T \n",b) // should show as double
	fmt.Printf("%T \n",c) // should show as char
	fmt.Printf("%T \n",d) // should show as string
	fmt.Printf("%T \n",e) // should show as bool
}
