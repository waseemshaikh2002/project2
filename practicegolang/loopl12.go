package main

import "fmt"

func Maintwelve() {
	/*	----(infinite for loop)-----

		for {
			fmt.Println("Welcome")
		}*/

	//-----(for with condition / for like while loop)---
	/*var count = 1
	for count <= 10 {
		fmt.Println(count)
		count++
	}*/

	//------(write a golang programe to print square of 1 to 10 numbers)-----
	/*var count = 1
	for count <= 10 {
		fmt.Printf("square of %d = %d \n", count, count*count)
		count++
	}*/

	//-----(common for loop)------

	/*for start := 1; start <= 10; start++ {
		fmt.Println(start)
	}*/

	//-----(w. g. p. to print table of given Number)---

	/*var num int
	var messege = "please Enter Any Number : "
	fmt.Println(messege)
	fmt.Scan(&num)
	fmt.Printf("Table of %d as follow \n", num)
	for start := 1; start <= 10; start++ {
		fmt.Println(start * num)
	}*/

	//-----(Write a golang programme print table of given range)----

	/*var start, end int
	var messege string = "please Enter start and end "

	fmt.Println(messege)
	fmt.Scan(&start, &end)
	fmt.Printf("\nYour table from %d to %d\n", start, end)

	for row := 1; row <= 10; row++ {
		for column := start; column <= end; column++ {
			fmt.Printf("%d\t", column*row)
		}
		fmt.Println()
	}*/

	//---------(Loop control/Jumping statements)---
	// 1) Break       2)continue
	//---(w.g.p. to print only file number among 10 number)

	/*for start := 1; start <= 10; start++ {
		if start > 5 {
			break
		}
		fmt.Println(start)
	}*/

	//---(w.g.p. to check given Number is prime or not)---
	//5/15/5
	/*message := "please enter Any Number :"
	fmt.Println(message)
	var num int
	fmt.Scan(&num)
	var start int
	for start = 2; start <= num-1; start++ {
		if num%start == 0 {
			fmt.Printf("\n %d is not prime Number\n", num)
			break
		}

	}
	if num == start {
		fmt.Printf("\n%d is Prime Number \n", num)
	} else {
		fmt.Printf("\n%d is Not prime Number \n", num)
	}*/

	//continue stament//
	//w.g.p. to print last 5 number from 10 number

	/*for start := 1; start <= 10; start++ {
		if start < 6 {
			continue
		}
		fmt.Println(start)
	}*/
}

func Printnumsmain() {
	/*var count = 1
	for count <= 10 {
		fmt.Println(count)
		count++
	}*/
}

func Simmain() {
	var count = 1
	for count <= 10 {
		fmt.Printf("square of %d = %d \n", count, count*count)
		count++
	}
}

// common for loop

func Somain() {
	for start := 1; start <= 10; start++ {
		fmt.Println(start)

	}
}

// common for loop
func Comain() {
	for start := 1; start <= 10; start++ {
		fmt.Println(start)
	}

}

//Print table using for loop

func Tamain() {
	var num int
	messege := "Enter Any Number :"
	fmt.Println(messege)
	fmt.Scan(&num)
	fmt.Println("Table of Numbers As follows \n ")
	for start := 1; start <= 10; start++ {
		fmt.Println(start * num)
	}
}

// Print table using for loop

func main() {
	var num int
	messege := "Enter Any Number :"
	fmt.Print(messege)
	fmt.Scan(&num)
	fmt.Println("Table of NUmber As follows")
	for start := 1; start <= 10; start++ {
		fmt.Println(start * num)

	}
}
