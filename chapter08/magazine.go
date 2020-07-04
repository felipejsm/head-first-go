package magazine

import "fmt"

type Subscriber struct {
	Name        string
	Rate        float64
	Active      bool
	HomeAddress Address
}
type Employee struct {
	Name        string
	Salary      float64
	HomeAddress Address
}
type Address struct {
	Street     string
	City       string
	State      string
	PostalCode string
func main() {
	sub := defaultSubscriber("Amy Poehler")
	applyDiscount(sub)
	printInfo(sub)

	sub2 := defaultSubscriber("Leslie Knope")
	applyDiscount(sub2)
	printInfo(sub2)

	var e Employee
	e.Name = "Thomas Shelby"
	e.Salary = 60000
	fmt.Println(e.Name)
	fmt.Println(e.Salary)

}
func printInfo(sub *Subscriber) {
	fmt.Println("Name:", sub.Name)
	fmt.Println("Rate:", sub.Rate)
	fmt.Println("Address:"sub.HomeAddress)
	fmt.Println("Active?", sub.Active)
}

func defaultSubscriber(name string) *Subscriber {
	var s Subscriber
	s.Name = name
	s.Rate = 5.99
	s.Active = true
	return &s
}
func applyDiscount(s *Subscriber) {
	s.Rate = 3.99
}
