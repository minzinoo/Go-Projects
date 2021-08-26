package packageone

import "fmt"

var PackageVar = "I am PackageVar in package level of packageone."

func PrintMe(s1, s2 string) {

	fmt.Println(s1, s2, PackageVar)
}

/* Used with Fourth version of main.go
package packageone

var privateVar = "I am private"
var PublicVar = "I am public (or exported)"

func notExported() {

}

func Exported() {

}

// *privateVar in package.go is declared with small letter on "private" word which means it can be used internally.
// *PublicVar is with capital letter can be used outside too.
// *The above also applies same thing to function as well. Look at "notExported" and "Exported" functions.
*/
