package main

import (
	"math/rand"
	"runtime"
	"time"

	"github.com/xm-tech/go-guess-game/internal"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	rand.Seed(time.Now().UnixNano())

	internal.G = internal.NewGame()
	internal.G.Run()
}
