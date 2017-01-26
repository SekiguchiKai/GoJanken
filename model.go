package main

import (
	"math/rand"
	"time"
)

type Player struct {
	PlayerHand   string
	PlayerName   string
	PlayerResult string
}

func (p *Player) setPlayerHand(intHand int) {

	switch intHand {
	case 0:
		p.PlayerHand = "グー"
	case 1:
		p.PlayerHand = "チョキ"
	default:
		p.PlayerHand = "パー"
	}
}



// ランダムでcouの値を作成する
func DecideCpuHand() (int) {
	// 擬似乱数生成器のソースを現時刻から生成
	src := rand.NewSource(time.Now().UnixNano())

	// srcをseedとして擬似乱数生成器を作成
	ran := rand.New(src)
	// 0~2の擬似乱数を生成
	cpuHand := ran.Intn(3)

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











