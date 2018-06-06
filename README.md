# pwcheck
[![GoDoc](https://godoc.org/github.com/dwin/pwCheck?status.svg)](https://godoc.org/github.com/dwin/pwCheck)
[![cover.run](https://cover.run/go/github.com/dwin/pwCheck.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=github.com%2Fdwin%2FpwCheck)
[![Build Status](https://travis-ci.org/dwin/pwCheck.svg?branch=master)](https://travis-ci.org/dwin/pwCheck)
[![Go Report Card](https://goreportcard.com/badge/github.com/dwin/pwCheck)](https://goreportcard.com/report/github.com/dwin/pwCheck)

pwCheck is a utility package that gives password strength and verifies passphrase 
has not been compromised in a previous breach using the https://haveibeenpwned.com API.

---

## Get Started
```
go get github.com/dwin/pwCheck
```

The function checks the password using the https://haveibeenpwned.com API.

It returns a pointer to a struct like so:

```go
// Pwd is returned as a struct pointer when calling CheckForPwnage
type Pwd struct {
	Pwned      bool   // Pwned returns true if passphrase is found pwned via API
	Pass       string // Pass returns the passphrase string passed to the function
	TimesPwned int    // TimesPwned returns the number of times the passphrase was found in the database
}
```

`Pwned` is true if the password has been pwned. <br>
`Pass` is the original password passed to the fucntion. <br>
`TimesPwned` is an int with the number of times the password has been pwned.

**Example Usage:**
```go
func Example() {
	userPass := form.Data("password")

	pwc, err := pwCheck.CheckForPwnage(userPass)
	if err != nil {
		fmt.Println("Password Check error: %s", err)
	}
	if pwc.Pwnd {
		fmt.Println("Password is found in database")
	}
}
```

# ToDo
- [ ] HTTP Client Timeout

Inspired by: [https://github.com/masonj88/pwchecker](https://github.com/masonj88/pwchecker)