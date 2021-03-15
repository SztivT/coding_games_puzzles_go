package main

import (
	"container/list"
	"fmt"
)

func roundDeck(deckOne *list.List, deckTwo *list.List, roundDeckOne *list.List, roundDeckTwo *list.List, pcs int) {
	roundDeckOne.PushBack(deckOne.Front().Value.(string))
	roundDeckTwo.PushBack(deckTwo.Front().Value.(string))
	deckOne.Remove(deckOne.Front())
	deckTwo.Remove(deckTwo.Front())
	if pcs > 1 {
		roundDeck(deckOne, deckTwo, roundDeckOne, roundDeckTwo, pcs - 1)
	}
}

func cardPower(card list.Element) int {
	var face string
	face = card.Value.(string)
	switch face[0] {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case '1':
		return 10
	default:
		return int(face[0] - 48)
	}
}

func fight(cardOne list.Element, cardTwo list.Element) int {
	powerOne := cardPower(cardOne)
	powerTwo := cardPower(cardTwo)
	if powerOne == powerTwo {
		return 0
	} else if powerOne > powerTwo {
		return 1
	} else {
		return 2
	}
}

func gathering(roundDeckOne *list.List, roundDeckTwo *list.List, targetDeck *list.List) {

	for roundDeckOne.Front() != nil {
		targetDeck.PushBack(roundDeckOne.Front().Value.(string))
		roundDeckOne.Remove(roundDeckOne.Front())
	}
	for roundDeckTwo.Front() != nil {
		targetDeck.PushBack(roundDeckTwo.Front().Value.(string))
		roundDeckTwo.Remove(roundDeckTwo.Front())
	}

}

func round(deckOne *list.List, deckTwo *list.List, r *int, pat *bool) {
	roundDeckOne := list.New()
	roundDeckTwo := list.New()
	if deckOne.Front() == nil || deckTwo.Front() == nil {
		return
	}
	roundDeck(deckOne, deckTwo, roundDeckOne, roundDeckTwo, 1)
	outcome := fight(*roundDeckOne.Back(), *roundDeckTwo.Back())
	for outcome == 0 {
		if deckOne.Len() < 4 || deckTwo.Len() < 4 {
			*pat = true
			return
		}
		roundDeck(deckOne, deckTwo, roundDeckOne, roundDeckTwo, 4)
		outcome = fight(*roundDeckOne.Back(), *roundDeckTwo.Back())
	}
	*r++
	switch outcome {
	case 1:
		gathering(roundDeckOne, roundDeckTwo, deckOne)
		round(deckOne, deckTwo, r, pat)
	case 2:
		gathering(roundDeckOne, roundDeckTwo, deckTwo)
		round(deckOne, deckTwo, r, pat)
	}
}

func main() {
	var deckSizePlayerOne int
	var deckSizePlayerTwo int
	var r int
	pat := false
	deckOne := list.New()
	deckTwo := list.New()
	var card string
	_, _ = fmt.Scanf("%d", &deckSizePlayerOne)
	for i := 0; i < deckSizePlayerOne; i++ {
		_, _ = fmt.Scanf("%s", &card)
		deckOne.PushBack(card)
	}
	_, _ = fmt.Scanf("%d", &deckSizePlayerTwo)
	for i := 0; i < deckSizePlayerTwo; i++ {
		_, _ = fmt.Scanf("%s", &card)
		deckTwo.PushBack(card)
	}
	round(deckOne, deckTwo, &r, &pat)
	if pat {
		fmt.Printf("PAT\n")
		return
	} else if deckTwo.Len() == 0 {
		fmt.Printf("1 %d\n", r)
	} else if deckOne.Len() == 0 {
		fmt.Printf("2 %d", r)
	}
}
