package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!") // This is a comment
	/* The code below will print Hello World
	to the screen, and it is amazing */

	variables()
	constants()
	output()
	datetypes()
	arrays()
	operators()
	conditions()
	switchs()
	loops()
	functions()
	structs()
	maps()
}

func maps() {
	/*
		Create Maps Using var and :=
		Syntax
		var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
		b := map[KeyType]ValueType{key1:value1, key2:value2,...}
	*/
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"}
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)

	/*
		Create Maps Using make()Function:
		Syntax
		var a = make(map[KeyType]ValueType)
		b := make(map[KeyType]ValueType)
	*/
	var c = make(map[string]string) // The map is empty now
	c["brand"] = "Ford"
	c["model"] = "Mustang"
	c["year"] = "1964"
	// c is no longer empty
	d := make(map[string]int)
	d["Oslo"] = 1
	d["Bergen"] = 2
	d["Trondheim"] = 3
	d["Stavanger"] = 4

	fmt.Printf("c\t%v\n", c)
	fmt.Printf("d\t%v\n", d)

	/*
		Create an Empty Map
		There are two ways to create an empty map. One is by using the make()function and the other is by using the following syntax.

		Syntax
		var a map[KeyType]ValueType
	*/
	// Note: The make()function is the right way to create an empty map. If you make an empty map in a different way and write to it, it will causes a runtime panic.
	var e = make(map[string]string)
	var f map[string]string
	// f["name"] = "David" // Error

	fmt.Println(e == nil) // false
	fmt.Println(f == nil) // true
	fmt.Printf("f\t%v\n", f)

	/*
		Allowed Key Types
		The map key can be of any data type for which the equality operator (==) is defined. These include:

		Booleans
		Numbers
		Strings
		Arrays
		Pointers
		Structs
		Interfaces (as long as the dynamic type supports equality)
		Invalid key types are:

		Slices
		Maps
		Functions
		These types are invalid because the equality operator (==) is not defined for them.

		Allowed Value Types
		The map values can be any type.
	*/

	// Remove Element from Map
	/*
		Removing elements is done using the delete() function.

		Syntax
		delete(map_name, key)
	*/
	fmt.Println(a)
	delete(a, "year")
	fmt.Println(a)

	/*
		Check For Specific Elements in a Map
		You can check if a certain key exists in a map using:

		Syntax
		val, ok :=map_name[key]
	*/
	val1, ok1 := a["brand"] // Checking for existing key and its value
	val2, ok2 := a["color"] // Checking for non-existing key and its value
	val3, ok3 := a["day"]   // Checking for existing key and its value
	_, ok4 := a["model"]    // Only checking for existing key and not its value

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

	/*
		Maps Are References
		Maps are references to hash tables.

		If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.
	*/

	//	Iterate Over Maps
	// You can use range to iterate over maps.
	g := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	for k, v := range g {
		fmt.Printf("%v : %v, ", k, v)
	}

	// Iterate Over Maps in a Specific Order
	// Maps are unordered data structures. If you need to iterate over a map in a specific order, you must have a separate data structure that specifies that order.
	h := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var i []string // defining the order
	i = append(i, "one", "two", "three", "four")

	for k, v := range h { // loop with no order
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println()

	for _, element := range i { // loop with the defined order
		fmt.Printf("%v : %v, ", element, g[element])
	}
}

func structs() {
	/*
	   To declare a structure in Go, use the type and struct keywords:

	   Syntax
	   type struct_name struct {
	     member1 datatype;
	     member2 datatype;
	     member3 datatype;
	     ...
	   }
	*/
	type Person struct {
		name   string
		age    int
		job    string
		salary int
	}

	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2 specification
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary = 4500

	// Access and print Pers1 info
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)

	// Access and print Pers2 info
	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)

}

