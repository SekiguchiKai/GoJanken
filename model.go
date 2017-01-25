package main

import (
	"math/rand"
	"time"
)


// ランダムでcouの値を作成する
func DecideCpuHand() (int) {
	// 一定のランダムではなく、毎回異なる擬似乱数を作成するためにseedを食わせる
	rand.Seed(time.Now().UnixNano())
	// 0~2にするために
	cpuHand := rand.Intn(3)

	return cpuHand
}


// 2人のプレーヤの勝負の結果を診断する
func Judge(player1Hand, player2Hand int) (string, string) {
	if player1Hand == player2Hand {
		return "引き分け", "引き分け"
	} else if (player1Hand + 1) % 3 == player2Hand {
		return "勝ち", "負け"
	}
	return "負け", "勝ち"

}











