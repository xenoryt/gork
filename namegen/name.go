//namegen generates random fantasy-like names
package namegen

import (
	"math/rand"
	"strings"
	"time"
)

var (
	vowels     = "aeiouy"
	consonants = "bcdfghjklmnpqrstvwxz"
	alphabet   = "abcdefghijklmnopqrstuvwxyz"
)

var vchain = map[byte][]byte{
	'a': {'u', 'i', 'y'},
	'e': {'i', 'a'},
	//'i': {},
	//'o': {},
	'u': {'e'},
}
var cchain = map[byte][]byte{
	'b': {'r', 'l'},
	'c': {'r', 'l', 'k'},
	'd': {'r'},
	'f': {'r', 'l', 't'},
	'g': {'r'},
	'h': {'n'},
	'k': {'r'},
	'l': {'d', 'n'},
	'n': {'k', 'c'},
	'p': {'h', 'l', 'r', 't'},
	'r': {'t'},
	's': {'t', 'h', 'w', 'p', 'm', 'c', 'k'},
	't': {'h', 'r'},
}

var rng *rand.Rand

func init() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func Gen() string {
	ret := ""
	nextCharType := 0
	for i := 0; i < rng.Intn(5)+5; i++ {
		lastChar := byte('0')
		char := byte('0')
		if nextCharType != 0 {
			lastChar = ret[len(ret)-1]
		}
		switch nextCharType {
		case 0:
			char = alphabet[rng.Intn(26)]
			if strings.Contains(vowels, string(char)) {
				nextCharType = 2
			} else {
				nextCharType = 1
			}
		case 1:
			if items, ok := vchain[lastChar]; ok {
				char = items[rng.Intn(len(items))]
				nextCharType = 2
			} else {
				char = vowels[rng.Intn(6)]
				if _, ok := vchain[char]; !ok || rng.Intn(10) < 8 {
					nextCharType = 2
				}
			}
		case 2:
			if items, ok := cchain[lastChar]; ok {
				char = items[rng.Intn(len(items))]
				nextCharType = 1
			} else {
				char = consonants[rng.Intn(20)]
				if _, ok := cchain[char]; !ok || rng.Intn(10) < 8 {
					nextCharType = 1
				}
			}
		}
		if lastChar != char {
			ret += string(char)
		} else {
			i--
		}
	}
	return string(ret[0]-32) + string(ret[1:])
}
