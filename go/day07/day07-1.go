package day07

import (
	_ "embed"
	"sort"
)

//go:embed test-1.data
var Test_1 []byte
var TestResult_1 = int64(6440)

//go:embed input-1.data
var Input_1 []byte

func Day07_1(dataB []byte) (sum int64) {
	cards := loadCards(dataB)
	sort.Slice(cards, func(i, j int) bool {
		ci := cards[i].hand
		cj := cards[j].hand
		ti := ci.getType()
		tj := cj.getType()
		if ti != tj {
			return ti > tj
		}
		for k := 0; k < 5; k++ {
			si, oki := strengthOfCard[ci[k]]
			sj, okj := strengthOfCard[cj[k]]
			if !oki || !okj {
				panic("oops")
			}
			if si == sj {
				continue
			}
			return si < sj
		}
		panic("oops")
	})
	for rank, card := range cards {
		sum += (int64(rank) + 1) * card.bid
	}
	return sum
}