func functions() {
	/*
		Naming Rules for Go Functions

		A function name must start with a letter
		A function name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ )
		Function names are case-sensitive
		A function name cannot contain spaces
		If the function name consists of multiple words, techniques introduced for multi-word variable naming can be used
		Tip: Give the function a name that reflects what the function does!

		Syntax
		func FunctionName(param1 type, param2 type, param3 type) {
		  // code to be executed
		}
	*/

	familyName("Liam")
	familyName("Jenny")
	familyName("Anja")

	// Named Return Values
	fmt.Println(namedReturnFunction(1, 2))

	// Multiple Return Values
	// Here, we store the two return values into two variables (a and b):
	a, b := multipleReturnFunction(5, "Hello")
	fmt.Println(a, b)

	// If we (for some reason) do not want to use some of the returned values, we can add an underscore (_), to omit this value.
	_, c := multipleReturnFunction(5, "Hello")
	fmt.Println(c)
	d, _ := multipleReturnFunction(5, "Hello")
	fmt.Println(d)
}

func multipleReturnFunction(x int, y string) (result int, txt1 string) {
	result = x + x
	txt1 = y + " World!"
	return
}

func namedReturnFunction(x int, y int) (result int) {
	result = x + y
	return
}

func familyName(fname string) int {
	fmt.Println("Hello", fname, "Refsnes")
	return 1
}

func loops() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// The Range Keyword
	/*
		The range keyword is used to more easily iterate over an array, slice or map. It returns both the index and the value.

		The range keyword is used like this:

		Syntax
		for index, value := array|slice|map {
		   // code to be executed for each iteration
		}
	*/
	fruits := [3]string{"apple", "orange", "banana"}
	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	}

	// Tip: To only show the value or the index, you can omit the other output using an underscore (_).
	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}

	for idx, _ := range fruits {
		fmt.Printf("%v\n", idx)
	}

}

func switchs() {
	day := 4

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
		// default:
		// 	fmt.Println("Not a weekday")
	}

	// The Multi-case switch Statement
	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekday")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}
}

func conditions() {
	// The if Statement
	if 20 > 18 {
		fmt.Println("20 is greater than 18")
	}
}

func operators() {
}

func arrays() {
	// This example declares two arrays (arr1 and arr2) with defined lengths
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)

	// This example declares two arrays (arr1 and arr2) with inferred lengths
	var arr3 = [...]int{1, 2, 3}
	arr4 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr3)
	fmt.Println(arr4)

	// This example declares an array of strings
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Print(cars, "\n")

	// This example shows how to access the first and third elements in the prices array
	prices := [3]int{10, 20, 30}

	fmt.Println(prices[0])
	fmt.Println(prices[2])

	// This example shows how to change the value of the third element in the prices array
	prices[2] = 50
	fmt.Println(prices)

	/*
		If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.
		Tip: The default value for int is 0, and the default value for string is "".
	*/
	arr5 := [5]int{}              //not initialized
	arr6 := [5]int{1, 2}          //partially initialized
	arr7 := [5]int{1, 2, 3, 4, 5} //fully initialized

	fmt.Println(arr5)
	fmt.Println(arr6)
	fmt.Println(arr7)

	// It is possible to initialize only specific elements in an array.
	arr8 := [5]int{1: 10, 2: 40}
	fmt.Println(arr8)
	fmt.Println(len(arr8)) // The len() function is used to find the length of an array

	slices()
}

