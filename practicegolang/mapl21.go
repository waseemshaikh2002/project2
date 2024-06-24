package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*func mapDecleration() {
	var salaries map[string]int //empty or nill map
	fmt.Println(salaries)
}
func main() {
	mapDecleration()
}*/
//---------------------------------------
/*func createMap() {
	var salaries map[string]int = make(map[string]int)
	salaries["waseem"] = 200000
	salaries["deepak"] = 210000
	salaries["aksahy"] = 220000
	fmt.Println(salaries)
}

func main() {
	createMap()

}*/
//--------------------------------------------
/*func updateValues() {
	var salaries map[string]int = make(map[string]int)
	salaries["Deepak"] = 200000
	salaries["aksahy"] = 210000
	salaries["waseem"] = 220000
	fmt.Println(salaries)
	fmt.Println("------After Update-----")

	salaries["deepak"] = 220000
	fmt.Println(salaries)

}

func main() {
	updateValues()
}*/
/*func deleteEntryFromMap() {
	var salaries map[string]int = make(map[string]int)
	salaries["Waseem"] = 2200000
	salaries["deepak"] = 2300000
	salaries["akshay"] = 2500000
	fmt.Println(salaries)
	fmt.Println("----After delete-------")
	delete(salaries, "akshay")
	fmt.Println(salaries)
}

func main() {
	deleteEntryFromMap()
}*/

/*func existanceMap() bool {
	var salaries map[string]int = make(map[string]int)
	salaries["waseem"] = 400000
	salaries["deepak"] = 500000
	salaries["akshay"] = 600000
	fmt.Println(salaries)
	_, isExist := salaries["waseem"]
	//_, waseemExist := salaries["waseem"]
	//_, deepakExist := salaries["deepak"]
	//fmt.Println("waseem is Exist")
	if isExist {
		fmt.Println("Given name is exist")
	} else {
		fmt.Println("Name is NOT exist")
	}
	return isExist

}
func main() {
	existanceMap()
}*/

/*func LengthOfMap() {
	var salaries map[string]int = make(map[string]int)
	salaries["waseem"] = 12000000
	salaries["deepak"] = 23000000
	salaries["akshay"] = 34000000
	fmt.Println(salaries)
	lenght := len(salaries)
	fmt.Println("Length Of Map is: ", lenght)
}
func main() {
	LengthOfMap()
}*/

/*func IteratingOverMap() {
	var salaries map[string]int = make(map[string]int)
	salaries["waseem"] = 12300000
	salaries["deepak"] = 13200000
	salaries["akshay"] = 14200000
	fmt.Println(salaries)
	for key, value := range salaries {
		fmt.Println(key, value)
	}
}
func main() {
	IteratingOverMap()
}*/

/*func AssesingKayValueFromMap() {
	var salaries map[string]int = make(map[string]int)
	salaries["waseem"] = 2300000
	salaries["deepak"] = 4500000
	salaries["akshay"] = 5400000
	fmt.Println(salaries)
	salary := salaries["waseem"] //Asssing key value from map
	fmt.Println(salary)
}
func main() {
	AssesingKayValueFromMap()
}*/

/*func MapLiteral() {
	student := map[int]string{
		101: "Raju",
		102: "Rani",
		103: "rama",
		105: "Michel",
	}
	for key, value := range student {
		fmt.Println(key, value)
	}
	fmt.Println(student)
	fmt.Println(student[102])
}
func main() {
	MapLiteral()
}*/

/*func DisplayMapEntries(data map[int]string) {
	for key, value := range data {
		fmt.Println(key, value)
	}
}

func PassingMapToFunction() {
	students := map[int]string{
		101: "raju",
		102: "rani",
		103: "michel",
	}
	DisplayMapEntries(students)
}

func main() {
PassingMapToFunction()
}*/

/*
	func ReturningMapFromFunction() map[int]string {
		students := map[int]string{
			101: "raju",
			102: "chutki",
			103: "chotta bheem",
			104: "jaggu",
		}
		return students
	}

func main() {

		//ReturningMapFromFunction()
	}
*/
func Dataset() map[string]string {
	data := map[string]string{
		"hi ":                     "hi good morning , how may i help you ?",
		"what is jio care number": " call on toll free 0042157895",
		"internet is not working": "please call on our service center on 999",
		"thank you":               " same to you. please feel free to ask me !",
		"want to call manager":    "sorry this is not allowed you can call to customer care",
	}
	return data
}
func chatBoat() {
	var input string
	message := "How May I Help You!...\t"
	fmt.Println(message)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		input = scanner.Text()
		data := Dataset()
		for key, value := range data {
			if strings.Contains(key, input) {
				fmt.Println(value)
				break
			}
		}
	}
}
func Mapmain() {
	chatBoat()
}
