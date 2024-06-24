package main

func Mainfirst() {
	//fmt.Println("welcome in golang")
	//fmt.Println("hello world")

	/*fmt.Printf("My Name is = %s", "waseem")
	  fmt.Printf("\nMy Age is = %d", 60)
	  fmt.Printf("\n My salary is = %f", 6000.20)
	  fmt.Printf("\nHe  is Not Good Boy = %t", false)
	  fmt.Printf("\nFirst Letter is = %c", 'w')
	  fmt.Printf("\nMy city Name is = %v", "Parli")*/

	//variable decleration in golang//
	//1) Type annotation
	//2) Type infrence

	//1) sysntax:var <variable-name> data type =>
	//type veriable declearation

	/*var city string
	  // variable  declearation using type//
	  city = "pune"
	  fmt.Println(city)*/

	/*var age int
	  age = 60
	  fmt.Println(age)
	  fmt.Printf("My Age is =%d", age)*/

	/*var name string = "sister"
	  var age int = 18
	  var salary float64 = 20000
	  var alphabetName rune = 'A'
	  var isSheGood bool = true
	  var complexNumber complex128 = 2 + 3i
	  fmt.Println(name, "\n", age, "\n", salary, "\n", alphabetName, "\n", isSheGood, "\n", complexNumber)*/
	//constant Type Variable Declaration//

	/*const address string = "waseem"
	  const gender string = "male"
	  const height float64 = 6.2
	  const isOkey bool = true
	  fmt.Println(address, "\n", gender, "\n", height, "\n", isOkey)*/

	// DECLARE CONSTANT USING TYPE INFERENCE

	// syntax : const <variable-name> = value

	/*const name = "Pinky"
	  const age = 56
	  const salary = 800.36
	  const isFair = false
	  const firstChar = 'P'
	  const complexNumber = 45 + 21i
	  fmt.Printf("%c\n", firstChar)
	  fmt.Println(name, age, salary, isFair, firstChar, complexNumber)*/

	/*var name1 string = "wellcome in codelangs"
	  fmt.Printf("\n**********(%s)************", name1)
	  fmt.Printf("\n+++++++++++++========={%s}==========+++++++++++", name1)
	  var hlo, dho string = "Hello", "dude....How are you....!"
	  fmt.Printf("\n%s ***((((((( %s", hlo, dho)

	  var name string = "raju"
	  var age int = 60
	  var grade rune = 'c'
	  fmt.Printf("\nMy Name is = %s", name)
	  fmt.Printf("\nMy age is = %d", age)
	  fmt.Printf("\nMy grade is = %c", grade)*/

	/*var plus string = "+++++++++++"
	  var roundBracket string = ")))))))))((((("
	  var wlcm string = "Wellcome"
	  var dot string = "........"
	  var star string = "*******"
	  fmt.Printf("%s%s%s%s%s%s", plus, roundBracket, wlcm, dot, wlcm, star)*/
	/*var name4 string = "ikbal"
	  var age int = 21
	  var height1 float64= 5.7
	  var skincolour string = "brown"
	  var education rune = 'a'*/

	//short variable//
	//syntax : <variable name > ;= value//

	//waseem := "Engineer"
	//fmt.Printf("Waseem is an %s", waseem)

	//(2)
	/*name := "akshay"
	  age := 25
	  address := "katraj"
	  height := 6.2
	  salary := 25000
	  single := "yes"
	  fmt.Printf("My Name is %s\n i am %d year old\n i live in %s\n my height is %f\n my salary is %d\n i am single", name, age, address, height, salary, single)*/

	//variable decleration using type inference
	//var gender string= "male"
	//var address = 40
	//var adress int = 40<------- internal type inference

	//var salary ="pinky"<------ variable using type inference
	//var salary string = "pinky"<------- internal type inference

	//constant//
	/*const name = "waseem"
	  fmt.Printf("Name : %s", name)*/
	/*const name1 <-------- it will show error because  i an not given value
	  name1="waseem"
	  fmt.Printf("name : %s",&name1)*/

	//Multiple variable declaration in golang

	/*	var num1, num2, addition, mul, substaction, division int64

		fmt.Println("Enter your first numbeer : ")
		fmt.Scan(&num1)

		fmt.Println("Enter your second number : ")
		fmt.Scan(&num2)

		addition = num1 + num2
		mul = num1 * num2
		substaction = num1 - num2
		division = num1 / num2

		fmt.Printf("Addition = %d \n Multiplication = %d \n Substraction = %d \n Division = %d ", addition, mul, substaction, division)*/

	//same programme using short hand variable and if else condition

	/*var num1, num2 int64

	  fmt.Println("Enter the The first Number : ")
	  fmt.Scan(&num1)

	  fmt.Println("Enter the second number : ")
	  fmt.Scan(&num2)

	  add := num1 + num2
	  mul := num1 * num2
	  subs := num1 - num2
	  div := num1 / num2

	  fmt.Printf("Addition = %d \n Multiplication = %d \n Substraction = %d \n Division = %d ", add, mul, subs, div)*/

	/*var number, square, cube int64

	  fmt.Print("Enter your Number  : ")
	  fmt.Scan(&number)

	  square = number * number
	  cube = square * number

	  fmt.Printf("The Sqare  Of Given Number's is = %d \nThe cube of given number's is = %d ", square, cube)*/

	//Even or Odd programme

	/*var number int

	  fmt.Print("Enter you Number : ")
	  fmt.Scan(&number)

	  if number%2 == 0 {

	  	fmt.Println(number, " Is Even")
	  } else {
	  	fmt.Println(number, "Is Odd")
	  }*/

	/*var num1, num2, sum int

	  fmt.Print("Enter your number : ")
	  fmt.Scan(&num1)

	  fmt.Println("Enter your second number : ")
	  fmt.Scan(&num2)

	  sum = num1 + num2
	  fmt.Printf("sum = %d", sum)*/

	//Using Typeof keyword know the data type of variable
	/*sys := "Waseem"
	  dataType := reflect.TypeOf(sys)
	  fmt.Println(dataType)*/

	/*number := 15
	  dataType := reflect.TypeOf(number)
	  fmt.Println(dataType)*/

	/*character := 'A'
	  dt := reflect.TypeOf(character)
	  fmt.Println(dt)*/
	// Out put of this programme ?
	// it will be return int 32 in golang because
	/*This is because the reflect.TypeOf() function returns the reflection Type of the value passed to it.
	  In this case, the variable character is of type rune, which is an alias for int32 in Go.
	  Trherefore, reflect.TypeOf(character) will return int32.*/

	//calculate square root

	// package << math.sqrt()
	// using float Type
	/*var num float64
	  fmt.Println("Please Enter any number")
	  fmt.Scan(&num)
	  //convert integer to float
	  //var floatNum float64 = float64(num)
	  var root = math.Sqrt(num)
	  fmt.Printf("\n Square Root of %.2f  = %.2f", num, root)*/

	// conver intiger to float
	/*var num int
	  fmt.Println("Please Enter any number")
	  fmt.Scan(&num)
	  var floatNum float64 = float64(num)
	  var root = math.Sqrt(floatNum)
	  fmt.Printf("\n Square Root of %.2f  = %f", floatNum, root)*/

	//blank identifier

	/*var city string = "pune"
	  _ = city

	  _ = 16
	  fmt.Println(_) */
	// black identifier cant print

	//short variable

	/*age := 10
	  id := 101
	  fmt.Println(age, id)*/

	//Read the data from bufio line of text
	/*var line string
	  var age string
	  var city string
	  var address string
	  var mobilenumber string

	  reader := bufio.NewReader(os.Stdin)
	  fmt.Println("Please enter your slogan : ")
	  line, _ = reader.ReadString('\n')
	  fmt.Println("Enter your age : ")
	  age, _ = reader.ReadString('\n')
	  fmt.Println("Enter your city : ")
	  city, _ = reader.ReadString('\n')
	  fmt.Println("Enter your Adress :  ")
	  address, _ = reader.ReadString('\n')
	  fmt.Println("Enter your Mobile Number")
	  mobilenumber, _ = reader.ReadString('\n')

	  fmt.Println("**********************************")

	  fmt.Println()
	  fmt.Printf("\n slogan : %s", line)
	  fmt.Printf("\n Age : %s", age)
	  fmt.Printf("\n city : %s", city)
	  fmt.Printf("\n Address : %s", address)
	  fmt.Printf("\n Mobile Number : %s", mobilenumber)*/

	/*var name, price, weight, color, ssd, ram, osType string

	  scanner := bufio.NewScanner(os.Stdin)

	  fmt.Println("Enter the laptop name HP/DLL/LENOVO :")
	  scanner.Scan()
	  name = scanner.Text()
	  if name == "hp" {
	  	fmt.Println("YOU Choose Hp Laptop...")
	  } else if name == "dell" {
	  	fmt.Println("you Selected Dell laptop...")
	  } else if name == "lenovo" {
	  	fmt.Println("You selected Lenovo laptop...")
	  } else {
	  	fmt.Println("You Entered Wrong choice")
	  }

	  fmt.Println("Enter the laptop Price 40k/50k/60k : ")
	  scanner.Scan()
	  price = scanner.Text()
	  if price == "40k" {
	  	fmt.Println("You selected 4ok's Laptop...")
	  } else if price == "50k" {
	  	fmt.Println("You selected 50k' laptop...")
	  } else if price == "70k" {
	  	fmt.Println("You selected 70k's Laptop...")
	  } else {
	  	fmt.Println("You Entered wrong Choice...")
	  }

	  fmt.Println("Enter the Laptop weight 2kg/3kg ; ")
	  scanner.Scan()
	  weight = scanner.Text()
	  if weight == "2kg" {
	  	fmt.Println("You selected 2kg's Laptop...")
	  } else if weight == "3kg" {
	  	fmt.Println("You selected 3kg's Laptop...")
	  } else {
	  	fmt.Println("You Entered Wrong choice...")
	  }

	  fmt.Println("Choose the laptop color Black/Silver/gray : ")
	  scanner.Scan()
	  color = scanner.Text()
	  if color == "black" {
	  	fmt.Println("You selected Balck color...")
	  } else if color == "silver" {
	  	fmt.Println("you selected silver color...")
	  } else if color == "gray" {
	  	fmt.Println("you selected Gray color...")
	  } else {
	  	fmt.Println("You entered Wrong choice")
	  }

	  fmt.Println("Enter the SSD 256/512 : ")
	  scanner.Scan()
	  ssd = scanner.Text()
	  if ssd == "256" {
	  	fmt.Println("You selected 256 SSD Laptop...")
	  } else if ssd == "512" {
	  	fmt.Println("You selected 512 ssd La[top...")
	  } else {
	  	fmt.Println("You entered wrong Choice")
	  }

	  fmt.Println("Enter the Ram 8gb/16gb/32gb : ")
	  scanner.Scan()
	  ram = scanner.Text()
	  if ram == "8gb" {
	  	fmt.Println("You selected 8gb Ram's Laptop...")
	  } else if ram == "16gb" {
	  	fmt.Println("You selected 16bg Ram's Laptop...")

	  } else if ram == "32" {
	  	fmt.Println("You Slected 32gb Ram's Laptop...")
	  } else {
	  	fmt.Println("You entered Wrong Choice")
	  }

	  fmt.Println("Enter the Operating system Type linux/Mac/Windows : ")
	  scanner.Scan()
	  osType = scanner.Text()
	  if osType == "linux" {
	  	fmt.Println("You seelceted Linux OS laptop...")
	  } else if osType == "mac" {
	  	fmt.Println("You selected Mac os Laptop...")
	  } else if osType == "windows" {
	  	fmt.Println("You selected Windwos os Laptop")
	  } else {
	  	fmt.Println("You Entered Wrong Choice")
	  }

	  fmt.Println("\n_____________________________________")
	  fmt.Printf("\n Laptop Name = %s", name)
	  fmt.Println("\n_____________________________________")
	  fmt.Printf("\n Laptop Price = %s", price)
	  fmt.Println("\n_____________________________________")
	  fmt.Printf("\n Laptop Weight = %s", weight)
	  fmt.Println("\n_____________________________________")
	  fmt.Printf("\n Laptop color = %s", color)
	  fmt.Println("\n_____________________________________")
	  fmt.Printf("\nLaptop SSD = %s", ssd)
	  fmt.Println("\n_____________________________________")
	  fmt.Printf("\nLaptop Ram  = %s", ram)
	  fmt.Println("\n_____________________________________")
	  fmt.Printf("\nLaptop Operating System = %s", osType)
	  fmt.Println("\n_____________________________________")
	  fmt.Println("...Congratulations You Purchased Laptop...")*/

	//Negative And Positive Number

	/*var num1 int64

	  fmt.Println("Enter the first Number : ")
	  fmt.Scan(&num1)

	  if num1 >= 0 {
	  	fmt.Println("Positive Number")
	  } else {
	  	fmt.Println("Negative Number")
	  }*/

	//operator
	//unary
	//Plus (+)Operator
	/*var x int = 30
	  var y = +x
	  fmt.Printf("%d", y)*/

	//Unary minus Operator

	//var x int = -30
	//var y = +x
	//fmt.Printf("%d", y)

	//var x = true
	//var y = !x
	//fmt.Println(y)

	//Increment operator (++)

	//var x = 10
	//x++
	//fmt.Println(x)

	//decrement Operator (--)

	//var x = 10
	//x--
	//fmt.Println(x)

	//Addres of operatoer (&)
	/*var x = 100
	  var y = &x
	  fmt.Println(y)*/

	//pointer operator (*)

	//var x = 10
	//var y = &x
	//var valueOfx = *y
	//fmt.Println(valueOfx)

	//***BINARY OPERATOR***

	//Arthmatic operator

	//plus (+)
	/*var x = 10 + 10
	  fmt.Println(x)*/

	//Minus (-)

	//var x = 20 - 10
	//fmt.Println(x)

	//Multiplication

	//var x = 10 * 5
	//fmt.Println(x)

	//Division ( / )

	//var x = 100 / 2
	//fmt.Println(x)

	//Modulus ( % )

	/*var x = 100 % 10
	  fmt.Println(x)*/

	//Comparision Operator

	//equal to (==)

	/*var x, y = 10, 50
	  var z = x == y
	  fmt.Println(z)*/

	//not equal to (1=)

	//var x, y = 10, 50
	//var z = x != y
	//fmt.Println(z)

	//Less than (<)

	//var x, y = 10, 5
	//var z = x < y
	//fmt.Println(z)

	//less than r equal (<=)

	//x, y := 10, 5
	//z := x <= y
	//fmt.Println(z)

	// Greater than

	//x, y := 10, 5
	//z := x > y
	//fmt.Println(z

	//Grenter than or equal

	//x, y := 10, 5
	//z := x >= y
	//fmt.Println(z)

	//**Logical Operator** (&&) And //If give True When Then
	//Both conditions Will be true
	//if pur one conditions is true and second
	//condition is false then it will be heat false

	/*x, y := 10, 20
	  z := x >= y && x == y
	  fmt.Println(z)*/

	//** logical operatoer or ( || )
	//or operator  returns ture only
	//either of operand returns true.
	//otherwise return false

	//x, y := 10, 5
	//z := x >= y || x == y
	//fmt.Println(z)

	//Combine Operator

	//x := 10 > 5 && 20 < 50 || 5 < 10 && 5 > 10 || 50 == 50 && 40 >= 50
	//fmt.Println(x)

	//Assignment Operator

	//calculator
	/*var num1, num2 int64

	  fmt.Println("Enter Your First Number : ")
	  fmt.Scan(&num1)

	  fmt.Println("Enter Your ssecond Number ")
	  fmt.Println(&num2)*/

	//Assignment Operator
	// Assignment operator
	// 1) equal to =
	// 2) plus and assignment +=
	// 3) minus and assignment -=
	// 4) multiplication and assignment *=
	// 5) division and assignment /=
	// 6) modulus and assignment %=

	//operator
	//unary
	//unary plus (+)
	/*var x = 10
	  var y = +x
	  fmt.Println(y)*/

	//Unary minus

	// var x = 10
	// var y = -x
	// fmt.Println(y)

	//3) Increment Operatoe (++)
	// var x int  = 10
	// x++
	// fmt.Println(x)

	//decrement Operator (--)
	// var x int  = 10
	// x--
	// fmt.Println(x)

	//logical not (!)

	//var isSheGood = !true
	//fmt.Println(isSheGood)

	// address operator  (&)

	//var name = "waseem"
	//var x = &name
	//fmt.Println(x)

	// pointer operator (*)

	// var x = 200
	// var y = &x
	// var valueOfName = *y
	// fmt.Println(valueOfName)

	//Binary Operator

	// Arthmatic operators
	//1) plus (+)

	// var x int = 10
	// var y int = 20
	// var z = x + y
	// fmt.Println(z)

	//2) minus (-)

	// var x int = 10
	// var y int = 20
	// var z = x - y
	// fmt.Println(z)

	// division operator (/)

	//Lecture (10)
	//using if condition
	/*var num int
	  fmt.Println("Enter you Number")
	  fmt.Scan(&num)
	  var result string = "its Odd Number"
	  if num%2 == 0 {
	  	result = "Number is Even"
	  }
	  fmt.Printf("%d is = %s ", num, result)*)*/

	// 2) if -- else
	//using if else

	// const num = 5
	// if num%2 == 0 {
	// 	fmt.Println("Number is Even")
	// } else {
	// 	fmt.Println("Number is odd")
	// }

	// if--else--if lader

	/*var marks int
	  fmt.Println("Enter Your Number")
	  fmt.Scan(&marks)
	  if marks >= 75 {
	  	fmt.Println("Distinction")
	  } else if marks >= 65 {
	  	fmt.Println("A")
	  } else if marks >= 45 {
	  	fmt.Println("B")
	  } else if marks >= 35 {
	  	fmt.Println("C")
	  } else {
	  	fmt.Println("Fail")
	  }*/

	//Nested If else

	// const x, y, z = 40, 10, 50
	// if x > y {
	// 	if x > z {
	// 		fmt.Println(x)
	// 	} else {
	// 		fmt.Println(z)
	// 	}
	// } else {
	// 	if y > z {
	// 		fmt.Println(y)
	// 	} else {
	// 		fmt.Println(z)
	// 	}
	// }

	//if else programme w p input number and know its even or odd

	// var num int
	// fmt.Print("Enter your Number :")
	// fmt.Scan(&num)
	// if num%2 == 0 {
	// 	fmt.Println("Even")
	// } else {
	// 	fmt.Println("odd")
	// }

	// --------------- {w.g.p. convert 3 digit amount in word }-------------

	/*var amount int
	  fmt.Println("Please Enter Amount : ")
	  fmt.Scan(&amount)

	  hundreds := amount / 100
	  reminder := amount % 100
	  tens := reminder / 10
	  reminder = reminder % 10

	  if hundreds == 1 {
	  	fmt.Print(" One Hundred")
	  } else if hundreds == 2 {
	  	fmt.Print(" two Hundred")
	  } else if hundreds == 3 {
	  	fmt.Print(" three Hundred")
	  } else if hundreds == 4 {
	  	fmt.Print(" four Hundred")
	  } else if hundreds == 5 {
	  	fmt.Print(" five Hundred")
	  }

	  if tens == 1 {
	  	fmt.Print(" ten")
	  } else if tens == 2 {
	  	fmt.Print(" twenty")
	  } else if tens == 3 {
	  	fmt.Print(" thirty")
	  } else if tens == 4 {
	  	fmt.Print(" fourty")
	  } else if tens == 5 {
	  	fmt.Print(" fifty")
	  } else if tens == 6 {
	  	fmt.Print(" sixty")
	  } else if tens == 7 {
	  	fmt.Print(" seventy")
	  } else if tens == 8 {
	  	fmt.Print(" eighty")
	  } else if tens == 9 {
	  	fmt.Print(" ninty")
	  }
	  if reminder == 1 {
	  	fmt.Print(" one")
	  } else if reminder == 2 {
	  	fmt.Print(" two ")
	  } else if reminder == 3 {
	  	fmt.Print(" three ")
	  } else if reminder == 4 {
	  	fmt.Print(" four ")
	  } else if reminder == 5 {
	  	fmt.Print(" five ")
	  } else if reminder == 6 {
	  	fmt.Print(" six  ")
	  } else if reminder == 7 {
	  	fmt.Print(" seven ")
	  } else if reminder == 8 {
	  	fmt.Print(" eight ")
	  } else if reminder == 9 {
	  	fmt.Print(" nine ")
	  }
	  fmt.Println() */

	//if else lader

	/*var percentage int

	  fmt.Println("Please Enter your percentage : ")
	  fmt.Scan(&percentage)

	  if percentage >= 75 {
	  	fmt.Println("...congratulations you are inDistinction...")
	  } else if percentage >= 60 {
	  	fmt.Println("...you g ot ( A ) Grade...")
	  } else if percentage >= 45 {
	  	fmt.Println("...you Got ( B ) Grade...")
	  } else if percentage >= 35 {
	  	fmt.Println("...you Got ( C ) Grade...")
	  } else {
	  	fmt.Println("...Ufffff you Are Fail...")
	  }*/
	//------------------------------------------------
	//( initialize/declare/Assignment of variable in if statement)

	//Question w.g.p square of number is even or odd

	/*if num := 7; (num*num)%2 == 0 {
	  	fmt.Println("Square of 7 is Even number... ")
	  } else {
	  	fmt.Println("Square of 7 is Odd...")
	  }*/

	/*	var num int
		fmt.Println("Please enter your Number")

		if fmt.Scan(&num); (num*num)%2 == 0 {
			fmt.Printf(" Square of %d is Even Number", num)
		}*/

	// w.g.p conert 3 digit Number in Golang
	/*var amount int
	  fmt.Println("please enter Your Amount : ")
	  fmt.Scan(&amount)

	  hundreds := amount / 100
	  reminder := amount % 100
	  tens := reminder / 10
	  reminder = reminder % 10

	  if hundreds == 1 {
	  	fmt.Println("One hundered")
	  } else if hundreds == 2 {
	  	fmt.Println("Two Hundred")
	  } else if hundreds == 3 {
	  	fmt.Println("Three Hundered")
	  } else if hundreds == 4 {
	  	fmt.Println("Fpur Hundred")
	  } else if hundreds == 5 {
	  	fmt.Println("Five Hundred")
	  }
	  if tens == 1 {
	  	fmt.Println("Ten")
	  } else if tens == 2 {
	  	fmt.Println("Twenty")
	  } else if tens == 3 {
	  	fmt.Println("Thirty")
	  } else if tens == 4 {
	  	fmt.Println("Fourty")
	  } else if tens == 5 {
	  	fmt.Println("Fifty")
	  } else if tens == 6 {
	  	fmt.Println("Sixty")
	  } else if tens == 7 {
	  	fmt.Println("Seventy")
	  } else if tens == 8 {
	  	fmt.Println("Eighty")
	  } else if tens == 9 {
	  	fmt.Println("Ninety")
	  }
	  if reminder == 1 {
	  	fmt.Println("One ")
	  } else if reminder == 2 {
	  	fmt.Println("Two")
	  } else if reminder == 3 {
	  	fmt.Println("Three")
	  } else if reminder == 4 {
	  	fmt.Println("Four")
	  } else if reminder == 5 {
	  	fmt.Println("Five")
	  } else if reminder == 6 {
	  	fmt.Println("Six")
	  } else if reminder == 7 {
	  	fmt.Println("Seven")
	  } else if reminder == 8 {
	  	fmt.Println("Eight")
	  } else if reminder == 9 {
	  	fmt.Println("Nine")
	  }
	  fmt.Println()*/

	/*var num1, num2 float64
	var operator string

	fmt.Println("Enter Your First Number : ")
	fmt.Scan(&num1)

	fmt.Println("Enter Your Second Number : ")
	fmt.Scan(&num2)

	fmt.Println("Choose Operator + , - , * , / ")
	fmt.Scan(&operator)

	var result float64
	switch operator {
	case "+":
		result = num1 + num2
		fmt.Printf("Result = %.2f, %s, %f, %f ", num1, operator, num2, result)
	}*/

}
