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
		syms  map[string]string
		gears map[string][]numberT
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
		symbols: symbolsT{
			syms:  map[string]string{},
			gears: map[string][]numberT{},
		},
	}
}

func (syms *symbolsT) add(x, y int, ch rune) {
	key := fmt.Sprintf("%d,%d", x, y)
	syms.syms[key] = string(ch)
}

func (syms *symbolsT) at(x, y int) string {
	key := fmt.Sprintf("%d,%d", x, y)
	return syms.syms[key]
}

func (syms *symbolsT) addGear(x, y int, num numberT) {
	key := fmt.Sprintf("%d,%d", x, y)
	syms.gears[key] = append(syms.gears[key], num)
}

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
