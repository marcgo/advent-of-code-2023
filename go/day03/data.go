package day03

import (
	_ "embed"
	"fmt"

	"AdventOfCode2023/internal/util"
)

//go:embed input.data
var Input []byte

type (
	engineT struct {
		lines   []string
		symbols symbolsT
		numbers []numberT
	}
	symbolsT struct {
		m map[string]string
	}
	numberT struct {
		s       string
		val     int64
		x, y, l int
	}
)

func loadData(dataB []byte) (engine engineT) {
	engine = newEngine()
	engine.lines = util.LoadLines(dataB)
	return engine
}

func newEngine() engineT {
	return engineT{
		//lines:   nil,
		symbols: symbolsT{m: map[string]string{}},
		//numbers: nil,
	}
}

func (e *engineT) sumPartNumbers() (sum int64) {
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

//func (e *engineT) findPartNumbers() (partNums []int64) {
//	//syms := symbolsT{m: map[string]bool{}}
//	for i, line := range e.lines {
//		num := numberT{
//			x: i,
//			y: -1,
//		}
//		for j, x := range line {
//			switch {
//			case x >= '0' && x <= '9':
//				num.addChar(j, x)
//			case x == '.':
//				e.addNumber(&num)
//			default:
//				e.addNumber(&num)
//				e.symbols.add(i, j)
//			}
//		}
//		e.addNumber(&num)
//	}
//
//numberLoop:
//	for _, num := range e.numbers {
//		for x := num.x - 1; x <= num.x+num.l; x++ {
//			for y := num.y - 1; y <= num.y+1; y++ {
//				if e.symbols.at(x, y) {
//					partNums = append(partNums, num.val)
//					continue numberLoop
//				}
//			}
//		}
//	}
//
//	return partNums
//}

func (syms *symbolsT) add(x, y int, ch rune) {
	key := fmt.Sprintf("%d,%d", x, y)
	syms.m[key] = string(ch)
}

func (syms *symbolsT) at(x, y int) string {
	key := fmt.Sprintf("%d,%d", x, y)
	return syms.m[key]
}

//func (syms *symbolsT) at(x, y int) bool {
//	key := fmt.Sprintf("%d,%d", x, y)
//	_, ok := syms.m[key]
//	return ok
//}

func (num *numberT) addChar(x int, ch rune) {
	if num.x < 0 {
		num.x = x
	}
	num.s += string(ch)
	num.l++
}

func (e *engineT) addNumber(num *numberT) {
	if num == nil || num.s == "" {
		return
	}
	num.val = util.MustAtoi(num.s)
	e.numbers = append(e.numbers, *num)
	num.x = -1
	num.s = ""
	num.l = 0
}
