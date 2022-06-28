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
	defer file.Close()

	var line int
	var passLine string
	var passVal string

	for txtScanner.Scan() {
		if line == 0 {
			passLine = txtScanner.Text()
			break
		}
	}

	if len(passLine) > 9 {
		passVal = passLine[10:]
	} else {
		passVal = ""
	}

	return passVal

}

// used to update txt file

func writePass(filePath string, newPassword string) {

	file, err := os.ReadFile("password.txt")

	if err != nil {
		log.Fatal(err)
	}

	rows := strings.Split(string(file), "\n")

	for i, row := range rows {
		if strings.Contains(row, "password: ") {
			updatedPassLine := fmt.Sprintf("password: %s", newPassword)
			rows[i] = updatedPassLine
		}
	}

	var output = strings.Join(rows, "\n")
	err = os.WriteFile("password.txt", []byte(output), 0644)
	if err != nil {
		log.Fatal(err)
	}

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

	var currentPassPath = readPass(`password.txt`)

	var tries int
	var loginVerified bool

	for i := 0; i < 3; i++ {
		var myInput = readConsoleInput()

		if myInput == currentPassPath {
			fmt.Println("Login succcessful")
			loginVerified = true
			break
		} else {
			tries = 2 - i
			toDisplay := fmt.Sprintf("Incorrect login, you have %d tries left", tries)
			fmt.Println(toDisplay)
		}
	}

	if loginVerified {
		fmt.Println("Would you like to change your password ? Y/N")
		var answer = readConsoleInput()
		if answer == "Y" {
			fmt.Println("Please provide new password")
			var newPasswordValue = readConsoleInput()

			writePass(currentPassPath, newPasswordValue)
			fmt.Println("Password updated ! \n Have a nice day !")
		}
	}

}
