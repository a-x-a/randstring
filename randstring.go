package randstring

import (
	crand "crypto/rand"
	"math/big"
	mrand "math/rand"
	"strings"
)

const (
	// minimum length
	minLength = 1
)

var (
	defaultCharSet = []string{
		"abcdefghijklmnopqrstuvwxyz",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		"0123456789",
		"~!@#$%^&*()-_=+[{]};:\\|/"}
)

func Generate(length int, charSet ...string) (string, error) {
	rands := ""

	if len(charSet) == 0 || len(charSet[0]) == 0 {
		charSet = defaultCharSet
	}
	if length < minLength {
		length = minLength
	}

	for _, s := range charSet {
		if length == 0 {
			break
		}
		r, err := crand.Int(crand.Reader, big.NewInt(int64(len(s))))
		if err != nil {
			return "", err
		}
		rands += string(s[r.Int64()])
		length--
	}

	s := strings.Join(charSet, "")

	for length > 0 {
		r, err := crand.Int(crand.Reader, big.NewInt(int64(len(s))))
		if err != nil {
			return "", err
		}
		rands += string(s[r.Int64()])
		length--
	}

	p := []byte(rands)
	mrand.Shuffle(len(p), func(i, j int) {
		p[i], p[j] = p[j], p[i]
	})
	rands = string(p)

	return rands, nil
}
