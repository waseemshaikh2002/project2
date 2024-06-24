package main

import "fmt"

//declaration of struct

/*type employee struct {
	age    int
	salary int
	name   string
	gender string
}

// -------------------------------------------------
func structLiteral() {
	emp1 := employee{age: 20, salary: 20000, name: "michal", gender: "male"}
	fmt.Println(emp1)

	emp2 := employee{age: 30, salary: 40000, name: "mitchel", gender: "male"}
	fmt.Println(emp2)

}
func main() {
	structLiteral()
}*/

//--------------------------------------------
/*type employee struct {
	name   string
	age    int
	salary int
	gender string
}

func createStructObject() {

	var emp1 employee    //create struct object
	emp1.name = "waseem" //struct field assignment
	emp1.age = 22        //struct field assignment
	emp1.salary = 240000 //struct field assignment
	emp1.gender = "male" //struct field assignment
	fmt.Println(emp1)

	var emp2 employee
	emp2.name = "deepak"
	emp2.age = 22
	emp2.salary = 240000
	emp2.gender = "male"
	fmt.Println(emp2)
}
func main() {
	createStructObject()
}*/

//take data feom user into struct

/*type computer struct {
	price       float64
	color       string
	CompanyName string
	Ramsize     int
}

func ReadDataIntoStruct() {
	var computer computer
	fmt.Println("Enter Computer price: ")
	fmt.Scan(&computer.price)
	fmt.Println("Enter the color of computer: ")
	fmt.Scan(&computer.color)
	fmt.Println("Enter the company Nmae of laptop: ")
	fmt.Scan(&computer.CompanyName)
	fmt.Println("Enter the Ramsize of Laptop: ")
	fmt.Scan(&computer.Ramsize)

	fmt.Println("------After getting data")
	fmt.Println(computer)
}
func main() {

	ReadDataIntoStruct()
}*/

/*func AnonymousStruct() {
	// it is struct without name calls as anonymous struct
	emp1 := struct {
		age    int
		name   string
		salary int
		gender string
	}{age: 20, name: "Raja", salary: 7000, gender: "Male"}
	fmt.Println(emp1)

}
func main() {
	AnonymousStruct()
}*/

// -----------------------------------
/*type computer struct {
	name        string
	color       string
	CompanyName string
	price       int
	Ramsize     int
}

// Passing Struct To the fucntion
func PrintStructData(computer computer) {
	fmt.Println(computer)
}
func ComputerData() {
	computer := computer{name: "waseem", color: "black", CompanyName: "Hp", price: 65000, Ramsize: 8}
	PrintStructData(computer)

}
func main() {
	ComputerData()
}*/
//------------------------------------------

//Returning Struct From the Function

//Declaration of struct

/*type pen struct {
	price int
	color string
	brand string
}

//multiple retrun value function

func CreateStruct() (pen, pen) {
	var pen1 pen //object created
	pen1.brand = "cello"
	pen1.color = "blue"
	pen1.price = 10

	pen2 := pen{price: 20, color: "red", brand: "styllo"} //struct literal

	return pen1, pen2
}
func ReturnStruct() {
	pen1, pen2 := CreateStruct()
	fmt.Println(pen1)
	fmt.Println(pen2)
}
func main() {
	ReturnStruct()
}*/
//------------------------------------------

/*type table struct {
	MaterialType string
	price        int
	weight       float64
}

//Array Of Structure

func ArrayOfStructure() {
	table1 := table{MaterialType: "Plastic", price: 2300, weight: 2}
	table2 := table{MaterialType: "Metal", price: 6300, weight: 100}
	table3 := table{MaterialType: "cement", price: 250, weight: 25}
	table4 := table{MaterialType: "marbel", price: 25000, weight: 1000}
	var tables [4]table
	tables[0] = table1
	tables[1] = table2
	tables[2] = table3
	tables[3] = table4

	fmt.Println("-------Array OF Table-----")

	fmt.Println(tables)
	fmt.Println("------------------------- Iterate Over Array -----------------")
	for _, table := range tables {
		fmt.Println("Material Type : " + table.MaterialType)
		fmt.Println("Price  : " + strconv.Itoa(table.price))
		fmt.Println("Weight : " + strconv.FormatFloat(table.weight, 'f', -1, 64))
		fmt.Println("-------------------------------------------------------------")
	}
}
func main() {
	ArrayOfStructure()
}*/

//-------------------------------------------------------------------
// Search Books

/*type Book struct {
	Name   string
	Price  int
	Author string
	Pages  int
}

func CreateBook() []Book {
	book1 := Book{Name: "java", Price: 700, Author: "James Goslin", Pages: 700}
	book2 := Book{Name: "c", Price: 800, Author: "Dennis Ritche", Pages: 450}
	book3 := Book{Name: "c++", Price: 900, Author: "Bjarne Stoustrup0", Pages: 600}
	book4 := Book{Name: "golang", Price: 1500, Author: "Robert", Pages: 950}
	book5 := Book{Name: "Javascript", Price: 1200, Author: "Natescape Navigator", Pages: 1000}
	books := []Book{book1, book2, book3, book4, book5}
	return books
}

func searchBook() {
	books := CreateBook()
	var bookName string
	fmt.Println("Please Enter the Book Name: ")
	fmt.Scan(&bookName)
	var isavailable bool = false
	for _, book := range books {
		if strings.EqualFold(book.Name, bookName) {
			isavailable = true
			fmt.Println("-------------------------------")
			fmt.Println(bookName + "Book is Found")
			fmt.Println("----------------------------------")
			fmt.Println("Name" + book.Name)
			fmt.Println("Price" + strconv.Itoa(book.Price))
			fmt.Println("Author" + book.Author)
			fmt.Println("Pages" + strconv.Itoa(book.Pages))
			fmt.Println("--------------------------------------------")

		}
	}
	if isavailable {
		fmt.Println("-----------------------------------")
		fmt.Println(bookName + "Book is Not Available")
		fmt.Println("--------------------------------------")
	}
}

func main() {
	searchBook()
}*/