func slices() {
	/*
		In Go, there are two functions that can be used to return the length and capacity of a slice:

		len() function - returns the length of the slice (the number of elements in the slice)
		cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)
	*/
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	// This example shows how to create a slice from an array
	// 	var myarray = [length]datatype{values} // An array
	// myslice := myarray[start:end] // A slice made from the array
	arr1 := [6]int{10, 11, 12, 13, 14, 15}
	myslice := arr1[2:4]

	fmt.Printf("myslice = %v\n", myslice)
	fmt.Printf("length = %d\n", len(myslice))
	fmt.Printf("capacity = %d\n", cap(myslice))

	// Create a Slice With The make() Function
	// slice_name := make([]type, length, capacity)
	// Note: If the capacity parameter is not defined, it will be equal to length.
	myslice3 := make([]int, 5, 10)
	fmt.Printf("myslice3 = %v\n", myslice3)
	fmt.Printf("length = %d\n", len(myslice3))
	fmt.Printf("capacity = %d\n", cap(myslice3))

	// with omitted capacity
	myslice4 := make([]int, 5)
	fmt.Printf("myslice4 = %v\n", myslice4)
	fmt.Printf("length = %d\n", len(myslice4))
	fmt.Printf("capacity = %d\n", cap(myslice4))

	// Append Elements To a Slice
	// slice_name = append(slice_name, element1, element2, ...)
	myslice5 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice5 = %v\n", myslice5)
	fmt.Printf("length = %d\n", len(myslice5))
	fmt.Printf("capacity = %d\n", cap(myslice5))

	myslice5 = append(myslice5, 20, 21)
	fmt.Printf("myslice5 = %v\n", myslice5)
	fmt.Printf("length = %d\n", len(myslice5))
	fmt.Printf("capacity = %d\n", cap(myslice5))

	// Append One Slice To Another Slice
	// slice3 = append(slice1, slice2...)
	// Note: The '...' after slice2 is necessary when appending the elements of one slice to another.
	myslice6 := []int{1, 2, 3}
	myslice7 := []int{4, 5, 6}
	myslice8 := append(myslice6, myslice7...)

	fmt.Printf("myslice8=%v\n", myslice8)
	fmt.Printf("length=%d\n", len(myslice8))
	fmt.Printf("capacity=%d\n", cap(myslice8))

	// Change The Length of a Slice
	arr10 := [6]int{9, 10, 11, 12, 13, 14} // An array
	myslice10 := arr10[1:5]                // Slice array
	fmt.Printf("myslice10 = %v\n", myslice10)
	fmt.Printf("length = %d\n", len(myslice10))
	fmt.Printf("capacity = %d\n", cap(myslice10))

	myslice10 = arr1[1:3] // Change length by re-slicing the array
	fmt.Printf("myslice10 = %v\n", myslice10)
	fmt.Printf("length = %d\n", len(myslice10))
	fmt.Printf("capacity = %d\n", cap(myslice10))

	myslice10 = append(myslice10, 20, 21, 22, 23) // Change length by appending items
	fmt.Printf("myslice10 = %v\n", myslice10)
	fmt.Printf("length = %d\n", len(myslice10))
	fmt.Printf("capacity = %d\n", cap(myslice10))

	// Memory Efficiency
	// copy(dest, src)
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	// Original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	fmt.Printf("neededNumbers = %v\n", neededNumbers)
	fmt.Printf("length = %d\n", len(neededNumbers))
	fmt.Printf("capacity = %d\n", cap(neededNumbers))

	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}

