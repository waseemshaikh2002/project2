package main

import (
	"fmt"
	"io/ioutil"
)

// 1) Open File

/*func openFile() {
	const filePath = "/home/waseem2001/Desktop/practicegolang"
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Could Not Open File")
	}
	defer file.Close()
	fmt.Println("File Open SuccessFully")
	fmt.Println(file)
	//if we want to see difrent difrent File Properties
	fmt.Println(file.Name())
	//If we want to see Number of bytes in file
	fmt.Println(file.Stat())

}

func main() {
	openFile()

}*/

// 2) Read data from a file

/*func readFileData() {
	const filePath = "/home/waseem2001/Desktop/practicegolang/2DArrayl17.go"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Could Not Open File", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
func main() {
	readFileData()
}*/

// Read file Data using IO util package

func readDataUsingIoUtilPackage() {
	const filePath = "/home/waseem2001/Desktop/practicegolang/2DArrayl17.go"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Could Not Open file")
		return
	}
	fmt.Println(string(data))

}
func FileIOmain() {
	readDataUsingIoUtilPackage()
}
