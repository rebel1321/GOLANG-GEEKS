package main

import (
	"fmt"
)

func main() {
//Basic
// fmt.Println("Basic for loop")
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("Iteration:", i)
// 	}

// //for as while
// fmt.Println("For as while")
// x := 5
// for x<10{
// 	fmt.Println(x)
// 	x++;
// }
// //Infinite loop
// fmt.Println("Infinite loop ")
// x=0
// for{
// 	if x>10{
// 		break
// 	}
// 	fmt.Println(x)
// 	x++
// }
// fmt.Println("range for loops")
// numbers:=[]int{1,2,3}
// for idx,val := range numbers{
// 	fmt.Println(idx,val)
// }


//basic if else
// x:=5
// if x < 10 {
// 	fmt.Println("x is less than 10")
// }
// fmt.Println("x is greater than or equal to 10")

// // if else if else
// if x < 5 {
// 	fmt.Println("x is less than 5")
// } else if x < 10 {
// 	fmt.Println("x is less than 10 but greater than or equal to 5")
// } else {
// 	fmt.Println("x is greater than or equal to 10")
// }
// //Short statement if
// if _,e:=strconv.Atoi("adsgf");e!=nil{
// 	fmt.Println(e)
// }

//Basic switch statement

day:=2
switch day{
case 1:
	fmt.Println("It's a Monday")
case 2:
	fmt.Println("It's a Tuesday")

default:
	fmt.Println("Any other day")
}
//Multiple switch statement
switch day{
case 1,2,3:
	fmt.Println("It's a week start")
case 4,5:
	fmt.Println("Almost end of week")

default:
	fmt.Println("weekand")
}

//type switch statement
var x interface{}=6
switch x.(type){
case int:
	fmt.Println("X is integer")
case string:
	fmt.Println("X is string")
default:
	fmt.Println("I don't know what type it is")
}
}