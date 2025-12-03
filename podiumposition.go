package piscine

func PodiumPosition(podium [][]string) [][]string {
	n := len(podium)
	result := make([][]string, n)

	for i := 0; i < n; i++ {
		result[i] = podium[n-1-i]
	}

	return result
}
