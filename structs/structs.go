package structs

import "fmt"

type rocket struct{}

func (r *rocket) Fly() {
	fmt.Println("Flying...")
}

// Run is the method that should be used to Run examples
func Run() {
	fmt.Println("Running structs example")

	r := &rocket{}
	fmt.Println(r)
	r.Fly()

	r = new(rocket)
	fmt.Println(r)
	r.Fly()

	var ro *rocket
	ro = &rocket{}
	fmt.Println(ro)

	fmt.Println("Structs example finished")
}
