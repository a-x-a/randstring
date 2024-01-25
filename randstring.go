package randstring

import (
	"math/rand"
	"strings"
	"time"
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

func Generate(length int, charSet ...string) (result string) {
  rnd := rand.NewSource(time.Now().Unix())

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

    rndn := int(rnd.Int63())
    result += string(s[rndn%len(s)])
    length--
  }

  s := strings.Join(charSet, "")

  for length > 0 {
    rndn := int(rnd.Int63())
    result += string(s[rndn%len(s)])
    length--
  }

  p := []byte(result)
  rand.Shuffle(len(p), func(i, j int) {
    p[i], p[j] = p[j], p[i]
  })

  result = string(p)

  return
}
