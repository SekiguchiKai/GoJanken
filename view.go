package main

import (
	"net/http"
	"fmt"
)

// 2回目の送信
func resultView(res http.ResponseWriter, req *http.Request, user *Player, cpu *Player) {

	fmt.Fprintf(res, "<!DOCTYPE html>")
	fmt.Fprintf(res, "<html lang=\"ja\">")

	fmt.Fprintf(res, "<head>")
	fmt.Fprintf(res, "<meta charset=\"UTF-8\">")
	fmt.Fprintf(res, "<title>じゃんけん結果</title>")
	fmt.Fprintf(res, "</head>")

	fmt.Fprintf(res, "<body>")
	fmt.Fprintf(res, "あなたの打ち手は" + user.PlayerHand + "<br>")
	fmt.Fprintf(res, "AIの打ち手は" + cpu.PlayerHand + "<br>")
	fmt.Fprintf(res, "よって")
	fmt.Fprintf(res, "あなたの" + user.PlayerResult + "<br>")
	fmt.Fprintf(res, "AIの" + cpu.PlayerResult + "<br>")
	fmt.Fprintf(res, "<a href =\" / \">戻る<a>")

	fmt.Fprintf(res, "</body>")
	fmt.Fprintf(res, "</html>")

}

