package main

import "fmt"

/*func Iterateover2DArray() {
	var data [3][3]int

	data[0][0] = 10
	data[0][1] = 20
	data[0][2] = 30

	data[1][0] = 40
	data[1][1] = 50
	data[1][2] = 60

	data[2][0] = 70
	data[2][1] = 80
	data[2][2] = 90
	//using simple for loop
	/*for row := 0; row < len(data); row++ {
		for column := 0; column < len(data[row]); column++ {

			fmt.Printf(" %d", data[row][column])
		}
		fmt.Println()

	}

	// using for range loop
	fmt.Println()
	for _, row := range data {
		for _, column := range row {
			fmt.Printf("%d ", column)
		}
		fmt.Println()
	}
}*/

// ============================
/*func Create2DArray() {
	var data [3][3]int
	data[0][0] = 10
	data[0][1] = 20
	data[0][2] = 30

	data[1][0] = 40
	data[1][1] = 50
	data[1][2] = 60

	data[2][0] = 70
	data[2][1] = 80
	data[2][2] = 90

	fmt.Println(data)
	//if we want to known Where Which value is stored in our Array
	fmt.Println(data[2][1]) //we can write this

}*/

//write a go programe to take 3*3 value and print as it is

func Readvalue() {
	var data [3][3]int
	fmt.Println("Please Enter 3*3 Values :")
	for rowIndex := 0; rowIndex < len(data); rowIndex++ {
		for colIndex := 0; colIndex < len(data); colIndex++ {
			fmt.Scan(&data[rowIndex][colIndex])
		}

	}
	//---------------- sum Using Range --------
	var sum int = 0
	for _, row := range data {
		for _, column := range row {
			sum += column
			fmt.Printf(" %d", column)
		}
		fmt.Println()
	}
	fmt.Printf("\nsum is = %d ", sum)

}

func Read2DArray() [3][3]int {
	var data [3][3]int
	const messege = "Please Enter 3*3 Matrix"
	fmt.Println(messege)
	for rowIndex := 0; rowIndex < len(data); rowIndex++ {
		for colIndex := 0; colIndex < len(data[rowIndex]); colIndex++ {
			fmt.Scan(&data[rowIndex][colIndex])
		}
	}
	return data
}
func Print2darray(data [3][3]int) {
	for _, row := range data {
		for _, col := range row {
			fmt.Printf(" %d", col)
		}
		fmt.Println()
	}

}
func GenerateMatrix() {
	var matrix1 [3][3]int = Read2DArray()
	var matrix2 [3][3]int = Read2DArray()
	Print2darray(matrix1)
	fmt.Println("\n-----------------------------")
	Print2darray(matrix2)
}

// write a go programe for addtion of two matrix
func TwoMatrixAddition() {
	var matrix1 [3][3]int = Read2DArray()
	var matrix2 [3][3]int = Read2DArray()
	var matrix3 [3][3]int

	for rowIndex := 0; rowIndex < len(matrix1); rowIndex++ {
		for colIndex := 0; colIndex < len(matrix1[rowIndex]); colIndex++ {
			sum := matrix1[rowIndex][colIndex] + matrix2[rowIndex][colIndex]
			matrix3[rowIndex][colIndex] = sum
		}
	}
	Print2darray(matrix1)
	fmt.Println()
	Print2darray(matrix2)
	fmt.Println()
	Print2darray(matrix3)
}

func Array2Dmain() {
	//Iterateover2DArray()
	//Create2DArray()
	//Readvalue() //it will only Read the value
	//Read2DArray() //Its Also read the value
	//GenerateMatrix()
	TwoMatrixAddition()
}
