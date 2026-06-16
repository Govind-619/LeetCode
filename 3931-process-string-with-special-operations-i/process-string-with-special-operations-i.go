package main

import (
	"strings"
)

func processStr(s string) string {
	// chunks stores the valid string segments processed so far
	chunks := make([]string, 0, len(s))

	for i := 0; i < len(s); i++ {
		ch := s[i]

		switch ch {
		case '*': // Backspace operation
			// Find the last non-empty chunk and remove its last character
			for len(chunks) > 0 {
				lastIdx := len(chunks) - 1
				if len(chunks[lastIdx]) > 0 {
					chunks[lastIdx] = chunks[lastIdx][:len(chunks[lastIdx])-1]
					// If the chunk becomes empty, pop it entirely
					if len(chunks[lastIdx]) == 0 {
						chunks = chunks[:lastIdx]
					}
					break
				} else {
					chunks = chunks[:lastIdx]
				}
			}

		case '%': // Reverse operation
			if len(chunks) == 0 {
				continue
			}
			// 1. Reverse the order of the chunks themselves
			for left, right := 0, len(chunks)-1; left < right; left, right = left+1, right-1 {
				chunks[left], chunks[right] = chunks[right], chunks[left]
			}
			// 2. Reverse the internal characters of each individual chunk
			for j := 0; j < len(chunks); j++ {
				runes := []rune(chunks[j])
				for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
					runes[left], runes[right] = runes[right], runes[left]
				}
				chunks[j] = string(runes)
			}

		case '#': // Duplicate operation
			if len(chunks) == 0 {
				continue
			}
			// Safely duplicate the current state by appending the existing chunks to themselves
			chunks = append(chunks, chunks...)

		default: // Normal lowercase letters
			chunks = append(chunks, string(ch))
		}
	}

	// Join all processed chunks into the final string
	return strings.Join(chunks, "")
}