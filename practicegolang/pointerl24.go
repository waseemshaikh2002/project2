package main

import "fmt"

/*func main() {
//type Employee struct {
	Name string
	Age  int
}
/*employee := Employee{"Raju", 10}
var p *Employee = &employee
fmt.Println(employee)
fmt.Println(p.Age, p.Name)*/

/*var age int = 10
var p *int = &age

fmt.Println(p)
fmt.Println(*p)*/

/*var name string = "codelangs"
var p *string = &name
fmt.Println(p)
fmt.Println(*p)*/

//Dynamically Allocation of Memory

/*p := new(int)
fmt.Println(p)
*p = 50
fmt.Println(*p)*/

// Zero pointer

/*var p *int
fmt.Println(p)*/

//}

/*func changeData(num *int) {
	*num = 50
}
func main() {
	var num int = 100
	fmt.Println(num)
	fmt.Println("----Before calling function----")
	changeData(&num)
	fmt.Println("-----After calling Function-----")
	fmt.Println(num)

}*/

//---------------------------------------------------
//-------------------- Pointer of Pointer------------

/*func pointrOfPointer() {
	var num int = 50
	var p *int = &num
	var pp **int = &p
	fmt.Println(pp)
	fmt.Println(**pp)
}
func main() {
	pointrOfPointer()
}*/

//---------------------------------------------------
//-----------------Types of Pointer-----------------

// ----------------Pointer to struct------------
/*type Person struct {
	Name string
	Age  int
}

func main() {
	var person Person
	person.Age = 30
	person.Name = "Raju"
	var ptr *Person = &person
	fmt.Println(ptr.Name, ptr.Age)

}*/

// ---------------------------------------------------
// ----------pointer to interface---------------------
/*type Person struct {
	Name string
	Age  int
}
type friendship interface {
	friend()
}

func (Person) friend() {
	fmt.Println("We Are Friend")

}
func pointerToInterface() {
	var person = Person{"Raju", 30}
	var friend friendship = person
	var pi *friendship = &friend

	(*pi).friend()

}
func main() {
	pointerToInterface()

}*/

// -----------------------------------------------
// ---------------Pointer To Array----------------
// -------------------------------------------------
/*func pointerToArray() {
	var data [3]int = [3]int{10, 20, 30}
	var p *[3]int = &data
	fmt.Println(p)
	fmt.Println((*p)[2])

}
func main() {
	pointerToArray()
}*/

//----------------------------------------------
//---------Pointer To slice--------------------

/*func pointerToSlice() {
	var slice = []int{10, 20, 30, 40}
	var p *[]int = &slice

	fmt.Println(p)
	fmt.Println((*p)[1])

}
func main() {
	pointerToSlice()
}*/

//------------------------------------------------
//-----------Pointer To Function----------------

/*func pointerToslice() {
	var slice = []int{10, 20, 30}
	var p *[]int = &slice

	fmt.Println(p)
	fmt.Println((*p)[1])
}
func pointerToFunction() {
	var f func() = pointerToslice
	f()

}
func main() {
	pointerToFunction()
}*/

//-------------------------------------------------
//-------------Pointer receiver Method-------------

/*type Person struct {
	Name string
	Age  int
}

func (p *Person) walk() {
	p.Name = "Raju"

}
func pointerReceiverMethod() {
	var person = Person{}
	person.walk()
	fmt.Println(person.Name)

}
func Pointermain() {
	pointerReceiverMethod()
}*/

//-----------------------------------------------------
//-----------------------------------------------------
//-----------Assigmment on Pointer and struct in golag---

// Q1) Declare an integer variable and pointer to it
// print the address and value using pointer.
/*func printAddressAndValue(num *int) {
	fmt.Printf("Address of num: %p\n", num)
	fmt.Printf("Value of num: %d\n", *num)
}

func main() {
	var number int = 42
	var ptr *int = &number

	printAddressAndValue(ptr)
}*/

