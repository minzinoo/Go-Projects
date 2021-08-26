package main

import (
	"myapp/packageone"
)

var myVar = "I am MyVar in package level of Main."

func main() {

	var blockVar = "I am blockVar from Block Level in Main."

	packageone.PrintMe(myVar, blockVar)

	// variables
	// declare a package level variable for the main
	// package named myVar

	// declare a block variable for the main function
	// called blockVar

	// declare a package level variable in the packageone
	// package named PackageVar

	// create an exported function in packageone called PrintMe

	// in the main function, print out the values of myVar,
	// blockVar, and PackageVar on on line using the PrintMe
	// function in packageone

}

/*Fourth version of Scope

package main

import (
	"fmt"
	"myapp/packageone"
)

var one = "One"

func main() {
	var somethingElse = "this is a block level variable"
	fmt.Println(somethingElse)
	myFunc()

	newString := packageone.PublicVar
	fmt.Println("From packageone:", newString)

	packageone.Exported()

}

func myFunc() {
	fmt.Println(one)
}

// *privateVar in package.go is declared with small letter on "private" word which means it can be used internally.
// *PublicVar is with capital letter can be used outside too.
// *The above also applies same thing to function as well. Look at "notExported" and "Exported" functions.

*/

/*Third version of Scope
// declare variable one in package level area.
package main

import "fmt"

var one = "One"

func main() {

	fmt.Println(one)

	myFunc()

}

func myFunc() {
	fmt.Println(one)
}
//Output is
//One
//One
*/

/*Second version of Scope
package main

import "fmt"

func main() {

	var one = "One"
	var two = "Two"
	fmt.Println(one)

	myFunc(two)

}

func myFunc(one string) {
	fmt.Println(one)
}
// Output is
// One
// Two
*/

/* First version of Scope
package main

import "fmt"

func main() {

	var one = "One"
	fmt.Println(one)

	myFunc()

}

func myFunc() {
	var one = "the number one"
	fmt.Println(one)
}
//Output is
//One
//the number one

*/

// command to create "go.mod" is "go mod init mapp"
