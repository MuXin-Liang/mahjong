package majong

import (
	"fmt"
	"sort"
)

type Player struct {
	Name  string
	Cards []Card
	Gangs []Card
	Pongs []Card
}

// PlayCard 出牌
func (p *Player) PlayCard() Card {
	var card Card
	for {
		card = InputInt(fmt.Sprintf("玩家%s出牌:", p.Name))
		if InList(card, p.Cards) {
			break
		} else {
			fmt.Println("牌", card, "不在玩家手牌中，不可出牌")
		}
	}
	p.Cards = RemoveCard(p.Cards, card)
	return card
}

// 摸牌
func (p *Player) DrawCard(card Card) {
	p.ReceiveCard([]Card{card})
	BroadCast(p.Name, fmt.Sprintf("抽牌 %d", card), p.Cards)
}

// 进入手牌
func (p *Player) ReceiveCard(cards []Card) {
	p.Cards = append(p.Cards, cards...)
	sort.Ints(p.Cards)
}

func (p *Player) CheckHu(card Card) bool {
	if b := CanHu(append(p.Cards, card)); b {
		BroadCast(p.Name, "手牌", p.Cards)
		res := InputString("玩家" + p.Name + ",你可以胡牌，请选择是否胡牌：")
		if res == "是" || res == "y" {
			p.Hu(card)
			return true
		}
	}
	return false
}

func (p *Player) CheckPong(card Card) bool {
	if CanPong(p.Cards, card) {
		BroadCast(p.Name, "手牌", p.Cards)
		res := InputString("玩家" + p.Name + ",你可以碰，请选择是否碰牌：")
		if res == "是" || res == "y" {
			p.Pong(card)
			return true
		}
	}
	return false
}

func (p *Player) CheckGang(card Card) bool {
	if CanGang(p.Cards, card) {
		BroadCast(p.Name, "手牌", p.Cards)
		res := InputString("玩家" + p.Name + ",你可以杠，请选择是否杠牌：")
		if res == "是" || res == "y" {
			p.Gang(card)
			return true
		}
	}
	return false
}

func (p *Player) Hu(card Card) {
	hupais := append(append(append(p.Cards, p.Pongs...), p.Gangs...), card)
	sort.Ints(hupais)
	BroadCast(p.Name, "胡牌 "+CheckHuType(hupais), hupais)
}

func (p *Player) Gang(card Card) {
	BroadCast(p.Name, "杠牌", nil)
	p.Gangs = append(p.Gangs, []Card{card, card, card}...)
	for i := 0; i < 3; i++ {
		p.Cards = RemoveCards(p.Cards, []Card{card, card, card})
	}
}

func (p *Player) Pong(card Card) {
	BroadCast(p.Name, "碰牌", nil)
	p.Pongs = append(p.Pongs, []Card{card, card, card}...)
	for i := 0; i < 3; i++ {
		p.Cards = RemoveCards(p.Cards, []Card{card, card})
	}
}
