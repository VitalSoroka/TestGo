package model

import (
	//"test/interfaces"
	"fmt"
)

type Sparrow struct {
	Hight float64
}

func (this *Sparrow) Sing() {
	fmt.Println("chick chirick")
	fmt.Println(this.Hight)
}
