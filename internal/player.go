package internal

import (
	"fmt"
	"math/rand"
	"strings"
)

type Player struct {
	Name string
	// 获胜轮数
	WinedCnt int
	// 是否最终获胜
	Wined bool
	// 出拳次数
	GuessCnt int
	// 是否是AI
	AI bool
}

func (self *Player) Guess() string {
	self.GuessCnt++
	var ret string
	if self.AI {
		// 是AI
		ret = G.Alternatives[rand.Intn(len(G.Alternatives))]
	} else {
		// 是玩家
		var input string
		fmt.Scanf("%s\n", &input)
		ret = input
	}
	fmt.Println(self.Name, "guessed", ret)
	return strings.TrimSpace(ret)
}

func (self *Player) Win() {
	self.WinedCnt++
	if self.WinedCnt >= G.WinCnt {
		self.Wined = true
	}
}

func (self *Player) Celebrate() {
	fmt.Println("Congratulations! ", self.Name, ", u Win! loops:", G.loops)
}
