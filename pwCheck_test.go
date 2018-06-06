package pwcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckForPwnage(t *testing.T) {
	// Check Password "password" for pwnage, should return no error and pwned true
	pass := "password"
	pwd, err := CheckForPwnage(pass)
	assert.Nil(t, err)
	assert.True(t, pwd.Pwned)
	assert.Equal(t, pass, pwd.Pass)
	assert.NotZero(t, pwd.TimesPwned)

	// Check Password "", should return Passphrase Empty error
	pwd2, err := CheckForPwnage("")
	assert.Error(t, err, ErrPassphraseEmpty.Error())
	assert.EqualValues(t, false, pwd2.Pwned)
	assert.EqualValues(t, 0, pwd2.TimesPwned)
}

func TestIsPwned(t *testing.T) {
	err := IsPwned("password")
	assert.EqualError(t, err, "Password is pwned")
}

func TestCheckPass(t *testing.T) {
	pass := "password"
	res, err := CheckPass(pass)
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	assert.Equal(t, pass, res.Pass)
	assert.Zero(t, res.Score)
	assert.Zero(t, res.CrackTimeSeconds)
}
