package main

// ---[1]--write golang program that take two number and
//func MulNumber(num1, num2 int64) int64 {
//return (num1 * num2)
//}
//---[2]---write golang program that check given number is even or odd
//func IsEven(num int64) bool {
//return (num%2 == 0)
//}
//---[3]---implemement function that return maximum of three number

/*func MaximumNumber(num1, num2, num3 int64) int64 {
	if num1 > num2 {
		if num1 > num3 {
			return num1
		} else {
			return num3
		}
	} else {
		if num2 > num3 {
			return num2
		} else {
			return num3
		}
	}
}*/
//-----[---4---]-write function that return given year is leap year or not---
/*func isleapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return true
	}
	return false
}*/
//----[---5---]---write function that return area of circle

//func AreaOfCircle(radius float64) float64 {
//formula =pi*r2
/*area*/
//return math.Pi * radius * radius //math.pi its inbuilt function
//return area
//}
//---[---6---]----------------------------------------
/*write function that convert temperature
celcius to fahrenheit voice versa*/
/*func CelciusTofahrenheit(celcius float64) float64 {
	fahrenheit := (celcius * 9 / 5) + 32 //formula to convert celcius to fahrenheit
	return fahrenheit
}
func fahrenheitToCelcious(fahrenheit float64) float64 {
	Celcious := (fahrenheit - 32) * 5 / 9 //formula to  convert fahrenhit to calcious
	return Celcious
}*/
//----[---7---]---------------------------------------
//write golang fuction that calculate and return power of number
/*func PowerOfNumber(base, exponent int) int {
	//2 3 = 8 sum = 2*2*2
	power := 1
	for start := 1; start <= exponent; start++ {
		power = power * base
	}
	return power
}*/
//--[---8---]---write funtion that check given squar is perfect squar or not

/*func IsPerfectSqrt(squar int) bool {
	root := int(math.Sqrt(float64(squar)))
	return (root*root == squar)
}*/

func Mainfourteen() {
	/*--[---1--]----------------[---1---]---------------
	var num1, num2 int64
	message := "Enter Any Two Number : "
	fmt.Println(message)
	fmt.Scan(&num1, &num2)
	mul := MulNumber(num1, num2)
	fmt.Printf("\n Multiplication of %d and %d is =%d", num1, num2, mul)*/
	//--[--2--]-------------------------------------------------------------------------
	/*var num int64
	message := "Enter Any Number : "
	fmt.Println(message)
	fmt.Scan(&num)
	result := IsEven(num)
	if result {
		fmt.Printf("%d is Even ", num)
	} else {
		fmt.Printf("%d is Odd", num)
	}*/

	//---[---3---]------------------------------
	/*var num1, num2, num3 int64
	message := "Enter Your Three Numbers ;"
	fmt.Println(message)
	fmt.Scan(&num1, &num2, &num3)
	result := MaximumNumber(num1, num2, num3)
	fmt.Printf("Maximum Number Between %d %d and %d is = %d", num1, num2, num3, result)*/

	//---[---4---]-------------------------------

	/*var year int
	message := "Enter Any Year"
	fmt.Println(message)
	fmt.Scan(&year)
	result := isleapYear(year)
	fmt.Printf("\n %d is Leap Year =  ", year)
	fmt.Println(result)*/

	//---[5]-------------------------------------
	//---method 1---
	//radius := 5.0
	//area := AreaOfCircle(radius)
	//fmt.Printf("The area of the circle with radius %.2f is %.2f\n", radius, area)
	//}

	//----Method 2-----
	/*var radius float64
	message := "Enter Your Radious "
	fmt.Println(message)
	fmt.Scan(&radius)
	area := AreaOfCircle(radius)
	fmt.Printf("\nRadious of Circle  = %f \n Area of Circle = %f", radius, area)
	fmt.Println()*/
	//---Method 3----
	/*var radius float64
	fmt.Print("Enter your Radius: ")
	fmt.Scan(&radius)

	area := AreaOfCircle(radius)
	fmt.Printf("Area of circle = %.2f\n", area)

	fmt.Printf("Enter the area of circle = %.2f\n", area)
	fmt.Printf("Your radius of circle is = %.2f\n", radius)
	*/

	//---[--6--]-----------------------------------
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

	//---[---7---]-----------------------------------

	/*var base, exponent int
	message := "Enter The Base And Exponent : "
	fmt.Println(message)
	fmt.Scan(&base, &exponent)
	power := PowerOfNumber(base, exponent)
	fmt.Printf("\nPower of the Base %d and Exponent %d = %d", base, exponent, power)
	fmt.Println()*/

	//------[---8---]----------------------------------

	/*var squar int
	message := "Enter squar Number"
	fmt.Println(message)
	fmt.Scan(&squar)
	Perfectsquar := IsPerfectSqrt(squar)
	if Perfectsquar {
		fmt.Printf("\n %d is Perfect Square Number ", squar)
	} else {
		fmt.Printf("\n %d is Not perfect square Number", squar)
	}
	fmt.Println()*/

}
