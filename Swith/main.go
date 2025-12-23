package main

// switch is to check a value against multiple possibilities and execute different code depending on which one matches
import (
	"fmt"
	"time"
	// "golang.org/x/text/number"
)


func main(){
	fmt.Println("im coming")
	t := time.Now()
	fmt.Println(t)
	
	day := 3 
	switch day {
	case 1:
		fmt.Print(6)
	case 2 :
		fmt.Println(4)
	}

	switch time.Now().Weekday(){
	case time.Friday , time.Saturday,time.Sunday : fmt.Println("its friday") 
	default: fmt.Println("its week days")
	} 

	fmt.Println(traffic("red"))
	fmt.Println(traffic("yellow"))
	fmt.Println(traffic("green"))
	fmt.Println(traffic("white"))

	fmt.Println(grade(100))

	fmt.Println(dayCheck(2))
} 


//Write a function that takes a traffic light color (string: "red", "yellow", "green") and returns what action a driver should take. 
// Handle invalid colors with a default case.


func traffic( colour string) string {
	switch colour {
	case  "red" : return ("stop")
	case "yellow":return ("get ready")
	case "green" : return ("go and go")

	default : return ("invalid colour")

	}

}

//Create a function that takes a numerical score (0-100) and returns the letter grade:

func grade( score int) string {	
	switch{
	case score >= 90 :
		 return ("A")
	case score >=80 :
		return ("B")
	case score >= 70 :
		return ("C")
	case score >= 60:
		return ("D")

	default :
		return ("F")
	}

}



//Write a function that takes a day number (1-7, where 1 is Monday) and returns:
/*"Work from office" for Monday and Wednesday
"Work from home" for Tuesday and Thursday
"Regular work day" for Friday
"Weekend" for Saturday and Sunday
"Invalid day" for anything else

Use multiple values in your cases where appropriate.*/

func dayCheck(arg int ) string {

	
		switch  arg{
		case 1, 3:
			return "Work from office"
		case 2, 4:
			return  "Work from home"	
		case 5:
			return "Regular work day"
		case 6, 7:
			return "Weekend"
		}

	return "Invalid day"
}

/* KNOWING when to add an expression to a switch or not

Quick tip: If you're using >=, <=, &&, || in your cases, you probably need switch without an expression!,
i.e when you have different conditions to check

 Switch WITH Expression (Testing a specific value)
 when comparing one variable against different values

*/
