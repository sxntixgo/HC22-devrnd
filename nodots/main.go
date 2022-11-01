package main

import (
	"fmt"
	"time"
)

func main() {

	sub := map[int]string{
		1: "c", 2: "d", 3: "e", 4: "f", 5: "g", 6: "h", 7: "i", 8: "j", 9: "k", 10: "l", 11: "m", 12: "n", 13: "o", 14: "p", 15: "q", 16: "r", 17: "s", 18: "t", 19: "u", 20: "v", 21: "w", 22: "x", 23: "y", 24: "z", 25: "a", 26: "b", 27: "{", 28: "}",
	}

	c_flag := [10]int{4,10,25,5,27,18,7,11,3,28}
	var flag = ""
	var size = 0

	for i := 0; i < len(c_flag); i++ {
		flag = flag + sub[c_flag[i]]
	}

	var input = ""
	fmt.Println("Type the flag and I will tell you if it is right")
	fmt.Scanln(&input)

	if len(input) < len(flag) {
		size = len(input)
	} else {
		size = len(flag)
	}

	for i := 0; i < size; i++ {
		first_i := input[0]
		first_p := flag[0]
		if first_i == first_p {
			fmt.Printf("")
			time.Sleep(500 * time.Millisecond)
		} else {
			fmt.Printf("Sorry, wrong flag\n")
			return
		}
		time.Sleep(100 * time.Millisecond)
		input = input[1:]
		flag = flag[1:]
	}
	fmt.Println("")
	if len(input) > 0 || len(flag) > 0 {
		fmt.Printf("Sorry, wrong flag")
	} else {
		fmt.Printf("You got the right flag\n")
	}
}
