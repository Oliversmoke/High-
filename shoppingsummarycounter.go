package piscine

import (
	"fmt"
	"strings"
)

func Main() {
	summary := "Burger Water Carrot Coffee Water Water Chips Carrot Carrot Burger Carrot Water"
	counts := ShoppingSummaryCounter(summary)

	seen := make(map[string]bool)
	for _, w := range strings.Fields(summary) {
		if !seen[w] {
			fmt.Println(w, "=>", counts[w])
			seen[w] = true
		}
	}
}

func ShoppingSummaryCounter(summary string) map[string]int {
	counts := make(map[string]int)
	for _, w := range strings.Fields(summary) {
		counts[w]++
	}
	return counts
}
