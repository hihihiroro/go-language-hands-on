package main

import (
	"fmt"
	"hello"
	"strconv"
	"strings"
)

func main() {
	x := hello.Input("Input data")
	ar := strings.Split(x, " ")
	t := 0
	//	for i := 0; i < len(ar); i++ {
	for _, V := range ar {
		//		n, er := strconv.Atoi(ar[i])
		n, er := strconv.Atoi(V)
		if er != nil {
			goto err
		}
		t += n
	}
	fmt.Println("total:", t)
	return

err:
	fmt.Println("ERROR!")
}
