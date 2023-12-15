package day15

import (
	_ "embed"
	"strings"

	"AdventOfCode2023/internal/util"
)

type (
	boxesT [256]boxT
	boxT   struct {
		lenses []lensT
	}
	lensT struct {
		label string
		focus int
	}
)

func Part2(dataB []byte) (sum int64) {
	var boxes boxesT
	steps := loadData(dataB)
	for _, step := range steps.steps {
		switch {
		case strings.Contains(step, "-"):
			label := step[:len(step)-1]
			boxId := hash(label)
			boxes[boxId].removeLens(label)
		case strings.Contains(step, "="):
			stpL := strings.Split(step, "=")
			newLens := lensT{
				label: stpL[0],
				focus: int(util.MustAtoi(stpL[1])),
			}
			boxId := hash(newLens.label)
			boxes[boxId].insertLens(newLens)
		default:
			panic("oops")
		}
	}
	return boxes.focusingPower()
}

func (box *boxT) insertLens(newLens lensT) {
	for i, lens := range box.lenses {
		if lens.label == newLens.label {
			box.lenses[i] = newLens
			return
		}
	}
	box.lenses = append(box.lenses, newLens)
}

func (box *boxT) removeLens(oldLabel string) {
	for i, lens := range box.lenses {
		if lens.label == oldLabel {
			box.lenses = append(box.lenses[:i], box.lenses[i+1:]...)
			return
		}
	}
}

func (bxs boxesT) focusingPower() (sum int64) {
	for i, box := range bxs {
		sum += box.focusingPower(i)
	}
	return sum
}

func (box boxT) focusingPower(boxId int) (sum int64) {
	for i, lens := range box.lenses {
		sum += int64((boxId + 1) * (i + 1) * lens.focus)
	}
	return sum
}
