package main

/*type friendship interface {
	Goodluck()
	DoEnjoy()
}
type person struct{}

func (p person) Goodluck() {
	fmt.Println("i am Good Talking")
}
func (p person) DoEnjoy() {
	fmt.Println("Doing Enjoy")
}
func Interfacemain() {
	var friend friendship = person{}
	friend.Goodluck()
	friend.DoEnjoy()
}*/

//Assignment on Interface

//Q1 define a strcut named person  with fields name and Age.
// Write a function to create a new person instance and initialize its fiels.

/*type person struct {
	name string
	age  int
}

func (p person) setvalues() {
	var p1 person
	p1.name = "waseem"
	p1.age = 22

	fmt.Println(p1.name)
	fmt.Println(p1.age)
}
func main() {
	var p person
	p.setvalues()
}*/
//-----------------------------------------------------------------------

//Q2 Write a mathod fot the person struct called printdetails that prints
//out the person's name and age.

/*type person struct {
	Name string
	Age  int
}

func (p person) printDetails() {
	p.Name = "Waseem"
	p.Age = 22
	fmt.Println("Person Name :", p.Name)
	fmt.Println("Person Age :", p.Age)
}
func main() {
	var p1 person
	p1.printDetails()
}*/
//---------------------------------------------------------------------

// 3. careate a slice of Person Instance and
//Iterate over it, Printing the detail of each person

/*type person struct {
	Name    string
	Age     int
	Salary  int
	Address string
}

func main() {
	var p1 person
	p1.Name = "Waseem"
	p1.Age = 22
	p1.Salary = 3000000
	p1.Address = "parli"

	var p2 person
	p2.Name = "deepak"
	p2.Age = 22
	p2.Salary = 3000000
	p2.Address = "parli"





	slice := []person{p1, p2}
	for index, value := range slice {
		fmt.Println("Details Of Person", index, value)
	}
}*/

//---------------------------------------------------------------------------

//4
