package main

import (
	"fmt"
	"strconv"
	"test/interfaces"
	. "test/model"
	"test/utils"
)

type Person struct {
	name string
	age  int64
	real bool
}

func newPerson(name string, age int64) *Person {
	result := Person{name: name, age: age, real: true}
	fmt.Printf("The address of a is: %p\n", &result)

	return &result
}

type PersonBuilder struct {
	person *Person
}

func (this *PersonBuilder) Name(name string) *PersonBuilder {
	if this.person == nil {
		this.person = new(Person)
		this.person.real = true
	}
	this.person.name = name
	return this
}

func (this *PersonBuilder) Age(age int64) *PersonBuilder {
	if this.person == nil {
		this.person = new(Person)
		this.person.real = true
	}
	this.person.age = age
	return this
}

func (this *PersonBuilder) buid() *Person {
	if this.person == nil {
		this.person = new(Person)
		this.person.real = true
	}
	return this.person
}

func (this *Person) pretyPrint() {
	json := "{\n\t\"name\" : \"" + this.name + "\"\n\t\"age\" : " + strconv.Itoa(int(this.age)) + "\n}"
	fmt.Println(json)
}

func main() {
	fmt.Println("Hello, Arch Linux")
	fibonachi(10)
	fmt.Println(utils.Factorial(6))
	var person Person = Person{name: "Vital Saroka", age: 24}

	var birds = [...]interfaces.Bird{&Forty{Hight: 14.5}, &Sparrow{Hight: 21.1}}

	for _, val := range birds {
		//fmt.Println(val)
		val.Sing()
	}

	//var map1 *map[string]string = new(map[string]string)
	//fmt.Println("test map", map1)
	map1 := make(map[string]bool)
	fmt.Println("test map", map1)
	map1["vital"] = true
	fmt.Printf("The address of a is: %p\n", &map1)
	person.pretyPrint()
	var bird interfaces.Bird = &Forty{Hight: 13.3}

	bird.Sing()
	fmt.Println(bird)
	person1 := newPerson("233223", 12)
	fmt.Println(person1.age)
	fmt.Printf("The address of a is: %p\n", &person1)
	person2 := new(PersonBuilder).Name("1212").buid()

	fmt.Println(person2)
	//fmt.Println(a)
}

func fibonachi(a int) {
	previous := 1
	current := 1
	fmt.Println(previous)
	for i := 1; i < a; i++ {
		fmt.Println(current)
		temp := current
		current = previous + current
		previous = temp
	}
}
