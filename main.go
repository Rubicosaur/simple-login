package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// read password from txt file

func readPass(filePath string) string {

	file, err := os.Open(filePath)

	if err != nil {
		log.Fatal(err)
	}

	txtScanner := bufio.NewScanner(file)

	var line int
	var passLine string
	for txtScanner.Scan() {
		if line == 0 {
			passLine = txtScanner.Text()
			break
		}
	}

	if len(passLine) > 9 {
		var passVal = passLine[10:]
		return passVal

	} else {

		return ""

	}

}

// used to update txt file
func writePass(filePath string) {

}

// read user input and remove unnecessary, unprintable signs

func readConsoleInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\r\n", "", -1)
	input = strings.Replace(input, "\n", "", -1)
	return input
}

func main() {

	fmt.Println("Provide password and press ENTER...")

	var currentPass = readPass(`C:\Users\Omen\Desktop\Go\login\password.txt`)

	var tries int

	for i := 0; i < 3; i++ {
		var myInput = readConsoleInput()

		if myInput == currentPass {
			fmt.Println("Login succcessful")
			break
		} else {
			tries = 2 - i
			toDisplay := fmt.Sprintf("Incorrect login, you have %d tries left", tries)
			fmt.Println(toDisplay)
		}
	}
}
