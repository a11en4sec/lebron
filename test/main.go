package main

import "fmt"

func main() {
	aList := make([]*string,20)
	a := "a1"
	b := "b1"
	c := "c1"
	aList[0] = &a 
	aList[1] = &b 
	aList[2] = &c 

	for i := 0; i < 5; i++ {
		if p, ok := aList[i]; ok {
			fmt.Println(p)

	}
	

	}

	fmt.Println(aList)
}
