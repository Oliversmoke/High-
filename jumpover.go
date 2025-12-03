package piscine

func JumpOver(str string) string {
	n := len(str)

	if n < 3 {
		return "\n"
	}

	result := ""
	for i := 2; i < n; i += 3 {
		result += string(str[i])
	}

	return result + "\n"
}