//--------------------------------------------------------
// Q2) Modify the value of an intiger variable using a pointer
//and print the updated value.

/*func modifyvalue(num *int) {
	*num = 100
}
func printAddressAndValue(num *int) {
	fmt.Println("Address of Number :", num)
	fmt.Println("value Of Number :", num)

}
func main() {
	number := 42
	ptr := &number

	fmt.Println("Before Modification :")
	printAddressAndValue(ptr)

	modifyvalue(ptr)

	fmt.Println("After Modification")
	printAddressAndValue(ptr)
}*/

//-------------------------------------------------------
// 3) Create a nil Pointer and check if its nil

/*func main() {
	var ptr *int
	if ptr == nil {
		fmt.Println("Pointer is Nil")
	} else {
		fmt.Println("pointer is Not Nil")
	}
}*/

//-----------------------------------------------------
// Q4) Attempt to Derefrence a nil pointer and handle the error.

/*func checkAndDereference(ptr *int) {
	if ptr == nil {
		fmt.Println("pinter is Nill cannot Derefrence")
	} else {
		fmt.Println("Value:", *ptr)
	}
}
func main() {
	var ptr *int
	checkAndDereference(ptr)

	value := 42
	ptr = &value
	checkAndDereference(ptr)

}*/

//---------------------------------------------------------
//5)  create float variable , create pointer to it.
//and print the value using it.

/*func printFloatValue(ptr *float64) {
	if ptr == nil {
		fmt.Println("poiter is nil cannot derefrence")
	} else {
		fmt.Println("Value :", ptr)
	}
}
func main() {
	var number float64 = 50.20
	var ptr *float64 = &number

	printFloatValue(ptr)
}*/

//---------------------------------------------------------
//Q6) Swap the value of Two Integer Variable Using Pointer

/*func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
func main() {
	var x int = 10
	var y int = 20

	fmt.Println("Before swap:")
	fmt.Println("x = %d, y = %d\n", x, y)

	swap(&x, &y)

	fmt.Println("After Swap :")
	fmt.Println("x =%d, y =%d \n", x, y)
}*/
//error
//---------------------------------------------------
// Q7) While function that accept pointer to string and modify string.
/*func modifyingstring(b *[]byte) {
	if len(*b) > 0 {
		(*b)[0] = 'H'
	}
}
func main() {
	str := "hello"
	byteslice := []byte(str)

	modifyingstring(&byteslice)

	modifiedstring := string(byteslice)

	fmt.Println(modifiedstring)
}*/
//----------------------------------------------------
// Q8) Write a functin that accept 0 pointer to float and double its value.

/*func doubleValue(f *float64) {
	if f != nil {
		*f = *f * 2
	}
}
func main() {
	value := 3.14

	fmt.Println("Original value", value)

	doubleValue(&value)

	fmt.Println("Double Value :", value)

}*/
//-----------------------------------------------------
// Q9) Creaate Array Of Integer and pointer to the first element print the first element usinf pointer

/*func printFirstElement(ptr *int) {
	fmt.Println("First Element :", ptr)
}
func main() {
	array := [5]int{10, 20, 30, 40, 50}

	pointerToFirstElement := &array[0]

	printFirstElement(pointerToFirstElement)

}*/
//----------------------------------------------------------
// 10) increment a pointer to an array and print the next element.

/*func printNextElement(array *[5]int, index int) {
	if index < len(array)-1 {
		index++
		fmt.Println("Next Element :", (*array)[index])
	} else {
		fmt.Println("No Next Element.")
	}
}
func main() {
	array := [5]int{10, 20, 30, 40, 50}

	pointerToArray := &array

	printNextElement(pointerToArray, 0)

}*/
//---------------------------------------------------------
// 11)  Write a function that return a pointer to an integer initialized with a specific value.
// i write wrong question.
/*func newIntPointer(value int) *int {
	num := value
	return &num
}
func main() {
	ptr := newIntPointer(20)

	fmt.Println("Value", ptr)
}*/

