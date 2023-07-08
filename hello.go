package main

import (
	"fmt"
	"test/utils"
	 "strconv"
)

type person struct{
	name string
	age int64

}
func (this *person) pretyPrint() {
	json := "{\n\t\"name\" : \"" + this.name + "\"\n\t\"age\" : " + strconv.Itoa(int(this.age)) + "\n}"
	fmt.Println(json)
}


func main() {
	fmt.Println("Hello, Arch Linux")
	fibonachi(100)
	fmt.Println(utils.Factorial(6))
	//fmt.Println(a)
}

func fibonachi(a int){
	previous := 1
	current := 1
	fmt.Println(previous)
	for i:=1; i < a; i++ {
		fmt.Println(current)
		temp := current
		current = previous + current
		previous = temp
	}

	var person person = person{name : "Vital Saroka", age : 24}
	person.pretyPrint()
}
