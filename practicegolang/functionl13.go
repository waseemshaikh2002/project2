package main

//import "fmt"

/////-----------(Named Function)------/////

//-------(Paramitrized Named Function)------//

func Mainthirteen() {

	//Addition of two Numbers
	/*var add = func(num1 int, num2 int) {
		fmt.Println(num1 + num2)
	}
	add(20, 50)*/
	//out put = 70

	//---(invocation of functiion )----

	/*func(num1, num2 int) {
		fmt.Println(num1 + num2)
	}(10, 20)*/
	//out put = 30

	//------(lexical Scope)----
	/*var x = 50
	func() {

		fmt.Println(x * x)
	}()*/
	// out put 2500

	//---(parameter No return Value)---

	/*func(num int) {
		fmt.Println(num * num)
	}(5)*/
	//out put = 25

	//--------(parameter with return value)----

	/*square := func(num int) int {
		return (num * num)
	}(5)
	fmt.Println(square)*/
	//out put = 25

	//-------(No parameter with return value)---

	/*pi := func() float64 {
		return 3.14

	}()
	fmt.Println(pi)*/

	//------(No parameter No return value)----
	/*func() {
		fmt.Println("welcome in golang")
	}()*/
	//out put = welcome in golang

	/*var add = func(num1 int, num2 int) {
		fmt.Println(num1 + num2)

	}
	add(50, 20)*/

	/////////-----(Self Calling Function)-----//////////
	/*func() {
		fmt.Println("Self calling Function")
	}()*/

	///-----(Parametrized self calling Function)---///

	/*func(num int) {
		fmt.Println(num + num)
	}(6)*/

	//--(self calling Function with parameter and returning value)

	/*cube := func(num int) int {
		return (num * num * num)
	}(5)
	fmt.Println(cube)*/

	//parameter No return Value

	/*greetting := func(name string) {
		fmt.Printf("\nWelcome %s \n", name)
	}
	greetting("raju")*/

	//--(parameter with retrun value)---//
	/*square := func(num int) int {
		return num * num
	}
	result := square(5)
	fmt.Println(result)
	result = square(6)
	fmt.Println(result)*/

	///---(No Parameter Return value)---///

	/*pi := func() float64 {
		return 3.14

	}
	result := pi()
	fmt.Println(result)*/

	//---(No parameter No return value)----//

	/*fn := func() {
		fmt.Println("Welcome in no parameter No return value")
	}
	fn()*/

	///----(Multiple return Value)----///

	/*add := func(num int) (int, int) {
		square := num * num
		cube := square * num
		return square, cube

	}
	square, cube := add(4)
	fmt.Println(square, cube)*/

}
