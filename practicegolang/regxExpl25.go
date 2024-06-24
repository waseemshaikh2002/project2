package main

import (
	"fmt"
	"regexp"
)

/*func main() {
	var mobleNumber string
	fmt.Println("Enter Your Mobile Number :")
	fmt.Scan(&mobleNumber)
	//Pattern For Mobile Number
	pattern := `^[0-9]\d{9}$`
	reg, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Invalid Pattern")
		return
	}
	isValidMobileNo := reg.MatchString(mobleNumber)
	if isValidMobileNo {
		fmt.Println("Valid Mobile Number")

	} else {
		fmt.Println("Invalid Mobile Number")
	}

}*/

//--------------------------------------------------
//-----------------Email id pattern--------------

func isValidEmail(email string) bool {
	emailpattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	reg, err := regexp.Compile(emailpattern)
	if err != nil {
		fmt.Println("Invalid Email Pattern")
		return false

	}
	isValidEmailid := reg.MatchString(email)
	return isValidEmailid
}

func Regxpmain() {
	var email string
	fmt.Println("Enter Your Email id :")
	fmt.Scan(&email)
	if isValidEmail(email) {
		fmt.Println("Valid Email Id")
	} else {
		fmt.Println("Invalid Email Id")
	}
}
