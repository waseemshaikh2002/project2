package main

// 2) Write a golang programe to take number from user and display its sum
/*func DisplaySum() {
	var nums [5]int
	var sum = 0
	fmt.Println("Please enter any 5 numbers : ")
	for start := 0; start < 5; start++ {
		fmt.Scan(&nums[start])
		sum = sum + nums[start]
	}
	fmt.Println(sum)
}*/
//============================================
//================[===3===]===================
//Write Golang Program to display Average Of 5 number
//==============================================
/*func FinadAverage() {
	var num [5]int = Scanvalue()
	var sum = sum(num)
	var length = len(num)
	var average = sum / length
	fmt.Println(average)
}
func sum(data [5]int) int {
	var sum = 0
	for start := 0; start < len(data); start++ {
		sum = sum + data[start]
	}
	return sum
}
func Scanvalue() [5]int {
	var data [5]int
	fmt.Println("Please Enter 5 Number")
	for start := 0; start < 5; start++ {
		fmt.Scan(&data[start])
	}
	return data
}*/
//===============[===4===]==============
//-----Write a golang programe to display only even number from user input of 5 number

/*func Scanvalue() [5]int {
	var data [5]int
	fmt.Println("Please Enter 5 Number")
	for start := 0; start < 5; start++ {
		fmt.Scan(&data[start])
	}
	return data
}
func DisplayEvenNumber() {
	var numbers [5]int = Scanvalue()
	for index := 0; index < len(numbers); index++ {
		if numbers[index]%2 == 0 {
			fmt.Printf("\nEven Number is %d,",
				numbers[index])
		}
	}
}*/
//===========================================
//=================[===5===]===
//===========================================
/*func DisplayArrayElement(array [5]int) {
	for index, value := range array {
		fmt.Println(index, value)
	}
}*/

//===============================================
//=================[===6===]====================
//===============================================

// write golang prgorame to check prime number in given number
/*func Scanvalue() [5]int {
	var data [5]int
	fmt.Println("Please Enter 5 Number")
	for start := 0; start < 5; start++ {
		fmt.Scan(&data[start])
	}
	return data
}
func CheckNumberIsprime() {
	var data [5]int = Scanvalue()
	for _, value := range data {
		if IsPrimeNumber(value) {
			fmt.Printf("\n%d is Prime Number", value)
		} else {
			fmt.Printf("\n %d Is Not Prime NUmber", value)
		}
	}
}
func IsPrimeNumber(num int) bool {
	for start := 2; start <= num-1; start++ {
		if num%start == 0 {
			return false
		}
	}
	return true
}*/
//==============================================
//====================[===7===]================
//==============================================
//Write a golang programe to display maximum Number in Array

/*func Scanvalue() [5]int {
	var data [5]int
	fmt.Println("Enter Any Five Element :")
	for start := 0; start < 5; start++ {
		fmt.Scan(&data[start])
	}
	return data

}
func DisplayMaximumNumber() {
	var data [5]int = Scanvalue()
	var max = data[0]
	for index := 1; index < len(data); index++ {
		if max < data[index] {
			max = data[index]
		}
	}
	fmt.Printf("\n %d is Maximum NUmber", max)
}*/

//=================================================
//-----------------------------------------------
//======================[===8===]=================
//Write a golang programe to dispaly squarroot of given 5 number
/*func Scanvalue() [5]int {
	var data [5]int
	fmt.Println("Enter Any 5 Number :")
	for start := 0; start < 5; start++ {
		fmt.Scan(&data[start])
	}
	return data
}
func DispalySquareRoot() {
	var data [5]int = Scanvalue()
	for _, value := range data {
		root := math.Sqrt(float64(value))
		fmt.Printf("\n Root of %d is = %f", value, root)
	}
}*/
//============================================
//------------------------[---9---]------------
//==============================================
/*func ArrayLiteral() {
	var data [5]int = [5]int{10, 20, 30, 40, 50} //array literal
	for index, value := range data {
		fmt.Println(index, value)
	}
}*/

// ================================================
// --------------------------[---10---]-----------
// =================================================
/*func ArrayCompare() {
	array1 := [...]int{10, 20, 30, 40, 50}
	array2 := [...]int{10, 20, 30, 40, 50}

	if array1 == array2 {
		fmt.Println("Both Array Are Same")
	} else {
		fmt.Println("Both Array Are Not Same")
	}

}*/
//===============================================
//------------------[---11---]--------------------
//===============================================
//-------[Copy Array Non-Programitically and Programiticaly]
/*func CopyArrayNonPogramiticaly(){
	array1:=[...]int{10,20,30,40,50}
	array2:=array1
	fmt.Println(array1)
	fmt.Println(array2)
}*/

// ----------Array Function Copy Programitically-------
/*func CopyArrayProgramiticaly() {
	array1 := [...]int{10, 20, 30, 40, 50}
	var array2 [5]int
	for index, value := range array1 {
		array2[index] = value
	}
	fmt.Println(array1)
	fmt.Println(array2)
}*/
//==============================================
//--------------------[---12---]----------------
//===============================================
/*func ArrayWithIndexValueMap() {
	array := [...]int{5: 400, 4: 500, 7: 700, 8: 800}
	fmt.Println(array)
}*/

/*func CompieTimeArrayIndexValueMap() {
	const index = 6
	array := [...]int{index: 800, 4: 400, 3: 300, 1: 1000}
	fmt.Println(array)
}*/

func Arraymain() {

	//===========================================================
	//========Array Practice With sir============================
	//==========================================================

	// 1) Array declaration

	/*var salaries [5]int //array declaration
	fmt.Println(salaries)
	//array assignment
	salaries[0] = 200
	salaries[1] = 400
	salaries[2] = 600
	salaries[3] = 800
	salaries[4] = 1000
	fmt.Println("----------------------------")
	fmt.Println(salaries)*/
	//--------------------
	//------------[2]------------
	//DisplaySum()
	//==========================
	//---------[3]--------------
	//FinadAverage()
	//================================
	//===========[=== 4 ===]===============

	//DisplayEvenNumber()
	//====================================
	//==============[===5===]===========
	/*var data [5]int
	data[0] = 100
	data[1] = 200
	data[2] = 300
	data[3] = 400
	data[4] = 500
	DisplayArrayElement(data)*/
	//=================================
	//===========[===6===]============
	/*CheckNumberIsprime()*/
	//=================================
	//=========[7]====================
	//DisplayMaximumNumber()
	//--------------------------
	//==============[===8===]============
	//DispalySquareRoot()
	//================================
	//---------------[---9---]----------

	//===================================
	//ArrayLiteral()
	//===================================
	//--------------------[10]------------
	//ArrayCompare()
	//------------------------------------
	//====================================
	//----------------[---11---]-----------
	//=======================================

	//CopyArrayProgramiticaly()
	//=======================================
	//-----------------[---12---]----------
	//=======================================
	//ArrayWithIndexValueMap()
	//CompieTimeArrayIndexValueMap()

}
