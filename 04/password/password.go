package password

func IsValidPassword(password int) bool {
	return hasAdjacentRepeatingDigits(password) && hasNonDecreasingDigits(password)
}

func IsValidPasswordPartTwo(password int) bool {
	return IsValidPassword(password) && hasExactlyTwoAdjacentRepeatingDigits(password)
}

func hasAdjacentRepeatingDigits(password int) bool {
	lastDigit := password % 10

	for password > 0 {
		password = password / 10
		currentDigit := password % 10

		if lastDigit == currentDigit {
			return true
		}

		lastDigit = currentDigit
	}

	return false
}

func hasExactlyTwoAdjacentRepeatingDigits(password int) bool {
	count := 1
	lastDigit := password % 10

	for password > 0 {
		password = password / 10
		currentDigit := password % 10

		if lastDigit == currentDigit {
			count++
		} else {
			if count == 2 {
				return true
			}

			count = 1
		}

		lastDigit = currentDigit
	}

	return false
}

func hasNonDecreasingDigits(password int) bool {
	lastDigit := password % 10

	for password > 0 {
		password = password / 10
		currentDigit := password % 10

		if currentDigit > lastDigit {
			return false
		}

		lastDigit = currentDigit
	}

	return true
}
