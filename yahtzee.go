package main

import "fmt"

type Category int

const (
	Ones Category = iota
	Twos
	Threes
	Fours
	Fives
	Sixes
	Pair
	TwoPair
	ThreeKind
	FourKind
	SmallStraight
	LargeStraight
	FullHouse
	Yahtzee
	Chance
)

func scoreYahtzee(rolls []int, category Category) (score int) {
	// validate rolls
	err := validateRolls(rolls)
	if err != nil {
		fmt.Println("error validating rolls")
		return 0
	}

	rollCounts := make(map[int]int)
	rollSum := 0
	for _, r := range rolls {
		rollCounts[r]++
		rollSum += r
	}

	switch category {
	case Ones:
		return rollCounts[1] * 1
	case Twos:
		return rollCounts[2] * 2
	case Threes:
		return rollCounts[3] * 3
	case Fours:
		return rollCounts[4] * 4
	case Fives:
		return rollCounts[5] * 5
	case Sixes:
		return rollCounts[6] * 6
	case Pair:
		return multipleScore(rollCounts, 2)
	case TwoPair:
		return twoMultipleScore(rollCounts, 2, 2)
	case ThreeKind:
		return multipleScore(rollCounts, 3)
	case FourKind:
		return multipleScore(rollCounts, 4)
	case SmallStraight:
		return straightScore(rollCounts, true)
	case LargeStraight:
		return straightScore(rollCounts, false)
	case FullHouse:
		return twoMultipleScore(rollCounts, 2, 3)
	case Yahtzee:
		if multipleScore(rollCounts, 5) != 0 {
			return 50
		}
		return 0
	case Chance:
		return rollSum
	}

	return 0
}

func multipleScore(rollCounts map[int]int, multiple int) int {
	rollVal := 0
	for r, c := range rollCounts {
		if c >= multiple && r > rollVal {
			rollVal = r
		}
	}
	return rollVal * multiple
}

func twoMultipleScore(rollCounts map[int]int, multiple1, multiple2 int) int {
	rollVal1 := 0
	rollVal2 := 0
	for r, c := range rollCounts {
		if c >= multiple1 && r > rollVal1 {
			if rollCounts[rollVal1] >= multiple2 && rollVal1 > rollVal2 {
				rollVal2 = rollVal1
			}
			rollVal1 = r
		} else if c >= multiple2 && r > rollVal2 {
			rollVal2 = r
		}
	}
	if rollVal1 == 0 || rollVal2 == 0 {
		return 0
	}
	return rollVal1*multiple1 + rollVal2*multiple2
}

func straightScore(rollCounts map[int]int, small bool) int {
	sum := 0
	for r, c := range rollCounts {
		if c > 1 {
			return 0
		}
		sum += r
	}
	// small straight: 1, 2, 3, 4, 5
	if rollCounts[6] == 0 && small {
		return sum
	}
	// large straight: 2, 3, 4, 5, 6
	if rollCounts[1] == 0 && !small {
		return sum
	}
	return 0
}

func validateRolls(rolls []int) error {
	if len(rolls) != 5 {
		return fmt.Errorf("must have 5 rolls")
	}
	for _, r := range rolls {
		if r < 1 || r > 6 {
			return fmt.Errorf("invalid roll value")
		}
	}
	return nil
}

func main() {
	fmt.Println("1")
	rolls := []int{1, 1, 4, 3, 4}
	assertEqual(scoreYahtzee(rolls, TwoPair), 10)
	assertEqual(scoreYahtzee(rolls, Pair), 8)
	assertEqual(scoreYahtzee(rolls, ThreeKind), 0)
	assertEqual(scoreYahtzee(rolls, FourKind), 0)
	assertEqual(scoreYahtzee(rolls, Yahtzee), 0)
	assertEqual(scoreYahtzee(rolls, Chance), 13)
	assertEqual(scoreYahtzee(rolls, Ones), 2)

	fmt.Println("2")
	rolls = []int{3, 3, 2, 3, 2}
	assertEqual(scoreYahtzee(rolls, TwoPair), 10)
	assertEqual(scoreYahtzee(rolls, Pair), 6)
	assertEqual(scoreYahtzee(rolls, ThreeKind), 9)
	assertEqual(scoreYahtzee(rolls, FourKind), 0)
	assertEqual(scoreYahtzee(rolls, Yahtzee), 0)
	assertEqual(scoreYahtzee(rolls, Chance), 13)
	assertEqual(scoreYahtzee(rolls, Ones), 0)

	fmt.Println("3")
	rolls = []int{1, 2, 4, 3, 4}
	assertEqual(scoreYahtzee(rolls, TwoPair), 0)
	assertEqual(scoreYahtzee(rolls, Pair), 8)
	assertEqual(scoreYahtzee(rolls, ThreeKind), 0)
	assertEqual(scoreYahtzee(rolls, FourKind), 0)
	assertEqual(scoreYahtzee(rolls, Yahtzee), 0)
	assertEqual(scoreYahtzee(rolls, Chance), 14)

	fmt.Println("4")
	rolls = []int{1, 2, 5, 3, 4}
	assertEqual(scoreYahtzee(rolls, TwoPair), 0)
	assertEqual(scoreYahtzee(rolls, Pair), 0)
	assertEqual(scoreYahtzee(rolls, ThreeKind), 0)
	assertEqual(scoreYahtzee(rolls, FourKind), 0)
	assertEqual(scoreYahtzee(rolls, Yahtzee), 0)
	assertEqual(scoreYahtzee(rolls, Chance), 15)

	fmt.Println("5")
	rolls = []int{1, 2, 2, 2, 2}
	assertEqual(scoreYahtzee(rolls, TwoPair), 0)
	assertEqual(scoreYahtzee(rolls, Pair), 4)
	assertEqual(scoreYahtzee(rolls, ThreeKind), 6)
	assertEqual(scoreYahtzee(rolls, FourKind), 8)
	assertEqual(scoreYahtzee(rolls, Yahtzee), 0)
	assertEqual(scoreYahtzee(rolls, Chance), 9)

	fmt.Println("6")
	rolls = []int{2, 2, 2, 2, 2}
	assertEqual(scoreYahtzee(rolls, TwoPair), 0)
	assertEqual(scoreYahtzee(rolls, Pair), 4)
	assertEqual(scoreYahtzee(rolls, ThreeKind), 6)
	assertEqual(scoreYahtzee(rolls, FourKind), 8)
	assertEqual(scoreYahtzee(rolls, Yahtzee), 50)
	assertEqual(scoreYahtzee(rolls, Chance), 10)
}

func assertEqual(a, b int) {
	if a != b {
		fmt.Println("not equal!!!", a, b)
	}
}
