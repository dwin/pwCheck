package pwcheck

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/nbutton23/zxcvbn-go"
)

const pwnedURL = "https://api.pwnedpasswords.com/range/%s"

var (
	// ErrPassphraseEmpty indicates passphrase input was less than 1 character
	ErrPassphraseEmpty = errors.New("Passphrase Input Empty")
)

// Pwd is returned as a struct pointer when calling CheckForPwnage()
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

// CheckForPwnage takes passphrase as string, sends request to API and returns Pwd and error
func CheckForPwnage(pw string) (pwd *Pwd, err error) {
	// Check Passphrase not empty
	if len(pw) < 1 {
		return &Pwd{false, pw, 0}, ErrPassphraseEmpty
	}

	// Create SHA1 hash of passphrase
	hash := sha1.New()
	hash.Write([]byte((string(pw))))
	// Get Passphrase prefix
	pfx := strings.ToUpper(hex.EncodeToString(hash.Sum(nil))[0:5])
	sfx := strings.ToUpper(hex.EncodeToString(hash.Sum(nil))[5:])

	// Send request to pwnedpassword API
	response, err := http.Get(fmt.Sprintf(pwnedURL, pfx))
	if err != nil {
		return &Pwd{false, pw, 0}, fmt.Errorf("HTTP request failed with error; %s", err)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return &Pwd{false, pw, 0}, fmt.Errorf("HTTP request body read failed with error; %s", err)
	}
	// Check API Response
	resp := strings.Split((string(data)), "\n")
	// Check hash prefix against suffixes returned in API response
	for i := range resp {
		// if prefix and suffix match API response, return passphrase as pwned=true
		if sfx == resp[i][0:35] {
			reg := regexp.MustCompile("[^0-9]+")
			times, err := strconv.Atoi(reg.ReplaceAllString(string(resp[i][36:]), ""))
			if err != nil {
				return pwd, err
			}
			return &Pwd{true, pw, times}, err
		}
	}

	return &Pwd{false, pw, 0}, err
}

// CheckPass
func CheckPass(pw string) (result CheckResult, err error) {
	cfp, err := CheckForPwnage(pw)
	if err != nil {
		return
	}
	score := zxcvbn.PasswordStrength(pw, nil)

	result.Pass = cfp.Pass
	result.CrackTimeSeconds = score.CrackTime
	result.CrackTimeDisplay = score.CrackTimeDisplay
	result.Pwned = cfp.Pwned
	result.Score = score.Score

	return
}

// IsPwned check passphrase input string and returns error, returns nil if password is
// not pwned and no other errors occur.
func IsPwned(pw string) error {
	pwd, err := CheckForPwnage(pw)
	if err != nil {
		return err
	}
	if pwd.Pwned {
		return fmt.Errorf("Password is pwned")
	}
	return nil
}
