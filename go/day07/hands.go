package day07

import (
	"fmt"
	"strings"

	"AdventOfCode2023/internal/util"
)

type (
	cardT     rune
	strengthT int
	handT     [5]cardT
	handBidT  struct {
		hand handT
		bid  int64
	}
	handsT   []handBidT
	handType int
)

const (
	None handType = iota
	FiveOfaKind
	FourOfaKind
	FullHouse
	ThreeOfaKind
	TwoPair
	OnePair
	HighCard
)

var cards = []cardT{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}
var strengthOfCard = map[cardT]strengthT{}

func loadCards(dataB []byte) (hands handsT) {
	lines := util.LoadLines(dataB)
	for i, c := range cards {
		strength := len(cards) - i + 1
		strengthOfCard[c] = strengthT(strength)
	}
	for _, line := range lines {
		hbLS := strings.Split(line, " ")
		if len(hbLS) != 2 {
			panic("oops")
		}
		hand := handBidT{
			bid: util.MustAtoi(hbLS[1]),
		}
		if len(hbLS[0]) != 5 {
			panic("oops")
		}
		for i, cardR := range hbLS[0] {
			hand.hand[i] = cardT(cardR)
		}
		hands = append(hands, hand)
	}
	return hands
}

func (hd handT) getType() handType {
	sameCards := map[cardT]int{}
	for _, cd := range hd {
		sameCards[cd]++
	}
	for _, cnt := range sameCards {
		switch {
		case cnt == 5 && len(sameCards) == 1:
			return FiveOfaKind
		case len(sameCards) == 2:
			switch {
			case cnt == 4:
				return FourOfaKind
			case cnt == 3:
				return FullHouse
			}
		case len(sameCards) == 3:
			switch cnt {
			case 1:
				continue
			case 3:
				return ThreeOfaKind
			case 2:
				return TwoPair
			}
		case len(sameCards) == 4:
			return OnePair
		case len(sameCards) == 5:
			return HighCard
		}
		//panic(fmt.Sprintf("oops: unknown card type: %+v", sameCards))
	}
	panic(fmt.Sprintf("oops: unknown card type: %+v", sameCards))
}
