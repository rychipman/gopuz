package gopuz

import (
	"bufio"
	"os"
)

type Puzzle struct {
	title     string
	author    string
	copyright string
	version   string
	notes     string
	width     int
	height    int
	numClues  int
	clues     []string
	solution  [][]byte
	state     [][]byte
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (puz *Puzzle) New(fileName string) {
	f, err := os.Open(fileName)
	check(err)

	r := bufio.NewReader(f)

	header := make([]byte, 0x34)
	_, err := r.Read(header)
	check(err)

	checksum := int(header[0x0:0x2])
	magicString := string(header[0x2:0xE])
	cibChecksum := int(header[0xE:0x10])
	lowChecksums := header[0x10:0x14]
	highChecksums := header[0x14:0x18]
	puz.version = string(header[0x18:0x1C])
	reserved1C := header[0x1C:0x1E]
	scrambledChecksum := int(header[0x1E:0x20])
	reserved20 := header[0x20:0x2C]
	puz.width = int(header[0x2c])
	puz.height = int(header[0x2d])

	var tmp int16
	buf := bytes.NewReader(header[0x2e:0x30])
	err = binary.Read(buf, binary.LittleEndian, &tmp)
	check(err)
	numClues := int(tmp)

	unknownBitmask := int(header[0x30:0x32])
	scrambledTag := int(header[0x30:0x32])

	puz.solution = make([][]byte, puz.height)
	for i := 0; i < puz.height; i++ {
		row := make([]byte, puz.width)
		_, err := r.Read(row)
		check(err)
		append(puz.solution, row)
	}

	puz.state = make([][]byte, puz.height)
	for i := 0; i < puz.height; i++ {
		row := make([]byte, puz.width)
		_, err := r.Read(row)
		check(err)
		append(puz.state, row)
	}

	puz.title, err = r.ReadString('\x00')
	check(err)
	puz.author, err = r.ReadString('\x00')
	check(err)
	puz.copyright, err = r.ReadString('\x00')
	check(err)

	for i := 1; i <= numClues; i++ {
		clue, err := r.ReadString('\x00')
		check(err)
		fmt.Printf("%d - %s\n", i, clue)
	}

	puz.notes, err = r.ReadString('\x00')
	check(err)
}
