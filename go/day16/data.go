package day16

import (
	_ "embed"
	"fmt"

	"AdventOfCode2023/internal/util"
)

//go:embed input.data
var Input []byte

type (
	contraptionT struct {
		board  [][]byte
		energy [][]byte
	}
)

func loadData(dataB []byte) (contraption contraptionT) {
	lines := util.LoadLines(dataB)
	for _, line := range lines {
		contraption.board = append(contraption.board, []byte(line))
		contraption.energy = append(contraption.energy, make([]byte, len(line)))
	}
	return contraption
}

func (c contraptionT) Print() {
	for _, lineB := range c.board {
		fmt.Println(string(lineB))
	}
	fmt.Println()
}

type (
	directionT int
)

const (
	none directionT = iota
	north
	east
	south
	west
	eastWest
	northSouth
)

func (dir directionT) mark() byte {
	switch dir {
	case north:
		return '^'
	case east:
		return '>'
	case south:
		return 'v'
	case west:
		return '<'
	default:
		panic("invalid direction")
	}
}

func (dir directionT) nextCoords(x, y int) (int, int) {
	switch dir {
	case north:
		return x, y - 1
	case east:
		return x + 1, y
	case south:
		return x, y + 1
	case west:
		return x - 1, y
	default:
		panic("unknown direction")
	}
}

func (dir directionT) nextDir(field byte) directionT {
	switch dir {
	case north:
		switch field {
		case '/':
			return east
		case '\\':
			return west
		case '-':
			return eastWest
		case '|':
			return north
		default:
			panic("oops")
		}

	case east:
		switch field {
		case '/':
			return north
		case '\\':
			return south
		case '-':
			return east
		case '|':
			return northSouth
		default:
			panic("oops")
		}

	case south:
		switch field {
		case '/':
			return west
		case '\\':
			return east
		case '-':
			return eastWest
		case '|':
			return south
		default:
			panic("oops")
		}

	case west:
		switch field {
		case '/':
			return south
		case '\\':
			return north
		case '-':
			return west
		case '|':
			return northSouth
		default:
			panic("oops")
		}

	default:
		panic("unknown direction")
	}
}
