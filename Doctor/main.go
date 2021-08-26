package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {

	//var WhatToSay string. changed to := to set variable.

	reader := bufio.NewReader(os.Stdin)

	WhatToSay := doctor.Intro()

	fmt.Println(WhatToSay)

	for {

		fmt.Println("->")

		userInput, _ := reader.ReadString('\n')

		userInput = strings.Replace(userInput, "\r\n", "", -1)
		userInput = strings.Replace(userInput, "\n", "", -1)

		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}

	}

}
