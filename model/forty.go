package model

import (
	//"test/interfaces"
	"fmt"
)

type Forty struct {
	Hight float64
}

func (this Forty) Sing() {
	fmt.Println("kar kar")
}
