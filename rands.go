// Copyright 2016 Nemanja Zbiljic
//

// Random string generator
package rands

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	mrand "math/rand"
)

var (
	// Set of characters to use for generating random strings
	alphanumeric = []rune("0123456879ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	alphabetic   = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	numeric      = []rune("0123456879")

	maxInt64         = int64(^uint64(0) >> 1)
	max      big.Int = *big.NewInt(maxInt64)

	baseSource mrand.Source
)

func init() {
	baseSource = mrand.NewSource(cryptoRandInt64())
}

func cryptoRandInt64() int64 {
	n, err := crand.Int(crand.Reader, &max)
	if err != nil {
		panic(fmt.Errorf("crypto/rand.Int returned error: %v", err))
	}
	return n.Int64()
}

// RandomAlphanumeric creates a random string whose length is the number of
// characters specified.
//
// Characters will be chosen from the set of alpha-numeric characters.
func AlphanumericString(n int) string {
	return string(randomRunes(n, alphanumeric))
}

// RandomAlphabetic creates a random string whose length is the number of
// characters specified.
//
// Characters will be chosen from the set of alphabetic characters.
func AlphabeticString(n int) string {
	return string(randomRunes(n, alphabetic))
}

// RandomNumeric creates a random string whose length is the number of
// characters specified.
//
// Characters will be chosen from the set of numeric characters.
func NumericString(n int) string {
	return string(randomRunes(n, numeric))
}

func randomRunes(n int, charset []rune) []rune {
	if n < 0 {
		panic(fmt.Sprintf("Requested random length %d is less than 0.", n))
	}
	rand := mrand.New(baseSource)
	sid := make([]rune, n)
	rlen := len(charset)
	for i := range sid {
		// reseed rand once you've used up all available entropy bits
		if i%10 == 0 {
			rand.Seed(cryptoRandInt64()) // 64 bits of random!
		}
		sid[i] = charset[rand.Intn(rlen)]
	}
	return sid
}
