package main

import (
	"fmt"
	"strconv"
)

// primitive data types
// 1) int  2) float 3) bool 4) string 5) rune 6) complex 7)byte
// 1) Intiger to string
func Castingmain() {

	var x int = 5
	strData := strconv.Itoa(x)
	fmt.Println(strData)

}
