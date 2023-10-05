package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("Ping Pong")
		} else if i%3 == 0 {
			fmt.Println("Ping")
		} else if i%5 == 0 {
			fmt.Println("Pong")
		} else {
			fmt.Println(i)

		}

	}
}