func datetypes() {
	var a bool = true    // Boolean
	var b int = 5        // Integer
	var c float32 = 3.14 // Floating point number
	var d string = "Hi!" // String

	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("Float:   ", c)
	fmt.Println("String:  ", d)

	// Boolean Data Type
	var b1 bool = true // typed declaration with initial value
	var b2 = true      // untyped declaration with initial value
	var b3 bool        // typed declaration without initial value
	b4 := true         // untyped declaration with initial value

	fmt.Println(b1) // Returns true
	fmt.Println(b2) // Returns true
	fmt.Println(b3) // Returns false
	fmt.Println(b4) // Returns true

	// Signed Integers
	/*
		Go has five keywords/types of signed integers:

		Type	Size	Range
		int	Depends on platform:
		32 bits in 32 bit systems and
		64 bit in 64 bit systems	-2147483648 to 2147483647 in 32 bit systems and
		-9223372036854775808 to 9223372036854775807 in 64 bit systems
		int8	8 bits/1 byte	-128 to 127
		int16	16 bits/2 byte	-32768 to 32767
		int32	32 bits/4 byte	-2147483648 to 2147483647
		int64	64 bits/8 byte	-9223372036854775808 to 9223372036854775807
	*/

	var x int = 500
	var y int16 = -4500
	fmt.Printf("Type: %T, value: %v", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)

	// Unsigned Integers
	/*
		Go has five keywords/types of unsigned integers:

		Type	Size	Range
		uint	Depends on platform:
		32 bits in 32 bit systems and
		64 bit in 64 bit systems	0 to 4294967295 in 32 bit systems and
		0 to 18446744073709551615 in 64 bit systems
		uint8	8 bits/1 byte	0 to 255
		uint16	16 bits/2 byte	0 to 65535
		uint32	32 bits/4 byte	0 to 4294967295
		uint64	64 bits/8 byte	0 to 18446744073709551615
	*/
	var xx uint = 500
	var yy uint = 4500
	fmt.Printf("Type: %T, value: %v", xx, xx)
	fmt.Printf("Type: %T, value: %v", yy, yy)

	/*
		Go Float Data Types
		The float data types are used to store positive and negative numbers with a decimal point, like 35.3, -2.34, or 3597.34987.

		The float data type has two keywords:

		Type	Size	Range
		float32	32 bits	-3.4e+38 to 3.4e+38.
		float64	64 bits	-1.7e+308 to +1.7e+308.

		Tip: The default type for float is float64. If you do not specify a type, the type will be float64.
	*/
	var x1 float32 = 123.78
	var y1 float32 = 3.4e+38
	fmt.Printf("Type: %T, value: %v\n", x1, x1)
	fmt.Printf("Type: %T, value: %v", y1, y1)

	var x2 float64 = 1.7e+308
	fmt.Printf("Type: %T, value: %v\n", x2, x2)

	// String Data Type
	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World 1"

	fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)
}

func output() {
	var i, j string = "Hello", "World"

	fmt.Print(i)
	fmt.Print(j)

	// If we want to print the arguments in new lines, we need to use \n.
	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	// It is also possible to only use one Print() for printing multiple variables.
	fmt.Print(i, "\n", j)

	// Print() inserts a space between the arguments if neither are strings:
	var a, b = 10, 20
	fmt.Print(a, b)

	/*
		The Println() function is similar to Print() with the difference
		that a whitespace is added between the arguments, and a newline is added at the end.
	*/
	fmt.Println(i, j)

	/*
		The Printf() function first formats its argument based on the given formatting verb and then prints them.

		Here we will use two formatting verbs:

		%v is used to print the value of the arguments
		%T is used to print the type of the arguments
	*/
	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("j has value: %v and type: %T", j, j)

	/*
		General Formatting Verbs
		The following verbs can be used with all data types:

		Verb	Description
		%v	Prints the value in the default format
		%#v	Prints the value in Go-syntax format
		%T	Prints the type of the value
		%%	Prints the % sign
	*/

	var ii = 15.5
	var txt = "Hello World!"

	fmt.Printf("%v\n", ii)
	fmt.Printf("%#v\n", ii)
	fmt.Printf("%v%%\n", ii)
	fmt.Printf("%T\n", ii)

	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)

	/*
	   Integer Formatting Verbs
	   The following verbs can be used with the integer data type:

	   Verb	Description
	   %b	Base 2
	   %d	Base 10
	   %+d	Base 10 and always show sign
	   %o	Base 8
	   %O	Base 8, with leading 0o
	   %x	Base 16, lowercase
	   %X	Base 16, uppercase
	   %#x	Base 16, with leading 0x
	   %4d	Pad with spaces (width 4, right justified)
	   %-4d	Pad with spaces (width 4, left justified)
	   %04d	Pad with zeroes (width 4
	*/
	var iii = 15

	fmt.Printf("%b\n", iii)
	fmt.Printf("%d\n", iii)
	fmt.Printf("%+d\n", iii)
	fmt.Printf("%o\n", iii)
	fmt.Printf("%O\n", iii)
	fmt.Printf("%x\n", iii)
	fmt.Printf("%X\n", iii)
	fmt.Printf("%#x\n", iii)
	fmt.Printf("%4d\n", iii)
	fmt.Printf("%-4d\n", iii)
	fmt.Printf("%04d\n", iii)

	/*
		String Formatting Verbs
		The following verbs can be used with the string data type:

		Verb	Description
		%s	Prints the value as plain string
		%q	Prints the value as a double-quoted string
		%8s	Prints the value as plain string (width 8, right justified)
		%-8s	Prints the value as plain string (width 8, left justified)
		%x	Prints the value as hex dump of byte values
		% x	Prints the value as hex dump with spaces
	*/
	var txt1 = "Hello"

	fmt.Printf("%s\n", txt1)
	fmt.Printf("%q\n", txt1)
	fmt.Printf("%8s\n", txt1)
	fmt.Printf("%-8s\n", txt1)
	fmt.Printf("%x\n", txt1)
	fmt.Printf("% x\n", txt1)

	/*
		Boolean Formatting Verbs
		The following verb can be used with the boolean data type:

		Verb	Description
		%t	Value of the boolean operator in true or false format (same as using %v)
	*/
	var t = true
	var f = false

	fmt.Printf("%t\n", t)
	fmt.Printf("%t\n", f)

	/*
		Float Formatting Verbs
		The following verbs can be used with the float data type:

		Verb	Description
		%e	Scientific notation with 'e' as exponent
		%f	Decimal point, no exponent
		%.2f	Default width, precision 2
		%6.2f	Width 6, precision 2
		%g	Exponent as needed, only necessary digits
	*/
	var iiii = 31111.141111111

	fmt.Printf("%e\n", iiii)
	fmt.Printf("%f\n", iiii)
	fmt.Printf("%.2f\n", iiii)
	fmt.Printf("%6.2f\n", iiii)
	fmt.Printf("%g\n", iiii)
}

