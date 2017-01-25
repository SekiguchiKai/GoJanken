package main

import (
	"net/http"
	"fmt"
)

// 2回目の送信
func resultView(res http.ResponseWriter, req *http.Request, user *Player, cpu *Player) {
	fmt.Fprintf(res, "あなたの打ち手は" + user.PlayerHand + "\n")
	fmt.Fprintf(res, "AIの打ち手は" + cpu.PlayerHand + "\n")
	fmt.Fprintf(res, "よって")
	fmt.Fprintf(res, "あなたの" + user.PlayerResult + "\n")
	fmt.Fprintf(res, "AIの" + cpu.PlayerResult + "\n")
}

