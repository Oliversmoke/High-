package piscine

func Compact(ptr *[]string) int {
	s := *ptr
	newSlice := []string{}

	for _, v := range s {
		if v != "" {
			newSlice = append(newSlice, v)
		}
	}

	*ptr = newSlice
	return len(newSlice)
}
