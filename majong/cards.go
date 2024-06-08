package majong

import (
	"math/rand"
	"time"
)

func CreateCards() (res []Card) {
	for j := 0; j < 4; j++ {
		for i := 1; i <= 10; i++ {
			for _, v := range []int{1, 2, 3} {
				res = append(res, Card(v*100+i))
			}
		}

	}
	return Shuffle(res)
}

func Shuffle(arr []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	for i := len(arr) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func RemoveCards(cards []Card, removes []Card) (res []Card) {
	res = cards
	for _, card := range removes {
		res = RemoveCard(res, card)
	}
	return
}

func RemoveCard(cards []Card, card Card) (res []Card) {
	for i, c := range cards {
		if c == card {
			res = append(res, cards[i+1:]...)
			return
		} else {
			res = append(res, c)
		}
	}
	return
}

func CanHu(cards []Card) bool {
	// 检查是否胡了
	// 四面子，一雀头
	// 先判断是否有刻子，如果没有，再判断是否有顺子
	// 面子牌仅限同色牌，所以对每种同色牌单独处理
	if len(cards) < 2 {
		return false
	}
	if len(cards) == 2 {
		return cards[0] == cards[1]
	}

	for i := 0; i < len(cards)-2; i++ {
		for j := i + 1; j < len(cards)-1; j++ {
			for k := j + 1; k < len(cards); k++ {
				if IsMianzi([]Card{cards[i], cards[j], cards[k]}) {
					if CanHu(RemoveCards(cards, []Card{cards[i], cards[j], cards[k]})) {
						return true
					}
				}
			}
		}
	}
	return false
}

func CanPong(cards []Card, card Card) bool {
	sameCards := 0
	for _, c := range cards {
		if c == card {
			sameCards++
			if sameCards >= 2 {
				return true
			}
		}
	}
	return false
}

func CanGang(cards []Card, card Card) bool {
	sameCards := 0
	for _, c := range cards {
		if c == card {
			sameCards++
			if sameCards >= 3 {
				return true
			}
		}
	}
	return false
}

func IsMianzi(cards []Card) bool {
	return IsKezi(cards) || IsShunzi(cards)
}

func IsKezi(cards []Card) bool {
	if len(cards) != 3 {
		return false
	}
	if cards[0] == cards[1] && cards[1] == cards[2] {
		return true
	}
	return false
}

func IsShunzi(cards []Card) bool {
	if len(cards) != 3 {
		return false
	}
	minCard := cards[0]
	for _, card := range cards {
		if card < minCard {
			minCard = card
		}
	}
	if cards[0]+cards[1]+cards[2] == 3*minCard+3 {
		return true
	}
	return false
}

func CheckHuType(cards []Card) string {
	if IsQingYiSe(cards) {
		return "清一色"
	}

	if IsDuiDuiHu(cards) {
		return "对对胡"
	}
	if IsQiXiaoDui(cards) {
		return "七小对"
	}
	return "平胡"

}

func IsQingYiSe(cards []Card) bool {
	if len(cards) == 0 {
		return false
	}
	cardType := cards[0] / 100
	for _, card := range cards {
		if card/100 != cardType {
			return false
		}
	}
	return true
}

func IsQiXiaoDui(cards []Card) bool {
	m := map[int]int{}
	for _, card := range cards {
		m[card]++
	}
	duizis := 0
	for _, num := range m {
		if num == 2 {
			duizis++
		}
		if num == 4 {
			duizis += 2
		}
	}
	return duizis == 7
}

func IsDuiDuiHu(cards []Card) bool {
	m := map[int]int{}
	for _, card := range cards {
		m[card]++
	}
	kezis := 0
	for _, num := range m {
		if num == 3 {
			kezis++
		}
	}
	return kezis == 4
}

func ConvertCardNames(cards []Card) (res []string) {
	for _, card := range cards {
		t0 := card % 100
		t1 := card / 100
		m0 := map[int]string{1: "一", 2: "二", 3: "三", 4: "四", 5: "五", 6: "六", 7: "七", 8: "八", 9: "九", 10: "十"}
		m1 := map[int]string{1: "万", 2: "筒", 3: "条"}
		res = append(res, m0[t0]+m1[t1])
	}
	return
}

func ConvertCardName(card Card) (res string) {
	t0 := card % 100
	t1 := card / 100
	m0 := map[int]string{1: "一", 2: "二", 3: "三", 4: "四", 5: "五", 6: "六", 7: "七", 8: "八", 9: "九", 10: "十"}
	m1 := map[int]string{1: "万", 2: "筒", 3: "条"}
	return m0[t0] + m1[t1]
}
