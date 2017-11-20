package structs

import "fmt"

type rocket struct{}

func (r *rocket) Fly() {
	fmt.Println("Flying...")
}