func constants() {
	const PI = 3.14

	/*
		Constant Rules:

		Constant names follow the same naming rules as variables
		Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
		Constants can be declared both inside and outside of a function

		Constant Types:
		There are two types of constants:

		Typed constants
		Untyped constants
	*/

	const A int = 1 // Typed constants

	const B = 1 // Untyped constants

	// Multiple Constants Declaration
	const (
		C int = 1
		D     = 3.14
		E     = "Hi!"
	)
}

// var a string //OK
// a := 1 //Error
func variables() {
	fmt.Println("Variable!")

	// with initial value
	// You always have to specify either type or value (or both).
	//1. Can be used inside and outside of functions
	var var1 string = "str1"
	//2. Variable declaration and value assignment can be done separately
	var var2 string
	var2 = "str2"

	/*
		1. Can only be used inside functions
		2. Variable declaration and value assignment cannot be done separately (must be done in the same line)
	*/
	var3 := 100

	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)

	// without initial value
	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// Multiple Variable Declaration
	var d, e, f, g int = 1, 3, 5, 7

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	/*
		If you use the type keyword, it is only possible to declare one type of variable per line.
		If the type keyword is not specified, you can declare different types of variables in the same line.
	*/
	var h, i = 6, "Hello"
	j, k := 7, "World!"

	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	// Variable Declaration in a Block
	var (
		l int
		m int    = 1
		n string = "hello"
	)

	fmt.Println(l)
	fmt.Println(m)
	fmt.Println(n)

	// Variable Naming Rules
	/*
		Go variable naming rules:

		A variable name must start with a letter or an underscore character (_)
		A variable name cannot start with a digit
		A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
		Variable names are case-sensitive (age, Age and AGE are three different variables)
		There is no limit on the length of the variable name
		A variable name cannot contain spaces
		The variable name cannot be any Go keywords
	*/

	camelCase := "John"
	PascalCase := "John"
	snake_case := "John"

	fmt.Println(camelCase)
	fmt.Println(PascalCase)
	fmt.Println(snake_case)
}
