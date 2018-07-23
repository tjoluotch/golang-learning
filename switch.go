package main

import "fmt"
import "time"

func main() {

	i := 2
	// can intergrate varibales with Print() mentod
	fmt.Print("Write ", i, " as ")
	// basic switch
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	// can use commas to seperate multiple expressions in the same case statement
	case time.Saturday, time.Sunday:
		fmt.Println("Its the weekend")
	// default case used - self explanantory
	default:
		fmt.Println("Its a weekday")
	}

	t := time.Now()
	//switch without an expression is a =nother way to express if/else logic

	switch {
	// shows that case expressions can be non constants
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	default:
		fmt.Println("Its after noon")
	}

	// declare variable whatAmI as a function that takes
	// i which is of type interface as a parameter
	whatAmI := func(i interface{}) {
		// a type switch compares types instead of values
		// can use this to discover the type of an interface value
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
