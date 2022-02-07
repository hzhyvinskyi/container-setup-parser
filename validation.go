package main

// findClosing finds appropriate closing bracket for a given one
func findClosing(openBracket string) string {
	switch openBracket {
	case "(":
		return ")"
	case "[":
		return "]"
	case "{":
		return "}"
	default:
		return ""
	}
}

// validate validates parenthesis in given input recursively
func validate(input string, n int) bool {
	if n == 0 {
		return true
	}

	// One character can't be valid
	if n == 1 {
		return false
	}

	char := string(input[0])

	// If first character is closing bracket then given input isn't valid
	if inSlice(char, []string{")","]","}"}) {
		return false
	}

	// Need to find closing bracket for the first opening bracket
	closing := findClosing(char)

	// Count is used for cases where the same brackets go one by one
	i, count := 0, 0
	for i = 1; i < n; i++ {
		char := string(input[i])
		// Exclude character from the input if it's not a bracket
		if !inSlice(char, []string{"(", ")", "[", "]", "{", "}"}) {
			input = input[:i] + input[i+1:]
			n--
			continue
		}
		if input[i] == input[0] {
			count++
		}
		if char == closing {
			if count == 0 {
				break
			}
			count--
		}
	}

	// Check if closing bracket doesn't exist
	if i == n {
		return false
	}

	// Check if closing bracket is next to open bracket
	if i == 1 {
		return validate(input[2:], n-2)
	}

	// Closing bracket must be a little further
	return validate(input[1:], i-1) && validate(input[i+1:], n-i-1)
}
