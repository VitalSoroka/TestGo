package model

import(
	//"test/interfaces"
	"fmt"
)

type Sparrow struct{
	weight float64

}

func (this *Sparrow) Sing(){
	fmt.Println("chick chirick")
}

