package main

import (
	"fmt"
)

func longestCommonPrefix(strs []string) []rune {
	if len(strs) == 0 {
		
		return []rune{}
	}

	prefix := []rune{}

search:
	for i := 0; i < len(strs[0]); i++ {
		V := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != V {
				
				break search
			}
		}
		prefix = append(prefix, rune(V))
	}

	return prefix
}

func main() {
	A := []string{"butterfly", "butter", "buttress"}
	fmt.Println(string(longestCommonPrefix(A)))
}
