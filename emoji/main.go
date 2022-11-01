package main

import (
	"fmt"
	"math"
	"syscall"
	"unicode"
	"github.com/forPelevin/gomoji"
	"golang.org/x/term"
)

func isASCII(s string) bool {
    for i := 0; i < len(s); i++ {
        if s[i] > unicode.MaxASCII {
            return false
        }
    }
    return true
}

func main() {

	sub := map[int]string{
		1: "c", 2: "d", 3: "e", 4: "f", 5: "g", 6: "h", 7: "i", 8: "j", 9: "k", 10: "l", 11: "m", 12: "n", 13: "o", 14: "p", 15: "q", 16: "r", 17: "s", 18: "t", 19: "u", 20: "v", 21: "w", 22: "x", 23: "y", 24: "z", 25: "a", 26: "b", 27: "{", 28: "}", 29: "🧮️",
	}

	c_flag := [7]int{4,10,25,5,27,29,28}

	var flag = ""

	for i := 0; i < len(c_flag); i++ {
		flag = flag + sub[c_flag[i]]
	}

	var b = 3664.00
	var t = math.Log2(math.Pow(b, float64(42)))

	fmt.Printf("Type a password that only uses different emojis and whose entropy is %f.\n", t)
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
    if err != nil {
        fmt.Println(err)
		return
    }

    input := string(bytePassword)

	if isASCII(input) {
		fmt.Println("This password includes ASCII characters.")
		return
	}

	r := len(gomoji.FindAll(input))

	if math.Log2(math.Pow(b, float64(r))) == t {
		fmt.Printf("%s", flag)	
	} else {
		fmt.Printf("You used %d different emojis.\n", r)
		fmt.Printf("That's not a password whose entropy is %f.", t)
	}
}
