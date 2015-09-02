package gopuz

import (
	"bufio"
	"os"
)

type Puzzle struct {
	Title     string
	Author    string
	Copyright string
	Version   string
	Notes     string
	Width     int
	Height    int
	NumClues  int
	Clues     []string
	Solution  [][]byte
	State     [][]byte
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func NewPuzzle() *Puzzle {
	return &Puzzle{
		Title:    "Untitled",
		Version:  "0.1",
		Width:    15,
		Height:   15,
		NumClues: 30,
	}
}

func (puz *Puzzle) Load(fileName string) {
	buf := newPuzzleBuffer()
	err := buf.load(fileName)
	check(err)
	puz.Title = buf.title()
	puz.Author = buf.author()
	puz.Copyright = buf.copyright()
	puz.Version = buf.version()
	puz.Notes = buf.notes()
	puz.Width = buf.width()
	puz.Height = buf.height()
	puz.Clues = buf.clues()
	puz.Solution = buf.solution()
	puz.State = buf.state()
}

type puzzleBuffer struct {
	header   []byte
	solution []byte
	state    []byte
	strings  []byte
	extra    []byte
}

func newPuzzleBuffer() *puzzleBuffer {
	return &puzzleBuffer{}
}

func (buf *puzzleBuffer) load(fileName string) error {
	f, err := os.Open(fileName)
	check(err)
	r := bufio.NewReader(f)

	// read the header
	buf.header = make([]byte, 0x34)
	_, err = r.Read(buf.header)
	check(err)

	// read the solution and state
	numSquares := buf.width() * buf.height()
	buf.solution = make([]byte, numSquares)
	_, err = r.Read(buf.solution)
	check(err)

	buf.state = make([]byte, numSquares)
	_, err = r.Read(buf.state)
	check(err)

	// TODO: read strings et al
	// TODO: verify checksums
}

func (buf *puzzleBuffer) save(fileName string) {
	// TODO: everything
}

func (buf *puzzleBuffer) checksum() int {
	var tmp int
	r := bytes.NewReader(header[:0x2])
	err = binary.Read(r, binary.LittleEndian, &tmp)
	check(err)
	return tmp
}

func (buf *puzzleBuffer) cibChecksum() int {
	var tmp int
	r := bytes.NewReader(header[0xE:0x10])
	err = binary.Read(r, binary.LittleEndian, &tmp)
	check(err)
	return tmp
}

func (buf *puzzleBuffer) lowChecksums() int {
	var tmp int
	r := bytes.NewReader(header[0x10:0x14])
	err = binary.Read(r, binary.LittleEndian, &tmp)
	check(err)
	return tmp
}

func (buf *puzzleBuffer) highChecksums() int {
	var tmp int
	r := bytes.NewReader(header[0x14:0x18])
	err = binary.Read(r, binary.LittleEndian, &tmp)
	check(err)
	return tmp
}

func (buf *puzzleBuffer) version() string {
	return string(buf.header[0x18:0x1C])
}

func (buf *puzzleBuffer) reserved1C() int {
	var tmp int
	r := bytes.NewReader(header[0x1C:0x1E])
	err = binary.Read(r, binary.LittleEndian, &tmp)
	check(err)
	return tmp
}

func (buf *puzzleBuffer) scrambledChecksum() int {
	var tmp int
	r := bytes.NewReader(header[0x1E:0x20])
	err = binary.Read(r, binary.LittleEndian, &tmp)
	check(err)
	return tmp
}

func (buf *puzzleBuffer) reserved20() int {
	var tmp int
	r := bytes.NewReader(header[0x20:0x2C])
	err = binary.Read(r, binary.LittleEndian, &tmp)
	check(err)
	return tmp
}

func (buf *puzzleBuffer) width() int {
	return int(buf.header[0x2C])
}

func (buf *puzzleBuffer) height() int {
	return int(buf.header[0x2D])
}

func (buf *puzzleBuffer) numClues() int {
	var tmp int
	r := bytes.NewReader(header[0x2e:0x30])
	err = binary.Read(r, binary.LittleEndian, &tmp)
	check(err)
	return tmp
}

func (buf *puzzleBuffer) unknownBitmask() int {
	var tmp int
	r := bytes.NewReader(header[0x30:0x32])
	err = binary.Read(r, binary.LittleEndian, &tmp)
	check(err)
	return tmp
}

func (buf *puzzleBuffer) scrambledTag() int {
	var tmp int
	r := bytes.NewReader(header[0x32:0x34])
	err = binary.Read(r, binary.LittleEndian, &tmp)
	check(err)
	return tmp
}

func (buf *puzzleBuffer) cibChecksum() int {
	base := buf.header[0x2C:0x34]
	return checksum(base, 0)
}

func checksum(base []byte, cksum int) int {
	for i := 0; i < len(base); i++ {
		if cksum & 0x0001 {
			cksum = (cksum >> 1) + 0x8000
		} else {
			cksum = cksum >> 0x0001
		}

	}
}

func (puz *Puzzle) New(fileName string) {

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
