package main

// we create a struct and its object

/*type Employee struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	var employee Employee
	employee.Name = "raju"
	employee.Age = 25
	employee.Salary = 240000

	fmt.Println(employee)// employee its our object
	out put : {waseem 23 200000}


}*/
//-----------------------------------------------------

// json marshalling
// to conver the jeson format we have the inbuilt api
//That is called as json package
//encode the json format its called as json marchalling
// package "encoding/json" it will be automatically import

// for the marshalling we have json.marshal() method

//marshal method will take argument of any type
// and it will be retrun byte array and error
//Nyte array means jsonData and error means error
//It will give the out put in byte array

/*type Employee struct {
	Name   string
	Age    int
	Salary float64
}

func main() {
	var employee Employee
	employee.Name = "waseem"
	employee.Age = 23
	employee.Salary = 200000

	fmt.Println(employee)
	// json marchalling
	//we want to convert our golang object in json format
	// our object is employee
	//we can marshalling using the json inbuilt package
	// json ts format and marshal is method

	jsonData, err := json.Marshal(employee)
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Println(jsonData)
	//out put after converting object to json
	// out put : [123 34 78 97 109 101 34 58 34 119 97 115 101 101 109 34 44 34 65 103 101 34 58 50 51 44 34 83 97 108 97 114 121 34 58 50 48 48 48 48 48 125]
	// we got the out out he conveted in json and given the out out in byte array
	// we want to conver the json result in string format
	// because we can understand the string data
	// byte to string conversion

	fmt.Println(string(jsonData))

	//out put after converted json byte array to string its called as json marshalling
	// out put : {"Name":"waseem","Age":23,"Salary":200000}
	// we converted out object in json format
	// this json we can send to othe user or othe system.
	// after that they will decode our json and it will be read

}*/

//-----------------------------------------------------
//--------------Unmarshalling-------------------------

//convert(decode) json data into golang object its called as unmarshalling
// josn data {"Name":"waseem","Age":23,"Salary":200000}
// for the decoding we want to stoe this data in variable
// actually we ecoding the data we are taking data from another service and we can decode.
// json data always be in string format.
// {"Name":"waseem","Age":23,"Salary":200000}
// name, age , salary these are our keys.
// waseem, 23, 200000 there are our values
// accorging to keys and data type we want to create a struct like.

/*type Employee struct {
	Name   string
	Age    int
	Salary float64
	// we will convet to te this string  back to this string
}

func unmarshallingjson() {
	// our json string
	var jsonstring = `{"Name":"waseem","Age":23,"Salary":200000}`
	// struct object.
	var employee Employee
	// we have the method for unmarshalling its unmarhsal().
	// pass the first argument slice of byte.
	// and second argument is any. it will be retrun error if we are gettin eror
	// our json is string if we will pass the sting it will give the error
	// we want to conver string to byte slice.
	err := json.Unmarshal([]byte(jsonstring), &employee)
	if err != nil {
		fmt.Println("Error in Decoding Json")
		return

	}
	// unmarshal method it will be decode our json string and kept in emplooye
	//employee will be match our propery name age salary if its 100% mathing after that it will be decode

	fmt.Println(jsonstring) // it will prinnt json string
	// it will be convert and print its object
	// if we want o print seperate out put
	fmt.Println(employee)
	fmt.Println(employee.Name)
	fmt.Println(employee.Age)
	fmt.Println(employee.Salary)

}
func marchallingjson() {
	var employee Employee
	employee.Name = "waseem"
	employee.Age = 23
	employee.Salary = 200000

	fmt.Println(employee)
	// json marshalling
	jsonData, err := json.Marshal(employee)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(jsonData)
	fmt.Println(string(jsonData))

}
func main() {
	unmarshallingjson()

}*/

//-----------------------------------------------------
//-----------------Json Array------------------------

/*func jsonArrayunmarshalling() {
	type Location struct {
		Address string
		City    string
	}
	var location []Location
	var data = `[
	{"Adress":"shivne","City":"Pune"},
	{"Adress":"kothrud","City":"Pune"},
	{"Adress":"shivaji Nagar","City":"Pune"}]`
	//we want to conver data in byte code
	byteData := []byte(data)

	err := json.Unmarshal(byteData, &location)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, item := range location {
		fmt.Println(item.Address, item.City)

	}
}
func main() {
	jsonArrayunmarshalling()
}*/

//---------------------------------------------------------------------

//------------------------Custom-Json-Marshelling--------
//--------------------------------------------------
// Actuaaly inn this sinario we dont know about  whch value is stored in json
//and how to conver this structure and bined
// in this type json we will directly conver in our object
// without knowing which type of structure
// this concept is called as custom marshalling.
// in this type key always be string and value will be Any Type.

/*func CustomjsonunMarshelling() {
	var data = `[
	{"Address":"shivne","City":"Pune"},
	{"Address":"kothrud","City":"Pune"},
	{"Address":"shivaji Nagar","City":"Pune"}]`

	var decode []map[string]interface{}
	err := json.Unmarshal([]byte(data), &decode)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(decode)
	for _, item := range decode {
		fmt.Println(item["Address"], item["City"])
	}

}

func main() {
	CustomjsonunMarshelling()
}*/

//-----------------------------------------------------
//-------------------Custom Json Marshalling------------

/*func customjsonMarshalling() {
	var emplooye map[string]interface{} = make(map[string]interface{})
	emplooye["Name"] = "Raju"
	emplooye["Age"] = 30
	emplooye["city"] = "Pune"
	jsonByte, err := json.Marshal(emplooye)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonstring := string(jsonByte)
	fmt.Println(jsonstring)
}
func main() {
	customjsonMarshalling()
}*/

//----------------------------------------------------
//-----------Marshalling Embedded struct-------

/*func marshallingEmbeddedStruct() {
	type Human struct {
		Weight float64
		Height float64
	}
	type Employee struct {
		Name string
		Age  int
		Human
	}
	var employee Employee
	employee.Height = 6.7
	employee.Weight = 56.6
	employee.Name = "Pinky"
	employee.Age = 20

	jsonByte, err := json.Marshal(employee)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(jsonByte))

}
func main() {
	marshallingEmbeddedStruct()
}*/

//------------------------------------------------------
//--------------Decode embedded struct unmarshalling---------------
//--------------------------------------------------------

/*func EmbeddedStructUnMarshalling() {
	type Human struct {
		Weight float64
		Height float64
	}
	type Employee struct {
		Name string
		Age  int
		Human
	}
	data := `{"Name":"Pinky","Age":20,"Weight":56.6,"Height":6.7}`
	var employee Employee
	err := json.Unmarshal([]byte(data), &employee)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(employee.Age)
	fmt.Println(employee.Height)
	fmt.Println(employee.Name)
	fmt.Println(employee.Weight)

}
func Jsonmain() {
	EmbeddedStructUnMarshalling()
}*/

//----------------------------------------------------------------------------
// ------------ Struct Tags Json Marhsalling--------

// 1) Field Main

/*
	type Person struct {
		Name string
		Age  int
	}

	func main() {
		var person Person = Person{Name: "Waseem", Age: 22}
		// encoding/json
		// Converting Json in To the struct
		jsonByte, _ := json.Marshal(person)

		fmt.Println(string(jsonByte))
	}
*/
/*type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var person Person = Person{Name: "waseem", Age: 23}
	jsonByte, _ := json.Marshal(person)
	fmt.Println(string(jsonByte))
}*/

//---------------------------------------------------------
//---------------------OmIting The fields-----------------
