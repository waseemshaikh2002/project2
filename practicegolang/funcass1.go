package main

// -----1. Write a program that checks if a given number is even or odd.
/*func IsEven(num int64) bool {
	return (num%2 == 0)
}*/

// ----2. Create a program that determines if a user's input is a positive number.
/*func IsPositive(num int64) bool {
	if num > 0 {
		return true
	} else {
		return false
	}
}*/

//---[---3---]------------------------------------------------
//Build a program that uses a switch statement to categorize a day of the
//week based on a user's input.

/*func CatagerizeOfWeekDay(day string) string {
	switch strings.ToLower(day) {
	case "monday":
		return "WeekDay"
	case "tuesday":
		return "WeekDay"
	case "wednesday":
		return "WeekDay"
	case "Thusday":
		return "WeekDay"
	case "friday":
		return "WeekDay"
	case "saturday":
		return "WeeEnd"
	case "sunday":
		return "WeekEnd"
	default:
		return "unknow"

	}
}*/
//-------------------[---Assignment 4 is below---]-----------
//---------[---5---]----------------------------------------------------
/*func PrintNUmber(number int) {
	for count := 1; count <= number; count++ {
		fmt.Println(count)
	}
}*/

//---[---6---]

//6. Create a program that calculates the factorial of a given number using a
//for loop.

/*
	func CalculateFactorial(num1 int64) int64 {
		var num2 int64
		var multiplication int64
		multiplication = 1
		for num2 = 1; num2 <= num1; num2++ {
		}
		fmt.Printf("Factorial Number :- %d", multiplication)
		return num1

}*/
//----------------------------------------------------------------------------------
//---[8. Write a program that uses a for loop to print the
//Fibonacci sequence up toa certain limit.]
/*func fibonacci(limit int) []int {
	fib := make([]int, 0)
	a, b := 0, 1
	for a <= limit {
		fib = append(fib, a)
		a, b = b, a+b
	}
	return fib
}*/

// ---[---9---Build a program that prints a multiplication table using nested loops.]
/*func MultiplicationTable(limit int) {
	for i := 1; i <= limit; i++ {
		for j := 1; j <= limit; j++ {
			fmt.Printf("%d\t", i*j)
		}
		fmt.Println()
	}
}*/

// ----[---10.Create a pattern using nested loops (e.g., a pyramid of stars).]
// Method 1
/*func PyramidPattern(num int) int {
	for row := 1; row <= num; row++ {
		for colum := 1; colum <= row; colum++ {
			fmt.Printf("%3c", '*')

		}
		fmt.Println("\n")
	}
	return num
}*/

//Method 2
/*func PrintSpaces(spacecount int) {
	for j := 1; j <= spacecount; j++ {
		fmt.Println(" ")
	}
}
func PrintStars(starsCount int) {
	for k := 1; k <= starsCount; k++ {
		fmt.Println(" * ")
	}
}
func StarPyramid(rows int) {
	for i := 1; i <= rows; i++ {
		PrintSpaces(rows - 1)

		PrintStars(2*i - 1)

		fmt.Println()
	}*/

//----[---11---]-----------------------------------------------------
//Write a function to determine if a given year is a leap year.]----

/*func isleapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	}
	return false
}*/
//----------[---12---]-----------------------------------------------

//Create a function that calculates the area of a rectangle and displays
//a message based on the area.

/*func CalculateRectangleArea(length, width float64) float64 {
	return length * width
}
func displayMessege(area float64) {
	if area > 0 {
		fmt.Println("The area of Rectangle is : ", area)
		if area > 100 {
			fmt.Println("This is Large Rectangle.")
		} else {
			fmt.Println("This is Small Rectangle.")

		}
	} else {
		fmt.Println("Invalid Dimension, Area Cannot be calculated.")
	}

}*/

//----------[---13---]------------------------------------------
//Build a program that uses break to exit a loop when a certain
//condition is met.
//Wrong Programe its Not Working

/*func UsesBreaak(num int) {
	for i := 10; i <= num; i++ {
		if num%2 == 0 {
			fmt.Printf("Even Numbers Betwen %d and %d is ", i, num)
			break
		}
	}
}*/

//--------------[---14---]---------------------------------------------

//not got Solution

//--------------[---15---]------------------------------------------------

/*  15)Write a Go program that takes a student's score as input and prints
the corresponding grade (A, B, C, D, or F) based on the following criteria:
A: 90-100
B: 80-89
C: 70-79
D: 60-69
F: 0-59   */