//===========================================================================
//=================== Struct L2 ============================================
//===========================================================================

/*
	type Human struct {
		Weight float64
		Height float64
		Age    int
	}

	type employee struct {
		Name  string
		Age   int
		Human //Human Embeding Struct
	}

	func EmbeddedStruct() {
		var emp employee
		emp.Age = 10
		emp.Height = 6.12
		emp.Name = "Raju"
		fmt.Println(emp)

}

	func main() {
		// EmbeddedStruct()

}
*/
//====================================
/*type Human struct {
	Weight float64
	Height float64
	Age    int
}
type employee struct {
	name  string
	age   int
	Human //
}

func SameEmbeedeStructureProperty() {
	var emp employee
	emp.age = 10
	emp.Human.Age = 30
	emp.Height = 6.2
	emp.name = "raju"
	fmt.Println(emp)

}


	func main() {
		SameEmbeedeStructureProperty()
	}
*/
//========================================
/*type Human struct {
	Weight float64
	Height float64
	Age    int
}

func StructEqulity() {
	human1 := Human{Weight: 65, Height: 5.6, Age: 45}
	human2 := Human{Weight: 65, Height: 5.6, Age: 45}
	isEqual := human1 == human2
	fmt.Println(isEqual)
}
func main() {
	StructEqulity()
}*/
//========================================

/*type Human struct {
	Height float64
	Weight float64
	Age    int
}

func OmitFieldInStructLiteral() {
	human := Human{Weight: 100} //Omit Struct Liteal
	fmt.Println(human)

}
func main() {
	OmitFieldInStructLiteral()
}*/

//=============================================

/*type Human struct {
	Age    int
	Height float64
	Weight float64
}

func (h Human) Walking() {
	fmt.Println("Human is Walking")

}
func (h Human) Eating() {
	fmt.Println("Human is Eating")
}
func (h *Human) SetData() {
	h.Age = 30
	h.Height = 6.2
	h.Weight = 60
	fmt.Println(h)
}
func (h *Human) PrintData() {
	fmt.Println(h)
}
func CallingStructureMethod() {
	human := Human{}
	human.Eating()
	human.Walking()
}
func main() {
	CallingStructureMethod()
}*/
//=========================================
/*type Human struct {
	Age    int
	Weight float64
	Height float64
}

func (h *Human) SetData() {
	h.Age = 30
	h.Height = 6.2
	h.Weight = 60
	fmt.Println(h)
}
func (h *Human) PrintData() {
	fmt.Println(h)
}
func CallByValueStructureMethod() {
	human := Human{}
	human.SetData()
	human.PrintData()
}
func main() {
	CallByValueStructureMethod()

}

//==========================================

func BankingSystem() {

}*/

////////////////////////////////////////////////
//-----------Double pracetice struct---------
//--------------------------------------------
//----------declaration of strcut-------------

/*type employee struct{
	age int
	salary int
	gender string
}*/
//----------------------------------------------
/*type employee struct {
	age    int
	salary int
	gender string
}

func structLiteral() {
	emp1 := employee{age: 20, salary: 20000, gender: "Male"}
	fmt.Println(emp1)
}
func main() {
	structLiteral()
}*/

// --------------------------------------------------------
/*type employee struct {
	name   string
	age    int
	salary int
	gender string
}

func createStructObject() {
	var emp1 employee
	emp1.name = "waseem"
	emp1.age = 23
	emp1.salary = 20000
	emp1.gender = "Male"
	fmt.Println(emp1)

	var emp2 employee
	emp2.name = "akshay"
	emp2.age = 23
	emp2.salary = 20000
	emp2.gender = "male"
	fmt.Println(emp2)

}
func main() {
	createStructObject()
}*/

//-------------------------------------------------------
//----------Take Data from user into strcu

/*type comptuter struct {
	price       float64
	color       string
	CompanyName string
	Ramsize     int
}

func ReadDataIntoStruct() {
	var computer comptuter
	fmt.Println("Enter the computer Price: ")
	fmt.Scan(&computer.price)
	fmt.Println("Enter the Color of Computer: ")
	fmt.Scan(&computer.color)
	fmt.Println("Enter the Company Name of conputer")
	fmt.Scan(&computer.CompanyName)
	fmt.Println("Enter the Ramsize of Computer: ")
	fmt.Scan(&computer.Ramsize)

	fmt.Println("After Getting dataa")
	fmt.Println(computer)
}
func main() {
	ReadDataIntoStruct()
}*/

//-----------------------------------------------
//------Anonymous Struct ----------------------

//Its struct without Name Calls as Ananymous struct

func AnonymousStruct() {
	emp1 := struct {
		age    int
		name   string
		salary int
		gender string
	}{age: 20, name: "waseem", salary: 200000, gender: "male"}
	fmt.Println(emp1)

}
func Structmain() {
	AnonymousStruct()
}
