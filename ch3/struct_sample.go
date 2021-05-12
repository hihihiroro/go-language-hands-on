package main

import (
	"fmt"
)

// Mydata is structure.
type Mydata struct {
	Name string
	Data []int
}

//var mydata struct {
//	Name string
//	Data []int
//}

func main() {
	//	mydata.Name = "Taro"
	//	mydata.Data = []int{10, 20, 30}
	//	fmt.Println(mydata)
	taro := Mydata{"Taro", []int{10, 20, 30}}
	hanako := Mydata{
		Name: "Hanako",
		Data: []int{90, 80, 70},
	}
	fmt.Println(taro)
	fmt.Println(hanako)
}
