package srp

import "fmt"

type fox struct{}

func (f *fox) Say() {
	fmt.Println("Ring-ding-ding-ding-dingeringeding!")
}

func (f *fox) Meow() {
	fmt.Println("Meow")
}

// Run is the method that should be used to Run examples
func Run() {
	fmt.Println("\nRunning SRP example")
	f := new(fox)
	f.Say()
	f.Meow()

	fmt.Println("SRP example finished")
}
