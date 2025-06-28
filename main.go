/*
---------------------------------
Go Basics Practice
---------------------------------
*/

package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"time"
)

// MAIN __________________________________________________________________________________________________________________________________________________________________________________________________________________
func main() {
    fmt.Println("--------------------------------MAIN-----------------------------------")

	// Print a welcome message with strings package fmt
	fmt.Println("Welcome to the playground!")

	// Using the time package to print the current time
	fmt.Println("The time is", time.Now())

	// Using the math package to perform calculations
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))

	// Print the value of Pi (exported name from the math package requires capitalization)
	fmt.Println(math.Pi)

	//Functions can be acessed like this
	fmt.Println(add(5, 15))

	//a word swap function
	a, b := swap("hello", "world") //DECLARE A VARIABLE AND INITIALIZE IN SAME LINE WITH :=
	fmt.Println(a, b)

	//naked returns
	fmt.Println(split(17))

	//testing vars and void functions
	testingVars()
	testingTypes()
	testingConstants()

    //for loops
    forLoops()
}

//END MAIN __________________________________________________________________________________________________________________________________________________________________________________________________________________

//Functions Syntax
/*
func functionName(param1, param2) (return type)
*/

// ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
// Functions are declared with parameters like so (note the var names before type):
// ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

// ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
// Naked Returns: Only to be used in short functions because of readability issues, they returned the NAMED variables in the declaration without having specified what to return.
// ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

// ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
// var: Var can be used at package or function level and declares a list of variables
// ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
var someBool, anotherBool, oneMoreBool bool
var p, q int = 1, 2 // initialize and declare in package level

func testingVars() {
    fmt.Println("-----------------------------testingVars--------------------------------------")

	//single line initialization and declaration
	i := 10
	aTruth, noTruth, str := true, false, "no!"

	var j string                   //declaration
	j = "This is how to use vars!" //initialization

	fmt.Println(p, q, i, j, someBool, anotherBool, oneMoreBool, aTruth, noTruth, str)
}

/*
----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
Go Types:

		bool

		string

		int  int8  int16  int32  int64
		uint uint8 uint16 uint32 uint64 uintptr

		byte // alias for uint8

		rune // alias for int32
		    // represents a Unicode code point

		float32 float64

		complex64 complex128

		Variables declared without an explicit initial value are given their zero value.
	        The zero value is:

	        0 for numeric types,
	        false for the boolean type, and
	        "" (the empty string) for strings.

----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
*/
var (
	ToBe   bool       = false                //boolean
	MaxInt uint64     = 1<<64 - 1            //A 64-bit unsigned integer initialized to the maximum value possible for its type (2^64 - 1).
	z      complex128 = cmplx.Sqrt(-5 + 12i) // A complex number of type complex128 initialized to the square root of the complex number -5 + 12i using the cmplx. Sqrt function from the math/cmplx package.
)

func testingTypes() {
	// %T prints the type of the variable
	// %v prints the value of the variable in a default format
	// Other common format verbs include:
	// %d for integers
	// %s for strings
	// %f for floats
	// %t for booleans
    fmt.Println("------------------------testingTypes-------------------------------")

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	//type conversions
	//1. It declares two integers `x` and `y`
	//2. Uses `float64()` to convert the integer expression `x*x + y*y` to a float for use with `math.Sqrt()`
	//3. Converts the resulting float back to an unsigned integer with `uint()`
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	//Type inference (You don't always need to specify type)
	v := 42.6
	j := uint(v)
	fmt.Printf("v is of type %T\n", v)
	fmt.Printf("j is of type %T\n", j)
}

// ----------------------------------------------------------------------------------
// Constants: Constants are declared like variables, but with the const keyword.
//Constants can be character, string, boolean, or numeric values.
//Constants cannot be declared using the := syntax.
// ----------------------------------------------------------------------------------

func testingConstants() {
    fmt.Println("-------------------------testingConstants------------------------")
	const word1 = "lets"
	const word2 = "go"
	const word3 = "yurr"
	combined := word1 + " " + word2 + " " + word3

	fmt.Println(combined)
}

// ----------------------------------------------------------------------------------
// For loops: Syntax
// ----------------------------------------------------------------------------------

func forLoops() {
    fmt.Println("-------------------------------For Loops------------------------------------")
    //traditional forloop
	sum := 0
	for i := 0; i < 10; i++ {
        fmt.Println(sum)
		sum += i
	}
	fmt.Println(sum)

    //The init and post statements are optional.
    j := 1
    for ; j < 100; {
        j += j
    }
    fmt.Println(j)

    fmt.Println("-------------------------------IF ELSE------------------------------------")

    //If else statement with prompting for input
    var x int
    fmt.Print("Enter a number ")
    fmt.Scan(&x) //promting for input
    if x > 0 {
        fmt.Println("Greater than 0")
	} else{
        fmt.Println("Less than 0")
    }


}
