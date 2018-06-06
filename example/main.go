package main

import (
	"fmt"

	"github.com/dwin/pwcheck"
)

func main() {
	passFromUser := "password"

	result, err := pwcheck.CheckPass(passFromUser)
	if err != nil {
		// Handle Error
		fmt.Println("Error: ", err)
	}

	fmt.Printf("--Result---\nPassword: %s\nPwned: %v\nScore %v of 4\nCrack Time: %.2f seconds\nHow long to crack: %s\n",
		result.Pass, result.Pwned, result.Score, result.CrackTimeSeconds, result.CrackTimeDisplay)

	// Expected Output
	// --Result---
	// Password: password
	// Pwned: true
	// Score 0 of 4
	// Crack Time: 0.00 seconds
	// How long to crack: instant

	if result.Pwned {
		// If pwned this password was found in compromised password database and you should handle
		// or inform user.
		fmt.Println("* Password is pwned and has been previously compromised")
	}

	if result.Score < 1 {
		// If score is less than 1 this is a weak password and should not be used
		fmt.Println("* Weak Password")
	}
}
