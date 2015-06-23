package gopuz

import ()

type Puzzle struct {
	title     string
	author    string
	copyright string
	width     int
	height    int
	numClues  int
	clues     []string
	solution  [][]byte
	state     [][]byte
}
