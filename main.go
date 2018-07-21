package main

import (
	"fmt"
)

func init() {
}

func main() {
	// const (
	// 	myName     = "Pongsathorn"
	// 	myLastName = "Roonbong"
	// )
	// a := "string"
	// b := 10
	// c := 30.4
	// d := true
	// fmt.Println(a, b, c, d)
	// if a == "string" {
	// 	fmt.Printf("%s is stringify\n", a)
	// }
	// if aaa := 10; aaa < 100 {
	// 	fmt.Println("A is less than 100")
	// }
	// a := [...]int{1, 2, 3, 4, 5}
	// b := []int{1, 2, 3, 4, 5, 6, 7}
	// fmt.Println(a)
	// fmt.Println(b)

	// for _, bx := range b {
	// 	fmt.Println(bx)
	// }
	// b = append(b, 10)
	// a := 10
	// b := 20
	// switch {
	// case a > 10:
	// 	fmt.Println("a is greater than 10")
	// case b > 10:
	// 	fmt.Println("b is greater than 10")
	// default:
	// 	fmt.Println("I don't know how to Go")
	// }
	// defer fmt.Println("middle")
	// defer fmt.Println("in the")
	// fmt.Println("Caught")
	// var a int
	// defer func() {
	// 	a = 20
	// 	}()
	// fmt.Println(a)
	// a := make(map[string]int){
	// 	"one": 1,
	// 	"two": 2
	// 	"three": 3
	// }
	// me := Person{
	// 	FirstName: "ter",
	// 	LastName:  "najaa",
	// 	Age:       24,
	// }

	var people []Person
	people = append(people,
		Person{
			FirstName: "ter",
			LastName:  "najaa",
			Age:       24,
		})
	people = append(people,
		Person{
			FirstName: "terter",
			LastName:  "najakonde",
			Age:       26,
		})
	for _, v := range people {
		// res, _ := json.Marshal(v)
		// fmt.Printf("%+v \n", res)
		v.GrowUp()
		fmt.Println(v.FullName())
	}

}

// FullName ...
func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s, %d", p.FirstName, p.LastName, p.Age)
}

// GrowUp ...
func (p *Person) GrowUp() {
	p.Age = p.Age + 1
}

// Person struct
type Person struct {
	FirstName string
	LastName  string
	Age       int
}
