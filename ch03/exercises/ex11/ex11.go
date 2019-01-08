package main

import (
	"fmt"
	"reflect"
)

func isAnagram(x, y string) bool {
	freq1 := make(map[string]int)
	freq2 := make(map[string]int)

	for _, char := range x {
		freq1[string(char)]++
	}

	for _, char := range y {
		freq2[string(char)]++
	}

	fmt.Println(freq1)
	fmt.Println(freq2)
	return reflect.DeepEqual(freq1, freq2)
}

func main() {
	fmt.Println(isAnagram("dogod", "dog"))
}