//---------------------------------------------------------
// Q12)  Create a function that return a pointer to an
//integer initialized with a specific value.

/*func newIntPointer(value int) *int {
	num := value
	return &num
}

func main() {
	ptr := newIntPointer(20)

	fmt.Println("value", ptr)

}*/
//-------------------------------------------------------

// Q13) Write a function That Takes a pointer to a boolean
//and toggle its value.

/*func toggleBoolean(b *bool) {
	*b = !*b
}
func main() {
	value := true

	fmt.Println("Original Value: ", value)

	toggleBoolean(&value)

	fmt.Println("Toggled Value: ", value)

	toggleBoolean(&value)

	fmt.Println("Toggle Value Again.", value)
}*/
//---------------------------------------------------------
//Q14) Implement a function that accept a pointer to a struct
//and modify one of its fields.

/*type person struct {
	Name string
	Age  int
}

func modifyPersonName(p *person, newName string) {
	p.Name = newName
}
func main() {
	person := person{Name: "Waseem", Age: 23}
	fmt.Println("Original Struct: ", person)

	modifyPersonName(&person, "Nawab")

	fmt.Println("Modified struct :", person)
}*/

//-------------------------------------------------------
//15) Write function that takes a pointer to an array and sets its all element to aspecific value.

/*func serArrayElement(arrptr *[5]int, value int) {
	for i := 0; i < len(*arrptr); i++ {
		(*arrptr)[i] = value
	}
}
func main() {
	array := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Original Array :", array)

	serArrayElement(&array, 10)

	fmt.Println("Modified Array:", array)
}*/
//-------------------------------------------------------
// 16) Create function that accept a pointer to slice and Appends an element to it.

/*func appendToslice(sliceptr *[]int, element int) {
	*sliceptr = append(*sliceptr, element)
}
func main() {
	slice := []int{1, 2, 3}

	fmt.Println("Original Slice: ", slice)

	appendToslice(&slice, 4)

	fmt.Println("Modified slice", slice)

}*/

//-------------------------------------------------------
//Q17) Write function that takes pointer to map and key value pointe.

/*func updateMapValue(mapptr *map[string]int, key *string, valueptr *int) {
	(*mapptr)[*key] = *valueptr
}
func main() {
	myMap := map[string]int{
		"one":   1,
		"two":   2,
		"Three": 3,
	}
	fmt.Println("Original Map", &myMap)

	key := "two"
	value := 22

	updateMapValue(&myMap, &key, &value)

	fmt.Println("Modified Map:", myMap)
}*/

//---------------------------------------------------------
//Q18) Implement function that Accept A pointer to a complex
// number and sets its read imaginery park.

/*func setcomplexRealImag(cptr *complex128, real, imag float64) {
	*cptr = complex(real, imag)
}
func main() {
	complexNum := complex(1.0, 2.0)
	fmt.Println("Original Complex Number: ", complexNum)

	setcomplexRealImag(&complexNum, 3.0, 4.0)

	fmt.Println("Modified Complex Number: ", complexNum)

}*/

//--------------------------------------------------------
// 20) Define a Sruct within multipe fields and creae a pointer to an instance of it. modify its field usinfg te pointer.

/*type person struct {
	Name    string
	Age     int
	Country string
}

func modifyperson(p *person, name string, age int, country string) {
	p.Name = name
	p.Age = age
	p.Country = country
}
func main() {
	personptr := &person{Name: "Waseem", Age: 23, Country: "INDIA"}

	fmt.Println("Original Person:", *personptr)

	modifyperson(personptr, "Nawab", 23, "India")

	fmt.Println("Modified Person: ", *personptr)
}*/

//---------------------------------------------------------
//Q21) Ceate a function That makes a pointe struct and
//pints its fields.

