package cleanit

import "fmt"

// Run is the method that should be used to Run examples
func Run() {
	fmt.Println("\nRunning clean example")

	fmt.Println(returnUsername(1))
	fmt.Println(returnUsername(2))
	fmt.Println(returnUsername(3))
	fmt.Println(returnUsername(4))
	fmt.Println(returnUsername(5))

	fmt.Println("Clean example finished")
}
