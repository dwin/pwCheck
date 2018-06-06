# pwcheck
[![GoDoc](https://godoc.org/github.com/dwin/pwCheck?status.svg)](https://godoc.org/github.com/dwin/pwCheck)
[![cover.run](https://cover.run/go/github.com/dwin/pwCheck.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Fdwin%2FpwCheck)
[![Build Status](https://travis-ci.org/dwin/pwCheck.svg?branch=master)](https://travis-ci.org/dwin/pwCheck)
[![Go Report Card](https://goreportcard.com/badge/github.com/dwin/pwCheck)](https://goreportcard.com/report/github.com/dwin/pwCheck)

pwCheck is a utility package that gives password strength and verifies passphrase 
has not been compromised in a previous breach using the https://haveibeenpwned.com 
API and the [Dropbox zxcvbn](https://blogs.dropbox.com/tech/2012/04/zxcvbn-realistic-password-strength-estimation/) method for estimating passphrase strength.

---

## Get Started
```
go get github.com/dwin/pwCheck
```

## Types:
```go
// Pwd is returned as a struct pointer when calling CheckForPwnage
type Pwd struct {
	Pwned      bool   // Pwned returns true if passphrase is found pwned via API
	Pass       string // Pass returns the passphrase string passed to the function
	TimesPwned int    // TimesPwned returns the number of times the passphrase was found in the database
}


// CheckResult is returned as a struct when calling CheckPass()
type CheckResult struct {
	Pwned            bool    // Pwned indicates if the pass given was found in previous breach
	Pass             string  // Pass returns the string passed to the function
	Score            int     // Score returns a 0-4 score of password strength, useful for gauge etc.
	CrackTimeSeconds float64 // CrackTimeSeconds indicates the estimated time to crack this password at ~ 10ms per guess in seconds
	CrackTimeDisplay string  // CrackTimeDisplay indicates the estimated time in seconds to years or centuries to crack password at ~ 10ms per guess
}
```
## Functions:
```go
func CheckPass(pw string) (result CheckResult, err error)
```
CheckPass() sends SHA1 partial hash of password to HaveIBeenPwned.com API 
to check for previous compromise and also computes strength using the
Dropbox "zxcvbn: realistic password strength estimation" method using 
[zxcvbn-go](https://github.com/nbutton23/zxcvbn-go). 



## Example Usage:
See [example]()
```go
func Example() {
	userPass := form.Data("password")

	checkRes, err := pwcheck.CheckPass(passFromUser)
	if err != nil {
		// Handle Error
	}

	if result.Pwned {
		// If pwned this password was found in compromised password database 
		// and you should handle or inform user.
	}

	if result.Score < 1 {
		// If score is less than 1 this is a weak password and should not be used
	}
}
```

# ToDo:
- [ ] HTTP Client Timeout

## Credits:
- Inspired by: [https://github.com/masonj88/pwchecker](https://github.com/masonj88/pwchecker)
- [zxcvbn-go](https://github.com/nbutton23/zxcvbn-go) - [by Nathan Button](https://github.com/nbutton23)
- [Have I Been Pwned API](https://haveibeenpwned.com/API/v2) - [by Troy Hunt](https://www.troyhunt.com)
- [Testify](https://github.com/stretchr/testify) - [by Stretchr, Inc.](https://github.com/stretchr)