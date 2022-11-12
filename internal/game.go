package internal

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

type Game struct {
	// 备选的出拳集合
	Alternatives []string
	// 获胜的出拳组合
	Wins [][]string
	// 两名玩家
	players []*Player
	// 当前已完游戏轮数
	loops int
	// 获胜需要轮数
	WinCnt int
	// 获胜的玩家
	Winner *Player
}

func NewGame() *Game {
	// 3 轮 2 胜 玩法
	game := &Game{
		Alternatives: []string{"剪刀", "石头", "布"},
		Wins:         [][]string{{"剪刀", "布"}, {"布", "石头"}, {"石头", "剪刀"}},
		loops:        0,
		WinCnt:       2,
	}
	// test data
	playerA := &Player{Name: "maxm", WinedCnt: 0}
	playerB := &Player{Name: "AI", WinedCnt: 0, AI: true}

	game.PlayerJoin(playerA)
	game.PlayerJoin(playerB)
	return game
}

// 玩家会有1加入游戏的过程
func (self *Game) PlayerJoin(p *Player) {
	if len(self.players) >= 2 {
		fmt.Println(p.Name, "join game fail, player enough")
		return
	}
	self.players = append(self.players, p)
}

func (self *Game) Run() {
	fmt.Println("请出拳")

	go func() {
		for {
			if self.Winner != nil {
				fmt.Println("Game Over, ", self.Winner.Name, " win, exit by ctrl c")
				break
			}

			guess0 := self.players[0].Guess()
			guess1 := self.players[1].Guess()
			if !self.validate(guess0) || !self.validate(guess1) {
				fmt.Println("无效，请再次出拳")
				continue
			}

			self.loops++
			if self.loopWin([]string{guess0, guess1}) {
				self.players[0].Win()
				fmt.Println("本轮", self.players[0].Name, "胜")
				if self.players[0].Wined {
					// 0号胜出, 游戏结束
					self.Winner = self.players[0]
					break
				}
			} else if self.loopWin([]string{guess1, guess0}) {
				self.players[1].Win()
				fmt.Println("本轮", self.players[1].Name, "胜")
				if self.players[1].Wined {
					// 1号胜出, 游戏结束
					self.Winner = self.players[1]
					break
				}
			} else {
				fmt.Println("本轮平局, 游戏继续")
			}
			time.Sleep(time.Second * 1)
		}

		// 胜出者庆祝
		if self.Winner != nil {
			self.Winner.Celebrate()
			os.Exit(0)
		}

	}()

	for {
		runtime.Gosched()
	}
}

// 胜利组合是否包含该输入, loop 1轮输入
func (self *Game) loopWin(loop []string) bool {
	for _, item := range self.Wins {
		if loop[0] == item[0] && loop[1] == item[1] {
			return true
		}
	}
	return false
}

// 验证出拳是否合法
func (self *Game) validate(guess string) bool {
	for _, a := range self.Alternatives {
		if guess == a {
			return true
		}
	}
	return false
}

// func from_prompt() (string, error) {
// 	prompt := promptui.Prompt{
// 		Label: "请出拳[剪刀/石头/布]",
// 		Validate: func(s string) error {
// 			s = strings.TrimSpace(s)
// 			fmt.Println("invalid input:", s)
// 			return nil
// 		},
// 	}
// 	u, err := prompt.Run()
// 	return u, err
// }