/*type point struct {
	x, y int
}

func createAndPrintPoint() {
	pointptr := &point{}

	pointptr.x = 10
	pointptr.y = 20

	fmt.Println("Point Coordinates :", pointptr.x, pointptr.y)
}
func main() {
	createAndPrintPoint()
}*/

//---------------------------------------------------------
// Q22)  Write a function that returns a pointer to struct
//initialized with specific value.

/*type person struct {
	Name string
	Age  int
}

func newPeson(name string, age int) *person {
	person := person{Name: name, Age: age}

	return &person
}
func main() {
	personptr := newPeson("Waseem", 23)

	fmt.Println("Name: ", personptr.Name)
	fmt.Println("Age: ", personptr.Age)
}*/

//-----------------------------------------------------
//Q23) Create struct nested struct and modify
//the nested struct field using pointers.

/*type Address struct {
	City    string
	Country string
}
type Person struct {
	Name    string
	Age     int
	Address *Address // pointer to access struct
}

func modifyaddressCity(person *Person, city string) {
	if person.Address != nil {
		person.Address.City = city
	}
}

func Pointermain() {
	person := Person{
		Name: "Akshay",
		Age:  30,

		Address: &Address{
			City:    "Pune",
			Country: "India",
		},
	}
	fmt.Println("Original Person Detail:")
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Address:", person.Address.City+",", person.Address.Country)

	modifyaddressCity(&person, "Parli vaijanth")

	fmt.Println("\nModified Person Details:")
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Address:", person.Address.City+",", person.Address.Country)
}*/

//------------------------------------------------------------------------------
//-----------------------------------------------------------------------------

//Pointer Assignmen tAssignment with Sir

// First Assignmetn with sir
// Declare an integer variable and pointer to it
// print the address and value using pointer

/*func Assignment_1() {
	var intvar int = 10
	var prtvar *int = &intvar

	fmt.Println("Memory Address Of Int variable: ", prtvar)
	fmt.Println("Value Of Int Variable : ", *prtvar)
}
func main() {
	Assignment_1()

}*/

// -----------------------------------------------------------------------------------
// Assignmet 2 :  Modify the value of an intiger variable using a pointer
// and print the updated value.
/*func Assignment_2() {
	var x int = 50
	fmt.Println("Original Value Of x : ", x)

	var p *int = &x
	*p = 80
	fmt.Println("Modified Value of x : ", x)
}
func main() {
	Assignment_2()

}*/

//-------------------------------------------------------------------------------------------

//Assignment 3 : Create A Nill Pointer And check If it is Nil.

/*func Assigmment_3() {
	var p *int //Nil Pointer
	if p == nil {
		fmt.Println("Its Nil Pointer ")
	} else {
		fmt.Println("Its Not Nil Pointer")
	}
}
func main() {
	Assigmment_3()
}*/

//----------------------------------------------------------------------------------------------

// Assignment 4 : Attempt to derefrence a nill pointer And Handle The Error.

/*func Assignment_4() {
	var p *int
	defer func() {
		recover()
		fmt.Println("Could noit derefrence NIll pointer")
	}()

	*p = 50 // panic
}
func main() {
	Assignment_4()
}*/

// --------------------------------------------------------------
// ------------Assignment 5---------------------------
// create float varibole , create poiter to it.
//and print the value using it.
/*func Assignment_5() {
	var x float64 = 200.30
	var p *float64 = &x
	fmt.Println(*p)
}
func main() {
	Assignment_5()
}*/

//------------------------------------------------------------
//----------Assignment 6 -------------------------------------
//swap the value of two Intiger variable using pointer.

func Assignment_6() {
	var x, y int = 10, 20
	fmt.Println(x, y)
	var px, py *int = &x, &y
	*px = *px + *py
	*py = *py - *px
	*px = *px - *py
	fmt.Println(*px, *py)

}
func Pointermain() {
	Assignment_6()
}
