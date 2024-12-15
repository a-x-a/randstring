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

func Generate(length int, charSet ...string) string {
  if len(charSet) == 0 || len(charSet[0]) == 0 {
    charSet = defaultCharSet
  }

  if length < minLength {
    length = minLength
  }

  rnd := rand.NewSource(time.Now().Unix())
  var sb strings.Builder

  for _, s := range charSet {
    if length == 0 {
      break
    }

    n := len(s)
    rndn := int(rnd.Int63())
    sb.WriteByte(s[rndn%n])
    length--
  }

  s := strings.Join(charSet, "")

  n := len(s)
  for length > 0 {
    rndn := int(rnd.Int63())
    sb.WriteByte(s[rndn%n])
    length--
  }

  p := []byte(sb.String())
  rand.Shuffle(len(p), func(i, j int) {
    p[i], p[j] = p[j], p[i]
  })

  return string(p)
}
