package main

import "fmt"


func main(){
	fruits := []string {"water" , "pineapple"}

	fmt.Println(fruits)

	numbers:= []int{1,2,5}

	fmt.Println(numbers)

	//how to create slice
		//mth1 using the syntax
		boys:=[]string{"sulaimon" , "jibril"}
		fmt.Println(boys)

		//using the make()
		//make(type , length , capacity)  Capacity is the current allocated space in memory b4 d slice needs to grow 

		slice1:= make([]int, 5)
		slice2:= make([]string, 4 , 8)

		fmt.Println(len(slice1) , slice2)
		fmt.Println(cap(slice1))

		//from an array

		arr:= [5]int{1,2,4,6,7}

		arr1:= arr[2:4]

		fmt.Println(arr1)

		//\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\/\



	//adding to a slice using append()
		num:=[]int{2,4,5}
		numb:=[]int{7,8}
		num = append(num, 3)
		num = append(num, 5)

		num = append(num, numb...)
		fmt.Println(num)


	//iterating over slice
		// Using range
		fruitss:= []string {"apple" , "mango" , "banana"}

		for index, item := range fruitss{


			// fmt.Println(index , item)
			fmt.Printf("index: %d, item:%s\n " , index , item)
		}

		fmt.Println(fruitss)


		//Using for loop

		for i:=0 ; i<len(fruitss) ; i++ {
			fmt.Println(i , fruitss[i])
		}

		//implementing structures with "value type"(which contains the data itself) and "refrence type"(only points to the position of the data in the memory)
		arrrr:= [3]int{1,2,3}

		brrr:= arrrr

		brrr[0]= 6

		fmt.Println(arrrr, brrr)

		digit:=[]int{5,7,9}

		git:= digit

		git[2]= 5

		fmt.Println(digit , git)

		annum:= []int {2,9}
		fmt.Println(takeArray(annum))
}

 
func takeArray(s []int)[]int{
	s = append(s, 100)
	return  s
}