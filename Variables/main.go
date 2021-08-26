package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and don't type your number in, just press ENTER when ready!"

func main() {
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2 //Intn(8) is to avoid zero and one number to start. if it is zero, it will add 2 to it.
	var secondNumber = rand.Intn(8) + 2
	var substraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - substraction

	playTheGame(firstNumber, secondNumber, substraction, answer)

}

func playTheGame(firstNumber, secondNumber, substraction, answer int) {

	//create our reader variable, which reads input from standard in (the keyboard)
	reader := bufio.NewReader(os.Stdin)

	// display a welcome/instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now substract", substraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	fmt.Println("The answer is", answer)

	//one way - declare, then assign (two steps)
	//var firstNumber int
	//firstNumber = 2

	//another way, declare type and name and assign the value
	//var secondNumber = 5

	//one step variable: declare name, assign value and let Go figure out the type
	//subtraction :=7

}

/* Third version of Number game
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and don't type your number in, just press ENTER when ready!"

func main() {
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2 //Intn(8) is to avoid zero and one number to start. if it is zero, it will add 2 to it.
	var secondNumber = rand.Intn(8) + 2
	var substraction = rand.Intn(8) + 2
	var answer = firstNumber*secondNumber - substraction

	//create our reader variable, which reads input from standard in (the keyboard)
	reader := bufio.NewReader(os.Stdin)

	// display a welcome/instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now substract", substraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	fmt.Println("The answer is", answer)

	//one way - declare, then assign (two steps)
	//var firstNumber int
	//firstNumber = 2

	//another way, declare type and name and assign the value
	//var secondNumber = 5

	//one step variable: declare name, assign value and let Go figure out the type
	//subtraction :=7

}

*/

/* Second version of Number Game
package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

const prompt = "and don't type your number in, just press ENTER when ready!"

func main() {
	//seed the random number generator
	rand.Seed(time.Now().UnixNano())

	var firstNumber = rand.Intn(8) + 2 //Intn(8) is to avoid zero and one number to start. if it is zero, it will add 2 to it.
	var secondNumber = rand.Intn(8) + 2
	var substraction = rand.Intn(8) + 2
	var answer int

	//create our reader variable, which reads input from standard in (the keyboard)
	reader := bufio.NewReader(os.Stdin)

	// display a welcome/instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now substract", substraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	answer = firstNumber*secondNumber - substraction
	fmt.Println("The answer is", answer)

	//one way - declare, then assign (two steps)
	//var firstNumber int
	//firstNumber = 2

	//another way, declare type and name and assign the value
	//var secondNumber = 5

	//one step variable: declare name, assign value and let Go figure out the type
	//subtraction :=7

}

*/

/* First version of Number Game
package main

import (
	"bufio"
	"fmt"
	"os"
)

const prompt = "and press ENTER when ready!"

func main() {

	var firstNumber = 2
	var secondNumber = 5
	var substraction = 7
	var answer int

	//create our reader variable, which reads input from standard in (the keyboard)
	reader := bufio.NewReader(os.Stdin)

	// display a welcome/instructions
	fmt.Println("Guess the Number Game")
	fmt.Println("---------------------")
	fmt.Println("")

	fmt.Println("Think of a number between 1 and 10", prompt)
	reader.ReadString('\n')

	// take them through the games
	fmt.Println("Multiply your number by", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Now multiply the result by", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you originally thought of", prompt)
	reader.ReadString('\n')

	fmt.Println("Now substract", substraction, prompt)
	reader.ReadString('\n')

	// give them the answer
	answer = firstNumber*secondNumber - substraction
	fmt.Println("The answer is", answer)

	//one way - declare, then assign (two steps)
	//var firstNumber int
	//firstNumber = 2

	//another way, declare type and name and assign the value
	//var secondNumber = 5

	//one step variable: declare name, assign value and let Go figure out the type
	//subtraction :=7

}*/
