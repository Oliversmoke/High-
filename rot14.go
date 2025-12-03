package piscine

func Rot14(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			runes[i] = rune((int(r-'a')+14)%26 + int('a'))
		} else if r >= 'A' && r <= 'Z' {
			runes[i] = rune((int(r-'A')+14)%26 + int('A'))
		}
	}
	return string(runes)
}
