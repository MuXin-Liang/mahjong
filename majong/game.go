package majong

import (
	"fmt"
	"math/rand"
)

func init() {
	ThisGame = Game{}
	ThisGame.init()
}

var ThisGame Game

type Card = int

type Game struct {
	Cards      []Card    // 剩余牌堆
	Players    []*Player // 所有玩家
	nowPlayer  int       // 当前玩家
	fromBottom bool      // 摸牌方向
}

func (g *Game) init() {
	// 初始化玩家
	g.initPlayers()

	// 每个玩家发13张牌
	sends := 13
	cards := CreateCards()
	for i := 0; i < 4; i++ {
		g.Players[i].ReceiveCard(cards[i*sends : i*sends+sends])
	}
	g.nowPlayer = rand.Intn(4)
	fmt.Println("初始玩家是", g.Players[g.nowPlayer].Name)
	g.Cards = cards[4*sends:]
}

func (g *Game) initPlayers() {
	g.Players = make([]*Player, 4)
	for i, name := range []string{"东", "南", "西", "北"} {
		g.Players[i] = &Player{
			Name: name,
		}
	}
}

type OperateType string

const (
	OpGang = "杠"
	OpPlay = "出"
	OpHu   = "胡"
)

/*
摸牌

	->胡牌
	->自杠	->摸牌
	->出牌	->其他玩家胡
		    ->其他玩家杠 ->摸牌
			->其他玩家碰 ->出牌
			->下一个玩家
*/
func (g *Game) Start() {
	for turns := 1; ; turns++ {
		if len(g.Cards) == 0 {
			fmt.Println("牌已摸尽，游戏结束")
			return
		}

		// 摸牌
		nowPlayer := g.NowPlayer()
		card := g.DrawCard()

		// 自摸
		if nowPlayer.CheckHu(card) {
			fmt.Println("游戏结束")
			return
		}

		// 自杠
		if nowPlayer.CheckGang(card) {
			g.FromBottom()
			continue
		}

		// 收入手牌
		nowPlayer.DrawCard(card)

		for {
			// 出牌
			playCard := g.NowPlayer().PlayCard()
			// 其他玩家胡
			if g.CheckHu(playCard) {
				fmt.Println("游戏结束")
				return
			}
			// 其他玩家杠
			if g.CheckGang(playCard) {
				g.FromBottom()
				break
			}
			// 其他玩家碰
			if g.CheckPong(playCard) {
				continue
			}
			// 下一个玩家
			g.NextPlayer()
			break
		}
	}
}

func (g *Game) NowPlayer() *Player {
	return g.Players[g.nowPlayer]
}

func (g *Game) NextPlayer() {
	g.nowPlayer = (g.nowPlayer + 1) % 4
}

func (g *Game) FromBottom() {
	g.fromBottom = true
}

// 检查是否有其他玩家碰
func (g *Game) CheckPong(card Card) bool {
	for i := 1; i <= 3; i++ {
		p := (g.nowPlayer + i) % 4
		player := g.Players[p]
		if player.CheckPong(card) {
			g.nowPlayer = p
			return true
		}
	}
	return false
}

// 检查是否有其他玩家杠
func (g *Game) CheckGang(card Card) bool {
	for i := 1; i <= 3; i++ {
		p := (g.nowPlayer + i) % 4
		player := g.Players[p]
		if player.CheckGang(card) {
			g.nowPlayer = p
			return true
		}
	}
	return false
}

// 检查是否有其他玩家胡
func (g *Game) CheckHu(card Card) bool {
	for i := 1; i <= 3; i++ {
		p := (g.nowPlayer + i) % 4
		player := g.Players[p]
		if player.CheckHu(card) {
			g.nowPlayer = p
			return true
		}
	}
	return false
}

// DrawCard 从牌堆抽牌
func (g *Game) DrawCard() Card {
	var card Card
	if g.fromBottom {
		card = g.Cards[len(g.Cards)-1]
		g.Cards = g.Cards[:len(g.Cards)-1]
	} else {
		card = g.Cards[0]
		g.Cards = g.Cards[1:]
	}

	// 重置
	g.fromBottom = false
	return card
}
