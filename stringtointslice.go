package piscine

func StringToIntSlice(str string) []int {
	result := []int{}
	for _, r := range str {
		result = append(result, int(r))
	}
	return result
}
