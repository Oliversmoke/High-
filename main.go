package high

import (
	"fmt"
	"os"
)

// Print "Error" and exit.
func fail() {
	fmt.Println("Error")
	os.Exit(0)
}

// Check input: 9 lines, each 9 chars ('.' or '1'-'9')
func validInput(lines []string) bool {
	if len(lines) != 9 {
		return false
	}

	for _, line := range lines {
		if len(line) != 9 {
			return false
		}
		for _, ch := range line {
			if ch != '.' && (ch < '1' || ch > '9') {
				return false
			}
		}
	}

	return true
}

// Convert input into a 2D slice grid
func buildGrid(lines []string) [][]byte {
	grid := make([][]byte, 9)
	for r := 0; r < 9; r++ {
		grid[r] = make([]byte, 9)
		for c := 0; c < 9; c++ {
			grid[r][c] = lines[r][c]
		}
	}
	return grid
}

// Validate the starting grid for duplicates
func initialValid(grid [][]byte) bool {
	rowSeen := make([][]bool, 9)
	colSeen := make([][]bool, 9)
	boxSeen := make([][]bool, 9)

	for i := 0; i < 9; i++ {
		rowSeen[i] = make([]bool, 10)
		colSeen[i] = make([]bool, 10)
		boxSeen[i] = make([]bool, 10)
	}

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {

			ch := grid[r][c]
			if ch == '.' {
				continue
			}

			d := int(ch - '0')

			// row or column repeated
			if rowSeen[r][d] || colSeen[c][d] {
				return false
			}
			rowSeen[r][d] = true
			colSeen[c][d] = true

			// box repeated
			box := (r/3)*3 + (c / 3)
			if boxSeen[box][d] {
				return false
			}
			boxSeen[box][d] = true
		}
	}

	return true
}

// Check if placing digit d at (r, c) is allowed
func canPlace(grid [][]byte, r, c int, d byte) bool {

	// Row + column check
	for i := 0; i < 9; i++ {
		if grid[r][i] == d || grid[i][c] == d {
			return false
		}
	}

	// 3×3 box check
	br := (r / 3) * 3
	bc := (c / 3) * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if grid[br+i][bc+j] == d {
				return false
			}
		}
	}

	return true
}

// Backtracking Sudoku solver
func solve(grid [][]byte) bool {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {

			if grid[r][c] == '.' {

				// Try digits 1 to 9
				for d := byte('1'); d <= '9'; d++ {
					if canPlace(grid, r, c, d) {
						grid[r][c] = d

						if solve(grid) {
							return true
						}

						grid[r][c] = '.' // undo
					}
				}

				return false // nothing fits here
			}

		}
	}
	return true // solved
}

// Print final solved grid
func printGrid(grid [][]byte) {
	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if c > 0 {
				fmt.Print(" ")
			}
			fmt.Printf("%c", grid[r][c])
		}
		fmt.Println()
	}
}

func main() {
	input := os.Args[1:]

	if !validInput(input) {
		fail()
	}

	grid := buildGrid(input)

	if !initialValid(grid) {
		fail()
	}

	if !solve(grid) {
		fail()
	}

	printGrid(grid)
}
