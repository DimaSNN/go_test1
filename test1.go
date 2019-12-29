package test1

import "fmt"


var arr[10] int
var balance = []float32{100, 1.5, 15.2, 11.9 , 1, 0, -10}

var slice []int //срез

func main() {

	
	slice = make([]int,5,5)
	
	fmt.Println(slice)
	var slice2 []float32 =  balance[0:3]
	fmt.Println(slice2)
	

	
	
}
