package main

import (
	"fmt"
)

/*func SliceDeclaration() {
	var slice []int //Nil Slice
	fmt.Println(slice, len(slice), cap(slice))

	if slice == nil {
		fmt.Println("Its Nill Slice")
	} else {
		fmt.Println("Its Not Nill Slice")
	}

	// Append elements to the slice
	slice = append(slice, 10, 20)
	fmt.Println(slice, len(slice), cap(slice))
}*/
//write a golag function that take 10 number from user and store into the slice
func GetInput() {
	var data []int
	const messege = "Please Enter Any 10 Number: "
	var num int
	fmt.Println(messege)
	for start := 1; start <= 10; start++ {
		fmt.Scan(&num)
		data = append(data, num)
	}
	fmt.Println(data)

	//Print square of each Number in slice
	for _, value := range data {
		fmt.Printf(" %d", value*value)
	}
}
func MakingSliceAndAssigningSlice() {
	slice := make([]int, 5, 10)
	fmt.Println(slice)
	slice[0] = 10
	slice[1] = 20
	slice[2] = 30
	slice[3] = 40
	slice[4] = 50
	slice = append(slice, 60, 70, 80)
	fmt.Println(slice)
}

func SliceLiteral() {
	slice := []int{10, 20, 30, 40, 50, 60} //slice Literal
	var data []int = []int{10, 20, 30, 40}
	fmt.Println("Slice Declaration", data)

	//read slice element indivisually

	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3])
	fmt.Println(slice[4])
	fmt.Println(slice[5])

	slice[4] = 400

	fmt.Println("----------------------------------")

	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	fmt.Println(slice[3])
	fmt.Println(slice[4])
	fmt.Println(slice[5])

}
func CopySlice() {
	source := []int{10, 20, 30, 40, 50}
	destination := make([]int, len(source))
	fmt.Println("----------Before Copy Slice--------")
	fmt.Println(destination)
	copy(destination, source)
	fmt.Println("---------After Copy Slice-------")
	fmt.Println(destination)
}

func CreateSlicefromArray() {
	var array [5]int = [5]int{10, 20, 30, 40, 50}
	//create slice of all Array

	slice := array[:] //create slice of all element in an Array
	fmt.Println(slice)
	//create slice from start index it remaning element

	slice = array[2:]
	fmt.Println("Create slice from start index till end ", slice)

	//create slice in range (from start index to end )
	const startindex = 1
	const endindex = 4
	slice = array[startindex:endindex]
	fmt.Println("Create slice in range ", slice)

	//create slice from sstarting to end index
	slice = array[:4] //create slice from 0 to 4-1=3
	fmt.Println("Created Slice from 0 to 3 index", slice)
	slice = append(slice, 60, 70, 80)
	fmt.Println("By Appending 60 70 80 Element into the slice", slice)

}

// Passing Slice To TheFunction
func PrintSlice(slice []int) {
	for _, value := range slice {
		fmt.Printf(" %d", value)
	}
}
func PassingSliceToFunction() {
	slice := []int{10, 20, 30, 40, 40, 50, 60}
	PrintSlice(slice)
}

//Returning slice from function

func CreateSlice(length int) []int {

	slice := make([]int, length)
	return slice

}

/*func ReturnSliceFromFunction() {
	slice := CreateSlice(6)
	fmt.Println(slice, len(slice))

	slice = CreateSlice(10)
	fmt.Println(slice, len(slice))
}*/

// Write a golag programetio display only not duplication numbers
// 101% interview Question

func ScanIntoSlice(slice []int) []int {
	fmt.Printf("Please Enter %d Number :", len(slice))
	for index := 0; index < len(slice); index++ {
		fmt.Scan(&slice[index])
	}
	return slice
}
func DisplayNotDuplicateNumber() {
	slice := CreateSlice(10)
	slice = ScanIntoSlice(slice)
	fmt.Println(slice)

	var duplicateSlice []int

	for _, value := range slice {
		if !IsDuplicate(value, duplicateSlice) {
			fmt.Printf(" %d", value)
			duplicateSlice = append(duplicateSlice, value)
		}
	}
	fmt.Println()

}
func IsDuplicate(searchNum int, slice []int) bool {
	for _, value := range slice {
		if value == searchNum {
			return true
		}
	}
	return false
}

func Slicemain() {
	//SliceDeclaration()
	//GetInput()
	//MakingSliceAndAssigningSlice()
	//SliceLiteral()
	//CopySlice()
	//CreateSlicefromArray()
	//PassingSliceToFunction()
	//ReturnSliceFromFunction()
	DisplayNotDuplicateNumber()

}
