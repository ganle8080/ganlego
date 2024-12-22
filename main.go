package main

import "fmt"

func main() {

	friends := []string{"Tom", "jerry"}
	for _, friend := range friends {
		fmt.Printf("friend: %v\n", friend)
	}
}
