package majong

import (
	"fmt"
)

func InputInt(hint string) int {
	fmt.Print(hint)
	var input int
	fmt.Scan(&input)
	return input
}

func InputString(hint string) string {
	fmt.Print(hint)
	var input string
	fmt.Scan(&input)
	return input
}

var debugMode = true

func Debug(format string, a ...any) {
	if debugMode {
		fmt.Println(fmt.Sprintf(format, a...))
	}
}

func InList[T int](card T, cards []T) bool {
	for _, c := range cards {
		if card == c {
			return true
		}
	}
	return false
}

func BroadCast(name, operate string, cards []Card) {
	prefix := "玩家" + name + operate
	if len(cards) == 0 {
		fmt.Println(prefix)
	} else {
		fmt.Println(prefix+" 牌型为", cards)
	}
}
