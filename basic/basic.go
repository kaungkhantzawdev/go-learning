package main

import (
	"fmt"
	"strings"
)

// VARIABLES
var a string;
var f float32 = 123.45;

// MULTIPLE DECLARE
var e, g string = "e declare", "g declare";

// CONSTANT
const PI float32 = 3.14;

var (
	h int = 1
	j string = "j declare"
	i float64 = 234.34
)

// DATA TYPES
var first int = 3;
var firstOne int = -3;
var second string = "hello welcome";
var third bool = true;
var four float32 = 3.45;
var five uint = 34;

// ARRAY
var arrFirst = [3]int{1, 2, 3}
var arrSecond = [5]string{"car", "home", "phone", "ipad", "mac"}

// SLICE
var arr1 = [5]string{"car", "home", "phone", "ipad", "mac"};
var slice1 = arr1[1:3];

// OPERATORS
var x int = 10;
var ab int = 1 + 2;
var bc int = 1 * 2;
var cd int = 1 - 2;
var ef int = 1 / 2;

func hi() {
	fmt.Println("hi");
}

func dataTypes() {

	// Only use inside function
	b := 23.3;
	a = "hello var keywords";

	// Array update
	arrFirst[0] = 10;

	// Operators
	x++;
	x--;
	x += 5;
	x = x ^ 3
	x >>= 3
	x &= 3
	x *= 3

	// Print
	fmt.Printf("%v and %0.1f \n", a, b);
	fmt.Println(a, f, e, g, h, i, j, PI);
	fmt.Println(slice1)

}


// STATEMENT
func stateMent() {
	// IF
	m := 3;

	if m == 4 {
		fmt.Println("m is 4");
	} else {
		fmt.Println("m is ", m)
	}

	// SWITCH 
	x := 3;
	switch x {
		case 3:
			fmt.Println("x is 3 from switch statement");
		case 4:
			fmt.Println("x is 4 from switch statement");
		default:
			fmt.Println("x is", x, "from switch")
	}
	// MULTI CASE
	b := 3;
	switch b {
		case 1,3,5:
			fmt.Println("b is odd from switch statement");
		case 2,4:
			fmt.Println("b is even from switch statement");
		default:
			fmt.Println("b is", b, "from switch")
	}
}

// LOOP
func loopState() {

	// FOR LOOP
	for i :=0; i < 5; i++ {
		fmt.Println("i is", i);
	}

	// RANGE
	for index, value := range arr1 {
		fmt.Printf("Index is %v and value is %v \n", index, value);
	}
}

// FUNCTIONS
func test(a string, b int) string {
	c := strings.ToUpper(a);
	d := fmt.Sprintf("%v and int is %v", c, b)
	return d
}

// STRUCK
type Person struct {
	name string
	age int
	job bool
	home bool
	
}

func printPerson(pers Person) {
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.home)
}

func showYour() {
	var person1 Person
	var person2 Person

	person1.name = "one"
	person1.age = 20
	person1.job = true
	person1.home = true

	person2.age = 22
	person2.name = "two"
	person2.job = true
	person2.home = true

	printPerson(person1)
	printPerson(person2)
}

// MAP
var aMap = map[string]int{
	"one": 1,
	"two": 2,
}
func testMap() {

	aMap["three"] = 3
	aMap["four"] = 4

	delete(aMap, "one")

	v, ok := aMap["two"]

	fmt.Println(aMap, v, ok)

	for k, v := range aMap {
		fmt.Printf("%v : %v, ", k, v)
	}
}

func main() {
	dataTypes()
	stateMent()
	loopState()
	showYour()
	testMap()
	
	fmt.Println("hello go. nice to meet you", test("hi hello", 23))
}

