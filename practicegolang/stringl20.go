package main

import (
	"fmt"
	"strconv"
	"strings"
)

func CreateString() {
	var city string = "codelangs"
	name := "Waseem"
	fmt.Println(city, name)
}
func CalculoateStringLength() {
	var city string = "codelangs"
	length := len(city)
	fmt.Println(length)

}
func StringConCatination() {
	str1 := "Codelangs"
	str2 := " international"
	str3 := str1 + str2 + " Best IT Trainnig Center"
	fmt.Println(str3)
}
func StringIndexing() {
	str := "codelangs"
	char := str[2]
	fmt.Println(string(char))
}
func CreatesliceFormString() {
	str := "codelangs"
	strslice := str[2:]
	fmt.Println(strslice)
}
func IterateOverStringUsingForloopIndexVariable() {

	str := "codelangs"
	for index := 0; index < len(str); index++ {
		fmt.Print(" ")
		fmt.Print(string(str[index]))
	}

}
func IterateOverStirngUsingForrangeLoop() {
	str := "codelangs"
	for _, value := range str {
		fmt.Print(" ")
		char := string(value)
		fmt.Print(char)
	}
	fmt.Println()
}
func CompareString() {
	str1 := "codeLangs"
	str2 := "International"
	Issame := str1 == str2
	fmt.Println(Issame)
	Issame = str1 != str2
	fmt.Println(Issame)
}

// --------------[IMP InterViw Questin]-----
func ConvertStringINtoByteSlice() {
	str1 := "Codelangs"
	byteslice := []byte(str1)
	str2 := string(byteslice)

	fmt.Println(byteslice)
	fmt.Println(str2)

	data := []byte{100, 102, 54, 36, 42}
	str3 := string(data)
	fmt.Println(str3)
}

//-----------------------------------------------

/*func immutableStringDemonstration() {

	//str1 := "codelangs"
	//str1[0] = "x"

}*/

//-----------------[interview Question]--------

func SplitString() {
	str := "Hi Good Morning"
	data := strings.Split(str, " ")
	fmt.Println(data, len(data))
	strData := "10-20-30"
	data = strings.Split(strData, "-")
	fmt.Println(data)
	num1, _ := strconv.Atoi(data[0])
	num2, _ := strconv.Atoi(data[1])
	num3, _ := strconv.Atoi(data[2])
	sum := num1 + num2 + num3
	fmt.Println(sum)

}
func JoinString() {
	slice := []string{
		"pune",
		"mumbai",
		"nashik",
		"parli",
	}
	alldata := strings.Join(slice, "$")
	fmt.Println(alldata)
}
func StringConversion() {
	str1 := "CODELANGS"
	StrLower := strings.ToLower(str1)
	fmt.Println(StrLower)
}
func StringTrimSpace() {
	str := "     welcome     "
	length := len(str)
	fmt.Println(str, length)

	trimStr := strings.TrimSpace(str)
	length = len(trimStr)
	fmt.Println(trimStr, length)

}
func ReadString() {
	str1 := "Codelangs"
	strNew := strings.Repeat(str1, 4)
	fmt.Println(strNew)
}
func CompareStringLexicoGraphicaly() {
	// if str1 > str2 = 1
	// if str1 < str2 = -1
	// if str1 == str2 = 0
	// interview questions
	str1 := "deepak"
	str2 := "waseem"
	result := strings.Compare(str1, str2)
	fmt.Println(result)

}
func Stringmain() {
	//CreateString()
	//CalculoateStringLength()
	//StringConCatination()
	//StringIndexing()
	//CreatesliceFormString()
	//IterateOverStringUsingForloopIndexVariable()
	//IterateOverStirngUsingForrangeLoop()
	//CompareString()
	//ConvertStringINtoByteSlice()
	//SplitString()
	//JoinString()
	//StringConversion()
	//StringTrimSpace()
	//ReadString()
	CompareStringLexicoGraphicaly()

}
