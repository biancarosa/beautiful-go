package naming

import "fmt"

// Run is the method that should be used to Run examples
func Run() {
	s := "I am a string"

	p := parse(s)
	fmt.Println(p)

	p = removeSpaces(s)
	fmt.Println(p)

	p = remove(s, " ")
	fmt.Println(p)
}
