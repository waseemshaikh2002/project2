package main

// Write a go programme to make  calculator
func Maineleven() {

	/*var num1, num2 float64
	var operator string

	fmt.Println("Enter Your First Number : ")
	fmt.Scan(&num1)

	fmt.Println("Enter Your Second Number : ")
	fmt.Scan(&num2)

	fmt.Println("Choose Operator ( + , - , * , / ) ")
	fmt.Scan(&operator)

	var result float64
	switch operator {
	case "+":
		result = num1 + num2
		fmt.Printf("Result = %.2f ", result)
	case "-":
		result = num1 - num2
		fmt.Printf("Result = %.2f ", result)
	case "*":
		result = num1 * num2
		fmt.Printf("Result = %.2f ", result)
	case "/":
		if num2 != 0 {
			result = num1 / num2
			fmt.Printf("Result = %.2f ", result)
		} else {
			fmt.Println("Error : cant Division By zero")
		}

	default:
		fmt.Println("Error : Invalid Operator")
	}*/

	//--------- switch case practice ------//
	//------1) basic switch case ------//

	/*switch 1 {
	case 2:
		fmt.Println("One")
	case 3:
		fmt.Println("one")
	case 1:
		fmt.Println("One")
	}*/

	//............( Multiple value in case ).......

	/*switch 1 {
	case 2, 6, 5:
		fmt.Println("Case one ")
	case 1, 8, 4:
		fmt.Println("case Two ")
	}*/

	//---------------------------------
	//---(switch with initialization)---

	/*switch salary := 500; salary {
	case 600:
		fmt.Println("salary is 600 ")
	case 500:
		fmt.Println("salary is 500")
	}*/

	//----(Switch with Variable Assignment)---

	/*var age int
	switch age = 10; age {
	case 10:
		fmt.Println("Salary is 600")
	case 20:
		fmt.Println("salary is 7003")
	}*/

	//------(Switch with Default)----

	/*var age int
	switch age = 100; age {
	case 10:
		fmt.Println("Age is ten")
	case 20:
		fmt.Println("Age is twenty ")
	default:
		fmt.Println("This is default case")
	}*/

	//--------(Empty with Empty)----

	/*switch 2 {

	}*/

	//---------(Switch with Only default)----

	/*switch 2 {
	default:
		fmt.Println("Welcome is switch's default case")
	}*/

	//--------(function call in switch)--------

	//w.g.p check given number is even or odd

	/*var num int
	fmt.Println("Please Enter Any Number")
	switch fmt.Scan(&num); num % 2 {
	case 0:
		fmt.Println("Number is Even")
	case 1:
		fmt.Println("Number is Odd")
	}*/

	//-----(Case with condition)-------

	//w.g.p to check given number is even or odd
	/*var num int
	fmt.Println("Please Enter Your Number : ")
	switch fmt.Scan(&num); {
	case (num % 2) == 0:
		fmt.Println("Number is Even ")
	case (num % 2) == 1:
		fmt.Println("Number is Odd")
	}*/

	//-------(Switch with fallthrough)-----

	/*switch 2{
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
		fallthrough
	default:
		fmt.Println("Default case")
	}*/

	//--------(Switch With Expression)------

	// w.g.p to create season detector

	// 1 2 3 4 5 6 7 8 9 10
	// 1 summer , fail , winter

	//6 7 8 9 => fail

	//10 11 12 1=> winter

	// 2 3 4 5 => summer

	/*var month int
	fmt.Println("Enter the Month")
	switch fmt.Scan(&month); month {
	case 6, 7, 8, 9:
		fmt.Println("season of fall")
	case 10, 11, 12, 1:
		fmt.Println("season of Winter ")
	case 2, 3, 4, 5:
		fmt.Println("Season of Winter ")
	default:
		fmt.Println("Invalid Month ")

	}*/

	//Menu selector : you can perform any operation by giving menu option
	/*var option int
	for {

		fmt.Println("Menu:")
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Println("5. Exit")

		fmt.Print("Enter your choice: ")
		fmt.Scanln(&option)

		switch option {
		case 1:
			var num1, num2 float64
			fmt.Print("Enter first number: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter second number: ")
			fmt.Scanln(&num2)
			result := num1 + num2
			fmt.Printf("Result: %.2f\n", result)
		case 2:
			var num1, num2 float64
			fmt.Print("Enter first number: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter second number: ")
			fmt.Scanln(&num2)
			result := num1 - num2
			fmt.Printf("Result: %.2f\n", result)
		case 3:
			var num1, num2 float64
			fmt.Print("Enter first number: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter second number: ")
			fmt.Scanln(&num2)
			result := num1 * num2
			fmt.Printf("Result: %.2f\n", result)
		case 4:
			var num1, num2 float64
			fmt.Print("Enter first number: ")
			fmt.Scanln(&num1)
			fmt.Print("Enter second number: ")
			fmt.Scanln(&num2)
			if num2 == 0 {
				fmt.Println("Error: Cannot divide by zero")
			} else {
				result := num1 / num2
				fmt.Printf("Result: %.2f\n", result)
			}
		case 5:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please select a valid option.")
		}
	}*/

	//Shape Aria Calculator : circle , square , triangle

	/*var choice int
	fmt.Println("Shape Area Calculator")
	fmt.Println("1. Circle")
	fmt.Println("2. Square")
	fmt.Println("3. Triangle")
	fmt.Print("Enter your choice (1, 2, or 3): ")
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		fmt.Print("Enter the radius of the circle: ")
		var radius float64
		fmt.Scanln(&radius)
		area := math.Pi * radius * radius
		fmt.Printf("The area of the circle is: %.2f\n", area)
	case 2:
		fmt.Print("Enter the side length of the square: ")
		var side float64
		fmt.Scanln(&side)
		area := side * side
		fmt.Printf("The area of the square is: %.2f\n", area)
	case 3:
		fmt.Print("Enter the base length of the triangle: ")
		var base float64
		fmt.Scanln(&base)
		fmt.Print("Enter the height of the triangle: ")
		var height float64
		fmt.Scanln(&height)
		area := 0.5 * base * height
		fmt.Printf("The area of the triangle is: %.2f\n", area)
	default:
		fmt.Println("Invalid choice!")
	}*/

	//Temprature Convertor : celcious to Fahrenheiit

	//its not understanding

	//File Type Identifier: file extension inpute. print type of file

	/*fileTypeMap := map[string]string{
		"txt": "Text File",
		"pdf": "PDF Document",
		"doc": "Microsoft Word Document",
	}
	fmt.Print("Enter file extension (e.g., txt, pdf, doc): ")
	var fileExt string
	fmt.Scanln(&fileExt)

	fileExt = strings.ToLower(fileExt)

	fileType, found := fileTypeMap[fileExt]
	if found {
		fmt.Printf("File type: %s\n", fileType)
	} else {
		fmt.Println("Unknown file type")
	}*/

}
