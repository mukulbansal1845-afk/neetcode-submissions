func isPalindrome(s string) bool {

    i := 0
	j := len(s) - 1

	for i < j {
		// Skip non-alphanumeric from the left.
		for i < j && !isAlphaNumeric(rune(s[i])) {
			i++
		}

		// Skip non-alphanumeric from the right.
		for i < j && !isAlphaNumeric(rune(s[j])) {
			j--
		}

		if unicode.ToLower(rune(s[i])) != unicode.ToLower(rune(s[j])) {
			return false
		}

		i++
		j--
	}

	return true

}

func isAlphaNumeric(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}
