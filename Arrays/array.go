package main

import (
	"fmt"
	
)

func main(){
var stud [5]int
stud[3]= 4
fmt.Println(stud[0])

var fruit [3]string = [3]string{"orange","apple" , "banana"}
fmt.Println(fruit)

color := [2]string{"white" , "black"}
fmt.Println(color)

ages:=[5]int {1 , 2,4,6,3}

// for i:=0 ; i < len(ages) ; i++ {
// 	fmt.Printf("Index %d:%d\n" , i , ages[i])
// }

for i:=range len(ages) {
	fmt.Printf("Index %d:%d\n" , i , ages[i])
}

for index,item := range ages{

	if item%2 == 0 {
		fmt.Println("wow") 
	}

	if index == 0 {
		fmt.Println("shaka")
	}
}

name:= "zobo"
age:=24
fmt.Printf("%-15s|%10s\n","name","age" )
fmt.Printf("%-15s|%10d\n",name,age)
fmt.Println(name , age)

}




/* Arrays are mainly used when you need a fixed size or for low-level programming.*/