/*func StudentsGrade(marks float64) {
	if marks >= 90 && marks <= 100 {
		fmt.Println("Congratulations You Got A grade")
	} else if marks >= 80 && marks <= 89 {
		fmt.Println("You Got B Grade")
	} else if marks >= 70 && marks <= 79 {
		fmt.Println("You Got C Grade")
	} else if marks >= 60 && marks <= 69 {
		fmt.Println("You Got D Grade")
	} else {
		fmt.Println("You Are Failed")
	}
}*/
//------------[---16----]---------------------------------
/*Implement a program that simulates a traffic light. Use a switch
statement to control the light transitions (red to green, green to yellow,
	yellow to red) and print the current state*/

/*
	func TraficLight(color string) string {
		switch strings.ToLower(color) {
		case "red":
			return "Please Stop signal Is Off"
		case "green":
			return "YOu can Go Signal Is On"
		case "yellow":
			return "you can wait or go slow"
		default:
			return "Unknow"
		}
	}
*/
//--------------------[---17---]-----------------------------
/*Create a program that takes an integer as input and prints its prime
factorization. Implement control flow statements to handle different cases.*/

/*func PrimeFactorization(number int) {
	if number <= 1 {
		fmt.Println("Prime Factorization is Not Possible for Number Less Than or Equal to 1.")
		return
	}
	fmt.Printf("Prime Factorization of %d :", number)
	for number%2 == 0 {
		fmt.Print("2 ")
		number /= 2
	} //Now n must be odd. We can skip even numbers.
	for i := 3; i*i <= number; i += 2 {
		for number%i == 0 {
			fmt.Printf("%d", i)
			number /= i
		}

	}
	if number > 2 {
	fmt.Printf("%d", number)
	}
	fmt.Println()
}*/
//-----------------------[---18---]----------------------------
/*Write a program that simulates user authentication. Prompt the user
for a username and password, and use control flow statements to determine
if the credentials are correct.*/

/*func authentication(userName, passWord string) bool {
	correctUserName := "waseem12345"
	correctPassWord := "Pass@2001"

	if userName == correctUserName && passWord == correctPassWord {
		return true
	}
	return false
}*/

//------------------------[---19---]----------------------

/*
write function that convert temperature
celcius to fahrenheit voice versa
*/
/*func CelciusTofahrenheit(celcius float64) float64 {
	fahrenheit := (celcius * 9 / 5) + 32 //formula to convert celcius to fahrenheit
	return fahrenheit
}
func fahrenheitToCelcious(fahrenheit float64) float64 {
	Celcious := (fahrenheit - 32) * 5 / 9 //formula to  convert fahrenhit to calcious
	return Celcious
}
*/
//------------------[---20---]-------------------------------
/*func playNumberGuessingGame() {
	rand.Seed(time.Now().Unix())

	randomNumber := generateRandomNumber(1, 100)

	fmt.Println("Welcome to Number Guessing Game")
	fmt.Println("Please Picked Any Number Between 1 to 100")
	fmt.Println("Can You Guess What It is")

	for {
		guess := getUserGuess()

		if guess == randomNumber {
			fmt.Println("Congratulations! You Guessed It Right")
			break
		} else if guess < randomNumber {
			fmt.Println("Too Low Try Again")
		} else {
			fmt.Println("Too High Try Again")
		}
	}
}
func generateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}
func getUserGuess() int {
	var guess int
	fmt.Println("Enter Your Guess: ")
	_, err := fmt.Scan(&guess)
	if err != nil {
		fmt.Println("Please Enter a Valid Number")
		return getUserGuess()
	}
	return guess
}*/
//-----------------------[---21---]------------------------------------

/*Create a menu-driven calculator program that supports addition, subtraction,
multiplication, and division. Use switch statements to execute the selected
operation*/

/*func addition(a, b float64) float64 {
	return a + b
}
func subtraction(a, b float64) float64 {
	return a - b
}
func multiplication(a, b float64) float64 {
	return a * b
}
func Divide(a, b float64) float64 {
	if b != 0 {
		return a / b
	}
	return 0
}*/
//------------------------------------------------
//--------------------[---22---]------------------
//------------------------------------------------
// Write a program that determines if a given year is a leap year or not.
//Implement control flow statements to handle the conditions for a leap year.
/*func isleapYear(year int) bool {
	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		return true
	}
	return false

}*/
//------------------------------------------------------------------------------
//------------------------------[---23---]--------------------------------------
//------------------------------------------------------------------------------
/*mplement a program that checks if a given string is a palindrome. Use
control flow statements to compare characters from both ends of the string.*/
/*func isPalindroome(input string) bool {
	input = strings.ToLower(input)

	start := 0
	end := len(input) - 1

	for start < end {
		if input[start] != input[end] {
			return false

		}
		start++
		end--
	}
	return true

}*/

// -----------------------------------------------------------------------------
//--------------------------------------------------
//-------------------------[---24---]-----------------
//Build a program that takes a user's age as input and prints whether the user
//is eligible to vote or not based on the legal voting age.

