package day03

import (
	_ "embed"
)

func Part1(dataB []byte) (sum int64) {
	engine := loadData(dataB)
	engine.parsePartNumbers()
	return engine.sumPartNumbers()
}

func (e *engineT) parsePartNumbers() {
	for y, line := range e.lines {
		num := numberT{
			x: -1,
			y: y,
		}
		for x, ch := range line {
			switch {
			case ch >= '0' && ch <= '9':
				num.addChar(x, ch)
			case ch == '.':
				e.addNumber(&num)
			default:
				e.addNumber(&num)
				e.symbols.add(x, y, ch)
			}
		}
		e.addNumber(&num)
	}
}

func (e *engineT) sumPartNumbers() (sum int64) {
numberLoop:
	for _, num := range e.numbers {
		for y := num.y - 1; y <= num.y+1; y++ {
			for x := num.x - 1; x <= num.x+num.l; x++ {
				symCh := e.symbols.at(x, y)
				if symCh != "" {
					numVal := num.val
					sum += numVal
					continue numberLoop
				}
			}
		}
	}
	return sum
}
