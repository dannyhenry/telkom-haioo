package main

import (
	"fmt"
	"math"
)

func compareTwoString(s1, s2 string) bool {
	if math.Abs(float64(len(s1)-len(s2))) > 1 {
		return false
	}

	var i, j, edits int
	for i < len(s1) && j < len(s2) {
		if s1[i] != s2[j] {
			edits++
			if len(s1) > len(s2) {
				i++
			} else if len(s1) < len(s2) {
				j++
			} else {
				i++
				j++
			}
		} else {
			i++
			j++
		}

		if edits > 1 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(compareTwoString("telkom", "telecom")) // Output: false
	fmt.Println(compareTwoString("telkom", "tlkom"))   // Output: true
}
