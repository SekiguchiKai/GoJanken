package main

import (
	"net/http"
	"fmt"
	"os"
	"strconv"
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


// controller1
func IndexHandler(res http.ResponseWriter, req *http.Request) {
	file, err := os.Open("./resource/index.html")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	buf := make([] byte, 99999)
	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			// Readエラー処理
			fmt.Println(err)
			os.Exit(1)
			break
		}

		res.Write(buf[:n])
	}
}


// controller2
func ResultHandler(res http.ResponseWriter, req *http.Request) {
	user := new(Player)
	cpu := new(Player)

	hand := req.FormValue("hand")

	user.PlayerName = "あなた"
	cpu.PlayerName = "AI"

	intUserHand, err := strconv.Atoi(hand)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	user.setPlayerHand(intUserHand)

	intCpuHand := DecideCpuHand()
	cpu.setPlayerHand(intCpuHand)

	user.PlayerResult, cpu.PlayerResult = Judge(intUserHand, intCpuHand)

	resultView(res, req, user, cpu)

}