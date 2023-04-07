package main

import "fmt"

func main() {
	fmt.Println("Enter you text:")
	var text string
	fmt.Scan(&text)
	count := 0
	kount := 0
	for i := 0; i < 4; i++ {
		if text[i] >= 65 && text[i] <= 90 {
			count++
		}
	}
	if count == 4 {
		for i := 0; i < 4; i++ {
			for j := i + 1; j < 4; j++ {
				if text[i] == text[j] {
					kount++
				}
			}
		}
		if kount == 2 {
			fmt.Println("The text consists of two letters, each of which occurs twice")
		} else {
			fmt.Println("NO")
		}
	} else {
		fmt.Println("NO")
	}
}
