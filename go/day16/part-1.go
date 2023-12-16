package day16

import (
	_ "embed"
	"fmt"
)

func Part1(dataB []byte) (sum int64) {
	contraption := loadData(dataB)
	contraption.lightbeam(0, 0, east)
	//contraption.Print()
	return contraption.energized()
}

func (c *contraptionT) lightbeam(x, y int, dir directionT) {
	if x < 0 || y < 0 || x >= len(c.board[0]) || y >= len(c.board) {
		return
	}
	c.energy[y][x] = '#'
	//c.Print()
	currentField := (c.board[y][x])
	switch currentField {
	case '.':
		c.lightMarkMoveOn(x, y, dir, dir.mark())
	case '/', '\\', '-', '|':
		newDir := dir.nextDir(currentField)
		switch newDir {
		case north, east, south, west:
			c.lightMoveOn(x, y, newDir)
		case northSouth:
			c.lightMoveOn(x, y, north)
			c.lightMoveOn(x, y, south)
		case eastWest:
			c.lightMoveOn(x, y, east)
			c.lightMoveOn(x, y, west)
		default:
			panic("oops")
		}
	case '^', '>', 'v', '<':
		if dir.mark() == currentField {
			return
		}
		c.lightMarkMoveOn(x, y, dir, '2')
	case '2', '3', '4', '5', '6', '7', '8':
		c.lightMarkMoveOn(x, y, dir, currentField+1)
	case '9':
		c.lightMoveOn(x, y, dir)
	default:
		msg := fmt.Sprintf("unhandled field %q == %#v", string(currentField), currentField)
		panic(msg)
	}
}

func (c *contraptionT) lightMoveOn(x, y int, dir directionT) {
	x1, y1 := dir.nextCoords(x, y)
	c.lightbeam(x1, y1, dir)
}

func (c *contraptionT) lightMarkMoveOn(x, y int, dir directionT, newField byte) {
	c.board[y][x] = newField
	c.lightMoveOn(x, y, dir)
}

func (c *contraptionT) energized() (sum int64) {
	for _, row := range c.energy {
		for _, field := range row {
			if field == '#' {
				sum++
			}
		}
	}
	return sum
}
