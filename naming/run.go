package naming

import "fmt"

// Run is the method that should be used to Run examples
func Run() {
	fmt.Println("\nRunning naming example")

	s := "I am a string"
	fmt.Println(s)

	p := parse(s)
	fmt.Println(p)

	p = removeSpaces(s)
	fmt.Println(p)

	p = remove(s, " ")
	fmt.Println(p)

	fmt.Println("Naming example finished")
}
