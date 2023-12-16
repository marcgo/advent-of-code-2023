package day03

import (
	_ "embed"
)

func Part2(dataB []byte) (sum int64) {
	engine := loadData(dataB)
	engine.parsePartNumbers()
	return engine.sumGearRatios()
}

func (e *engineT) sumGearRatios() (sum int64) {
numberLoop:
	for _, num := range e.numbers {
		for y := num.y - 1; y <= num.y+1; y++ {
			for x := num.x - 1; x <= num.x+num.l; x++ {
				symCh := e.symbols.at(x, y)
				if symCh == "*" {
					e.symbols.addGear(x, y, num)
					continue numberLoop
				}
			}
		}
	}
	return e.symbols.sumGearRatios()
}

func (syms *symbolsT) sumGearRatios() (sum int64) {
	for _, nums := range syms.gears {
		if len(nums) == 2 {
			sum += nums[0].val * nums[1].val
		}
	}
	return sum
}