/*func isEligible(age int) bool {
	const ligalVoteAge = 18
	return age >= ligalVoteAge
}*/

/*func isvalidEmail(email string)bool{
	if !strings.Contains(email,"0"){
		return false
	}
	parts := strings.Split(email,"0")
	localpart := parts[0]
	domainpart := parts[1]

	if len(domainpart) ==0 || !strings.Contains(domainpart,"."){
		return false
	}
	return true

}*/
//------------------------------------------------------------
//-------------------[---25---]---------------------------------
//-------------------------------------------------------------
//Write a program that validates if a given email address is in the correct
//format. Use control flow statements to check for the presence of "@" and a valid
//domain

/*func isvalidEmail(email string)bool{
atIdex := strings.Index(email,"@")
if atIdex== -1 {
	return false
}
if strings.LastIndex(email,"@") !=atIdex{
	return false
}
localpart := email[:atIdex]
domainpart := email[atIdex+1:]

if localpart == "" || domainpart == "" {
	return false
}
dotIndex == -1 || dotIndex == 0 || dotIndex == len(domainpart)-1{
	return false
}
return true

}
*/

func Assignmentmain() {

	//--------[--1--]--------------------------------
	/*var num int64
	message := "Enter Your Number"
	fmt.Println(message)
	fmt.Scan(&num)
	result := IsEven(num)
	if result {
		fmt.Printf("%d is Even Number", num)
	} else {
		fmt.Printf("%d is Odd Number", num)
	}*/

	//---[---2---]-------------------------------------

	/*var num int64
	  message := "Enter Your Number :"
	  fmt.Println(message)
	  fmt.Scan(&num)
	  result := IsPositive(num)
	  if result {
	  	fmt.Printf("%d is Positive Number ", num)
	  } else {
	  	fmt.Printf("%d is Negative Number ", num)
	  }*/

	//----[---3--]----------------------------------------------]

	/*var input string
	  message := "Enter The Day"
	  fmt.Println(message)
	  fmt.Scan(&input)
	  catagerize := CatagerizeOfWeekDay(input)
	  fmt.Printf("\n %s is a %s  \n", input, catagerize)*/

	/*var input string
	  fmt.Println("Enter The day of Week")
	  fmt.Scan(&input)

	  catagerize := CatagerizeOfWeekDay(input)
	  fmt.Printf("\n %s is %s \n", input, catagerize)*/

	//---[---4---]---------------------------------------------------------------
	//--Implement a calculator using a switch statement for different arithmaticoperation.

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
	//----[5]---------------------------------------------------------
	/*number := 10
	  PrintNUmber(number)*/
	//---------------------------------------------------------------------
	//------------------[6]------------------------------------------------
	/*var num1 int64
	  fmt.Println("Enter YOur Number:")
	  fmt.Scan(&num1)
	  CalculateFactorial(num1)*/
	//----------[---8---]--------------------------------------------------------
	// Function to generate Fibonacci sequence up to a certain limit
	/*limit := 100 // Change this to set the limit for Fibonacci sequence
	  fibSeq := fibonacci(limit)
	  fmt.Printf("Fibonacci sequence up to %d:\n", limit)
	  for _, num := range fibSeq {
	  	fmt.Printf("%d ", num)
	  }
	  fmt.Println()*/

	/*limit := 100
	  fibSeq := fibonacci(limit)

	  fmt.Printf("Fiboacci sequence up to %d : \n", limit)
	  for _, num := range fibSeq {
	  	fmt.Printf(" %d ", num)
	  }
	  fmt.Println()*/
	//-----------[---6---]--------------------------------------------------------------
	/*limit := 10
	  MultiplicationTable(limit)*/
	//---------------------------------------------------------------------------------
	//-------[---10---]
	/*var num int
	  fmt.Println("Enter Any Number")
	  fmt.Scan(&num)
	  PyramidPattern(num)*/

	//---------[---11---]---------------------------------------

	/*var year int
	  message := "Enter Any Year"
	  fmt.Println(message)
	  fmt.Scan(&year)
	  result := isleapYear(year)
	  fmt.Printf("\n %d is Leap Year = ", year)
	  fmt.Println(result)*/

	//---------------[---12---]--------------------------------------

	/*var length, width float64
	  fmt.Println("Enter the Length of rectangle :")
	  fmt.Scan(&length)
	  fmt.Println("Enter the Width of Rectangle :")
	  fmt.Scan(&width)

	  rectanglearea := CalculateRectangleArea(length, width)
	  displayMessege(rectanglearea)*/
	//------------------[---13---]------------------------------------

	//num := 100
	//UsesBreaak(num)

	//------------------[---14---]-----------------------------------------
	//not Got Solution

	//------------------[---15---]----------------------------------------
	/*var marks float64

	  fmt.Println("Enter Any Number : ")
	  fmt.Scan(&marks)

	  StudentsGrade(marks)*/

	//----------------[---16---]-------------------------------

	/*var color string
	  fmt.Println("Enter Any color 1)Red 2)Green 3)Yello")
	  fmt.Scan(&color)

	  message := TraficLight(color)
	  fmt.Println(message)*/
	//-----------------[---17---]--------------------------------------
	/*var number int
	  fmt.Print("Enter Any Interger Number :")
	  _, err := fmt.Scan(&number)
	  if err != nil {
	  	fmt.Println("Error Reading Input :", err)
	  }
	  PrimeFactorization(number)*/

	//-----------[---18---]-----------------------------

	/*var userName, passWord string
	  fmt.Println("Enter Your User Name : ")
	  fmt.Scan(&userName)

	  fmt.Println("Enter Your PassWord : ")
	  fmt.Scan(&passWord)

	  if authentication(userName, passWord) {
	  	fmt.Println("Authentication Successful ! Congratulations ", userName)
	  } else {
	  	fmt.Println("Authentication Failed, Invalid UseName Or PassWord")
	  }*/
	/*var temperature float64
	  message := "Enter Temprerature"
	  fmt.Println(message)
	  fmt.Scan(&temperature)
	  fmt.Println("1.Celcius To Fahrenhit")
	  fmt.Println("2.Fahrenhit To Celcious")
	  fmt.Println("Enter Your Choice")
	  var choice int
	  fmt.Scan(&choice)
	  switch choice {
	  case 1:
	  	result := CelciusTofahrenheit(temperature)
	  	fmt.Printf("\nTemperature Celcius To Fahrenhit : %f", result)
	  case 2:
	  	result := fahrenheitToCelcious(temperature)
	  	fmt.Printf("\nTemperature Fahrenhit To celcius : %f", result)
	  default:
	  	fmt.Println("Invalid Choice")
	  }
	  fmt.Println()*/

	//--------------[---20---]----------------------------------------

	//playNumberGuessingGame()

	//--------------------------------------------------------------
	//---------------[21]--------------------------------------------

	/*var choice int
	var num1, num2 float64

	fmt.Println("Welcome In Calculator")
	fmt.Println("1: Addition.")
	fmt.Println("2: Substraction.")
	fmt.Println("3: Multiplication")
	fmt.Println("4: Divide")
	fmt.Println("Please Enter You choice [1-4]")
	fmt.Scan(&choice)

	fmt.Print("Please Enter Your First Number: ")
	fmt.Scan(&num1)
	fmt.Print("Please Enter Your Second Number: ")
	fmt.Scan(&num2)

	switch choice {
	case 1:
		fmt.Printf("Result : %2f/n", addition(num1, num2))
	case 2:
		fmt.Printf("Result : %2f/n", subtraction(num1, num2))
	case 3:
		fmt.Printf("Result : %2f/n", multiplication(num1, num2))
	case 4:
		fmt.Printf("Result : %2f/n", Divide(num1, num2))
	default:
		fmt.Println("Invalid Choice")

	}*/
	//--------------------------------------------
	//----------------------[---22---]-------------
	/*var year int
	fmt.Print("ENter Any Year:")
	fmt.Scan(&year)

	if isleapYear(year) {
		fmt.Printf("%d is a leap Year,\n", year)
	} else {
		fmt.Printf("%d Is Not a Leap Year,\n", year)
	}*/
	//----------------------------------------------------------
	//------------------------[---23---]---------------------

	/*var input string
	fmt.Println("Enter a String: ")
	fmt.Scan(&input)

	if isPalindroome(input) {
		fmt.Printf("%s Is Palindrome, \n", input)
	} else {
		fmt.Printf("%s is not a Palindrome \n", input)
	}*/

	//----------------------------------------------------------------
	//----------------------[---24---]--------------------------------

	/*var age int
	fmt.Print("Enter Your Age: ")
	fmt.Scan(&age)
	if isEligible(age) {
		fmt.Println("You Are Eligible to vote!...")
	} else {
		fmt.Println("You Are Not Eligible to vote!...")
	}*/

	//-------------------------------------------------------------------------
	//----------------------[---25---]-----------------------------------------
	/*emails:=[]string{
		"nawabwaseem53@gmail.com","nontech63@gmail.com","sigmarules1@gmail.com",
	}
	for _, email := range emails{
		if isValidEmail(email){
			fmt.Printf("%s is valid Email address \n",email)
		}else{
			fmt.Println("%s is Invalid Email address \n",email)
		}
	}*/
}
